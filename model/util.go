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

// InitDocumenter translates a bson.M received, to the Documenter
// structure received as a pointer. It fills the structure fields with
// the values of each key in the bson.M received.
func InitDocumenter(in bson.M, out *Documenter) (err error) {
	var marshalled []byte

	if marshalled, err = bson.Marshal(in); err == nil {
		err = bson.Unmarshal(marshalled, *out)
	}
	return
}

// MapDocumenter translates a Documenter in whathever structure
// it has, to a bson.M object, more easily read by mgo.Collection
// methods.
func MapDocumenter(in Documenter) (out bson.M, err error) {
	var buf []byte
	var target interface{}

	if buf, err = bson.Marshal(in); err == nil {
		if err = bson.Unmarshal(buf, &target); err == nil {
			out = target.(bson.M)
		}
	}

	return
}
