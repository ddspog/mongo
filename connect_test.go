package mongo

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/ddspog/bdd"
	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/mocks"

	"github.com/comail/colog"
)

// Feature Connection with MongoDB
// - As a developer,
// - I want to be able to connect to a MongoDB database,
// - So that I can use any handlers to manipulate data.
func Test_Connection_with_MongoDB(t *testing.T) {
	colog.Register()
	colog.SetMinLevel(colog.LError)

	createMGO, _ := mocks.NewMockMGOSetup(t)
	createMongo, _ := NewMockMongoSetup(t)
	defer finish(createMGO, createMongo)

	given, like, s := bdd.Sentences()

	given(t, "a database named '%[1]v' with a products collection with 10 elements", func(when bdd.When, args ...interface{}) {
		db := createMGO.DatabaseMock(args[0].(string), func(mcl *mocks.MockCollectioner) {
			mcl.ExpectCountReturn(10)
		})

		createMongo.ParseURL().Returns(elements.NewDatabaseInfo(args[0].(string)), nil)
		createMongo.Dial().Returns(createMGO.SessionMock(args[0].(string), db), nil)

		when("calling mongo.Connect()", func(it bdd.It) {
			it("should run with no problems", func(assert bdd.Assert) {
				assert.NoError(Connect())
			})
		})

		when("running mongo.ConsumeDatabaseOnSession() to link p handler on '%[1]v' collection", func(it bdd.It) {
			var n int
			var err error

			ConsumeDatabaseOnSession(func(db elements.Databaser) {
				p := newProductCount(args[0].(string))
				n, err = p.Link(db).Count()
			})

			it("p.Count() should return no errors", func(assert bdd.Assert) {
				assert.NoError(err)
			})
			it("p.Count() should return 10", func(assert bdd.Assert) {
				assert.Equal(n, 10)
			})
		})

		once = *new(sync.Once)
	}, like(
		s("test"), s("db01"), s("db02"),
	))
}

// Feature Connect only with valid URLs.
// - As a developer,
// - I want Connect to fail with invalid URLs,
// - So that I can connect on MongoDB without errors.
func Test_Connect_only_with_valid_URLs(t *testing.T) {
	colog.Register()
	colog.SetMinLevel(colog.LError)

	createMGO, _ := mocks.NewMockMGOSetup(t)
	createMongo, _ := NewMockMongoSetup(t)
	defer finish(createMGO, createMongo)

	given, _, _ := bdd.Sentences()

	given(t, "a valid url u as env MONGODB_URL with no problems", func(when bdd.When) {
		os.Setenv("MONGODB_URL", "validURL")

		db := createMGO.DatabaseMock("randomDB", func(mcl *mocks.MockCollectioner) {})

		createMongo.ParseURL().Returns(elements.NewDatabaseInfo("randomDB"), nil)
		createMongo.Dial().Returns(createMGO.SessionMock("randomDB", db), nil)

		when("calling mongo.Connect()", func(it bdd.It) {
			it("shouldn't return an error", func(assert bdd.Assert) {
				assert.NoError(Connect())
			})
		})

		once = *new(sync.Once)
	})

	given(t, "a valid url u as env MONGODB_URL with parsing problems", func(when bdd.When) {
		u := "badParsedURL"
		os.Setenv("MONGODB_URL", u)

		createMongo.ParseURL().Returns(nil, fmt.Errorf("anyReason"))

		when("calling mongo.Connect()", func(it bdd.It) {
			it("should return a parsing error", func(assert bdd.Assert) {
				err := Connect()

				assert.Error(err)
				assert.Contains(err.Error(), "Problem parsing Mongo URI. uri="+u)
			})
		})

		once = *new(sync.Once)
	})

	given(t, "a valid url u as env MONGODB_URL with dialing problems", func(when bdd.When) {
		u := "badDialingURL"
		os.Setenv("MONGODB_URL", u)

		createMongo.ParseURL().Returns(elements.NewDatabaseInfo("randomDB"), nil)
		createMongo.Dial().Returns(nil, fmt.Errorf("anyReason"))

		when("calling mongo.Connect()", func(it bdd.It) {
			it("should return a dialing error", func(assert bdd.Assert) {
				err := Connect()

				assert.Error(err)
				assert.Contains(err.Error(), "Problem dialing Mongo URI. uri="+u)
			})
		})

		once = *new(sync.Once)
	})
}

// Feature MockMongoSetup works only on Tests.
// - As a developer,
// - I want that MockMongoSetup returns an error when receiving a nil test element,
// - So that I could restrain the use of this Setup only to tests.
func Test_MockModelSetup_works_only_on_Tests(t *testing.T) {
	given, _, _ := bdd.Sentences()

	given(t, "the start of the test", func(when bdd.When) {
		when("calling NewMockMongoSetup(nil)", func(it bdd.It) {
			_, err := NewMockMongoSetup(nil)
			it("should return an error", func(assert bdd.Assert) {
				assert.Error(err)
			})
		})
	})
}
