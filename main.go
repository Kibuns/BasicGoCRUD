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

	myRouter.Use(CORS)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAll)
	myRouter.HandleFunc("/create", storeStrat)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func storeStrat(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	fmt.Println("Storing Strat")
	// parse the request body into a Strategy struct
	var strat Strategy
	err := json.NewDecoder(body).Decode(&strat)
	fmt.Println(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// insert the strategy into the database
	insertStrat(strat, w)

}

//service functions

func getAllStrats() (values []primitive.M) {
	return readAllStrats()
}

// other
// CORS Middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}
