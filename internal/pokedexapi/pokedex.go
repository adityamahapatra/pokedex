package pokedexapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Offset int
	Limit  int
}

func LocationAreas(conf *Config) []string {
	var areas []string

	url := fmt.Sprintf("%s/location-area?offset=%d&limit=%d", baseURL, conf.Offset, conf.Limit)
	request, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer request.Body.Close()

	response, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatalf("Invalid Response: Status Code - %d: %s\n", request.StatusCode, request.Body)
	}

	var locArea LocationArea
	if err = json.Unmarshal(response, &locArea); err != nil {
		log.Fatalln(err)
	}

	for _, result := range locArea.Results {
		areas = append(areas, result.Name)
	}

	return areas
}
