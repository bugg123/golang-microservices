package client

import (
	"testing"

	"github.com/bugg123/golang-microservices/product-api/sdk/client/products"
	"github.com/stretchr/testify/assert"
)


func TestListProducts(t *testing.T) {
	cfg := DefaultTransportConfig()
	cfg.WithHost("localhost:9090")
	c := NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	assert.NoError(t, err)

	prod.GetPayload()
}
