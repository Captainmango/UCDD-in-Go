// Code generated by MockGen. DO NOT EDIT.
// Source: ./repositories/UsersRepository.go

// Package mock_repositories is a generated GoMock package.
package mocks

import (
	domain "main/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUsersRepository is a mock of IUsersRepository interface.
type MockIUsersRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUsersRepositoryMockRecorder
}

// MockIUsersRepositoryMockRecorder is the mock recorder for MockIUsersRepository.
type MockIUsersRepositoryMockRecorder struct {
	mock *MockIUsersRepository
}

// NewMockIUsersRepository creates a new mock instance.
func NewMockIUsersRepository(ctrl *gomock.Controller) *MockIUsersRepository {
	mock := &MockIUsersRepository{ctrl: ctrl}
	mock.recorder = &MockIUsersRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUsersRepository) EXPECT() *MockIUsersRepositoryMockRecorder {
	return m.recorder
}

// GetAllUsers mocks base method.
func (m *MockIUsersRepository) GetAllUsers() []domain.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]domain.User)
	return ret0
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUsersRepositoryMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUsersRepository)(nil).GetAllUsers))
}

// GetUserById mocks base method.
func (m *MockIUsersRepository) GetUserById(id int32) (domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockIUsersRepositoryMockRecorder) GetUserById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockIUsersRepository)(nil).GetUserById), id)
}
