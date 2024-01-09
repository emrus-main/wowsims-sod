package druid

import (
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/stats"
	"github.com/wowsims/sod/sim/core/proto"
)

func (druid *Druid) getSavageRoarMultiplier() float64 {
	return 1.3
}

func (druid *Druid) applySavageRoar() {
	if !druid.HasRune(proto.DruidRune_RuneLegsSavageRoar) {
		return
	}

	actionID := core.ActionID{SpellID: 407988}

	srm := druid.getSavageRoarMultiplier()

	druid.SavageRoarDurationTable = [6]time.Duration{
		0,
		time.Second*(9+5),
		time.Second*(9+10),
		time.Second*(9+15),
		time.Second*(9+20),
		time.Second*(9+25),
	}

	druid.SavageRoarAura = druid.RegisterAura(core.Aura{
		Label:    "Savage Roar Aura",
		ActionID: actionID,
		Duration: 9,
		OnGain: func(aura *core.Aura, sim *core.Simulation) {
			druid.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexPhysical] *= srm
		},
		OnExpire: func(aura *core.Aura, sim *core.Simulation) {
			if druid.InForm(Cat) {
				druid.PseudoStats.SchoolDamageDealtMultiplier[stats.SchoolIndexPhysical] /= srm
			}
		},
	})

	srSpell := druid.RegisterSpell(Cat, core.SpellConfig{
		ActionID: actionID,
		Flags:    core.SpellFlagAPL,
		EnergyCost: core.EnergyCostOptions{
			Cost: 25,
		},
		Cast: core.CastConfig{
			DefaultCast: core.Cast{
				GCD: time.Second,
			},
			IgnoreHaste: true,
		},
		ExtraCastCondition: func(sim *core.Simulation, target *core.Unit) bool {
			return druid.ComboPoints() > 0
		},

		ApplyEffects: func(sim *core.Simulation, _ *core.Unit, spell *core.Spell) {
			druid.SavageRoarAura.Duration = druid.SavageRoarDurationTable[druid.ComboPoints()]
			druid.SavageRoarAura.Activate(sim)
			druid.SpendComboPoints(sim, spell.ComboPointMetrics())
		},
	})

	druid.SavageRoar = srSpell
}

func (druid *Druid) CurrentSavageRoarCost() float64 {
	return druid.SavageRoar.ApplyCostModifiers(druid.SavageRoar.DefaultCast.Cost)
}

