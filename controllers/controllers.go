package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
	"github.com/JoseLuisTorrentera/academy-go-q42021/usecases"

	"github.com/gorilla/mux"
)

type getSpells interface {
	GetAllSpells() ([]*models.Spell, error)
}

type getSpell interface {
	GetSpell(name string) (*models.Spell, error)
}
type controller struct {
	ucGetSpells getSpells
	ucGetSpell  getSpell
}

func NewController(ucGetSpells usecases.UcGetSpells, ucGetSpell usecases.UcGetSpell) controller {
	return controller{ucGetSpells: ucGetSpells, ucGetSpell: ucGetSpell}
}

// GetAllSpells - Get all spells from csv
func (c controller) GetAllSpells(w http.ResponseWriter, r *http.Request) {
	spells, err := c.ucGetSpells.GetAllSpells()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(spells)
}

// GetSpellsByName - Get spell by a given name
func (c controller) GetSpellByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	spell, err := c.ucGetSpell.GetSpell(name)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(spell)
}
