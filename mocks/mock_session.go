package mocks

import (
	"reflect"

	"github.com/ddspog/mongo/elements"
	"github.com/globalsign/mgo"
	"github.com/golang/mock/gomock"
)

// MockSessioner is a mock of Sessioner interface
type MockSessioner struct {
	ctrl     *gomock.Controller
	recorder *MockSessionerMockRecorder
}

// MockSessionerMockRecorder is the mock recorder for MockSessioner
type MockSessionerMockRecorder struct {
	mock *MockSessioner
}

// NewMockSessioner creates a new mock instance
func NewMockSessioner(ctrl *gomock.Controller) *MockSessioner {
	mock := &MockSessioner{ctrl: ctrl}
	mock.recorder = &MockSessionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessioner) EXPECT() *MockSessionerMockRecorder {
	return m.recorder
}

// Clone mocks base method
func (m *MockSessioner) Clone() elements.Sessioner {
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(elements.Sessioner)
	return ret0
}

// Clone indicates an expected call of Clone
func (mr *MockSessionerMockRecorder) Clone() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockSessioner)(nil).Clone))
}

// Close mocks base method
func (m *MockSessioner) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSessionerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSessioner)(nil).Close))
}

// DB mocks base method
func (m *MockSessioner) DB(arg0 string) elements.Databaser {
	ret := m.ctrl.Call(m, "DB", arg0)
	ret0, _ := ret[0].(elements.Databaser)
	return ret0
}

// DB indicates an expected call of DB
func (mr *MockSessionerMockRecorder) DB(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockSessioner)(nil).DB), arg0)
}

// SetSafe mocks base method
func (m *MockSessioner) SetSafe(arg0 *mgo.Safe) {
	m.ctrl.Call(m, "SetSafe", arg0)
}

// SetSafe indicates an expected call of SetSafe
func (mr *MockSessionerMockRecorder) SetSafe(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSafe", reflect.TypeOf((*MockSessioner)(nil).SetSafe), arg0)
}
