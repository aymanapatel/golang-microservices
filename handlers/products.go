package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AymanArif/golang-microservices/data"
)

//Products https://
type Products struct {
	l *log.Logger
}

//NewProducts https://
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	/* Marshalling */
	lp := data.GetProductsInterface()
	d, err := json.Marshal(lp) // Marshalling. READ MARSHAL(SLOWER) VS ENCODER(FAST)

	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}

	rw.Write(d)

}
