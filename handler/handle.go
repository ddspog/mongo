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
		var mapped bson.M
		if mapped, err = encode(doc); err == nil {
			var result interface{}
			if err = h.collectionV.Find(mapped).One(&result); err == nil {
				err = decode(result, out)
			}
		}
	}
	return
}

// FindAll search for all documents matching the doc data on
// collection connected to Handle.
func (h *Handle) FindAll(doc model.Documenter, newDoc func() model.Documenter, out *[]model.Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		var mapped bson.M
		if mapped, err = encode(doc); err == nil {
			var result []interface{}
			if err = h.collectionV.Find(mapped).All(&result); err == nil {
				err = decodeAll(result, newDoc, out)
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
		if mapped, err = encode(doc); err == nil {
			err = h.collectionV.Insert(mapped)
		}
	}
	return
}

// Remove delete a document on collection connected to Handle, matching
// id of doc.
func (h *Handle) Remove(doc model.Documenter) (err error) {
	if doc.ID().Hex() == "" {
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
		var mapped bson.M
		if mapped, err = encode(doc); err == nil {
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
			if mapped, err = encode(doc); err == nil {
				err = h.collectionV.UpdateID(id, mapped)
			}
		}
	}
	return
}

// checklink verifies if collection were already linked on the Handle.
func (h *Handle) checkLink() (err error) {
	if h.collectionV == nil {
		err = ErrHandlerNotLinked
	}
	return
}

// encode translates a model.Documenter in whathever structure it has,
// to a bson.M object, more easily read by mgo.Collection methods.
func encode(in model.Documenter) (out bson.M, err error) {
	var buf []byte
	var target interface{}

	if buf, err = bson.Marshal(in); err == nil {
		if err = bson.Unmarshal(buf, &target); err == nil {
			out = target.(bson.M)
		}
	}

	return
}

// decode translates a bson.M received as a interface{}, to the
// model.Documenter strucutre received as a pointer. It fills the
// structure fields with the values of each key in the bson.M received.
func decode(in interface{}, out *model.Documenter) (err error) {
	var marshalled []byte

	if marshalled, err = bson.Marshal(in); err == nil {
		err = bson.Unmarshal(marshalled, *out)
	}
	return
}

// decodeAll applies the decode on an array of bson.M objects,
// translating to an array of model.Documenter.
func decodeAll(in []interface{}, newDoc func() model.Documenter, out *[]model.Documenter) (err error) {
	outa := make([]model.Documenter, len(in))
	for i := range in {
		//noinspection GoNilContainerIndexing
		outa[i] = newDoc()
		if err := decode(in[i], &outa[i]); err != nil {
			break
		}
	}
	*out = outa
	return
}
