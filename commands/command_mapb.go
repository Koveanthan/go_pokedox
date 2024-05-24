package commands

import (
	"encoding/json"
	// "strconv"
	// "fmt"
    "github.com/koveanthan/go_pokedox/poke_api/location"
    "github.com/koveanthan/go_pokedox/config"
)

func CommandMapB(config *config.Config) error {
    config.UpdateBaseURL(location.LocationURL)

    resBytes := location.Get(config.GetPrevURL())

	response := location.LocResponse{}
	err := json.Unmarshal(resBytes, &response)
	if err == nil {
        response.PrintResults()
        config.UpdateNextURL(response.Next)
        config.UpdatePrevURL(response.Previous)
		return nil
	}

	return err
}
