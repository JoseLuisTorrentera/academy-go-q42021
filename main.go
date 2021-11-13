package main

import (
	"log"
	"net/http"

	"github.com/JoseLuisTorrentera/academy-go-q42021/routes"
)

func main() {
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":10000", r))
}
