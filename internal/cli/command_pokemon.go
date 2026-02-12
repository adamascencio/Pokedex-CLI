package cli

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func CommandExplore(cfg *AppState, cache *pokecache.Cache, area string) error {
	base, err := url.Parse(api.LocationsURL)
	if err != nil {
		panic(err)
	}
	arg, err := url.Parse(area)
	if err != nil {
		panic(err)
	}
	fullURL := base.ResolveReference(arg)
	res, err := api.GetPokemonInLocation(cache, fullURL.String())
	if err != nil {
		return err
	}
	for _, data := range res.PokemonEncounters {
		fmt.Println(data.Pokemon.Name)
	}
	fmt.Println("")
	return nil
}

func CommandCatch(cfg *AppState, cache *pokecache.Cache, pokemon string) error {
	base, err := url.Parse(api.PokemonURL)
	if err != nil {
		panic(err)
	}
	arg, err := url.Parse(pokemon)
	if err != nil {
		panic(err)
	}
	fullURL := base.ResolveReference(arg)
	res, err := api.GetPokemon(cache, fullURL.String())
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at %s...\n", pokemon)
	baseExp := res.BaseExperience
	catchRate := baseExp / 255
	time.Sleep(2 * time.Second)
	if rand.Float64() >= float64(catchRate) {
		cfg.Captures[pokemon] = res
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}
	fmt.Println("")
	return nil
}

func CommandInspect(cfg *AppState, cache *pokecache.Cache, pokemon string) error {
	data, ok := cfg.Captures[pokemon]
	if !ok {
		fmt.Println("You have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Println("Stats:")
	for _, stats := range data.Stats {
		fmt.Printf("   -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range data.Types {
		fmt.Printf("   - %s\n", types.Type.Name)
	}
	fmt.Println("")
	return nil
}

func CommandPokedex(cfg *AppState, cache *pokecache.Cache, pokemon string) error {
	if !(len(cfg.Captures) > 0) {
		fmt.Println("You have no pokemon...")
		return nil
	}
	for pokemon, _ := range cfg.Captures {
		fmt.Printf(" - %s\n", pokemon)
	}
	fmt.Println("")
	return nil
}
