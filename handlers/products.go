package handlers

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID    string  `json:"id" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
	Stock int     `json:"stock" bson:"stock"`
}

func GetProductsHandler(productsCol *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, err := productsCol.Find(context.TODO(), bson.M{})
		if err != nil {
			http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
			return
		}
		var products []Product
		if err := cursor.All(context.TODO(), &products); err != nil {
			http.Error(w, "Failed to parse products", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(products)
	}
}

func CreateProductHandler(productsCol *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		product.ID = strconv.Itoa(rand.Intn(1000000))
		if _, err := productsCol.InsertOne(context.TODO(), product); err != nil {
			http.Error(w, "Failed to add product", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}
