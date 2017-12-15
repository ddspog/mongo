package mongo

// Define embedded types that mimic mgo.Collection, mgo.Database and
// mgo.Query. They are embedded to use the interfaces defined on this
// package, allowing mocking and correct testing of them.

import (
	"time"

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

// Query it's an embedded type of mgo.Query, made to use the
// interfaces it implements, Querier, on it's methods signatures.
type Query struct {
	*mgo.Query
}

// Batch default size is defined by the database itself.  As of this
// writing, MongoDB will use an initial size of min(100 docs, 4MB) on the
// first batch, and 4MB on remaining ones.
func (q *Query) Batch(n int) elements.Querier {
	return &Query{q.Query.Batch(n)}
}

// Comment adds a comment to the query to identify it in the database profiler output.
//
// Relevant documentation:
//
//     http://docs.mongodb.org/manual/reference/operator/meta/comment
//     http://docs.mongodb.org/manual/reference/command/profile
//     http://docs.mongodb.org/manual/administration/analyzing-mongodb-performance/#database-profiling
//
func (q *Query) Comment(c string) elements.Querier {
	return &Query{q.Query.Comment(c)}
}

// Hint will include an explicit "hint" in the query to force the server
// to use a specified index, potentially improving performance in some
// situations.  The provided parameters are the fields that compose the
// key of the index to be used.  For details on how the indexKey may be
// built, see the EnsureIndex method.
//
// For example:
//
//     query := collection.Find(bson.M{"firstname": "Joe", "lastname": "Winter"})
//     query.Hint("lastname", "firstname")
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Optimization
//     http://www.mongodb.org/display/DOCS/Query+Optimizer
//
func (q *Query) Hint(idk ...string) elements.Querier {
	return &Query{q.Query.Hint(idk...)}
}

// Limit restricts the maximum number of documents retrieved to n, and also
// changes the batch size to the same value.  Once n documents have been
// returned by Next, the following call will return ErrNotFound.
func (q *Query) Limit(n int) elements.Querier {
	return &Query{q.Query.Limit(n)}
}

// LogReplay enables an option that optimizes queries that are typically
// made on the MongoDB oplog for replaying it. This is an internal
// implementation aspect and most likely uninteresting for other uses.
// It has seen at least one use case, though, so it's exposed via the API.
func (q *Query) LogReplay() elements.Querier {
	return &Query{q.Query.LogReplay()}
}

// Prefetch sets the point at which the next batch of results will be requested.
// When there are p*batch_size remaining documents cached in an Iter, the next
// batch will be requested in background. For instance, when using this:
//
//     query.Batch(200).Prefetch(0.25)
//
// and there are only 50 documents cached in the Iter to be processed, the
// next batch of 200 will be requested. It's possible to change this setting on
// a per-session basis as well, using the SetPrefetch method of Session.
//
// The default prefetch value is 0.25.
func (q *Query) Prefetch(p float64) elements.Querier {
	return &Query{q.Query.Prefetch(p)}
}

// Select enables selecting which fields should be retrieved for the results
// found. For example, the following query would only retrieve the name field:
//
//     err := collection.Find(nil).Select(bson.M{"name": 1}).One(&result)
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Retrieving+a+Subset+of+Fields
//
func (q *Query) Select(sel interface{}) elements.Querier {
	return &Query{q.Query.Select(sel)}
}

// SetMaxScan constrains the query to stop after scanning the specified
// number of documents.
//
// This modifier is generally used to prevent potentially long running
// queries from disrupting performance by scanning through too much data.
func (q *Query) SetMaxScan(n int) elements.Querier {
	return &Query{q.Query.SetMaxScan(n)}
}

// SetMaxTime constrains the query to stop after running for the specified time.
//
// When the time limit is reached MongoDB automatically cancels the query.
// This can be used to efficiently prevent and identify unexpectedly slow queries.
//
// A few important notes about the mechanism enforcing this limit:
//
//  - Requests can block behind locking operations on the server, and that blocking
//    time is not accounted for. In other words, the timer starts ticking only after
//    the actual start of the query when it initially acquires the appropriate lock;
//
//  - Operations are interrupted only at interrupt points where an operation can be
//    safely aborted – the total execution time may exceed the specified value;
//
//  - The limit can be applied to both CRUD operations and commands, but not all
//    commands are interruptible;
//
//  - While iterating over results, computing follow up batches is included in the
//    total time and the iteration continues until the allotted time is over, but
//    network roundtrips are not taken into account for the limit.
//
//  - This limit does not override the inactive cursor timeout for idle cursors
//    (default is 10 min).
//
// This mechanism was introduced in MongoDB 2.6.
//
// Relevant documentation:
//
//   http://blog.mongodb.org/post/83621787773/maxtimems-and-query-optimizer-introspection-in
//
func (q *Query) SetMaxTime(d time.Duration) elements.Querier {
	return &Query{q.Query.SetMaxTime(d)}
}

// Skip skips over the n initial documents from the query results.  Note that
// this only makes sense with capped collections where documents are naturally
// ordered by insertion time, or with sorted results.
func (q *Query) Skip(n int) elements.Querier {
	return &Query{q.Query.Skip(n)}
}

// Snapshot will force the performed query to make use of an available
// index on the _id field to prevent the same document from being returned
// more than once in a single iteration. This might happen without this
// setting in situations when the document changes in size and thus has to
// be moved while the iteration is running.
//
// Because snapshot mode traverses the _id index, it may not be used with
// sorting or explicit hints. It also cannot use any other index for the
// query.
//
// Even with snapshot mode, items inserted or deleted during the query may
// or may not be returned; that is, this mode is not a true point-in-time
// snapshot.
//
// The same effect of Snapshot may be obtained by using any unique index on
// field(s) that will not be modified (best to use Hint explicitly too).
// A non-unique index (such as creation time) may be made unique by
// appending _id to the index when creating it.
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/How+to+do+Snapshotted+Queries+in+the+Mongo+Database
//
func (q *Query) Snapshot() elements.Querier {
	return &Query{q.Query.Snapshot()}
}

// Sort asks the database to order returned documents according to the
// provided field names. A field name may be prefixed by - (minus) for
// it to be sorted in reverse order.
//
// For example:
//
//     query1 := collection.Find(nil).Sort("firstname", "lastname")
//     query2 := collection.Find(nil).Sort("-age")
//     query3 := collection.Find(nil).Sort("$natural")
//     query4 := collection.Find(nil).Select(bson.M{"score": bson.M{"$meta": "textScore"}}).Sort("$textScore:score")
//
// Relevant documentation:
//
//     http://www.mongodb.org/display/DOCS/Sorting+and+Natural+Order
//
func (q *Query) Sort(fields ...string) elements.Querier {
	return &Query{q.Query.Sort(fields...)}
}
