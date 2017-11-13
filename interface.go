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
	Count() (int, error)
	Create(*mgo.CollectionInfo) error
	DropCollection() error
	DropIndex(...string) error
	DropIndexName(string) error
	EnsureIndex(mgo.Index) error
	EnsureIndexKey(...string) error
	Find(interface{}) Querier
	FindID(interface{}) Querier
	Indexes() ([]mgo.Index, error)
	Insert(docs ...interface{}) error
	NewIter(*mgo.Session, []bson.Raw, int64, error) *mgo.Iter
	Pipe(interface{}) *mgo.Pipe
	Remove(interface{}) error
	RemoveAll(interface{}) (*ChangeInfo, error)
	RemoveID(interface{}) error
	Repair() *mgo.Iter
	Update(interface{}, interface{}) error
	UpdateAll(interface{}, interface{}) (*ChangeInfo, error)
	UpdateID(interface{}, interface{}) error
	Upsert(interface{}, interface{}) (*ChangeInfo, error)
	UpsertID(interface{}, interface{}) (*ChangeInfo, error)
	With(*mgo.Session) Collectioner
}

// Databaser is the interface that tries to enumerate all methods
// that mgo.Database have, with the difference of using the
// interfaces on this package, instead of mgo.Collection, mgo.Database
// and mgo.Query.
type Databaser interface {
	AddUser(string, string, bool) error
	C(string) Collectioner
	CollectionNames() ([]string, error)
	DropDatabase() error
	FindRef(*mgo.DBRef) Querier
	GridFS(string) *mgo.GridFS
	Login(string, string) error
	Logout()
	RemoveUser(string) error
	Run(interface{}, interface{}) error
	UpsertUser(*mgo.User) error
	With(*mgo.Session) Databaser
}

// Querier is the interface that tries to enumerate all methods
// that mgo.Query have, with the difference of using the interfaces
// on this package, instead of mgo.Collection, mgo.Database and
// mgo.Query.
type Querier interface {
	All(interface{}) error
	Apply(mgo.Change, interface{}) (*mgo.ChangeInfo, error)
	Batch(int) Querier
	Comment(string) Querier
	Count() (int, error)
	Distinct(string, interface{}) error
	Explain(interface{}) error
	For(interface{}, func() error) error
	Hint(...string) Querier
	Iter() *mgo.Iter
	Limit(int) Querier
	LogReplay() Querier
	MapReduce(*mgo.MapReduce, interface{}) (*mgo.MapReduceInfo, error)
	One(interface{}) error
	Prefetch(float64) Querier
	Select(interface{}) Querier
	SetMaxScan(int) Querier
	SetMaxTime(time.Duration) Querier
	Skip(int) Querier
	Snapshot() Querier
	Sort(...string) Querier
	Tail(time.Duration) *mgo.Iter
}
