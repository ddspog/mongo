package mongo

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddspog/mongo/elements"
	"github.com/globalsign/mgo/bson"
)

// MockMongoSetup it's a setup type for configuring mocking for connect
// functions.
type MockMongoSetup struct {
	originalParseURL func(string) (*elements.DialInfo, error)
	originalDial     func(string) (elements.Sessioner, error)
	tempParseURL     *fakeParseURL
	tempDial         *fakeDial
}

// NewMockMongoSetup returns a new MockMongoSetup, already configuring
// a FakeParseURL function, and a FakeDial function on the setup. It
// requires a test environment to be running.
func NewMockMongoSetup(t *testing.T) (s *MockMongoSetup, err error) {
	if t == nil {
		err = fmt.Errorf("run only on test environment")
	} else {
		s = &MockMongoSetup{
			originalParseURL: parseURL,
			originalDial:     dial,
			tempParseURL:     newFakeParseURL(),
			tempDial:         newFakeDial(),
		}
		s.tempParseURL.mockMongoSetupP = s
		s.tempDial.mockMongoSetupP = s
	}
	return
}

// Finish restore the functions mocked to the original ones.
func (s *MockMongoSetup) Finish() {
	parseURL = s.originalParseURL
	dial = s.originalDial
}

// ParseURL returns the fake ParseURL object on this setup.
func (s *MockMongoSetup) ParseURL() (f FakeParseURLer) {
	f = s.tempParseURL
	return
}

// Dial returns the fake Dial object on this setup.
func (s *MockMongoSetup) Dial() (f FakeDialer) {
	f = s.tempDial
	return
}

// updateParseURL updates the ParseURL function with a mocked one.
func (s *MockMongoSetup) updateParseURL() {
	parseURL = s.ParseURL().GetFunction()
}

// updateDial updates the dial function with a mocked one.
func (s *MockMongoSetup) updateDial() {
	dial = s.Dial().GetFunction()
}

// FakeParseURLer it's a function mocking object, needed for mock
// purposes.
type FakeParseURLer interface {
	Returns(*elements.DialInfo, error)
	GetFunction() func(string) (*elements.DialInfo, error)
}

// fakeParseURL it's a type that enable mocking of function parseURL.
type fakeParseURL struct {
	returnInfoV     *elements.DialInfo
	returnErrorV    error
	mockMongoSetupP *MockMongoSetup
}

// newFakeParseURL returns a new FakeParseURL object.
func newFakeParseURL() (f *fakeParseURL) {
	f = &fakeParseURL{}
	return
}

// Returns ensures values to be returned on calls to parseURL.
func (f *fakeParseURL) Returns(i *elements.DialInfo, err error) {
	f.returnInfoV = i
	f.returnErrorV = err
	f.mockMongoSetupP.updateParseURL()
}

// GetFunction creates a version of parseURL that returns value demanded.
func (f *fakeParseURL) GetFunction() (fn func(string) (*elements.DialInfo, error)) {
	fn = func(u string) (i *elements.DialInfo, err error) {
		i = f.returnInfoV
		err = f.returnErrorV
		return
	}
	return
}

// FakeDialer it's a function mocking object, needed for mock
// purposes.
type FakeDialer interface {
	Returns(elements.Sessioner, error)
	GetFunction() func(string) (elements.Sessioner, error)
}

// fakeDial it's a type that enable mocking of function dial.
type fakeDial struct {
	returnSessionV  elements.Sessioner
	returnErrorV    error
	mockMongoSetupP *MockMongoSetup
}

// newFakeDial returns a new FakeDial object.
func newFakeDial() (f *fakeDial) {
	f = &fakeDial{}
	return
}

// Returns ensures a value to be returned on calls to dial.
func (f *fakeDial) Returns(s elements.Sessioner, err error) {
	f.returnSessionV = s
	f.returnErrorV = err
	f.mockMongoSetupP.updateDial()
}

// GetFunction creates a version of dial that returns value demanded.
func (f *fakeDial) GetFunction() (fn func(string) (elements.Sessioner, error)) {
	fn = func(u string) (s elements.Sessioner, err error) {
		s = f.returnSessionV
		err = f.returnErrorV
		return
	}
	return
}

// MockModelSetup it's a setup type for configuring mocking for util
// functions.
type MockModelSetup struct {
	originalNow func() time.Time
	tempNow     *fakeNow
	tempNewID   *fakeNewID
}

// NewMockModelSetup returns a new MockModelSetup, already configuring
// a FakeNow function, and a FakeNewID function  on the setup. It
// requires a test environment to be running.
func NewMockModelSetup(t *testing.T) (s *MockModelSetup, err error) {
	if t == nil {
		err = fmt.Errorf("run only on test environment")
	} else {
		s = &MockModelSetup{
			originalNow: now,
			tempNow:     newFakeNow(),
			tempNewID:   newFakeNewID(),
		}
		s.tempNow.mockModelSetupP = s
		s.tempNewID.mockModelSetupP = s
	}
	return
}

// Finish restore the functions mocked to the original ones.
func (s *MockModelSetup) Finish() {
	now = s.originalNow
}

// Now returns the fake Now object on this setup.
func (s *MockModelSetup) Now() (f FakeNower) {
	f = s.tempNow
	return
}

// NewID returns the fake NewID object on this setup.
func (s *MockModelSetup) NewID() (f FakeNewIDer) {
	f = s.tempNewID
	return
}

// updateNow the now function with a mocked one.
func (s *MockModelSetup) updateNow() {
	now = s.Now().GetFunction()
}

// updateNewID the NewID function with a mocked one.
func (s *MockModelSetup) updateNewID() {
	newID = s.NewID().GetFunction()
}

// fakeNow it's a type that enable mocking of function now.
type fakeNow struct {
	returnV         *time.Time
	mockModelSetupP *MockModelSetup
}

// FakeNower it's a function mocking object, needed for mock purposes.
type FakeNower interface {
	Returns(time.Time)
	GetFunction() func() time.Time
}

// newFakeNow returns a new FakeNow object.
func newFakeNow() (f *fakeNow) {
	var i time.Time
	f = &fakeNow{
		returnV: &i,
	}
	return
}

// Returns ensures a value to be returned on calls to now.
func (f *fakeNow) Returns(t time.Time) {
	*f.returnV = t
	f.mockModelSetupP.updateNow()
}

// getFunction creates a version of now that returns value demanded.
func (f *fakeNow) GetFunction() (fn func() time.Time) {
	fn = func() (t time.Time) {
		t = *f.returnV
		return
	}
	return
}

// fakeNewID it's a type that enable mocking of function NewID.
type fakeNewID struct {
	returnV         *bson.ObjectId
	mockModelSetupP *MockModelSetup
}

// FakeNewIDer it's a function mocking object, needed for mock purposes.
type FakeNewIDer interface {
	Returns(bson.ObjectId)
	GetFunction() func() bson.ObjectId
}

// newFakeNewID returns a new FakeNewID object.
func newFakeNewID() (f *fakeNewID) {
	var id = newID()
	f = &fakeNewID{
		returnV: &id,
	}
	return
}

// Returns ensures a value to be returned on calls to NewID.
func (f *fakeNewID) Returns(id bson.ObjectId) {
	*f.returnV = id
	f.mockModelSetupP.updateNewID()
}

// getFunction creates a version of NewID that returns value demanded.
func (f *fakeNewID) GetFunction() (fn func() bson.ObjectId) {
	fn = func() (id bson.ObjectId) {
		id = *f.returnV
		return
	}
	return
}
