package cli

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

type inspectResponse struct {
	Name   string        `json:"name"`
	Height int           `json:"height"`
	Weight int           `json:"weight"`
	Stats  []inspectStat `json:"stats"`
	Types  []inspectType `json:"types"`
}

type inspectStat struct {
	Name     string `json:"name"`
	BaseStat int    `json:"base_stat"`
}

type inspectType struct {
	Name string `json:"name"`
}

func buildURL(endpoint string, args ...string) *url.URL {
	base, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}
	fullURL := base
	for _, arg := range args {
		if arg == "" {
			continue
		}
		urlArg, err := url.Parse(arg)
		if err != nil {
			panic(err)
		}
		fullURL = fullURL.ResolveReference(urlArg)
	}
	return fullURL
}

func CommandExplore(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Provide an area")
		return nil
	}
	fullURL := buildURL(api.LocationsURL, args[0])
	res, err := api.GetPokemonInLocation(cache, fullURL.String())
	if err != nil {
		return err
	}
	for _, data := range res.PokemonEncounters {
		fmt.Println(data.Pokemon.Name)
	}
	return nil
}

func CommandInspect(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	pokemonName := ""
	jsonOutput := false
	for _, arg := range args {
		switch arg {
		case "--json":
			jsonOutput = true
		default:
			if pokemonName == "" {
				pokemonName = arg
			}
		}
	}
	if pokemonName == "" {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	fullURL := buildURL(api.PokemonURL, pokemonName)
	data, err := api.GetPokemon(cache, fullURL.String())
	if err != nil {
		return err
	}
	if jsonOutput {
		response := inspectResponse{
			Name:   data.Name,
			Height: data.Height,
			Weight: data.Weight,
			Stats:  make([]inspectStat, 0, len(data.Stats)),
			Types:  make([]inspectType, 0, len(data.Types)),
		}
		for _, stat := range data.Stats {
			response.Stats = append(response.Stats, inspectStat{
				Name:     stat.Stat.Name,
				BaseStat: stat.BaseStat,
			})
		}
		for _, pokemonType := range data.Types {
			response.Types = append(response.Types, inspectType{
				Name: pokemonType.Type.Name,
			})
		}
		encoded, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(encoded))
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

func CommandWeakTo(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	pokemonURL := buildURL(api.PokemonURL, args[0])
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

func CommandSuperEffective(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	pokemonURL := buildURL(api.PokemonURL, args[0])
	pokemonData, err := api.GetPokemon(cache, pokemonURL.String())
	if err != nil {
		return err
	}
	typeURLs := make([]string, 0, 2)
	types := pokemonData.Types
	for _, t := range types {
		typeURLs = append(typeURLs, t.Type.URL)
	}
	superEffective := make([]string, 0)
	for _, link := range typeURLs {
		typeData, err := api.GetPokemonTypes(cache, link)
		if err != nil {
			return err
		}
		weak_to_slice := typeData.DamageRelations.DoubleDamageTo
		for _, poketype := range weak_to_slice {
			superEffective = append(superEffective, poketype.Name)
		}
	}
	for _, poketype := range superEffective {
		fmt.Println(poketype)
	}
	return nil
}

func CommandFindPokemon(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon name.")
		return nil
	}
	endpoint := "https://pokeapi.co/api/v2/pokemon/"
	pokemonURL := buildURL(endpoint, args[0]+"/", "encounters")
	locationData, err := api.GetPokemonEncounters(cache, pokemonURL.String())
	if err != nil {
		return err
	}
	locations := make([]string, 0)
	games := args[1:]
	gameFilter := make(map[string]struct{}, len(games))
	for _, game := range games {
		if game == "" {
			continue
		}
		gameFilter[game] = struct{}{}
	}
	for _, d := range locationData {
		if len(gameFilter) == 0 {
			locations = append(locations, d.LocationArea.Name)
			continue
		}
		for _, v := range d.VersionDetails {
			if _, ok := gameFilter[v.Version.Name]; ok {
				locations = append(locations, d.LocationArea.Name)
				break
			}
		}
	}
	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}
