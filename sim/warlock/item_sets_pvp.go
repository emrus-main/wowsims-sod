package warlock

import (
	"time"

	"github.com/wowsims/sod/sim/core"
	"github.com/wowsims/sod/sim/core/stats"
)

///////////////////////////////////////////////////////////////////////////
//                            SoD Phase 4 Item Sets
///////////////////////////////////////////////////////////////////////////

var ItemSetChampionsThreads = core.NewItemSet(core.ItemSet{
	Name: "Champion's Threads",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Reduces the casting time of your Immolate spell by 0.2 sec.
		4: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.GetOrRegisterAura(core.Aura{
				Label:    "Champion's Threads Immolate Cast Time Reduction",
				ActionID: core.ActionID{SpellID: 23047},
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Immolate {
						if spell != nil {
							spell.DefaultCast.CastTime -= time.Millisecond * 200
						}
					}
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Immolate {
						if spell != nil {
							spell.DefaultCast.CastTime += time.Millisecond * 200
						}
					}
				},
			})
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})

var ItemSetLieutenantCommandersThreads = core.NewItemSet(core.ItemSet{
	Name: "Lieutenant Commander's Threads",
	Bonuses: map[int32]core.ApplyEffect{
		// Increases damage and healing done by magical spells and effects by up to 23.
		2: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.SpellPower, 23)
		},
		// Reduces the casting time of your Immolate spell by 0.2 sec.
		4: func(agent core.Agent) {
			warlock := agent.(WarlockAgent).GetWarlock()
			warlock.GetOrRegisterAura(core.Aura{
				Label:    "Lieutenant Commander's Threads Immolate Cast Time Reduction",
				ActionID: core.ActionID{SpellID: 23047},
				Duration: core.NeverExpires,
				OnReset: func(aura *core.Aura, sim *core.Simulation) {
					aura.Activate(sim)
				},
				OnGain: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Immolate {
						if spell != nil {
							spell.DefaultCast.CastTime -= time.Millisecond * 200
						}
					}
				},
				OnExpire: func(aura *core.Aura, sim *core.Simulation) {
					for _, spell := range warlock.Immolate {
						if spell != nil {
							spell.DefaultCast.CastTime += time.Millisecond * 200
						}
					}
				},
			})
		},
		// +20 Stamina.
		6: func(agent core.Agent) {
			c := agent.GetCharacter()
			c.AddStat(stats.Stamina, 20)
		},
	},
})
