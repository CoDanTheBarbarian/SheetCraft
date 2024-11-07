package character

import (
	bg "SheetCraft/character/backgrounds"
	rd "SheetCraft/character/races"
	"SheetCraft/items"
	"fmt"
)

var ProficiencyBaseline = map[string]bool{
	"Strength":        false,
	"Dexterity":       false,
	"Constitution":    false,
	"Intelligence":    false,
	"Wisdom":          false,
	"Charisma":        false,
	"Athletics":       false,
	"Acrobatics":      false,
	"Sleight Of Hand": false,
	"Stealth":         false,
	"Arcana":          false,
	"History":         false,
	"Investigation":   false,
	"Nature":          false,
	"Religion":        false,
	"Animal Handling": false,
	"Insight":         false,
	"Medicine":        false,
	"Perception":      false,
	"Survival":        false,
	"Deception":       false,
	"Intimidation":    false,
	"Performance":     false,
	"Persuasion":      false,
	"Light Armor":     false,
	"Medium Armor":    false,
	"Heavy Armor":     false,
	"Shield":          false,
	"Simple Weapons":  false,
	"Martial Weapons": false,
	"Club":            false,
	"Dagger":          false,
	"Greatclub":       false,
	"Handaxe":         false,
	"Javelin":         false,
	"Light Hammer":    false,
	"Mace":            false,
	"Quarterstaff":    false,
	"Sickle":          false,
	"Spear":           false,
	"Unarmed Strike":  false,
	"Light Crossbow":  false,
	"Dart":            false,
	"Short Bow":       false,
	"Sling":           false,
	"Battle Axe":      false,
	"Flail":           false,
	"Glaive":          false,
	"Great Axe":       false,
	"Great Sword":     false,
	"Halberd":         false,
	"Lance":           false,
	"Long Sword":      false,
	"Maul":            false,
	"Morning Star":    false,
	"Pike":            false,
	"Rapier":          false,
	"Scimitar":        false,
	"Short Sword":     false,
	"Trident":         false,
	"War Pick":        false,
	"Warhammer":       false,
	"Whip":            false,
	"Blow Gun":        false,
	"Hand Crossbow":   false,
	"Heavy Crossbow":  false,
	"Long Bow":        false,
	"Net":             false,
}

var ResistancesBaseline = map[string]bool{
	"Bludgeoning": false,
	"Piercing":    false,
	"Slashing":    false,
	"Acid":        false,
	"Cold":        false,
	"Fire":        false,
	"Force":       false,
	"Lightning":   false,
	"Necrotic":    false,
	"Poison":      false,
	"Psychic":     false,
	"Radiant":     false,
	"Thunder":     false,
}

var SkillsCoreStatMap = map[string]string{
	"Athletics":       "Strength",
	"Acrobatics":      "Dexterity",
	"Sleight Of Hand": "Dexterity",
	"Stealth":         "Dexterity",
	"History":         "Intelligence",
	"Investigation":   "Intelligence",
	"Nature":          "Intelligence",
	"Religion":        "Intelligence",
	"Animal Handling": "Wisdom",
	"Insight":         "Wisdom",
	"Medicine":        "Wisdom",
	"Perception":      "Wisdom",
	"Survival":        "Wisdom",
	"Deception":       "Charisma",
	"Intimidation":    "Charisma",
	"Performance":     "Charisma",
	"Persuasion":      "Charisma",
}

func GetSkillCoreStat(skill string) (string, error) {
	coreStat, ok := SkillsCoreStatMap[skill]
	if !ok {
		return "", fmt.Errorf("skill %s not found in SkillsCoreStatMap", skill)
	}
	return coreStat, nil
}

var SpellSlotsBaseline = map[string]int{
	"Cantrip": 0,
	"Level 1": 0,
	"Level 2": 0,
	"Level 3": 0,
	"Level 4": 0,
	"Level 5": 0,
	"Level 6": 0,
	"Level 7": 0,
	"Level 8": 0,
	"Level 9": 0,
}

type Character struct {
	Name     string
	RaceName string
	RaceInfo []struct {
		Trait string
		Info  string
	}
	SubraceName         string
	ClassName           string
	SubclassName        string
	ClassInfo           map[string]string
	ClassAbilities      map[string]string
	Background          string
	Speed               int
	HP                  int
	HPBonus             int
	HitDie              int
	ProficiencyBonus    int
	Level               int
	XP                  int
	Strength            int
	Dexterity           int
	Constitution        int
	Intelligence        int
	Wisdom              int
	Charisma            int
	AC                  int
	Proficiencies       map[string]bool
	DamageResistance    map[string]bool
	Inventory           []string // TODO: turn this into an Item interface
	Weapon              *items.Weapon
	OffHandWeapon       *items.Weapon
	Armor               *items.Armor
	Shield              *items.Shield
	SpellCastingAbility string
	SpellSlots          map[string]int
	SpellsKnown         int
	SpellList           []string
	BreathWeapon        struct {
		BreathShape string
		BreathSize  []int
		BreathType  string
	}
}

// New Character constructor

func NewCharacter(Name string) (Character, error) {
	character := Character{
		Name:     Name,
		RaceName: "",
		RaceInfo: make([]struct {
			Trait string
			Info  string
		}, 0),
		SubraceName:         "",
		ClassName:           "",
		SubclassName:        "",
		ClassInfo:           make(map[string]string),
		ClassAbilities:      make(map[string]string),
		Background:          "",
		Speed:               0,
		HP:                  0,
		HPBonus:             0,
		HitDie:              0,
		ProficiencyBonus:    2,
		Level:               1,
		XP:                  0,
		Strength:            8,
		Dexterity:           8,
		Constitution:        8,
		Intelligence:        8,
		Wisdom:              8,
		Charisma:            8,
		AC:                  10,
		Proficiencies:       make(map[string]bool),
		DamageResistance:    make(map[string]bool),
		Inventory:           make([]string, 0),
		Weapon:              nil,
		OffHandWeapon:       nil,
		Armor:               nil,
		Shield:              nil,
		SpellCastingAbility: "",
		SpellSlots:          make(map[string]int),
		SpellsKnown:         0,
		SpellList:           make([]string, 0),
		BreathWeapon: struct {
			BreathShape string
			BreathSize  []int
			BreathType  string
		}{
			BreathShape: "",
			BreathSize:  make([]int, 0),
			BreathType:  "",
		},
	}
	if err := character.init(); err != nil {
		return Character{}, err
	}
	return character, nil
}

func (c *Character) init() error {
	for k := range ProficiencyBaseline {
		c.Proficiencies[k] = ProficiencyBaseline[k]
	}
	for k := range ResistancesBaseline {
		c.DamageResistance[k] = ResistancesBaseline[k]
	}
	for k := range SpellSlotsBaseline {
		c.SpellSlots[k] = SpellSlotsBaseline[k]
	}
	return nil
}

// Proficiencies and Resistances

func (c *Character) GainProficiencies(profs []string) error {
	for _, prof := range profs {
		if _, exists := c.Proficiencies[prof]; !exists {
			return fmt.Errorf("error: proficiency '%s' does not exist", prof)
		} else {
			c.Proficiencies[prof] = true
		}
	}
	return nil
}

func (c *Character) GainResistances(res []string) error {
	for _, res := range res {
		if _, exists := c.DamageResistance[res]; !exists {
			return fmt.Errorf("error: Damage type '%s' does not exist in Resistances", res)
		} else {
			c.DamageResistance[res] = true
		}
	}
	return nil
}

func (c *Character) ProfMod(prof string) (int, error) {
	if _, exists := c.Proficiencies[prof]; !exists {
		return 0, fmt.Errorf("error: proficiency '%s' does not exist", prof)
	}
	if c.Proficiencies[prof] {
		return c.ProficiencyBonus, nil
	} else {
		return 0, nil
	}
}

// Core Stats, Skills, and Saving Throws

func (c *Character) AssignStat(stat string, stat_num int) error {
	if stat_num > 20 || stat_num < 1 {
		return fmt.Errorf("invalid stat number %d: must be between 1 and 20", stat_num)
	}
	switch stat {
	case "Strength":
		c.Strength = stat_num
		return nil
	case "Dexterity":
		c.Dexterity = stat_num
		return nil
	case "Constitution":
		c.Constitution = stat_num
		return nil
	case "Intelligence":
		c.Intelligence = stat_num
		return nil
	case "Wisdom":
		c.Wisdom = stat_num
		return nil
	case "Charisma":
		c.Charisma = stat_num
		return nil
	default:
		return fmt.Errorf("invalid stat %s", stat)
	}
}

func (c *Character) IncreaseStat(stat string, amount int) error {
	max := 20
	switch stat {
	case "Strength":
		if c.Strength+amount > max {
			return fmt.Errorf("strength cannot be increased above 20")
		}
		c.Strength += amount
	case "Dexterity":
		if c.Dexterity+amount > max {
			return fmt.Errorf("dexterity cannot be increased above 20")
		}
		c.Dexterity += amount
	case "Constitution":
		if c.Constitution+amount > max {
			return fmt.Errorf("constitution cannot be increased above 20")
		}
		c.Constitution += amount
	case "Intelligence":
		if c.Intelligence+amount > max {
			return fmt.Errorf("intelligence cannot be increased above 20")
		}
		c.Intelligence += amount
	case "Wisdom":
		if c.Wisdom+amount > max {
			return fmt.Errorf("wisdom cannot be increased above 20")
		}
		c.Wisdom += amount
	case "Charisma":
		if c.Charisma+amount > max {
			return fmt.Errorf("charisma cannot be increased above %d", max)
		}
		c.Charisma += amount
	default:
		return fmt.Errorf("stat %s does not exist", stat)
	}
	return nil
}

func (c *Character) DecreaseStat(stat string, amount int) error {
	min := 1
	switch stat {
	case "Strength":
		if c.Strength-amount < min {
			return fmt.Errorf("strength cannot be decreased lower than %d", min)
		}
		c.Strength -= amount
	case "Dexterity":
		if c.Dexterity-amount < min {
			return fmt.Errorf("dexterity cannot be decreased lower than %d", min)
		}
		c.Dexterity -= amount
	case "Constitution":
		if c.Constitution-amount < min {
			return fmt.Errorf("constitution cannot be decreased lower than %d", min)
		}
		c.Constitution -= amount
	case "Intelligence":
		if c.Intelligence-amount < min {
			return fmt.Errorf("intelligence cannot be decreased lower than %d", min)
		}
		c.Intelligence -= amount
	case "Wisdom":
		if c.Wisdom-amount < min {
			return fmt.Errorf("wisdom cannot be decreased lower than %d", min)
		}
		c.Wisdom -= amount
	case "Charisma":
		if c.Charisma-amount < min {
			return fmt.Errorf("charisma cannot be decreased lower than %d", min)
		}
		c.Charisma -= amount
	default:
		return fmt.Errorf("stat %s does not exist", stat)
	}
	return nil
}

func (c *Character) GetAbilityMod(stat string) (int, error) {
	switch stat {
	case "Strength":
		return ((c.Strength / 2) - 5), nil
	case "Dexterity":
		return ((c.Dexterity / 2) - 5), nil
	case "Constitution":
		return ((c.Constitution / 2) - 5), nil
	case "Intelligence":
		return ((c.Intelligence / 2) - 5), nil
	case "Wisdom":
		return ((c.Wisdom / 2) - 5), nil
	case "Charisma":
		return ((c.Charisma / 2) - 5), nil
	default:
		return -6, fmt.Errorf("stat '%s' does not exist", stat)
	}
}

func (c *Character) GetSavingThrow(stat string) (int, error) {
	abilityMod, err := c.GetAbilityMod(stat)
	if err != nil {
		return -1, err
	}
	profMod, err := c.ProfMod(stat)
	if err != nil {
		return -1, err
	}

	return abilityMod + profMod, nil
}

func (c *Character) GetSkillMod(skill string) (int, error) {
	stat, err := GetSkillCoreStat(skill)
	if err != nil {
		return 0, err
	}

	mod, err := c.GetAbilityMod(stat)
	if err != nil {
		return 0, err
	}

	profMod, err := c.ProfMod(skill)
	if err != nil {
		return 0, err
	}

	return mod + profMod, nil
}

// Race, Subrace and Class applications

func (c *Character) AssignRace(r *rd.Race) {
	c.RaceName = r.Name
	c.Speed = r.Speed
	if len(r.StatBonus) > 0 {
		for _, increase := range r.StatBonus {
			c.IncreaseStat(increase.Stat, increase.Bonus)
		}
	}
	c.GainProficiencies(r.Proficiencies)
	c.GainResistances(r.Resistances)
	for _, trait := range r.RaceTraits {
		c.RaceInfo = append(c.RaceInfo, struct {
			Trait string
			Info  string
		}{
			Trait: trait.Trait,
			Info:  trait.Explanation,
		})
	}
}

func (c *Character) AssignSubrace(s *rd.SubRace) {
	c.SubraceName = s.Name
	if len(s.StatBonus) > 0 {
		for _, increase := range s.StatBonus {
			c.IncreaseStat(increase.Stat, increase.Bonus)
		}
	}
	c.GainProficiencies(s.Proficiencies)
	c.GainResistances(s.Resistances)
}

func (c *Character) AssignDragonColor(d *rd.DragonColor) {
	c.SubraceName = d.Name
	if len(d.StatBonus) > 0 {
		for _, increase := range d.StatBonus {
			c.IncreaseStat(increase.Stat, increase.Bonus)
		}
	}
	c.GainProficiencies(d.Proficiencies)
	c.GainResistances(d.Resistances)
	c.BreathWeapon.BreathShape = d.BreathWeapon.BreathShape
	c.BreathWeapon.BreathType = d.BreathWeapon.BreathType
	c.BreathWeapon.BreathSize = d.BreathWeapon.BreathSize
}

// Background application, Starting AC/HP calcs

func (c *Character) GainBackground(background string) error {
	b, ok := bg.Backgrounds[background]
	if !ok {
		return fmt.Errorf("background '%s' does not exist", background)
	}
	c.Background = background
	c.GainProficiencies(b.Proficiencies)
	return nil
}

func (c *Character) SetUnarmoredAC() {
	if c.AC != 10 {
		c.AC = 10
	}
	if c.Armor == nil {
		dexMod, err := c.GetAbilityMod("Dexterity")
		if err == nil {
			c.AC += dexMod
		}
		if c.ClassName == "Barbarian" {
			conMod, err := c.GetAbilityMod("Constitution")
			if err == nil {
				c.AC += conMod
			}
		}
	}
}

func (c *Character) SetMaxHP() {
	conMod, err := c.GetAbilityMod("Constitution")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		c.HP = c.HitDie + c.HPBonus
		return
	}
	c.HP = c.HitDie + conMod + c.HPBonus
}

// Spell Info

func (c *Character) SpellAttackMod() int {
	mod, err := c.GetAbilityMod(c.SpellCastingAbility)
	if err != nil {
		return 0
	}
	return mod + c.ProficiencyBonus
}

func (c *Character) SpellSaveDC() int {
	if c.SpellCastingAbility != "" {
		return c.SpellAttackMod() + 8
	}
	return 0
}
