package main

import (
	"log"
	"net/http"

	"github.com/JoseLuisTorrentera/academy-go-q42021/controllers"
	"github.com/JoseLuisTorrentera/academy-go-q42021/repository"
	"github.com/JoseLuisTorrentera/academy-go-q42021/routes"
	"github.com/JoseLuisTorrentera/academy-go-q42021/services"
	"github.com/JoseLuisTorrentera/academy-go-q42021/usecases"
	"github.com/JoseLuisTorrentera/academy-go-q42021/utils"
)

func main() {

	ucGetSpells := usecases.NewUCGetSpells(repository.NewSpellRepo("./commons/dnd-spells.csv"))
	ucGetSpell := usecases.NewUCGetSpell(services.NewSpellApiService("./commons/dnd-spells.csv"), utils.NewUpdateCSV())
	c := controllers.NewController(ucGetSpells, ucGetSpell)
	r := routes.NewRouter(c)
	log.Fatal(http.ListenAndServe(":10000", r))
}
