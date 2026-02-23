package cli

import (
	"fmt"
	"net/url"

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

func CommandInspect(cfg *AppState, cache *pokecache.Cache, pokemon string) error {
	if pokemon == "" {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	base, err := url.Parse(api.PokemonURL)
	if err != nil {
		panic(err)
	}
	arg, err := url.Parse(pokemon)
	if err != nil {
		panic(err)
	}
	fullURL := base.ResolveReference(arg)
	data, err := api.GetPokemon(cache, fullURL.String())
	if err != nil {
		return err
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
