package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/JoseLuisTorrentera/academy-go-q42021/repository"

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

// GetSpellsById - Get spell by a given id
func GetSpellsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	spell, err := repository.GetSpellById(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(spell)
}
