# model [![GoDoc](https://godoc.org/github.com/ddspog/mongo/model?status.svg)](https://godoc.org/github.com/ddspog/mongo/model) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo/model)](https://goreportcard.com/report/github.com/ddspog/mongo/model) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/)

## Overview

Package model contain utility functions to help modeling documents.

The package contains a interface Documenter and a Document type
implementing this interface, for embedding purposes. Documenter
interface contain getters for important attributes to any document
on MongoDB: \_id, created\_on and updated_on. The Document type contains
already functions that generates correctly the created_on and
updated_on attributes.

## Usage

The package can be used like this:

```go
// Create a type embedding the Document type
type product struct {
    *model.Document
    name string        `json:"name" form:"name" binding:"required" bson:"name"`
    price float32    `json:"price" form:"price" binding:"required" bson:"price"`
}

// Create a product variable, and try its methods.
p := product{}
p.CalculateCreatedOn()
t := p.CreatedOn()
```