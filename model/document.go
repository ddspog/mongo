package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Document it's a simples implementation of Documenter. Can be
// embedded to another struct. It contains some attributes important
// to any document on MongoDB.
type Document struct {
	IdV        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedOnV int64         `json:"created_on" bson:"created_on"`
	UpdatedOnV int64         `json:"updated_on" bson:"updated_on"`
}

var (
	// nowInMilli it's stores imported calculation of time in int64 in
	// Millisecond unit for mocking purposes.
	nowInMilli func() int64 = func() (t int64) {
		t = time.Now().UnixNano() / int64(time.Millisecond)
		return
	}
	// newId it's stores imported generation of new ids for documents
	// for mocking purposes.
	newId func() bson.ObjectId = func() (id bson.ObjectId) {
		id = bson.NewObjectId()
		return
	}
)

// Id returs the _id attribute of a Document.
func (p *Document) Id() (id bson.ObjectId) {
	id = p.IdV
	return
}

// SetId set the _id attribute of a Document.
func (p *Document) SetId(id bson.ObjectId) {
	p.IdV = id
}

// CreatedOn returs the created_on attribute of a Document.
func (p *Document) CreatedOn() (t int64) {
	t = p.CreatedOnV
	return
}

// SetCreatedOn set the created_on attribute of a Document.
func (p *Document) SetCreatedOn(t int64) {
	p.CreatedOnV = t
}

// UpdatedOn returs the updated_on attribute of a Document.
func (p *Document) UpdatedOn() (t int64) {
	t = p.UpdatedOnV
	return
}

// SetUpdatedOn set the updated_on attribute of a Document.
func (p *Document) SetUpdatedOn(t int64) {
	p.UpdatedOnV = t
}

// GenerateId creates a new id for a document.
func (p *Document) GenerateId() {
	p.IdV = NewId()
}

// CalculateCreatedOn update the created_on attribute with a value
// corresponding to actual time.
func (p *Document) CalculateCreatedOn() {
	p.CreatedOnV = NowInMilli()
}

// CalculateUpdatedOn update the updated_on attribute with a value
// corresponding to actual time.
func (p *Document) CalculateUpdatedOn() {
	p.UpdatedOnV = NowInMilli()
}

// NowInMilli returns the actual time, in a int64 value in Millisecond
// unit, used by the updaters of created_on and updated_on.
func NowInMilli() (t int64) {
	t = nowInMilli()
	return
}

// NewId generates a new id for documents.
func NewId() (id bson.ObjectId) {
	id = newId()
	return
}
