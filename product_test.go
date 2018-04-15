package mongo

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
	err = p.Handle.Find(p.Document(), doc)
	prod = doc.(*Product)
	return
}

// FindAll search on connected collection for all documents matching
// data stored on ProductHandle and returns it.
func (p *ProductHandle) FindAll() (proda []*Product, err error) {
	var da []model.Documenter
	err = p.Handle.FindAll(p.Document(), &da)
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

// Product it's a type embedding the Document struct.
type Product struct {
	IDV        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedOnV int64         `json:"created_on,omitempty" bson:"created_on,omitempty"`
	UpdatedOnV int64         `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
}

// NewProduct returns a empty Product.
func NewProduct() (p *Product) {
	p = &Product{}
	return
}

// New creates a new instance of the same Product, used on another
// functions for clone purposes.
func (p *Product) New() (doc model.Documenter) {
	doc = NewProduct()
	return
}

// Map translates a product to a bson.M object, more easily read by mgo
// methods.
func (p *Product) Map() (out bson.M, err error) {
	out, err = model.MapDocumenter(p)
	return
}

// Init translates a bson.M received, to the product strucutre. It
// fills the structure fields with the values of each key in the
// bson.M received.
func (p *Product) Init(in bson.M) (err error) {
	var doc model.Documenter = p
	err = model.InitDocumenter(in, &doc)
	return
}

// ID returns the _id attribute of a Document.
func (p *Product) ID() (id bson.ObjectId) {
	id = p.IDV
	return
}

// CreatedOn returns the created_on attribute of a Document.
func (p *Product) CreatedOn() (t int64) {
	t = p.CreatedOnV
	return
}

// UpdatedOn returns the updated_on attribute of a Document.
func (p *Product) UpdatedOn() (t int64) {
	t = p.UpdatedOnV
	return
}

// GenerateID creates a new id for a document.
func (p *Product) GenerateID() {
	p.IDV = model.NewID()
}

// CalculateCreatedOn update the created_on attribute with a value
// corresponding to actual time.
func (p *Product) CalculateCreatedOn() {
	p.CreatedOnV = model.NowInMilli()
}

// CalculateUpdatedOn update the updated_on attribute with a value
// corresponding to actual time.
func (p *Product) CalculateUpdatedOn() {
	p.UpdatedOnV = model.NowInMilli()
}
