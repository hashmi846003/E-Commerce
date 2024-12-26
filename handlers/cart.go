package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCartHandler(usersCol *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "Missing user_id query parameter", http.StatusBadRequest)
			return
		}

		var user User
		if err := usersCol.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user); err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user.Cart)
	}
}

func AddToCartHandler(usersCol, productsCol *mongo.Collection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "Missing user_id query parameter", http.StatusBadRequest)
			return
		}

		var cartItem CartItem
		if err := json.NewDecoder(r.Body).Decode(&cartItem); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		var product Product
		if err := productsCol.FindOne(context.TODO(), bson.M{"_id": cartItem.ProductID}).Decode(&product); err != nil {
			http.Error(w, "Product not found", http.StatusBadRequest)
			return
		}
		if product.Stock < cartItem.Quantity {
			http.Error(w, "Insufficient stock", http.StatusBadRequest)
			return
		}

		if _, err := productsCol.UpdateOne(context.TODO(), bson.M{"_id": cartItem.ProductID}, bson.M{"$inc": bson.M{"stock": -cartItem.Quantity}}); err != nil {
			http.Error(w, "Failed to update product stock", http.StatusInternalServerError)
			return
		}

		if _, err := usersCol.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$push": bson.M{"cart": cartItem}}); err != nil {
			http.Error(w, "Failed to update user cart", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cartItem)
	}
}
