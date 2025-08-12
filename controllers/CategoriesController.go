package controllers

import (
	"go-web-native/models"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := models.GetCategories()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
