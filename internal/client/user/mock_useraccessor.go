// Code generated by MockGen. DO NOT EDIT.
// Source: secret_keeper/internal/client/user (interfaces: UserAccessor)

// Package user is a generated GoMock package.
package user

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserAccessor is a mock of UserAccessor interface.
type MockUserAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockUserAccessorMockRecorder
}

// MockUserAccessorMockRecorder is the mock recorder for MockUserAccessor.
type MockUserAccessorMockRecorder struct {
	mock *MockUserAccessor
}

// NewMockUserAccessor creates a new mock instance.
func NewMockUserAccessor(ctrl *gomock.Controller) *MockUserAccessor {
	mock := &MockUserAccessor{ctrl: ctrl}
	mock.recorder = &MockUserAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserAccessor) EXPECT() *MockUserAccessorMockRecorder {
	return m.recorder
}

// UserCreate mocks base method.
func (m *MockUserAccessor) UserCreate(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserCreate", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserCreate indicates an expected call of UserCreate.
func (mr *MockUserAccessorMockRecorder) UserCreate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserCreate", reflect.TypeOf((*MockUserAccessor)(nil).UserCreate), arg0, arg1, arg2)
}

// UserGet mocks base method.
func (m *MockUserAccessor) UserGet(arg0 context.Context, arg1 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserGet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UserGet indicates an expected call of UserGet.
func (mr *MockUserAccessorMockRecorder) UserGet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserGet", reflect.TypeOf((*MockUserAccessor)(nil).UserGet), arg0, arg1)
}
