package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakshitha31/urlshortnerchallenge/pkg/controller"
)

func StartServer() {
	fmt.Println("Server running")
	r := mux.NewRouter()
	r.HandleFunc("/shorten", controller.ShortenUrl).Methods("POST")
	r.HandleFunc("/{key}", controller.RedirectUrl).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Error in starting the server", err)
	}
}
