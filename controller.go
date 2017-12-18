package mongo

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ddspog/mongo/elements"
	"gopkg.in/mgo.v2"
)

var (
	// Mutex that allows only one call to Connect.
	once sync.Once
)

const (
	// DBUrl is the default MongoDB url that will be used to
	// connect to the database.
	DBUrl = "mongodb://localhost:27017/test"
)

// Control implements functions necessary to connect to MongoDB to
// handlers defined on system.
type Control struct {
	session elements.Sessioner
	mongo   *elements.DialInfo
}

// Connect to MongoDB of server. It tries to connect with MONGODB_URL,
// but without defining this environment variable, tries to connect
// with default URL.
func (c *Control) Connect() (err error) {
	once.Do(func() {
		// Parse adequate MongoDB URI.
		u := mongoURI()

		// Capture Session and Mongo objects using URI.
		var m *mgo.DialInfo
		m, err = mgo.ParseURL(u)
		if err != nil {
			err = fmt.Errorf("Problem parsing Mongo URI. uri="+u, err.Error())
			return
		}
		var s *mgo.Session
		s, err = mgo.Dial(u)
		if err != nil {
			err = fmt.Errorf("Problem dialing Mongo URI. uri="+u, err.Error())
			return
		}

		// No errors showing, save objects.
		s.SetSafe(&mgo.Safe{})
		log.Printf("debug: - Connected to MongoDB URI. uri=%s", u)

		c.session = &Session{s}
		c.mongo = elements.NewDialInfo(m)
	})
	return
}

// Session return connected mongo session.
func (c *Control) Session() (s elements.Sessioner) {
	s = c.session
	return
}

// Mongo return the MongoDB connection string information.
func (c *Control) Mongo() (m *elements.DialInfo) {
	m = c.mongo
	return
}

// ConsumeDatabaseOnSession clones a session and use it to creates a
// Databaser object to be consumed in f function. Closes session after
// consume of Databaser object.
func (c *Control) ConsumeDatabaseOnSession(f func(elements.Databaser)) {
	s := c.Session().Clone()
	defer s.Close()

	f(s.DB(c.Mongo().Database))
}

// mongoURI load the selected MongoDB database url, or default.
func mongoURI() (u string) {
	u = os.Getenv("MONGODB_URL")
	if len(u) == 0 {
		u = DBUrl
	}
	return
}
