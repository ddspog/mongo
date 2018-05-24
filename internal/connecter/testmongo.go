package connecter

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/dbtest"
)

var (
	// ErrInvalidFixtureKeys it's an error received when keys on
	// fixture doesn't follow correct structure <col.id>.
	ErrInvalidFixtureKeys = errors.New("invalid fixture key name, use <col.id>")
	// ErrInvalidFixtureMap it's an error received when fixture
	// received can't be converted to a map.
	ErrInvalidFixtureMap = errors.New("fixtures received ain't a map")
)

// TestMongo is a MongoConnecter with use on testing purposes. Using a
// dbtest.DBServer to simulate a Database on a temp directory.
type TestMongo struct {
	dir      string
	prefix   string
	path     string
	session  *mgo.Session
	server   *dbtest.DBServer
	fixtures interface{}
	resetFn  *func() error
}

// NewTestable returns a TestMongo as MongoConnecter, using a temp
// directory, fixtures to init the database and a optional reset
// function address, to be set when connecting the MongoConnecter.
//
// It's important to send fixtures as a map[string] of any type,
// otherwise, an error will be returned on Connect(). Also, every
// element must have a key structure as <col.id> for identification
// of collection to insert.
func NewTestable(d, p string, fixtures interface{}, reset ...*func() error) (m MongoConnecter) {
	tm := &TestMongo{
		dir:      d,
		prefix:   p,
		fixtures: fixtures,
	}

	if len(reset) == 1 {
		tm.resetFn = reset[0]
	}

	tm.server = new(dbtest.DBServer)
	m = tm
	return
}

// Connect to MongoDB of server. It tries to connect with MONGODB_URL,
// but without defining this environment variable, tris to connect
// with default URL. It also defines a reset function if received on
// constructor, to reset and initialize database with fixtures.
func (m *TestMongo) Connect() (err error) {
	if m.path, err = ioutil.TempDir(m.dir, m.prefix); err == nil {
		m.server.SetPath(m.path)
		m.session = m.server.Session()

		// No errors showing, save objects.
		m.session.SetSafe(&mgo.Safe{})

		if err = m.insertFixtures(); err == nil && m.resetFn != nil {
			*m.resetFn = func() (err error) {
				m.Session().DB("test").DropDatabase()
				err = m.insertFixtures()
				return
			}
		}
	}

	return
}

// insertFixtures init the database with some documents, defined on the
// constructor as fixtures.
func (m *TestMongo) insertFixtures() (err error) {
	if v := reflect.ValueOf(m.fixtures); v.Kind() == reflect.Map {
		for _, key := range v.MapKeys() {
			doc := v.MapIndex(key)

			names := strings.Split(key.String(), ".")
			if len(names) == 2 {
				err = m.Session().DB("test").C(names[0]).Insert(doc.Interface())
			} else {
				err = ErrInvalidFixtureKeys
			}

			if err != nil {
				return
			}
		}
	} else {
		err = ErrInvalidFixtureMap
	}
	return
}

// Disconnect undo the connection made. Preparing package for a new
// connection.
func (m *TestMongo) Disconnect() {
	m.session.Close()
	m.session = nil
	m.server.Stop()
	os.RemoveAll(m.path)
}

// ConsumeDatabaseOnSession clones a session and use it to creates a
// Databaser object to be consumed in f function. Closes session after
// consume of Databaser object. Returns nil if no session is available.
func (m *TestMongo) ConsumeDatabaseOnSession(f func(*mgo.Database)) {
	if s := m.Session(); s != nil {
		s := s.Clone()
		defer s.Close()

		f(s.DB("test"))
	} else {
		f(nil)
	}
}

// Session return connected mongo session.
func (m *TestMongo) Session() (s *mgo.Session) {
	s = m.session
	return
}
