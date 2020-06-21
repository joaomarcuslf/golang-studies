package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"./person"
)

func main() {
	person.Seed()

	router := mux.NewRouter()

	router.HandleFunc("/", HealthCheck).Methods("GET")

	router.HandleFunc("/person", person.Index).Methods("GET")
	router.HandleFunc("/person/{id}", person.Get).Methods("GET")
	router.HandleFunc("/person/{id}", person.Create).Methods("POST")
	router.HandleFunc("/person/{id}", person.Delete).Methods("DELETE")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8000", loggedRouter)

	fmt.Println("Server started at: http://localhost:8000/")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
