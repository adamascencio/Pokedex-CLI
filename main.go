package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, string) error
}

type config struct {
	Previous string
	Next     string
	Captures map[string]api.Pokemon
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays the names of 20 areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the previous 20 names of areas",
		callback:    commandMapBack,
	},
	"explore": {
		name:        "explore",
		description: "List all pokemon in a specified area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a single pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect the stats of captured pokemon",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "List all captured pokemon",
		callback:    commandPokedex,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	urls := config{
		Next:     api.LocationsURL,
		Captures: make(map[string]api.Pokemon),
	}
	cache := pokecache.NewCache(10 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			text := cleanInput(scanner.Text())
			cmd := text[0]
			var arg string
			if len(text) > 1 {
				arg = text[1]
			}
			if cli, ok := commands[cmd]; ok {
				cli.callback(&urls, cache, arg)
			} else {
				fmt.Println("Unknown command")
			}
			fmt.Print("Pokedex > ")
		}
	}
}
