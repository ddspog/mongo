// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package mongo mask the connection to MongoDB using mgo package.

This is made with function Connect, that saves Session and Mongo object
which will be used later from other packages. Also, I've embedded the
Collection, Database and Query types, to allow mocking via interfaces.
The embedded was necessary for the functions to use the interfaces as
return values, that way, the code can use the original, or generate
a mock of them for testing purposes.

The package can be used like this:

	// To connect with MongoDB database.
	mongo.Connect()
	defer mongo.Disconnect()

	// You can use mgo known functions with mongo.CurrentSession() or
	// mongo.Mongo(). If you want to use only the Database object to
	// handle the operations on MongoDB with a handler, use:
	mongo.ConsumeDatabaseOnSession(func(db elements.Databaser) {
		// Make db object available on handlers.
		p := NewProductHandler()
		p.Link(db)

		// ... Do other operations.
	})

Other option of usage is through the use of mongo.DatabaseSocket:

	// To connect with MongoDB database.
	mongo.Connect()
	defer mongo.Disconnect()

	// Create socket
	s := mongo.NewSocket()
	defer s.Close()

	// Make db object available on handlers.
	p := NewProductHandler()
	p.Link(s.DB())

	// ... Do other operations.

Or even through the concept of LinkedHandlers, as described later:

	// To connect with MongoDB database.
	mongo.Connect()
	defer mongo.Disconnect()

	// Create a linked handler
	p, _ := NewLinkedProductHandler()

	// ... Do other operations.

Further usage it's the same way mgo package is used. Look into mgo
docs page: https://godoc.org/github.com/globalsign/mgo

The Connect function tries to connect to a MONGODB_URL environment
variable, but when it's not defined, it uses a default URL:

	mongodb://localhost:27017/severo-rest-db

You can mock some functions of this package, by mocking the mgo
called functions mgo.ParseURL and mgo.Dial. Use the MockMongoSetup
presented on this package (only in test environment), like:

	create, _ := mongo.NewMockMongoSetup(t)
	defer create.Finish()

	create.ParseURL().Returns(db, nil)
	create.Dial().Returns(info, nil)

	// Call any preparations on connection ...
	if err := mongo.Connect(); err != nil {
		t.fail()
	}


Documenter

Mongo package also contain utility functions to help modeling documents.

The package contains a interface Documenter which contain getters for
important attributes to any document on MongoDB: _id, created_on and
updated_on. It also contains functions that generates correctly the
created_on and updated_on attributes.

The Documenter can be used like this:

	// Create a type representing the Document type
	type Product struct {
		IDV			ObjectId	`json:"_id,omitempty" bson:"_id,omitempty"`
		CreatedOnV	int64			`json:"created_on,omitempty" bson:"created_on,omitempty"`
		UpdatedOnV	int64			`json:"updated_on,omitempty" bson:"updated_on,omitempty"`
		NameV		string			`json:"name" form:"name" binding:"required" bson:"name"`
		PriceV		float32			`json:"price" form:"price" binding:"required" bson:"price"`
	}

	// Implement the Documenter interface.
	func (p *Product) ID() (id ObjectId) {
		id = p.IDV
		return
	}

	func (p *Product) CreatedOn() (t int64) {
		t = p.CreatedOnV
		return
	}

	func (p *Product) UpdatedOn() (t int64) {
		t = p.UpdatedOnV
		return
	}

	func (p *Product) New() (doc mongo.Documenter) {
		doc = &Product{}
		return
	}

	// On these methods, you can use the functions implemented mongo
	// package.
	func (p *Product) Map() (out M, err error) {
		out, err = mongo.MapDocumenter(p)
		return
	}

	func (p *Product) Init(in M) (err error) {
		var doc mongo.Documenter = p
		err = mongo.InitDocumenter(in, &doc)
		return
	}

	func (p *Product) GenerateID() {
		p.IDV = mongo.NewID()
	}

	func (p *Product) CalculateCreatedOn() {
		p.CreatedOnV = mongo.NowInMilli()
	}

	func (p *Product) CalculateUpdatedOn() {
		p.UpdatedOnV = mongo.NowInMilli()
	}

	// Create a product variable, and try its methods.
	p := Product{}
	p.CalculateCreatedOn()
	t := p.CreatedOn()

You can also mock some other functions of this package, by mocking some
called functions time.Now and NewObjectId. Use the MockModelSetup
presented on this package (only in test environment), like:

	create, _ := mongo.NewMockModelSetup(t)
	defer create.Finish()

	create.Now().Returns(time.Parse("02-01-2006", "22/12/2006"))
	create.NewID().Returns(ObjectIdHex("anyID"))

	var d mongo.Documenter
	// Call any needed methods ...
	d.GenerateID()
	d.CalculateCreatedOn()


Handle

Mongo package also enable creation of Handle, a type that connects to
database collections and do some operations.

The Handle were made to be imported on embedding type, and through
overriding of some methods, to implement an adequate Handler for a
desired type of Document. The Handle type assumes to operate on a
Documenter type, that will contain information about the operation
to made with Handle.

The package should be used to create new types. Use the Handler type
for creating embedding types.

	type ProductHandle struct {
		*mongo.Handle
		DocumentV *product.Product
	}

For each new type, a constructor may be needed, and for that Handler
has a basic constructor.

	func New() (p *ProductHandle) {
		p = &ProductHandle{
			Handle: mongo.NewHandle("products"),
			DocumentV: product.New(),
		}
		return
	}

	func NewLinked() (p *ProductHandle, err error) {
		p = &ProductHandle{
			DocumentV: product.New(),
		}
		p.Handle, err = NewLinkedHandle("products")
	}

All functions were made to be overridden and rewrite. First thing to do
it's creating the Link method, as it follows:

	func (p *ProductHandle) Link(db mongo.Databaser) (err error) {
		err = p.Handle.Link(db)
		return
	}

The creation of Insert, Remove and RemoveAll are trivial. Call it with
a Document getter function defined like:

	func (p *ProductHandle) Document() (d *product.Product) {
		d = p.DocumentV
		return
	}

	func (p *ProductHandle) Insert() (err error) {
		err = p.Handle.Insert(p.Document())
		return
	}

The Clean function is simple and helps a lot:

	func (p *ProductHandle) Clean() {
		p.Handle.Clean()
		p.DocumentV = product.New()
	}

The Update function uses an id as an argument:

	func (p *ProductHandle) Update(id ObjectId) (err error) {
		err = p.Handle.Update(id, p.Document())
		return
	}

The complicated functions are Find and FindAll which requires casting
for the Document type:

	func (p *ProductHandle) Find() (prod *product.Product, err error) {
		var doc mongo.Documenter = product.New()
		err = p.Handle.Find(p.Document(), doc)
		prod = doc.(*product.Product)
		return
	}

	func (p *ProductHandle) FindAll() (proda []*product.Product, err error) {
		var da []mongo.Documenter
		err = p.Handle.FindAll(p.Document(), &da)
		proda = make([]*product.Product, len(da))
		for i := range da {
			//noinspection GoNilContainerIndexing
			proda[i] = da[i].(*product.Product)
		}
		return
	}

For all functions written, verification it's advisable.
*/
package mongo
