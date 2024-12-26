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

type User struct {
	ID       string     `json:"id" bson:"_id"`
	Username string     `json:"username" bson:"username"`
	Cart     []CartItem `json:"cart" bson:"cart"`
}

type CartItem struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}

func GetUsersHandler(usersCol *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, err := usersCol.Find(context.TODO(), bson.M{})
		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		var users []User
		if err := cursor.All(context.TODO(), &users); err != nil {
			http.Error(w, "Failed to parse users", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

func CreateUserHandler(usersCol *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		user.ID = strconv.Itoa(rand.Intn(1000000))
		if _, err := usersCol.InsertOne(context.TODO(), user); err != nil {
			http.Error(w, "Failed to add user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
