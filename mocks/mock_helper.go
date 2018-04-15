package mocks

import (
	"errors"
	"testing"

	"github.com/ddspog/mongo/elements"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/mock/gomock"
)

// MockMGOSetup it's a setup type for configuring mocking of mgo
// package.
type MockMGOSetup struct {
	mockController *gomock.Controller
}

// NewMockMGOSetup returns a new MockMGOSetup, already configuring mock
// environment for the mgo classes mocked with gomock. It requires a
// test environment to be running.
func NewMockMGOSetup(t *testing.T) (s *MockMGOSetup, err error) {
	if t.Name() == "" {
		err = errors.New("run only on test environment")
	} else {
		s = &MockMGOSetup{
			mockController: gomock.NewController(t),
		}
	}
	return
}

// Finish restore mocking classes to original behavior.
func (s *MockMGOSetup) Finish() {
	s.controller().Finish()
}

// controller gets gomock controller.
func (s *MockMGOSetup) controller() (c *gomock.Controller) {
	c = s.mockController
	return
}

// SessionMock create a Session mock that runs like normally, returning
// a Databaser mock.
func (s *MockMGOSetup) SessionMock(db string, mdb *MockDatabaser) (ms *MockSessioner) {
	ms = NewMockSessioner(s.controller())

	ms.EXPECT().SetSafe(gomock.Any()).AnyTimes()
	ms.EXPECT().Clone().AnyTimes().Return(ms)
	ms.EXPECT().Close().AnyTimes()
	ms.EXPECT().DB(db).AnyTimes().Return(mdb)

	return
}

// DatabaseMock create a Database mock that expect an C to return the
// Collectioner mocked from function.
func (s *MockMGOSetup) DatabaseMock(name string, f func(*MockCollectioner)) (mdb *MockDatabaser) {
	mdb = NewMockDatabaser(s.controller())
	mcl := NewMockCollectioner(s.controller())
	f(mcl)
	mdb.EXPECT().C(name).AnyTimes().Return(mcl)
	return
}

// CollectionMock create a new Collection mock.
func (s *MockMGOSetup) CollectionMock() (mcl *MockCollectioner) {
	mcl = NewMockCollectioner(s.controller())
	return
}

// ExpectCountReturn make a Collectioner expects an Count to return
// defined number n.
func (m *MockCollectioner) ExpectCountReturn(n int) {
	m.EXPECT().Count().Return(n, nil)
}

// ExpectCountFail make a Collectioner expects an Count to return
// an error with message m.
func (m *MockCollectioner) ExpectCountFail(mes string) {
	m.EXPECT().Count().Return(0, errors.New(mes))
}

// ExpectFindReturn make a Collectioner expects an Find to return
// defined document.
func (m *MockCollectioner) ExpectFindReturn(ret bson.M) {
	mqr := NewMockQuerier(m.controller())
	mqr.EXPECT().One(gomock.Any()).Return(nil).Do(func(d *interface{}) {
		*d = ret
	})
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectFindFail make a Collectioner expects an Find to return an
// error for whatever reason.
func (m *MockCollectioner) ExpectFindFail(mes string) {
	mqr := NewMockQuerier(m.controller())
	mqr.EXPECT().One(gomock.Any()).Return(errors.New(mes))
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectFindAllReturn make a Collectioner expects an FindAll to return
// defined documents.
func (m *MockCollectioner) ExpectFindAllReturn(ret []interface{}) {
	mqr := NewMockQuerier(m.controller())
	mqr.EXPECT().All(gomock.Any()).Return(nil).Do(func(da *[]interface{}) {
		*da = ret
	})
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectFindAllFail make a Collectioner expects an FindAll to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectFindAllFail(mes string) {
	mqr := NewMockQuerier(m.controller())
	mqr.EXPECT().All(gomock.Any()).Return(errors.New(mes))
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectInsertReturn make a Collectioner expects an Insert to return
// no error and do nothing.
func (m *MockCollectioner) ExpectInsertReturn() {
	m.EXPECT().Insert(gomock.Any()).Return(nil)
}

// ExpectInsertFail make a Collectioner expects an Insert to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectInsertFail(mes string) {
	m.EXPECT().Insert(gomock.Any()).Return(errors.New(mes))
}

// ExpectRemoveIDReturn make a Collectioner expects an RemoveId to return
// no error and do nothing.
func (m *MockCollectioner) ExpectRemoveIDReturn() {
	m.EXPECT().RemoveID(gomock.Any()).Return(nil)
}

// ExpectRemoveIDFail make a Collectioner expects an RemoveId to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectRemoveIDFail(mes string) {
	m.EXPECT().RemoveID(gomock.Any()).Return(errors.New(mes))
}

// ExpectRemoveAllReturn make a Collectioner expects an RemoveAll to
// return no error and do nothing.
func (m *MockCollectioner) ExpectRemoveAllReturn(ret *elements.ChangeInfo) {
	m.EXPECT().RemoveAll(gomock.Any()).Return(ret, nil)
}

// ExpectRemoveAllFail make a Collectioner expects an RemoveAll to
// return an error for whatever reason.
func (m *MockCollectioner) ExpectRemoveAllFail(mes string) {
	m.EXPECT().RemoveAll(gomock.Any()).Return(nil, errors.New(mes))
}

// ExpectUpdateIDReturn make a Collectioner expects an UpdateId to return
// no error and do nothing.
func (m *MockCollectioner) ExpectUpdateIDReturn() {
	m.EXPECT().UpdateID(gomock.Any(), gomock.Any()).Return(nil)
}

// ExpectUpdateIDFail make a Collectioner expects an UpdateId to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectUpdateIDFail(mes string) {
	m.EXPECT().UpdateID(gomock.Any(), gomock.Any()).Return(errors.New(mes))
}

// controller gets gomock controller.
func (m *MockCollectioner) controller() (c *gomock.Controller) {
	c = m.ctrl
	return
}
