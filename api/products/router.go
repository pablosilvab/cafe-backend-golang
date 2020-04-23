package products

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pablosilvab/cafe-backend-golang/api"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	client := api.DBConnect()
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

	response, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
