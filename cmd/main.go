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

func main() {

	api.DBConnect()

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
	log.Println("Running on port " + configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
