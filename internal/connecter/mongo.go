package connecter

// Define Connect function, easy way to connect the MongoDB database,
// storing the Session, allowing easy access to Database object.

import (
	"fmt"
	"os"
	"sync"

	"github.com/globalsign/mgo"
)

// MongoConnecter represents a minimal set of functions to run a
// project using MongoDB. Connects and Disconnects, access Session
// information and returns DB on a parallel session clone.
type MongoConnecter interface {
	Connect() error
	Disconnect()
	ConsumeDatabaseOnSession(f func(*mgo.Database))
	Session() *mgo.Session
}

// Mongo is a MongoConnecter that functions with a real MongoDB
// connection.
type Mongo struct {
	once    sync.Once
	session *mgo.Session
	mongo   *mgo.DialInfo
}

// New returns a Mongo connecter for production purposes. It connects
// to a real MongoDB database to perform operations of search and
// manipulation of data.
func New() (m MongoConnecter) {
	m = &Mongo{}
	return
}

// Connect to MongoDB of server.
// It tries to connect with MONGODB_URL, but without defining this
// environment variable, tris to connect with default URL.
func (m *Mongo) Connect() (err error) {
	m.once.Do(func() {
		// Parse adequate MongoDB URI.
		u := m.mongoURI()

		// Capture Session and Mongo objects using URI.
		var d *mgo.DialInfo
		d, err = parseURL(u)
		if err != nil {
			err = fmt.Errorf("problem parsing Mongo URI uri=%[1]s err='%[2]v'", u, err.Error())
			return
		}
		var s *mgo.Session
		s, err = dial(u)
		if err != nil {
			err = fmt.Errorf("problem dialing Mongo URI uri=%[1]s err='%[2]v'", u, err.Error())
			return
		}

		// No errors showing, save objects.
		s.SetSafe(&mgo.Safe{})
		//log.Printf("debug: - Connected to MongoDB URI. uri=%s", u)

		m.session = s
		m.mongo = d
	})
	return
}

// Disconnect undo the connection made. Preparing package for a new
// connection.
func (m *Mongo) Disconnect() {
	m.once = *new(sync.Once)

	if m.Session() != nil {
		m.Session().Close()
	}

	m.session = nil
	m.mongo = nil
}

// ConsumeDatabaseOnSession clones a session and use it to creates a
// Databaser object to be consumed in f function. Closes session after
// consume of Databaser object.
func (m *Mongo) ConsumeDatabaseOnSession(f func(*mgo.Database)) {
	if s := m.Session(); s != nil {
		s := s.Clone()
		defer s.Close()

		f(s.DB(m.mongo.Database))
	} else {
		f(nil)
	}
}

// Session return connected mongo session.
func (m *Mongo) Session() (s *mgo.Session) {
	s = m.session
	return
}

const (
	// DBUrl is the default MongoDB url that will be used to
	// connect to the database.
	DBUrl = "mongodb://localhost:27017/test"
)

var (
	// parseURL returns the URL information collected.
	parseURL = mgo.ParseURL
	// dial returns mongo session after connecting.
	dial = mgo.Dial
)

// mongoURI load the selected MongoDB database url, or default.
func (m *Mongo) mongoURI() (u string) {
	u = os.Getenv("MONGODB_URL")
	if len(u) == 0 {
		u = DBUrl
	}
	return
}
