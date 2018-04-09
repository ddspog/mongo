# mongo [![GoDoc](https://godoc.org/github.com/ddspog/mongo?status.svg)](https://godoc.org/github.com/ddspog/mongo) [![Go Report Card](https://goreportcard.com/badge/github.com/ddspog/mongo)](https://goreportcard.com/report/github.com/ddspog/mongo) [![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](http://commitizen.github.io/cz-cli/) [![Travis CI](https://travis-ci.org/ddspog/mongo.svg?branch=master)](https://travis-ci.org/ddspog/mongo)

by [ddspog](http://hithub.com/ddspog)

Package **mongo** helps you mask the connection to MongoDB using mgo package.

## License

You are free to copy, modify and distribute **mongo** package with attribution under the terms of the MIT license. See the [LICENSE](https://github.com/ddspog/mongo/blob/master/LICENSE) file for details.

## Installation

Install **mongo** package with:

```shell
go get github.com/ddspog/mongo
```

## How to use

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

## Mocking

You can mock some functionalities of this package, by mocking the mgo
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

## Contribution

This package has some objectives from now:

* Being incorporate on mgo package (possible fork) on possible future.
* Creating real tests with MongoDB connections.

Any interest in help is much appreciated.