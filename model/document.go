package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Document it's a simples implementation of Documenter. Can be
// embedded to another struct. It contains some attributes important
// to any document on MongoDB.
type Document struct {
	IDV        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedOnV int64         `json:"created_on" bson:"created_on"`
	UpdatedOnV int64         `json:"updated_on" bson:"updated_on"`
}

var (
	// now it's stores imported calculation of time for mocking
	// purposes.
	now = time.Now
	// newID it's stores imported generation of new ids for documents
	// for mocking purposes.
	newID = bson.NewObjectId
)

// ID returns the _id attribute of a Document.
func (p *Document) ID() (id bson.ObjectId) {
	id = p.IDV
	return
}

// SetID set the _id attribute of a Document.
func (p *Document) SetID(id bson.ObjectId) {
	p.IDV = id
}

// CreatedOn returns the created_on attribute of a Document.
func (p *Document) CreatedOn() (t int64) {
	t = p.CreatedOnV
	return
}

// SetCreatedOn set the created_on attribute of a Document.
func (p *Document) SetCreatedOn(t int64) {
	p.CreatedOnV = t
}

// UpdatedOn returns the updated_on attribute of a Document.
func (p *Document) UpdatedOn() (t int64) {
	t = p.UpdatedOnV
	return
}

// SetUpdatedOn set the updated_on attribute of a Document.
func (p *Document) SetUpdatedOn(t int64) {
	p.UpdatedOnV = t
}

// GenerateID creates a new id for a document.
func (p *Document) GenerateID() {
	p.SetID(NewID())
}

// CalculateCreatedOn update the created_on attribute with a value
// corresponding to actual time.
func (p *Document) CalculateCreatedOn() {
	p.SetCreatedOn(NowInMilli())
}

// CalculateUpdatedOn update the updated_on attribute with a value
// corresponding to actual time.
func (p *Document) CalculateUpdatedOn() {
	p.SetUpdatedOn(NowInMilli())
}

// NowInMilli returns the actual time, in a int64 value in Millisecond
// unit, used by the updaters of created_on and updated_on.
func NowInMilli() (t int64) {
	t = now().UnixNano() / int64(time.Millisecond)
	return
}

// NewID generates a new id for documents.
func NewID() (id bson.ObjectId) {
	id = newID()
	return
}
