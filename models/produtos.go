package models

import (
	"fmt"
	"loja/db"
	"strings"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts(sort string) ([]Product, error) {
	db := db.ConnectDB()

	sortMap := map[string]string{
		"nome_asc":        "name ASC",
		"preco_asc":       "price ASC",
		"quantidade_asc":  "quantity ASC",
		"nome_desc":       "name DESC",
		"descricao_desc":  "description DESC",
		"preco_desc":      "price DESC",
		"quantidade_desc": "quantity DESC",
	}

	orderBy := sortMap[strings.ToLower(sort)]

	if orderBy == "" {
		orderBy = "id ASC"
	}

	query := fmt.Sprintf("SELECT * FROM products ORDER BY %s", orderBy)
	selectAllProducts, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer selectAllProducts.Close()

	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			return nil, err
		}

		p := Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		}

		products = append(products, p)
	}

	return products, nil
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	insertDataInDB, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataInDB.Exec(name, description, price, quantity)
	defer db.Close()

}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()

}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	productOfDB, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productForUpdate := Product{}

	for productOfDB.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productOfDB.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		productForUpdate.Id = id
		productForUpdate.Name = name
		productForUpdate.Description = description
		productForUpdate.Price = price
		productForUpdate.Quantity = quantity
	}
	defer db.Close()
	return productForUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDB()

	UpdateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	UpdateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()
}
