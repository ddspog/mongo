package model

import "gopkg.in/mgo.v2/bson"

// Documenter it's an interface that could be common to any documents
// types used to store values on a MongoDB. It contains getters and
// generaters to important documents values: _id, created_on and
// updated_on
type Documenter interface {
	ID() bson.ObjectId
	SetID(id bson.ObjectId)
	CreatedOn() int64
	SetCreatedOn(t int64)
	UpdatedOn() int64
	SetUpdatedOn(t int64)
	GenerateID()
	CalculateCreatedOn()
	CalculateUpdatedOn()
}
