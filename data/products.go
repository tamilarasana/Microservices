package data

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	SKU         string
	CreatedOn   string
	DeletedOn   string
}

func GetProducts() []*Product {
	return productList

}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "tamil",
		Description: "milk coffe",
		Price:       250,
		SKU:         "abcdes",
		CreatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "tamil",
		Description: "milk tea",
		Price:       200,
		SKU:         "abcdefgh",
		CreatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
