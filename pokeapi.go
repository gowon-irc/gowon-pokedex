package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	example       = "slowbro"
	apiURL        = "https://pokeapi.co/api/v2/"
	speciesAPIURL = "pokemon-species/%s"
)

var (
	errIncorrectUsage = errors.New("species name must be included")
)

func pk(args string) (msg string, err error) {
	w := strings.Fields(args)
	if len(w) < 2 {
		return "Error: species name must be included", nil
	}
	species, game := w[0], w[1]

	url := apiURL + fmt.Sprintf(speciesAPIURL, species)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if string(body) == "Not Found" {
		return fmt.Sprintf("Species %s not found", species), nil
	}

	j := &Species{}

	err = json.Unmarshal(body, &j)

	if err != nil {
		return "", err
	}

	f, err := j.FlavorTextEntries.Select("en", game)

	if err != nil && err.Error() == "no results" {
		return fmt.Sprintf("No %s entry found for game %s", species, game), nil
	}

	if err != nil {
		return "", err
	}

	return f.FlavorText, nil
}
