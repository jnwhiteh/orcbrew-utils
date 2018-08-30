package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

var rawSourceJSON string
var testSource OrcbrewSource
var loaded bool

func getTestDataJSON(t *testing.T, key string) string {
	s := fmt.Sprintf(`(?s)\n  "%s": ({.*?\n  }),?\n`, key)
	re := regexp.MustCompile(s)
	matches := re.FindStringSubmatch(rawSourceJSON)

	if len(matches) <= 0 {
		return ""
	}

	return strings.TrimSpace(matches[1])
}

func getTestData(t *testing.T) OrcbrewSource {
	if !loaded {
		filename := "example.json"
		file, err := os.Open(filename)
		if err != nil {
			t.Fatalf("Error loading %s: %s", filename, err)
		}
		defer file.Close()

		buf, err := ioutil.ReadAll(file)
		rawSourceJSON = string(buf)

		err = json.Unmarshal(buf, &testSource)
		if err != nil {
			t.Fatal(err)
		}
		loaded = true
	}

	return testSource
}

func TestLanguages(t *testing.T) {
	source := getTestData(t)

	if len(source.Languages) != 1 {
		t.Errorf("Expected 1 languages got %v", len(source.Languages))
	}

	result := source.Languages["pig-latin"]
	expected := LanguageConfig{
		Key:         "pig-latin",
		OptionPack:  "Test",
		Name:        "Pig latin",
		Description: "Pig latin",
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestClasses(t *testing.T) {
	source := getTestData(t)

	if len(source.Classes) != 2 {
		t.Errorf("Expected 2 classes, got %v", len(source.Classes))
	}

	result := source.Classes["myclass"]
	expected := ClassConfig{
		Key: "myclass",

		LevelModifiers: LevelModifierList{
			&ModifierDamageImmunity{Value: Bludgeoning},
			&ModifierSavingThrowAdvantage{Level: 2, Value: Blinded},
			&ModifierSpell{Value: SpellWithAbility{Key: "druidcraft", Ability: Strength}},
		},

		Name:       "MyClass",
		OptionPack: "Test",

		SubclassLevel: 2,
		SubclassTitle: "MySubclassTitle",
		LevelSelections: []LevelSelection{
			LevelSelection{Num: 2, Type: "my-selection-thingy"},
		},
		Spellcasting: &SpellcastingConfig{
			Ability:     Charisma,
			KnownMode:   "schedule",
			SpellListKw: "bard",
			LevelFactor: 3,
			SpellsKnown: map[int]int{
				3:  3,
				4:  1,
				7:  1,
				8:  1,
				10: 1,
				11: 1,
				13: 1,
				14: 1,
				16: 1,
				19: 1,
				20: 1,
			},
		},
		AbilityIncreaseLevels: []int{4, 8, 12, 16, 19},
		Profs: &ClassProficiencies{
			Save: map[Ability]bool{Strength: true},
			SkillExpertiseOptions: &SkillOptions{
				Choose: 2,
				Options: map[Skill]bool{
					Acrobatics:     true,
					AnimalHandling: true,
					Arcana:         true,
				},
			},
			SkillOptions: &SkillOptions{
				Options: map[Skill]bool{
					Acrobatics: true,
				},
			},
		},
		HitDie: 6,
		Traits: []LevelTrait{
			LevelTrait{
				Level:       2,
				Name:        "MyClassTrait",
				Description: "This is a class trait",
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestSubClasses(t *testing.T) {
	source := getTestData(t)

	if len(source.Subclasses) != 1 {
		t.Errorf("Expected 1 subclass, got %v", len(source.Subclasses))
	}

	result := source.Subclasses["mysubclass"]
	expected := SubclassConfig{
		Key: "mysubclass",

		LevelModifiers: LevelModifierList{
			&ModifierArmorProficiency{Value: LightArmor},
			&ModifierSkillProficiency{Level: 3, Value: Athletics},
		},

		Name:       "MySubClass",
		OptionPack: "Test",

		LevelSelections: []LevelSelection{
			LevelSelection{Level: 3, Num: 2, Type: "my-selection-thingy"},
		},
		Class: "barbarian",
		Profs: &ClassProficiencies{
			SkillExpertiseOptions: &SkillOptions{
				Choose: 2,
				Options: map[Skill]bool{
					Acrobatics:     true,
					AnimalHandling: true,
					Arcana:         true,
				},
			},
			SkillOptions: &SkillOptions{
				Options: map[Skill]bool{
					Acrobatics: false,
				},
			},
		},
		Traits: []LevelTrait{
			LevelTrait{
				Name:        "My subclass trait",
				Description: "This is a subclass trait",
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestMonsters(t *testing.T) {
	source := getTestData(t)

	if len(source.Monsters) != 1 {
		t.Errorf("Expected 1 monster, got %v", len(source.Subclasses))
	}

	result := source.Monsters["mymonster"]
	expected := MonsterConfig{
		Key:        "mymonster",
		OptionPack: "Test",

		Name:        "MyMonster",
		Description: "Description of MyMonster",
		LegendaryActions: &LegendaryActionDescription{
			Description: "I have some legendary actions",
		},

		Str: 10,
		Dex: 10,
		Con: 10,
		Int: 10,
		Wis: 10,
		Cha: 10,

		Alignment:  "neutral",
		Speed:      "30 ft.",
		HitPoints:  &HitDieCount{Die: 8, DieCount: 1},
		Type:       "aberration",
		Size:       Large,
		ArmorClass: 10,
		Skills: map[Skill]int{
			Performance: 2,
			Survival:    2,
			History:     4,
			Stealth:     2,
		},
		SavingThrows: map[Ability]int{
			Strength:     0,
			Dexterity:    0,
			Constitution: 1,
			Intelligence: 1,
			Wisdom:       0,
			Charisma:     1,
		},
		Challenge: 1,
		Props: &MonsterProperties{
			DamageVulnerability: map[Damage]bool{Acid: true},
			DamageImmunity:      map[Damage]bool{Acid: true},
			Language: map[string]bool{
				"elvish":  true,
				"abyssal": true,
			},
			ConditionImmunity: map[Condition]bool{Blinded: true},
			DamageResistance: map[Damage]bool{
				Lightning: false,
				Slashing:  false,
				Traps:     true,
			},
		},
		Traits: []MonsterTrait{
			MonsterTrait{
				Type:        "action",
				Name:        "Multi-attack",
				Description: "Beat someone up.",
			},
			MonsterTrait{
				Type:        "legendary-action",
				Name:        "Dragon breath",
				Description: "I bathe you in fire!!!",
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestFeats(t *testing.T) {
	source := getTestData(t)

	if len(source.Feats) != 1 {
		t.Errorf("Expected 1 feat, got %v", len(source.Feats))
	}

	result := source.Feats["myfeat"]
	expected := FeatConfig{
		Description: "This is my feat",
		PathPrereqs: FeatPathPrereqs{
			Race: map[string]bool{
				"elf":        true,
				"half-orc":   true,
				"human":      true,
				"gnome":      true,
				"myrace":     true,
				"tiefling":   true,
				"dragonborn": true,
				"dwarf":      true,
				"halfling":   true,
				"half-elf":   true,
			},
		},
		Prereqs: []string{
			"str",
			"con",
			"dex",
			"medium",
			"wis",
			"heavy",
			"spellcasting",
			"int",
			"cha",
			"light",
		},
		Key:        "myfeat",
		Name:       "MyFeat",
		OptionPack: "Test",
		AbilityIncreases: []string{
			"str",
			"saves?",
			"con",
			"int",
			"dex",
			"wis",
			"cha",
		},
		Props: map[string]interface{}{
			"max-hp-bonus":                 float64(2),
			"passive-investigation-5":      true,
			"magic-novice":                 true,
			"medium-armor-max-dex-3":       true,
			"speed":                        float64(15),
			"ritual-casting":               true,
			"passive-perception-5":         true,
			"initiative":                   float64(2),
			"weapon-prof-choice":           float64(3),
			"two-weapon-any-one-handed":    true,
			"saving-throw-advantage-traps": true,
			"language-choice":              float64(2),
			"armor-prof": map[string]interface{}{
				"medium":  true,
				"heavy":   true,
				"light":   true,
				"shields": true,
			},
			"improvised-weapons-prof": true,
			"attack-spell":            true,
			"two-weapon-ac-1":         true,
			"skill-tool-choice":       float64(2),
			"medium-armor-stealth":    true,
			"skill-prof-or-expertise": map[string]interface{}{
				"religion":        true,
				"persuasion":      true,
				"investigation":   true,
				"acrobatics":      true,
				"performance":     true,
				"perception":      true,
				"sleight-of-hand": true,
				"survival":        true,
				"history":         true,
				"animal-handling": true,
				"nature":          true,
				"deception":       true,
				"intimidation":    true,
				"arcana":          true,
				"athletics":       true,
				"insight":         true,
				"medicine":        true,
				"stealth":         true,
			},
			"damage-resistance": map[string]interface{}{
				"fire":        true,
				"acid":        true,
				"psychic":     true,
				"force":       true,
				"bludgeoning": true,
				"radiant":     true,
				"lightning":   true,
				"slashing":    true,
				"piercing":    true,
				"thunder":     true,
				"cold":        true,
				"traps":       true,
				"poison":      true,
				"necrotic":    true,
			},
			"tool-prof-or-expertise": map[string]interface{}{
				"cartographers-tools":    true,
				"painters-supplies":      true,
				"navigators-tools":       true,
				"glassblowers-tools":     true,
				"flute":                  true,
				"poisoners-tools":        true,
				"dice-set":               true,
				"horn":                   true,
				"herbalism-kit":          true,
				"dulcimer":               true,
				"disguise-kit":           true,
				"masons-tools":           true,
				"land-vehicles":          true,
				"thieves-tools":          true,
				"jewelers-tools":         true,
				"leatherworkers-tools":   true,
				"smiths-tools":           true,
				"drum":                   true,
				"cobblers-tools":         true,
				"potters-tools":          true,
				"dragonchess-set":        true,
				"playing-card-set":       true,
				"brewers-supplies":       true,
				"three-dragon-ante-set":  true,
				"forgery-kit":            true,
				"pan-flute":              true,
				"bagpipes":               true,
				"woodcarvers-tools":      true,
				"carpenters-tools":       true,
				"tinkers-tools":          true,
				"alchemists-supplies":    true,
				"water-vehicles":         true,
				"weavers-tools":          true,
				"shawm":                  true,
				"cooks-utensils":         true,
				"lute":                   true,
				"calligraphers-supplies": true,
				"lyre":                   true,
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestBackgrounds(t *testing.T) {
	source := getTestData(t)

	if len(source.Backgrounds) != 1 {
		t.Errorf("Expected 1 background, got %v", len(source.Backgrounds))
	}

	result := source.Backgrounds["mybackground"]
	expected := BackgroundConfig{
		Key:  "mybackground",
		Name: "MyBackground",
		Equipment: map[string]int{
			"clothes-common":      1,
			"abacus":              1,
			"disguise-kit":        1,
			"amulet":              1,
			"alchemists-supplies": 1,
		},
		Treasure: map[Currency]int{
			"gp": 10,
		},
		EquipmentChoices: []BackgroundEquipmentOption{
			BackgroundEquipmentOption{
				Name: "Musical Instruments",
				Options: map[string]int{
					"flute":     1,
					"horn":      1,
					"dulcimer":  1,
					"viol":      1,
					"drum":      1,
					"pan-flute": 1,
					"bagpipes":  1,
					"shawm":     1,
					"lute":      1,
					"lyre":      1,
				},
			},
			BackgroundEquipmentOption{
				Name: "Artisan's Tools",
				Options: map[string]int{
					"cartographers-tools":    1,
					"painters-supplies":      1,
					"glassblowers-tools":     1,
					"masons-tools":           1,
					"jewelers-tools":         1,
					"leatherworkers-tools":   1,
					"smiths-tools":           1,
					"cobblers-tools":         1,
					"potters-tools":          1,
					"brewers-supplies":       1,
					"woodcarvers-tools":      1,
					"carpenters-tools":       1,
					"tinkers-tools":          1,
					"alchemists-supplies":    1,
					"weavers-tools":          1,
					"cooks-utensils":         1,
					"calligraphers-supplies": 1,
				},
			},
		},
		OptionPack: "Test",
		Profs: &BackgroundProfs{
			LanguageOptions: &BackgroundLanguageOptions{
				Choose: 1,
				Options: BackgroundLanguageOptionConfig{
					Any: true,
				},
			},
			Skill: map[Skill]bool{
				Acrobatics: true,
			},
			Tool: map[string]bool{
				"disguise-kit":        true,
				"alchemists-supplies": true,
				"water-vehicles":      true,
			},
			ToolOptions: &BackgroundToolOptions{
				GamingSet:         2,
				MusicalInstrument: 1,
			},
		},
		Traits: []BackgroundTrait{
			BackgroundTrait{
				Name:        "Some trait",
				Description: "This is a trait",
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestInvocations(t *testing.T) {
	source := getTestData(t)

	if len(source.Invocations) != 1 {
		t.Errorf("Expected 1 invocation, got %v", len(source.Invocations))
	}

	result := source.Invocations["myeldritchinvocation"]
	expected := InvocationConfig{
		Key:         "myeldritchinvocation",
		OptionPack:  "Test",
		Name:        "MyEldritchInvocation",
		Description: "I don't know why this is a separate concept, instead of being a selection itself.",
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestSubraces(t *testing.T) {
	source := getTestData(t)

	if len(source.Subraces) != 1 {
		t.Errorf("Expected 1 subrace, got %v", len(source.Invocations))
	}

	result := source.Invocations["myeldritchinvocation"]
	expected := InvocationConfig{
		Key:         "myeldritchinvocation",
		OptionPack:  "Test",
		Name:        "MyEldritchInvocation",
		Description: "I don't know why this is a separate concept, instead of being a selection itself.",
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}

}

func TestSpells(t *testing.T) {
	source := getTestData(t)

	if len(source.Spells) != 1 {
		t.Errorf("Expected 1 spell, got %v", len(source.Spells))
	}

	result := source.Spells["myspell"]
	expected := SpellConfig{
		Key:        "myspell",
		OptionPack: "Test",

		Name:        "MySpell",
		Description: "Casting a spell",

		School:   "necromancy",
		Duration: "1 hour",
		Level:    0,
		Components: &SpellComponents{
			MaterialComponent: "A pinch of salt",
			Verbal:            true,
			Material:          true,
			Somatic:           true,
		},
		Ritual:      true,
		AttackRoll:  true,
		CastingTime: "1 action",
		SpellLists: map[string]bool{
			"cleric":   true,
			"druid":    true,
			"ranger":   true,
			"sorcerer": true,
			"wizard":   true,
			"warlock":  true,
			"paladin":  true,
			"bard":     true,
		},
		Range: "10 ft.",
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestEncounters(t *testing.T) {
	source := getTestData(t)

	if len(source.Encounters) != 1 {
		t.Errorf("Expected 1 encounter, got %v", len(source.Encounters))
	}

	result := source.Encounters["goblin-ambush"]
	expected := EncounterConfig{
		Key:        "goblin-ambush",
		OptionPack: "Test",
		Name:       "Goblin Ambush",
		Creatures: []EncounterCreature{
			EncounterCreature{
				Type: "monster",
				Creature: EncounterCreatureConfig{
					Monster: "goblin",
					Num:     3,
				},
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}
func TestSelections(t *testing.T) {
	source := getTestData(t)

	if len(source.Selections) != 1 {
		t.Errorf("Expected 1 selection, got %v", len(source.Encounters))
	}

	result := source.Selections["my-selection-thingy"]
	expected := SelectionConfig{
		Key:        "my-selection-thingy",
		OptionPack: "Test",

		Name: "My selection thingy",
		Options: []SelectionOption{
			SelectionOption{
				Name:        "Fire thingy",
				Description: "I have a fire thingy",
			},
			SelectionOption{
				Name:        "Frost thingy",
				Description: "I've got another frost thingy",
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}

func TestRaces(t *testing.T) {
	source := getTestData(t)

	if len(source.Races) != 1 {
		t.Errorf("Expected 1 race, got %v", len(source.Races))
	}

	result := source.Races["myrace"]
	expected := RaceConfig{
		Key:        "myrace",
		OptionPack: "Test",

		Name:       "MyRace",
		Speed:      30,
		Darkvision: 60,
		Spells: []RaceSpellConfig{
			RaceSpellConfig{
				Value: SpellWithAbilityLevel{
					Key:     "acid-splash",
					Ability: Charisma,
				},
			},
			RaceSpellConfig{
				Value: SpellWithAbilityLevel{
					Key:     "aid",
					Ability: Charisma,
					Level:   2,
				},
			},
		},
		Abilities: map[Ability]int{
			Strength:     1,
			Constitution: 1,
		},
		Size: Medium,
		Profs: &RaceProficiencies{
			LanguageOptions: &RaceLanguageOptions{
				Options: map[string]bool{
					"common": false,
				},
			},
			ToolOptions: map[string]bool{
				"bagpipes": true,
			},
			WeaponProficiencyOptions: &RaceWeaponProficiencyOptions{
				Options: map[string]bool{
					"shortbow":       true,
					"crossbow-light": true,
					"dart":           false,
				},
			},
			SkillOptions: &RaceSkillOptions{
				Options: map[Skill]bool{
					Acrobatics: false,
				},
			},
		},
		Languages: []string{"Abyssal"},
		Props: &RaceProperties{
			SkillProficiency: map[Skill]bool{
				Acrobatics: true,
			},
			DamageImmunity: map[Damage]bool{
				Acid: true,
			},
			WeaponProficiency: map[string]bool{
				"simple":       true,
				"light-hammer": true,
			},
			ArmorProficiency: map[Armor]bool{
				LightArmor: true,
			},
			LizardfolkAC: false,
			TortleAC:     true,
			FlyingSpeed:  10,
			DamageResistance: map[Damage]bool{
				Traps: true,
			},
		},
		Traits: []LevelTrait{
			LevelTrait{
				Name:        "Awesome trait",
				Description: "This is truly an awesome trait",
			},
		},
	}

	if diff := deep.Equal(result, expected); diff != nil {
		t.Error(diff)
	}
}
