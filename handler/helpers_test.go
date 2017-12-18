package handler

import (
	"time"

	"github.com/ddspog/mongo/model"
)

const (
	// Const values to help tests legibility.
	testid     = "000000000000746573746964"
	product1id = "000070726f64756374316964"
	product2id = "000070726f64756374326964"
	anyReason  = "Whatever reason."
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
	p.GenerateID()
	time.Sleep(10 * time.Millisecond)

	p.CalculateCreatedOn()
	return
}

// Finish call Finish() from all objects received.
func Finish(fs ...interface {
	Finish()
}) {
	for i := range fs {
		fs[i].Finish()
	}
}

// timeFmt parses time well formatted.
func timeFmt(s string) (t time.Time) {
	t, _ = time.Parse("02-01-2006 15:04:05", s)
	return
}

// expectedNowInMilli returns expected return from NowInMilli function,
// given the time returned by time.Now().
func expectedNowInMilli(t time.Time) (r int64) {
	r = t.UnixNano() / int64(time.Millisecond)
	return
}
