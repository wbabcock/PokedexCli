package pokeapi

type Config struct {
	PokeapiClient    Client
	NextLocationsURL *string
	PrevLocationsURL *string
	CaughtPokemon    map[string]Pokemon
}

const (
	baseURL = "https://pokeapi.co/api/v2"
)
