# mongo [![GoDoc](https://godoc.org/github.com/ddspog/mongo?status.svg)](https://godoc.org/github.com/ddspog/mongo) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo)](https://goreportcard.com/report/github.com/ddspog/mongo) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![Travis CI](https://travis-ci.org/ddspog/mongo.svg?branch=master)](https://travis-ci.org/ddspog/mongo)

by [ddspog](http://github.com/ddspog)

Package **mongo** helps you mask the connection to MongoDB using mgo package.

## License

You are free to copy, modify and distribute **mongo** package with attribution under the terms of the MIT license. See the [LICENSE](https://github.com/ddspog/mongo/blob/master/LICENSE) file for details.

## Installation

Install **mongo** package with:

```shell
go get github.com/ddspog/mongo
```

## How to use

This package mask the connection to MongoDB using mgo package.

This is made with function Connect, that saves Session and Mongo object which will be used later from other packages. Also, I've embedded the Collection, Database and Query types, to allow mocking via interfaces. The embedded was necessary for the functions to use the interfaces as return values, that way, the code can use the original, or generate a mock of them for testing purposes.

The package can be used like this:

```go
// To connect with MongoDB database.
mongo.Connect()

// Clone the Session generated with Connect method, to allow use
// on other packages.
s := mongo.Session.Clone()
defer s.Close()

// You can use mgo known functions with mongo.Session() or
// mongo.Mongo(). If you want to use only the Database object to
// handle the operations on MongoDB with a handler, use:
mongo.ConsumeDatabaseOnSession(func(db elements.Databaser) {
    // Make db object available on handlers.
    handler.Link(db)
    // ... Do other operations.
})
```

Further usage it's the same way mgo package is used. Look into mgo
docs page: <https://godoc.org/github.com/globalsign/mgo>

The Connect function tries to connect to a MONGODB_URL environment
variable, but when it's not defined, it uses a default URL:

mongodb://localhost:27017/test

### Mocking

You can mock some functions of this package, by mocking the mgo
called functions mgo.ParseURL and mgo.Dial. Use the MockMongoSetup
presented on this package (only in test environment), like:

```go
create, _ := mongo.NewMockMongoSetup(t)
defer create.Finish()

create.ParseURL().Returns(db, nil)
create.Dial().Returns(info, nil)

// Call any preparations on connection ...
if err := mongo.Connect(); err != nil {
    t.fail()
}
```

## Documenter

Mongo package also contain utility functions to help modeling documents.

The package contains a interface Documenter which contain getters for important attributes to any document on MongoDB: _id, created_on and updated_on. It also contains functions that generates correctly the created_on and updated_on attributes.

The Documenter can be used like this:

```go
// Create a type representing the Document type
type Product struct {
    IDV bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
    CreatedOnV  int64   `json:"created_on,omitempty" bson:"created_on,omitempty"`
    UpdatedOnV  int64   `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
    NameV   string  `json:"name" form:"name" binding:"required" bson:"name"`
    PriceV  float32 `json:"price" form:"price" binding:"required" bson:"price"`
}

// Implement the Documenter interface.
func (p *Product) ID() (id bson.ObjectId) {
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
func (p *Product) Map() (out bson.M, err error) {
    out, err = mongo.MapDocumenter(p)
    return
}

func (p *Product) Init(in bson.M) (err error) {
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

You can also mock some other functions of this package, by mocking some called functions time.Now and bson.NewObjectId. Use the MockModelSetup presented on this package (only in test environment), like:

```go
create, _ := mongo.NewMockModelSetup(t)
defer create.Finish()

create.Now().Returns(time.Parse("02-01-2006", "22/12/2006"))
create.NewID().Returns(bson.ObjectIdHex("anyID"))

var d mongo.Documenter
// Call any needed methods ...
d.GenerateID()
d.CalculateCreatedOn()
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
        Handle: mongo.NewHandle(),
        DocumentV: product.New(),
    }
    return
}
```

All functions were made to be overridden and rewrite. First thing to do it's creating a Name function.

```go
func (p *ProductHandle) Name() (n string) {
    n = "products"
    return
}
```

With Name function, the creation of Link method goes as it follows:

```go
func (p *ProductHandle) Link(db mongo.Databaser) (h *ProductHandle) {
    p.Handle.Link(db, p.Name())
    h = p
    return
}
```

The creation of Insert, Remove and RemoveAll are trivial. Call it with a Document getter function defined like:

```go
func (p *ProductHandle) Document() (d *product.Product) {
    d = p.DocumentV
    return
}

func (p *ProductHandle) Insert() (err error) {
    err = p.Handle.Insert(p.Document())
    return
}
```

The Clean function is simple and helps a lot:

```go
func (p *ProductHandle) Clean() {
    p.Handle.Clean()
    p.DocumentV = product.New()
}
```

The Update function uses an id as an argument:

```go
func (p *ProductHandle) Update(id bson.ObjectId) (err error) {
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
```

For all functions written, verification it's advisable.

## Testing

This package contains a nice coverage with the unit tests, within the objectives of the project.

The elements, embedded and mocks sub-packages have low coverage because they fulfill a need to mock mgo elements. These packages only embedded mgo objects to mock, and by this a lot of unused functions were created to fulfill interface requisites.

On the other hand, model, handler and mongo package have full coverage, being the focus of this project.

The project also contains a set of acceptance tests. I've have set the test-acceptance task with the commands to run it. These tests requires a mongo test database to be available. It creates, search and remove elements from it, being reusable without broking the database.

## Contribution

This package has some objectives from now:

* Being incorporate on mgo package (possible fork) on possible future.
* Incorporate any new ideas about possible improvements.

Any interest in help is much appreciated.