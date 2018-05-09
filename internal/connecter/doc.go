// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package connecter implements production and test connecters for MongoDB.

I've created this package to implement two models for same interface
MongoConnecter. Then created two objects implementing this interface,
the Mongo and TestMongo types.

The main reason because this code needed to be at an internal package
was due to restrictions on testing. Since TestMain on mongo package
would interfere with the testing of these connecters.

Mongo package use the interface and constructors, redeclaring on mongo
package.
*/
package connecter