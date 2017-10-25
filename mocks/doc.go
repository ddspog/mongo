// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package mocks enable mgo objects mocking.

This is made using gomock package, that generated the mocks using the
interface defined on mongo package. The mocks follow interfaces
Collectioner, Databaser and Querier, that are capable of running same
functions as the ones in mgo package.

Usage

The package can be used like this:

	// Create mock controller first.
	c = gomock.NewController(t)
	defer c.Finish()

	// Create Database mock, and make the desired collection to return
	// a Collection Mock.
	mdb := mocks.NewMockDatabaser(c)
	mcl := mocks.NewMockCollectioner(c)

	mcl.EXPECT().Insert(gomock.Any()).Return(nil)

	mdb.EXPECT().C("collection").AnyTimes().Return(mcl)

For further usage look onto GoMock Framework Docs: https://godoc.org/github.com/golang/mock/gomock
*/
package mongo
