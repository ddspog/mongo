package example

import "github.com/ddspog/mongo/model"

// Product it's a type embedding the Document struct.
type Product struct {
	*model.Document
}

// NewProduct returns a empty Product.
func NewProduct() (p Product) {
	p = Product{
		Document: &model.Document{},
	}
	return
}
