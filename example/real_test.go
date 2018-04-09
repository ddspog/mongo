package example

import (
	"os"
	"testing"

	"github.com/ddspog/mongo"

	"github.com/ddspog/mspec/bdd"
)

// Feature Real connection to MongoDB
// - As a developer,
// - I want to be able to connect to MongoDB with this package,
// - So that I could use the Handle methods to to various operations on a real DB.
func Test_Real_connection_to_MongoDB(t *testing.T) {
	given, _, _ := bdd.Sentences()

	given(t, "a local database", func(when bdd.When) {
		when("connecting with its url", func(it bdd.It) {
			os.Setenv("MONGODB_URL", "mongodb://localhost:27017/test")
			err := mongo.Connect()
			defer mongo.Disconnect()

			it("should connect with no problems", func(assert bdd.Assert) {
				assert.NoError(err)
			})

			conn := NewDBSocket()
			defer conn.Close()

			db := conn.DB()

			it("should open a socket containing valid DB", func(assert bdd.Assert) {
				assert.NotNil(db)
			})

			p, err := NewProductHandle().Link(db)

			it("should link correctly with products collection", func(assert bdd.Assert) {
				assert.NoError(err)
			})

			n, err := p.Count()

			it("should enable counting the documents on products collection", func(assert bdd.Assert) {
				assert.NoError(err)
				assert.Equal(1, n)
			})
		})
	})
}
