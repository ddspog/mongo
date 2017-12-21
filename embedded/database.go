package embedded

import (
	"github.com/ddspog/mongo/elements"
	"gopkg.in/mgo.v2"
)

// Database it's an embedded type of mgo.Database, made to use the
// interfaces it implements, Databaser, on it's methods signatures.
type Database struct {
	*mgo.Database
}

// C returns a value representing the named collection.
//
// Creating this value is a very lightweight operation, and
// involves no network communication.
func (d *Database) C(n string) elements.Collectioner {
	return &Collection{d.Database.C(n)}
}

// FindRef returns a query that looks for the document in the provided
// reference. If the reference includes the DB field, the document will
// be retrieved from the respective database.
//
// See also the DBRef type and the FindRef method on Session.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Database+References
//
func (d *Database) FindRef(ref *mgo.DBRef) elements.Querier {
	return &Query{d.Database.FindRef(ref)}
}

// With returns a copy of db that uses session s.
func (d *Database) With(s *mgo.Session) elements.Databaser {
	return &Database{d.Database.With(s)}
}
