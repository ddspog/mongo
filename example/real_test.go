package example

import (
	"os"
	"testing"

	"github.com/ddspog/mongo"

	"github.com/ddspog/mspec/bdd"

	"github.com/globalsign/mgo/bson"
)

const (
	// Const values to help tests legibility.
	testid01 = "000000007465737469643031"
	testid02 = "000000007465737469643032"
	testid03 = "000000007465737469643033"
)

// Feature Real connection to MongoDB
// - As a developer,
// - I want to be able to connect to MongoDB with this package,
// - So that I could use the Handle methods to to various operations on a real DB.
func Test_Real_connection_to_MongoDB(t *testing.T) {
	given, _, _ := bdd.Sentences()

	given(t, "a local database with a products collection with 3 documents", func(when bdd.When) {
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

			it("should count 3 documents on products collection", func(assert bdd.Assert) {
				assert.NoError(err)
				assert.Equal(3, n)
			})
		})
	})
}

// Feature Read data on MongoDB
// - As a developer,
// - I want to be able to connect and retrieve data from MongoDB,
// - So I can use these functionalities on real applications.
func Test_Read_data_on_MongoDB(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a local database with a products collection that contains documents with ids: 'testid01', 'testid02', 'testid03'", func(when bdd.When) {
		os.Setenv("MONGODB_URL", "mongodb://localhost:27017/test")
		_ = mongo.Connect()
		defer mongo.Disconnect()

		conn := NewDBSocket()
		defer conn.Close()
		db := conn.DB()

		p := NewProductHandle()
		p.Link(db)

		when("using ph.Find() with '%[1]v' as document id'", func(it bdd.It, args ...interface{}) {
			p.DocumentV.IDV = bson.ObjectIdHex(args[0].(string))
			product, errFind := p.Find()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFind)
			})

			it("should return a product with id '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(args[0].(string), product.IDV.Hex())
			})
		}, like(
			s(testid01), s(testid02), s(testid03),
		))

		when("using ph.FindAll() with a empty Document", func(it bdd.It) {
			p.DocumentV = NewProduct()
			products, errFindAll := p.FindAll()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFindAll)
			})

			it("should return 3 documents", func(assert bdd.Assert) {
				assert.Equal(3, len(products))
			})

			if len(products) == 3 {
				it("should return the %[1]vth document with id '%[2]v'", func(assert bdd.Assert, args ...interface{}) {
					assert.Equal(args[1].(string), products[args[0].(int)-1].IDV.Hex())
				}, like(
					s(1, testid01), s(2, testid02), s(3, testid03),
				))
			}
		})

		when("using ph.FindAll() on a Document with id '%[1]v'", func(it bdd.It, args ...interface{}) {
			p.DocumentV.IDV = bson.ObjectIdHex(args[0].(string))
			products, errFindAll := p.FindAll()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFindAll)
			})

			it("should return 1 documents", func(assert bdd.Assert) {
				assert.Equal(1, len(products))
			})

			if len(products) == 1 {
				it("should return the 1th document with id '%[1]v'", func(assert bdd.Assert) {
					assert.Equal(args[0].(string), products[0].IDV.Hex())
				})
			}
		}, like(
			s(testid01), s(testid02), s(testid03),
		))
	})
}
