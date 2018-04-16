package mongo

import "github.com/ddspog/mongo/elements"

func newDBSocket() (db databaseSocketer) {
	db = &databaseSocket{
		db:   make(chan elements.Databaser),
		quit: make(chan bool),
	}
	return
}

type databaseSocketer interface {
	DB() elements.Databaser
	Close()
}

type databaseSocket struct {
	db   chan elements.Databaser
	quit chan bool
}

func (d *databaseSocket) DB() (db elements.Databaser) {
	go ConsumeDatabaseOnSession(func(db elements.Databaser) {
		d.db <- db
		<-d.quit
	})

	return <-d.db
}

func (d *databaseSocket) Close() {
	d.quit <- true
}
