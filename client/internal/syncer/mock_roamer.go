// Code generated by MockGen. DO NOT EDIT.
// Source: client/internal/syncer (interfaces: RoamerAccessor)

// Package syncer is a generated GoMock package.
package syncer

import (
	models "client/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRoamerAccessor is a mock of RoamerAccessor interface.
type MockRoamerAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockRoamerAccessorMockRecorder
}

// MockRoamerAccessorMockRecorder is the mock recorder for MockRoamerAccessor.
type MockRoamerAccessorMockRecorder struct {
	mock *MockRoamerAccessor
}

// NewMockRoamerAccessor creates a new mock instance.
func NewMockRoamerAccessor(ctrl *gomock.Controller) *MockRoamerAccessor {
	mock := &MockRoamerAccessor{ctrl: ctrl}
	mock.recorder = &MockRoamerAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoamerAccessor) EXPECT() *MockRoamerAccessorMockRecorder {
	return m.recorder
}

// SecretGet mocks base method.
func (m *MockRoamerAccessor) SecretGet(arg0 string) ([]models.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretGet", arg0)
	ret0, _ := ret[0].([]models.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecretGet indicates an expected call of SecretGet.
func (mr *MockRoamerAccessorMockRecorder) SecretGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretGet", reflect.TypeOf((*MockRoamerAccessor)(nil).SecretGet), arg0)
}

// SecretSet mocks base method.
func (m *MockRoamerAccessor) SecretSet(arg0 []models.Secret) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretSet", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretSet indicates an expected call of SecretSet.
func (mr *MockRoamerAccessorMockRecorder) SecretSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretSet", reflect.TypeOf((*MockRoamerAccessor)(nil).SecretSet), arg0)
}
