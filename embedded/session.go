package embedded

import (
	"github.com/ddspog/mongo/elements"
	"gopkg.in/mgo.v2"
)

// Session it's an embedded type of mgo.Session, made to use the
// interfaces it implements, Sessioner, on it's methods signatures.
type Session struct {
	*mgo.Session
}

// SetSafe changes the session safety mode.
//
// If the safe parameter is nil, the session is put in unsafe mode, and writes
// become fire-and-forget, without error checking.  The unsafe mode is faster
// since operations won't hold on waiting for a confirmation.
//
// If the safe parameter is not nil, any changing query (insert, update, ...)
// will be followed by a getLastError command with the specified parameters,
// to ensure the request was correctly processed.
//
// The default is &Safe{}, meaning check for errors and use the default
// behavior for all fields.
//
// The safe.W parameter determines how many servers should confirm a write
// before the operation is considered successful.  If set to 0 or 1, the
// command will return as soon as the primary is done with the request.
// If safe.WTimeout is greater than zero, it determines how many milliseconds
// to wait for the safe.W servers to respond before returning an error.
//
// Starting with MongoDB 2.0.0 the safe.WMode parameter can be used instead
// of W to request for richer semantics. If set to "majority" the server will
// wait for a majority of members from the replica set to respond before
// returning. Custom modes may also be defined within the server to create
// very detailed placement schemas. See the data awareness documentation in
// the links below for more details (note that MongoDB internally reuses the
// "w" field name for WMode).
//
// If safe.J is true, servers will block until write operations have been
// committed to the journal. Cannot be used in combination with FSync. Prior
// to MongoDB 2.6 this option was ignored if the server was running without
// journaling. Starting with MongoDB 2.6 write operations will fail with an
// exception if this option is used when the server is running without
// journaling.
//
// If safe.FSync is true and the server is running without journaling, blocks
// until the server has synced all data files to disk. If the server is running
// with journaling, this acts the same as the J option, blocking until write
// operations have been committed to the journal. Cannot be used in
// combination with J.
//
// Since MongoDB 2.0.0, the safe.J option can also be used instead of FSync
// to force the server to wait for a group commit in case journaling is
// enabled. The option has no effect if the server has journaling disabled.
//
// For example, the following statement will make the session check for
// errors, without imposing further constraints:
//
//     session.SetSafe(&mgo.Safe{})
//
// The following statement will force the server to wait for a majority of
// members of a replica set to return (MongoDB 2.0+ only):
//
//     session.SetSafe(&mgo.Safe{WMode: "majority"})
//
// The following statement, on the other hand, ensures that at least two
// servers have flushed the change to disk before confirming the success
// of operations:
//
//     session.EnsureSafe(&mgo.Safe{W: 2, FSync: true})
//
// The following statement, on the other hand, disables the verification
// of errors entirely:
//
//     session.SetSafe(nil)
//
// See also the EnsureSafe method.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/getLastError+Command
//     http://www.mongodb.org/display/DOCS/Verifying+Propagation+of+Writes+with+getLastError
//     http://www.mongodb.org/display/DOCS/Data+Center+Awareness
//
func (s *Session) SetSafe(sf *mgo.Safe) {
	s.Session.SetSafe(sf)
}

// Clone works just like Copy, but also reuses the same socket as the original
// session, in case it had already reserved one due to its consistency
// guarantees.  This behavior ensures that writes performed in the old session
// are necessarily observed when using the new session, as long as it was a
// strong or monotonic session.  That said, it also means that long operations
// may cause other goroutines using the original session to wait.
func (s *Session) Clone() (session elements.Sessioner) {
	session = &Session{s.Session.Clone()}
	return
}

// Close terminates the session.  It's a runtime error to use a session
// after it has been closed.
func (s *Session) Close() {
	s.Session.Close()
}

// DB returns a value representing the named database. If name is
// empty, the database name provided in the dialed URL is used instead.
// If that is also empty, "test" is used as a fallback in a way
// equivalent to the mongo shell.
//
// Creating this value is a very lightweight operation, and involves no
// network communication.
func (s *Session) DB(n string) (db elements.Databaser) {
	db = &Database{s.Session.DB(n)}
	return
}
