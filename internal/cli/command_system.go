package cli

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

var ErrExitRequested = errors.New("exit requested")

func CleanInput(text string) []string {
	words := strings.Fields(text)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return words
}

func CommandExit(cfg *AppState, cache *pokecache.Cache, area string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return ErrExitRequested
}

func CreateHistoryFile() string {
	f, err := os.OpenFile("/tmp/readline.tmp", os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		panic("Unable to open tmp file")
	}
	f.Close()
	return "/tmp/readline.tmp"
}

func DeleteHistoryFile() error {
	err := os.Remove("/tmp/readline.tmp")
	return err
}
