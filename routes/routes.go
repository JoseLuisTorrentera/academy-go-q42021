package routes

import (
	"github.com/JoseLuisTorrentera/academy-go-q42021/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/spells", controllers.GetAllSpells)
	r.HandleFunc("/spells/{name}", controllers.GetSpellByName)
	return r
}
