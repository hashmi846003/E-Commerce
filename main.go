package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
  	"ecommerce/handlers"
	"ecommerce/middleware"
	

)

var (
	client         *mongo.Client
	databaseName   = "ecommerce"
	productsColName = "products"
	usersColName    = "users"
)

func main() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")

	productsCol := client.Database(databaseName).Collection(productsColName)
	usersCol := client.Database(databaseName).Collection(usersColName)

	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)

	r.HandleFunc("/products", handlers.GetProductsHandler(productsCol)).Methods(http.MethodGet)
	r.HandleFunc("/products", handlers.CreateProductHandler(productsCol)).Methods(http.MethodPost)
	r.HandleFunc("/users", handlers.GetUsersHandler(usersCol)).Methods(http.MethodGet)
	r.HandleFunc("/users", handlers.CreateUserHandler(usersCol)).Methods(http.MethodPost)
	r.HandleFunc("/cart", handlers.GetCartHandler(usersCol)).Methods(http.MethodGet)
	r.HandleFunc("/cart", handlers.AddToCartHandler(usersCol, productsCol)).Methods(http.MethodPost)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}