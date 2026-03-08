package main

import (
	"fmt"
	"os"
	"time"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/cli"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

const commandUsage = "Usage: pokedexcli <command> [arguments]"

func main() {
	state := cli.AppState{Next: api.LocationsURL}
	cache := pokecache.NewCache(10 * time.Second)

	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		fmt.Println(commandUsage)
		cli.Commands["help"].Callback(&state, cache)
		return
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	cmdFunc, ok := cli.Commands[cmd]
	if !ok {
		fmt.Printf("Unknown command: %s\n", cmd)
		fmt.Println(commandUsage)
		return
	}

	if err := cmdFunc.Callback(&state, cache, args...); err != nil {
		fmt.Println(err)
		return
	}
}
