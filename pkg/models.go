package pkg

type responseData struct {
	RegionName     string         `json:"name"`
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
}

// PokemonEntry structure corresponding to the Json
type PokemonEntry struct {
	EntryNumber    int `json:"entry_number"`
	PokemonSpecies Ref `json:"pokemon_species"`
}

// Ref to reuse the common pattern of Name linked to an Url in the json
type Ref struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pokemon Structure from the Json of the API
type Pokemon struct {
	Name                string `json:"name"`
	ID                  int    `json:"id"`
	BaseHappiness       int    `json:"base_happines"`
	CaptureRate         int    `json:"capture_rate"`
	Color               Ref    `json:"color"`
	EggGroups           []Ref  `json:"egg_groups"`
	EvolvesFrom         Ref    `json:"evolves_from_species"`
	GenderRate          int    `json:"gender_rate"`
	Generation          Ref    `json:"generation"`
	GrowthRate          Ref    `json:"growth_rate"`
	HasGenderDifference bool   `json:"has_gender_differences"`
	HatchCounter        int    `json:"hatch_counter"`
}
