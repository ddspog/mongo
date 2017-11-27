package model

import (
	"testing"

	"gopkg.in/ddspog/mspec.v1/bdd"

	"gopkg.in/mgo.v2/bson"
)

// Feature Enable embedding with Document
// - As a developer,
// - I want to be able to embedded Document in other defined types,
// - So that I could use the Document methods to abstract data on it.
func Test_Enable_embedding_with_Document(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a new embedded Product p with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProduct()
		p.IDV = bson.ObjectIdHex(args[0].(string))

		when("casting to Documenter interface d", func(it bdd.It) {
			var d Documenter = p
			it("d.ID().Hex() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(d.ID().Hex(), args[0].(string))
			})
		})
	}, like(
		s(testid), s(product1id), s(product2id),
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
		p.IDV = bson.ObjectIdHex(args[0].(string))
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
		s(testid, int64(123), int64(321)),
		s(product1id, int64(10), int64(1)),
		s(product2id, int64(20), int64(2)),
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

		when("p.SetID(bson.ObjectIdHex(%[1]v)) is called", func(it bdd.It) {
			p.SetID(bson.ObjectIdHex(args[0].(string)))
			it("p.ID().Hex() should return '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(p.ID().Hex(), args[0].(string))
			})
		})
		when("p.SetCreatedOn(%[2]v) is called", func(it bdd.It) {
			p.SetCreatedOn(args[1].(int64))
			it("p.CreatedOn() should return %[2]v", func(assert bdd.Assert) {
				assert.Equal(p.CreatedOn(), args[1].(int64))
			})
		})
		when("p.SetUpdatedOn(%[3]v) is called", func(it bdd.It) {
			p.SetUpdatedOn(args[2].(int64))
			it("p.UpdatedOn() should return %[3]v", func(assert bdd.Assert) {
				assert.Equal(p.UpdatedOn(), args[2].(int64))
			})
		})
	}, like(
		s(testid, int64(123), int64(321)),
		s(product1id, int64(10), int64(1)),
		s(product2id, int64(20), int64(2)),
	))
}

// Feature Calculate Document values
// - As a developer,
// - I want to be able to call calculation methods to set some values with current time,
// - So that I could use these values later for data analysis.
func Test_Calculate_Document_values(t *testing.T) {
	make, _ := NewMockModelSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a empty Product p at current time %[1]v", func(when bdd.When, args ...interface{}) {
		make.NowInMilli().Returns(args[0].(int64))
		p := newProduct()

		when("p.CalculateCreatedOn() is called", func(it bdd.It) {
			p.CalculateCreatedOn()
			it("p.CreatedOn() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(p.CreatedOn(), args[0].(int64))
			})
		})
		when("p.CalculateUpdatedOn() is called", func(it bdd.It) {
			p.CalculateUpdatedOn()
			it("p.UpdatedOn() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(p.UpdatedOn(), args[0].(int64))
			})
		})
	}, like(
		s(int64(123)), s(int64(321)), s(int64(10)), s(int64(1)),
	))
}

// Feature Generate ID of Document
// - As a developer,
// - I want to be able to call generation method to set random ID for Document,
// - So that I could use Document later for indexing.
func Test_Generate_ID_of_Document(t *testing.T) {
	make, _ := NewMockModelSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a empty Product p", func(when bdd.When, args ...interface{}) {
		p := newProduct()

		when("p.GenerateID() is called", func(it bdd.It) {
			make.NewID().Returns(bson.ObjectIdHex(args[0].(string)))
			p.GenerateID()
			it("p.ID().Hex() should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(p.ID().Hex(), args[0].(string))
			})
		})
	}, like(
		s(testid), s(product1id), s(product2id),
	))
}
