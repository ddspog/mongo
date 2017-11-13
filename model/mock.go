package model

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// MockModelSetup it's a setup type for configuring functions mocking.
type MockModelSetup struct {
	originalNowInMilli func() int64
	tempNowInMilli     *FakeNowInMilli
	tempNewID          *FakeNewID
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
			tempNowInMilli:     NewFakeNowInMilli(),
			tempNewID:          NewFakeNewID(),
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
func (s *MockModelSetup) NowInMilli() (f *FakeNowInMilli) {
	f = s.tempNowInMilli
	return
}

// NewID returns the fake NewID object on this setup.
func (s *MockModelSetup) NewID() (f *FakeNewID) {
	f = s.tempNewID
	return
}

// updateNowInMilli the nowInMilli function with a mocked one.
func (s *MockModelSetup) updateNowInMilli() {
	nowInMilli = s.NowInMilli().getFunction()
}

// updateNewID the NewID function with a mocked one.
func (s *MockModelSetup) updateNewID() {
	newID = s.NewID().getFunction()
}

// FakeNowInMilli it's a type that enable mocking of function nowInMilli.
type FakeNowInMilli struct {
	returnV         *int64
	mockModelSetupP *MockModelSetup
}

// NewFakeNowInMilli returns a new FakeNowInMilli object.
func NewFakeNowInMilli() (f *FakeNowInMilli) {
	var i = int64(0)
	f = &FakeNowInMilli{
		returnV: &i,
	}
	return
}

// Returns ensures a value to be returned on calls to nowInMilli.
func (f *FakeNowInMilli) Returns(t int64) {
	*f.returnV = t
	f.mockModelSetupP.updateNowInMilli()
}

// getFunction creates a version of nowInMilli that returns value demanded.
func (f *FakeNowInMilli) getFunction() (fn func() int64) {
	fn = func() (t int64) {
		t = *f.returnV
		return
	}
	return
}

// FakeNewID it's a type that enable mocking of function NewID.
type FakeNewID struct {
	returnV         *bson.ObjectId
	mockModelSetupP *MockModelSetup
}

// NewFakeNewID returns a new FakeNewID object.
func NewFakeNewID() (f *FakeNewID) {
	var id = newID()
	f = &FakeNewID{
		returnV: &id,
	}
	return
}

// Returns ensures a value to be returned on calls to NewID.
func (f *FakeNewID) Returns(id bson.ObjectId) {
	*f.returnV = id
	f.mockModelSetupP.updateNewID()
}

// getFunction creates a version of NewID that returns value demanded.
func (f *FakeNewID) getFunction() (fn func() bson.ObjectId) {
	fn = func() (id bson.ObjectId) {
		id = *f.returnV
		return
	}
	return
}
