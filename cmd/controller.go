package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pokemon-rest-api/listing"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Controller struct {
	repo PokemonRepository
}

func (c *Controller) setRepository(repo PokemonRepository) {
	c.repo = repo
}

func (c *Controller) setupRoutes(router *httprouter.Router) {
	router.GET("/pokemons", c.getPokemons)
	router.GET("/pokemons/:id", c.getPokemon)
	router.POST("/pokemons", c.addPokemon)
}

func (c *Controller) getPokemons(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("GET /pokemons")

	pokemons, err := c.repo.getPokemons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(pokemons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *Controller) getPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("GET /pokemons/%s\n", ps.ByName("id"))

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pokemon, err := c.repo.getPokemon(id)
	if err != nil {
		if _, ok := err.(*NotFoundError); ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(pokemon)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (c *Controller) addPokemon(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("POST /pokemons")

	pokemonData, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	var pokemon listing.Pokemon
	err = json.Unmarshal(pokemonData, &pokemon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = c.repo.addPokemon(pokemon)
	if err != nil {
		if _, ok := err.(*ConflictError); ok {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
