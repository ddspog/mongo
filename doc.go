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

Since mongo functions depends on MongoDB connection, the functions do
panic instead of running a error, since the program shouldn't run if
the DB connection doesn't apply.

Usage

The package can be used like this:

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

Further usage it's the same way mgo package is used. Look into mgo
docs page: https://godoc.org/gopkg.in/mgo.v2

The Connect function tries to connect to a MONGODB_URL environment
variable, but when it's not defined, it uses a default URL:

	mongodb://localhost:27017/severo-rest-db
*/
package mongo