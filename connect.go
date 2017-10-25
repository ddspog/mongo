package mongo

// Define Connect function, easy way to connect the MongoDB database,
// storing the Session, allowing easy acess to Database object.

import (
	"log"
	"os"
	"sync"

	"gopkg.in/mgo.v2"
)

var (
	// Mutex that allows only one call to Connect.
	once sync.Once

	// session stores mongo session.
	session *mgo.Session

	// mongo stores the MongoDB connection string information.
	mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default MongoDB url that will be used to
	// connect to the database.
	MongoDBUrl = "mongodb://localhost:27017/test"
)

// Connect to MongoDB of server.
// It tries to connect with MONGODB_URL, but without defining this
// environment variable, tris to connect with default URL.
func Connect() {
	once.Do(func(){
		// Parse adequate MongoDB URI.
		u := mongoURI()

		// Capture Session and Mongo objects using URI.
		m, err := mgo.ParseURL(u)
		if err != nil {
			showError("Problem parsing Mongo URI. uri="+u, err.Error())
		}
		s, err := mgo.Dial(u)
		if err != nil {
			showError("Problem dialing Mongo URI. uri="+u, err.Error())
		}

		// No errors showing, save objects.
		s.SetSafe(&mgo.Safe{})
		log.Printf("debug: - Connected to MongoDB URI. uri=%s", u)

		session = s
		mongo = m
	})	
}

// Session return connected mongo session.
func Session() (s *mgo.Session) {
	s := session
	return
}

// Mongo return the MongoDB connection string information.
func Mongo() (m *mgo.DialInfo) {
	m := mongo
	return
}

// mongoURI load the selected MongoDB database url, or default.
func mongoURI() (u string) {
	u = os.Getenv("MONGODB_URL")
	if len(u) == 0 {
		u = MongoDBUrl
	}
	return
}

// showError log Panic message, and error captured on panic.
func showError(m, e string) {
	log.Print("panic: " + m)
	panic(e)
}
