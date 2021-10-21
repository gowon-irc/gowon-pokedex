package main

import (
	"errors"
)

// Species represents the return content from the pokemon-species endpoint
type Species struct {
	BaseHappiness        int                 `json:"base_happiness"`
	CaptureRate          int                 `json:"capture_rate"`
	Color                color               `json:"color"`
	EggGroups            []eggGroups         `json:"egg_groups"`
	EvolutionChain       evolutionChain      `json:"evolution_chain"`
	EvolvesFromSpecies   evolvesFromSpecies  `json:"evolves_from_species"`
	FlavorTextEntries    flavorTextEntries   `json:"flavor_text_entries"`
	FormDescriptions     []interface{}       `json:"form_descriptions"`
	FormsSwitchable      bool                `json:"forms_switchable"`
	GenderRate           int                 `json:"gender_rate"`
	Genera               []genera            `json:"genera"`
	Generation           generation          `json:"generation"`
	GrowthRate           growthRate          `json:"growth_rate"`
	Habitat              habitat             `json:"habitat"`
	HasGenderDifferences bool                `json:"has_gender_differences"`
	HatchCounter         int                 `json:"hatch_counter"`
	ID                   int                 `json:"id"`
	IsBaby               bool                `json:"is_baby"`
	Name                 string              `json:"name"`
	Names                []names             `json:"names"`
	Order                int                 `json:"order"`
	PalParkEncounters    []palParkEncounters `json:"pal_park_encounters"`
	PokedexNumbers       []pokedexNumbers    `json:"pokedex_numbers"`
	Shape                shape               `json:"shape"`
	Varieties            []varieties         `json:"varieties"`
}

type color struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type eggGroups struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type evolutionChain struct {
	URL string `json:"url"`
}

type evolvesFromSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type flavorTextEntry struct {
	FlavorText string   `json:"flavor_text"`
	Language   language `json:"language"`
	Version    version  `json:"version"`
}

type flavorTextEntries []flavorTextEntry

func (f flavorTextEntries) Select(l, v string) (flavorTextEntry, error) {
	for _, i := range f {
		if i.Language.Name == l && i.Version.Name == v {
			return i, nil
		}
	}
	return flavorTextEntry{}, errors.New("no results")
}

type genera struct {
	Genus    string   `json:"genus"`
	Language language `json:"language"`
}

type generation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type growthRate struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type habitat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type names struct {
	Language language `json:"language"`
	Name     string   `json:"name"`
}

type area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type palParkEncounters struct {
	Area      area `json:"area"`
	BaseScore int  `json:"base_score"`
	Rate      int  `json:"rate"`
}

type pokedex struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokedexNumbers struct {
	EntryNumber int     `json:"entry_number"`
	Pokedex     pokedex `json:"pokedex"`
}

type shape struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type varieties struct {
	IsDefault bool    `json:"is_default"`
	Pokemon   pokemon `json:"pokemon"`
}
