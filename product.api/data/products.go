package data

import "time"

// Product defines the structure for an API product
type Product struct {
	ID				int
	Name 			string
	Description		string
	Price			string
	SKU				float32
	CreatedOn		string
	UpdatedOn		string
	DeletedOn		string
}

var productList = [] * Product{
	&Product{
		ID			:	1
		Name		:	"Latte",
		Description	:	"Frothy milky coffe",
		Price		:	2.45
		SKU			:	"abc323",
		CreatedOn	:	time.Now().UTC().String(),
		UpdatedOn	:	time.Now().UTC().String()
	},
	&Product{
		ID			:	2
		Name		:	"Espresso",
		Description	:	"Short and strong coffe without milk",
		Price		:	1.99
		SKU			:	"fjd34",
		CreatedOn	:	time.Now().UTC().String(),
		UpdatedOn	:	time.Now().UTC().String()
	}
}