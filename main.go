package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/cli"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	urls := cli.AppState{
		Next:     api.LocationsURL,
		Captures: make(map[string]api.Pokemon),
	}
	cache := pokecache.NewCache(10 * time.Second)
	for {
		fmt.Print("Pokedex > ")
		for scanner.Scan() {
			text := cli.CleanInput(scanner.Text())
			if len(text) == 0 {
				continue
			}
			cmd := text[0]
			var arg string
			if len(text) > 1 {
				arg = text[1]
			}
			if cliFunc, ok := cli.Commands[cmd]; ok {
				cliFunc.Callback(&urls, cache, arg)
			} else {
				fmt.Println("Unknown command")
			}
			fmt.Print("Pokedex > ")
		}
	}
}
