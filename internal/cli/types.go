package cli

import (
	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

type Command struct {
	Name        string
	Description string
	Callback    func(*AppState, *pokecache.Cache, string) error
}

type AppState struct {
	Previous string
	Next     string
	Captures map[string]api.Pokemon
}
