package handler

import (
	"fmt"
	"testing"

	"github.com/ddspog/bdd"
	"github.com/ddspog/mongo"
	"github.com/ddspog/mongo/model"
	"github.com/ddspog/trialtbl"
	"gopkg.in/mgo.v2/bson"
)

// Feature Enable embedding with Handle
// - As a developer,
// - I want to be able to embedded Handle in other defined types,
// - So that I could use the Handle methods to store general data on DB.
func Test_Enable_embedding_with_Handle(t *testing.T) {
	given, _, _ := bdd.Sentences()

	given(t, "a new embedded ProductHandle p", func(when bdd.When) {
		p := newProductHandle()

		when("casting to ProductHandler interface h", func(it bdd.It) {
			var h productHandler = p

			it("h.Name() should return 'products'", func(assert bdd.Assert) {
				assert.Equal(h.Name(), "products")
			})
		})
	})
}

// Feature Create Handle with functional Getters
// - As a developer,
// - I want to be able to create a Handle, and access data with its getters,
// - So that I could use these getters to manipulate and read data.
func Test_Create_Handle_with_functional_Getters(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a empty ProductHandler h with Idv=bson.ObjectIdHex('%[1]v')", func(when bdd.When, args ...interface{}) {
		p := newProduct()
		p.IdV = bson.ObjectIdHex(args[0].(string))

		ph := newProductHandle()
		ph.DocumentV = p
		var h productHandler = ph

		when("h.Name() is called", func(it bdd.It) {
			it("should return 'products'", func(assert bdd.Assert) {
				assert.Equal(h.Name(), "products")
			})
		})

		when("h.Document().Id().Hex() is called", func(it bdd.It) {
			it("should return '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(h.Document().Id().Hex(), args[0].(string))
			})
		})

	}, like(
		s(testid), s(product1id), s(product2id),
	))
}

// Feature Link Handle to Database
// - As a developer,
// - I want to link Handle to database,
// - So that I can use database methods on handler.
func Test_Link_Handle_to_Database(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a empty ProductHandler h and products collection has %[1]v documents", func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			mcl.ExpectCountReturn(args[0].(int))
		})

		var h productHandler = newProductHandle()

		when("h.Link(db).Count() is called", func(it bdd.It) {
			n, err := h.Link(db).Count()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})
			it("should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(n, args[0].(int))
			})
		})
	}, like(
		s(5), s(10), s(15), s(150), s(3000), s(12301293029130),
	))
}

// Feature Find documents with Handle
// - As a developer,
// - I want to Find documents using Handle,
// - So that I can user Handler to search on database.
func Test_Find_documents_with_Handle(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	p := productCollection
	col := fmt.Sprintf("{'%[1]v', '%[2]v'}", p[0].Id().Hex(), p[1].Id().Hex())

	given(t, "a empty ProductHandler h with Id '%[1]v' and products collection with documents "+col, func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			switch args[0] {
			case p[0].Id().Hex():
				mcl.ExpectFindReturn(p[0])
			case p[1].Id().Hex():
				mcl.ExpectFindReturn(p[1])
			default:
				mcl.ExpectFindFail(anyReason)
			}
		})

		var h productHandler = newProductHandle()
		h.Document().SetId(bson.ObjectIdHex(args[0].(string)))

		when("d, err := h.Link(db).Find() is called", func(it bdd.It) {
			d, err := h.Link(db).Find()

			if args[1].(bool) {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("d.Id().Hex() should return %[1]v", func(assert bdd.Assert) {
					assert.Equal(d.Id().Hex(), args[0].(string))
				})
				it("d.CreatedOn() should return %[3]v", func(assert bdd.Assert) {
					assert.Equal(d.CreatedOn(), args[2].(int64))
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Error(err)
				})
			}
		})
	}, like(
		s(p[0].Id().Hex(), true, p[0].CreatedOn()),
		s(p[1].Id().Hex(), true, p[1].CreatedOn()),
		s(testid, false),
		s(product1id, false),
		s(product2id, false),
	))
}

// Feature Find various documents with Handle
// - As a developer,
// - I want to Find various documents using Handle,
// - So that I can use Handler to iterate through data.
func Test_Find_various_documents_with_Handle(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	p := productCollection
	col := fmt.Sprintf("{'%[1]v', '%[2]v'}", p[0].Id().Hex(), p[1].Id().Hex())

	given(t, "a empty ProductHandler h with Id '%[1]v' and products collection with documents "+col, func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			switch args[0] {
			case "":
				mcl.ExpectFindAllReturn([]model.Documenter{p[0], p[1]})
			case p[0].Id().Hex():
				mcl.ExpectFindAllReturn([]model.Documenter{p[0]})
			case p[1].Id().Hex():
				mcl.ExpectFindAllReturn([]model.Documenter{p[1]})
			default:
				mcl.ExpectFindAllFail(anyReason)
			}
		})

		var h productHandler = newProductHandle()
		if args[0].(string) != "" {
			h.Document().SetId(bson.ObjectIdHex(args[0].(string)))
		}

		when("da, err := h.Link(db).FindAll() is called", func(it bdd.It) {
			da, err := h.Link(db).FindAll()

			if args[1].(bool) {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})

				for i := range da {
					dstr := fmt.Sprintf("da[%d]", i)

					aID := args[(2*i)+2].(string)
					it(dstr+".Id().Hex() should  return "+aID, func(assert bdd.Assert) {
						assert.Equal(da[i].Id().Hex(), aID)
					})

					aCreatedOn := args[(2*i)+3].(int64)
					sCreatedOn := fmt.Sprintf("%v", aCreatedOn)
					it(dstr+".CreatedOn() should  return "+sCreatedOn, func(assert bdd.Assert) {
						assert.Equal(da[i].CreatedOn(), aCreatedOn)
					})
				}
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Error(err)
				})
			}
		})
	}, like(
		s(p[0].Id().Hex(), true, p[0].Id().Hex(), p[0].CreatedOn()),
		s(p[1].Id().Hex(), true, p[1].Id().Hex(), p[1].CreatedOn()),
		s("", true, p[0].Id().Hex(), p[0].CreatedOn(), p[1].Id().Hex(), p[1].CreatedOn()),
		s(testid, false),
		s(product1id, false),
		s(product2id, false),
	))
}

// TestHandleInsert checks if a type embedding Handle runs Insert
// correctly.
func TestHandleInsert(t *testing.T) {
	makeMGO, _ := mongo.NewMockMGOSetup(t)
	makeModel, _ := model.NewMockModelSetup(t)
	defer Finish(makeMGO, makeModel)

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(100)),
			trialtbl.NewTrial(true, int64(100)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(50)),
			trialtbl.NewTrial(false, int64(100)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(150)),
			trialtbl.NewTrial(false, int64(100)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(200), productCollection[0].Id().Hex()),
			trialtbl.NewTrial(true, int64(200)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(15), productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true, int64(15)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(1), testid),
			trialtbl.NewTrial(true, int64(1)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(1), testid),
			trialtbl.NewTrial(false, int64(2)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(2110), product1id),
			trialtbl.NewTrial(true, int64(2110)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(10), product2id),
			trialtbl.NewTrial(true, int64(10)),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var h productHandler

		// Test Insert() execution.
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db := makeMGO.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				mcl.ExpectInsertReturn()
			})

			ph := newProductHandle()
			if len(f) == 2 {
				ph.Document().SetId(bson.ObjectIdHex(f[1].(string)))
			}

			// Cast productHandle to productHandler.
			h = ph

			makeModel.NowInMilli().Returns(f[0].(int64))
			err := h.Link(db).Insert()
			val := err == nil
			sig := "err := h.Link(db).Insert(); err == nil /* Inserted? */"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test CreatedOn() attribution.
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := h.Document().CreatedOn() == f[0].(int64)
			sig := "h.Document().CreatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleRemove checks if a type embedding Handle runs Remove
// correctly.
func TestHandleRemove(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(false),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(false),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
			trialtbl.NewTrial(false),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product1id),
			trialtbl.NewTrial(false),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product2id),
			trialtbl.NewTrial(false),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(false),
			trialtbl.NewTrial(true),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var err error

		// Test Remove() execution.
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				if len(f) == 1 {
					mcl.ExpectRemoveIdReturn()
				}
			})

			ph := newProductHandle()
			if len(f) == 1 {
				ph.Document().SetId(bson.ObjectIdHex(f[0].(string)))
			}

			// Cast productHandle to productHandler.
			var h productHandler = ph

			err = h.Link(db).Remove()
			val := err == nil
			sig := "err := h.Link(db).Remove(); err == nil /* Removed? */"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test error signature.
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := err == ErrIDNotDefined
			sig := fmt.Sprintf("err := h.Link(db).Remove(); err == %v", ErrIDNotDefined)
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleRemoveAll checks if a type embedding Handle runs RemoveAll
// correctly.
func TestHandleRemoveAll(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[0].Id().Hex()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product1id),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product2id),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		// Test Remove() execution.
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				mcl.ExpectRemoveAllReturn(mongo.NewRemoveInfo(0))
			})

			ph := newProductHandle()
			if len(f) == 1 {
				ph.Document().SetId(bson.ObjectIdHex(f[0].(string)))
			}

			// Cast productHandle to productHandler.
			var h productHandler = ph

			_, err := h.Link(db).RemoveAll()
			val := err == nil
			sig := "_, err := h.Link(db).RemoveAll(); err == nil /* Removed all? */"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleUpdate checks if a type embedding Handle runs Update
// correctly.
func TestHandleUpdate(t *testing.T) {
	makeMGO, _ := mongo.NewMockMGOSetup(t)
	makeModel, _ := model.NewMockModelSetup(t)
	defer Finish(makeMGO, makeModel)

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(1), productCollection[0].Id().Hex()),
			trialtbl.NewTrial(false),
			trialtbl.NewTrial(true, int64(1)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(10), productCollection[1].Id().Hex()),
			trialtbl.NewTrial(false),
			trialtbl.NewTrial(true, int64(10)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(20), testid),
			trialtbl.NewTrial(false),
			trialtbl.NewTrial(true, int64(20)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(30), product1id),
			trialtbl.NewTrial(false),
			trialtbl.NewTrial(true, int64(30)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, int64(15), product2id),
			trialtbl.NewTrial(false),
			trialtbl.NewTrial(true, int64(15)),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(false, int64(1)),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(false, int64(1)),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var h productHandler
		var err error

		// Test Remove() execution.
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db := makeMGO.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				if len(f) == 2 {
					mcl.ExpectUpdateIdReturn()
				}
			})

			ph := newProductHandle()
			if len(f) == 2 {
				ph.Document().SetId(bson.ObjectIdHex(f[1].(string)))
			}

			// Cast productHandle to productHandler.
			h = ph

			makeModel.NowInMilli().Returns(f[0].(int64))
			err = h.Link(db).Update()
			val := err == nil
			sig := "err := h.Link(db).Update(); err == nil /* Updated? */"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test error signature.
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := err == ErrIDNotDefined
			sig := fmt.Sprintf("err := h.Link(db).Remove(); err == %v", ErrIDNotDefined)
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test UpdatedOn attribute.
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			val := h.Document().UpdatedOn() == f[0].(int64)
			sig := "h.Document().UpdatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}
