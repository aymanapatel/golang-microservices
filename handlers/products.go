// Package classification of Product API.
//
// Documenting for Product API
//
//
//
// Schemes: http, https
// BasePath: /
// Version: 0.0.1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AymanArif/golang-microservices/data"
	"github.com/gorilla/mux"
)

//Products https://
type Products struct {
	l *log.Logger
}

//NewProducts https://
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Println: Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProductsInterface()

	// serialize the list to JSON
	err := lp.ToJSON(rw) // READ MARSHAL(SLOWER) VS ENCODER(FAST)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// postProducts returns the products from the data store
func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Println: Handle POST Products")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

type KeyProduct struct{}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Update Request Read from Regex")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		p.l.Panicln(err)
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product")

	// MiddleWarej
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

// Context:
// Use type(KeyProduct{}) instead of String.
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := prod.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserialization", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// Validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add the product to the context.
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain
		// or the final handler
		next.ServeHTTP(rw, r)

	})
}
