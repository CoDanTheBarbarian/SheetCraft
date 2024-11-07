package character

import (
	"reflect"
	"testing"
)

func TestGetSubraceData(t *testing.T) {
	tests := []struct {
		name          string
		expectedData  SubRace //injected from the map in subraces file
		expectedError bool
	}{
		{
			name: "Hill Dwarf",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "Mountain Dwarf",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "High Elf",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "Wood Elf",
			expectedData: SubRace{
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
				SpeedBonus:  0,
			},
			expectedError: false,
		},
		{
			name: "Dark Elf (Drow)",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "Lightfoot Halfling",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "Stout Halfling",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "Forest Gnome",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name: "Rock Gnome",
			expectedData: SubRace{
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
			expectedError: false,
		},
		{
			name:          "Glock Gnome",
			expectedData:  SubRace{},
			expectedError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			subraceData, err := GetSubraceData(test.name)
			if (err != nil) != test.expectedError {
				t.Errorf("unexpected error state: got %v, want %v", err != nil, test.expectedError)
			}
			if !reflect.DeepEqual(subraceData, test.expectedData) {
				t.Errorf("expected: %+v, got %+v", test.expectedData, subraceData)
			}
		})
	}
}

func TestGetDragonColorData(t *testing.T) {
	tests := []struct {
		name          string
		expectedData  DragonColor //injected from the map in subraces file
		expectedError bool
	}{
		{
			name: "Black Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Blue Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Brass Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Bronze Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Copper Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Gold Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Green Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Red Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "Silver Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
		{
			name: "White Dragonborn",
			expectedData: DragonColor{
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
			expectedError: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dragonData, err := GetDragonColorData(test.name)
			if (err != nil) != test.expectedError {
				t.Errorf("unexpected error state: got %v, want %v", err != nil, test.expectedError)
			}
			if !reflect.DeepEqual(dragonData, test.expectedData) {
				t.Errorf("expected: %+v, got %+v", test.expectedData, dragonData)
			}
		})
	}
}
