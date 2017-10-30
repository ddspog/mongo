package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Document it's a simples implementation of Documenter. Can be
// embedded to another struct. It contains some attributes important
// to any document on MongoDB.
type Document struct {
	IdV	bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedOnV	int64 `json:"created_on" bson:"created_on"`
	UpdatedOnV	int64 `json:"updated_on" bson:"updated_on"`
}

// Id returs the _id attribute of a Document.
func (p *Document) Id() (id bson.ObjectId) {
	id = p.IdV
	return
}

// CreatedOn returs the created_on attribute of a Document.
func (p *Document) CreatedOn() (t int64) {
	t = p.CreatedOnV
	return
}

// UpdatedOn returs the updated_on attribute of a Document.
func (p *Document) UpdatedOn() (t int64) {
	t = p.UpdatedOnV
	return
}

// GenerateId creates a new id for a document.
func (p *Document) GenerateId() {
	p.IdV = bson.NewObjectId()
}

// CalculateCreatedOn update the created_on attribute with a value
// corresponding to actual time.
func (p *Document) CalculateCreatedOn() {
	p.CreatedOnV = NowInInt64()
}

// CalculateUpdatedOn update the updated_on attribute with a value
// corresponding to actual time.
func (p *Document) CalculateUpdatedOn() {
	p.UpdatedOnV = NowInInt64()
}

// NowInInt64 returns the actual time, in a int64 value, used by the
// updaters of created_on and updated_on.
func NowInInt64() (t int64) {
	t = time.Now().UnixNano() / int64(time.Millisecond)
	return
}
