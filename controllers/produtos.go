package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"loja/models"
	"loja/utils"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	sortParam := r.URL.Query().Get("sort")

	if !utils.IsValidSort(sortParam) {
		sortParam = ""
	}

	allProducts, err := models.GetAllProducts(sortParam)
	if err != nil {
		http.Error(w, "Erro ao buscar products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		qtyProducts, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantity:", err)
		}

		models.CreateNewProduct(name, description, priceFloat, qtyProducts)
	}
	http.Redirect(w, r, "/produtos", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/produtos", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float64:", err)
		}

		qtyProducts, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da quantity para int:", err)
		}

		models.UpdateProduct(idInt, name, description, priceFloat, qtyProducts)
	}
	http.Redirect(w, r, "/produtos", http.StatusSeeOther)
}
