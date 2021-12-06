package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

type getSpellQuery interface {
	GetSpellsByQuery(itemType string, numItems int, numItemsWorker int) ([]*models.Spell, error)
}
type controller struct {
	ucGetSpells     getSpells
	ucGetSpell      getSpell
	ucGetSpellQuery getSpellQuery
}

func NewController(ucGetSpells usecases.UcGetSpells, ucGetSpell usecases.UcGetSpell, ucGetSpellQuery usecases.UcGetSpellsQuery) controller {
	return controller{ucGetSpells: ucGetSpells, ucGetSpell: ucGetSpell, ucGetSpellQuery: ucGetSpellQuery}
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

func (c controller) GetSpellsByQuery(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	t := vars["type"]
	items := vars["items"]
	iw := vars["items_workers"]

	itemsInt, err := strconv.Atoi(items[0])
	if err != nil {
		log.Fatal(err)
	}
	iwInt, err := strconv.Atoi(iw[0])
	if err != nil {
		log.Fatal(err)
	}

	spells, err := c.ucGetSpellQuery.GetSpellsByQuery(t[0], itemsInt, iwInt)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(spells)
}
