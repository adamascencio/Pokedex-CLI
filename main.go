package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/cli"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
	"github.com/chzyer/readline"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func main() {
	historyFile := cli.CreateHistoryFile()
	defer cli.DeleteHistoryFile()
	l, err := readline.NewEx(&readline.Config{
		Prompt:              "\033[31m»\033[0m ",
		HistoryFile:         historyFile,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	l.CaptureExitSignal()
	state := cli.AppState{
		Next: api.LocationsURL,
	}
	cache := pokecache.NewCache(10 * time.Second)
	log.SetOutput(l.Stderr())
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				if err := cli.DeleteHistoryFile(); err != nil {
					fmt.Println(err)
				}
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			if err := cli.DeleteHistoryFile(); err != nil {
				fmt.Println(err)
			}
			break
		}

		args := cli.CleanInput(line)

		if len(args) == 0 {
			continue
		}
		cmd := args[0]
		arg := ""
		if len(args) > 1 {
			arg = args[1]
		}

		if cliFunc, ok := cli.Commands[cmd]; ok {
			if err := cliFunc.Callback(&state, cache, arg); err != nil {
				if errors.Is(err, cli.ErrExitRequested) {
					cli.DeleteHistoryFile()
					break
				}
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
