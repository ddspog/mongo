# model [![GoDoc](https://godoc.org/github.com/ddspog/mongo/model?status.svg)](https://godoc.org/github.com/ddspog/mongo/model) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo/model)](https://goreportcard.com/report/github.com/ddspog/mongo/model) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)

## Overview

Package model contain utility functions to help modeling documents.

The package contains a interface Documenter which contain getters for
important attributes to any document on MongoDB: \_id, created\_on and
updated\_on. It also contains functions that generates correctly the
created\_on and updated\_on attributes.

## Usage

The package can be used like this:

```go
// Create a type embedding the Document type
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

func (p *Product) New() (doc model.Documenter) {
    doc = &Product{}
    return
}

// On these methods, you can use the functions implemented on this
// model package.
func (p *Product) Map() (out bson.M, err error) {
    out, err = model.MapDocumenter(p)
    return
}

func (p *Product) Init(in bson.M) (err error) {
    var doc model.Documenter = p
    err = model.InitDocumenter(in, &doc)
    return
}

func (p *Product) GenerateID() {
    p.IDV = model.NewID()
}

func (p *Product) CalculateCreatedOn() {
    p.CreatedOnV = model.NowInMilli()
}

func (p *Product) CalculateUpdatedOn() {
    p.UpdatedOnV = model.NowInMilli()
}

// Create a product variable, and try its methods.
p := Product{}
p.CalculateCreatedOn()
t := p.CreatedOn()
```

## Mocking

You can mock some functions of this package, by mocking some
called functions time.Now and bson.NewObjectId. Use the MockModelSetup presented on this package (only in test environment), like:

```go
create, _ := model.NewMockModelSetup(t)
defer create.Finish()

create.Now().Returns(time.Parse("02-01-2006", "22/12/2006"))
create.NewID().Returns(bson.ObjectIdHex("anyID"))

var d model.Documenter
// Call any needed methods ...
d.GenerateID()
d.CalculateCreatedOn()
```