package items

import "fmt"

type Weapon struct {
	Name               string
	WeaponType         string
	DamageDie          []int
	DamageType         string
	CoreAttribute      []string
	Properties         []string
	Range              []int
	VersatileDamageDie int
}

type Armor struct {
	Name           string
	ArmorType      string
	AC             int
	DexMod         bool
	DexMax         int
	StealthPenalty bool
	StrengthMin    int
	Properties     []string
}

type Shield struct {
	Name       string
	ArmorType  string
	ACBonus    int
	Properties []string
}

var Weapons = map[string]Weapon{
	// Melee Weapons
	"Club": {
		Name:               "Club",
		WeaponType:         "Simple",
		DamageDie:          []int{4},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Light"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Greatclub": {
		Name:               "Greatclub",
		WeaponType:         "Simple",
		DamageDie:          []int{8},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Mace": {
		Name:               "Mace",
		WeaponType:         "Simple",
		DamageDie:          []int{6},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Sickle": {
		Name:               "Sickle",
		WeaponType:         "Simple",
		DamageDie:          []int{4},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Light"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Flail": {
		Name:               "Flail",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Glaive": {
		Name:               "Glaive",
		WeaponType:         "Martial",
		DamageDie:          []int{10},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Heavy", "Reach", "Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Great Axe": {
		Name:               "Great Axe",
		WeaponType:         "Martial",
		DamageDie:          []int{12},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Heavy", "Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Great Sword": {
		Name:               "Great Sword",
		WeaponType:         "Martial",
		DamageDie:          []int{6, 6},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Heavy", "Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Halberd": {
		Name:               "Halberd",
		WeaponType:         "Martial",
		DamageDie:          []int{10},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Heavy", "Reach", "Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Lance": {
		Name:               "Lance",
		WeaponType:         "Martial",
		DamageDie:          []int{12},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Reach"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Maul": {
		Name:               "Maul",
		WeaponType:         "Martial",
		DamageDie:          []int{6, 6},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Heavy", "Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Morning Star": {
		Name:               "Morning Star",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Pike": {
		Name:               "Pike",
		WeaponType:         "Martial",
		DamageDie:          []int{10},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Heavy", "Reach", "Two Handed"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Rapier": {
		Name:               "Rapier",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength", "Dexterity"},
		Properties:         []string{"Finesse"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Scimitar": {
		Name:               "Scimitar",
		WeaponType:         "Martial",
		DamageDie:          []int{6},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength", "Dexterity"},
		Properties:         []string{"Finesse", "Light"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Short Sword": {
		Name:               "Short Sword",
		WeaponType:         "Martial",
		DamageDie:          []int{6},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength", "Dexterity"},
		Properties:         []string{"Finesse", "Light"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"War Pick": {
		Name:               "War Pick",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	"Whip": {
		Name:               "Whip",
		WeaponType:         "Martial",
		DamageDie:          []int{4},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength", "Dexterity"},
		Properties:         []string{"Finesse", "Reach"},
		Range:              []int{},
		VersatileDamageDie: 0,
	},
	// Versatile Melee Weapons
	"Quarterstaff": {
		Name:               "Quarterstaff",
		WeaponType:         "Simple",
		DamageDie:          []int{6},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Versatile"},
		Range:              []int{},
		VersatileDamageDie: 8,
	},
	"Battle Axe": {
		Name:               "Battle Axe",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Versatile"},
		Range:              []int{},
		VersatileDamageDie: 10,
	},
	"Long Sword": {
		Name:               "Long Sword",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Slashing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Versatile"},
		Range:              []int{},
		VersatileDamageDie: 10,
	},
	"Warhammer": {
		Name:               "Warhammer",
		WeaponType:         "Martial",
		DamageDie:          []int{8},
		DamageType:         "Bludgeoning",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Versatile"},
		Range:              []int{},
		VersatileDamageDie: 10,
	},
	// Ranged Weapons
	"Light Crossbow": {
		Name:          "Light Crossbow",
		WeaponType:    "Simple",
		DamageDie:     []int{8},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged", "Loading", "Two Handed"},
		Range:         []int{80, 320},
	},
	"Dart": {
		Name:          "Dart",
		WeaponType:    "Simple",
		DamageDie:     []int{4},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Finesse", "Thrown"},
		Range:         []int{20, 60},
	},
	"Short Bow": {
		Name:          "Short Bow",
		WeaponType:    "Simple",
		DamageDie:     []int{6},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged", "Two Handed"},
		Range:         []int{80, 320},
	},
	"Sling": {
		Name:          "Sling",
		WeaponType:    "Simple",
		DamageDie:     []int{4},
		DamageType:    "Bludgeoning",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged"},
		Range:         []int{80, 320},
	},
	"Blow Gun": {
		Name:          "Blow Gun",
		WeaponType:    "Martial",
		DamageDie:     []int{1},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged", "Loading"},
		Range:         []int{25, 100},
	},
	"Hand Crossbow": {
		Name:          "Hand Crossbow",
		WeaponType:    "Martial",
		DamageDie:     []int{6},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged", "Light", "Loading"},
		Range:         []int{30, 120},
	},
	"Heavy Crossbow": {
		Name:          "Heavy Crossbow",
		WeaponType:    "Martial",
		DamageDie:     []int{10},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged", "Heavy", "Loading", "Two Handed"},
		Range:         []int{100, 400},
	},
	"Long Bow": {
		Name:          "Long Bow",
		WeaponType:    "Martial",
		DamageDie:     []int{8},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Ammunition", "Ranged", "Heavy", "Two Handed"},
		Range:         []int{150, 600},
	},
	"Net": {
		Name:          "Net",
		WeaponType:    "Martial",
		DamageDie:     []int{0},
		DamageType:    "",
		CoreAttribute: []string{"Dexterity"},
		Properties:    []string{"Special", "Thrown"},
		Range:         []int{5, 15},
	},
	"Dagger": {
		Name:          "Dagger",
		WeaponType:    "Simple",
		DamageDie:     []int{4},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Strength", "Dexterity"},
		Properties:    []string{"Finesse", "Light", "Thrown"},
		Range:         []int{20, 60},
	},
	"Handaxe": {
		Name:          "Handaxe",
		WeaponType:    "Simple",
		DamageDie:     []int{6},
		DamageType:    "Slashing",
		CoreAttribute: []string{"Strength"},
		Properties:    []string{"Light", "Thrown"},
		Range:         []int{20, 60},
	},
	"Javelin": {
		Name:          "Javelin",
		WeaponType:    "Simple",
		DamageDie:     []int{6},
		DamageType:    "Piercing",
		CoreAttribute: []string{"Strength"},
		Properties:    []string{"Thrown"},
		Range:         []int{20, 60},
	},
	"Light Hammer": {
		Name:          "Light Hammer",
		WeaponType:    "Simple",
		DamageDie:     []int{4},
		DamageType:    "Bludgeoning",
		CoreAttribute: []string{"Strength"},
		Properties:    []string{"Light", "Thrown"},
		Range:         []int{20, 60},
	},
	// Versatile Ranged Weapons
	"Trident": {
		Name:               "Trident",
		WeaponType:         "Martial",
		DamageDie:          []int{6},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Thrown", "Versatile"},
		Range:              []int{20, 60},
		VersatileDamageDie: 8,
	},
	"Spear": {
		Name:               "Spear",
		WeaponType:         "Simple",
		DamageDie:          []int{6},
		DamageType:         "Piercing",
		CoreAttribute:      []string{"Strength"},
		Properties:         []string{"Thrown", "Versatile"},
		Range:              []int{20, 60},
		VersatileDamageDie: 8,
	},
}

func GetWeaponData(name string) (Weapon, error) {
	if weaponData, ok := Weapons[name]; ok {
		return weaponData, nil
	}
	return Weapon{}, fmt.Errorf("could not retrieve weapon data for %s", name)
}

var Armors = map[string]Armor{
	"Padded Armor": {
		Name:           "Padded Armor",
		ArmorType:      "Light",
		AC:             11,
		DexMod:         true,
		DexMax:         0, // None is represented as 0
		StealthPenalty: true,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Leather Armor": {
		Name:           "Leather Armor",
		ArmorType:      "Light",
		AC:             11,
		DexMod:         true,
		DexMax:         0, // None is represented as 0
		StealthPenalty: false,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Studded Leather Armor": {
		Name:           "Studded Leather Armor",
		ArmorType:      "Light",
		AC:             12,
		DexMod:         true,
		DexMax:         0, // None is represented as 0
		StealthPenalty: false,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Hide Armor": {
		Name:           "Hide Armor",
		ArmorType:      "Medium",
		AC:             12,
		DexMod:         true,
		DexMax:         2,
		StealthPenalty: false,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Chain Shirt": {
		Name:           "Chain Shirt",
		ArmorType:      "Medium",
		AC:             13,
		DexMod:         true,
		DexMax:         2,
		StealthPenalty: false,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Scale Mail": {
		Name:           "Scale Mail",
		ArmorType:      "Medium",
		AC:             14,
		DexMod:         true,
		DexMax:         2,
		StealthPenalty: true,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Breastplate": {
		Name:           "Breastplate",
		ArmorType:      "Medium",
		AC:             14,
		DexMod:         true,
		DexMax:         2,
		StealthPenalty: false,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Half Plate": {
		Name:           "Half Plate",
		ArmorType:      "Medium",
		AC:             15,
		DexMod:         true,
		DexMax:         2,
		StealthPenalty: true,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Ring Mail": {
		Name:           "Ring Mail",
		ArmorType:      "Heavy",
		AC:             14,
		DexMod:         false,
		DexMax:         0, // None is represented as 0
		StealthPenalty: true,
		StrengthMin:    0, // None is represented as 0
		Properties:     []string{},
	},
	"Chain Mail": {
		Name:           "Chain Mail",
		ArmorType:      "Heavy",
		AC:             16,
		DexMod:         false,
		DexMax:         0, // None is represented as 0
		StealthPenalty: true,
		StrengthMin:    13,
		Properties:     []string{},
	},
	"Splint Armor": {
		Name:           "Splint Armor",
		ArmorType:      "Heavy",
		AC:             17,
		DexMod:         false,
		DexMax:         0, // None is represented as 0
		StealthPenalty: true,
		StrengthMin:    15,
		Properties:     []string{},
	},
	"Plate Armor": {
		Name:           "Plate Armor",
		ArmorType:      "Heavy",
		AC:             18,
		DexMod:         false,
		DexMax:         0, // None is represented as 0
		StealthPenalty: true,
		StrengthMin:    15,
		Properties:     []string{},
	},
}

func GetArmorData(name string) (Armor, error) {
	if armorData, ok := Armors[name]; ok {
		return armorData, nil
	}
	return Armor{}, fmt.Errorf("could not get armor data for %s", name)
}

var Shields = map[string]Shield{
	"Shield": {
		Name:       "Shield",
		ArmorType:  "Shield",
		ACBonus:    2,
		Properties: []string{},
	},
}

func GetShieldData(name string) (Shield, error) {
	if sData, ok := Shields[name]; ok {
		return sData, nil
	}
	return Shield{}, fmt.Errorf("could not get shield data for %s", name)
}
