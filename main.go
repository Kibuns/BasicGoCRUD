package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	handleRequests()
}

//controllers

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAll")
	json.NewEncoder(w).Encode(getAllStrats())
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAll).Methods("GET")
	myRouter.HandleFunc("/create", storeStrat).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func storeStrat(w http.ResponseWriter, r *http.Request) {

	// parse the request body into a Strategy struct
	var strat Strategy
	err := json.NewDecoder(r.Body).Decode(&strat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert the strategy into the database
	insertStrat(strat, w)

}

//service functions

func getAllStrats() (values []primitive.M) {
	return readAllStrats()
}
