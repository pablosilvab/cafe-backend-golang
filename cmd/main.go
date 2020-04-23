package main

import (
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

func main() {
	api.LoadConfig()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	log.Println("Running on port " + data.Port)
	log.Fatal(http.ListenAndServe(":"+data.Port, router))
}
