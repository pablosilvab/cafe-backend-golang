package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablosilvab/cafe-backend-golang/api"
	"github.com/tkanos/gonfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const defaultConfigPath = "./config/config.json"

type Configuration struct {
	Port string
}

type Product struct {
	//Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Available   bool               `json:"disponible,omitempty" bson:"disponible,omitempty"`
	Name        string             `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Description string             `json:"descripcion,omitempty" bson:"descripcion,omitempty"`
	Price       int32              `json:"precioUni,omitempty" bson:"precioUni,omitempty"`
	Category    primitive.ObjectID `json:"categoria,omitempty" bson:"categoria,omitempty"`
	User        primitive.ObjectID `json:"usuario,omitempty" bson:"usuario,omitempty"`
	Image       string             `json:"img,omitempty" bson:"img,omitempty"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Link!")
}

func products(w http.ResponseWriter, r *http.Request) {
	client := api.DBConnect() // debo mantener el cliente conectado . revisar
	collection := client.Database("cafe").Collection("productos")

	var products []*Product

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var product Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, &product)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(products) // encode similar to serialize process.
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
	router.HandleFunc("/products", products).Methods("GET")
	log.Println("Running on port " + configuration.Port)
	log.Fatal(http.ListenAndServe(":"+configuration.Port, router))
}
