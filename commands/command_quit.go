package commands

import (
    "os"
    "github.com/koveanthan/go_pokedox/config"
)

func CommandQuit(config *config.Config) error {
    os.Exit(0)
    return nil
}



