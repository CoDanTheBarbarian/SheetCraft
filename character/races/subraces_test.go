package character

import (
	"reflect"
	"testing"
)

func TestGetSubraceData(t *testing.T) {
	tests := []struct {
		name          string
		expectedData  SubRace
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
