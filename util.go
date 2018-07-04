package mongo

import (
	"time"

	"github.com/ddspog/mongo/internal/bsonutils"
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

// M is a convenient alias for a map[string]interface{} map, useful for
// dealing with BSON in a native way.  For instance:
//
//     M{"a": 1, "b": true}
//
// There's no special handling for this type in addition to what's done anyway
// for an equivalent map type.  Elements in the map will be dumped in an
// undefined ordered.
type M = bson.M

// ObjectId is a unique ID identifying a BSON value. It must be exactly 12 bytes
// long. MongoDB objects by default have such a property set in their "_id"
// property.
//
// http://www.mongodb.org/display/DOCS/Object+Ids
type ObjectId = bson.ObjectId

// ObjectIdHex returns an ObjectId from the provided hex representation.
// Calling this function with an invalid hex representation will
// cause a runtime panic.
func ObjectIdHex(s string) ObjectId {
	return bson.ObjectIdHex(s)
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

// InitDocumenter translates a M received, to the Documenter
// structure received as a pointer. It fills the structure fields with
// the values of each key in the M received.
func InitDocumenter(in M, out *Documenter) (err error) {
	var marshalled []byte

	bsonutils.SetOmitEmptyAsDefault(true)

	if marshalled, err = bsonutils.Marshal(in); err == nil {
		err = bsonutils.Unmarshal(marshalled, *out)
	}

	bsonutils.SetOmitEmptyAsDefault(false)

	return
}

// MapDocumenter translates a Documenter in whatever structure
// it has, to a M object, more easily read by mgo.Collection
// methods.
func MapDocumenter(in Documenter) (out M, err error) {
	var buf []byte
	var target interface{}

	bsonutils.SetOmitEmptyAsDefault(true)

	if buf, err = bsonutils.Marshal(in); err == nil {
		if err = bsonutils.Unmarshal(buf, &target); err == nil {
			out = target.(M)
		}
	}

	bsonutils.SetOmitEmptyAsDefault(false)

	return
}

// MarshalM applies marshal to an M object and returns the buffer
// result and error if any.
func MarshalM(in M) (out []byte, err error) {
	out, err = bsonutils.Marshal(in)
	return
}

// UnmarshalToM applies unmarshal to a new M object, returning with an
// error if received.
func UnmarshalToM(in []byte) (out M, err error) {
	out = M{}
	err = bsonutils.Unmarshal(in, &out)
	return
}
