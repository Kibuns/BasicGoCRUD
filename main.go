package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Strategy struct {
	Name    string    `bson:"name"`
	Script  string    `bson:"script"`
	Created time.Time `bson:"created"`
}

// global variable mongodb connection client
var client mongo.Client = newClient()

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
	scriptsCollection := client.Database("testing").Collection("strategies")
	json.NewEncoder(w).Encode(readAll(*scriptsCollection))
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAll).Methods("GET")
	myRouter.HandleFunc("/create", createNewScript).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func createNewScript(w http.ResponseWriter, r *http.Request) {
	scriptsCollection := client.Database("testing").Collection("strategies")

	// parse the request body into a Strategy struct
	var strat Strategy
	err := json.NewDecoder(r.Body).Decode(&strat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// insert the script into the database
	strat.Created = time.Now()
	result, err := scriptsCollection.InsertOne(context.TODO(), strat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the ID of the newly inserted script
	fmt.Fprintf(w, "New script created with ID: %s", result.InsertedID)
}

//service functions

func readAll(collection mongo.Collection) (values []primitive.M) {
	// retrieve all the documents (empty filter)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results from the search query")
	for _, result := range results {
		fmt.Println(result)
	}

	values = results
	return
}

// other
func newClient() (value mongo.Client) {
	//connection mongo db
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	value = *client
	return
}
