package main

import "fmt"

func commandHelp() error {
	message := "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n"
	fmt.Println(message)
	return nil
}
