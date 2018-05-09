// +build !acceptance

package mongo

import (
	"testing"
	"time"

	"github.com/ddspog/mspec/bdd"
)

// Feature Enable embedding with Document
// - As a developer,
// - I want to be able to embedded Document in other defined types,
// - So that I could use the Document methods to abstract data on it.
func Test_Enable_embedding_with_Document(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a new embedded Product p with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProduct()
		p.IDV = ObjectIdHex(args[0].(string))

		when("casting to Documenter interface d", func(it bdd.It) {
			var d Documenter = p
			it("d.ID().Hex() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(d.ID().Hex(), args[0].(string))
			})
		})
	}, like(
		s(id1), s(id2), s(id3),
	))
}

// Feature Create Document with functional Getters
// - As a developer,
// - I want to be able to create a Document and access data with its getters,
// - So that I could use these getters to manipulate and read data.
func Test_Create_Document_with_functional_Getters(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a Product p with ID '%[1]v', CreatedOn = %[2]v, UpdatedOn = %[3]v", func(when bdd.When, args ...interface{}) {
		p := newProduct()
		p.IDV = ObjectIdHex(args[0].(string))
		p.CreatedOnV = args[1].(int64)
		p.UpdatedOnV = args[2].(int64)

		when("p.ID().Hex() is called", func(it bdd.It) {
			it("should return '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(p.ID().Hex(), args[0].(string))
			})
		})
		when("p.CreatedOn() is called", func(it bdd.It) {
			it("should return %[2]v", func(assert bdd.Assert) {
				assert.Equal(p.CreatedOn(), args[1].(int64))
			})
		})
		when("p.UpdatedOn() is called", func(it bdd.It) {
			it("should return %[3]v", func(assert bdd.Assert) {
				assert.Equal(p.UpdatedOn(), args[2].(int64))
			})
		})
	}, like(
		s(id1, int64(123), int64(321)),
		s(id2, int64(10), int64(1)),
		s(id3, int64(20), int64(2)),
	))
}

// Feature Create Document with functional Setters
// - As a developer,
// - I want to be able to create a Document and modify data with its setters,
// - So that I could use these setters to manipulate data.
func Test_Create_Document_with_functional_Setters(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a Product p with ID '%[1]v', CreatedOn = %[2]v, UpdatedOn = %[3]v", func(when bdd.When, args ...interface{}) {
		p := newProduct()

		when("p.IDV = ObjectIdHex(%[1]v)", func(it bdd.It) {
			p.IDV = ObjectIdHex(args[0].(string))
			it("p.ID().Hex() should return '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(p.ID().Hex(), args[0].(string))
			})
		})
		when("p.CreatedOnV = %[2]v is called", func(it bdd.It) {
			p.CreatedOnV = args[1].(int64)
			it("p.CreatedOn() should return %[2]v", func(assert bdd.Assert) {
				assert.Equal(p.CreatedOn(), args[1].(int64))
			})
		})
		when("p.UpdatedOnV = %[3]v is called", func(it bdd.It) {
			p.UpdatedOnV = args[2].(int64)
			it("p.UpdatedOn() should return %[3]v", func(assert bdd.Assert) {
				assert.Equal(p.UpdatedOn(), args[2].(int64))
			})
		})
	}, like(
		s(id1, int64(123), int64(321)),
		s(id2, int64(10), int64(1)),
		s(id3, int64(20), int64(2)),
	))
}

// Feature Calculate Document values
// - As a developer,
// - I want to be able to call calculation methods to set some values with current time,
// - So that I could use these values later for data analysis.
func Test_Calculate_Document_values(t *testing.T) {

	given, like, s := bdd.Sentences()

	given(t, "a empty Product p at current time %[1]v", func(when bdd.When, args ...interface{}) {
		now = func() (t time.Time) {
			t = args[0].(time.Time)
			return
		}
		defer resetUtils()

		p := newProduct()

		when("p.CalculateCreatedOn() is called", func(it bdd.It) {
			p.CalculateCreatedOn()
			it("p.CreatedOn() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(p.CreatedOn(), expectedNowInMilli(args[0].(time.Time)))
			})
		})
		when("p.CalculateUpdatedOn() is called", func(it bdd.It) {
			p.CalculateUpdatedOn()
			it("p.UpdatedOn() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(p.UpdatedOn(), expectedNowInMilli(args[0].(time.Time)))
			})
		})
	}, like(
		s(timeFmt("13-01-1870 14:00:30")), s(timeFmt("22-03-2000 10:12:21")), s(timeFmt("15-06-1995 08:50:20")),
	))
}

// Feature Generate ID of Document
// - As a developer,
// - I want to be able to call generation method to set random ID for Document,
// - So that I could use Document later for indexing.
func Test_Generate_ID_of_Document(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a empty Product p", func(when bdd.When, args ...interface{}) {
		p := newProduct()

		when("p.GenerateID() is called", func(it bdd.It) {
			newID = func() (id ObjectId) {
				id = ObjectIdHex(args[0].(string))
				return
			}
			defer resetUtils()

			p.GenerateID()
			it("p.ID().Hex() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(p.ID().Hex(), args[0].(string))
			})
		})
	}, like(
		s(id1), s(id2), s(id3),
	))
}

// Feature Encoding to map object.
// - As a developer,
// - I want that Documenter to be able to convert to M object,
// - So that I can use to ease call on mgo methods.
func Test_Encoding_to_map_object(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a Product p with id '%[1]s'", func(when bdd.When, args ...interface{}) {
		p := newProduct()
		p.IDV = ObjectIdHex(args[0].(string))

		when("out, errMap := p.Map() is called", func(it bdd.It) {
			out, errMap := p.Map()

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errMap)
			})

			it("out['_id'] should be ObjectId equal to '%[1]s'", func(assert bdd.Assert) {
				id, ok := out["_id"].(ObjectId)
				assert.True(ok)
				assert.Equal(id.Hex(), args[0].(string))
			})
		})
	}, like(
		s(id1), s(id2), s(id3),
	))

	given(t, "a map m with m['_id'] equal to id '%[1]s' and p an empty Product", func(when bdd.When, args ...interface{}) {
		m := M{
			"_id": ObjectIdHex(args[0].(string)),
		}

		p := newProduct()

		when("errInit := p.Init(m) is called", func(it bdd.It) {
			errInit := p.Init(m)

			it("should return no errors", func(assert bdd.Assert) {
				assert.NoError(errInit)
			})

			it("p.ID() should equal to '%[1]s'", func(assert bdd.Assert) {
				assert.Equal(p.ID().Hex(), args[0].(string))
			})
		})
	}, like(
		s(id1), s(id2), s(id3),
	))
}
