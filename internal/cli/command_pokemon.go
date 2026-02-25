package cli

import (
	"fmt"
	"net/url"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func buildURL(endpoint string, arg string) *url.URL {
	base, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	urlArg, err := url.Parse(arg)
	if err != nil {
		panic(err)
	}
	fullURL := base.ResolveReference(urlArg)
	return fullURL
}

func CommandExplore(cfg *AppState, cache *pokecache.Cache, area string) error {
	fullURL := buildURL(api.LocationsURL, area)
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
	fullURL := buildURL(api.PokemonURL, pokemon)
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

func CommandWeakTo(cfg *AppState, cache *pokecache.Cache, pokemon string) error {
	if pokemon == "" {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	pokemonURL := buildURL(api.PokemonURL, pokemon)
	pokemonData, err := api.GetPokemon(cache, pokemonURL.String())
	if err != nil {
		return err
	}
	typeURLs := make([]string, 0, 2)
	types := pokemonData.Types
	for _, t := range types {
		typeURLs = append(typeURLs, t.Type.URL)
	}
	weaknesses := make([]string, 0)
	for _, link := range typeURLs {
		typeData, err := api.GetPokemonTypes(cache, link)
		if err != nil {
			return err
		}
		weak_to_slice := typeData.DamageRelations.DoubleDamageFrom
		for _, poketype := range weak_to_slice {
			weaknesses = append(weaknesses, poketype.Name)
		}
	}
	for _, poketype := range weaknesses {
		fmt.Println(poketype)
	}
	return nil
}
