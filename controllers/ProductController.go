package controllers

import (
	"go-web-native/models"
	"html/template"
	"net/http"
)

func IndexProduct(w http.ResponseWriter, r *http.Request) {
	products := models.GetDataProducts()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/products/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

}

func DetailProduct(w http.ResponseWriter, r *http.Request) {

}

func EditProduct(w http.ResponseWriter, r *http.Request) {

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
