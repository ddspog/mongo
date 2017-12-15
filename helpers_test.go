package mongo

import (
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/handler"
)

type productCounter interface {
	Name() string
	Link(elements.Databaser) productCounter
	Count() (int, error)
}

type productCount struct {
	*handler.Handle
}

// newProductCount returns a empty productCount.
func newProductCount() (p *productCount) {
	p = &productCount{
		Handle: handler.New(),
	}
	return
}

// Name returns the name of connection that productHandle can connect.
func (p *productCount) Name() (n string) {
	n = "products"
	return
}

// Link connects the productHandle to collection.
func (p *productCount) Link(db elements.Databaser) (h productCounter) {
	p.Handle.Link(db, p.Name())
	h = p
	return
}
