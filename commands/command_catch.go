package commands

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/koveanthan/go_pokedox/config"
	"github.com/koveanthan/go_pokedox/poke_api/location"
)


func CommandCatch(config *config.Config) error {
    config.UpdateBaseURL(location.PokemonURL)

	if len(config.Params) <= 0 {
		return errors.New("Enter a pokemon name to catch")
	}

	name := config.Params[0]
	resp := location.GetPokemon(name)
	if resp == nil {
		return errors.New("Issue in API fetch.. Check your params and try later.")
	}

	res := location.PokeResponse{}
	err := json.Unmarshal(resp, &res)
	if err == nil {
        if !config.IsPokemonCaught(name) {
            baseExp := res.GetBaseExperience()
            config.CatchPokemon(name, baseExp)
        } else {
            fmt.Printf("Use inspect to fetch details of %v\n", name)
        }
		return nil
	}

	return err
}

