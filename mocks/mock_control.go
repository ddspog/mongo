package mocks

import (
	"reflect"

	"github.com/ddspog/mongo/elements"
	"github.com/golang/mock/gomock"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockController) Connect() error {
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockControllerMockRecorder) Connect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockController)(nil).Connect))
}

// ConsumeDatabaseOnSession mocks base method
func (m *MockController) ConsumeDatabaseOnSession(arg0 func(elements.Databaser)) {
	m.ctrl.Call(m, "ConsumeDatabaseOnSession", arg0)
}

// ConsumeDatabaseOnSession indicates an expected call of ConsumeDatabaseOnSession
func (mr *MockControllerMockRecorder) ConsumeDatabaseOnSession(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeDatabaseOnSession", reflect.TypeOf((*MockController)(nil).ConsumeDatabaseOnSession), arg0)
}

// Mongo mocks base method
func (m *MockController) Mongo() *elements.DialInfo {
	ret := m.ctrl.Call(m, "Mongo")
	ret0, _ := ret[0].(*elements.DialInfo)
	return ret0
}

// Mongo indicates an expected call of Mongo
func (mr *MockControllerMockRecorder) Mongo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mongo", reflect.TypeOf((*MockController)(nil).Mongo))
}

// Session mocks base method
func (m *MockController) Session() elements.Sessioner {
	ret := m.ctrl.Call(m, "Session")
	ret0, _ := ret[0].(elements.Sessioner)
	return ret0
}

// Session indicates an expected call of Session
func (mr *MockControllerMockRecorder) Session() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockController)(nil).Session))
}
