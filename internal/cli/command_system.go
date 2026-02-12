package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func CleanInput(text string) []string {
	words := strings.Fields(text)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return words
}

func CommandExit(cfg *AppState, cache *pokecache.Cache, area string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
