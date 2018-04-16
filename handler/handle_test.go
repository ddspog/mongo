package handler

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/mocks"
	"github.com/ddspog/mongo/model"
	"github.com/ddspog/mspec/bdd"
	"github.com/globalsign/mgo/bson"
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
	create, _ := mocks.NewMockMGOSetup(t)
	defer create.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a empty ProductHandler h and products collection has %[1]v documents", func(when bdd.When, args ...interface{}) {
		db := create.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
			mcl.ExpectCountReturn(args[0].(int))
		})

		var h productHandler = newProductHandle()

		when("h.Count() is called", func(it bdd.It) {
			n, err := h.Count()

			it(fmt.Sprintf("should return error '%[1]s'", ErrHandlerNotLinked.Error()), func(assert bdd.Assert) {
				assert.Error(err)
				assert.Equal(err.Error(), ErrHandlerNotLinked.Error())
			})

			it("should return 0", func(assert bdd.Assert) {
				assert.Equal(0, n)
			})
		})

		when("h.Link(nil) is called", func(it bdd.It) {
			errLink := h.Link(nil)

			it(fmt.Sprintf("should return error '%[1]s'", ErrDBNotDefined.Error()), func(assert bdd.Assert) {
				assert.Error(errLink)
				assert.Equal(errLink.Error(), ErrDBNotDefined.Error())
			})
		})

		when("errLink := h.Link(db) and n, errCount := h.Cound() is called", func(it bdd.It) {
			errLink := h.Link(db)
			n, errCount := h.Count()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(errLink)
				assert.Nil(errCount)
			})
			it("should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(n, args[0].(int))
			})
		})
	}, like(
		s(5), s(10), s(15), s(150), s(3000), s(12301293029130),
	))
}

// Feature Clean documents with Handle
// - As a developer,
// - I want to Clean documents on Handle,
// - So that I can reset Handle after use.
func Test_Clean_documents_with_Handle(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a ProductHandler h with Document with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		var h productHandler = newProductHandle()
		h.Document().IDV = bson.ObjectIdHex(args[0].(string))

		when("h.Clean() is called", func(it bdd.It) {
			h.Clean()

			it("should have h.Document().ID() return empty", func(assert bdd.Assert) {
				assert.Empty(h.Document().ID().Hex())
			})
		})
	}, like(
		s(testid), s(product1id), s(product2id),
	))
}

// Feature Find documents with Handle
// - As a developer,
// - I want to Find documents using Handle,
// - So that I can user Handler to search on database.
func Test_Find_documents_with_Handle(t *testing.T) {
	create, _ := mocks.NewMockMGOSetup(t)
	defer create.Finish()

	given, like, s := bdd.Sentences()

	p := productCollection
	col := fmt.Sprintf("{'%[1]v', '%[2]v'}", p[0].ID().Hex(), p[1].ID().Hex())

	given(t, "a linked ProductHandler h and products collection with documents "+col, func(when bdd.When, args ...interface{}) {
		db := func() *mocks.MockDatabaser {
			return create.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
				switch args[0] {
				case p[0].ID().Hex():
					mcl.ExpectFindReturn(bson.M{"_id": p[0].IDV, "created_on": p[0].CreatedOnV, "updated_on": p[0].UpdatedOnV})
				case p[1].ID().Hex():
					mcl.ExpectFindReturn(bson.M{"_id": p[1].IDV, "created_on": p[1].CreatedOnV, "updated_on": p[1].UpdatedOnV})
				default:
					mcl.ExpectFindFail(anyReason)
				}
			})
		}

		var h productHandler = newProductHandle()
		_ = h.Link(db())

		when("d, err := h.Find() is called with Document id '%[1]v'", func(it bdd.It) {
			h.Document().IDV = bson.ObjectIdHex(args[0].(string))
			d, err := h.Find()

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

		h.Clean()
		_ = h.Link(db())

		when("d, err := h.Find() is called with Search '_id' equal '%[1]v'", func(it bdd.It) {
			h.SearchM()["_id"] = bson.ObjectIdHex(args[0].(string))
			d, err := h.Find()

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
	create, _ := mocks.NewMockMGOSetup(t)
	defer create.Finish()

	given, like, s := bdd.Sentences()

	p := productCollection
	col := fmt.Sprintf("{'%[1]v', '%[2]v'}", p[0].ID().Hex(), p[1].ID().Hex())

	given(t, "a linked ProductHandler h with products collection with documents "+col, func(when bdd.When, args ...interface{}) {
		db := func() *mocks.MockDatabaser {
			return create.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
				switch args[0] {
				case "":
					mcl.ExpectFindAllReturn([]interface{}{
						bson.M{"_id": p[0].IDV, "created_on": p[0].CreatedOnV, "updated_on": p[0].UpdatedOnV},
						bson.M{"_id": p[1].IDV, "created_on": p[1].CreatedOnV, "updated_on": p[1].UpdatedOnV},
					})
				case p[0].ID().Hex():
					mcl.ExpectFindAllReturn([]interface{}{
						bson.M{"_id": p[0].IDV, "created_on": p[0].CreatedOnV, "updated_on": p[0].UpdatedOnV},
					})
				case p[1].ID().Hex():
					mcl.ExpectFindAllReturn([]interface{}{
						bson.M{"_id": p[1].IDV, "created_on": p[1].CreatedOnV, "updated_on": p[1].UpdatedOnV},
					})
				default:
					mcl.ExpectFindAllFail(anyReason)
				}
			})
		}

		var h productHandler = newProductHandle()
		_ = h.Link(db())

		when("da, err := h.FindAll() is called with document id '%[1]v'", func(it bdd.It) {
			if args[0].(string) != "" {
				h.Document().IDV = bson.ObjectIdHex(args[0].(string))
			}
			da, err := h.FindAll()

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

		h.Clean()
		_ = h.Link(db())

		when("da, err := h.FindAll() is called with Search '_id' equal '%[1]v'", func(it bdd.It) {
			if args[0].(string) != "" {
				h.SearchM()["_id"] = bson.ObjectIdHex(args[0].(string))
			}
			da, err := h.FindAll()

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
	makeMGO, _ := mocks.NewMockMGOSetup(t)
	makeModel, _ := model.NewMockModelSetup(t)
	defer Finish(makeMGO, makeModel)

	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := makeMGO.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
			mcl.ExpectInsertReturn()
		})

		var h productHandler = newProductHandle()
		_ = h.Link(db)
		if args[0].(string) != "" {
			h.Document().IDV = bson.ObjectIdHex(args[0].(string))
		}

		when("h.Insert() is called", func(it bdd.It) {
			makeModel.Now().Returns(args[1].(time.Time))
			err := h.Insert()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})
			it("h.Document().CreatedOn() should return %[2]v", func(assert bdd.Assert) {
				assert.Equal(h.Document().CreatedOn(), expectedNowInMilli(args[1].(time.Time)))
			})
		})
	}, like(
		s("", timeFmt("01-01-2000 00:00:01")), s("", timeFmt("02-05-2014 13:36:42")), s("", timeFmt("19-12-2017 22:59:00")),
		s(productCollection[0].ID().Hex(), timeFmt("11-11-2011 11:11:11")),
		s(productCollection[1].ID().Hex(), timeFmt("12-12-2012 00:00:00")),
		s(testid, timeFmt("01-01-0001 02:14:16")), s(product1id, timeFmt("03-10-2004 15:03:02")), s(product2id, timeFmt("15-06-1995 10:00:00")),
	))
}

// Feature Remove documents with Handle
// - As a developer,
// - I want to Remove documents using Handle,
// - So that I can use Handler to remove data.
func Test_Remove_documents_with_Handle(t *testing.T) {
	create, _ := mocks.NewMockMGOSetup(t)
	defer create.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := create.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
			if args[0].(string) != "" {
				mcl.ExpectRemoveIDReturn()
			}
		})

		var h productHandler = newProductHandle()
		_ = h.Link(db)
		if args[0].(string) != "" {
			h.Document().IDV = bson.ObjectIdHex(args[0].(string))
		}

		when("h.Remove() is called", func(it bdd.It) {
			err := h.Remove()

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
	create, _ := mocks.NewMockMGOSetup(t)
	defer create.Finish()

	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := create.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
			mcl.ExpectRemoveAllReturn(elements.NewRemoveInfo(0))
		})

		var h productHandler = newProductHandle()
		_ = h.Link(db)
		if args[0].(string) != "" {
			h.Document().IDV = bson.ObjectIdHex(args[0].(string))
		}

		when("h.RemoveAll() is called", func(it bdd.It) {
			_, err := h.RemoveAll()

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
	makeMGO, _ := mocks.NewMockMGOSetup(t)
	makeModel, _ := model.NewMockModelSetup(t)
	defer Finish(makeMGO, makeModel)

	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandler h with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		db := makeMGO.DatabaseMock("products", func(mcl *mocks.MockCollectioner) {
			if args[0].(string) != "" {
				mcl.ExpectUpdateIDReturn()
			}
		})

		var h productHandler = newProductHandle()
		_ = h.Link(db)
		if args[0].(string) != "" {
			h.Document().IDV = bson.ObjectIdHex(args[0].(string))
		}

		when("h.Update() is called", func(it bdd.It) {
			if args[0].(string) != "" {
				makeModel.Now().Returns(args[1].(time.Time))
			}
			err := h.Update(h.Document().ID())

			if args[0].(string) != "" {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("should have h.Document().UpdatedOn() return %[2]v", func(assert bdd.Assert) {
					assert.Equal(h.Document().UpdatedOn(), expectedNowInMilli(args[1].(time.Time)))
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Equal(err, ErrIDNotDefined)
				})
			}
		})
	}, like(
		s(productCollection[0].ID().Hex(), timeFmt("14-03-1998 12:15:06")), s(productCollection[1].ID().Hex(), timeFmt("22-10-1974 03:11:02")),
		s(testid, timeFmt("07-12-2007 02:48:59")), s(product1id, timeFmt("31-12-1999 23:59:59")), s(product2id, timeFmt("30-06-2019 22:14:06")), s(""),
	))
}
