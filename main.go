package main

import (
	"go-web-native/config"
	"go-web-native/controllers"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()

	// Homepage
	http.HandleFunc("/", controllers.Welcome)

	// Categories
	http.HandleFunc("/categories", controllers.IndexCategory)
	http.HandleFunc("/categories/create", controllers.CreateCategory)
	http.HandleFunc("/categories/edit", controllers.EditCategory)
	http.HandleFunc("/categories/delete", controllers.DeleteCategory)

	// Products
	http.HandleFunc("/products", controllers.IndexProduct)
	http.HandleFunc("/products/create", controllers.CreateProduct)
	http.HandleFunc("/products/detail", controllers.DetailProduct)
	http.HandleFunc("/products/edit", controllers.EditProduct)
	http.HandleFunc("/products/delete", controllers.DeleteProduct)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
