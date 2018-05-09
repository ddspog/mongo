# mongo [![GoDoc](https://godoc.org/github.com/ddspog/mongo?status.svg)](https://godoc.org/github.com/ddspog/mongo) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo)](https://goreportcard.com/report/github.com/ddspog/mongo) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![Travis CI](https://travis-ci.org/ddspog/mongo.svg?branch=master)](https://travis-ci.org/ddspog/mongo)

by [ddspog](http://github.com/ddspog)

Package **mongo** helps you mask the connection to MongoDB using mgo package.

## License

You are free to copy, modify and distribute **mongo** package with attribution under the terms of the MIT license. See the [LICENSE](https://github.com/ddspog/mongo/blob/master/LICENSE) file for details.

## Installation

Install **mongo** package with:

```shell
go get github.com/ddspog/mongo.v2
```

## How to use

This package mask the connection to MongoDB using mgo package.

This is made with function Connect, that saves Session and Mongo object which will be used later from other packages. It's possible to choose to work with a real database, or a temporary using mongo.InitConnecter function. Calling mongo.NewTestableConnecter enables use of a database loaded in a temp folder, initialized with documents given. This turns the testing process much easier.

The package can be used like this:

```go
// To connect with MongoDB database.
mongo.Connect()
defer mongo.Disconnect()

// You can use mgo known functions with mongo.Session(). If you
// want to use only the Database object to handle the operations on
// MongoDB with a handler, use:
mongo.ConsumeDatabaseOnSession(func(db elements.Databaser) {
    // Use db for operation on collections.
    db.C("products").Insert(document)

    // ... Do other operations.
})
```

Other option of usage is through the use of mongo.DatabaseSocket:

```go
// To connect with MongoDB database.
mongo.Connect()
defer mongo.Disconnect()

// Create socket
s := mongo.NewSocket()
defer s.Close()

// Make db object available on handlers.
s.DB().C("products").Insert(document)

// ... Do other operations.
```

Or even through the concept of Handlers, as described later:

```go
// To connect with MongoDB database.
mongo.Connect()
defer mongo.Disconnect()

// Create a linked handler
p := handler.NewProductHandler()
p.SetDocument(&product{
    Name: "bread",
    Price: 0.5,
}).Insert()

// ... Do other operations.
```

Further usage on some objects, are the same way mgo package is used. Look into mgo docs page: <https://godoc.org/github.com/globalsign/mgo>

The Connect function tries to connect to a MONGODB_URL environment variable, but when it's not defined, it uses a default URL:

mongodb://localhost:27017/test

## TestableConnecter

Instead of calling Connect() on production, in test environment it's advisable to use temp database using:

```go
var resetDB func() error
conn := mongo.NewTestableConnecter("", "testing", map[string]*product.Product{
    product.NewProduct("bread", "0.5"),
    product.NewProduct("cake", "2.2"),
    product.NewProduct("soda", "1.2"),
}, &resetDB)
mongo.InitConnecter(conn)

Connect()
defer Disconnect()

// Start any tests...
```

Note that the first two parameters of NewTestableConnecter are related to the temp path where database will locate. The following parameter are the fixtures, a map of documents to populate on this temp database. Lastly is a optional address of a Reset function to drop database and repopulate it.

## Documenter

Mongo package also contain utility functions to help modeling documents.

The package contains a interface Documenter which contain getters for important attributes to any document on MongoDB: _id, created_on and updated_on. It also contains functions that generates correctly the created_on and updated_on attributes.

The Documenter can be used like this:

```go
// Create a type representing the Document type
type Product struct {
    IDV ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
    CreatedOnV  int64   `json:"created_on,omitempty" bson:"created_on,omitempty"`
    UpdatedOnV  int64   `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
    NameV   string  `json:"name" form:"name" binding:"required" bson:"name"`
    PriceV  float32 `json:"price" form:"price" binding:"required" bson:"price"`
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
```

## Handle

Mongo package also enable creation of Handle, a type that connects to database collections and do some operations.

The Handle were made to be imported on embedding type, and through overriding of some methods, to implement an adequate Handler for a desired type of Document. The Handle type assumes to operate on a Documenter type, that will contain information about the operation to made with Handle.

The package should be used to create new types. Use the Handler type for creating embedding types.

```go
type ProductHandle struct {
    *mongo.Handle
    DocumentV *product.Product
}
```

For each new type, a constructor may be needed, and for that Handler has a basic constructor.

```go
func New() (p *ProductHandle) {
    p = &ProductHandle{
        Handle: mongo.NewHandle("products"),
        DocumentV: product.New(),
    }
    return
}
```

All functions were made to be overridden and rewrite. First thing to do it's creating the Clean method (returning itself is optional), as it follows:

```go
func (p *ProductHandle) Clean() (ph *ProductHandle) {
    p.Handle.Clean()
    p.DocumentV = product.New()
    ph = p
    return
}
```

Create a Document, or SearchMap getter and setters functions improving use of Handle:

```go
func (p *ProductHandle) SetDocument(d *product.Product) (r *ProductHandle) {
    p.DocumentV = d
    r = p
    return
}

func (p *ProductHandle) Document() (d *product.Product) {
    d = p.DocumentV
    return
}

// Note that SearchMap, the getter is already defined on Handle.
func (p *ProductHandle) SearchFor(s mongo.M) (r *ProductHandle) {
    p.SearchMapV = s
    r = p
    return
}
```

The creation of Insert, Remove and RemoveAll are trivial.

```go
func (p *ProductHandle) Insert() (err error) {
    err = p.Handle.Insert(p.Document())
    return
}

func (p *ProductHandle) Remove() (err error) {
    err = p.Handle.Remove(p.Document().ID())
    return
}

func (p *ProductHandle) RemoveAll() (info *mgo.ChangeInfo, err error) {
    info, err = p.Handle.RemoveAll(p.Document())
    return
}
```

The Update function uses an id as an argument:

```go
func (p *ProductHandle) Update(id mongo.ObjectId) (err error) {
    err = p.Handle.Update(id, p.Document())
    return
}
```

The complicated functions are Find and FindAll which requires casting for the Document type:

```go
func (p *ProductHandle) Find() (prod *product.Product, err error) {
    var doc mongo.Documenter = product.New()
    err = p.Handle.Find(p.Document(), doc)
    prod = doc.(*product.Product)
    return
}

// QueryOptions serve to add options on returning the query.
func (p *ProductHandle) FindAll(opts ...mongo.QueryOptions) (proda []*product.Product, err error) {
    var da []mongo.Documenter
    err = p.Handle.FindAll(p.Document(), &da, opts...)
    proda = make([]*product.Product, len(da))
    for i := range da {
        //noinspection GoNilContainerIndexing
        proda[i] = da[i].(*product.Product)
    }
    return
}
```

For all functions written, verification it's advisable.

## Testing

This package contains a nice coverage with the unit tests (currently at 99%), within the objectives of the project.

The project also contains a set of acceptance tests. I've have set the test-acceptance task with the commands to run it. These tests requires a mongo test database to be available. It creates, search and remove elements from it, being reusable without broking the database.

## Contribution

This package has some objectives from now:

* Being incorporate on mgo package (possible fork) on possible future.
* Incorporate any new ideas about possible improvements.

Any interest in help is much appreciated.