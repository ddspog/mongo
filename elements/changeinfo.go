package elements

import "gopkg.in/mgo.v2"

// ChangeInfo it's an embedded type of mgo.ChangeInfo, made to reduce
// importing to this package only, since it's necessary on other
// functions.
type ChangeInfo struct {
	*mgo.ChangeInfo
}

// NewChangeInfo creates a new embedded ChangeInfo, with all attributes
// for class.
func NewChangeInfo(u, r, m int, id interface{}) (c *ChangeInfo) {
	c = &ChangeInfo{
		ChangeInfo: &mgo.ChangeInfo{
			Updated:    u,
			Removed:    r,
			Matched:    m,
			UpsertedId: id,
		},
	}
	return
}

// NewRemoveInfo creates a new embedded ChangeInfo, with all attributes
// returned on calls to Remove.
func NewRemoveInfo(r int) (c *ChangeInfo) {
	c = NewChangeInfo(0, r, 0, nil)
	return
}
