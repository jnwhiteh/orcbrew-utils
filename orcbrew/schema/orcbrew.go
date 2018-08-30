package schema

//go:generate go run internal/gen_modifiers/main.go -output modifiers.go

// OrcbrewExportAll is a map from source name to OrcbrewSource, used by the
// "Export All" functionality in Orcbrew
type OrcbrewExportAll map[string]OrcbrewSource

// OrcbrewSource is the data structure used for an individual source that has
// been exported.
type OrcbrewSource struct {
	Languages   map[string]LanguageConfig
	Classes     map[string]ClassConfig
	Subclasses  map[string]SubclassConfig
	Monsters    map[string]MonsterConfig
	Feats       map[string]FeatConfig
	Backgrounds map[string]BackgroundConfig
	Invocations map[string]InvocationConfig
	Subraces    map[string]SubraceConfig
	Spells      map[string]SpellConfig
	Encounters  map[string]EncounterConfig
	Selections  map[string]SelectionConfig
	Races       map[string]RaceConfig
}

// LanguageConfig defines a language that can be spoken/written/read
type LanguageConfig struct {
	Key         string `json:"key"`
	OptionPack  string `json:"option-pack"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ClassConfig contains the configuration for a character class
type ClassConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name string `json:"name"`

	HitDie int `json:"hit-die"` // the type of hit die

	SubclassLevel int    `json:"subclass-level,omitempty"` // the level at which subclasses are unlocked
	SubclassTitle string `json:"subclass-title,omitempty"` // the name of the subclass paths/types (path, domain, etc.)

	// The spellcasting configuration for this class (if any)
	Spellcasting *SpellcastingConfig `json:"spellcasting"`

	AbilityIncreaseLevels []int `json:"ability-increase-levels"` // the levels at which abilities increase

	LevelModifiers  LevelModifierList `json:"level-modifiers"`            // modifiers that apply to the class (by level)
	LevelSelections []LevelSelection  `json:"level-selections,omitempty"` // options/selections available upon taking the class (by level)

	Profs  *ClassProficiencies `json:"profs,omitempty"` // proficiencies in skills and saving throws
	Traits []LevelTrait        `json:"traits"`
}

// SubclassConfig contains the configuration for a character subclass
type SubclassConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name  string `json:"name"`
	Class string `json:"class"` // the class this is a subclass for

	// The spellcasting configuration for this class (if any)
	Spellcasting *SpellcastingConfig `json:",omitempty"`

	LevelModifiers  LevelModifierList `json:"level-modifiers"`            // modifiers that apply to the class (by level)
	LevelSelections []LevelSelection  `json:"level-selections,omitempty"` // options/selections available upon taking the class (by level)

	Profs  *ClassProficiencies `json:"profs,omitempty"` // proficiencies in skills and saving throws
	Traits []LevelTrait        `json:"traits"`
}

// SpellcastingConfig defines the spellcasting rules and progression for class
type SpellcastingConfig struct {
	SpellListKw string  `json:"spell-list-kw,omitempty"` // the key for the list of spells available
	LevelFactor int     `json:"level-factor"`            // what level is spellcasting unlocked
	Ability     Ability `json:"ability"`

	// The method used for computing how many spells are known
	KnownMode string `json:"known-mode"` // "all" or "schedule"

	// TODO: Determine if/what/how/when/why this should be used
	// Appears to contain the number of spells unlocked at each level
	SpellsKnown map[int]int `json:"spells-known"`

	// The list of spell options available at each level
	SpellList map[int][]string `json:"spell-list,omitempty"`
}

// SpellWithAbility defines a known spell paired with the spellcasting ability
// used for that spell, which may differ from the core spellcasting ability
// for the character, or for non-spellcasting classes.
type SpellWithAbility struct {
	Ability Ability `json:"ability"` // the ability used as spellcasting ability for this spell
	Key     string  `json:"key"`     // the key for the spell
}

// LevelSelection defines an option/selection that gets unlocked at a given
// level, with potentially more being unlocked at subsequent levels.
type LevelSelection struct {
	Level int    `json:"level,omitempty"` // the level the selection is unlocked
	Num   int    `json:"num"`             // the number of selections allowed
	Type  string `json:"type"`            // the key/type of selections
}

// ClassProficiencies defines a limited set of proficiencies/options for a class
type ClassProficiencies struct {
	Save                  map[Ability]bool `json:"save,omitempty"`          // saving throws
	SkillExpertiseOptions *SkillOptions    `json:"skill-expertise-options"` // skill expertise (double proficiency)
	SkillOptions          *SkillOptions    `json:"skill-options"`           // skill proficiency
}

// SkillOptions enable choosing from a set list of skills
type SkillOptions struct {
	Choose  int            `json:"choose,omitempty"` // how many to choose
	Options map[Skill]bool `json:"options"`          // skills to choose from (empty means no options)
}

// LevelTrait desscribes a trait that is granted at a given level
type LevelTrait struct {
	Level       int    `json:"level,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MonsterConfig defines a new monster type
type MonsterConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name             string                      `json:"name"`
	Description      string                      `json:"description"`
	LegendaryActions *LegendaryActionDescription `json:"legendary-actions"`

	// Ability scores
	Str int `json:"str"`
	Dex int `json:"dex"`
	Con int `json:"con"`
	Int int `json:"int"`
	Wis int `json:"wis"`
	Cha int `json:"cha"`

	HitPoints  *HitDieCount `json:"hit-points"`
	Speed      string       `json:"speed"`
	Alignment  string       `json:"alignment"`
	Size       Size         `json:"size"`
	ArmorClass int          `json:"armor-class"`
	Type       string       `json:"type"`

	Skills       map[Skill]int   `json:"skills"`
	SavingThrows map[Ability]int `json:"saving-throws"`

	Challenge float32 `json:"challenge"`

	Props  *MonsterProperties `json:"props"`
	Traits []MonsterTrait     `json:"traits"`
}

type LegendaryActionDescription struct {
	Description string `json:"description"`
}

// HitDieCount defines a number of a specific type of dice
type HitDieCount struct {
	DieCount int `json:"die-count"` // the number of dice
	Die      int `json:"die"`       // the type of die
}

// MonsterProperties contains some core configurable properties for monsters
type MonsterProperties struct {
	Language            map[string]bool    `json:"language"`
	DamageResistance    map[Damage]bool    `json:"damage-resistance"`
	DamageImmunity      map[Damage]bool    `json:"damage-immunity"`
	DamageVulnerability map[Damage]bool    `json:"damage-vulnerability"`
	ConditionImmunity   map[Condition]bool `json:"condition-immunity"`
}

// MonsterTrait defines a feature/trait for a monster, which distinguishes
// between normal and legendary actions.
type MonsterTrait struct {
	Type        MonsterTraitAction `json:"type"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
}

// FeatConfig defines a new feat that can be taken by characters
type FeatConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name        string `json:"name"`
	Description string `json:"description"`

	// TODO: Should be Ability, but "saves?" can be used to indicate that
	// they also gain proficiency on saving throws with the given abilities
	AbilityIncreases []string `json:"ability-increases"`

	// Pre-requisites for taking this feat. Values seem to be:
	// spellcasting, str, dex, con, int, wis, cha, light, medium, heavy
	Prereqs []string `json:"prereqs"`

	// Racial pre-requisites for taking this feat
	PathPrereqs FeatPathPrereqs `json:"path-prereqs"`

	// TODO: Find a way to properly type this...
	Props map[string]interface{} `json:"props"`
}

// FeatPathPrereqs contains any racial pre-requisites for taking a feat
type FeatPathPrereqs struct {
	Race map[string]bool `json:"race"`
}

// BackgroundConfig defines a new character background
type BackgroundConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name string `json:"name"`

	Equipment map[string]int   `json:"equipment"`
	Treasure  map[Currency]int `json:"treasure"`

	// TODO: Borked
	EquipmentChoices []BackgroundEquipmentOption `json:"equipment-choices"`
	Profs            *BackgroundProfs            `json:"profs"`

	Traits []BackgroundTrait `json:"traits"`
}

type BackgroundProfs struct {
	LanguageOptions *BackgroundLanguageOptions `json:"language-options"`
	Skill           map[Skill]bool             `json:"skill"`

	// TODO: Includes vehicles?
	Tool map[string]bool `json:"tool"`

	// How many tools do you have of each type?
	ToolOptions *BackgroundToolOptions `json:"tool-options"`
}

type BackgroundEquipmentOption struct {
	Name    string         `json:"name"`
	Options map[string]int `json:"options"`
}

type BackgroundLanguageOptions struct {
	Choose  int                            `json:"choose"`
	Options BackgroundLanguageOptionConfig `json:"options"`
}

type BackgroundLanguageOptionConfig struct {
	Any bool `json:"any"`
}
type BackgroundToolOptions struct {
	MusicalInstrument int `json:"musical-instrument"`
	GamingSet         int `json:"gaming-set"`
}

type BackgroundTrait struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// InvocationConfig describes an option for Eldritch Invocations
type InvocationConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name        string `json:"name"`
	Description string `json:"description"`
}

// SubraceConfig defines a new sub-race option
type SubraceConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name       string            `json:"name"`
	Race       string            `json:"race"`
	Abilities  map[Ability]int   `json:"abilities,omitempty"`
	Languages  []string          `json:"languages,omitempty"`
	Darkvision int               `json:"darkvision,omitempty"`
	Size       Size              `json:"size,omitempty"`
	Speed      int               `json:"speed,omitempty"`
	Spells     []RaceSpellConfig `json:"spells,omitempty"`

	Props *SubraceProperties `json:"props,omitempty"`
	Profs *RaceProficiencies `json:"profs,omitempty"`

	Traits []LevelTrait `json:"traits"`
}

type SubraceProperties struct {
	FlyingSpeed          int                `json:"flying-speed,omitempty"`
	SkillProficiency     map[Skill]bool     `json:"skill-prof,omitempty"`
	DamageImmunity       map[Damage]bool    `json:"damage-immunity,omitempty"`
	DamageResistance     map[Damage]bool    `json:"damage-resistance,omitempty"`
	ArmorProficiency     map[Armor]bool     `json:"armor-prof,omitempty"`
	WeaponProficiency    map[string]bool    `json:"weapon-prof,omitempty"`
	SavingThrowAdvantage map[Condition]bool `json:"saving-throw-advantage,omitempty"`
	Language             map[string]bool    `json:"language,omitempty"`
	MaxHpBonus           int                `json:"max-hp-bonus,omitempty"`
}

type ToolOptions struct {
	Choose  int             `json:"choose"`
	Options map[string]bool `json:"options,omitempty"`
}

type AbilityIncreaseOptions struct {
	Choose  int              `json:"choose"`
	Amount  int              `json:"amount"`
	Options map[Ability]bool `json:"options,omitempty"`
}

type FeatOptions struct {
	Choose  int             `json:"choose"`
	Options map[string]bool `json:"options,omitempty"`
}

type SpellOptions struct {
	Choose    int             `json:"choose"`
	Options   map[string]bool `json:"options"`              // A list of specific options
	SpellList []string        `json:"spell-list,omitempty"` // The name of a pre-defined spell list
}

type SpellConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name        string `json:"name"`
	Description string `json:"description"`

	Level       int              `json:"level"`
	School      string           `json:"school"`
	Duration    string           `json:"duration"`
	Components  *SpellComponents `json:"components"`
	Ritual      bool             `json:"ritual"`
	AttackRoll  bool             `json:"attack-roll?"` // Spelling incorrect in orcbrew
	CastingTime string           `json:"casting-time"`
	SpellLists  map[string]bool  `json:"spell-lists"`
	Range       string           `json:"range"`
}

type SpellComponents struct {
	MaterialComponent string `json:"material-component,omitempty"`
	Verbal            bool   `json:"verbal"`
	Material          bool   `json:"material"`
	Somatic           bool   `json:"somatic"`
}

type EncounterConfig struct {
	Key        string `json:"key"`
	OptionPack string `json:"option-pack"`

	Name      string              `json:"name"`
	Creatures []EncounterCreature `json:"creatures"`
}

type EncounterCreature struct {
	Type string `json:"type"` // always "monster" except for NPCs

	// TODO: Support NPCs?
	Creature EncounterCreatureConfig `json:"creature"`
}

type EncounterCreatureConfig struct {
	Num     int    `json:"num"`
	Monster string `json:"monster"`
}

type SelectionConfig struct {
	Key        string            `json:"key"`
	Name       string            `json:"name"`
	OptionPack string            `json:"option-pack"`
	Options    []SelectionOption `json:"options"`
}

type SelectionOption struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RaceConfig struct {
	Key        string `json:"key"`
	Name       string `json:"name"`
	OptionPack string `json:"option-pack"`

	Abilities  map[Ability]int `json:"abilities"`
	Languages  []string        `json:"languages"`
	Darkvision int             `json:"darkvision"`
	Size       Size            `json:"size"`
	Speed      int             `json:"speed"`

	Spells []RaceSpellConfig  `json:"spells"`
	Props  *RaceProperties    `json:"props"`
	Profs  *RaceProficiencies `json:"profs"`

	Traits []LevelTrait `json:"traits"`
}

type RaceSpellConfig struct {
	Value SpellWithAbilityLevel `json:"value"`
}

type SpellWithAbilityLevel struct {
	Level   int     `json:"level,omitempty"`
	Key     string  `json:"key"`
	Ability Ability `json:"ability"`
}

type RaceProperties struct {
	FlyingSpeed          int                `json:"flying-speed,omitempty"`
	SkillProficiency     map[Skill]bool     `json:"skill-prof,omitempty"`
	DamageImmunity       map[Damage]bool    `json:"damage-immunity,omitempty"`
	DamageResistance     map[Damage]bool    `json:"damage-resistance,omitempty"`
	ArmorProficiency     map[Armor]bool     `json:"armor-prof,omitempty"`
	WeaponProficiency    map[string]bool    `json:"weapon-prof,omitempty"`
	SavingThrowAdvantage map[Condition]bool `json:"saving-throw-advantage,omitempty"`
	Language             map[string]bool    `json:"language,omitempty"`
	MaxHpBonus           int                `json:"max-hp-bonus,omitempty"`
	LizardfolkAC         bool               `json:"lizardfolk-ac"`       // TODO: What is this?
	TortleAC             bool               `json:"tortle-ac,omitempty"` // TODO: What is this?
}

type RaceProficiencies struct {
	LanguageOptions          *RaceLanguageOptions          `json:"language-options,omitempty"`
	ToolOptions              map[string]bool               `json:"tool,omitempty"`
	SkillOptions             *RaceSkillOptions             `json:"skill-options,omitempty"`
	WeaponProficiencyOptions *RaceWeaponProficiencyOptions `json:"weapon-proficiency-options,omitempty"`
}

type RaceLanguageOptions struct {
	// TODO: This seems to be false? Huh?
	Options map[string]bool `json:"options"`
}

type RaceWeaponProficiencyOptions struct {
	Options map[string]bool `json:"options"`
}

type RaceSkillOptions struct {
	Choose  int            `json:"choose,omitempty"`
	Options map[Skill]bool `json:"options"`
}
