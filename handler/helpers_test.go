package handler

import (
	"time"

	"github.com/ddspog/mongo/model"
	"github.com/globalsign/mgo/bson"
)

const (
	// Const values to help tests legibility.
	testid     = "000000000000746573746964"
	product1id = "000070726f64756374316964"
	product2id = "000070726f64756374326964"
	anyReason  = "Whatever reason."
)

// productCollection mimics a collection of products.
var productCollection = []*product{
	newProductStored(),
	newProductStored(),
}

// product it's a simples implementation of Documenter. It contains
// some attributes important to any document on MongoDB.
type product struct {
	IDV        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedOnV int64         `json:"created_on" bson:"created_on"`
	UpdatedOnV int64         `json:"updated_on" bson:"updated_on"`
}

// newProduct returns a empty product.
func newProduct() (p *product) {
	p = &product{}
	return
}

// newProductStored returns a product with some attributes.
func newProductStored() (p *product) {
	p = newProduct()
	p.GenerateID()
	time.Sleep(10 * time.Millisecond)

	p.CalculateCreatedOn()
	return
}

// New creates a new instance of the same Product, used on another
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
