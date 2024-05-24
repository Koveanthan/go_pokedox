package commands

import (
	"encoding/json"
	"errors"

	// "strconv"
	// "fmt"
	"github.com/koveanthan/go_pokedox/config"
	"github.com/koveanthan/go_pokedox/poke_api/location"
)

func CommandMap(config *config.Config) error {
    config.UpdateBaseURL(location.LocationURL)

    resBytes := location.Get(config.GetNextURL())
    if resBytes == nil {
        return errors.New("Issue in API fetch.. Check your params and try later.")
    }

    response := location.LocResponse{}
    err := json.Unmarshal(resBytes, &response)
    if err == nil {
        config.UpdateNextURL(response.Next)
        config.UpdatePrevURL(response.Previous)
        response.PrintResults()
        return nil
    }

    return err
}
