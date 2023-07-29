package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	connection := "user=postgres dbname=loja password=postgresql host=db sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateProductsTable() error {
	db := ConnectDB()

	createTableIfNotExist := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			price NUMERIC(10, 2) NOT NULL,
			quantity INT NOT NULL
		)
	`

	_, err := db.Exec(createTableIfNotExist)
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func PopulateProductsTable() error {
	db := ConnectDB()

	products := []struct {
		name        string
		description string
		price       float64
		quantity    int
	}{
		{"Produto 1", "Descrição do Produto 1", 10.99, 100},
		{"Produto 2", "Descrição do Produto 2", 20.50, 50},
		{"Produto 3", "Descrição do Produto 3", 5.99, 200},
	}

	insertDataInDb, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	for _, p := range products {
		_, err = insertDataInDb.Exec(p.name, p.description, p.price, p.quantity)
		if err != nil {
			return err
		}
	}

	defer db.Close()

	return nil
}
