package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)


type Response struct {
	Name    string    `json:"name"`
	Pokeman []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

  var responseObject Response
  json.Unmarshal(responseData, &responseObject)

  fmt.Println(responseObject.Name)
  fmt.Println(len(responseObject.Pokeman))

  for _, pokemon := range responseObject.Pokeman {
    fmt.Println(pokemon.Species.Name)
  }

}
