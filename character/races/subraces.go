package character

import (
	"fmt"
)

type SubRace struct {
	Name      string
	StatBonus []struct {
		Stat  string
		Bonus int
	}
	Proficiencies []string
	Resistances   []string
	HpBonus       int
	SpeedBonus    int
}

type DragonColor struct {
	Name      string
	StatBonus []struct {
		Stat  string
		Bonus int
	}
	Proficiencies []string
	Resistances   []string
	BreathWeapon  struct {
		BreathShape string
		BreathSize  []int
		BreathType  string
	}
}

var SubRaces = map[string]SubRace{
	"Hill Dwarf": {
		Name: "Hill Dwarf",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Wisdom", 1},
		},
		Proficiencies: []string{},
		Resistances:   []string{},
		HpBonus:       1,
		SpeedBonus:    0,
	},
	"Mountain Dwarf": {
		Name: "Mountain Dwarf",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Strength", 2},
		},
		Proficiencies: []string{
			"Light Armor", "Medium Armor",
		},
		Resistances: []string{},
		HpBonus:     0,
		SpeedBonus:  0,
	},
	"High Elf": {
		Name: "High Elf",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Intelligence", 1},
		},
		Proficiencies: []string{
			"Long Sword", "Short Sword", "Long Bow", "Short Bow",
		},
		Resistances: []string{},
		HpBonus:     0,
		SpeedBonus:  0,
	},
	"Wood Elf": {
		Name: "Wood Elf",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Wisdom", 1},
		},
		Proficiencies: []string{
			"Long Sword", "Short Sword", "Long Bow", "Short Bow",
		},
		Resistances: []string{},
		HpBonus:     0,
		SpeedBonus:  5,
	},
	"Dark Elf (Drow)": {
		Name: "Dark Elf (Drow)",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Charisma", 1},
		},
		Proficiencies: []string{
			"Rapier", "Short Sword", "Hand Crossbow",
		},
		Resistances: []string{},
		HpBonus:     0,
		SpeedBonus:  0,
	},
	"Lightfoot Halfling": {
		Name: "Lightfoot Halfling",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Charisma", 1},
		},
		Proficiencies: []string{},
		Resistances:   []string{},
		HpBonus:       0,
		SpeedBonus:    0,
	},
	"Stout Halfling": {
		Name: "Stout Halfling",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Charisma", 1},
		},
		Proficiencies: []string{},
		Resistances:   []string{"Poison"},
		HpBonus:       0,
		SpeedBonus:    0,
	},
	"Forest Gnome": {
		Name: "Forest Gnome",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Dexterity", 1},
		},
		Proficiencies: []string{},
		Resistances:   []string{},
		HpBonus:       0,
		SpeedBonus:    0,
	},
	"Rock Gnome": {
		Name: "Rock Gnome",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{
			{"Constitution", 1},
		},
		Proficiencies: []string{},
		Resistances:   []string{},
		HpBonus:       0,
		SpeedBonus:    0,
	},
}

func GetSubraceData(name string) (SubRace, error) {
	if subraceData, ok := SubRaces[name]; ok {
		return subraceData, nil
	} else {
		return SubRace{}, fmt.Errorf("could not get subrace data for %s", name)
	}
}

var DragonColors = map[string]DragonColor{
	"Black Dragonborn": {
		Name: "Black Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"acid"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Line",
			BreathSize:  []int{5, 30},
			BreathType:  "Acid",
		},
	},
	"Blue Dragonborn": {
		Name: "Blue Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Lightning"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Line",
			BreathSize:  []int{5, 30},
			BreathType:  "Lightning",
		},
	},
	"Brass Dragonborn": {
		Name: "Brass Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Fire"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Line",
			BreathSize:  []int{5, 30},
			BreathType:  "Fire",
		},
	},
	"Bronze Dragonborn": {
		Name: "Bronze Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Lightning"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Line",
			BreathSize:  []int{5, 30},
			BreathType:  "Lightning",
		},
	},
	"Copper Dragonborn": {
		Name: "Copper Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Acid"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Line",
			BreathSize:  []int{5, 30},
			BreathType:  "Acid",
		},
	},
	"Gold Dragonborn": {
		Name: "Gold Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Fire"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Cone",
			BreathSize:  []int{15},
			BreathType:  "Fire",
		},
	},
	"Green Dragonborn": {
		Name: "Green Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Poison"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Cone",
			BreathSize:  []int{15},
			BreathType:  "Poison",
		},
	},
	"Red Dragonborn": {
		Name: "Red Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Fire"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Cone",
			BreathSize:  []int{15},
			BreathType:  "Fire",
		},
	},
	"Silver Dragonborn": {
		Name: "Silver Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Cold"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Cone",
			BreathSize:  []int{15},
			BreathType:  "Cold",
		},
	},
	"White Dragonborn": {
		Name: "White Dragonborn",
		StatBonus: []struct {
			Stat  string
			Bonus int
		}{}, // No Stat bonuses
		Proficiencies: []string{}, // No proficiencies
		Resistances:   []string{"Cold"},
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "Cone",
			BreathSize:  []int{15},
			BreathType:  "Cold",
		},
	},
}

func GetDragonColorData(name string) (DragonColor, error) {
	if colorData, ok := DragonColors[name]; ok {
		return colorData, nil
	}
	return DragonColor{}, fmt.Errorf("could not get dragon color data for %s", name)
}
