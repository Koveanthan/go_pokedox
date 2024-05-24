package commands 

import (
    "fmt"
    "github.com/koveanthan/go_pokedox/config"
)

func CommandHelp(config *config.Config) error {
	fmt.Println("The following are the command you can input")
    for _, cli := range GetCommands(){
        fmt.Println(cli.Name, "-", cli.Description)
    }
	return nil
}


