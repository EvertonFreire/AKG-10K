package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", hello).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World -> Olá Go 10K")
}
