// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package mocks defines mocks of mgo objects.

This package creates mocks objects that implements the same functions
as the mgo objects. It implements interface defined on
github.com/ddspog/mongo/elements package.

Usage

The package encourages to use the mock helper object MockMGOSetup, to
call all the mocking methods returning the desired values. It was made
to reduce number of lines when using the package.
*/package mocks
