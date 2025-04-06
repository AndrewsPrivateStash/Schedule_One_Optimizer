/*
	The static global data used by various functions
*/

package main

type Product struct {
	Name        string
	Base_price  float64
	Base_effect string
}

type Ingredient struct {
	Name   string
	Effect string
	Cost   float64
	Rank   string
}

type Rank struct {
	Name    string
	Level   int
	Unlocks []string
}

// products
var PRODS = [...]Product{
	{"OG Kush", 35.00, "Calming"},
	{"Sour Diesel", 35.00, "Refreshing"},
	{"Green Crack", 35.00, "Energizing"},
	{"Granddaddy Purple", 35.00, "Sedating"},
	{"Meth", 70.00, ""},
	{"Cocaine", 150.00, ""},
}

// Ingredients
var INGRS = [...]Ingredient{
	{"Cuke", "Energizing", 2.00, "Street rat I"},
	{"Flu Medicine", "Sedating", 5.00, "Hoodlum IV"},
	{"Gasoline", "Toxic", 5.00, "Hoodlum V"},
	{"Donut", "Calorie-Dense", 3.00, "Street rat I"},
	{"Energy Drink", "Athletic", 6.00, "Peddler I"},
	{"Mouth Wash", "Balding", 4.00, "Hoodlum III"},
	{"Motor Oil", "Slippery", 6.00, "Peddler II"},
	{"Banana", "Gingeritis", 2.00, "Street rat I"},
	{"Chili", "Spicy", 7.00, "Peddler IV"},
	{"Iodine", "Jennerising", 8.00, "Hustler I"},
	{"Paracetamol", "Sneaky", 3.00, "Street rat I"},
	{"Viagra", "Tropic Thunder", 4.00, "Hoodlum II"},
	{"Horse Semen", "Long Faced", 9.00, "Hustler III"},
	{"Mega Bean", "Foggy", 7.00, "Peddler III"},
	{"Addy", "Thought-Provoking", 9.00, "Hustler II"},
	{"Battery", "Bright-Eyed", 8.00, "Peddler V"},
}

// Effect Multipliers
var MULT = map[string]float64{
	"Anti-Gravity":      0.54,
	"Athletic":          0.32,
	"Balding":           0.30,
	"Bright-Eyed":       0.40,
	"Calming":           0.10,
	"Calorie-Dense":     0.28,
	"Cyclopean":         0.56,
	"Disorienting":      0.00,
	"Electrifying":      0.50,
	"Energizing":        0.22,
	"Euphoric":          0.18,
	"Explosive":         0.00,
	"Focused":           0.16,
	"Foggy":             0.36,
	"Gingeritis":        0.20,
	"Glowing":           0.48,
	"Jennerising":       0.42,
	"Laxative":          0.00,
	"Lethal":            0.00,
	"Long Faced":        0.52,
	"Munchies":          0.12,
	"Paranoia":          0.00,
	"Refreshing":        0.14,
	"Schizophrenia":     0.00,
	"Sedating":          0.26,
	"Seizure-Inducing":  0.00,
	"Shrinking":         0.60,
	"Slippery":          0.34,
	"Smelly":            0.00,
	"Sneaky":            0.24,
	"Spicy":             0.38,
	"Thought-Provoking": 0.44,
	"Toxic":             0.00,
	"Tropic Thunder":    0.46,
	"Zombifying":        0.58,
}

// Ingredient Effects
var ING_EFFECTS = map[string]string{
	"Cuke":         "Energizing",
	"Flu Medicine": "Sedating",
	"Gasoline":     "Toxic",
	"Donut":        "Calorie-Dense",
	"Energy Drink": "Athletic",
	"Mouth Wash":   "Balding",
	"Motor Oil":    "Slippery",
	"Banana":       "Gingeritis",
	"Chili":        "Spicy",
	"Iodine":       "Jennerising",
	"Paracetamol":  "Sneaky",
	"Viagra":       "Tropic Thunder",
	"Horse Semen":  "Long Faced",
	"Mega Bean":    "Foggy",
	"Addy":         "Thought-Provoking",
	"Battery":      "Bright-Eyed",
}

// ranks
var RANKS = map[string]Rank{
	"Street Rat I":    {"Street Rat I", 1, []string{"OG Kush", "Cuke", "Banana"}},
	"Street Rat II":   {"Street Rat II", 2, []string{}},
	"Street Rat III":  {"Street Rat III", 3, []string{"Jar", "PGR", "Speed grow"}},
	"Street Rat IV":   {"Street Rat IV", 4, []string{"Long-Life Soil", "Sour Diesel seeds", "Sour Diesel"}},
	"Street Rat V":    {"Street Rat V", 5, []string{"Pot Sprinkler", "Electric Plant Trimmers"}},
	"Hoodlum I":       {"Hoodlum I", 6, []string{"Westville region", "Low-Quality Pseudo", "Flu Medicine", "Gasoline", "Donut", "Meth"}},
	"Hoodlum II":      {"Hoodlum II", 7, []string{"Soil Pourer", "Green Crack seeds", "Energy Drink", "Green Crack"}},
	"Hoodlum III":     {"Hoodlum III", 8, []string{"Mouth Wash"}},
	"Hoodlum IV":      {"Hoodlum IV", 9, []string{"Extra Long-Life Soil", "Granddaddy purple seeds", "Motor Oil", "Granddaddy Purple"}},
	"Hoodlum V":       {"Hoodlum V", 10, []string{"Packaging station Mk II", "Mixing station", "Banana"}},
	"Peddler I":       {"Peddler I", 11, []string{"Iodine"}},
	"Peddler II":      {"Peddler II", 12, []string{"Paracetamol", "Mixing Station Mk II"}},
	"Peddler III":     {"Peddler III", 13, []string{"Viagra"}},
	"Peddler IV":      {"Peddler IV", 14, []string{}},
	"Peddler V":       {"Peddler V", 15, []string{}},
	"Hustler I":       {"Hustler I", 16, []string{"Downtown region"}},
	"Hustler II":      {"Hustler II", 17, []string{}},
	"Hustler III":     {"Hustler III", 18, []string{"Pseudo", "Horse Semen"}},
	"Hustler IV":      {"Hustler IV", 19, []string{}},
	"Hustler V":       {"Hustler V", 20, []string{}},
	"Bagman I":        {"Bagman I", 21, []string{}},
	"Bagman II":       {"Bagman II", 22, []string{}},
	"Bagman III":      {"Bagman III", 23, []string{}},
	"Bagman IV":       {"Bagman IV", 24, []string{}},
	"Bagman V":        {"Bagman V", 25, []string{"Brick press", "High-Quality Pseudo"}},
	"Enforcer I":      {"Enforcer I", 26, []string{"Docks region", "Cauldron", "Cocaine"}},
	"Enforcer II":     {"Enforcer II", 27, []string{}},
	"Enforcer III":    {"Enforcer III", 28, []string{}},
	"Enforcer IV":     {"Enforcer IV", 29, []string{}},
	"Enforcer V":      {"Enforcer V", 30, []string{}},
	"Shot Caller I":   {"Shot Caller I", 31, []string{}},
	"Shot Caller II":  {"Shot Caller II", 32, []string{}},
	"Shot Caller III": {"Shot Caller III", 33, []string{}},
	"Shot Caller IV":  {"Shot Caller IV", 34, []string{}},
	"Shot Caller V":   {"Shot Caller V", 35, []string{}},
	"Block Boss I":    {"Block Boss I", 36, []string{"Suburbia region"}},
	"Block Boss II":   {"Block Boss II", 37, []string{}},
	"Block Boss III":  {"Block Boss III", 38, []string{}},
	"Block Boss IV":   {"Block Boss IV", 39, []string{}},
	"Block Boss V":    {"Block Boss V", 40, []string{}},
	"Underlord I":     {"Underlord I", 41, []string{}},
	"Underlord II":    {"Underlord II", 42, []string{}},
	"Underlord III":   {"Underlord III", 43, []string{}},
	"Underlord IV":    {"Underlord IV", 44, []string{}},
	"Underlord V":     {"Underlord V", 45, []string{}},
	"Baron I":         {"Baron I", 46, []string{"Uptown region"}},
	"Baron II":        {"Baron II", 47, []string{}},
	"Baron III":       {"Baron III", 48, []string{}},
	"Baron IV":        {"Baron IV", 49, []string{}},
	"Baron V":         {"Baron V", 50, []string{}},
	"Kingpin I":       {"Kingpin I", 51, []string{}},
}

// transforms (effect interactions)
// assume ingredient itself is irrelevent and only the effect it transmits matters
var TRANS = map[string]map[string]string{
	"Energizing": {
		"Toxic":      "Euphoric",
		"Slippery":   "Munchies",
		"Sneaky":     "Paranoia",
		"Foggy":      "Cyclopean",
		"Gingeritis": "Thought-Provoking",
		"Munchies":   "Athletic",
		"Euphoric":   "Laxative",
	},
	"Sedating": {
		"Calming":           "Bright-Eyed",
		"Athletic":          "Munchies",
		"Thought-Provoking": "Gingeritis",
		"Cyclopean":         "Foggy",
		"Munchies":          "Slippery",
		"Laxative":          "Euphoric",
		"Euphoric":          "Toxic",
		"Focused":           "Calming",
		"Electrifying":      "Refreshing",
		"Shrinking":         "Paranoia",
	},
	"Toxic": {
		"Gingeritis":   "Smelly",
		"Jennerising":  "Sneaky",
		"Sneaky":       "Tropic Thunder",
		"Munchies":     "Sedating",
		"Energizing":   "Euphoric",
		"Euphoric":     "Energizing",
		"Laxative":     "Foggy",
		"Disorienting": "Glowing",
		"Paranoia":     "Calming",
		"Electrifying": "Disorienting",
		"Shrinking":    "Focused",
	},
	"Calorie-Dense": {
		"Calorie-Dense": "Explosive",
		"Balding":       "Sneaky",
		"Anti-Gravity":  "Slippery",
		"Jennerising":   "Gingeritis",
		"Focused":       "Euphoric",
		"Shrinking":     "Energizing",
	},
	"Athletic": {
		"Sedating":       "Munchies",
		"Euphoric":       "Energizing",
		"Spicy":          "Euphoric",
		"Tropic Thunder": "Sneaky",
		"Glowing":        "Disorienting",
		"Foggy":          "Laxative",
		"Disorienting":   "Electrifying",
		"Schizophrenia":  "Balding",
		"Focused":        "Shrinking",
	},
	"Balding": {
		"Calming":       "Anti-Gravity",
		"Calorie-Dense": "Sneaky",
		"Explosive":     "Sedating",
		"Focused":       "Jennerising",
	},
	"Slippery": {
		"Energizing": "Munchies",
		"Foggy":      "Toxic",
		"Euphoric":   "Sedating",
		"Paranoia":   "Anti-Gravity",
		"Munchies":   "Schizophrenia",
	},
	"Gingeritis": {
		"Energizing":   "Thought-Provoking",
		"Calming":      "Sneaky",
		"Toxic":        "Smelly",
		"Long Faced":   "Refreshing",
		"Cyclopean":    "Thought-Provoking",
		"Disorienting": "Focused",
		"Focused":      "Seizure-Inducing",
		"Paranoia":     "Jennerising",
		"Smelly":       "Anti-Gravity",
	},
	"Spicy": {
		"Athletic":     "Euphoric",
		"Anti-Gravity": "Tropic Thunder",
		"Sneaky":       "Bright-Eyed",
		"Munchies":     "Toxic",
		"Laxative":     "Long Faced",
		"Shrinking":    "Refreshing",
	},
	"Jennerising": {
		"Calming":       "Balding",
		"Toxic":         "Sneaky",
		"Foggy":         "Paranoia",
		"Calorie-Dense": "Gingeritis",
		"Euphoric":      "Seizure-Inducing",
		"Refreshing":    "Thought-Provoking",
	},
	"Sneaky": {
		"Energizing":   "Paranoia",
		"Calming":      "Slippery",
		"Toxic":        "Tropic Thunder",
		"Spicy":        "Bright-Eyed",
		"Glowing":      "Toxic",
		"Foggy":        "Calming",
		"Munchies":     "Anti-Gravity",
		"Paranoia":     "Balding",
		"Electrifying": "Athletic",
		"Focused":      "Gingeritis",
	},
	"Tropic Thunder": {
		"Athletic":     "Sneaky",
		"Euphoric":     "Bright-Eyed",
		"Laxative":     "Calming",
		"Disorienting": "Toxic",
	},
	"Long Faced": {
		"Anti-Gravity":      "Calming",
		"Gingeritis":        "Refreshing",
		"Thought-Provoking": "Electrifying",
	},
	"Foggy": {
		"Energizing":        "Cyclopean",
		"Calming":           "Glowing",
		"Sneaky":            "Calming",
		"Jennerising":       "Paranoia",
		"Athletic":          "Laxative",
		"Slippery":          "Toxic",
		"Thought-Provoking": "Energizing",
		"Seizure-Inducing":  "Focused",
		"Focused":           "Disorienting",
		"Shrinking":         "Electrifying",
	},
	"Thought-Provoking": {
		"Sedating":   "Gingeritis",
		"Long Faced": "Electrifying",
		"Glowing":    "Refreshing",
		"Foggy":      "Energizing",
		"Explosive":  "Euphoric",
	},
	"Bright-Eyed": {
		"Munchies":     "Tropic Thunder",
		"Euphoric":     "Zombifying",
		"Electrifying": "Euphoric",
		"Laxative":     "Calorie-Dense",
		"Cyclopean":    "Glowing",
		"Shrinking":    "Munchies",
	},
}
