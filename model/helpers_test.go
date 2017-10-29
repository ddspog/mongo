package model

const (
	// Const values to help tests legibility.
	testid     string = "000000000000746573746964"
	product1id string = "000070726f64756374316964"
	product2id string = "000070726f64756374326964"
	before     int64  = NowInInt64()
)

// product it's a type embedding the Document struct.
type product struct {
	*Document
	name  string
	price float32
}

// newProduct returns a empty product.
func newProduct() (p product) {
	p = product{
		Document: &Document{},
	}
	return
}
