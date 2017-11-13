package model

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// MockModelSetup it's a setup type for configuring functions mocking.
type MockModelSetup struct {
	originalNowInMilli func() int64
	tempNowInMilli     *fakeNowInMilli
	tempNewID          *fakeNewID
}

// NewMockModelSetup returns a new MockModelSetup, already configuring
// a FakeNowInMilli function, and a FakeNewID function  on the setup.
// It requires a test environment to be running.
func NewMockModelSetup(t *testing.T) (s *MockModelSetup, err error) {
	if t.Name() == "" {
		err = fmt.Errorf("Run only on test environment")
	} else {
		s = &MockModelSetup{
			originalNowInMilli: nowInMilli,
			tempNowInMilli:     newFakeNowInMilli(),
			tempNewID:          newFakeNewID(),
		}
		s.tempNowInMilli.mockModelSetupP = s
		s.tempNewID.mockModelSetupP = s
	}
	return
}

// Finish restore the functions mocked to the original ones.
func (s *MockModelSetup) Finish() {
	nowInMilli = s.originalNowInMilli
}

// NowInMilli returns the fake NowInMilli object on this setup.
func (s *MockModelSetup) NowInMilli() (f FakeNowInMillier) {
	f = s.tempNowInMilli
	return
}

// NewID returns the fake NewID object on this setup.
func (s *MockModelSetup) NewID() (f FakeNewIDer) {
	f = s.tempNewID
	return
}

// updateNowInMilli the nowInMilli function with a mocked one.
func (s *MockModelSetup) updateNowInMilli() {
	nowInMilli = s.NowInMilli().GetFunction()
}

// updateNewID the NewID function with a mocked one.
func (s *MockModelSetup) updateNewID() {
	newID = s.NewID().GetFunction()
}

// fakeNowInMilli it's a type that enable mocking of function nowInMilli.
type fakeNowInMilli struct {
	returnV         *int64
	mockModelSetupP *MockModelSetup
}

// FakeNowInMillier it's a function mocking object, needed for mock purposes.
type FakeNowInMillier interface {
	Returns(int64)
	GetFunction() func() int64
}

// newFakeNowInMilli returns a new FakeNowInMilli object.
func newFakeNowInMilli() (f *fakeNowInMilli) {
	var i int64
	f = &fakeNowInMilli{
		returnV: &i,
	}
	return
}

// Returns ensures a value to be returned on calls to nowInMilli.
func (f *fakeNowInMilli) Returns(t int64) {
	*f.returnV = t
	f.mockModelSetupP.updateNowInMilli()
}

// getFunction creates a version of nowInMilli that returns value demanded.
func (f *fakeNowInMilli) GetFunction() (fn func() int64) {
	fn = func() (t int64) {
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
