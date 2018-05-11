package mongo

import (
	"errors"

	"github.com/globalsign/mgo"
)

var (
	// ErrIDNotDefined it's an error received when an ID isn't defined.
	ErrIDNotDefined = errors.New("ID not defined")
)

// Handle it's a type implementing the Handler interface, responsible
// of taking documents and using them to manipulate collections.
type Handle struct {
	safely            bool
	socket            *DatabaseSocket
	collection        *mgo.Collection
	collectionName    string
	collectionIndexes []mgo.Index
	SearchMapV        M
}

// NewHandle creates a new Handle to be embedded onto handle for other
// types. It also accept optional indexes to be loaded onto collection.
func NewHandle(name string, indexes ...mgo.Index) (h *Handle) {
	sk := NewSocket()
	h = &Handle{
		collectionName:    name,
		socket:            sk,
		collection:        sk.DB().C(name),
		safely:            false,
		collectionIndexes: indexes,
	}

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
	n, err = h.collection.Count()

	if h.safely {
		h.Close()
	}
	return
}

// Find search for a document matching the doc data on collection
// connected to Handle.
func (h *Handle) Find(doc Documenter, out Documenter) (err error) {
	var mapped M

	if h.IsSearchEmpty() {
		mapped, err = doc.Map()
	} else {
		mapped = h.SearchMap()
	}

	if err == nil {
		var result interface{}
		if err = h.collection.Find(mapped).One(&result); err == nil {
			err = out.Init(result.(M))
		}
	}

	if h.safely {
		h.Close()
	}
	return
}

// QueryOptions enumerates different options altering result on queries.
type QueryOptions struct {
	Sort []string
}

// FindAll search for all documents matching the doc data on
// collection connected to Handle. Accepts options to alter result.
func (h *Handle) FindAll(doc Documenter, out *[]Documenter, opts ...QueryOptions) (err error) {
	var mapped M

	if h.IsSearchEmpty() {
		mapped, err = doc.Map()
	} else {
		mapped = h.SearchMap()
	}

	if err == nil {
		var result []interface{}
		qry := h.collection.Find(mapped)

		if len(opts) == 1 {
			if opts[0].Sort != nil {
				qry = qry.Sort(opts[0].Sort...)
			}
		}

		if err = qry.All(&result); err == nil {
			tempArr := make([]Documenter, len(result))
			for i := range result {
				//noinspection GoNilContainerIndexing
				tempArr[i] = doc.New()
				if err := tempArr[i].Init(result[i].(M)); err != nil {
					break
				}
			}

			*out = tempArr
		}
	}

	if h.safely {
		h.Close()
	}
	return
}

// Insert puts a new document on collection connected to Handle, using
// doc data.
func (h *Handle) Insert(doc Documenter) (err error) {
	if doc.ID().Hex() == "" {
		doc.GenerateID()
	}

	doc.CalculateCreatedOn()

	var mapped M
	if mapped, err = doc.Map(); err == nil {
		err = h.collection.Insert(mapped)
	}

	if h.safely {
		h.Close()
	}
	return
}

// Remove delete a document on collection connected to Handle, matching
// id received.
func (h *Handle) Remove(id ObjectId) (err error) {
	if id == "" {
		err = ErrIDNotDefined
	} else {
		err = h.collection.RemoveId(id)
	}

	if h.safely {
		h.Close()
	}
	return
}

// RemoveAll delete all documents on collection connected to Handle,
// matching the doc data.
func (h *Handle) RemoveAll(doc Documenter) (info *mgo.ChangeInfo, err error) {
	var mapped M

	if h.IsSearchEmpty() {
		mapped, err = doc.Map()
	} else {
		mapped = h.SearchMap()
	}

	if err == nil {
		info, err = h.collection.RemoveAll(mapped)
	}

	if h.safely {
		h.Close()
	}
	return
}

// Update updates a document on collection connected to Handle,
// matching id received, updating with the information on doc.
func (h *Handle) Update(id ObjectId, doc Documenter) (err error) {
	if id == "" {
		err = ErrIDNotDefined
	} else {
		doc.CalculateUpdatedOn()

		var mapped M
		if mapped, err = doc.Map(); err == nil {
			delete(mapped, "_id")
			idSelector := M{
				"_id": id,
			}

			err = h.collection.Update(idSelector, mapped)
		}
	}

	if h.safely {
		h.Close()
	}
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
	for _, index := range h.collectionIndexes {
		h.collection.EnsureIndex(index)
	}
}
