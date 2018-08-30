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
	Key        string
	OptionPack string `json:"option-pack"`

	Name       string
	Race       string
	Abilities  map[Ability]int
	Languages  []string
	Darkvision int
	Size       Size
	Speed      int

	Props *SubraceProperties
	Profs *SubraceProficiencies

	Traits []LevelTrait
}

type SubraceProperties struct {
	FlyingSpeed          int                `json:"flying-speed"`
	SkillProficiency     map[Skill]bool     `json:"skill-prof"`
	DamageImmunity       map[Damage]bool    `json:"damage-immunity"`
	DamageResistance     map[Damage]bool    `json:"damage-resistance"`
	SavingThrowAdvantage map[Condition]bool `json:"saving-throw-advantage"`
	ArmorProficiency     map[Armor]bool     `json:"armor-prof"`
	WeaponProficiency    map[string]bool    `json:"weapon-prof"`
	HitPointLevelBonus   int                `json:"hit-point-level-bonus"`
	ToolProficiency      map[string]bool    `json:"tool-prof"`
	LizardfolkAC         bool               `json:"lizardfolk-ac"` // TODO: What is this?
	TortleAC             bool               `json:"tortle-ac"`     // TODO: What is this?
}

type SubraceProficiencies struct {
	LanguageOptions        *LanguageConfig         `json:"language-options"`
	SkillOptions           *SkillOptions           `json:"skill-options"`
	ToolOptions            *ToolOptions            `json:"tool-options"`
	AbilityIncreaseOptions *AbilityIncreaseOptions `json:"ability-increase-options"`
	FeatOptions            *FeatOptions            `json:"feat-options"`
	SpellOptions           *SpellOptions           `json:"spell-options"`
}

type ToolOptions struct {
	Choose  int
	Options map[string]bool
}

type AbilityIncreaseOptions struct {
	Choose  int
	Amount  int
	Options map[Ability]bool
}

type FeatOptions struct {
	Choose  int
	Options map[string]bool
}

type SpellOptions struct {
	Choose    int
	Options   map[string]bool // A list of specific options
	SpellList []string        // The name of a pre-defined spell list
}

type SpellConfig struct {
	Key        string
	OptionPack string `json:"option-pack"`

	Name        string
	Description string

	Level       int
	School      string
	Duration    string
	Components  *SpellComponents
	Ritual      bool
	AttackRoll  bool            `json:"attack-roll?"` // Spelling incorrect in orcbrew
	CastingTime string          `json:"casting-time"`
	SpellLists  map[string]bool `json:"spell-lists"`
	Range       string
}

type SpellComponents struct {
	MaterialComponent string `json:"material-component,omitempty"`
	Verbal            bool
	Material          bool
	Somatic           bool
}

type EncounterConfig struct {
	Key        string
	OptionPack string `json:"option-pack"`

	Name      string
	Creatures []EncounterCreature
}

type EncounterCreature struct {
	Type string // always "monster" except for NPCs

	// TODO: Support NPCs?
	Creature EncounterCreatureConfig
}

type EncounterCreatureConfig struct {
	Num     int
	Monster string
}

type SelectionConfig struct {
	Key        string
	Name       string
	OptionPack string `json:"option-pack"`
	Options    []SelectionOption
}

type SelectionOption struct {
	Name        string
	Description string
}

type RaceConfig struct {
	Key        string
	Name       string
	OptionPack string `json:"option-pack"`

	Abilities  map[Ability]int
	Languages  []string
	Darkvision int
	Size       Size
	Speed      int

	Spells []RaceSpellConfig
	Props  *RaceProperties
	Profs  *RaceProficiencies

	Traits []LevelTrait
}

type RaceSpellConfig struct {
	Value SpellWithAbilityLevel
}

type SpellWithAbilityLevel struct {
	Level   int
	Key     string
	Ability Ability
}

type RaceProperties struct {
	FlyingSpeed       int             `json:"flying-speed"`
	SkillProficiency  map[Skill]bool  `json:"skill-prof"`
	DamageImmunity    map[Damage]bool `json:"damage-immunity"`
	DamageResistance  map[Damage]bool `json:"damage-resistance"`
	ArmorProficiency  map[Armor]bool  `json:"armor-prof"`
	WeaponProficiency map[string]bool `json:"weapon-prof"`
	LizardfolkAC      bool            `json:"lizardfolk-ac"` // TODO: What is this?
	TortleAC          bool            `json:"tortle-ac"`     // TODO: What is this?
}

type RaceProficiencies struct {
	LanguageOptions          *RaceLanguageOptions          `json:"language-options"`
	ToolOptions              map[string]bool               `json:"tool"`
	SkillOptions             *RaceSkillOptions             `json:"skill-options"`
	WeaponProficiencyOptions *RaceWeaponProficiencyOptions `json:"weapon-proficiency-options"`
}

type RaceLanguageOptions struct {
	// TODO: This seems to be false? Huh?
	Options map[string]bool
}

type RaceWeaponProficiencyOptions struct {
	Options map[string]bool
}

type RaceSkillOptions struct {
	Options map[Skill]bool
}
