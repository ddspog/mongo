package handler

import (
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/model"
)

// productHandler it's an interface describing operations common to
// handler's of MongoDB products collection.
type productHandler interface {
	Link(elements.Databaser) *productHandle
	Count() (int, error)
	Find() (*product, error)
	FindAll() ([]*product, error)
	Insert() error
	Remove() error
	RemoveAll() (*elements.ChangeInfo, error)
	Update() error
	Document() *product
	Name() string
}

// productHandle it's a type embedding the Handle struct, it's capable
// of storing products.
type productHandle struct {
	*Handle
	DocumentV *product
}

// newProductHandle returns a empty productHandle.
func newProductHandle() (p *productHandle) {
	p = &productHandle{
		Handle:    New(),
		DocumentV: newProduct(),
	}
	return
}

// Name returns the name of connection that productHandle can connect.
func (p *productHandle) Name() (n string) {
	n = "products"
	return
}

// Link connects the productHandle to collection.
func (p *productHandle) Link(db elements.Databaser) (h *productHandle) {
	p.Handle.Link(db, p.Name())
	h = p
	return
}

// Find search on connected collection for a document matching data
// stored on productHandle and returns it.
func (p *productHandle) Find() (prod *product, err error) {
	var doc model.Documenter = newProduct()
	err = p.Handle.Find(p.Document(), &doc)
	prod = doc.(*product)
	return
}

// FindAll search on connected collection for all documents matching
// data stored on productHandle and returns it.
func (p *productHandle) FindAll() (proda []*product, err error) {
	var da []model.Documenter
	err = p.Handle.FindAll(p.Document(), func() model.Documenter {
		return newProduct()
	}, &da)
	proda = make([]*product, len(da))
	for i := range da {
		//noinspection GoNilContainerIndexing
		proda[i] = da[i].(*product)
	}
	return
}

// Insert creates a new document with data stored on productHandle
// and put on connected collection.
func (p *productHandle) Insert() (err error) {
	err = p.Handle.Insert(p.Document())
	return
}

// Remove delete a document from connected collection matching the id
// of data stored on Handle.
func (p *productHandle) Remove() (err error) {
	err = p.Handle.Remove(p.Document())
	return
}

// Remove deletes all document from connected collection matching the
// data stored on Handle.
func (p *productHandle) RemoveAll() (info *elements.ChangeInfo, err error) {
	info, err = p.Handle.RemoveAll(p.Document())
	return
}

// Update updates document from connected collection matching the id
// of data stored on Handle, and uses further info to update.
func (p *productHandle) Update() (err error) {
	err = p.Handle.Update(p.Document())
	return
}

// Document returns the Document of Handle with correct type.
func (p *productHandle) Document() (d *product) {
	d = p.DocumentV
	return
}
