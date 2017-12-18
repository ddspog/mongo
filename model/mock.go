package model

import (
	"fmt"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// MockModelSetup it's a setup type for configuring functions mocking.
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
