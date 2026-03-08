package cli

import (
	"fmt"

	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func CommandHelp(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("  pokedexcli <command> [arguments]")
	fmt.Println("help: Display help message")
	fmt.Println("map: Displays the names of 20 areas")
	fmt.Println("explore: List all pokemon in a specified area")
	fmt.Println("inspect: Inspect the stats of any pokemon")
	fmt.Println("weak: Find types a pokemon is weak against")
	fmt.Println("super: Find types a pokemon is super effective against")
	fmt.Println("find: Find locations where you can catch a specific pokemon")
	return nil
}
