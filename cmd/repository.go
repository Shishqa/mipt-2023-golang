package main

import (
	"log"
	"os"
	"pokemon-rest-api/listing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PokemonRepository interface {
	getPokemons() (listing.Pokemons, error)
	getPokemon(id int) (listing.Pokemon, error)

	addPokemon(pokemon listing.Pokemon) error
}

type Database struct {
	conn *gorm.DB
}

func (db *Database) initConnection() error {
	dsn := os.Getenv("POSTGRES_DSN")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	conn.AutoMigrate(&listing.Pokemon{})
	db.conn = conn
	return nil
}

func (db *Database) getPokemons() (listing.Pokemons, error) {
	var count int64
	err := db.conn.Find(&listing.Pokemon{}).Count(&count).Error
	if err != nil {
		return nil, err
	}
	pokemons := make(listing.Pokemons, count)
	err = db.conn.Model(&listing.Pokemon{}).Find(&pokemons).Error
	return pokemons, err
}

func (db *Database) getPokemon(id int) (listing.Pokemon, error) {
	var pokemon listing.Pokemon
	err := db.conn.First(&pokemon, id).Error
	return pokemon, err
}

func (db *Database) addPokemon(pokemon listing.Pokemon) error {
	log.Print("add pokemon")
	err := db.conn.Create(&pokemon).Error
	return err
}
