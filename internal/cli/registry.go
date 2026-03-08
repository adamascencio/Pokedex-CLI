package cli

var Commands = map[string]Command{
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
		Callback:    CommandSuperEffective,
	},
	"find": {
		Name:        "find",
		Description: "Find locations where you can catch a specific pokemon",
		Callback:    CommandFindPokemon,
	},
}
