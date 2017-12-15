package mongo

// Define Connect function, easy way to connect the MongoDB database,
// storing the Session, allowing easy access to Database object.

import (
	"github.com/ddspog/mongo/elements"
)

var (
	// control stores package controller.
	control elements.Controller
)

// InitController initializes the package controller with a object of
// caller choice. This enables more power in mocking this package
// functions.
func InitController(c elements.Controller) {
	control = c
}

// Connect to MongoDB of server.
// It tries to connect with MONGODB_URL, but without defining this
// environment variable, tris to connect with default URL.
func Connect() (err error) {
	ensuresControlIsDefined()
	err = control.Connect()
	return
}

// CurrentSession return connected mongo session.
func CurrentSession() (s elements.Sessioner) {
	ensuresControlIsDefined()
	return control.Session()
}

// ConsumeDatabaseOnSession clones a session and use it to creates a
// Databaser object to be consumed in f function. Closes session after
// consume of Databaser object.
func ConsumeDatabaseOnSession(f func(elements.Databaser)) {
	ensuresControlIsDefined()
	control.ConsumeDatabaseOnSession(f)
}

// Mongo return the MongoDB connection string information.
func Mongo() (m *elements.DialInfo) {
	ensuresControlIsDefined()
	return control.Mongo()
}

// ensuresControlIsDefined asserts that control is defined when being
// called. Even if it's still empty.
func ensuresControlIsDefined() {
	if control == nil {
		InitController(&Control{})
	}
}
