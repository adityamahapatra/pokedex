package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	message := "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n"
	fmt.Println(message)
	return nil
}

var commandMap = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
}

func PokedexPrompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			input := scanner.Text()

			if len(input) == 0 {
				continue
			}

			output := cleanInput(input)

			if len(output) == 0 {
				continue
			}

			command, ok := commandMap[output[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			} else {
				if err := command.callback(); err != nil {
					fmt.Println(err)
				}
			}

		}
	}
}

func cleanInput(text string) []string {
	output := make([]string, 0)
	for _, word := range strings.Fields(text) {
		output = append(output, strings.ToLower(word))
	}

	return output
}
