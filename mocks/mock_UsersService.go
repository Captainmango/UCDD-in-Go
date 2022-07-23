// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/UsersService.go

// Package mock_services is a generated GoMock package.
package mocks

import (
	domain "main/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUsersService is a mock of IUsersService interface.
type MockIUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockIUsersServiceMockRecorder
}

// MockIUsersServiceMockRecorder is the mock recorder for MockIUsersService.
type MockIUsersServiceMockRecorder struct {
	mock *MockIUsersService
}

// NewMockIUsersService creates a new mock instance.
func NewMockIUsersService(ctrl *gomock.Controller) *MockIUsersService {
	mock := &MockIUsersService{ctrl: ctrl}
	mock.recorder = &MockIUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUsersService) EXPECT() *MockIUsersServiceMockRecorder {
	return m.recorder
}

// GetAllUsers mocks base method.
func (m *MockIUsersService) GetAllUsers() []domain.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]domain.User)
	return ret0
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUsersServiceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUsersService)(nil).GetAllUsers))
}

// GetUserById mocks base method.
func (m *MockIUsersService) GetUserById(id int32) domain.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(domain.User)
	return ret0
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockIUsersServiceMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockIUsersService)(nil).GetUserById), id)
}
