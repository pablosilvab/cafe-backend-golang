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
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	router.HandleFunc("/products", products.CreateProduct).Methods("POST")
	log.Println("Running on port " + data.Port)
	log.Fatal(http.ListenAndServe(":"+data.Port, router))
}
