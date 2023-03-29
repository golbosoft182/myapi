package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ID    string  `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

var products []Product

func main() {
	router := mux.NewRouter()

	// create sample data
	products = append(products, Product{ID: "1", Name: "Product 1", Price: 99.99})
	products = append(products, Product{ID: "2", Name: "Product 2", Price: 149.99})

	// GET all products
	router.HandleFunc("/api/products", getProducts).Methods("GET")
	// GET a single product
	router.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	// POST a new product
	router.HandleFunc("/api/products", createProduct).Methods("POST")
	// DELETE a product
	router.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")
	// UPDATE a product
	router.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			products[index] = product
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}
