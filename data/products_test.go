package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "nice",
		Price: 1.00,
		SKU: "asdf-asdf-asdf",
	}

	err := p.Validate()
	if err != nil {
		t.Error(err)
	}
}
