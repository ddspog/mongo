# mongo [![GoDoc](https://godoc.org/github.com/ddspog/mongo?status.svg)](https://godoc.org/github.com/ddspog/mongo) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo)](https://goreportcard.com/report/github.com/ddspog/mongo) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![Travis CI](https://travis-ci.org/ddspog/mongo.svg?branch=master)](https://travis-ci.org/ddspog/mongo) [![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

## Overview

Package mongo mask the connection to MongoDB using mgo package.

This is made with function Connect, that saves Session and Mongo object
which will be used later from other packages. Also, I've embedded the
Collection, Database and Query types, to allow mocking via interfaces.
The embedded was necessary for the functions to use the interfaces as
return values, that way, the code can use the original, or generate
a mock of them for testing purposes.

## Usage

The package can be used like this:

```go
// To connect with MongoDB database.
mongo.Connect()

// Clone the Session generated with Connect method, to allow use
// on other packages.
s := mongo.Session.Clone()
defer s.Close()

// The Database object can be used further with the methods more
// common to mgo usage.
db := mongo.Mongo.Database

c := db.C("articles")
c.Insert(Article{
    Title: "newtitle",
    Content: "newcontent"
})
```

Further usage it's the same way mgo package is used. Look into mgo
docs page: <https://godoc.org/gopkg.in/mgo.v2>

The Connect function tries to connect to a MONGODB_URL environment
variable, but when it's not defined, it uses a default URL:

mongodb://localhost:27017/test
