package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name           string           `json:"name"`
	PokemonEntries []PokemonEntries `json:"Pokemon_entries"`
}

type PokemonEntries struct {
	EntryNumber    int            `json:"entry_number"`
	PokemonSpecies PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

	var responseObject Response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		return
	}

	fmt.Println(responseObject.Name)
	fmt.Println(responseObject.PokemonEntries)

	for _, pokemon := range responseObject.PokemonEntries {
		fmt.Println(pokemon.EntryNumber)
		fmt.Println(pokemon.PokemonSpecies.Name)
	}

}
