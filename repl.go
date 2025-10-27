package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/adityamahapatra/pokedex/internal/pokedexapi"
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

var conf = pokedexapi.Config{
	Offset: 0,
	Limit:  20,
}

func commandMap() error {
	for _, area := range pokedexapi.LocationAreas(&conf) {
		fmt.Println(area)
	}
	conf.Offset += 20
	return nil
}

func commandMapBack() error {
	if conf.Offset < 20 {
		fmt.Println("you're on the first page")
		return nil
	}
	fmt.Println(conf.Offset)
	conf.Offset -= 40
	fmt.Println(conf.Offset)
	for _, area := range pokedexapi.LocationAreas(&conf) {
		fmt.Println(area)
	}

	return nil
}

var commandMapping = map[string]cliCommand{
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
	"map": {
		name:        "map",
		description: "Get 20 location areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Get the previous 20 location areas",
		callback:    commandMapBack,
	},
}

func PokedexPrompt() {
	const prompt = "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		if scanner.Scan() {
			input := scanner.Text()

			if len(input) == 0 {
				continue
			}

			output := cleanInput(input)

			if len(output) == 0 {
				continue
			}

			command, ok := commandMapping[output[0]]
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
