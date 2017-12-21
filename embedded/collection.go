package embedded

import (
	"github.com/ddspog/mongo/elements"
	"gopkg.in/mgo.v2"
)

// Collection it's an embedded type of mgo.Collection, made to use the
// interfaces it implements, Collectioner, on it's methods signatures.
type Collection struct {
	*mgo.Collection
}

// Find prepares a query using the provided document.  The document may be a
// map or a struct value capable of being marshalled with bson.  The map
// may be a generic one using interface{} for its key and/or values, such as
// bson.M, or it may be a properly typed map.  Providing nil as the document
// is equivalent to providing an empty document such as bson.M{}.
//
// Further details of the query may be tweaked using the resulting Query value,
// and then executed to retrieve results using methods such as One, For,
// Iter, or Tail.
//
// In case the resulting document includes a field named $err or errmsg, which
// are standard ways for MongoDB to return query errors, the returned err will
// be set to a *QueryError value including the Err message and the Code.  In
// those cases, the result argument is still unmarshalled into with the
// received document so that any other custom values may be obtained if
// desired.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Querying
//     http://www.mongodb.org/display/DOCS/Advanced+Queries
//
func (c *Collection) Find(q interface{}) elements.Querier {
	return &Query{c.Collection.Find(q)}
}

// FindID is a convenience helper equivalent to:
//
//     query := collection.Find(bson.M{"_id": id})
//
// See the Find method for more details.
func (c *Collection) FindID(id interface{}) elements.Querier {
	return &Query{c.Collection.FindId(id)}
}

// RemoveID is a convenience helper equivalent to:
//
//     err := collection.Remove(bson.M{"_id": id})
//
// See the Remove method for more details.
func (c *Collection) RemoveID(id interface{}) error {
	err := c.Collection.RemoveId(id)
	return err
}

// RemoveAll finds all documents matching the provided selector document
// and removes them from the database.  In case the session is in safe mode
// (see the SetSafe method) and an error happens when attempting the change,
// the returned error will be of type *LastError.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Removing
//
func (c *Collection) RemoveAll(s interface{}) (*elements.ChangeInfo, error) {
	ci, err := c.Collection.RemoveAll(s)
	return &elements.ChangeInfo{ChangeInfo: ci}, err
}

// UpdateID is a convenience helper equivalent to:
//
//     err := collection.Update(bson.M{"_id": id}, update)
//
// See the Update method for more details.
func (c *Collection) UpdateID(id interface{}, u interface{}) error {
	err := c.Collection.UpdateId(id, u)
	return err
}

// UpdateAll finds all documents matching the provided selector document
// and modifies them according to the update document.
// If the session is in safe mode (see SetSafe) details of the executed
// operation are returned in info or an error of type *LastError when
// some problem is detected. It is not an error for the update to not be
// applied on any documents because the selector doesn't match.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func (c *Collection) UpdateAll(s interface{}, u interface{}) (*elements.ChangeInfo, error) {
	ci, err := c.Collection.UpdateAll(s, u)
	return &elements.ChangeInfo{ChangeInfo: ci}, err
}

// Upsert finds a single document matching the provided selector document
// and modifies it according to the update document.  If no document matching
// the selector is found, the update document is applied to the selector
// document and the result is inserted in the collection.
// If the session is in safe mode (see SetSafe) details of the executed
// operation are returned in info, or an error of type *LastError when
// some problem is detected.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Updating
//     http://www.mongodb.org/display/DOCS/Atomic+Operations
//
func (c *Collection) Upsert(s interface{}, u interface{}) (*elements.ChangeInfo, error) {
	ci, err := c.Collection.Upsert(s, u)
	return &elements.ChangeInfo{ChangeInfo: ci}, err
}

// UpsertID is a convenience helper equivalent to:
//
//     info, err := collection.Upsert(bson.M{"_id": id}, update)
//
// See the Upsert method for more details.
func (c *Collection) UpsertID(id interface{}, u interface{}) (*elements.ChangeInfo, error) {
	ci, err := c.Collection.UpsertId(id, u)
	return &elements.ChangeInfo{ChangeInfo: ci}, err
}

// With returns a copy of c that uses session s.
func (c *Collection) With(s *mgo.Session) elements.Collectioner {
	return &Collection{c.Collection.With(s)}
}