package model

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

var (
	// now it's stores imported calculation of time for mocking
	// purposes.
	now = time.Now
	// newID it's stores imported generation of new ids for documents
	// for mocking purposes.
	newID = bson.NewObjectId
)

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
