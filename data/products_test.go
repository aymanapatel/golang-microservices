package data

import "testing"

func TestCheckValidation(t *testing.T) {

	p := &Product{
		Name:  "ayman",
		Price: 1.0,
		SKU:   "avd-vd-fdfd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
