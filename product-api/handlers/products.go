// Package classification of Product APIVersion
//
// Documentation for Product APIVersion
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bugg123/golang-microservices/data"
	"github.com/gorilla/mux"
)

type KeyProduct struct{}

// Products is a http.Handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts returns a new Products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Message string `json:"message"`
}

func getProductID(r *http.Request) int {
	vars := mux.Vars(r)

	// convert the id into an integer and return nil, err
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}
	return id
}
