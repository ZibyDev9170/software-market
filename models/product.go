package models

import (
	"database/sql"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
}

func AllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query(`SELECT id, name, price, description FROM "Product"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func GetProduct(db *sql.DB, id string) (Product, error) {
	var p Product
	row := db.QueryRow(`SELECT id, name, price, description FROM "Product" WHERE id = $1`, id)
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description); err != nil {
		return p, err
	}
	return p, nil
}
