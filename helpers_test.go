package mongo

import (
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/handler"
)

// productCounter it's a helping handler that only serves to count
// documents on collection.
type productCounter interface {
	Name() string
	Link(elements.Databaser) productCounter
	Count() (int, error)
}

// productCount it's a handle that implements productCounter.
type productCount struct {
	*handler.Handle
	name string
}

// newProductCount returns a empty productCount.
func newProductCount(n string) (p *productCount) {
	p = &productCount{
		Handle: handler.New(),
		name:   n,
	}
	return
}

// Name returns the name of connection that productHandle can connect.
func (p *productCount) Name() (n string) {
	n = p.name
	return
}

// Link connects the productHandle to collection.
func (p *productCount) Link(db elements.Databaser) (h productCounter) {
	p.Handle.Link(db, p.Name())
	h = p
	return
}

// finisher defines a type that can be finished, closing all pendant
// operations.
type finisher interface {
	Finish()
}

// finish calls Finish for all finishers received.
func finish(fs ...finisher) {
	for _, f := range fs {
		f.Finish()
	}
}
