package mongo

import (
	"errors"
	"testing"

	"github.com/ddspog/mongo/model"
	"github.com/golang/mock/gomock"
)

// mockMGOSetup it's a setup type for configurating mocking of mgo
// package.
type mockMGOSetup struct {
	mockController *gomock.Controller
}

// NewMockMGOSetup returns a new mockMGOSetup, already configuring mock
// environment for the mgo classes mocked with gomock. It requires a
// test environment to be running.
func NewMockMGOSetup(t *testing.T) (s *mockMGOSetup, err error) {
	if t.Name() == "" {
		err = errors.New("Run only on test environment.")
	} else {
		s = &mockMGOSetup{
			mockController: gomock.NewController(t),
		}
	}
	return
}

// Finish restore mocking classes to original behavior.
func (s *mockMGOSetup) Finish() {
	s.controller().Finish()
}

// controller gets gomock controller.
func (s *mockMGOSetup) controller() (c *gomock.Controller) {
	c = s.mockController
	return
}

// DatabaseMock create a Database mock that expect an C to return the
// Collectioner mocked from function.
func (s *mockMGOSetup) DatabaseMock(n string, f func(*MockCollectioner)) (mdb *MockDatabaser) {
	mdb = newMockDatabaser(s.controller())
	mcl := newEmbeddedMockCollectioner(s.controller())
	f(mcl)
	mdb.EXPECT().C(n).AnyTimes().Return(mcl)
	return
}

// CollectionMock create a new Collection mock.
func (s *mockMGOSetup) CollectionMock() (mcl *MockCollectioner) {
	mcl = newEmbeddedMockCollectioner(s.controller())
	return
}

// MockCollectioner redirect the naming to this package, reducing
// importing. This new embedded class extends the mocked class
// to enable using of more helpful functions.
type MockCollectioner struct {
	*mockCollectioner
	mockController *gomock.Controller
}

// newEmbeddedMockCollectioner creates a new MockCollectioner embedded type.
func newEmbeddedMockCollectioner(c *gomock.Controller) (m *MockCollectioner) {
	m = &MockCollectioner{
		mockCollectioner: newMockCollectioner(c),
		mockController:   c,
	}
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
func (m *MockCollectioner) ExpectFindReturn(ret model.Documenter) {
	mqr := newMockQuerier(m.controller())
	mqr.EXPECT().One(gomock.Any()).Return(nil).Do(func(d model.Documenter) {
		d.SetId(ret.Id())
		d.SetCreatedOn(ret.CreatedOn())
	})
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectFindFail make a Collectioner expects an Find to return an
// error for whatever reason.
func (m *MockCollectioner) ExpectFindFail(mes string) {
	mqr := newMockQuerier(m.controller())
	mqr.EXPECT().One(gomock.Any()).Return(errors.New(mes))
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectFindAllReturn make a Collectioner expects an FindAll to return
// defined documents.
func (m *MockCollectioner) ExpectFindAllReturn(ret []model.Documenter) {
	mqr := newMockQuerier(m.controller())
	mqr.EXPECT().All(gomock.Any()).Return(nil).Do(func(da *[]model.Documenter) {
		*da = ret
	})
	m.EXPECT().Find(gomock.Any()).Return(mqr)
}

// ExpectFindAllFail make a Collectioner expects an FindAll to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectFindAllFail(mes string) {
	mqr := newMockQuerier(m.controller())
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

// ExpectRemoveIdReturn make a Collectioner expects an RemoveId to return
// no error and do nothing.
func (m *MockCollectioner) ExpectRemoveIdReturn() {
	m.EXPECT().RemoveId(gomock.Any()).Return(nil)
}

// ExpectRemoveIdFail make a Collectioner expects an RemoveId to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectRemoveIdFail(mes string) {
	m.EXPECT().RemoveId(gomock.Any()).Return(errors.New(mes))
}

// ExpectRemoveAllReturn make a Collectioner expects an RemoveAll to
// return no error and do nothing.
func (m *MockCollectioner) ExpectRemoveAllReturn(ret *ChangeInfo) {
	m.EXPECT().RemoveAll(gomock.Any()).Return(ret, nil)
}

// ExpectRemoveAllFail make a Collectioner expects an RemoveAll to
// return an error for whatever reason.
func (m *MockCollectioner) ExpectRemoveAllFail(mes string) {
	m.EXPECT().RemoveAll(gomock.Any()).Return(nil, errors.New(mes))
}

// ExpectUpdateIdReturn make a Collectioner expects an UpdateId to return
// no error and do nothing.
func (m *MockCollectioner) ExpectUpdateIdReturn() {
	m.EXPECT().UpdateId(gomock.Any(), gomock.Any()).Return(nil)
}

// ExpectUpdateIdFail make a Collectioner expects an UpdateId to return
// an error for whatever reason.
func (m *MockCollectioner) ExpectUpdateIdFail(mes string) {
	m.EXPECT().UpdateId(gomock.Any(), gomock.Any()).Return(errors.New(mes))
}

// controller gets gomock controller.
func (m *MockCollectioner) controller() (c *gomock.Controller) {
	c = m.mockController
	return
}
