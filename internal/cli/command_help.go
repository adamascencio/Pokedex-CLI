package cli

import (
	"fmt"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func CommandHelp(cfg *AppState, cache *pokecache.Cache, area string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}
