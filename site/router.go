package main

import (
	"api/config"
	"api/middlewares"
	"api/rest"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Carregar()

	router := mux.NewRouter()
	r := router.PathPrefix("/api").Subrouter()

	r.HandleFunc("/login", rest.LoginHandler)
	r.HandleFunc("/recebe-dados", middlewares.Autenticar(rest.RecebeDadosHandler))

	http.Handle("/", router)

	log.Printf("Padronizando para porta %d", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil); err != nil {
		log.Fatal(err)
	}
}
