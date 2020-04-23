package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultConfigPath = "./config/config.json"

type Configuration struct {
	Port string
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Link!")
}

func main() {
	// MongoDB

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Configuration
	configuration := Configuration{}
	err = gonfig.GetConf(defaultConfigPath, &configuration)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	// Routing
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Println("Running on port " + configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
