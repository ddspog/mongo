package mongo

import (
	"errors"

	"github.com/ddspog/mongo/elements"
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
	// ErrTryRelinkWithNoSocket it's an error received when the Handler
	// tries a second Link with no socket defined on SetSocket().
	ErrTryRelinkWithNoSocket = errors.New("no socket defined for relink")
)

// Handle it's a type implementing the Handler interface, responsible
// of taking documents and using them to manipulate collections.
type Handle struct {
	socket         *DatabaseSocket
	collection     elements.Collectioner
	collectionName string
	SearchMV       M
}

// NewHandle creates a new Handle to be embedded onto handle for other types.
func NewHandle(name string) (h *Handle) {
	h = &Handle{
		collectionName: name,
	}
	return
}

// NewLinkedHandle creates a new linked Handle to be embedded onto
// handle for other types. It fits for use on real application, since
// it tries for a direct connection with the MongoDB.
func NewLinkedHandle(name string) (h *Handle, err error) {
	h = NewHandle(name)
	h.SetSocket(NewSocket())
	err = h.Link()
	return
}

// Close ends connection with the MongoDB collection for this handle.
// Resets socket to be used again after relinked with Link. If no
// Socket is defined, do nothing to avoid errors.
func (h *Handle) Close() {
	if h.Socket() != nil {
		h.Socket().Close()
	}
}

// Name returns the name of connection that Handle can connect.
func (h *Handle) Name() (n string) {
	n = h.collectionName
	return
}

// Link connects the database to the Handle, enabling operations.
func (h *Handle) Link(db ...elements.Databaser) (err error) {
	if len(db) >= 1 {
		if db[0] != nil {
			h.collection = db[0].C(h.Name())
			err = nil
		} else {
			err = ErrDBNotDefined
		}
	} else {
		if h.Socket() != nil {
			err = h.Link(h.Socket().DB())
		} else {
			err = ErrTryRelinkWithNoSocket
		}
	}
	return
}

// Clean resets SearchM value.
func (h *Handle) Clean() {
	h.SearchMV = make(map[string]interface{})
}

// IsSearchEmpty verify if there aren't any key defined on the SearchM
// value.
func (h *Handle) IsSearchEmpty() (result bool) {
	result = len(h.SearchM()) == 0
	return
}

// Count returns the number of documents on collection connected to
// Handle.
func (h *Handle) Count() (n int, err error) {
	if err = h.checkLink(); err == nil {
		n, err = h.collection.Count()
	}
	return
}

// Find search for a document matching the doc data on collection
// connected to Handle.
func (h *Handle) Find(doc Documenter, out Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		var mapped M

		if h.IsSearchEmpty() {
			mapped, err = doc.Map()
		} else {
			mapped = h.SearchM()
		}

		if err == nil {
			var result interface{}
			if err = h.collection.Find(mapped).One(&result); err == nil {
				err = out.Init(result.(M))
			}
		}
	}
	return
}

// FindAll search for all documents matching the doc data on
// collection connected to Handle.
func (h *Handle) FindAll(doc Documenter, out *[]Documenter) (err error) {
	if err = h.checkLink(); err == nil {
		var mapped M

		if h.IsSearchEmpty() {
			mapped, err = doc.Map()
		} else {
			mapped = h.SearchM()
		}

		if err == nil {
			var result []interface{}
			if err = h.collection.Find(mapped).All(&result); err == nil {
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

	if err = h.checkLink(); err == nil {
		var mapped M
		if mapped, err = doc.Map(); err == nil {
			err = h.collection.Insert(mapped)
		}
	}
	return
}

// Remove delete a document on collection connected to Handle, matching
// id received.
func (h *Handle) Remove(id ObjectId) (err error) {
	if id.Hex() == "" {
		err = ErrIDNotDefined
	} else {
		if err = h.checkLink(); err == nil {
			err = h.collection.RemoveID(id)
		}
	}
	return
}

// RemoveAll delete all documents on collection connected to Handle,
// matching the doc data.
func (h *Handle) RemoveAll(doc Documenter) (info *elements.ChangeInfo, err error) {
	if err = h.checkLink(); err == nil {
		var mapped M

		if h.IsSearchEmpty() {
			mapped, err = doc.Map()
		} else {
			mapped = h.SearchM()
		}

		if err == nil {
			info, err = h.collection.RemoveAll(mapped)
		}
	}
	return
}

// Update updates a document on collection connected to Handle,
// matching id received, updating with the information on doc.
func (h *Handle) Update(id ObjectId, doc Documenter) (err error) {
	if id.Hex() == "" {
		err = ErrIDNotDefined
	} else {
		doc.CalculateUpdatedOn()

		if err = h.checkLink(); err == nil {
			var mapped M
			if mapped, err = doc.Map(); err == nil {
				err = h.collection.UpdateID(id, mapped)
			}
		}
	}
	return
}

// SearchM return the search map value of Handle.
func (h *Handle) SearchM() (s M) {
	s = h.SearchMV
	return
}

// SetSocket defines a database socket to use on Handle for linking.
func (h *Handle) SetSocket(s *DatabaseSocket) {
	h.socket = s
}

// Socket returns the database socket used on Handle for linking.
func (h *Handle) Socket() (s *DatabaseSocket) {
	s = h.socket
	return
}

// checkLink verifies if collection were already linked on the Handle.
func (h *Handle) checkLink() (err error) {
	if h.collection == nil {
		err = ErrHandlerNotLinked
	}
	return
}
