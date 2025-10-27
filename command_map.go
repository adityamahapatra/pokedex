package main

import (
	"fmt"

	"github.com/adityamahapatra/pokedex/internal/pokedexapi"
)

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
