package character

type Background struct {
	Name          string
	Proficiencies []string
}

var Backgrounds = map[string]Background{
	"Acolyte": {
		Name:          "Acolyte",
		Proficiencies: []string{"Insight", "Religion"},
	},
	"Charlatan": {
		Name:          "Charlatan",
		Proficiencies: []string{"Deception", "Sleight of Hand"},
	},
	"Criminal": {
		Name:          "Criminal",
		Proficiencies: []string{"Deception", "Stealth"},
	},
	"Entertainer": {
		Name:          "Entertainer",
		Proficiencies: []string{"Acrobatics", "Performance"},
	},
	"Folk Hero": {
		Name:          "Folk Hero",
		Proficiencies: []string{"Animal Handling", "Survival"},
	},
	"Guild Artisan": {
		Name:          "Guild Artisan",
		Proficiencies: []string{"Insight", "Persuasion"},
	},
	"Hermit": {
		Name:          "Hermit",
		Proficiencies: []string{"Medicine", "Religion"},
	},
	"Noble": {
		Name:          "Noble",
		Proficiencies: []string{"History", "Persuasion"},
	},
	"Outlander": {
		Name:          "Outlander",
		Proficiencies: []string{"Athletics", "Survival"},
	},
	"Sage": {
		Name:          "Sage",
		Proficiencies: []string{"Arcana", "History"},
	},
	"Sailor": {
		Name:          "Sailor",
		Proficiencies: []string{"Athletics", "Perception"},
	},
	"Soldier": {
		Name:          "Soldier",
		Proficiencies: []string{"Athletics", "Intimidation"},
	},
	"Urchin": {
		Name:          "Urchin",
		Proficiencies: []string{"Sleight of Hand", "Stealth"},
	},
}
