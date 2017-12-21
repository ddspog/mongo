package embedded

// Define embedded types that mimic mgo.Collection, mgo.Database and
// mgo.Query. They are embedded to use the interfaces defined on this
// package, allowing mocking and correct testing of them.

import (
	"time"

	"github.com/ddspog/mongo/elements"
	"gopkg.in/mgo.v2"
)

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
