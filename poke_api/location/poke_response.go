package location

import (
	"fmt"
	"strconv"
)

type PokeResponse struct {
	BaseExperience int `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	Name          string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (res PokeResponse) GetBaseExperience() int{
    return res.BaseExperience
}

func (res PokeResponse) GetDetails() string{
    spec := `Name: %v
Id: %v
Base Experience: %v
Height: %v
Weight: %v
Stat: %v
Types: %v`
    final := fmt.Sprintf(spec, res.Name, res.ID, res.BaseExperience, res.Height, res.Weight, getStats(res), getTypes(res)) 
    return final
}

func getStats(res PokeResponse) string {
    final := "\n\t"
    for i, stat := range res.Stats {
        final += "-" + stat.Stat.Name + ": " + strconv.Itoa(stat.BaseStat)
        if i < len(res.Stats) - 1 {
            final += "\n\t"
        }
    }
    return final
}

func getTypes(res PokeResponse) string {
    final := "\n\t"
    for i, typ := range res.Types {
        final += "-" + typ.Type.Name
        if i < len(res.Types) - 1 {
            final += "\n\t"
        }
    }
    return final
}

