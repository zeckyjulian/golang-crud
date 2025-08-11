package main

import (
	"go-web-native/config"
	"log"
	"net/http"
)

func main() {
	config.ConnectDb()

	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
