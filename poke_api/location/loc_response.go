package location

import "fmt"

type LocResponse struct {
	Count    int       `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []LocResults `json:"results"`
}

type LocResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (res LocResponse) PrintResults() {
	for _, result := range res.Results {
		fmt.Println(result.Name)
	}
}

