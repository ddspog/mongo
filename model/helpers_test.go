package model

const (
	// Const values to help tests legibility.
	testid     = "000000000000746573746964"
	product1id = "000070726f64756374316964"
	product2id = "000070726f64756374326964"
)

// product it's a type embedding the Document struct.
type product struct {
	*Document
}

// newProduct returns a empty product.
func newProduct() (p product) {
	p = product{
		Document: &Document{},
	}
	return
}
