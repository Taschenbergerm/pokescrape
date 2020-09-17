package pkg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/taschenbergerm/pokescraper/log"
)

// Main will start the scrape loop of the region and then for each pokemon
func Main() {
	log.Info("Start Main")
	entries := ScrapeRegion()
	log.Infof("Scraped for %v entries", len(entries))
	pokemonChannel := make(chan PokemonEntry)
	quit := make(chan bool)
	go InsertPokemonLink(pokemonChannel, quit)
	for i, entry := range entries {
		log.Infof("Loop over entry nr. %v - %v",

			i,
			entry.PokemonSpecies.Name)

		p := ScrapePokemon(entry)
		log.Infof("Found %v ", p.Name)
		pokemonChannel <- entry

	}
	quit <- true
	log.Info("Shutting Down")
}

// ScrapeRegion will call the PokeApi to retrieve the list of Pokemons from the webiste
func ScrapeRegion() []PokemonEntry {
	log.Infoln("Scraper is going to start")
	url := "https://pokeapi.co/api/v2/pokedex/kanto/"

	resp, err := http.Get(url)
	HandleErrorStrictly(err)

	response, err := ioutil.ReadAll(resp.Body)
	HandleErrorStrictly(err)

	var responseObject responseData
	json.Unmarshal(response, &responseObject)

	log.Infoln(responseObject.RegionName)
	log.Infoln(responseObject.PokemonEntries[0].PokemonSpecies.Name)
	return responseObject.PokemonEntries
}

// ScrapePokemon retrieves all the individual Facts from the Pokemon via the API
func ScrapePokemon(pokemon PokemonEntry) Pokemon {
	log.Infof("Start Scraping for %s", pokemon.PokemonSpecies.Name)
	resp, err := http.Get(pokemon.PokemonSpecies.URL)
	HandleErrorSoftly(err)

	response, err := ioutil.ReadAll(resp.Body)
	HandleErrorSoftly(err)

	var PokemonInstance Pokemon
	json.Unmarshal(response, &PokemonInstance)
	log.Infof("Scraper %s  with a BaseHappines of %s and GrowthRate of %s",
		PokemonInstance.Name,
		PokemonInstance.BaseHappiness,
		PokemonInstance.GrowthRate.Name)
	return PokemonInstance
}
