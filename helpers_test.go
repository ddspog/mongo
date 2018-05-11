package mongo

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	// Const values to help tests legibility.
	idE = "000000000000746573746964"
	id1 = "000070726f64756374316964"
	id2 = "000070726f64756374326964"
	id3 = "000070726f64756374336964"
)

var (
	// Helpful vars for testing with database.
	fixtures = map[string]*product{
		"products.id1": newProductWithID(id1),
		"products.id2": newProductWithID(id2),
		"products.id3": newProductWithID(id3),
	}
	colFixtures = "{id1, id2, id3}"
	resetDB     func() error
)

// fixture returns the i-st  element from fixtures.
func fixture(i int) (p *product) {
	p = fixtures[fmt.Sprintf("products.id%[1]v", i)]
	return
}

// cleanChanges call resetDB ignoring its error.
func cleanChanges() {
	_ = resetDB()
}

// PrepareTestMongoAndRun setup test to run with a temporary database
// initialized with fixtures. It also tests Disconnect and Session.
func PrepareTestMongoAndRun(m *testing.M) {
	// Prepare clean env
	Disconnect()

	InitConnecter(NewTestableConnecter("", "testing", fixtures, &resetDB))

	_ = Connect()
	defer Disconnect()

	// Check session
	_ = Session()

	retCode := m.Run()
	defer os.Exit(retCode)
}

// product it's a type embedding the Document struct.
type product struct {
	IDV        ObjectId `bson:"_id,omitempty"`
	CreatedOnV int64    `bson:"created_on,omitempty"`
	UpdatedOnV int64    `bson:"updated_on,omitempty"`
}

// newProduct returns a empty product.
func newProduct() (p *product) {
	p = &product{}
	return
}

// newProductWithID returns a product with ID and CreatedOn defined.
func newProductWithID(id string) (p *product) {
	p = &product{
		IDV: ObjectIdHex(id),
	}
	p.CalculateCreatedOn()
	return
}

// New creates a new instance of the same product, used on another
// functions for clone purposes.
func (p *product) New() (doc Documenter) {
	doc = newProduct()
	return
}

// Map translates a product to a M object, more easily read by mgo
// methods.
func (p *product) Map() (out M, err error) {
	out, err = MapDocumenter(p)
	return
}

// Init translates a M received, to the product structure. It
// fills the structure fields with the values of each key in the
// M received.
func (p *product) Init(in M) (err error) {
	var doc Documenter = p
	err = InitDocumenter(in, &doc)
	return
}

// ID returns the _id attribute of a Document.
func (p *product) ID() (id ObjectId) {
	id = p.IDV
	return
}

// CreatedOn returns the created_on attribute of a Document.
func (p *product) CreatedOn() (t int64) {
	t = p.CreatedOnV
	return
}

// UpdatedOn returns the updated_on attribute of a Document.
func (p *product) UpdatedOn() (t int64) {
	t = p.UpdatedOnV
	return
}

// GenerateID creates a new id for a document.
func (p *product) GenerateID() {
	p.IDV = NewID()
}

// CalculateCreatedOn update the created_on attribute with a value
// corresponding to actual time.
func (p *product) CalculateCreatedOn() {
	p.CreatedOnV = NowInMilli()
}

// CalculateUpdatedOn update the updated_on attribute with a value
// corresponding to actual time.
func (p *product) CalculateUpdatedOn() {
	p.UpdatedOnV = NowInMilli()
}

// productHandle it's a type embedding the Handle struct, it's capable
// of storing Products.
type productHandle struct {
	*Handle
	DocumentV *product
}

// newProductHandle returns a empty productHandle.
func newProductHandle() (p *productHandle) {
	p = &productHandle{
		Handle: NewHandle("products", mgo.Index{
			Key: []string{"created_on"},
		}),
		DocumentV: newProduct(),
	}
	return
}

// Safely sets Handler to close after any operation.
func (p *productHandle) Safely() (ph *productHandle) {
	p.Handle.Safely()
	ph = p
	return
}

// Clean documents and search map values, returns Handle for chaining
// purposes.
func (p *productHandle) Clean() (ph *productHandle) {
	p.Handle.Clean()
	p.DocumentV = newProduct()
	ph = p
	return
}

// Find search on connected collection for a document matching data
// stored on productHandle and returns it.
func (p *productHandle) Find() (prod *product, err error) {
	var doc Documenter = newProduct()
	err = p.Handle.Find(p.Document(), doc)
	prod = doc.(*product)
	return
}

// FindAll search on connected collection for all documents matching
// data stored on productHandle and returns it. Accept options to alter
// query results.
func (p *productHandle) FindAll(opts ...QueryOptions) (proda []*product, err error) {
	var da []Documenter
	err = p.Handle.FindAll(p.Document(), &da, opts...)
	proda = make([]*product, len(da))
	for i := range da {
		//noinspection GoNilContainerIndexing
		proda[i] = da[i].(*product)
	}
	return
}

// Insert creates a new document with data stored on productHandle
// and put on connected collection.
func (p *productHandle) Insert() (err error) {
	err = p.Handle.Insert(p.Document())
	return
}

// Remove delete a document from connected collection matching the id
// of data stored on Handle.
func (p *productHandle) Remove() (err error) {
	err = p.Handle.Remove(p.Document().ID())
	return
}

// Remove deletes all document from connected collection matching the
// data stored on Handle.
func (p *productHandle) RemoveAll() (info *mgo.ChangeInfo, err error) {
	info, err = p.Handle.RemoveAll(p.Document())
	return
}

// Update updates document from connected collection matching the id
// received, and uses document info to update.
func (p *productHandle) Update(id ObjectId) (err error) {
	err = p.Handle.Update(id, p.Document())
	return
}

// SetDocument sets product on Handle and returns Handle for chaining
// purposes.
func (p *productHandle) SetDocument(d *product) (r *productHandle) {
	p.DocumentV = d
	r = p
	return
}

// Document returns the Document of Handle with correct type.
func (p *productHandle) Document() (d *product) {
	d = p.DocumentV
	return
}

// Set search map value for Handle and returns Handle for chaining
// purposes.
func (p *productHandle) SearchFor(s M) (r *productHandle) {
	p.SearchMapV = s
	r = p
	return
}

// timeFmt parses time well formatted.
func timeFmt(s string) (t time.Time) {
	t, _ = time.Parse("02-01-2006 15:04:05", s)
	return
}

// expectedNowInMilli returns expected return from NowInMilli function,
// given the time returned by time.Now().
func expectedNowInMilli(t time.Time) (r int64) {
	r = t.UnixNano() / int64(time.Millisecond)
	return
}

// resetUtils reset the functions named now and newID.
func resetUtils() {
	now = time.Now
	newID = bson.NewObjectId
}
