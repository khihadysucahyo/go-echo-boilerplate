// Code generated by MockGen. DO NOT EDIT.
// Source: repositories/user_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/khihadysucahyo/go-echo-boilerplate/models"
)

// MockUserRepositoryQ is a mock of UserRepositoryQ interface.
type MockUserRepositoryQ struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryQMockRecorder
}

// MockUserRepositoryQMockRecorder is the mock recorder for MockUserRepositoryQ.
type MockUserRepositoryQMockRecorder struct {
	mock *MockUserRepositoryQ
}

// NewMockUserRepositoryQ creates a new mock instance.
func NewMockUserRepositoryQ(ctrl *gomock.Controller) *MockUserRepositoryQ {
	mock := &MockUserRepositoryQ{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryQMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepositoryQ) EXPECT() *MockUserRepositoryQMockRecorder {
	return m.recorder
}

// GetUserByEmail mocks base method.
func (m *MockUserRepositoryQ) GetUserByEmail(user *models.User, email string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetUserByEmail", user, email)
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryQMockRecorder) GetUserByEmail(user, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepositoryQ)(nil).GetUserByEmail), user, email)
}
