package handlers

import (
	"net/http"
	"strconv"

	"github.com/bugg123/golang-microservices/data"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Products", id)
	prod, ok := r.Context().Value(KeyProduct{}).(*data.Product)
	if !ok {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
