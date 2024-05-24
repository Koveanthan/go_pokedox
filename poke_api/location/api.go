package location

import (
	"fmt"
	"io"
	"net/http"
	"time"
	//
	"github.com/koveanthan/go_pokedox/poke_cache"
)

const (
    BaseURL = "https://pokeapi.co/api/v2/"
    LocationURL = "https://pokeapi.co/api/v2/location-area"
    PokemonURL = "https://pokeapi.co/api/v2/pokemon"
)

var cache poke_cache.Cache= poke_cache.NewCache(5 * time.Minute)

func Get(url string) []byte { 
    urlCache, ok := cache.Get(url)
        if ok {
            fmt.Println("(Serving from cache)")
            return urlCache
        }
        
        res, err := http.Get(url)
        if err != nil {
            fmt.Println("Fetch Error", err)
            return nil
        }

        defer res.Body.Close()

        if res.StatusCode != 200 {
            fmt.Println("Response Error", res.StatusCode)
            return nil
        }

        content, err := io.ReadAll(res.Body)
        if err != nil {
            fmt.Println("Read Error", err)
            return nil
        }

        cache.Add(url, content)

        return content}

func GetLocation(name string) []byte {
    url := fmt.Sprintf("%v/%v", LocationURL, name)
    return Get(url)
}

func GetPokemon(name string) []byte {
    url := fmt.Sprintf("%v/%v", PokemonURL, name)
    return Get(url)
}
