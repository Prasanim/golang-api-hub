package handlers

import (
	"log"
	"net/http"

	"github.com/Prasanim/golang-api-hub/product.api/data"
)

// Product
type Products struct {
	l *log.Logger
}

//List New product
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Export it in http
func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshali(lp)
	if err != nil {
		http.Error(rw, "Undable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
