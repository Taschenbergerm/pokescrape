package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taschenbergerm/pokescraper/config"
	"github.com/taschenbergerm/pokescraper/log"
)

type responseData struct {
	RegionName     string         `json:"name"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
}

type PokemonEntry struct {
	EntryNumber    int        `json:"entry_number"`
	PokemonSpecies PokemonRef `json:"pokemon_species"`
}

type PokemonRef struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Scrape the website once
func Scrape() {
	log.Info("Scraper is going to start")
	Config := config.Config()
	log.Info(Config.GetString("KICKER"))
	url := "https://pokeapi.co/api/v2/pokedex/kanto/"

	resp, err := http.Get(url)
	HandleErrorStrictly(err)

	response, err := ioutil.ReadAll(resp.Body)
	HandleErrorStrictly(err)

	var responseObject responseData
	json.Unmarshal(response, &responseObject)

	log.Info(responseObject.RegionName)
	log.Info(responseObject.PokemonEntries[0].PokemonSpecies.Name)
}
