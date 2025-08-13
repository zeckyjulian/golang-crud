package controllers

import (
	"go-web-native/entities"
	"go-web-native/models"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func IndexCategory(w http.ResponseWriter, r *http.Request) {
	categories := models.GetDataCategories()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/categories/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := models.CreateDataCategory(category); !ok {
			temp, _ := template.ParseFiles("views/categories/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/categories/edit.html")
		if err != nil {
			panic(err)
		}

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category := models.DetailDataCategory(id)
		data := map[string]any{
			"category": category,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var category entities.Category

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := models.UpdateDataCategory(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := models.DeleteDataCategory(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
