package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			fmt.Printf("Your command was: %s\n", output[0])
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
