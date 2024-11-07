package character

import (
	"reflect"
	"testing"
)

func TestGetRaceData(t *testing.T) {
	tests := []struct {
		raceName      string
		expectedData  Race
		expectedError bool
	}{
		{
			raceName: "Dwarf",
			expectedData: Race{
				Name:  "Dwarf",
				Speed: 25,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Constitution", 2},
				},
				Proficiencies: []string{
					"Battle Axe", "Handaxe", "Light Hammer", "Warhammer",
				},
				Resistances: []string{
					"Poison",
				},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{
					"Hill Dwarf", "Mountain Dwarf",
				},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Darkvision", "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."},
					{"Dwarven Resilience", "You have advantage on saving throws against poison, and resistance against poison damage."},
					{"Stonecunning", "Whenever you make an Intelligence (History) check related to the origin of stonework, you are considered proficient in the History skill and add double your proficiency bonus to the check."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Elf",
			expectedData: Race{
				Name:  "Elf",
				Speed: 30,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Dexterity", 2},
				},
				Proficiencies: []string{
					"Perception",
				},
				Resistances: []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{
					"High Elf", "Wood Elf", "Dark Elf",
				},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Darkvision", "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."},
					{"Fey Ancestry", "You have advantage on saving throws against being charmed, and magic can't put you to sleep."},
					{"Trance", "You don't need to sleep. Instead you meditate deeply for 4 hours each day."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Halfling",
			expectedData: Race{
				Name:  "Halfling",
				Speed: 25,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Dexterity", 2},
				},
				Proficiencies: []string{},
				Resistances:   []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{
					"Lightfoot Halfling", "Stout Halfling",
				},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Lucky", "When you roll a 1 on the d20 for an attack roll, ability check, or saving throw, you can reroll the die and must use the new roll."},
					{"Brave", "You have advantage on saving throws against being frightened."},
					{"Halfling Nimbleness", "You can move through the space of any creature that is of a size larger than yours."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Human",
			expectedData: Race{
				Name:  "Human",
				Speed: 30,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Strength", 1},
					{"Dexterity", 1},
					{"Constitution", 1},
					{"Intelligence", 1},
					{"Wisdom", 1},
					{"Charisma", 1},
				},
				Proficiencies: []string{},
				Resistances:   []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{},
			},
			expectedError: false,
		},
		{
			raceName: "Dragonborn",
			expectedData: Race{
				Name:  "Dragonborn",
				Speed: 30,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Strength", 2},
					{"Charisma", 1},
				},
				Proficiencies: []string{},
				Resistances:   []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{
					"Black", "Blue", "Brass", "Bronze", "Copper", "Gold", "Green", "Red", "Silver", "White",
				},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Draconic Ancestry", "You have draconic ancestry. Your ancestry determines your breath weapon and damage resistance."},
					{"Breath Weapon", "You can use your action to exhale destructive energy. Your draconic ancestry determines the size, shape and damade type of exhalation. Each creature in it's path makes saving throw, the type of which is determined by your draconic ancestry."},
					{"Damage Resistance", "You have resistance to the damage type associated with your draconic ancestry."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Gnome",
			expectedData: Race{
				Name:  "Gnome",
				Speed: 25,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Dexterity", 2},
				},
				Proficiencies: []string{},
				Resistances:   []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{
					"Forest Gnome", "Rock Gnome",
				},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Darkvision", "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."},
					{"Gnome Cunning", "You have advantage on all Intelligence, Wisdom, and Charisma saving throws against magic."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Half-Elf",
			expectedData: Race{
				Name:  "Half-Elf",
				Speed: 30,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Charisma", 2},
				},
				Proficiencies: []string{},
				Resistances:   []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{
					{
						Choices: []string{
							"Athletics",
							"Acrobatics",
							"Sleight of Hand",
							"Stealth",
							"Arcana",
							"History",
							"Investigation",
							"Nature",
							"Religion",
							"Animal Handling",
							"Insight",
							"Medicine",
							"Perception",
							"Survival",
							"Deception",
							"Intimidation",
							"Performance",
							"Persuasion",
						},
						ChoiceNumber: 2,
					},
				},
				Subraces: []string{},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Darkvision", "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."},
					{"Fey Ancestry", "You have advantage on saving throws against being charmed, and magic can't put you to sleep."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Half-Orc",
			expectedData: Race{
				Name:  "Half-Orc",
				Speed: 30,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Strength", 2},
					{"Constitution", 1},
				},
				Proficiencies: []string{},
				Resistances:   []string{},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Darkvision", "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."},
					{"Relentless Endurance", "When you are reduced to 0 hit points but not killed outright, you can drop to 1 hit point instead. You can't use this feature again until you finish a long rest."},
					{"Savage Attacks", "When you score a critical hit with a melee weapon attack, you can roll one of the weapon's damage dice one additional time and add it to the extra damage of the critical hit."},
				},
			},
			expectedError: false,
		},
		{
			raceName: "Tiefling",
			expectedData: Race{
				Name:  "Tiefling",
				Speed: 30,
				StatBonus: []struct {
					Stat  string
					Bonus int
				}{
					{"Intelligence", 1},
					{"Charisma", 2},
				},
				Proficiencies: []string{},
				Resistances: []string{
					"Fire",
				},
				ProfChoices: []struct {
					Choices      []string
					ChoiceNumber int
				}{},
				Subraces: []string{},
				RaceTraits: []struct {
					Trait       string
					Explanation string
				}{
					{"Darkvision", "You can see in dim light within 60 feet of you as if it were bright light, and in darkness as if it were dim light. You can't discern color in darkness, only shades of gray."},
					{"Hellish Resistance", "You have resistance to fire damage."},
					{"Infernal Legacy", "You know the thaumaturgy cantrip."},
				},
			},
			expectedError: false,
		},
		{
			raceName:      "Blurg",
			expectedData:  Race{},
			expectedError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.raceName, func(t *testing.T) {
			raceData, err := GetRaceData(test.raceName)
			if (err != nil) != test.expectedError {
				t.Errorf("unexpected error state: got %v, want %v", err != nil, test.expectedError)
			}
			if !reflect.DeepEqual(raceData, test.expectedData) {
				t.Errorf("expected: %+v, got: %+v", test.expectedData, raceData)
			}
		})
	}
}
