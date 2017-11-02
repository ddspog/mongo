package model

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

// mockSetup it's a setup type for configuring functions mocking.
type mockModelSetup struct {
	originalNowInMilli func() int64
	tempNowInMilli     *fakeNowInMilli
	tempNewId          *fakeNewId
}

// NewMockModelSetup returns a new mockModelSetup, already configuring
// a fakeNowInMilli function, and a fakeNewId function  on the setup.
// It requires a test environment to be running.
func NewMockModelSetup(t *testing.T) (s *mockModelSetup, err error) {
	if t.Name() == "" {
		err = fmt.Errorf("Run only on test environment.")
	} else {
		s = &mockModelSetup{
			originalNowInMilli: nowInMilli,
			tempNowInMilli:     newFakeNowInMilli(),
			tempNewId:          newFakeNewId(),
		}
		s.tempNowInMilli.mockModelSetupP = s
		s.tempNewId.mockModelSetupP = s
	}
	return
}

// Finish restore the functions mocked to the original ones.
func (s *mockModelSetup) Finish() {
	nowInMilli = s.originalNowInMilli
}

// NowInMilli returns the fake NowInMilli object on this setup.
func (s *mockModelSetup) NowInMilli() (f *fakeNowInMilli) {
	f = s.tempNowInMilli
	return
}

// NewId returns the fake NewId object on this setup.
func (s *mockModelSetup) NewId() (f *fakeNewId) {
	f = s.tempNewId
	return
}

// updateNowInMilli the nowInMilli function with a mocked one.
func (s *mockModelSetup) updateNowInMilli() {
	nowInMilli = s.NowInMilli().getFunction()
}

// updateNewId the newId function with a mocked one.
func (s *mockModelSetup) updateNewId() {
	newId = s.NewId().getFunction()
}

// fakeNowInMilli it's a type that enable mocking of function nowInMilli.
type fakeNowInMilli struct {
	returnV         *int64
	mockModelSetupP *mockModelSetup
}

// newFakeNowInMilli returns a new fakeNowInMilli object.
func newFakeNowInMilli() (f *fakeNowInMilli) {
	var i int64 = 0
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
func (f *fakeNowInMilli) getFunction() (fn func() int64) {
	fn = func() (t int64) {
		t = *f.returnV
		return
	}
	return
}

// fakeNewId it's a type that enable mocking of function newId.
type fakeNewId struct {
	returnV         *bson.ObjectId
	mockModelSetupP *mockModelSetup
}

// newFakeNewId returns a new fakeNewId object.
func newFakeNewId() (f *fakeNewId) {
	var id bson.ObjectId = newId()
	f = &fakeNewId{
		returnV: &id,
	}
	return
}

// Returns ensures a value to be returned on calls to newId.
func (f *fakeNewId) Returns(id bson.ObjectId) {
	*f.returnV = id
	f.mockModelSetupP.updateNewId()
}

// getFunction creates a version of newId that returns value demanded.
func (f *fakeNewId) getFunction() (fn func() bson.ObjectId) {
	fn = func() (id bson.ObjectId) {
		id = *f.returnV
		return
	}
	return
}
