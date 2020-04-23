package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/pablosilvab/cafe-backend-golang/api"
	"github.com/pablosilvab/cafe-backend-golang/api/products"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Link!")
}

var data api.Configuration

func init() {
	data = api.LoadConfig()
}

func main() {
	api.LoadConfig()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/products", products.GetProducts).Methods("GET")
	log.Println("Running on port " + data.Port)
	log.Fatal(http.ListenAndServe(":"+data.Port, router))
}
