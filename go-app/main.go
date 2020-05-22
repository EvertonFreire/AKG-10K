package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello() string {
	return "Hello World -> Olá Go 10K"
}
