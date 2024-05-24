package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/koveanthan/go_pokedox/config"
	"github.com/koveanthan/go_pokedox/poke_api/location"
)


func CommandInspect(config *config.Config) error {
    config.UpdateBaseURL(location.PokemonURL)

	if len(config.Params) <= 0 {
		return errors.New("Enter a pokemon name to catch")
	}

	name := config.Params[0]
    if !config.IsPokemonCaught(name) {
        fmt.Printf("Catch the pokemon %v to inspect further\n", name)
        return nil
    }

	resp := location.GetPokemon(name)
	if resp == nil {
		return errors.New("Issue in API fetch.. Check your params and try later.")
	}

	res := location.PokeResponse{}
	err := json.Unmarshal(resp, &res)
	if err == nil {
        fmt.Println(res.GetDetails())
	}

	return err
}

