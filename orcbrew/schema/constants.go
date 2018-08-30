package schema

// Ability is a type alias for abilities
type Ability string

// Symbolic constants for abilities
const (
	Strength     Ability = "str"
	Dexterity    Ability = "dex"
	Constitution Ability = "con"
	Intelligence Ability = "int"
	Wisdom       Ability = "wis"
	Charisma     Ability = "cha"
)

// Size is a type alias for entity size
type Size string

// Symbolic constants for sizes
const (
	Tiny       Size = "tiny"
	Small      Size = "small"
	Medium     Size = "medium"
	Large      Size = "large"
	Huge       Size = "huge"
	Gargantuan Size = "gargantuan"
)

// Skill is a type alias for skill proficiencies
type Skill string

// Symbolic constants for skills
const (
	Acrobatics     Skill = "acrobatics"
	AnimalHandling Skill = "animal-handling"
	Arcana         Skill = "arcana"
	Athletics      Skill = "athletics"
	Deception      Skill = "deception"
	History        Skill = "history"
	Insight        Skill = "insight"
	Intimidation   Skill = "intimidation"
	Investigation  Skill = "investigation"
	Medicine       Skill = "medicine"
	Nature         Skill = "nature"
	Perception     Skill = "perception"
	Performance    Skill = "performance"
	Persuasion     Skill = "persuasion"
	Religion       Skill = "religion"
	SleightOfHand  Skill = "sleight-of-hand"
	Stealth        Skill = "stealth"
	Survival       Skill = "survival"
)

// Damage is a type alias for damage types
type Damage string

// Symbolic constants for damage types
const (
	Acid        Damage = "acid"
	Bludgeoning Damage = "bludgeoning"
	Cold        Damage = "cold"
	Fire        Damage = "fire"
	Lightning   Damage = "lightning"
	Necrotic    Damage = "necrotic"
	Piercing    Damage = "piercing"
	Poison      Damage = "poison"
	Psychic     Damage = "psychic"
	Radiant     Damage = "radiant"
	Slashing    Damage = "slashing"
	Thunder     Damage = "thunder"
	Traps       Damage = "traps"
)

// Condition is a type alias for character conditions
type Condition string

// Symbolic constants for condition types
const (
	Blinded       Condition = "blinded"
	Charmed       Condition = "charmed"
	Deafened      Condition = "deafened"
	Frightened    Condition = "frightened"
	Grapped       Condition = "grappled"
	Incapacitated Condition = "incapacitated"
	Invisible     Condition = "invisible"
	Paralyzed     Condition = "paralyzed"
	Petrified     Condition = "petrified"
	Poisoned      Condition = "poisoned"
	Prone         Condition = "prone"
	Restrained    Condition = "restrained"
	Stunned       Condition = "stunned"
	Unconscious   Condition = "unconscious"
)

// Armor is a type alias for various armor classes, including shields
type Armor string

// Symbolic constants for armor classes
const (
	LightArmor  Armor = "light"
	MediumArmor Armor = "medium"
	HeavyArmor  Armor = "heavy"
	Shields     Armor = "shields"
	Unarmored   Armor = "unarmored"
)

// Weapon is a type alias for classes of weapons
type Weapon string

// Symbolic constants for weapon types
const (
	Simple  Weapon = "simple"
	Martial Weapon = "martial"
)

// MonsterTraitAction is a type alias to define the valid values for the
// "type" field of monster traits
type MonsterTraitAction string

// Symbolic constants for monster trait action types
const (
	MonsterTraitActionAction          MonsterTraitAction = "action"
	MonsterTraitActionLegendaryAction MonsterTraitAction = "legendary-action"
)

// Currency is a currency abbreviation
type Currency string

// Symbolic constants for currency types
const (
	Copper   Currency = "cp"
	Silver   Currency = "sp"
	Electrum Currency = "ep"
	Gold     Currency = "gp"
	Platinum Currency = "pp"
)
