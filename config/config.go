package config

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
    finalRange = 10
    startRange = 0
)

var endRange = 10

type Config struct {
	BaseURL string
	PrevURL *string
	NextURL *string
	Params  []string
    Pokemon []string
}

func (config Config) GetNextURL() string {
	if config.NextURL != nil {
		return *config.NextURL
	}

	return config.BaseURL
}

func (config Config) GetPrevURL() string {
	if config.PrevURL != nil {
		return *config.PrevURL
	}

	return config.BaseURL
}

func (config *Config) UpdateBaseURL(url string) {
	config.BaseURL = url
}

func (config *Config) UpdateNextURL(url *string) {
	config.NextURL = url
}

func (config *Config) UpdatePrevURL(url *string) {
	config.PrevURL = url
}

func (config *Config) IsPokemonCaught(name string) bool{
    for _, v := range config.Pokemon {
        if strings.ToLower(name) == strings.ToLower(v) {
            return true
        }
    }

    return false
}

func (config *Config) CatchPokemon(name string, baseExp int) {
    resetRange(baseExp)
    result := rand.Intn(finalRange)
    fmt.Printf("Throwing a pokeball at %v with probability %v for base experience %v and endrange %v\n", name, result, baseExp, endRange)
    if result <= endRange {
        config.Pokemon = append(config.Pokemon, name)
        fmt.Printf("%v is caught!\n", name)
    } else {
        fmt.Printf("%v escaped!\n", name)
    }
}

func (config Config) PrintPokedox() {
    fmt.Println("The caught pokedox are")
    for _, pokemon := range config.Pokemon {
        fmt.Println("\t-", pokemon)
    }
}

func GenerateConfig() Config {
	return Config{
        "",
		nil,
		nil,
		[]string{},
        []string{},
	}
}

func resetRange(exp int) {
	switch {
	case exp <= 50:
		endRange = 8
	case exp > 50 && exp <= 250:
		endRange = 6
	case exp > 250 && exp <= 300:
		endRange = 4
	case exp > 300:
		endRange = 2
	}
}
