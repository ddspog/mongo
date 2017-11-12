package handler

import (
	"fmt"
	"testing"

	"github.com/ddspog/mongo"
	"github.com/ddspog/mongo/model"
	"github.com/ddspog/trialtbl"
	"gopkg.in/mgo.v2/bson"
)

// TestHandleCast checks if the type Handle can be casted without any
// problems.
func TestHandleCast(t *testing.T) {
	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			p := newProductHandle()

			var h productHandler = p

			// Verify if Name function return correct value.
			val := h.Name() == "products"
			sig := "h.Name() == \"products\""
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleCreation checks if a type embedding Handle has functional
// getters.
func TestHandleCreation(t *testing.T) {
	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, testid),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product1id),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, product1id),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, product2id),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, product2id),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(false, product1id),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var h productHandler

		// Utility Trial
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			p := newProduct()
			p.IdV = bson.ObjectIdHex(f[0].(string))

			ph := newProductHandle()
			ph.DocumentV = p

			// Cast productHandle to productHandler.
			h = ph

			r = trialtbl.NewResult(true, "true")
			return
		})

		// Test Name() method
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			val := h.Name() == "products"
			sig := "h.Name() == \"products\""
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test Document() method
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			val := h.Document().Id() == bson.ObjectIdHex(f[0].(string))
			sig := "h.Document().Id() == bson.ObjectIdHex(\"%s\")"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleLink checks if a type embedding Handle runs Link correctly.
func TestHandleLink(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, 10),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, 15),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, 5),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, 150),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		// Test Link() method
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				mcl.ExpectCountReturn(f[0].(int))
			})

			ph := newProductHandle()

			var h productHandler = ph

			n, err := h.Link(db).Count()

			val := err == nil && n == f[0].(int)
			sig := "n, err := h.Link(db).Count(); err == nil && n == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleFind checks if a type embedding Handle runs Find correctly
// returning a document.
func TestHandleFind(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(true, productCollection[0].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true, productCollection[1].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(false, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(false, productCollection[1].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(false, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(false, productCollection[0].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
			trialtbl.NewTrial(false),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var db *mongo.MockDatabaser
		var h productHandler

		// Utility Trial
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db = make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				switch f[0] {
				case productCollection[0].Id().Hex():
					mcl.ExpectFindReturn(productCollection[0])
				case productCollection[1].Id().Hex():
					mcl.ExpectFindReturn(productCollection[1])
				default:
					mcl.ExpectFindFail(anyReason)
				}
			})

			ph := newProductHandle()
			ph.Document().SetId(bson.ObjectIdHex(f[0].(string)))

			// Cast productHandle to productHandler.
			h = ph

			r = trialtbl.NewResult(true, "true")
			return
		})

		var d model.Documenter

		// Test Find() execution
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			var err error
			d, err = h.Link(db).Find()
			val := err == nil
			sig := "d, err := h.Link(db).Find(); err == nil /* Found? */"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test Id() of Document returned
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.Id() == bson.ObjectIdHex(f[0].(string))
			sig := "d.Id() == bson.ObjectIdHex(\"%s\")"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test CreatedOn() of Document returned
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			val := d.CreatedOn() == f[0].(int64)
			sig := "d.CreatedOn() == %v"
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
}

// TestHandleFindAll checks if a type embedding Handle runs FindAll
// correctly returning documents.
func TestHandleFindAll(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	trialtbl.NewSuite(
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, productCollection[0].Id().Hex(), productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true, productCollection[0].CreatedOn(), productCollection[1].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true, productCollection[1].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(false, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(false, productCollection[1].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, productCollection[1].Id().Hex()),
			trialtbl.NewTrial(true),
			trialtbl.NewTrial(false, productCollection[0].Id().Hex()),
			trialtbl.NewTrial(false, productCollection[0].CreatedOn()),
		),
		trialtbl.NewExperiment(
			trialtbl.NewTrial(true, testid),
			trialtbl.NewTrial(false),
		),
	).Test(t, func(e *trialtbl.Experiment) {
		var db *mongo.MockDatabaser
		var h productHandler

		// Utility Trial
		e.RegisterResult(0, func(f ...interface{}) (r *trialtbl.Result) {
			db = make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
				if len(f) == 1 {
					switch f[0] {
					case productCollection[0].Id().Hex():
						mcl.ExpectFindAllReturn([]model.Documenter{productCollection[0]})
					case productCollection[1].Id().Hex():
						mcl.ExpectFindAllReturn([]model.Documenter{productCollection[1]})
					default:
						mcl.ExpectFindAllFail(anyReason)
					}
				} else {
					mcl.ExpectFindAllReturn([]model.Documenter{productCollection[0], productCollection[1]})
				}
			})

			ph := newProductHandle()
			if len(f) == 1 {
				ph.Document().SetId(bson.ObjectIdHex(f[0].(string)))
			}

			// Cast productHandle to productHandler.
			h = ph

			r = trialtbl.NewResult(true, "true")
			return
		})

		var da []product

		// Test FindAll() execution
		e.RegisterResult(1, func(f ...interface{}) (r *trialtbl.Result) {
			var err error
			da, err = h.Link(db).FindAll()
			val := err == nil
			sig := "da, err := h.Link(db).FindAll(); err == nil /* Found? */"
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test Id() of Documents returned
		e.RegisterResult(2, func(f ...interface{}) (r *trialtbl.Result) {
			var val = true
			var sig = "true"

			if val = len(f) == len(da); val {
				for i := range da {
					val = val && da[i].Id() == bson.ObjectIdHex(f[i].(string))
					newpiece := fmt.Sprintf(" && da[%v].Id() == ", i)
					sig = sig + newpiece + "bson.ObjectIdHex(\"%s\")"
				}
			} else {
				sig = "len(f) == len(da)"
			}
			r = trialtbl.NewResult(val, sig)
			return
		})

		// Test CreatedOn() of Documents returned
		e.RegisterResult(3, func(f ...interface{}) (r *trialtbl.Result) {
			var val = true
			var sig = "true"

			if val = len(f) == len(da); val {
				for i := range da {
					val = val && da[i].CreatedOn() == f[i].(int64)
					newpiece := fmt.Sprintf(" && da[%v].CreatedOn() == ", i)
					sig = sig + newpiece + "%v"
				}
			} else {
				sig = "len(f) == len(da)"
			}
			r = trialtbl.NewResult(val, sig)
			return
		})
	})
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
			val := err == ErrIdNotDefined
			sig := fmt.Sprintf("err := h.Link(db).Remove(); err == %v", ErrIdNotDefined)
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
			val := err == ErrIdNotDefined
			sig := fmt.Sprintf("err := h.Link(db).Remove(); err == %v", ErrIdNotDefined)
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
