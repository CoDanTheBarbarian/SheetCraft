package character

import (
	"reflect"
	"testing"
)

func TestGetSkillCoreStat(t *testing.T) {
	_, err := GetSkillCoreStat("Athletic")
	if err == nil {
		t.Error("Expected error, but got nil")
	}
	stat, err := GetSkillCoreStat("Athletics")
	if err != nil {
		t.Errorf("Expected no error, gut got: %v", err)
	}
	if stat != "Strength" {
		t.Errorf("Expected stat to be Strength, but got %s", stat)
	}
}

func TestNewCharacter(t *testing.T) {
	name := "Test Character"
	c, err := NewCharacter(name)
	if err != nil {
		t.Errorf("Expected name: %s, got: %s", name, c.Name)
	}
	if c.RaceName != "" {
		t.Errorf("Expected race: '', got: %s", c.RaceName)
	}
	if len(c.RaceInfo) != 0 {
		t.Errorf("Expected RaceInfo to be empty, but got: %v", c.RaceInfo)
	}
	if c.SubraceName != "" {
		t.Errorf("Expected SubraceName to be empty, but got: '%s'", c.SubraceName)
	}
	if c.ClassName != "" {
		t.Errorf("Expected ClassName to be empty, but got: '%s'", c.ClassName)
	}
	if c.SubclassName != "" {
		t.Errorf("Expected SubclassName to be empty, but got: '%s'", c.SubclassName)
	}
	if len(c.ClassInfo) != 0 {
		t.Errorf("Expected ClassIngo to be empty, but got: %v", c.ClassInfo)
	}
	if len(c.ClassAbilities) != 0 {
		t.Errorf("Expected ClassIngo to be empty, but got: %v", c.ClassAbilities)
	}
	if c.Background != "" {
		t.Errorf("Expected Background to be empty, but got: '%s'", c.Background)
	}
	if c.Speed != 0 {
		t.Errorf("Expected Speed to be 0, but got: %d", c.Speed)
	}
	if c.HP != 0 {
		t.Errorf("Expected HP to be 0, but got: %d", c.HP)
	}
	if c.HPBonus != 0 {
		t.Errorf("Expected HPBonus to be 0, but got: %d", c.HPBonus)
	}
	if c.HitDie != 0 {
		t.Errorf("Expected HitDie to be 0, but got: %d", c.HitDie)
	}
	if c.ProficiencyBonus != 2 {
		t.Errorf("Expected ProficiencyBonus to be 2, but got: %d", c.ProficiencyBonus)
	}
	if c.Level != 1 {
		t.Errorf("Expected Level to be 1, but got: %d", c.Level)
	}
	if c.XP != 0 {
		t.Errorf("Expected XP to be 0, but got: %d", c.XP)
	}
	if c.Strength != 8 {
		t.Errorf("Expected Strength to be 8, but got: %d", c.Strength)
	}
	if c.Dexterity != 8 {
		t.Errorf("Expected Dexterity to be 8, but got: %d", c.Dexterity)
	}
	if c.Constitution != 8 {
		t.Errorf("Expected Constitution to be 8, but got: %d", c.Constitution)
	}
	if c.Intelligence != 8 {
		t.Errorf("Expected Intelligence to be 8, but got: %d", c.Intelligence)
	}
	if c.Wisdom != 8 {
		t.Errorf("Expected Wisdom to be 8, but got: %d", c.Wisdom)
	}
	if c.Charisma != 8 {
		t.Errorf("Expected Charisma to be 8, but got: %d", c.Charisma)
	}
	if c.AC != 10 {
		t.Errorf("Expected AC to be 10, but got: %d", c.AC)
	}
	if len(c.Proficiencies) == 0 {
		t.Error("Expected Proficiencies to not be empty, but it is")
	}
	if !reflect.DeepEqual(c.Proficiencies, ProficiencyBaseline) {
		t.Errorf("Expected Proficiencies to be %v, but got: %v", ProficiencyBaseline, c.Proficiencies)
	}
	if len(c.DamageResistance) == 0 {
		t.Error("Expected DamageResistance to not be empty, but it is")
	}
	if !reflect.DeepEqual(c.DamageResistance, ResistancesBaseline) {
		t.Errorf("Expected Resistances to be %v, but got: %v", ResistancesBaseline, c.DamageResistance)
	}
	if len(c.Inventory) != 0 {
		t.Errorf("Expected Inventory to be empty, but got: %v", c.Inventory)
	}
	if c.Weapon != nil {
		t.Errorf("Expected Weapon to be nil, but got: %v", c.Weapon)
	}
	if c.OffHandWeapon != nil {
		t.Errorf("Expected OffHandWeapon to be nil, but got: %v", c.OffHandWeapon)
	}
	if c.Armor != nil {
		t.Errorf("Expected Armor to be nil, but got: %v", c.Armor)
	}
	if c.Shield != nil {
		t.Errorf("Expected Shield to be nil, but got: %v", c.Shield)
	}
	if c.SpellCastingAbility != "" {
		t.Errorf("Expected SpellCastingAbility to be empty, but got: %s", c.SpellCastingAbility)
	}
	if len(c.SpellSlots) == 0 {
		t.Error("Expected SpellSlots to not be empty, but it is")
	}
	if !reflect.DeepEqual(c.SpellSlots, SpellSlotsBaseline) {
		t.Errorf("Expected SpellSlots to be %v, but got: %v", SpellSlotsBaseline, c.SpellSlots)
	}
	if c.SpellsKnown != 0 {
		t.Errorf("Expected SpellsKnown to be 0, but got: %d", c.SpellsKnown)
	}
	if len(c.SpellList) != 0 {
		t.Errorf("Expected SpellList to be empty, but got: %v", c.SpellList)
	}
	if c.BreathWeapon.BreathShape != "" {
		t.Errorf("Expected BreathShape to be empty. but got: %s", c.BreathWeapon.BreathShape)
	}
	if len(c.BreathWeapon.BreathSize) != 0 {
		t.Errorf("Expected BreathSize to be empty, but got: %v", c.BreathWeapon.BreathSize)
	}
	if c.BreathWeapon.BreathType != "" {
		t.Errorf("Expected BreathType to be empty. but got: %s", c.BreathWeapon.BreathType)
	}

}

func TestGainProficiencies(t *testing.T) {
	c, _ := NewCharacter("Test")
	if c.Proficiencies["Martial Weapons"] == true {
		t.Error("Expected Martial Weapons to be false, but got true.")
	}
	if c.Proficiencies["Simple Weapons"] == true {
		t.Error("Expected Simple Weapons to be false, but got true.")
	}
	if c.Proficiencies["Light Armor"] == true {
		t.Error("Expected Light Armor to be false, but got true.")
	}
	if c.Proficiencies["Athletics"] == true {
		t.Error("Expected Martial Weapons to be false, but got true.")
	}
	if c.Proficiencies["Dexterity"] == true {
		t.Error("Expected Martial Weapons to be false, but got true.")
	}
	profs := []string{
		"Martial Weapons", "Simple Weapons", "Light Armor", "Athletics", "Dexterity",
	}
	err := c.GainProficiencies(profs)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if c.Proficiencies["Martial Weapons"] != true {
		t.Error("Expected Martial Weapons to be true, but got false.")
	}
	if c.Proficiencies["Simple Weapons"] != true {
		t.Error("Expected Simple Weapons to be true, but got false.")
	}
	if c.Proficiencies["Light Armor"] != true {
		t.Error("Expected Light Armor to be true, but got false.")
	}
	if c.Proficiencies["Athletics"] != true {
		t.Error("Expected Martial Weapons to be true, but got false.")
	}
	if c.Proficiencies["Dexterity"] != true {
		t.Error("Expected Martial Weapons to be true, but got false.")
	}
	profs2 := []string{
		"Martial Rewards", "light armor", "Short Swords", "Shields",
	}
	err = c.GainProficiencies(profs2)
	if err == nil {
		t.Errorf("Expected error, instead got nil")
	}
}

func TestGainResistances(t *testing.T) {
	c, _ := NewCharacter("Test")
	if c.DamageResistance["Acid"] == true {
		t.Error("Expected Acid to be false, but it was true")
	}
	if c.DamageResistance["Radiant"] == true {
		t.Error("Expected Radiant to be false, but it was true")
	}
	if c.DamageResistance["Fire"] == true {
		t.Error("Expected Fire to be false, but it was true")
	}
	if c.DamageResistance["Cold"] == true {
		t.Error("Expected Cold to be false, but it was true")
	}
	if c.DamageResistance["Thunder"] == true {
		t.Error("Expected Thunder to be false, but it was true")
	}
	res := []string{
		"Acid", "Radiant", "Fire", "Cold",
	}
	err := c.GainResistances(res)
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if c.DamageResistance["Acid"] != true {
		t.Error("Expected Acid to be true, but it was false")
	}
	if c.DamageResistance["Radiant"] != true {
		t.Error("Expected Radiant to be true, but it was false")
	}
	if c.DamageResistance["Fire"] != true {
		t.Error("Expected Fire to be true, but it was false")
	}
	if c.DamageResistance["Cold"] != true {
		t.Error("Expected Cold to be true, but it was false")
	}
	if c.DamageResistance["Thunder"] == true {
		t.Error("Expected Thunder to be false, but it was true")
	}
	res2 := []string{
		"Acid", "Thunder Damage", "Bludgening",
	}
	err = c.GainResistances(res2)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

}

func TestProfMod(t *testing.T) {
	c, _ := NewCharacter("Test")
	_, err := c.ProfMod("Strengtth")
	if err == nil {
		t.Error("Expected to return an err, but it did not.")
	}
	mod, err := c.ProfMod("Strength")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 0 {
		t.Errorf("Expected Strength mod to be 0, but got: %d", mod)
	}
	profs := []string{
		"Strength", "Athletics", "Perception", "Short Sword",
	}
	err = c.GainProficiencies(profs)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	mod, err = c.ProfMod("Strength")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 2 {
		t.Errorf("Expected Strength mod to be 0, but got: %d", mod)
	}
	mod, err = c.ProfMod("Dexterity")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 0 {
		t.Errorf("Expected Dexterity mod to be 0, but got: %d", mod)
	}
	mod, err = c.ProfMod("Athletics")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 2 {
		t.Errorf("Expected Athletics mod to be 2, but got: %d", mod)
	}
	mod, err = c.ProfMod("Perception")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 2 {
		t.Errorf("Expected Perception mod to be 2, but got: %d", mod)
	}
	mod, err = c.ProfMod("Short Sword")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 2 {
		t.Errorf("Expected Short Sword mod to be 2, but got: %d", mod)
	}

}

func TestAssignStat(t *testing.T) {
	c, _ := NewCharacter("Test")
	err := c.AssignStat("Strengt", 2)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.AssignStat("Strength", 0)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.AssignStat("Strength", 35)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.AssignStat("Strength", 20)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if c.Strength != 20 {
		t.Errorf("Expected Strength to be 20, but got: %d", c.Strength)
	}
}

func TestIncreaseStat(t *testing.T) {
	c, _ := NewCharacter("Test")
	err := c.IncreaseStat("Strengh", 10)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.IncreaseStat("Strength", 13)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.IncreaseStat("Strength", 12)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if c.Strength != 20 {
		t.Errorf("Expected Strength to be 20, but got: %d", c.Strength)
	}
}

func TestDecreaseStat(t *testing.T) {
	c, _ := NewCharacter("Test")
	err := c.DecreaseStat("Stregt", 6)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.DecreaseStat("Strength", 8)
	if err == nil {
		t.Error("Expected error, but got none")
	}
	err = c.DecreaseStat("Strength", 6)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if c.Strength != 2 {
		t.Errorf("Expected Strength to be 2, but got: %d", c.Strength)
	}
}

func TestGetAbilityMod(t *testing.T) {
	c, _ := NewCharacter("Test")
	_, err := c.GetAbilityMod("Strengt")
	if err == nil {
		t.Error("Expected error, but got none")
	}
	mod, err := c.GetAbilityMod("Strength")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != -1 {
		t.Errorf("Expected mod to be -1, but got: %d", mod)
	}
	err = c.AssignStat("Strength", 20)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	mod, err = c.GetAbilityMod("Strength")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 5 {
		t.Errorf("Expected mod to be 5, but got: %d", mod)
	}

}

func TestGetSavingThrow(t *testing.T) {
	c, _ := NewCharacter("Test")
	mod, err := c.GetSavingThrow("Strength")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != -1 {
		t.Errorf("Expected mod to be -1, but got: %d", mod)
	}
	_ = c.GainProficiencies([]string{"Strength"})
	mod, err = c.GetSavingThrow("Strength")
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	if mod != 1 {
		t.Errorf("Expected mod to be 1, but got: %d", mod)
	}
}

func TestGetSkillMod(t *testing.T) {
	c, _ := NewCharacter("Test")
	_, err := c.GetSkillMod("Athletic")
	if err == nil {
		t.Error("Expected error but got nil")
	}
	mod, err := c.GetSkillMod("Athletics")
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if mod != -1 {
		t.Errorf("Expected athletics mod to be -1, but got: %d", mod)
	}
	_ = c.AssignStat("Strength", 20)
	mod, err = c.GetSkillMod("Athletics")
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if mod != 5 {
		t.Errorf("Expected athletics mod to be 5, but got: %d", mod)
	}
	_ = c.GainProficiencies([]string{"Athletics"})
	mod, err = c.GetSkillMod("Athletics")
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if mod != 7 {
		t.Errorf("Expected athletics mod to be 7, but got: %d", mod)
	}
}

func TestGainBackground(t *testing.T) {
	c, _ := NewCharacter("Test")
	err := c.GainBackground("Acolite")
	if err == nil {
		t.Error("Expected error but got nil.")
	}
	bg := "Acolyte"
	err = c.GainBackground(bg)
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if c.Background != bg {
		t.Errorf("Expected background to be %v,but got: %v", bg, c.Background)
	}
	if !c.Proficiencies["Insight"] {
		t.Error("Expected c.Proficiences['Insight'] to be true")
	}
	if !c.Proficiencies["Religion"] {
		t.Error("Expected c.Proficiences['Religion'] to be true")
	}
}

func TestSetUnarmoredAC(t *testing.T) {
	c, _ := NewCharacter("Test")
	if c.AC != 10 {
		t.Errorf("Expected AC to be 10 but got:%d", c.AC)
	}
	dexMod, err := c.GetAbilityMod("Dexterity")
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if dexMod != -1 {
		t.Errorf("Expected dex mod to be -1, but got: %d", dexMod)
	}
	c.SetUnarmoredAC()
	if c.AC != 9 {
		t.Errorf("Expected AC to be 9, but got: %d", c.AC)
	}
	c.ClassName = "Barbarian"
	conMod, err := c.GetAbilityMod("Constitution")
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if conMod != -1 {
		t.Errorf("Expected dex mod to be -1, but got: %d", conMod)
	}
	c.SetUnarmoredAC()
	if c.AC != 8 {
		t.Errorf("Expected AC to be 8, but got: %d", c.AC)
	}
}

func TestSetMaxHP(t *testing.T) {
	c, _ := NewCharacter("Test")
	c.HitDie = 12
	c.SetMaxHP()
	conMod, err := c.GetAbilityMod("Constitution")
	if err != nil {
		t.Errorf("Expected error to be nil, but got: %v", err)
	}
	if conMod != -1 {
		t.Errorf("Expected con mod to be -1, but got: %d", conMod)
	}
	if c.HP != 11 {
		t.Errorf("Expected HP to be 11, but got %d", c.HP)
	}
	_ = c.AssignStat("Constitution", 20)
	c.SetMaxHP()
	if c.HP != 17 {
		t.Errorf("Expected HP to be 17, but got %d", c.HP)
	}
	c.HPBonus = 5
	c.SetMaxHP()
	if c.HP != 22 {
		t.Errorf("Expected HP to be 22, but got %d", c.HP)
	}
}
