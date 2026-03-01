package cli

import (
	"fmt"
	"github.com/adamascencio/pokedexcli/internal/api"
	"github.com/adamascencio/pokedexcli/internal/pokecache"
)

func CommandMap(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	res, err := api.GetLocations(cache, cfg.Next)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, data := range res.Results {
		fmt.Println(data.Name)
	}
	cfg.Previous = res.Previous
	cfg.Next = res.Next
	fmt.Println("")
	return nil
}

func CommandMapBack(cfg *AppState, cache *pokecache.Cache, args ...string) error {
	if cfg.Previous == "" {
		fmt.Println("Call map first before calling mapb")
		return nil
	}
	res, err := api.GetLocations(cache, cfg.Previous)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, data := range res.Results {
		fmt.Println(data.Name)
	}
	cfg.Previous = res.Previous
	cfg.Next = res.Next
	fmt.Println("")
	return nil
}
