package commands

import (
	// "strings"
	"github.com/koveanthan/go_pokedox/config"
)

type CLICommand struct {
	Name        string
	Description string
	Command     func(*config.Config) error
}

func GetCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"help": {
            "help",
			"List all possible that can be used with pokedox",
            CommandHelp,
		},
		"quit": {
			"quit",
			"Quits the command line prompt",
			CommandQuit,
		},
		"map": {
			"map",
			"Fetches the pokemon location. Retype again to paginate forward.",
			CommandMap,
		},
		"mapb": {
			"mapb",
			"Fetches the pokemon location as map. But paginates backwards.",
			CommandMapB,
		},
        "explore": {
            "explore",
            "List the pokemon available in the given area",
            CommandExplore,
        },
        "catch": {
            "catch",
            "Catch the pokemon",
            CommandCatch,
        },
        "inspect": {
            "inspect",
            "Inspect a caught pokemon",
            CommandInspect,
        },
        "pokedox":{
            "pokedox",
            "List all caught pokemon",
            CommandPokedox,
        },
	}
}

