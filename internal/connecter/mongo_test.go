// +build !acceptance

package connecter

import (
	"errors"
	"testing"

	"github.com/ddspog/bdd"
	"github.com/globalsign/mgo"
)

// Feature Mongo consume database on session
// - As a developer,
// - I want to be able to use Mongo to get database object between sessions,
// - So that I can use it in parallel in isolate cases.
func Test_Mongo_consume_database_on_session(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, "a new test MongoConnecter m, and DBCollecter(el){db = el}", func(when bdd.When) {
		mc := New()

		var db *mgo.Database
		DBCollecter := func(el *mgo.Database) {
			db = el
		}

		mc.ConsumeDatabaseOnSession(DBCollecter)

		when("mc.ConsumeDatabaseOnSession(DBCollecter) is called before mc.Connect()", func(it bdd.It) {
			it("db should still be nil", func(assert bdd.Assert) {
				assert.Nil(db)
			})
		})

		err := mc.Connect()
		defer mc.Disconnect()

		when("err := mc.Connect() is called", func(it bdd.It) {
			it("shouldn't return any error", func(assert bdd.Assert) {
				assert.Nil(err)
			})
		})

		mc.ConsumeDatabaseOnSession(DBCollecter)

		when("mc.ConsumeDatabaseOnSession(DBCollecter) is called after mc.Connect()", func(it bdd.It) {
			it("db should partially equal to mc.Session().DB('test')", func(assert bdd.Assert) {
				assert.NotNil(mc.Session())
				if mc.Session() != nil {
					exp := mc.Session().DB("test")
					assert.Equal(exp.Name, db.Name)
				}
			})
		})
	})
}

// Feature Mongo returns URL information
// - As a developer,
// - I want to be able to use Mongo to acess URL informatino,
// - So that it can use this to connect to databases safely.
func Test_Mongo_returns_URL_information(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, "a new test MongoConnecter m, and parseURL returning error.New('any reason')", func(when bdd.When) {
		mc := New()

		parseURL = func(u string) (i *mgo.DialInfo, err error) {
			err = errors.New("any reason")
			return
		}
		defer resetUtils()

		err := mc.Connect()
		defer mc.Disconnect()

		when("err := mc.Connect() is called", func(it bdd.It) {
			it("should return an error", func(assert bdd.Assert) {
				assert.Error(err)
			})
		})
	})
}

// Feature Mongo dials connection
// - As a developer,
// - I want to be able to use Mongo to dial a connection,
// - So that it can access databases without problem.
func Test_Mongo_dials_connection(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, "a new test MongoConnecter m, and dial returning error.New('any reason')", func(when bdd.When) {
		mc := New()

		dial = func(u string) (s *mgo.Session, err error) {
			err = errors.New("any reason")
			return
		}
		defer resetUtils()

		err := mc.Connect()
		defer mc.Disconnect()

		when("err := mc.Connect() is called", func(it bdd.It) {
			it("should return an error", func(assert bdd.Assert) {
				assert.Error(err)
			})
		})
	})
}

// resetUtils reset the functions defined to the initial purpose.
// Making mocking really simple.
func resetUtils() {
	parseURL = mgo.ParseURL
	dial = mgo.Dial
}
