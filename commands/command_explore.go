package commands

import (
	"encoding/json"
	"errors"

	"github.com/koveanthan/go_pokedox/config"
	"github.com/koveanthan/go_pokedox/poke_api/location"
)

func CommandExplore(config *config.Config) error {
    config.UpdateBaseURL(location.LocationURL)

    if len(config.Params) <= 0 {
        return errors.New("Enter a city name to explore Pokemon") 
    }
   
    area := config.Params[0]
    resp := location.GetLocation(area)
    if resp == nil {
        return errors.New("Issue in API fetch.. Check your params and try later.")
    }
    
    res := location.ExpResponse{}
    err := json.Unmarshal(resp, &res)
    if err == nil {
        res.PrintResults()
        return nil
    }

    return err
}
