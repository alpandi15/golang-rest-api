package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project1/service"

	"github.com/gorilla/mux"
	"github.com/novalagung/gubrak"
)

func findAll(w http.ResponseWriter, r *http.Request) {
	var response service.Response
	var meta service.Meta

	user := service.GetAll()
	meta.TotalData = len(user)

	response.Status = 1
	response.Message = "Success"
	response.Meta = meta
	response.Data = user

	fmt.Println(user)
	fmt.Println(gubrak.RandomInt(10, 20))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user", findAll).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
