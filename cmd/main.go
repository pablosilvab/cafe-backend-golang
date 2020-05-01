package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablosilvab/cafe-backend-golang/api"
	"github.com/pablosilvab/cafe-backend-golang/api/products"
)

var data api.Configuration

func init() {
	data = api.LoadConfig()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Link!")
}

func main() {
	api.LoadConfig()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")

	router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/ready", readinessHandler)

	router.HandleFunc("/api/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/api/products", products.CreateProduct).Methods("POST")

	server := &http.Server{
		Addr:    ":" + data.Port,
		Handler: router,
	}

	log.Println("Running on port " + data.Port)
	server.ListenAndServe()
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
