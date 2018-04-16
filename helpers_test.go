package mongo

import (
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/handler"
	"github.com/ddspog/mongo/model"
	"github.com/globalsign/mgo/bson"
)

// product it's a type embedding the Document struct.
type product struct {
	IDV        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedOnV int64         `json:"created_on,omitempty" bson:"created_on,omitempty"`
	UpdatedOnV int64         `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
}

// newProduct returns a empty product.
func newProduct() (p *product) {
	p = &product{}
	return
}

// New creates a new instance of the same product, used on another
// functions for clone purposes.
func (p *product) New() (doc model.Documenter) {
	doc = newProduct()
	return
}

// Map translates a product to a bson.M object, more easily read by mgo
// methods.
func (p *product) Map() (out bson.M, err error) {
	out, err = model.MapDocumenter(p)
	return
}

// Init translates a bson.M received, to the product structure. It
// fills the structure fields with the values of each key in the
// bson.M received.
func (p *product) Init(in bson.M) (err error) {
	var doc model.Documenter = p
	err = model.InitDocumenter(in, &doc)
	return
}

// ID returns the _id attribute of a Document.
func (p *product) ID() (id bson.ObjectId) {
	id = p.IDV
	return
}

// CreatedOn returns the created_on attribute of a Document.
func (p *product) CreatedOn() (t int64) {
	t = p.CreatedOnV
	return
}

// UpdatedOn returns the updated_on attribute of a Document.
func (p *product) UpdatedOn() (t int64) {
	t = p.UpdatedOnV
	return
}

// GenerateID creates a new id for a document.
func (p *product) GenerateID() {
	p.IDV = model.NewID()
}

// CalculateCreatedOn update the created_on attribute with a value
// corresponding to actual time.
func (p *product) CalculateCreatedOn() {
	p.CreatedOnV = model.NowInMilli()
}

// CalculateUpdatedOn update the updated_on attribute with a value
// corresponding to actual time.
func (p *product) CalculateUpdatedOn() {
	p.UpdatedOnV = model.NowInMilli()
}

// productHandler it's an interface describing operations common to
// handler's of MongoDB Products collection.
type productHandler interface {
	Link(elements.Databaser) (productHandler, error)
	Count() (int, error)
	Find() (*product, error)
	FindAll() ([]*product, error)
	Insert() error
	Remove() error
	RemoveAll() (*elements.ChangeInfo, error)
	Update(bson.ObjectId) error
	Document() *product
	Name() string
}

// productHandle it's a type embedding the Handle struct, it's capable
// of storing Products.
type productHandle struct {
	*handler.Handle
	DocumentV *product
}

// newProductHandle returns a empty productHandle.
func newProductHandle() (p *productHandle) {
	p = &productHandle{
		Handle:    handler.New(),
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
func (p *productHandle) Link(db elements.Databaser) (h productHandler, err error) {
	err = p.Handle.Link(db, p.Name())
	h = p
	return
}

// Find search on connected collection for a document matching data
// stored on productHandle and returns it.
func (p *productHandle) Find() (prod *product, err error) {
	var doc model.Documenter = newProduct()
	err = p.Handle.Find(p.Document(), doc)
	prod = doc.(*product)
	return
}

// FindAll search on connected collection for all documents matching
// data stored on productHandle and returns it.
func (p *productHandle) FindAll() (proda []*product, err error) {
	var da []model.Documenter
	err = p.Handle.FindAll(p.Document(), &da)
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
	err = p.Handle.Remove(p.Document().ID())
	return
}

// Remove deletes all document from connected collection matching the
// data stored on Handle.
func (p *productHandle) RemoveAll() (info *elements.ChangeInfo, err error) {
	info, err = p.Handle.RemoveAll(p.Document())
	return
}

// Update updates document from connected collection matching the id
// received, and uses document info to update.
func (p *productHandle) Update(id bson.ObjectId) (err error) {
	err = p.Handle.Update(id, p.Document())
	return
}

// Document returns the Document of Handle with correct type.
func (p *productHandle) Document() (d *product) {
	d = p.DocumentV
	return
}

// finisher defines a type that can be finished, closing all pendant
// operations.
type finisher interface {
	Finish()
}

// finish calls finish for all finishers received.
func finish(fs ...finisher) {
	for _, f := range fs {
		f.Finish()
	}
}

func newDBSocket() (db databaseSocketer) {
	db = &databaseSocket{
		db:   make(chan elements.Databaser),
		quit: make(chan bool),
	}
	return
}

type databaseSocketer interface {
	DB() elements.Databaser
	Close()
}

type databaseSocket struct {
	db   chan elements.Databaser
	quit chan bool
}

func (d *databaseSocket) DB() (db elements.Databaser) {
	go ConsumeDatabaseOnSession(func(db elements.Databaser) {
		d.db <- db
		<-d.quit
	})

	return <-d.db
}

func (d *databaseSocket) Close() {
	d.quit <- true
}
