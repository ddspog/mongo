// Interface for mgo.Collection to allow mocking.
package mongo

// Define interfaces that the types mgo.Collection, mgo.Database and
// mgo.Query implements. They are necessary to allow mocking of these
// types, needed for testing.

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Collectioner is the interface that tries to enumerate all methods
// that mgo.Collection have, with the difference of using the
// interfaces on this package, instead of mgo.Collection, mgo.Database
// and mgo.Query.
type Collectioner interface {
	Bulk() *mgo.Bulk
	Count() (n int, err error)
	Create(info *mgo.CollectionInfo) error
	DropCollection() error
	DropIndex(key ...string) error
	DropIndexName(name string) error
	EnsureIndex(index mgo.Index) error
	EnsureIndexKey(key ...string) error
	Find(query interface{}) Querier
	FindId(id interface{}) Querier
	Indexes() (indexes []mgo.Index, err error)
	Insert(docs ...interface{}) error
	NewIter(session *mgo.Session, firstBatch []bson.Raw, cursorId int64, err error) *mgo.Iter
	Pipe(pipeline interface{}) *mgo.Pipe
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error)
	RemoveId(id interface{}) error
	Repair() *mgo.Iter
	Update(selector interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpdateId(id interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	With(s *mgo.Session) Collectioner
}

// Databaser is the interface that tries to enumerate all methods
// that mgo.Database have, with the difference of using the
// interfaces on this package, instead of mgo.Collection, mgo.Database
// and mgo.Query.
type Databaser interface {
	AddUser(username, password string, readOnly bool) error
	C(name string) Collectioner
	CollectionNames() (names []string, err error)
	DropDatabase() error
	FindRef(ref *mgo.DBRef) Querier
	GridFS(prefix string) *mgo.GridFS
	Login(user, pass string) error
	Logout()
	RemoveUser(user string) error
	Run(cmd interface{}, result interface{}) error
	UpsertUser(user *mgo.User) error
	With(s *mgo.Session) Databaser
}

// Querier is the interface that tries to enumerate all methods
// that mgo.Query have, with the difference of using the interfaces
// on this package, instead of mgo.Collection, mgo.Database and
// mgo.Query.
type Querier interface {
	All(result interface{}) error
	Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error)
	Batch(n int) Querier
	Comment(comment string) Querier
	Count() (n int, err error)
	Distinct(key string, result interface{}) error
	Explain(result interface{}) error
	For(result interface{}, f func() error) error
	Hint(indexKey ...string) Querier
	Iter() *mgo.Iter
	Limit(n int) Querier
	LogReplay() Querier
	MapReduce(job *mgo.MapReduce, result interface{}) (info *mgo.MapReduceInfo, err error)
	One(result interface{}) (err error)
	Prefetch(p float64) Querier
	Select(selector interface{}) Querier
	SetMaxScan(n int) Querier
	SetMaxTime(d time.Duration) Querier
	Skip(n int) Querier
	Snapshot() Querier
	Sort(fields ...string) Querier
	Tail(timeout time.Duration) *mgo.Iter
}
