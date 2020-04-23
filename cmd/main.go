package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablosilvab/cafe-backend-golang/api"
	"github.com/tkanos/gonfig"
)

const defaultConfigPath = "./config/config.json"

type Configuration struct {
	Port string
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Link!")
}

func products(w http.ResponseWriter, r *http.Request) {
	client := api.DBConnect() // debo mantener el cliente conectado .
	collection := client.Database("cafe").Collection("productos")
	log.Println(collection)
	fmt.Fprintf(w, "products endpoints!")
}

func main() {

	// Configuration
	configuration := Configuration{}
	err := gonfig.GetConf(defaultConfigPath, &configuration)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return
	}

	// Routing
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/products", products)
	log.Println("Running on port " + configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
