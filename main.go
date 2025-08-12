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
	http.HandleFunc("/categories", controllers.Index)
	http.HandleFunc("/categories/create", controllers.Create)
	http.HandleFunc("/categories/edit", controllers.Edit)
	http.HandleFunc("/categories/delete", controllers.Delete)

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
