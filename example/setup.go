package example

import "github.com/ddspog/mongo"
import "github.com/ddspog/mongo/elements"

func NewDBSocket() (db DatabaseSocketer) {
	db = &DatabaseSocket{
		db:   make(chan elements.Databaser),
		quit: make(chan bool),
	}
	return
}

type DatabaseSocketer interface {
	DB() elements.Databaser
	Close()
}

type DatabaseSocket struct {
	db   chan elements.Databaser
	quit chan bool
}

func (d *DatabaseSocket) DB() (db elements.Databaser) {
	go mongo.ConsumeDatabaseOnSession(func(db elements.Databaser) {
		d.db <- db
		<-d.quit
	})

	return <-d.db
}

func (d *DatabaseSocket) Close() {
	d.quit <- true
}
