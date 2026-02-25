package cli

var Commands = map[string]Command{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    CommandExit,
	},
	"help": {
		Name:        "help",
		Description: "Display help message",
		Callback:    CommandHelp,
	},
	"map": {
		Name:        "map",
		Description: "Displays the names of 20 areas",
		Callback:    CommandMap,
	},
	"mapb": {
		Name:        "mapb",
		Description: "Displays the previous 20 names of areas",
		Callback:    CommandMapBack,
	},
	"explore": {
		Name:        "explore",
		Description: "List all pokemon in a specified area",
		Callback:    CommandExplore,
	},
	"inspect": {
		Name:        "inspect",
		Description: "Inspect the stats of any pokemon",
		Callback:    CommandInspect,
	},
	"weak": {
		Name:        "weak",
		Description: "Find types a pokemon is weak against",
		Callback:    CommandWeakTo,
	},
	"super": {
		Name:        "super",
		Description: "Find types a pokemon is super effective against",
		Callback:    CommandWeakTo,
	},
}
