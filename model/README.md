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
type product struct {
    IDV bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
    CreatedOnV int64 `json:"created_on" bson:"created_on"`
    UpdatedOnV int64 `json:"updated_on" bson:"updated_on"`
    name string `json:"name" form:"name" binding:"required" bson:"name"`
    price float32 `json:"price" form:"price" binding:"required" bson:"price"`
}

// Implement the Documenter interface.
func (p *product) ID() (id bson.ObjectId) {
    id = p.IDV
    return
}

func (p *product) CreatedOn() (t int64) {
    t = p.CreatedOnV
    return
}

func (p *product) UpdatedOn() (t int64) {
    t = p.UpdatedOnV
    return
}

// On these methods, you can use the functions implemented on this
// model package.
func (p *product) GenerateID() {
    p.IDV = model.NewID()
}

func (p *product) CalculateCreatedOn() {
    p.CreatedOnV = model.NowInMilli()
}

func (p *product) CalculateUpdatedOn() {
    p.UpdatedOnV = model.NowInMilli()
}

// Create a product variable, and try its methods.
p := product{}
p.CalculateCreatedOn()
t := p.CreatedOn()
```

## Mocking

You can mock some functionalities of this package, by mocking some
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