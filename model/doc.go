// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package model contain utility functions to help modeling documents.

The package contains a interface Documenter and a Document type
implementing this interface, for embedding purposes. Documenter
interface contain getters for important attributes to any document
on MongoDB: _id, created_on and updated_on. The Document type contains
already functions that generates correctly the created_on and
updated_on attributes.

Usage

The package can be used like this:

	// Create a type embedding the Document type
	type product struct {
		*model.Document
		name string		`json:"name" form:"name" binding:"required" bson:"name"`
		price float32	`json:"price" form:"price" binding:"required" bson:"price"`
	}

	// Create a product variable, and try its methods.
	p := product{}
	p.CalculateCreatedOn()
	t := p.CreatedOn()

Mocking

You can mock some functionalities of this package, by mocking some
called functions time.Now and bson.NewObjectId. Use the MockModelSetup
presented on this package (only in test environment), like:

	create, _ := model.NewMockModelSetup(t)
	defer create.Finish()

	create.Now().Returns(time.Parse("02-01-2006", "22/12/2006"))
	create.NewID().Returns(bson.ObjectIdHex("anyID"))

	var d model.Documenter
	// Call any needed methods ...
	d.GenerateID()
	d.CalculateCreatedOn()
*/
package model
