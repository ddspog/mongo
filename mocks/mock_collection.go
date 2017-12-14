package mocks

import (
	"reflect"

	"github.com/ddspog/mongo/elements"
	"github.com/golang/mock/gomock"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MockCollectioner is a mock of Collectioner interface
type MockCollectioner struct {
	ctrl     *gomock.Controller
	recorder *mockCollectionerMockRecord
}

// mockCollectionerMockRecord is the mock recorder for mockCollectioner
type mockCollectionerMockRecord struct {
	mock *MockCollectioner
}

// MockCollectionerMockRecorder is a recorder used for mocking purposes.
// It's needed for the EXPECT method to work.
type MockCollectionerMockRecorder interface {
	Bulk() *gomock.Call
	Count() *gomock.Call
	Create(interface{}) *gomock.Call
	DropCollection() *gomock.Call
	DropIndex(...interface{}) *gomock.Call
	DropIndexName(interface{}) *gomock.Call
	EnsureIndex(interface{}) *gomock.Call
	EnsureIndexKey(...interface{}) *gomock.Call
	Find(interface{}) *gomock.Call
	FindID(interface{}) *gomock.Call
	Indexes() *gomock.Call
	Insert(...interface{}) *gomock.Call
	NewIter(interface{}, interface{}, interface{}, interface{}) *gomock.Call
	Pipe(interface{}) *gomock.Call
	Remove(interface{}) *gomock.Call
	RemoveAll(interface{}) *gomock.Call
	RemoveID(interface{}) *gomock.Call
	Repair() *gomock.Call
	Update(interface{}, interface{}) *gomock.Call
	UpdateAll(interface{}, interface{}) *gomock.Call
	UpdateID(interface{}, interface{}) *gomock.Call
	Upsert(interface{}, interface{}) *gomock.Call
	UpsertID(interface{}, interface{}) *gomock.Call
	With(interface{}) *gomock.Call
}

// NewMockCollectioner creates a new mock instance
func NewMockCollectioner(ctrl *gomock.Controller) *MockCollectioner {
	mock := &MockCollectioner{ctrl: ctrl}
	mock.recorder = &mockCollectionerMockRecord{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollectioner) EXPECT() MockCollectionerMockRecorder {
	return m.recorder
}

// Bulk mocks base method
func (m *MockCollectioner) Bulk() *mgo.Bulk {
	ret := m.ctrl.Call(m, "Bulk")
	ret0, _ := ret[0].(*mgo.Bulk)
	return ret0
}

// Bulk indicates an expected call of Bulk
func (mr *mockCollectionerMockRecord) Bulk() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bulk", reflect.TypeOf((*MockCollectioner)(nil).Bulk))
}

// Count mocks base method
func (m *MockCollectioner) Count() (int, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *mockCollectionerMockRecord) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCollectioner)(nil).Count))
}

// Create mocks base method
func (m *MockCollectioner) Create(arg0 *mgo.CollectionInfo) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *mockCollectionerMockRecord) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCollectioner)(nil).Create), arg0)
}

// DropCollection mocks base method
func (m *MockCollectioner) DropCollection() error {
	ret := m.ctrl.Call(m, "DropCollection")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCollection indicates an expected call of DropCollection
func (mr *mockCollectionerMockRecord) DropCollection() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCollection", reflect.TypeOf((*MockCollectioner)(nil).DropCollection))
}

// DropIndex mocks base method
func (m *MockCollectioner) DropIndex(arg0 ...string) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropIndex", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndex indicates an expected call of DropIndex
func (mr *mockCollectionerMockRecord) DropIndex(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndex", reflect.TypeOf((*MockCollectioner)(nil).DropIndex), arg0...)
}

// DropIndexName mocks base method
func (m *MockCollectioner) DropIndexName(arg0 string) error {
	ret := m.ctrl.Call(m, "DropIndexName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndexName indicates an expected call of DropIndexName
func (mr *mockCollectionerMockRecord) DropIndexName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndexName", reflect.TypeOf((*MockCollectioner)(nil).DropIndexName), arg0)
}

// EnsureIndex mocks base method
func (m *MockCollectioner) EnsureIndex(arg0 mgo.Index) error {
	ret := m.ctrl.Call(m, "EnsureIndex", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndex indicates an expected call of EnsureIndex
func (mr *mockCollectionerMockRecord) EnsureIndex(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndex", reflect.TypeOf((*MockCollectioner)(nil).EnsureIndex), arg0)
}

// EnsureIndexKey mocks base method
func (m *MockCollectioner) EnsureIndexKey(arg0 ...string) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnsureIndexKey", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndexKey indicates an expected call of EnsureIndexKey
func (mr *mockCollectionerMockRecord) EnsureIndexKey(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndexKey", reflect.TypeOf((*MockCollectioner)(nil).EnsureIndexKey), arg0...)
}

// Find mocks base method
func (m *MockCollectioner) Find(arg0 interface{}) elements.Querier {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(elements.Querier)
	return ret0
}

// Find indicates an expected call of Find
func (mr *mockCollectionerMockRecord) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollectioner)(nil).Find), arg0)
}

// FindID mocks base method
func (m *MockCollectioner) FindID(arg0 interface{}) elements.Querier {
	ret := m.ctrl.Call(m, "FindID", arg0)
	ret0, _ := ret[0].(elements.Querier)
	return ret0
}

// FindID indicates an expected call of FindID
func (mr *mockCollectionerMockRecord) FindID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindID", reflect.TypeOf((*MockCollectioner)(nil).FindID), arg0)
}

// Indexes mocks base method
func (m *MockCollectioner) Indexes() ([]mgo.Index, error) {
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].([]mgo.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Indexes indicates an expected call of Indexes
func (mr *mockCollectionerMockRecord) Indexes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockCollectioner)(nil).Indexes))
}

// Insert mocks base method
func (m *MockCollectioner) Insert(arg0 ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *mockCollectionerMockRecord) Insert(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollectioner)(nil).Insert), arg0...)
}

// NewIter mocks base method
func (m *MockCollectioner) NewIter(arg0 *mgo.Session, arg1 []bson.Raw, arg2 int64, arg3 error) *mgo.Iter {
	ret := m.ctrl.Call(m, "NewIter", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mgo.Iter)
	return ret0
}

// NewIter indicates an expected call of NewIter
func (mr *mockCollectionerMockRecord) NewIter(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIter", reflect.TypeOf((*MockCollectioner)(nil).NewIter), arg0, arg1, arg2, arg3)
}

// Pipe mocks base method
func (m *MockCollectioner) Pipe(arg0 interface{}) *mgo.Pipe {
	ret := m.ctrl.Call(m, "Pipe", arg0)
	ret0, _ := ret[0].(*mgo.Pipe)
	return ret0
}

// Pipe indicates an expected call of Pipe
func (mr *mockCollectionerMockRecord) Pipe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockCollectioner)(nil).Pipe), arg0)
}

// Remove mocks base method
func (m *MockCollectioner) Remove(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *mockCollectionerMockRecord) Remove(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCollectioner)(nil).Remove), arg0)
}

// RemoveAll mocks base method
func (m *MockCollectioner) RemoveAll(arg0 interface{}) (*elements.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(*elements.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *mockCollectionerMockRecord) RemoveAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockCollectioner)(nil).RemoveAll), arg0)
}

// RemoveID mocks base method
func (m *MockCollectioner) RemoveID(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RemoveID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveID indicates an expected call of RemoveID
func (mr *mockCollectionerMockRecord) RemoveID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveID", reflect.TypeOf((*MockCollectioner)(nil).RemoveID), arg0)
}

// Repair mocks base method
func (m *MockCollectioner) Repair() *mgo.Iter {
	ret := m.ctrl.Call(m, "Repair")
	ret0, _ := ret[0].(*mgo.Iter)
	return ret0
}

// Repair indicates an expected call of Repair
func (mr *mockCollectionerMockRecord) Repair() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repair", reflect.TypeOf((*MockCollectioner)(nil).Repair))
}

// Update mocks base method
func (m *MockCollectioner) Update(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *mockCollectionerMockRecord) Update(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCollectioner)(nil).Update), arg0, arg1)
}

// UpdateAll mocks base method
func (m *MockCollectioner) UpdateAll(arg0, arg1 interface{}) (*elements.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpdateAll", arg0, arg1)
	ret0, _ := ret[0].(*elements.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *mockCollectionerMockRecord) UpdateAll(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockCollectioner)(nil).UpdateAll), arg0, arg1)
}

// UpdateID mocks base method
func (m *MockCollectioner) UpdateID(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "UpdateID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateID indicates an expected call of UpdateID
func (mr *mockCollectionerMockRecord) UpdateID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateID", reflect.TypeOf((*MockCollectioner)(nil).UpdateID), arg0, arg1)
}

// Upsert mocks base method
func (m *MockCollectioner) Upsert(arg0, arg1 interface{}) (*elements.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(*elements.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *mockCollectionerMockRecord) Upsert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockCollectioner)(nil).Upsert), arg0, arg1)
}

// UpsertID mocks base method
func (m *MockCollectioner) UpsertID(arg0, arg1 interface{}) (*elements.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpsertID", arg0, arg1)
	ret0, _ := ret[0].(*elements.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertID indicates an expected call of UpsertID
func (mr *mockCollectionerMockRecord) UpsertID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertID", reflect.TypeOf((*MockCollectioner)(nil).UpsertID), arg0, arg1)
}

// With mocks base method
func (m *MockCollectioner) With(arg0 *mgo.Session) elements.Collectioner {
	ret := m.ctrl.Call(m, "With", arg0)
	ret0, _ := ret[0].(elements.Collectioner)
	return ret0
}

// With indicates an expected call of With
func (mr *mockCollectionerMockRecord) With(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockCollectioner)(nil).With), arg0)
}
