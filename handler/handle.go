package handler

import (
	"errors"

	"github.com/ddspog/mongo"
	"github.com/ddspog/mongo/model"
)

// ErrIDNotDefined it's an error received when an ID isn't defined.
var ErrIDNotDefined = errors.New("ID not defined")

// Handle it's a type implementing the Handler interface, responsible
// of taking documents and using them to manipulate collections.
type Handle struct {
	collectionV mongo.Collectioner
}

// New creates a new Handle to be embedded onto handle for other types.
func New() (h *Handle) {
	h = &Handle{}
	return
}

// Link connects the database to the Handle, enabling operations.
func (h *Handle) Link(db mongo.Databaser, n string) {
	h.collectionV = db.C(n)
}

// Count returns the number of documents on collection connected to
// Handle.
func (h *Handle) Count() (n int, err error) {
	n, err = h.collectionV.Count()
	return
}

// Find search for a document matching the doc data on collection
// connected to Handle.
func (h *Handle) Find(doc model.Documenter, out *model.Documenter) (err error) {
	err = h.collectionV.Find(doc).One(out)
	return
}

// FindAll search for alls documents matching the doc data on
// collection connected to Handle.
func (h *Handle) FindAll(doc model.Documenter, out *[]model.Documenter) (err error) {
	err = h.collectionV.Find(doc).All(out)
	return
}

// Insert puts a new document on collection connected to Handle, using
// doc data.
func (h *Handle) Insert(doc model.Documenter) (err error) {
	doc.CalculateCreatedOn()
	err = h.collectionV.Insert(doc)
	return
}

// Remove delete a document on collection connected to Handle, matching
// id of doc.
func (h *Handle) Remove(doc model.Documenter) (err error) {
	if doc.ID() == "" {
		err = ErrIDNotDefined
	} else {
		err = h.collectionV.RemoveId(doc.ID())
	}
	return
}

// RemoveAll delete all documents on collection connected to Handle,
// matching the doc data.
func (h *Handle) RemoveAll(doc model.Documenter) (info *mongo.ChangeInfo, err error) {
	info, err = h.collectionV.RemoveAll(doc)
	return
}

// Update updates a document on collection connected to Handle,
// matching id on doc data, updataing with the extra information on doc.
func (h *Handle) Update(doc model.Documenter) (err error) {
	if doc.ID() == "" {
		err = ErrIDNotDefined
	} else {
		doc.CalculateUpdatedOn()
		err = h.collectionV.UpdateId(doc.ID(), doc)
	}
	return
}
