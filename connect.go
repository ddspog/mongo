package mongo

//noinspection GoInvalidPackageImport
import (
	"github.com/ddspog/mongo/internal/connecter"
	"github.com/globalsign/mgo"
)

// Connecter represents method of connection to Mongo, using real
// database of a temporary database.
type Connecter = connecter.MongoConnecter

var (
	// NewConnecter returns a new real database connecter.
	NewConnecter = connecter.New
	// NewTestableConnecter returns a temporary database connecter.
	NewTestableConnecter = connecter.NewTestable
	// conn connects with MongoDB
	conn = NewConnecter()
)

// InitConnecter with the real database connecter, or with testable
// version if given as parameter.
func InitConnecter(c ...Connecter) {
	conn = connecter.New()

	if len(c) == 1 && c[0] != nil {
		conn = c[0]
	}
}

// Connect to MongoDB of server.
// It tries to connect with MONGODB_URL, but without defining this
// environment variable, tries to connect with default URL.
func Connect() (err error) {
	err = conn.Connect()
	return
}

// Disconnect undo the connection made. Preparing package for a new
// connection.
func Disconnect() {
	conn.Disconnect()
}

// ConsumeDatabaseOnSession clones a session and use it to creates a
// Databaser object to be consumed in f function. Closes session after
// consume of Databaser object.
func ConsumeDatabaseOnSession(f func(*mgo.Database)) {
	conn.ConsumeDatabaseOnSession(f)
}

// Session return connected mongo session.
func Session() (s *mgo.Session) {
	s = conn.Session()
	return
}
