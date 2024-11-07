package character

import (
	"fmt"
)

type Class struct {
	ClassName     string
	HitDie        int
	Proficiencies []string
	ProfChoice    []struct {
		Choices      []string
		ChoiceNumber int
	}
	ClassInfo []struct {
		Name string
		Info string
	}
	StartingEquipment []string
	ClassAbilities    []struct {
		Name string
		Uses string
	}
	SpellCastingAbility string
	SpellSlots          map[string]int
	SpellsKnown         int
	Subclass            string
	SubclassInfo        []struct {
		Name string
		Info string
	}
}

var Classes = map[string]Class{
	"Barbarian": {
		ClassName: "Barbarian",
		HitDie:    12,
		Proficiencies: []string{
			"Light Armor", "Medium Armor", "Shield", "Simple Weapons", "Martial Weapons", "Strength", "Constitution",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Animal Handling", "Athletics", "Intimidation", "Nature", "Perception", "Survival"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Unarmored Defense", Info: "While you are not wearing any armor, your AC equals 10 + your Dexterity modifier + your Constitution modifier. You can use a shield and still gain this benefit."},
			{Name: "Rage", Info: "During your turn in combat, you can rage as a bonus action. While raging you gain the following benefits:\n - You have advantage of Stregth checks and saving throws\n - When you make a strength based weapon attack add 2 to your attack roll\n - You have resistance to bludgeoning, piercing, and slashing damage\nYou can't cast spells while raging. Your rage lasts for 1 minute and ends if you are knocked unconscious or you haven't attacked a hostile creature since your last turn."},
		},
		StartingEquipment: []string{
			"Great Axe", "Halberd", "Javelin",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{
			{Name: "Rage Charges", Uses: "2 Charges"},
		},
		SpellCastingAbility: "",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Bard": {
		ClassName: "Bard",
		HitDie:    8,
		Proficiencies: []string{
			"Light Armor", "Simple Weapons", "Martial Rewards", "Dexterity", "Charisma",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Spellcasting", Info: "You've learned to cast spells using your music."},
			{Name: "Bardic Inspiration", Info: "You can use a bonus action to give one creature other than yourself a Bardic Inspiration die to add to their next ability check, attack roll or saving throw in the next 10 minutes. The target can wait until after they roll to use the Bardic die, but must decide before the DM declares success or failure. You regain any expended uses on a long rest."},
		},
		StartingEquipment: []string{
			"Rapier", "Dagger",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{
			{Name: "Bardic Inspiration", Uses: "A number of times equal to your Charisma modifier (minimum 1)."},
		},
		SpellCastingAbility: "Charisma",
		SpellSlots: map[string]int{
			"Cantrips": 4,
			"Level 1":  2,
		},
		SpellsKnown: 4,
		Subclass:    "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Cleric": {
		ClassName: "Cleric",
		HitDie:    8,
		Proficiencies: []string{
			"Light Armor", "Medium Armor", "Shield", "Simple Weapons", "Martial Weapons", "Wisdom", "Charisma",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"History", "Insight", "Medicine", "Persuasion", "Religion"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Spellcasting", Info: "You can cast spells granted to you by your Divine Domain. Your domain spells are always prepared and don't count against the number of spells you can have prepared."},
		},
		StartingEquipment: []string{
			"Light Crossbow", "Mace", "Shield", "Scale Mail",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{},
		SpellCastingAbility: "Wisdom",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Druid": {
		ClassName: "Druid",
		HitDie:    8,
		Proficiencies: []string{
			"Light Armor", "Medium Armor", "Shield", "Club", "Dagger", "Dart", "Javelin", "Mace", "Quarterstaff", "Scimitar", "Sickle", "Sling", "Spear", "Intelligence", "Wisdom",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Arcana", "Animal Handling", "Insight", "Medicine", "Nature", "Perception", "Survival"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Spellcasting", Info: "You can cast spells drawn from the divine essence of nature itself."},
		},
		StartingEquipment: []string{
			"Scimitar", "Leather Armor", "Shield",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{},
		SpellCastingAbility: "Charisma",
		SpellSlots: map[string]int{
			"Cantrips": 2,
			"Level 1":  2,
		},
		SpellsKnown: 0,
		Subclass:    "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Fighter": {
		ClassName: "Fighter",
		HitDie:    10,
		Proficiencies: []string{
			"Light Armor", "Medium Armor", "Heavy Armor", "Shield", "Simple Weapons", "Martial Weapons", "Strength", "Constitution",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Athletics", "Acrobatics", "Sleight of Hand", "Stealth", "Performance", "Persuasion"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Second Wind", Info: "On your turn you can use a bonus action to regain hit points eqaul to 1d10 + your fighter level. Once you use this feature you can't use it again until you finish a short or long rest."},
		},
		StartingEquipment: []string{
			"Great Axe", "Long Bow", "Light Crossbow", "Long Sword", "Leather Armor", "Shield",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{
			{Name: "Second Wind", Uses: "Once per short rest."},
		},
		SpellCastingAbility: "",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Monk": {
		ClassName: "Monk",
		HitDie:    8,
		Proficiencies: []string{
			"Simple Weapons", "Short Sword", "Strength", "Dexterity", "Unarmed Strike",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Acrobatics", "Athletics", "History", "Insight", "Religion", "Stealth"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Martial Arts", Info: "While you are unarmed or wielding a monk weapon (short swords or any melee weapon that does not have the 'heavy' or 'two-handed' property) and you aren't wearing armor, you gain the following benefits."},
			{Name: "Deft Strike", Info: "You can use Dexterity instead of Strength for the attack and damage rolls of your unarmed strikes and monk weapons."},
			{Name: "Well Trained", Info: "You can roll a d4 in place of the normal damage of your unarmed strike or monk weapon. This die changes as you gain levels."},
			{Name: "Quick Hands", Info: "When you use the Attack action with your unarmed strike or monk weapon on your turn, you can use a bonus action to make an unarmed strike."},
		},
		StartingEquipment: []string{
			"Dagger", "Short Sword", "Leather Armor", "Shield",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{},
		SpellCastingAbility: "",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Paladin": {
		ClassName: "Paladin",
		HitDie:    10,
		Proficiencies: []string{
			"Light Armor", "Medium Armor", "Heavy Armor", "Shield", "Simple Weapons", "Martial Weapons", "Wisdom", "Charisma",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Athletics", "Insight", "Intimidation", "Medicine", "Persuasion", "Religion"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Divine Sense", Info: "On your turn as an action you invoke your divine sense. Until the end of your next turn, you know the location of any celestial, fiend or undead within 60 feet of you that is not behind total cover. You know the type of creature sensed, but not it's identity. Regain all uses on a long rest."},
			{Name: "Lay on Hands", Info: "As an action, you can can touch a creature and restore their hit points up to the amount in your pool, or expend 5 hit points from your pool to heal that creature of a disease or poison. Your pool is 5 times your paladin level and replenishes on a long rest."},
		},
		StartingEquipment: []string{
			"Short Bow", "Long Sword", "Ring Mail", "Shield",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{
			{Name: "Divine Sense", Uses: "One plus your Charisma modifier. Replenishes on a long rest."},
			{Name: "Lay on Hands", Uses: "Five times your paladin level. Replenishes on a long rest."},
		},
		SpellCastingAbility: "Wisdom",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Ranger": {
		ClassName: "Ranger",
		HitDie:    10,
		Proficiencies: []string{
			"Light Armor", "Medium Armor", "Heavy Armor", "Shield", "Simple Weapons", "Martial Weapons", "Strength", "Dexterity",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Animal Handling", "Athletics", "Insight", "Investigation", "Nature", "Perception", "Stealth", "Survival"}, ChoiceNumber: 3},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Favored Enemy", Info: "You have advantage on Survival checks to track your favored enemy and on Intelligence checks to recall information about them."},
			{Name: "Favored Terrain", Info: "While in your favored terrain, when you make Intelligence and Wisdom checks related to that terrain, your proficiency bonus is doubled if applicable.\nWhile traveling for an hour or more in your favored terrain you gain the following benefits:\n - Unaffected by difficult terrain\n - Can't get lost except by magical means\n - Move stealthily at a normal pace"},
		},
		StartingEquipment: []string{
			"Short Sword", "Short Sword", "Long Bow", "Ring Mail",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{
			{Name: "Favored Enemy", Uses: ""},
			{Name: "Favored Terrain", Uses: ""},
		},
		SpellCastingAbility: "",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Rogue": {
		ClassName: "Rogue",
		HitDie:    8,
		Proficiencies: []string{
			"Light Armor", "Hand Crossbow", "Long Sword", "Rapier", "Short Sword", "Simple Weapons", "Dexterity", "Intelligence",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Acrobatics", "Athletics", "Deception", "Insight", "Intimidation", "Investigation", "Perception", "Performance", "Persuasion", "Sleight of Hand", "Stealth"}, ChoiceNumber: 4},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Expertise", Info: "Choose two proficiencies or one and thieves tools. Your proficiency bonus is doubled when using those proficiencies."},
			{Name: "Sneak Attack", Info: "Once per turn you can deal an extra 1d6 damage when you hit and have advantage on the attack roll or if the target has an enemy within 5 feet of it. The attack must use a finesse or a ranged weapon."},
		},
		StartingEquipment: []string{
			"Dagger", "Dagger", "Short Sword", "Short Bow", "Leather Armor",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{
			{Name: "Sneak Attack", Uses: "1d6"},
		},
		SpellCastingAbility: "",
		SpellSlots:          map[string]int{},
		SpellsKnown:         0,
		Subclass:            "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Sorcerer": {
		ClassName: "Sorcerer",
		HitDie:    6,
		Proficiencies: []string{
			"Dagger", "Dart", "Sling", "Quarterstaff", "Light Crossbow", "Constitution", "Charisma",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Arcana", "Deception", "Insight", "Intimidation", "Persuasion", "Religion"}, ChoiceNumber: 2}, {Choices: []string{"Arcana", "Deception", "Insight", "Intimidation", "Persuasion", "Religion"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Spellcasting", Info: "You can cast spells from your sorcerous bloodline."},
		},
		StartingEquipment: []string{
			"Dagger", "Dagger", "Light Crossbow", "Quarterstaff",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{},
		SpellCastingAbility: "Charisma",
		SpellSlots: map[string]int{
			"Cantrip": 4,
			"Level 1": 2,
		},
		SpellsKnown: 2,
		Subclass:    "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Warlock": {
		ClassName: "Warlock",
		HitDie:    8,
		Proficiencies: []string{
			"Light Armor", "Simple Weapons", "Wisdom", "Charisma",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Arcana", "Deception", "History", "Intimidation", "Investigation", "Nature", "Religion"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Pact Magic", Info: "You can cast spells granted to you by your Patron."},
		},
		StartingEquipment: []string{
			"Light Crossbow", "Handaxe", "Dagger", "Dagger", "Leather Armor",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{},
		SpellCastingAbility: "Charisma",
		SpellSlots: map[string]int{
			"Cantrip": 2,
			"Level 1": 2,
		},
		SpellsKnown: 2,
		Subclass:    "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
	"Wizard": {
		ClassName: "Wizard",
		HitDie:    6,
		Proficiencies: []string{
			"Dagger", "Dart", "Sling", "Quarterstaff", "Light Crossbow", "Intelligence", "Wisdom",
		},
		ProfChoice: []struct {
			Choices      []string
			ChoiceNumber int
		}{
			{Choices: []string{"Arcana", "History", "Insight", "Investigation", "Medicine", "Religion"}, ChoiceNumber: 2},
		},
		ClassInfo: []struct {
			Name string
			Info string
		}{
			{Name: "Spellcasting", Info: "You can cast spells from your research into the arcane."},
			{Name: "Arcan Recovery", Info: "Once per day when you finish a short rest you can choose expended spell slots to recover. The spell slots can have a combined total of no more than half your wizard level (rounded up)."},
		},
		StartingEquipment: []string{
			"Quarterstaff",
		},
		ClassAbilities: []struct {
			Name string
			Uses string
		}{},
		SpellCastingAbility: "Intelligence",
		SpellSlots: map[string]int{
			"Cantrip": 3,
			"Level 1": 2,
		},
		SpellsKnown: 6,
		Subclass:    "",
		SubclassInfo: []struct {
			Name string
			Info string
		}{},
	},
}

func GetClassData(className string) (classData Class, err error) {
	if classData, ok := Classes[className]; ok {
		return classData, nil
	}
	return Class{}, fmt.Errorf("could not get class data for %s", className)
}
