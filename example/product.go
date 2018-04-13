package example

import (
	"github.com/ddspog/mongo/model"
	"github.com/globalsign/mgo/bson"
)

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
