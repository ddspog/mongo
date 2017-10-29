package model

import "gopkg.in/mgo.v2/bson"

// Documenter it's an interface that could be common to any documents
// types used to store values on a MongoDB. It contains getters and
// generaters to important documents values: _id, created_on and
// updated_on
type Documenter interface {
	Id() bson.ObjectId
	CreatedOn() int64
	UpdatedOn() int64
	GenerateId()
	CalculateCreatedOn()
	CalculateUpdatedOn()
}