package mongo

import (
	"testing"

	"github.com/ddspog/bdd"
	"github.com/ddspog/mongo/elements"
)

// Feature Connection with MongoDB
// - As a developer,
// - I want to be able to connect to a MongoDB database,
// - So that I can use any handlers to manipulate data.
func Test_Connection_with_MongoDB(t *testing.T) {
	create, _ := NewMockMGOSetup(t)
	defer create.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a database named '%[1]v' with a products collection with 10 elements", func(when bdd.When, args ...interface{}) {
		db := create.DatabaseMock("products", func(mcl *MockCollectioner) {
			mcl.ExpectCountReturn(10)
		})
		InitController(create.ControlMock(elements.NewDatabaseInfo(args[0].(string)), db))

		when("calling mongo.Connect()", func(it bdd.It) {
			it("should run with no problems", func(assert bdd.Assert) {
				assert.NoError(Connect())
			})
		})

		when("running mongo.ConsumeDatabaseOnSession() to link p handler on products collection", func(it bdd.It) {
			var n int
			var err error

			ConsumeDatabaseOnSession(func(db elements.Databaser) {
				p := newProductCount()
				n, err = p.Link(db).Count()
			})

			it("p.Count() should return no errors", func(assert bdd.Assert) {
				assert.NoError(err)
			})
			it("p.Count() should return 10", func(assert bdd.Assert) {
				assert.Equal(n, 10)
			})
		})
	}, like(
		s("test"), s("db01"), s("db02"),
	))
}
