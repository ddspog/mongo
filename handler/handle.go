package handler

import (
	"errors"

	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/model"
)

var (
	// ErrIDNotDefined it's an error received when an ID isn't defined.
	ErrIDNotDefined = errors.New("ID not defined")
	// ErrDBNotDefined it's an error received when an DB is nil or
	// undefined.
	ErrDBNotDefined = errors.New("DB not defined")
	// ErrHandlerNotLinked it's an error receibed when the Handler
	// isn't linked to any collection.
	ErrHandlerNotLinked = errors.New("Handler not linked to collection")
)

// Handle it's a type implementing the Handler interface, responsible
// of taking documents and using them to manipulate collections.
type Handle struct {
	collectionV elements.Collectioner
}

// New creates a new Handle to be embedded onto handle for other types.
func New() (h *Handle) {
	h = &Handle{}
	return
}

// Link connects the database to the Handle, enabling operations.
func (h *Handle) Link(db elements.Databaser, n string) (err error) {
	if db != nil {
		h.collectionV = db.C(n)
		err = nil
	} else {
		err = ErrDBNotDefined
	}
	return
}

// Count returns the number of documents on collection connected to
// Handle.
func (h *Handle) Count() (n int, err error) {
	if err = h.checkLink(); err == nil {
		n, err = h.collectionV.Count()
	}
	return
}

// Find search for a document matching the doc data on collection
// connected to Handle.
func (h *Handle) Find(doc model.Documenter, out *model.Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		err = h.collectionV.Find(doc).One(out)
	}
	return
}

// FindAll search for all documents matching the doc data on
// collection connected to Handle.
func (h *Handle) FindAll(doc model.Documenter, out []model.Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		err = h.collectionV.Find(doc).All(out)
	}
	return
}

// Insert puts a new document on collection connected to Handle, using
// doc data.
func (h *Handle) Insert(doc model.Documenter) (err error) {
	doc.CalculateCreatedOn()

	if err = h.checkLink(); err == nil {
		err = h.collectionV.Insert(doc)
	}
	return
}

// Remove delete a document on collection connected to Handle, matching
// id of doc.
func (h *Handle) Remove(doc model.Documenter) (err error) {
	if doc.ID() == "" {
		err = ErrIDNotDefined
	} else {
		if err = h.checkLink(); err == nil {
			err = h.collectionV.RemoveID(doc.ID())
		}
	}
	return
}

// RemoveAll delete all documents on collection connected to Handle,
// matching the doc data.
func (h *Handle) RemoveAll(doc model.Documenter) (info *elements.ChangeInfo, err error) {
	if err = h.checkLink(); err == nil {
		info, err = h.collectionV.RemoveAll(doc)
	}
	return
}

// Update updates a document on collection connected to Handle,
// matching id on doc data, updating with the extra information on doc.
func (h *Handle) Update(doc model.Documenter) (err error) {
	if doc.ID() == "" {
		err = ErrIDNotDefined
	} else {
		doc.CalculateUpdatedOn()

		if err = h.checkLink(); err == nil {
			err = h.collectionV.UpdateID(doc.ID(), doc)
		}
	}
	return
}

func (h *Handle) checkLink() (err error) {
	if h.collectionV == nil {
		err = ErrHandlerNotLinked
	}
	return
}
