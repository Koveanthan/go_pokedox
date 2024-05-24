package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/koveanthan/go_pokedox/commands"
	"github.com/koveanthan/go_pokedox/config"
)

func StartREPL(config *config.Config) {
	scanner := bufio.NewScanner(os.Stdin)
    
	for {
		fmt.Print("pokedox>")

		scanner.Scan()

		words := cleanText(scanner.Text())

		if len(words) == 0 {
			continue
		}

		input := words[0]
        config.Params = words[1:] 

		command, exists := commands.GetCommands()[input]
		if !exists {
			fmt.Println("Not a command. Try help for available commands!")
            continue
        }

        err := command.Command(config)
        if err != nil {
            fmt.Println(err)
        }
	}
}

func cleanText(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, " ")
	words := strings.Fields(text)
	return words
}

