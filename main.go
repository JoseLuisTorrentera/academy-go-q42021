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
	repo := repository.NewSpellRepo()
	servApi := services.NewSpellApiService()
	updateCSV := utils.NewUpdateCSV()
	ucGetSpells := usecases.NewUCGetSpells(repo)
	ucGetSpell := usecases.NewUCGetSpell(servApi, updateCSV)
	c := controllers.NewController(ucGetSpells, ucGetSpell)
	r := routes.NewRouter(c)
	log.Fatal(http.ListenAndServe(":10000", r))
}
