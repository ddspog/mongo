package handler

import (
	"errors"

	"github.com/ddspog/mongo/elements"
	"github.com/ddspog/mongo/model"
	"github.com/globalsign/mgo/bson"
)

var (
	// ErrIDNotDefined it's an error received when an ID isn't defined.
	ErrIDNotDefined = errors.New("ID not defined")
	// ErrDBNotDefined it's an error received when an DB is nil or
	// undefined.
	ErrDBNotDefined = errors.New("DB not defined")
	// ErrHandlerNotLinked it's an error received when the Handler
	// isn't linked to any collection.
	ErrHandlerNotLinked = errors.New("handler not linked to collection")
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
func (h *Handle) Find(doc model.Documenter, out model.Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		var mapped bson.M
		if mapped, err = doc.Map(); err == nil {
			var result interface{}
			if err = h.collectionV.Find(mapped).One(&result); err == nil {
				err = out.Init(result.(bson.M))
			}
		}
	}
	return
}

// FindAll search for all documents matching the doc data on
// collection connected to Handle.
func (h *Handle) FindAll(doc model.Documenter, out *[]model.Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		var mapped bson.M
		if mapped, err = doc.Map(); err == nil {
			var result []interface{}
			if err = h.collectionV.Find(mapped).All(&result); err == nil {
				tempArr := make([]model.Documenter, len(result))
				for i := range result {
					//noinspection GoNilContainerIndexing
					tempArr[i] = doc.New()
					if err := tempArr[i].Init(result[i].(bson.M)); err != nil {
						break
					}
				}

				*out = tempArr
			}
		}
	}
	return
}

// Insert puts a new document on collection connected to Handle, using
// doc data.
func (h *Handle) Insert(doc model.Documenter) (err error) {
	if doc.ID().Hex() == "" {
		doc.GenerateID()
	}

	doc.CalculateCreatedOn()

	if err = h.checkLink(); err == nil {
		var mapped bson.M
		if mapped, err = doc.Map(); err == nil {
			err = h.collectionV.Insert(mapped)
		}
	}
	return
}

// Remove delete a document on collection connected to Handle, matching
// id received.
func (h *Handle) Remove(id bson.ObjectId) (err error) {
	if id.Hex() == "" {
		err = ErrIDNotDefined
	} else {
		if err = h.checkLink(); err == nil {
			err = h.collectionV.RemoveID(id)
		}
	}
	return
}

// RemoveAll delete all documents on collection connected to Handle,
// matching the doc data.
func (h *Handle) RemoveAll(doc model.Documenter) (info *elements.ChangeInfo, err error) {
	if err = h.checkLink(); err == nil {
		var mapped bson.M
		if mapped, err = doc.Map(); err == nil {
			info, err = h.collectionV.RemoveAll(mapped)
		}
	}
	return
}

// Update updates a document on collection connected to Handle,
// matching id received, updating with the information on doc.
func (h *Handle) Update(id bson.ObjectId, doc model.Documenter) (err error) {
	if id.Hex() == "" {
		err = ErrIDNotDefined
	} else {
		doc.CalculateUpdatedOn()

		if err = h.checkLink(); err == nil {
			var mapped bson.M
			if mapped, err = doc.Map(); err == nil {
				err = h.collectionV.UpdateID(id, mapped)
			}
		}
	}
	return
}

// checkLink verifies if collection were already linked on the Handle.
func (h *Handle) checkLink() (err error) {
	if h.collectionV == nil {
		err = ErrHandlerNotLinked
	}
	return
}
