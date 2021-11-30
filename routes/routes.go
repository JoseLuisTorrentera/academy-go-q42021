package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type controller interface {
	GetAllSpells(w http.ResponseWriter, r *http.Request)
	GetSpellByName(w http.ResponseWriter, r *http.Request)
}

func NewRouter(c controller) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/spells", c.GetAllSpells)
	r.HandleFunc("/spells/{name}", c.GetSpellByName)
	return r
}
