// Copyright 2009 Dênnis Dantas de Sousa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package handler enable creation of Handle, a type that connects to
database collections and do some operations.

The Handle were made to be imported on embedding type, and through
overriding of some methods, to implement an adequate Handler for a
desired type of Document. The Handle type assumes to operate on a
model.Document type, that will contain informations about the operation
to made with Handle.

Usage

The package should be used to create new types. Use the Handler type
for creating embedding types.

	type ProductHandle struct {
		*handler.Handle
		DocumentV product.Product
	}

For each new type, a constructor may be needed, and for that Handler
has a basic constructor.

	func New() (p *ProductHandle) {
		p = &ProductHandle{
			Handle: handler.New(),
			DocumentV: product.New(),
		}
		return
	}

All functions were made to be overriden and rewrited. First thing to do
it's creating a Name function.

	func (p *ProductHandle) Name() (n string) {
		n = "products"
		return
	}

With Name function, the creation of Link method goes as it follows:

	func (p *ProductHandle) Link(db mongo.Databaser) (h ProductHandler) {
		p.Handle.Link(db, p.Name())
		h = p
		return
	}

The creation of Insert, Remove and RemoveAll are trivial. Call it with
a Document getter function defined like:

	func (p *ProductHandle) Document() (d product.Product) {
		d = p.DocumentV
		return
	}

	func (p *ProductHandle) Insert() (err error) {
		err = p.Handle.Insert(p.Document())
		return
	}

The complicated fucntions are Find and FindAll which requires casting
for the Document type:

	func (p *ProductHandle) Find() (prod product.Product, err error) {
		prod = product.New()
		err = p.Handle.Find(p.Document(), prod)
		return
	}

	func (p *productHandle) FindAll() (proda []product.Product, err error) {
		var da []model.Documenter
		err = p.Handle.FindAll(p.Document(), &da)
		proda = make([]product.Product, len(da))
		for i := range da {
			proda[i] = da[i].(product.Product)
		}
		return
	}

For all functions written, verification it's advisable.
*/
package handler
