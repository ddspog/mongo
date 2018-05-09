package mongo

import (
	"github.com/globalsign/mgo"
)

// DatabaseSocket it's a socket connection with a specified MongoDB
// database, that can be closed after using it. It's used to make calls
// to the mongo collections parallel and independent.
type DatabaseSocket struct {
	db   chan *mgo.Database
	quit chan bool
}

// NewSocket creates a new DatabaseSocket, initializing channel values,
// supporting the DB calls.
func NewSocket() (db *DatabaseSocket) {
	db = &DatabaseSocket{
		db:   make(chan *mgo.Database),
		quit: make(chan bool),
	}
	return
}

// DB returns the database object returned by a cloned session of Mongo
// connection. Requires closing after operation is done, to avoid
// memory leak.
func (d *DatabaseSocket) DB() (db *mgo.Database) {
	go ConsumeDatabaseOnSession(func(db *mgo.Database) {
		d.db <- db
		<-d.quit
	})

	return <-d.db
}

// Close the socket open when DB is called.
func (d *DatabaseSocket) Close() {
	d.quit <- true
}
