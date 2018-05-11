// +build !acceptance

package mongo

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddspog/mspec/bdd"
)

// TestMain setup the testable mongo connecter to run a temp database.
func TestMain(m *testing.M) {
	PrepareTestMongoAndRun(m)
}

// Feature Create Handle with functional Getters
// - As a developer,
// - I want to be able to create a Handle, and access data with its getters,
// - So that I could use these getters to manipulate and read data.
func Test_Create_Handle_with_functional_Getters(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a ProductHandle p with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProductHandle().SetDocument(&product{
			IDV: ObjectIdHex(args[0].(string)),
		})
		defer p.Close()

		when("p.Name() is called", func(it bdd.It) {
			it("should return 'products'", func(assert bdd.Assert) {
				assert.Equal("products", p.Name())
			})
		})

		when("p.Document().ID().Hex() is called", func(it bdd.It) {
			it("should return '%[1]v'", func(assert bdd.Assert) {
				assert.Equal(args[0].(string), p.Document().ID().Hex())
			})
		})

	}, like(
		s(id1), s(id2), s(id3),
	))
}

// Feature Count documents with Handle
// - As a developer,
// - I want to count documents with Handle,
// - So that I can use this to perform verifications and such.
func Test_Count_documents_with_Handle(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a empty ProductHandle and products collection has %[1]v documents", func(when bdd.When, args ...interface{}) {
		when("p.Count() is called", func(it bdd.It) {
			n, err := newProductHandle().Safely().Count()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})
			it("should return %[1]v", func(assert bdd.Assert) {
				assert.Equal(args[0].(int), n)
			})
		})
	}, like(
		s(len(fixtures)),
	))
}

// Feature Clean documents with Handle
// - As a developer,
// - I want to Clean documents on Handle,
// - So that I can reset Handle after use.
func Test_Clean_documents_with_Handle(t *testing.T) {
	given, like, s := bdd.Sentences()

	given(t, "a ProductHandle p with Document with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProductHandle().SetDocument(&product{
			IDV: ObjectIdHex(args[0].(string)),
		})
		defer p.Close()

		when("h.Clean() is called", func(it bdd.It) {
			p.Clean()

			it("should have p.Document().ID() return empty", func(assert bdd.Assert) {
				assert.Empty(p.Document().ID().Hex())
			})
		})
	}, like(
		s(id1), s(id2), s(id3),
	))
}

// Feature Find documents with Handle
// - As a developer,
// - I want to Find documents using Handle,
// - So that I can user Handler to search on database.
func Test_Find_documents_with_Handle(t *testing.T) {
	defer cleanChanges()
	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandle p and products collection with documents "+colFixtures, func(when bdd.When, args ...interface{}) {
		p := newProductHandle()

		VerifyFind := func(it bdd.It, args ...interface{}) {
			d, err := p.Safely().Find()

			if args[1].(bool) {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("d.ID().Hex() should return %[1]v", func(assert bdd.Assert) {
					assert.Equal(args[0].(string), d.ID().Hex())
				})
				it("d.CreatedOn() should return %[3]v", func(assert bdd.Assert) {
					assert.Equal(args[2].(int64), d.CreatedOn())
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Error(err)
				})
			}
		}

		when("d, err := p.Find() is called with Document id '%[1]v'", func(it bdd.It) {
			p.Document().IDV = ObjectIdHex(args[0].(string))

			VerifyFind(it, args...)
		})

		p.Clean()

		when("d, err := p.Find() is called with Search '_id' equal '%[1]v'", func(it bdd.It) {
			p.SearchFor(M{
				"_id": ObjectIdHex(args[0].(string)),
			})

			VerifyFind(it, args...)
		})
	}, like(
		s(fixture(1).ID().Hex(), true, fixture(1).CreatedOn()),
		s(fixture(2).ID().Hex(), true, fixture(2).CreatedOn()),
		s(fixture(3).ID().Hex(), true, fixture(3).CreatedOn()),
		s(idE, false),
	))
}

// Feature Find various documents with Handle
// - As a developer,
// - I want to Find various documents using Handle,
// - So that I can use Handler to iterate through data.
func Test_Find_various_documents_with_Handle(t *testing.T) {
	defer cleanChanges()
	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandle p with products collection with documents "+colFixtures, func(when bdd.When, args ...interface{}) {
		p := newProductHandle()

		VerifyFindAll := func(it bdd.It, args ...interface{}) {
			da, err := p.Safely().FindAll(QueryOptions{
				Sort: []string{"_id"},
			})

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})

			if args[1].(bool) {
				it("should return some documents", func(assert bdd.Assert) {
					assert.NotEqual(0, len(da))
				})

				for i := range da {
					dstr := fmt.Sprintf("da[%d]", i)

					aID := args[(2*i)+2].(string)
					it(dstr+".ID().Hex() should  return "+aID, func(assert bdd.Assert) {
						assert.Equal(aID, da[i].ID().Hex())
					})

					aCreatedOn := args[(2*i)+3].(int64)
					sCreatedOn := fmt.Sprintf("%v", aCreatedOn)
					it(dstr+".CreatedOn() should  return "+sCreatedOn, func(assert bdd.Assert) {
						assert.Equal(aCreatedOn, da[i].CreatedOn())
					})
				}
			} else {
				it("should return no documents", func(assert bdd.Assert) {
					assert.Equal(0, len(da))
				})
			}
		}

		when("da, err := p.FindAll() is called with document id '%[1]v'", func(it bdd.It) {
			if args[0].(string) != "" {
				p.Document().IDV = ObjectIdHex(args[0].(string))
			}

			VerifyFindAll(it, args...)
		})

		p.Clean()

		when("da, err := h.FindAll() is called with Search '_id' equal '%[1]v'", func(it bdd.It) {
			if args[0].(string) != "" {
				p.SearchMap()["_id"] = ObjectIdHex(args[0].(string))
			}

			VerifyFindAll(it, args...)
		})
	}, like(
		s(fixture(1).ID().Hex(), true, fixture(1).ID().Hex(), fixture(1).CreatedOn()),
		s(fixture(2).ID().Hex(), true, fixture(2).ID().Hex(), fixture(2).CreatedOn()),
		s(fixture(3).ID().Hex(), true, fixture(3).ID().Hex(), fixture(3).CreatedOn()),
		s(
			"", true,
			fixture(1).ID().Hex(), fixture(1).CreatedOn(),
			fixture(2).ID().Hex(), fixture(2).CreatedOn(),
			fixture(3).ID().Hex(), fixture(3).CreatedOn(),
		),
		s(idE, false),
	))
}

// Feature Insert documents with Handle
// - As a developer,
// - I want to Insert documents using Handle,
// - So that I can use Handler to insert data.
func Test_Insert_documents_with_Handle(t *testing.T) {
	defer cleanChanges()
	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandle p with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProductHandle()

		if args[0].(string) != "" {
			p.Document().IDV = ObjectIdHex(args[0].(string))
		}

		when("p.Insert() is called", func(it bdd.It) {
			now = func() (t time.Time) {
				t = args[1].(time.Time)
				return
			}
			defer resetUtils()
			err := p.Safely().Insert()

			if args[2].(bool) {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("p.Document().CreatedOn() should return %[2]v", func(assert bdd.Assert) {
					assert.Equal(expectedNowInMilli(args[1].(time.Time)), p.Document().CreatedOn())
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Error(err)
				})
				it("p.Document().CreatedOn() should return %[2]v", func(assert bdd.Assert) {
					assert.Equal(expectedNowInMilli(args[1].(time.Time)), p.Document().CreatedOn())
				})
			}
		})
	}, like(
		s("", timeFmt("01-01-2000 00:00:01"), true),
		s("", timeFmt("02-05-2014 13:36:42"), true),
		s("", timeFmt("19-12-2017 22:59:00"), true),
		s(fixture(1).ID().Hex(), timeFmt("11-11-2011 11:11:11"), false),
		s(fixture(2).ID().Hex(), timeFmt("12-12-2012 00:00:00"), false),
		s(fixture(3).ID().Hex(), timeFmt("01-01-0001 02:14:16"), false),
	))
}

// Feature Remove documents with Handle
// - As a developer,
// - I want to Remove documents using Handle,
// - So that I can use Handler to remove data.
func Test_Remove_documents_with_Handle(t *testing.T) {
	defer cleanChanges()
	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandle p with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProductHandle().SetDocument(&product{
			IDV: args[0].(ObjectId),
		})

		when("p.Remove() is called", func(it bdd.It) {
			err := p.Safely().Remove()

			if args[0].(ObjectId) != "" {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Equal(ErrIDNotDefined, err)
				})
			}
		})
	}, like(
		s(fixture(1).ID()),
		s(fixture(2).ID()),
		s(fixture(3).ID()),
		s(ObjectId("")),
	))
}

// Feature Remove various documents with Handle
// - As a developer,
// - I want to Remove various documents using Handle,
// - So that I can use Handler to remove lots of data.
func Test_Remove_various_documents_with_Handle(t *testing.T) {
	defer cleanChanges()
	given, like, s := bdd.Sentences()

	given(t, "a linked ProductHandle p with ID '%[1]v'", func(when bdd.When, args ...interface{}) {
		p := newProductHandle()

		VerifyRemoveAll := func(it bdd.It, args ...interface{}) {
			info, err := p.Safely().RemoveAll()

			it("should return no errors", func(assert bdd.Assert) {
				assert.Nil(err)
			})

			it("should have removed %[2]v documents", func(assert bdd.Assert) {
				assert.Equal(args[1].(int), info.Removed)
			})

		}

		cleanChanges()

		when("_, err := p.RemoveAll() is called with document id '%[1]v'", func(it bdd.It) {
			if args[0].(string) != "" {
				p.Document().IDV = ObjectIdHex(args[0].(string))
			}
			VerifyRemoveAll(it, args...)
		})

		p.Clean()
		cleanChanges()

		when("_, err := p.RemoveAll() is called with Search '_id' equal '%[1]v'", func(it bdd.It) {
			if args[0].(string) != "" {
				p.SearchMap()["_id"] = ObjectIdHex(args[0].(string))
			}
			VerifyRemoveAll(it, args...)
		})
	}, like(
		s(fixture(1).ID().Hex(), 1),
		s(fixture(2).ID().Hex(), 1),
		s(fixture(3).ID().Hex(), 1),
		s("", 3),
	))
}

// Feature Update documents with Handle
// - As a developer,
// - I want to Update documents using Handle,
// - So that I can use Handler to update data.
func Test_Update_documents_with_Handle(t *testing.T) {
	defer cleanChanges()
	given, like, s := bdd.Sentences()

	given(t, "a linked empty ProductHandle p", func(when bdd.When, args ...interface{}) {
		p := newProductHandle()

		when("p.Update('%[1]v') is called", func(it bdd.It) {
			now = func() (t time.Time) {
				t = args[1].(time.Time)
				return
			}
			defer resetUtils()

			err := p.Safely().Update(args[0].(ObjectId))

			if args[0].(ObjectId) != "" {
				it("should return no errors", func(assert bdd.Assert) {
					assert.Nil(err)
				})
				it("should have p.Document().UpdatedOn() return %[2]v", func(assert bdd.Assert) {
					assert.Equal(expectedNowInMilli(args[1].(time.Time)), p.Document().UpdatedOn())
				})
			} else {
				it("should return an error", func(assert bdd.Assert) {
					assert.Equal(ErrIDNotDefined, err)
				})
			}
		})
	}, like(
		s(fixture(1).ID(), timeFmt("14-03-1998 12:15:06")),
		s(fixture(2).ID(), timeFmt("22-10-1974 03:11:02")),
		s(fixture(3).ID(), timeFmt("07-12-2007 02:48:59")),
		s(ObjectId(""), timeFmt("11-2-2037 01:53:21")),
	))

}
