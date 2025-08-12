package controllers

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
