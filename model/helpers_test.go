package model

import "time"

const (
	// Const values to help tests legibility.
	testid     = "000000000000746573746964"
	product1id = "000070726f64756374316964"
	product2id = "000070726f64756374326964"
)

// product it's a type embedding the Document struct.
type product struct {
	*Document
}

// newProduct returns a empty product.
func newProduct() (p product) {
	p = product{
		Document: &Document{},
	}
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
