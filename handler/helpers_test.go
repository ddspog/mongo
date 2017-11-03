package handler

import (
	"time"

	"github.com/ddspog/mongo/model"
)

const (
	// Const values to help tests legibility.
	testid     string = "000000000000746573746964"
	product1id string = "000070726f64756374316964"
	product2id string = "000070726f64756374326964"
	anyReason  string = "Whatever reason."
)

// productCollection mimics a collection of products.
var productCollection = []product{
	newProductStored(),
	newProductStored(),
}

// product it's a type embedding the Document struct.
type product struct {
	*model.Document
}

// newProduct returns a empty product.
func newProduct() (p product) {
	p = product{
		Document: &model.Document{},
	}
	return
}

// newProductStored returns a product with some attributes.
func newProductStored() (p product) {
	p = newProduct()
	p.GenerateId()
	time.Sleep(10 * time.Millisecond)

	p.CalculateCreatedOn()
	return
}


// Finish call Finish() from all objects received.
func Finish(fs ...interface{Finish()}) {
	for i := range fs {
		fs[i].Finish()
	}
}