package main

import (
	"log"
	"net/http"
	usercontroller "project1/controllers/user"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", usercontroller.FindAll).Methods("GET")

	log.Print("Server Linten localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
