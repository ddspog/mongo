package example

import (
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/handler"
	"github.com/ddspog/mongo/model"
	"github.com/globalsign/mgo/bson"
)

// ProductHandler it's an interface describing operations common to
// handler's of MongoDB Products collection.
type ProductHandler interface {
	Link(elements.Databaser) (ProductHandler, error)
	Count() (int, error)
	Find() (*Product, error)
	FindAll() ([]*Product, error)
	Insert() error
	Remove() error
	RemoveAll() (*elements.ChangeInfo, error)
	Update(bson.ObjectId) error
	Document() *Product
	Name() string
}

// ProductHandle it's a type embedding the Handle struct, it's capable
// of storing Products.
type ProductHandle struct {
	*handler.Handle
	DocumentV *Product
}

// NewProductHandle returns a empty ProductHandle.
func NewProductHandle() (p *ProductHandle) {
	p = &ProductHandle{
		Handle:    handler.New(),
		DocumentV: NewProduct(),
	}
	return
}

// Name returns the name of connection that ProductHandle can connect.
func (p *ProductHandle) Name() (n string) {
	n = "products"
	return
}

// Link connects the ProductHandle to collection.
func (p *ProductHandle) Link(db elements.Databaser) (h ProductHandler, err error) {
	err = p.Handle.Link(db, p.Name())
	h = p
	return
}

// Find search on connected collection for a document matching data
// stored on ProductHandle and returns it.
func (p *ProductHandle) Find() (prod *Product, err error) {
	var doc model.Documenter = NewProduct()
	err = p.Handle.Find(p.Document(), &doc)
	prod = doc.(*Product)
	return
}

// FindAll search on connected collection for all documents matching
// data stored on ProductHandle and returns it.
func (p *ProductHandle) FindAll() (proda []*Product, err error) {
	var da []model.Documenter
	err = p.Handle.FindAll(p.Document(), func() model.Documenter {
		return NewProduct()
	}, &da)
	proda = make([]*Product, len(da))
	for i := range da {
		//noinspection GoNilContainerIndexing
		proda[i] = da[i].(*Product)
	}
	return
}

// Insert creates a new document with data stored on ProductHandle
// and put on connected collection.
func (p *ProductHandle) Insert() (err error) {
	err = p.Handle.Insert(p.Document())
	return
}

// Remove delete a document from connected collection matching the id
// of data stored on Handle.
func (p *ProductHandle) Remove() (err error) {
	err = p.Handle.Remove(p.Document())
	return
}

// Remove deletes all document from connected collection matching the
// data stored on Handle.
func (p *ProductHandle) RemoveAll() (info *elements.ChangeInfo, err error) {
	info, err = p.Handle.RemoveAll(p.Document())
	return
}

// Update updates document from connected collection matching the id
// received, and uses document info to update.
func (p *ProductHandle) Update(id bson.ObjectId) (err error) {
	err = p.Handle.Update(id, p.Document())
	return
}

// Document returns the Document of Handle with correct type.
func (p *ProductHandle) Document() (d *Product) {
	d = p.DocumentV
	return
}
