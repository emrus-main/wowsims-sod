package database

import (
	"regexp"

	"github.com/wowsims/sod/sim/core/proto"
	"github.com/wowsims/sod/sim/core/stats"
)

var OtherItemIdsToFetch = []string{}

var ItemOverrides = []*proto.UIItem{
	// Valentine's day event rewards
	// {Id: 51804, Phase: 2},

	// SOD Items
	{Id: 211848, Name: "Blackfathom Mana Oil", Icon: "inv_potion_99", Stats: stats.Stats{stats.MP5: 12, stats.SpellHit: 2}.ToFloatArray()},
	{Id: 211845, Name: "Blackfathom Sharpening Stone", Icon: "inv_misc_rune_04", Stats: stats.Stats{stats.MeleeHit: 2}.ToFloatArray()},

	{Id: 10019, Sources: []*proto.UIItemSource{{
		Source: &proto.UIItemSource_Crafted{
			Crafted: &proto.CraftedSource{
				Profession: proto.Profession_Tailoring, SpellId: 3759,
			},
		},
	}}},

	// Updated profession items not updated in the AtlasLoot DB
	// Crimson Silk Robe
	{Id: 217245, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439085}}}}},
	// Black Mageweave Vest
	{Id: 217246, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439086}}}}},
	// Long Silken Cloak
	{Id: 217252, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439094}}}}},
	// Enchanter's Cowl
	{Id: 217257, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Tailoring, SpellId: 439102}}}}},
	// Big Voodoo Mask
	{Id: 217259, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439105}}}}},
	// Big Voodoo Robe
	{Id: 217261, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439108}}}}},
	// Turtle Scale Breastplate
	{Id: 217268, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439116}}}}},
	// Turtle Scale Gloves
	{Id: 217270, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Leatherworking, SpellId: 439118}}}}},
	// Golden Scale Cuirass
	{Id: 217277, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Blacksmithing, SpellId: 439124}}}}},
	// Golden Scale Coif
	{Id: 217279, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Blacksmithing, SpellId: 439126}}}}},
	// Golden Scale Leggings
	{Id: 217285, Sources: []*proto.UIItemSource{{Source: &proto.UIItemSource_Crafted{Crafted: &proto.CraftedSource{Profession: proto.Profession_Blacksmithing, SpellId: 439132}}}}},

	// The item tooltip is missing the usual Libram tag
	{Id: 221457, RangedWeaponType: proto.RangedWeaponType_RangedWeaponTypeLibram},

	// The item tooltip is missing the usual Totem tag
	{Id: 221464, RangedWeaponType: proto.RangedWeaponType_RangedWeaponTypeTotem},

	// SoD Gnomeregan Quest Necklaces are missing quest info from the gear planner DB
	{Id: 213343, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},
	{Id: 213344, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},
	{Id: 213345, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},
	{Id: 213346, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80324, Name: "The Mad King"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 80325, Name: "The Mad King"}}},
	}},

	// SoD Sunken Temple Drakeclaw Bands are missing quest info from the gear planner DB
	{Id: 220626, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220627, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220628, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220629, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},
	{Id: 220630, Sources: []*proto.UIItemSource{
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82081, Name: "A Broken Ritual"}}},
		{Source: &proto.UIItemSource_Quest{Quest: &proto.QuestSource{Id: 82083, Name: "A Broken Ritual"}}},
	}},

	// Heirloom Dwarven Handcannon, Wowhead partially glitchs out and shows us some other lvl calc for this
	// {Id: 44093, Stats: stats.Stats{stats.MeleeCrit: 30, stats.SpellCrit: 30, stats.Resilience: 13, stats.AttackPower: 34}.ToFloatArray()},
}

// Keep these sorted by item ID.
var ItemAllowList = map[int32]struct{}{
	11815: {}, // Hand of Justice
	12590: {}, // Felstriker
	15808: {}, // Fine Light Crossbow (for hunter testing).
	18843: {},
	18844: {},
	18847: {},
	18848: {},
	19019: {}, // Thunderfury
	19808: {}, // Rockhide Strongfish
	20837: {}, // Sunstrider Axe
	20966: {}, // Jade Pendant of Blasting
	21625: {}, // Scarab Brooch
	21685: {}, // Petrified Scarab
	24114: {}, // Braided Eternium Chain
	28572: {}, // Blade of the Unrequited
	28830: {}, // Dragonspine Trophy
	29383: {}, // Bloodlust Brooch
	29387: {}, // Gnomeregan Auto-Blocker 600
	29994: {}, // Thalassian Wildercloak
	29996: {}, // Rod of the Sun King
	30032: {}, // Red Belt of Battle
	30627: {}, // Tsunami Talisman
	30720: {}, // Serpent-Coil Braid
	31193: {}, // Blade of Unquenched Thirst
	32387: {}, // Idol of the Raven Goddess
	32658: {}, // Badge of Tenacity
	33135: {}, // Falling Star
	33140: {}, // Blood of Amber
	33143: {}, // Stone of Blades
	33144: {}, // Facet of Eternity
	33504: {}, // Libram of Divine Purpose
	33506: {}, // Skycall Totem
	33507: {}, // Stonebreaker's Totem
	33508: {}, // Idol of Budding Life
	33510: {}, // Unseen moon idol
	33829: {}, // Hex Shrunken Head
	33831: {}, // Berserkers Call
	34472: {}, // Shard of Contempt
	34473: {}, // Commendation of Kael'thas
	37032: {}, // Edge of the Tuskarr
	37574: {}, // Libram of Furious Blows
	38072: {}, // Thunder Capacitor
	38287: {}, // Empty Mug of Direbrew
	38289: {}, // Coren's Lucky Coin
	39208: {}, // Sigil of the Dark Rider
	41752: {}, // Brunnhildar Axe
	6360:  {}, // Steelscale Crushfish
	8345:  {}, // Wolfshead Helm
	9449:  {}, // Manual Crowd Pummeler

	// SOD
	211848: {},
	211845: {},
	215111: {}, // Gneuro-Linked Arcano-Filament Monocle
	215114: {}, // Glowing Hyperconductive Scale Coif
}

// Keep these sorted by item ID.
var ItemDenyList = map[int32]struct{}{
	9653:   {}, // Speedy Racer Goggles
	12104:  {}, // Brindlethorn Tunic
	12805:  {}, // Orb of Fire
	17782:  {}, // talisman of the binding shard
	17783:  {}, // talisman of the binding fragment
	17802:  {}, // Deprecated version of Thunderfury
	20522:  {}, // Feral Staff
	33350:  {},
	34576:  {}, // Battlemaster's Cruelty
	34577:  {}, // Battlemaster's Depreavity
	34578:  {}, // Battlemaster's Determination
	34579:  {}, // Battlemaster's Audacity
	34580:  {}, // Battlemaster's Perseverence
	206382: {}, // Tempest Icon
	206387: {}, // Kajaric Icon
	206954: {}, // Idol of Ursine Rage
	208689: {}, // Ferocious Idol
	208849: {}, // Libram of Blessings
	208851: {}, // Libram of Justice
	210195: {}, // Unbalanced Idol
	210534: {}, // Idol of the Wild
	215116: {}, // UNUSED - Hyperconductive Speed Belt
	211472: {}, // Libram of Banishment
	213513: {}, // Libram of Deliverance
	213594: {}, // Idol of the Heckler
	220915: {}, // Idol of the Raging Shambler
}

// Item icons to include in the DB, so they don't need to be separately loaded in the UI.
var ExtraItemIcons = []int32{
	// Demonic Rune
	12662,

	// Explosives
	13180,
	11566,
	8956,
	10646,
	18641,
	15993,
	16040,

	// Food IDs
	13928,
	20452,
	13931,
	18254,
	21023,
	13813,
	13810,

	// Flask IDs
	13510,
	13511,
	13512,
	13513,

	// Zanza
	20079,

	// Blasted Lands
	8412,
	8423,
	8424,
	8411,

	// Agility Elixer IDs
	13452,
	9187,

	// Single Elixirs
	20007, // Mana Regen Elixir
	20004, // Major Troll's Blood Potion
	9088,  // Gift of Arthas

	// Armor Elixirs
	3389,  // Defense
	8951,  // Greater
	13445, // Superior Defense

	// Health Elixirs
	2458, // Minor Fortitude
	3825, // Fortitude

	// Strength
	12451,
	9206,

	// AP
	12460,
	12820,

	// Random
	5206, // Bogling Root

	// SP
	13454,
	9264,
	21546,
	17708,

	// Crystal
	11564, // Armor

	// Alcohol Buff
	18284,
	18269,
	20709,
	21114,
	21151,

	// Potions / In Battle Consumes
	13444,

	// Thistle Tea
	7676,

	// Weapon Oils
	20748,
	20749,
	12404,
	18262,
}

var SpellIconoverrides = []*proto.IconData{
	{Id: 415068, Name: "Exorcism (Rank 1)"},
	{Id: 415069, Name: "Exorcism (Rank 2)"},
	{Id: 415070, Name: "Exorcism (Rank 3)"},
	{Id: 415071, Name: "Exorcism (Rank 4)"},
	{Id: 415072, Name: "Exorcism (Rank 5)"},
	{Id: 415073, Name: "Exorcism (Rank 6)"},
}

// Raid buffs / debuffs
var SharedSpellsIcons = []int32{
	// World Buffs
	22888, // Ony / Nef
	24425, // Spirit
	16609, // Warchief
	23768, // DMF Damage
	23736, // DMF Agi
	23766, // DMF Int
	23738, // DMF Spirit
	23737, // DMF Stam

	22818, // DM Stam
	22820, // DM Spell Crit
	22817, // DM AP

	15366, // Songflower

	29534, // Silithus

	18264, // Headmasters

	// Registered CD's
	10060, // Power Infusion
	29166, // Innervate

	// Mark
	1126,
	5232,
	6756,
	5234,
	8907,
	9884,
	9885,
	17055,

	20217, // Kings (Talent)
	25898, // Greater Kings
	25899, // Sanctuary

	10293, // Devo Aura
	20142, // Imp. Devo

	// Stoneskin Totem
	10408,
	16293,

	// Fort
	1243,
	1244,
	1245,
	2791,
	10937,
	10938,
	14767,

	// Spirit
	14752,
	14818,
	14819,
	27841,

	// Might
	19740,
	19834,
	19835,
	19836,
	19837,
	19838,
	25291,
	20048,

	// Commanding Shout
	6673,
	5242,
	6192,
	11549,
	11550,
	11551,
	25289,
	12861,

	// AP
	30811, // Unleashed Rage
	19506, // Trueshot

	// Battle Shout
	6673,
	5242,
	6192,
	11549,
	11550,
	11551,
	25289,
	12861, // Imp

	// Wisdom
	19742,
	19850,
	19852,
	19853,
	19854,
	25290,
	20245,

	// Mana Spring
	5675,
	10495,
	10496,
	10497,

	17007, // Leader of the Pack
	24858, // Moonkin

	// Windfury
	8512,
	10613,
	10614,
	29193, // Imp WF

	// Raid Debuffs
	8647,
	7386,
	7405,
	8380,
	11596,
	11597,

	770,
	778,
	9749,
	9907,
	11708,
	18181,

	26016,
	12879,
	9452,
	26021,
	16862,
	9747,
	9898,

	3043,
	14275,
	14276,
	14277,

	17800,
	17803,
	12873,
	28593,

	11374,
	15235,

	24977,

	// Runes
	// Mage
	401729,
	401731,
	401732,
	401734,
	401741,
	401743,
	401744,
	412325,
	412326,
	415467,
	425168,
	425169,
	401556,
	400574,
	400573,

	// Gnomeregan on-use item effects
	437327,
	437362,
}

// If any of these match the item name, don't include it.
var DenyListNameRegexes = []*regexp.Regexp{
	regexp.MustCompile(`30 Epic`),
	regexp.MustCompile(`63 Blue`),
	regexp.MustCompile(`63 Green`),
	regexp.MustCompile(`66 Epic`),
	regexp.MustCompile(`90 Epic`),
	regexp.MustCompile(`90 Green`),
	regexp.MustCompile(`Boots 1`),
	regexp.MustCompile(`Boots 2`),
	regexp.MustCompile(`Boots 3`),
	regexp.MustCompile(`Bracer 1`),
	regexp.MustCompile(`Bracer 2`),
	regexp.MustCompile(`Bracer 3`),
	regexp.MustCompile(`DB\d`),
	regexp.MustCompile(`DEPRECATED`),
	regexp.MustCompile(`Deprecated: Keanna`),
	regexp.MustCompile(`Indalamar`),
	regexp.MustCompile(`Monster -`),
	regexp.MustCompile(`NEW`),
	regexp.MustCompile(`PH`),
	regexp.MustCompile(`QR XXXX`),
	regexp.MustCompile(`TEST`),
	regexp.MustCompile(`Test`),
	regexp.MustCompile(`zOLD`),
}
