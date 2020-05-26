package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product https://
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"` // Skip in curl response
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products https:// Array of Product
type Products []*Product

// GetProductsInterface: returns a list of products
// Can swap out hard-coded values with DB i future
func GetProductsInterface() Products {
	return productList
}

// Encoding
func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Ayman",
		Description: "Golang programmer",
		Price:       3.9,
		SKU:         "abc123",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Gagan",
		Description: "Python",
		Price:       5.4,
		SKU:         "xyz123",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
