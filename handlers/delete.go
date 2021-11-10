package handlers

import (
	"net/http"
	"strconv"

	"github.com/bugg123/golang-microservices/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the database
// responses:
//   201: noContent

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	 vars := mux.Vars(r)
	 id, _ := strconv.Atoi(vars["id"])

	 p.l.Println("Handle DELETE Product", id)

	 err := data.DeleteProduct(id)

	 if err == data.ErrProductNotFound {
		 http.Error(w, "Product not found", http.StatusNotFound)
		 return
	 }
	 if err != nil {
		 http.Error(w, "Product not found", http.StatusInternalServerError)
		 return
	 }
}
