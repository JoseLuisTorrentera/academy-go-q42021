package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoseLuisTorrentera/academy-go-q42021/repository"
	"github.com/JoseLuisTorrentera/academy-go-q42021/services"

	"github.com/gorilla/mux"
)

// GetAllSpells - Get all spells from csv
func GetAllSpells(w http.ResponseWriter, r *http.Request) {
	spells, err := repository.GetAllSpells()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(spells)
}

// GetSpellsByName - Get spell by a given name
func GetSpellByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	spell, err := services.GetSpellByName(name)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(spell)
}
