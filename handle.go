package mongo

import (
	"errors"
	"reflect"

	"github.com/globalsign/mgo"
)

var (
	// ErrIDNotDefined it's an error received when an ID isn't defined.
	ErrIDNotDefined = errors.New("ID not defined")
	// DocNotDefined it's an error received when the document received
	// is nil.
	DocNotDefined = errors.New("Document not defined")
)

// Handle it's a type implementing the Handler interface, responsible
// of taking documents and using them to manipulate collections.
type Handle struct {
	safely            bool
	socket            *DatabaseSocket
	collection        *mgo.Collection
	collectionName    string
	collectionIndexes []mgo.Index
	DocumentV         Documenter
	InternalErr       error
	SearchMapV        M
}

// NewHandle creates a new Handle to be embedded onto handle for other
// types. It needs the name for collection to link, and a document not
// nil to perform some operations. It also accept optional indexes to
// be loaded onto collection.
func NewHandle(name string, doc Documenter, indexes ...mgo.Index) (h *Handle) {
	sk := NewSocket()

	h = &Handle{
		safely:            false,
		socket:            sk,
		collection:        sk.DB().C(name),
		collectionName:    name,
		collectionIndexes: indexes,
	}

	h.SetDocument(doc)
	h.ensureIndexes()
	return
}

// Close ends connection with the MongoDB collection for this handle.
// Resets socket to be used again after relinked with Link. If no
// Socket is defined, do nothing to avoid errors.
func (h *Handle) Close() {
	if h.socket != nil {
		h.socket.Close()
		h.socket = nil
	}
}

// Safely sets Handle to close after any operation.
func (h *Handle) Safely() {
	h.safely = true
}

// Clean resets handler values.
func (h *Handle) Clean() {
	h.SearchMapV = make(map[string]interface{})

	if h.Document() != nil {
		h.SetDocument(h.Document().New())
	}

	h.Close()
	sk := NewSocket()
	h.socket = sk
	h.safely = false
	h.collection = sk.DB().C(h.Name())
	h.ensureIndexes()
}

// Name returns the name of connection that Handle can connect.
func (h *Handle) Name() (n string) {
	n = h.collectionName
	return
}

// IsSearchEmpty verify if there aren't any key defined on the SearchM
// value.
func (h *Handle) IsSearchEmpty() (result bool) {
	result = len(h.SearchMap()) == 0
	return
}

// Count returns the number of documents on collection connected to
// Handle.
func (h *Handle) Count() (n int, err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		n, err = h.collection.Count()
	}

	return
}

// Find search for a document matching the doc data on collection
// connected to Handle.
func (h *Handle) Find() (out Documenter, err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		out = h.Document().New()

		var mapped M
		if mapped, err = h.mapped(); err == nil {
			var result interface{}
			if err = h.collection.Find(mapped).One(&result); err == nil {
				err = out.Init(result.(M))
			}
		}
	}
	return
}

// QueryOptions enumerates different options altering result on queries.
type QueryOptions struct {
	Sort []string
}

// FindAll search for all documents matching the document data on
// collection connected to Handle. Accepts options to alter result.
func (h *Handle) FindAll(opts ...QueryOptions) (out []Documenter, err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		var mapped M
		if mapped, err = h.mapped(); err == nil {
			var result []interface{}
			qry := h.collection.Find(mapped)

			if len(opts) == 1 {
				if opts[0].Sort != nil {
					qry = qry.Sort(opts[0].Sort...)
				}
			}

			if err = qry.All(&result); err == nil {
				out = make([]Documenter, len(result))
				for i := 0; i < len(result) && err == nil; i++ {
					out[i] = h.Document().New()
					err = out[i].Init(result[i].(M))
				}
			}
		}
	}

	return
}

// Insert puts a new document on collection connected to Handle, using
// document data.
func (h *Handle) Insert() (err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		if h.Document().ID() == "" {
			h.Document().GenerateID()
		}

		h.Document().CalculateCreatedOn()

		var mapped M
		if mapped, err = h.mapped(); err == nil {
			// Even if the new document were made with SearchFor, it
			// add these attributes, since they're important.
			mapped["_id"] = h.Document().ID()
			mapped["created_on"] = h.Document().CreatedOn()

			err = h.collection.Insert(mapped)
		}
	}

	return
}

// Remove delete a document on collection connected to Handle, matching
// id received.
func (h *Handle) Remove(id ObjectId) (err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		if id == "" {
			err = ErrIDNotDefined
		} else {
			err = h.collection.RemoveId(id)
		}
	}

	return
}

// RemoveAll delete all documents on collection connected to Handle,
// matching the document data.
func (h *Handle) RemoveAll() (info *mgo.ChangeInfo, err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		var mapped M
		if mapped, err = h.mapped(); err == nil {
			info, err = h.collection.RemoveAll(mapped)
		}
	}

	return
}

// Update updates a document on collection connected to Handle,
// matching id received, updating with the information on doc.
func (h *Handle) Update(id ObjectId) (err error) {
	defer h.ifSafelyClose()

	if err = h.InternalErr; err == nil {
		if id == "" {
			err = ErrIDNotDefined
		} else {
			h.Document().CalculateUpdatedOn()

			var mapped M
			if mapped, err = h.mapped(); err == nil {
				delete(mapped, "_id")
				mapped["updated_on"] = h.Document().UpdatedOn()

				idSelector := M{
					"_id": id,
				}

				err = h.collection.Update(idSelector, mapped)
			}
		}
	}

	return
}

// SetDocument sets product on Handle.
func (h *Handle) SetDocument(d Documenter) {
	if reflect.ValueOf(d).IsNil() {
		h.InternalErr = DocNotDefined
	} else {
		h.InternalErr = d.Validate()
	}

	h.DocumentV = d
	return
}

// Document returns the Document of Handle.
func (h *Handle) Document() (d Documenter) {
	d = h.DocumentV
	return
}

// Set search map value for Handle and returns Handle for chaining
// purposes.
func (h *Handle) SearchFor(s M) {
	h.SearchMapV = s
	return
}

// SearchMap return the search map value of Handle.
func (h *Handle) SearchMap() (s M) {
	s = h.SearchMapV
	return
}

// ensureIndexes search for any loaded index on Handle, and set it on
// collection.
func (h *Handle) ensureIndexes() {
	for i := 0; i < len(h.collectionIndexes) && h.InternalErr == nil; i++ {
		h.InternalErr = h.collection.EnsureIndex(h.collectionIndexes[i])
	}
}

// mapped returns SearchMap if it isn't empty, or the Document mapped.
func (h *Handle) mapped() (m M, err error) {
	if h.IsSearchEmpty() {
		m, err = h.Document().Map()
	} else {
		m = h.SearchMap()
	}

	return
}

// ifSafelyClose checks if safely was activated to close socket.
func (h *Handle) ifSafelyClose() {
	if h.safely {
		h.Close()
	}
}
