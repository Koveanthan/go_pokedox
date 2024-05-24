package commands

import (
	"github.com/koveanthan/go_pokedox/config"
)


func CommandPokedox(config *config.Config) error {
    config.PrintPokedox()
	return nil
}

