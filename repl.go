package main

import (
	"strings"
)

func cleanInput(text string) []string {
	output := make([]string, 0)
	for _, word := range strings.Fields(text) {
		output = append(output, strings.ToLower(word))
	}

	return output
}
