package items

import (
	"reflect"
	"testing"
)

func TestGetWeaponData(t *testing.T) {
	tests := []struct {
		name          string
		expectedData  Weapon
		expectedError bool
	}{
		{
			name: "Club",
			expectedData: Weapon{
				Name:               "Club",
				WeaponType:         "Simple",
				DamageDie:          []int{4},
				DamageType:         "Bludgeoning",
				CoreAttribute:      []string{"Strength"},
				Properties:         []string{"Light"},
				Range:              []int{},
				VersatileDamageDie: 0,
			},
			expectedError: false,
		},
		{
			name: "Battle Axe",
			expectedData: Weapon{
				Name:               "Battle Axe",
				WeaponType:         "Martial",
				DamageDie:          []int{8},
				DamageType:         "Slashing",
				CoreAttribute:      []string{"Strength"},
				Properties:         []string{"Versatile"},
				Range:              []int{},
				VersatileDamageDie: 10,
			},
			expectedError: false,
		},
		{
			name: "Javelin",
			expectedData: Weapon{
				Name:          "Javelin",
				WeaponType:    "Simple",
				DamageDie:     []int{6},
				DamageType:    "Piercing",
				CoreAttribute: []string{"Strength"},
				Properties:    []string{"Thrown"},
				Range:         []int{20, 60},
			},
			expectedError: false,
		},
		{
			name: "Trident",
			expectedData: Weapon{
				Name:               "Trident",
				WeaponType:         "Martial",
				DamageDie:          []int{6},
				DamageType:         "Piercing",
				CoreAttribute:      []string{"Strength"},
				Properties:         []string{"Thrown", "Versatile"},
				Range:              []int{20, 60},
				VersatileDamageDie: 8,
			},
			expectedError: false,
		},
		{
			name:          "Grenade",
			expectedData:  Weapon{},
			expectedError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wData, err := GetWeaponData(test.name)
			if (err != nil) != test.expectedError {
				t.Errorf("unexpected error state: got %v, want %v", err != nil, test.expectedError)
			}
			if !reflect.DeepEqual(wData, test.expectedData) {
				t.Errorf("expected: %+v, got %+v", test.expectedData, wData)
			}
		})
	}
}
