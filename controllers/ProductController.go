package controllers

import (
	"go-web-native/entities"
	"go-web-native/models"
	"html/template"
	"net/http"
	"strconv"
	"time"
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
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/products/create.html")
		if err != nil {
			panic(err)
		}

		categories := models.GetDataCategories()
		data := map[string]any{
			"categories": categories,
		}

		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		categoryId, err := strconv.Atoi(r.FormValue("category_id"))
		if err != nil {
			panic(err)
		}

		stock, err := strconv.Atoi(r.FormValue("stock"))
		if err != nil {
			panic(err)
		}

		product.Name = r.FormValue("name")
		product.Category.Id = uint(categoryId)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := models.CreateDataProduct(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func DetailProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := models.DetailDataProduct(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/products/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
