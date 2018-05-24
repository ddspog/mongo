// +build !acceptance

package connecter

import (
	"testing"

	"github.com/ddspog/bdd"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Feature TstMongo accept only map fixtures
// - As a developer,
// - I want to be able to use Testmongo only with map fixtures,
// - So that I can access info on keys and data of each fixture.
func Test_TstMongo_accept_only_map_fixtures(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, "a new test MongoConnecter m with an array as fixtures", func(when bdd.When, args ...interface{}) {
		mc := NewTestable("", "internalTesting", []interface{}{newProduct(), newProduct()})

		err := mc.Connect()
		defer mc.Disconnect()

		when("err := mc.Connect() is called", func(it bdd.It) {
			it("should return an error", func(assert bdd.Assert) {
				assert.Equal(ErrInvalidFixtureMap, err)
			})
		})
	})
}

// Feature TstMongo accept only correct fixtures
// - As a developer,
// - I want to be able to use Testmongo only with correct naming fixtures,
// - So that I can differentiate collections withing Testmongo fixtures.
func Test_TstMongo_accept_only_correct_fixtures(t *testing.T) {
	given, like, s := bdd.Sentences().All()

	given(t, "a new test MongoConnecter m with %[1]v style of fixtures key naming", func(when bdd.When, args ...interface{}) {
		mc := NewTestable("", "internalTesting", args[2].(map[string]interface{}))
		err := mc.Connect()
		defer mc.Disconnect()

		when("err := mc.Connect() is called", func(it bdd.It) {

			if args[1].(bool) {
				it("shouldn't return any error", func(assert bdd.Assert) {
					assert.Nil(err)
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Equal(ErrInvalidFixtureKeys, err)
				})
			}
		})
	}, like(
		s("<db.col.id>", false, dbColIdFixtures),
		s("<col.id>", true, colIdFixtures),
		s("<id>", false, idFixtures),
	))
}

// Feature TstMongo sets a reset function
// - As a developer,
// - I want to be able to use Testmongo and set a reset function,
// - So that I can reset the temp database between tests.
func Test_TstMongo_sets_a_reset_function(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, "a new test MongoConnecter m with a function 'Reset'", func(when bdd.When) {
		var Reset func() error
		mc := NewTestable("", "internalTesting", colIdFixtures, &Reset)
		err := mc.Connect()
		defer mc.Disconnect()

		when("err := mc.Connect() is called", func(it bdd.It) {
			it("shouldn't return any error", func(assert bdd.Assert) {
				assert.Nil(err)
			})
		})

		err = Reset()

		when("err := Reset() is called", func(it bdd.It) {
			it("shouldn't return any error", func(assert bdd.Assert) {
				assert.Nil(err)
			})
		})
	})
}

// Feature TstMongo consume database on session
// - As a developer,
// - I want to be able to use Testmongo to get database object between sessions,
// - So that I can use it in parallel in isolate cases.
func Test_TstMongo_consume_database_on_session(t *testing.T) {
	given := bdd.Sentences().Given()

	given(t, "a new test MongoConnecter m, and DBCollecter(el){db = el}", func(when bdd.When) {
		mc := NewTestable("", "internalTesting", colIdFixtures)

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

		mc.Connect()
		defer mc.Disconnect()

		mc.ConsumeDatabaseOnSession(DBCollecter)

		when("mc.ConsumeDatabaseOnSession(DBCollecter) is called after mc.Connect()", func(it bdd.It) {
			it("db should partially equal to mc.Session().DB('test')", func(assert bdd.Assert) {
				exp := mc.Session().DB("test")
				assert.Equal(exp.Name, db.Name)
			})
		})
	})
}

var (
	// Various fixtures to help on tests.
	dbColIdFixtures = map[string]interface{}{
		"test.products.testid1": newProduct(),
		"test.products.testid2": newProduct(),
	}
	colIdFixtures = map[string]interface{}{
		"products.testid1": newProduct(),
		"products.testid2": newProduct(),
	}
	idFixtures = map[string]interface{}{
		"testid1": newProduct(),
		"testid2": newProduct(),
	}
)

// product it's an random type.
type product struct {
	IDV bson.ObjectId `bson:"_id"`
}

// newProduct creates a product with random id.
func newProduct() (p *product) {
	p = &product{
		IDV: bson.NewObjectId(),
	}
	return
}
