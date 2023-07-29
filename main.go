package main

import (
	"loja/db"
	"loja/routes"
	"net/http"
)

func main() {
	err := db.CreateProductsTable()
	if err != nil {
		panic(err)
	}

	err = db.PopulateProductsTable()
	if err != nil {
		panic(err)
	}

	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
