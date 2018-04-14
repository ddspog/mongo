package example

import (
	"fmt"
	"os"
	"testing"

	"github.com/ddspog/mongo/model"

	"github.com/ddspog/mongo"

	"github.com/ddspog/mspec/bdd"

	"github.com/globalsign/mgo/bson"
)

const (
	// Const values to help tests legibility.
	testid01 = "000000007465737469643031"
	testid02 = "000000007465737469643032"
	testid03 = "000000007465737469643033"
	testid04 = "000000007465737469643034"
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

	given(t, fmt.Sprintf("a local database with a products collection that contains documents with ids: '%[1]s', '%[2]s', '%[3]s'", testid01, testid02, testid03), func(when bdd.When) {
		os.Setenv("MONGODB_URL", "mongodb://localhost:27017/test")
		_ = mongo.Connect()
		defer mongo.Disconnect()

		conn := NewDBSocket()
		defer conn.Close()
		db := conn.DB()

		p := NewProductHandle()
		p.Link(db)

		when("using p.Find() with '%[1]v' as document id'", func(it bdd.It, args ...interface{}) {
			p.DocumentV.IDV = bson.ObjectIdHex(args[0].(string))
			product, errFind := p.Find()

			it("should run without errors", func(assert bdd.Assert) {
				assert.NoError(errFind)
			})

			it("should return a product with id '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(args[0].(string), product.ID().Hex())
			})
		}, like(
			s(testid01), s(testid02), s(testid03),
		))

		when("using p.FindAll() with a empty Document", func(it bdd.It) {
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
					assert.Equal(args[1].(string), products[args[0].(int)-1].ID().Hex())
				}, like(
					s(1, testid01), s(2, testid02), s(3, testid03),
				))
			}
		})

		when("using p.FindAll() on a Document with id '%[1]v'", func(it bdd.It, args ...interface{}) {
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
					assert.Equal(args[0].(string), products[0].ID().Hex())
				})
			}
		}, like(
			s(testid01), s(testid02), s(testid03),
		))
	})
}

// Feature Manipulate data on MongoDB
// - As a developer,
// - I want to be able to connect and manipulate data from MongoDB,
// - So I can create an real application on this, using a DB.
func Test_Manipulate_data_on_MongoDB(t *testing.T) {
	given, _, _ := bdd.Sentences()

	given(t, fmt.Sprintf("a local database with a products collection that contains documents with ids: '%[1]s', '%[2]s', '%[3]s'", testid01, testid02, testid03), func(when bdd.When) {
		os.Setenv("MONGODB_URL", "mongodb://localhost:27017/test")
		_ = mongo.Connect()
		defer mongo.Disconnect()

		conn := NewDBSocket()
		defer conn.Close()
		db := conn.DB()

		p := NewProductHandle()
		p.Link(db)

		when("using p.RemoveAll()", func(it bdd.It) {
			removeInfo, errRemoveAll := p.RemoveAll()

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errRemoveAll)
			})

			it("should remove 3 documents", func(assert bdd.Assert) {
				assert.Equal(3, removeInfo.Removed)
			})

			it("should have p.Count() return 0", func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				n, errCount := p.Count()

				assert.NoError(errCount)
				assert.Equal(0, n)
			})

			it("should have p.FindAll() return nothing", func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				products, errFindAll := p.FindAll()

				assert.NoError(errFindAll)
				assert.Equal(0, len(products))
			})
		})

		var newId bson.ObjectId

		when(fmt.Sprintf("using p.Insert() to add documents with ids: '%[1]s', '%[2]s', '%[3]s', and a new one", testid01, testid02, testid03), func(it bdd.It) {
			p.DocumentV = NewProduct()
			p.DocumentV.IDV = bson.ObjectIdHex(testid01)
			errInsertDoc01 := p.Insert()

			p.DocumentV = NewProduct()
			p.DocumentV.IDV = bson.ObjectIdHex(testid02)
			errInsertDoc02 := p.Insert()

			p.DocumentV = NewProduct()
			p.DocumentV.IDV = bson.ObjectIdHex(testid03)
			errInsertDoc03 := p.Insert()

			p.DocumentV = NewProduct()
			errInsertDoc04 := p.Insert()
			newId = p.DocumentV.ID()

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errInsertDoc01)
				assert.NoError(errInsertDoc02)
				assert.NoError(errInsertDoc03)
				assert.NoError(errInsertDoc04)
			})

			it("should have p.Count() return 4", func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				n, errCount := p.Count()

				assert.NoError(errCount)
				assert.Equal(4, n)
			})

			it("should have p.Find() return all documents", func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				p.DocumentV.IDV = bson.ObjectIdHex(testid01)
				product01, errFindDoc01 := p.Find()

				assert.NoError(errFindDoc01)
				assert.Equal(testid01, product01.ID().Hex())

				p.DocumentV = NewProduct()
				p.DocumentV.IDV = bson.ObjectIdHex(testid02)
				product02, errFindDoc02 := p.Find()

				assert.NoError(errFindDoc02)
				assert.Equal(testid02, product02.ID().Hex())

				p.DocumentV = NewProduct()
				p.DocumentV.IDV = bson.ObjectIdHex(testid03)
				product03, errFindDoc03 := p.Find()

				assert.NoError(errFindDoc03)
				assert.Equal(testid03, product03.ID().Hex())

				p.DocumentV = NewProduct()
				p.DocumentV.IDV = newId
				newProduct, errFindNewDoc := p.Find()

				assert.NoError(errFindNewDoc)
				assert.Equal(newId.Hex(), newProduct.ID().Hex())
			})
		})

		now := model.NowInMilli()

		when(fmt.Sprintf("using p.Update() to change created_on doc with id '%[1]s' to '%[2]v'", newId.Hex(), now), func(it bdd.It) {
			p.DocumentV = NewProduct()
			p.DocumentV.CreatedOnV = now
			errUpdate := p.Update(newId)

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errUpdate)
			})

			it(fmt.Sprintf("should have p.Find() with id '%[1]s' return with new value", newId.Hex()), func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				p.DocumentV.IDV = newId
				product, errFind := p.Find()

				assert.NoError(errFind)
				assert.NotNil(product)
				assert.Equal(now, product.CreatedOn())
				assert.NotEqual(0, product.UpdatedOn())
			})

			it("should have p.Count() return 4", func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				n, errCount := p.Count()

				assert.NoError(errCount)
				assert.Equal(4, n)
			})
		})

		when(fmt.Sprintf("using p.Remove() to remove doc with id '%[1]s'", newId.Hex()), func(it bdd.It) {
			p.DocumentV = NewProduct()
			p.DocumentV.IDV = newId
			errRemove := p.Remove()

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errRemove)
			})

			it(fmt.Sprintf("should have p.Find() with id '%[1]s' return nothing", newId.Hex()), func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				p.DocumentV.IDV = newId
				product, errFind := p.Find()

				assert.Error(errFind)
				assert.Contains(errFind.Error(), "not found")
				assert.Empty(product.ID().Hex())
			})

			it("should have p.Count() return 3", func(assert bdd.Assert) {
				p.DocumentV = NewProduct()
				n, errCount := p.Count()

				assert.NoError(errCount)
				assert.Equal(3, n)
			})
		})
	})
}
