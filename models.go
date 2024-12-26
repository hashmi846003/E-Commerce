package main


type Product struct {
	ID    string  `json:"id" bson:"_id"`
	Name  string  `json:"name" bson:"name"`
	Price float64 `json:"price" bson:"price"`
	Stock int     `json:"stock" bson:"stock"`
}

type User struct {
	ID       string     `json:"id" bson:"_id"`
	Username string     `json:"username" bson:"username"`
	Cart     []CartItem `json:"cart" bson:"cart"`
}

type CartItem struct {
	ProductID string `json:"product_id" bson:"product_id"`
	Quantity  int    `json:"quantity" bson:"quantity"`
}
