package mongo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
)

// MockPCollectioner is a mock of Collectioner interface
type MockPCollectioner struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionerMockRecorder
}

// MockCollectionerMockRecorder is the mock recorder for mockCollectioner
type MockCollectionerMockRecorder struct {
	mock *MockPCollectioner
}

// newMockPCollectioner creates a new mock instance
func newMockPCollectioner(ctrl *gomock.Controller) *MockPCollectioner {
	mock := &MockPCollectioner{ctrl: ctrl}
	mock.recorder = &MockCollectionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPCollectioner) EXPECT() *MockCollectionerMockRecorder {
	return m.recorder
}

// Bulk mocks base method
func (m *MockPCollectioner) Bulk() *mgo_v2.Bulk {
	ret := m.ctrl.Call(m, "Bulk")
	ret0, _ := ret[0].(*mgo_v2.Bulk)
	return ret0
}

// Bulk indicates an expected call of Bulk
func (mr *MockCollectionerMockRecorder) Bulk() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bulk", reflect.TypeOf((*MockPCollectioner)(nil).Bulk))
}

// Count mocks base method
func (m *MockPCollectioner) Count() (int, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockCollectionerMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockPCollectioner)(nil).Count))
}

// Create mocks base method
func (m *MockPCollectioner) Create(arg0 *mgo_v2.CollectionInfo) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCollectionerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPCollectioner)(nil).Create), arg0)
}

// DropCollection mocks base method
func (m *MockPCollectioner) DropCollection() error {
	ret := m.ctrl.Call(m, "DropCollection")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCollection indicates an expected call of DropCollection
func (mr *MockCollectionerMockRecorder) DropCollection() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropCollection", reflect.TypeOf((*MockPCollectioner)(nil).DropCollection))
}

// DropIndex mocks base method
func (m *MockPCollectioner) DropIndex(arg0 ...string) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DropIndex", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndex indicates an expected call of DropIndex
func (mr *MockCollectionerMockRecorder) DropIndex(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndex", reflect.TypeOf((*MockPCollectioner)(nil).DropIndex), arg0...)
}

// DropIndexName mocks base method
func (m *MockPCollectioner) DropIndexName(arg0 string) error {
	ret := m.ctrl.Call(m, "DropIndexName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndexName indicates an expected call of DropIndexName
func (mr *MockCollectionerMockRecorder) DropIndexName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndexName", reflect.TypeOf((*MockPCollectioner)(nil).DropIndexName), arg0)
}

// EnsureIndex mocks base method
func (m *MockPCollectioner) EnsureIndex(arg0 mgo_v2.Index) error {
	ret := m.ctrl.Call(m, "EnsureIndex", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndex indicates an expected call of EnsureIndex
func (mr *MockCollectionerMockRecorder) EnsureIndex(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndex", reflect.TypeOf((*MockPCollectioner)(nil).EnsureIndex), arg0)
}

// EnsureIndexKey mocks base method
func (m *MockPCollectioner) EnsureIndexKey(arg0 ...string) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnsureIndexKey", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndexKey indicates an expected call of EnsureIndexKey
func (mr *MockCollectionerMockRecorder) EnsureIndexKey(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndexKey", reflect.TypeOf((*MockPCollectioner)(nil).EnsureIndexKey), arg0...)
}

// Find mocks base method
func (m *MockPCollectioner) Find(arg0 interface{}) Querier {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCollectionerMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockPCollectioner)(nil).Find), arg0)
}

// FindID mocks base method
func (m *MockPCollectioner) FindID(arg0 interface{}) Querier {
	ret := m.ctrl.Call(m, "FindID", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// FindID indicates an expected call of FindID
func (mr *MockCollectionerMockRecorder) FindID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindID", reflect.TypeOf((*MockPCollectioner)(nil).FindID), arg0)
}

// Indexes mocks base method
func (m *MockPCollectioner) Indexes() ([]mgo_v2.Index, error) {
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].([]mgo_v2.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Indexes indicates an expected call of Indexes
func (mr *MockCollectionerMockRecorder) Indexes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockPCollectioner)(nil).Indexes))
}

// Insert mocks base method
func (m *MockPCollectioner) Insert(arg0 ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockCollectionerMockRecorder) Insert(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPCollectioner)(nil).Insert), arg0...)
}

// NewIter mocks base method
func (m *MockPCollectioner) NewIter(arg0 *mgo_v2.Session, arg1 []bson.Raw, arg2 int64, arg3 error) *mgo_v2.Iter {
	ret := m.ctrl.Call(m, "NewIter", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

// NewIter indicates an expected call of NewIter
func (mr *MockCollectionerMockRecorder) NewIter(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIter", reflect.TypeOf((*MockPCollectioner)(nil).NewIter), arg0, arg1, arg2, arg3)
}

// Pipe mocks base method
func (m *MockPCollectioner) Pipe(arg0 interface{}) *mgo_v2.Pipe {
	ret := m.ctrl.Call(m, "Pipe", arg0)
	ret0, _ := ret[0].(*mgo_v2.Pipe)
	return ret0
}

// Pipe indicates an expected call of Pipe
func (mr *MockCollectionerMockRecorder) Pipe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockPCollectioner)(nil).Pipe), arg0)
}

// Remove mocks base method
func (m *MockPCollectioner) Remove(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCollectionerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPCollectioner)(nil).Remove), arg0)
}

// RemoveAll mocks base method
func (m *MockPCollectioner) RemoveAll(arg0 interface{}) (*ChangeInfo, error) {
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockCollectionerMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockPCollectioner)(nil).RemoveAll), arg0)
}

// RemoveID mocks base method
func (m *MockPCollectioner) RemoveID(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RemoveID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveID indicates an expected call of RemoveID
func (mr *MockCollectionerMockRecorder) RemoveID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveID", reflect.TypeOf((*MockPCollectioner)(nil).RemoveID), arg0)
}

// Repair mocks base method
func (m *MockPCollectioner) Repair() *mgo_v2.Iter {
	ret := m.ctrl.Call(m, "Repair")
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

// Repair indicates an expected call of Repair
func (mr *MockCollectionerMockRecorder) Repair() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repair", reflect.TypeOf((*MockPCollectioner)(nil).Repair))
}

// Update mocks base method
func (m *MockPCollectioner) Update(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCollectionerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPCollectioner)(nil).Update), arg0, arg1)
}

// UpdateAll mocks base method
func (m *MockPCollectioner) UpdateAll(arg0, arg1 interface{}) (*ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpdateAll", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockCollectionerMockRecorder) UpdateAll(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockPCollectioner)(nil).UpdateAll), arg0, arg1)
}

// UpdateID mocks base method
func (m *MockPCollectioner) UpdateID(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "UpdateID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateID indicates an expected call of UpdateID
func (mr *MockCollectionerMockRecorder) UpdateID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateID", reflect.TypeOf((*MockPCollectioner)(nil).UpdateID), arg0, arg1)
}

// Upsert mocks base method
func (m *MockPCollectioner) Upsert(arg0, arg1 interface{}) (*ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockCollectionerMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockPCollectioner)(nil).Upsert), arg0, arg1)
}

// UpsertID mocks base method
func (m *MockPCollectioner) UpsertID(arg0, arg1 interface{}) (*ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpsertID", arg0, arg1)
	ret0, _ := ret[0].(*ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertID indicates an expected call of UpsertID
func (mr *MockCollectionerMockRecorder) UpsertID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertID", reflect.TypeOf((*MockPCollectioner)(nil).UpsertID), arg0, arg1)
}

// With mocks base method
func (m *MockPCollectioner) With(arg0 *mgo_v2.Session) Collectioner {
	ret := m.ctrl.Call(m, "With", arg0)
	ret0, _ := ret[0].(Collectioner)
	return ret0
}

// With indicates an expected call of With
func (mr *MockCollectionerMockRecorder) With(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockPCollectioner)(nil).With), arg0)
}
