// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ddspog/mongo (interfaces: Collectioner)

// Package mocks is a generated GoMock package.
package mocks

import (
	mongo "github.com/ddspog/mongo"
	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
	bson "gopkg.in/mgo.v2/bson"
	reflect "reflect"
)

// MockCollectioner is a mock of Collectioner interface
type MockCollectioner struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionerMockRecorder
}

// MockCollectionerMockRecorder is the mock recorder for MockCollectioner
type MockCollectionerMockRecorder struct {
	mock *MockCollectioner
}

// NewMockCollectioner creates a new mock instance
func NewMockCollectioner(ctrl *gomock.Controller) *MockCollectioner {
	mock := &MockCollectioner{ctrl: ctrl}
	mock.recorder = &MockCollectionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollectioner) EXPECT() *MockCollectionerMockRecorder {
	return m.recorder
}

// Bulk mocks base method
func (m *MockCollectioner) Bulk() *mgo_v2.Bulk {
	ret := m.ctrl.Call(m, "Bulk")
	ret0, _ := ret[0].(*mgo_v2.Bulk)
	return ret0
}

// Bulk indicates an expected call of Bulk
func (mr *MockCollectionerMockRecorder) Bulk() *gomock.Call {
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
func (mr *MockCollectionerMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCollectioner)(nil).Count))
}

// Create mocks base method
func (m *MockCollectioner) Create(arg0 *mgo_v2.CollectionInfo) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCollectionerMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCollectioner)(nil).Create), arg0)
}

// DropCollection mocks base method
func (m *MockCollectioner) DropCollection() error {
	ret := m.ctrl.Call(m, "DropCollection")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropCollection indicates an expected call of DropCollection
func (mr *MockCollectionerMockRecorder) DropCollection() *gomock.Call {
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
func (mr *MockCollectionerMockRecorder) DropIndex(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndex", reflect.TypeOf((*MockCollectioner)(nil).DropIndex), arg0...)
}

// DropIndexName mocks base method
func (m *MockCollectioner) DropIndexName(arg0 string) error {
	ret := m.ctrl.Call(m, "DropIndexName", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DropIndexName indicates an expected call of DropIndexName
func (mr *MockCollectionerMockRecorder) DropIndexName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIndexName", reflect.TypeOf((*MockCollectioner)(nil).DropIndexName), arg0)
}

// EnsureIndex mocks base method
func (m *MockCollectioner) EnsureIndex(arg0 mgo_v2.Index) error {
	ret := m.ctrl.Call(m, "EnsureIndex", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureIndex indicates an expected call of EnsureIndex
func (mr *MockCollectionerMockRecorder) EnsureIndex(arg0 interface{}) *gomock.Call {
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
func (mr *MockCollectionerMockRecorder) EnsureIndexKey(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureIndexKey", reflect.TypeOf((*MockCollectioner)(nil).EnsureIndexKey), arg0...)
}

// Find mocks base method
func (m *MockCollectioner) Find(arg0 interface{}) mongo.Querier {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(mongo.Querier)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCollectionerMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollectioner)(nil).Find), arg0)
}

// FindId mocks base method
func (m *MockCollectioner) FindId(arg0 interface{}) mongo.Querier {
	ret := m.ctrl.Call(m, "FindId", arg0)
	ret0, _ := ret[0].(mongo.Querier)
	return ret0
}

// FindId indicates an expected call of FindId
func (mr *MockCollectionerMockRecorder) FindId(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockCollectioner)(nil).FindId), arg0)
}

// Indexes mocks base method
func (m *MockCollectioner) Indexes() ([]mgo_v2.Index, error) {
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].([]mgo_v2.Index)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Indexes indicates an expected call of Indexes
func (mr *MockCollectionerMockRecorder) Indexes() *gomock.Call {
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
func (mr *MockCollectionerMockRecorder) Insert(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollectioner)(nil).Insert), arg0...)
}

// NewIter mocks base method
func (m *MockCollectioner) NewIter(arg0 *mgo_v2.Session, arg1 []bson.Raw, arg2 int64, arg3 error) *mgo_v2.Iter {
	ret := m.ctrl.Call(m, "NewIter", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

// NewIter indicates an expected call of NewIter
func (mr *MockCollectionerMockRecorder) NewIter(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewIter", reflect.TypeOf((*MockCollectioner)(nil).NewIter), arg0, arg1, arg2, arg3)
}

// Pipe mocks base method
func (m *MockCollectioner) Pipe(arg0 interface{}) *mgo_v2.Pipe {
	ret := m.ctrl.Call(m, "Pipe", arg0)
	ret0, _ := ret[0].(*mgo_v2.Pipe)
	return ret0
}

// Pipe indicates an expected call of Pipe
func (mr *MockCollectionerMockRecorder) Pipe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipe", reflect.TypeOf((*MockCollectioner)(nil).Pipe), arg0)
}

// Remove mocks base method
func (m *MockCollectioner) Remove(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockCollectionerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockCollectioner)(nil).Remove), arg0)
}

// RemoveAll mocks base method
func (m *MockCollectioner) RemoveAll(arg0 interface{}) (*mongo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(*mongo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveAll indicates an expected call of RemoveAll
func (mr *MockCollectionerMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockCollectioner)(nil).RemoveAll), arg0)
}

// RemoveId mocks base method
func (m *MockCollectioner) RemoveId(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RemoveId", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveId indicates an expected call of RemoveId
func (mr *MockCollectionerMockRecorder) RemoveId(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveId", reflect.TypeOf((*MockCollectioner)(nil).RemoveId), arg0)
}

// Repair mocks base method
func (m *MockCollectioner) Repair() *mgo_v2.Iter {
	ret := m.ctrl.Call(m, "Repair")
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

// Repair indicates an expected call of Repair
func (mr *MockCollectionerMockRecorder) Repair() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repair", reflect.TypeOf((*MockCollectioner)(nil).Repair))
}

// Update mocks base method
func (m *MockCollectioner) Update(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockCollectionerMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCollectioner)(nil).Update), arg0, arg1)
}

// UpdateAll mocks base method
func (m *MockCollectioner) UpdateAll(arg0, arg1 interface{}) (*mongo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpdateAll", arg0, arg1)
	ret0, _ := ret[0].(*mongo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAll indicates an expected call of UpdateAll
func (mr *MockCollectionerMockRecorder) UpdateAll(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAll", reflect.TypeOf((*MockCollectioner)(nil).UpdateAll), arg0, arg1)
}

// UpdateId mocks base method
func (m *MockCollectioner) UpdateId(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "UpdateId", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateId indicates an expected call of UpdateId
func (mr *MockCollectionerMockRecorder) UpdateId(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateId", reflect.TypeOf((*MockCollectioner)(nil).UpdateId), arg0, arg1)
}

// Upsert mocks base method
func (m *MockCollectioner) Upsert(arg0, arg1 interface{}) (*mongo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(*mongo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert
func (mr *MockCollectionerMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockCollectioner)(nil).Upsert), arg0, arg1)
}

// UpsertId mocks base method
func (m *MockCollectioner) UpsertId(arg0, arg1 interface{}) (*mongo.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpsertId", arg0, arg1)
	ret0, _ := ret[0].(*mongo.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertId indicates an expected call of UpsertId
func (mr *MockCollectionerMockRecorder) UpsertId(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertId", reflect.TypeOf((*MockCollectioner)(nil).UpsertId), arg0, arg1)
}

// With mocks base method
func (m *MockCollectioner) With(arg0 *mgo_v2.Session) mongo.Collectioner {
	ret := m.ctrl.Call(m, "With", arg0)
	ret0, _ := ret[0].(mongo.Collectioner)
	return ret0
}

// With indicates an expected call of With
func (mr *MockCollectionerMockRecorder) With(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockCollectioner)(nil).With), arg0)
}
