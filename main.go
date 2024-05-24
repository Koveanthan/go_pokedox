package main

import (
	"fmt"

	"github.com/koveanthan/go_pokedox/config"
)

func main() {
    fmt.Println("In Main")
    configVal := config.GenerateConfig()
    StartREPL(&configVal)
}
