package elements

import "github.com/globalsign/mgo"

// DialInfo it's an embedded type of mgo.DialInfo, made to reduce
// importing to this package only, since it's necessary on other
// functions.
type DialInfo struct {
	*mgo.DialInfo
}

// NewDialInfo creates a new embedded DialInfo, with all attributes
// for class.
func NewDialInfo(info *mgo.DialInfo) (d *DialInfo) {
	d = &DialInfo{
		DialInfo: info,
	}
	return
}

// NewDatabaseInfo creates a new simplified DialInfo, with only
// Database attribute set.
func NewDatabaseInfo(db string) (d *DialInfo) {
	d = &DialInfo{
		DialInfo: &mgo.DialInfo{
			Database: db,
		},
	}
	return
}
