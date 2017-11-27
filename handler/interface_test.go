package handler

import (
	"fmt"
	"testing"

	"github.com/ddspog/mspec/bdd"
	"github.com/ddspog/mongo"
	"github.com/ddspog/mongo/model"
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

	given(t, "a ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProduct()
		p.IDV = bson.ObjectIdHex(args[0].(string))

		ph := newProductHandle()
		ph.DocumentV = p
		var h productHandler = ph

		when("h.Name() is called", func(it bdd.It) {
			it("should return 'products'", func(assert bdd.Assert) {
				assert.Equal(h.Name(), "products")
			})
		})

		when("h.Document().ID().Hex() is called", func(it bdd.It) {
			it("should return '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(h.Document().ID().Hex(), args[0].(string))
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
	col := fmt.Sprintf("{'%[1]v', '%[2]v'}", p[0].ID().Hex(), p[1].ID().Hex())

	given(t, "a ProductHandler h with ID '%[1]v' and products collection with documents "+col, func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			switch args[0] {
			case p[0].ID().Hex():
				mcl.ExpectFindReturn(p[0])
			case p[1].ID().Hex():
				mcl.ExpectFindReturn(p[1])
			default:
				mcl.ExpectFindFail(anyReason)
			}
		})

		var h productHandler = newProductHandle()
		h.Document().SetID(bson.ObjectIdHex(args[0].(string)))

		when("d, err := h.Link(db).Find() is called", func(it bdd.It) {
			d, err := h.Link(db).Find()

			if args[1].(bool) {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("d.ID().Hex() should return %[1]v", func(assert bdd.Assert) {
					assert.Equal(d.ID().Hex(), args[0].(string))
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
		s(p[0].ID().Hex(), true, p[0].CreatedOn()),
		s(p[1].ID().Hex(), true, p[1].CreatedOn()),
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
	col := fmt.Sprintf("{'%[1]v', '%[2]v'}", p[0].ID().Hex(), p[1].ID().Hex())

	given(t, "a ProductHandler h with ID '%[1]v' and products collection with documents "+col, func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			switch args[0] {
			case "":
				mcl.ExpectFindAllReturn([]model.Documenter{p[0], p[1]})
			case p[0].ID().Hex():
				mcl.ExpectFindAllReturn([]model.Documenter{p[0]})
			case p[1].ID().Hex():
				mcl.ExpectFindAllReturn([]model.Documenter{p[1]})
			default:
				mcl.ExpectFindAllFail(anyReason)
			}
		})

		var h productHandler = newProductHandle()
		if args[0].(string) != "" {
			h.Document().SetID(bson.ObjectIdHex(args[0].(string)))
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
					it(dstr+".ID().Hex() should  return "+aID, func(assert bdd.Assert) {
						assert.Equal(da[i].ID().Hex(), aID)
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
		s(p[0].ID().Hex(), true, p[0].ID().Hex(), p[0].CreatedOn()),
		s(p[1].ID().Hex(), true, p[1].ID().Hex(), p[1].CreatedOn()),
		s("", true, p[0].ID().Hex(), p[0].CreatedOn(), p[1].ID().Hex(), p[1].CreatedOn()),
		s(testid, false),
		s(product1id, false),
		s(product2id, false),
	))
}

// Feature Insert documents with Handle
// - As a developer,
// - I want to Insert documents using Handle,
// - So that I can use Handler to insert data.
func Test_Insert_documents_with_Handle(t *testing.T) {
	makeMGO, _ := mongo.NewMockMGOSetup(t)
	makeModel, _ := model.NewMockModelSetup(t)
	defer Finish(makeMGO, makeModel)

	given, like, s := bdd.Sentences()

	given(t, "a ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := makeMGO.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			mcl.ExpectInsertReturn()
		})

		var h productHandler = newProductHandle()
		if args[0].(string) != "" {
			h.Document().SetID(bson.ObjectIdHex(args[0].(string)))
		}

		when("h.Link(db).Insert() is called", func(it bdd.It) {
			makeModel.NowInMilli().Returns(args[1].(int64))
			err := h.Link(db).Insert()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})
			it("h.Document().CreatedOn() should return %[2]v", func(assert bdd.Assert) {
				assert.Equal(h.Document().CreatedOn(), args[1].(int64))
			})
		})
	}, like(
		s("", int64(50)), s("", int64(100)), s("", int64(150)),
		s(productCollection[0].ID().Hex(), int64(200)),
		s(productCollection[1].ID().Hex(), int64(15)),
		s(testid, int64(1)), s(product1id, int64(2110)), s(product2id, int64(10)),
	))
}

// Feature Remove documents with Handle
// - As a developer,
// - I want to Remove documents using Handle,
// - So that I can use Handler to remove data.
func Test_Remove_documents_with_Handle(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			if args[0].(string) != "" {
				mcl.ExpectRemoveIDReturn()
			}
		})

		var h productHandler = newProductHandle()
		if args[0].(string) != "" {
			h.Document().SetID(bson.ObjectIdHex(args[0].(string)))
		}

		when("h.Link(db).Remove() is called", func(it bdd.It) {
			err := h.Link(db).Remove()

			if args[0].(string) != "" {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Equal(err, ErrIDNotDefined)
				})
			}
		})
	}, like(
		s(productCollection[0].ID().Hex()), s(productCollection[1].ID().Hex()),
		s(testid), s(product1id), s(product2id), s(""),
	))
}

// Feature Remove various documents with Handle
// - As a developer,
// - I want to Remove various documents using Handle,
// - So that I can use Handler to remove lots of data.
func Test_Remove_various_documents_with_Handle(t *testing.T) {
	make, _ := mongo.NewMockMGOSetup(t)
	defer make.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := make.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			mcl.ExpectRemoveAllReturn(mongo.NewRemoveInfo(0))
		})

		var h productHandler = newProductHandle()
		if args[0].(string) != "" {
			h.Document().SetID(bson.ObjectIdHex(args[0].(string)))
		}

		when("h.Link(db).RemoveAll() is called", func(it bdd.It) {
			_, err := h.Link(db).RemoveAll()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})
		})
	}, like(
		s(productCollection[0].ID().Hex()), s(productCollection[1].ID().Hex()),
		s(testid), s(product1id), s(product2id), s(""),
	))
}

// Feature Update documents with Handle
// - As a developer,
// - I want to Update documents using Handle,
// - So that I can use Handler to update data.
func Test_Update_documents_with_Handle(t *testing.T) {
	makeMGO, _ := mongo.NewMockMGOSetup(t)
	makeModel, _ := model.NewMockModelSetup(t)
	defer Finish(makeMGO, makeModel)

	given, like, s := bdd.Sentences()

	given(t, "a ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := makeMGO.DatabaseMock("products", func(mcl *mongo.MockCollectioner) {
			if args[0].(string) != "" {
				mcl.ExpectUpdateIDReturn()
			}
		})

		var h productHandler = newProductHandle()
		if args[0].(string) != "" {
			h.Document().SetID(bson.ObjectIdHex(args[0].(string)))
		}

		when("h.Link(db).Update() is called", func(it bdd.It) {
			if args[0].(string) != "" {
				makeModel.NowInMilli().Returns(args[1].(int64))
			}
			err := h.Link(db).Update()

			if args[0].(string) != "" {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("should have h.Document().UpdatedOn() return %[2]v", func(assert bdd.Assert) {
					assert.Equal(h.Document().UpdatedOn(), args[1].(int64))
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Equal(err, ErrIDNotDefined)
				})
			}
		})
	}, like(
		s(productCollection[0].ID().Hex(), int64(10)), s(productCollection[1].ID().Hex(), int64(30)),
		s(testid, int64(1)), s(product1id, int64(101)), s(product2id, int64(102)), s(""),
	))
}
