package mongo

// Define Connect function, easy way to connect the MongoDB database,
// storing the Session, allowing easy access to Database object.

import (
	"fmt"
	"os"
	"sync"

	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/embedded"
	"github.com/globalsign/mgo"
)

var (
	// Mutex that allows only one call to Connect.
	once sync.Once
	// session represents the current session running MongoDB.
	session elements.Sessioner
	// mongo object holds info about the current mongo connection.
	mongo *elements.DialInfo
	// parseURL it's the function that will validate URL and return
	// connection information.
	parseURL = func(u string) (i *elements.DialInfo, err error) {
		info, err := mgo.ParseURL(u)
		i = elements.NewDialInfo(info)
		return
	}
	// dial it's the function that will connects program with a MongoDB
	// session.
	dial = func(u string) (s elements.Sessioner, err error) {
		session, err := mgo.Dial(u)
		s = &embedded.Session{
			Session: session,
		}
		return
	}
)

const (
	// DBUrl is the default MongoDB url that will be used to
	// connect to the database.
	DBUrl = "mongodb://localhost:27017/test"
)

// Connect to MongoDB of server.
// It tries to connect with MONGODB_URL, but without defining this
// environment variable, tris to connect with default URL.
func Connect() (err error) {
	once.Do(func() {
		// Parse adequate MongoDB URI.
		u := mongoURI()

		// Capture Session and Mongo objects using URI.
		var m *elements.DialInfo
		m, err = parseURL(u)
		if err != nil {
			err = fmt.Errorf("problem parsing Mongo URI uri=%[1]s err='%[2]v'", u, err.Error())
			return
		}
		var s elements.Sessioner
		s, err = dial(u)
		if err != nil {
			err = fmt.Errorf("problem dialing Mongo URI uri=%[1]s err='%[2]v'", u, err.Error())
			return
		}

		// No errors showing, save objects.
		s.SetSafe(&mgo.Safe{})
		//log.Printf("debug: - Connected to MongoDB URI. uri=%s", u)

		session = s
		mongo = m
	})
	return
}

func Disconnect() {
	once = *new(sync.Once)
	session = new(embedded.Session)
	mongo = new(elements.DialInfo)
}

// ConsumeDatabaseOnSession clones a session and use it to creates a
// Databaser object to be consumed in f function. Closes session after
// consume of Databaser object.
func ConsumeDatabaseOnSession(f func(elements.Databaser)) {
	if s := CurrentSession(); s != nil {
		s := s.Clone()
		defer s.Close()

		f(s.DB(Mongo().Database))
	} else {
		f(nil)
	}
}

// CurrentSession return connected mongo session.
func CurrentSession() (s elements.Sessioner) {
	s = session
	return
}

// Mongo return the MongoDB connection string information.
func Mongo() (m *elements.DialInfo) {
	m = mongo
	return
}

// mongoURI load the selected MongoDB database url, or default.
func mongoURI() (u string) {
	u = os.Getenv("MONGODB_URL")
	if len(u) == 0 {
		u = DBUrl
	}
	return
}
