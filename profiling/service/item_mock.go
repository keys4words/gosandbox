// Code generated by MockGen. DO NOT EDIT.
// Source: service/item.go

// Package service is a generated GoMock package.
package service

import (
	entity "profiling/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockItemRepoInterface is a mock of ItemRepoInterface interface.
type MockItemRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepoInterfaceMockRecorder
}

// MockItemRepoInterfaceMockRecorder is the mock recorder for MockItemRepoInterface.
type MockItemRepoInterfaceMockRecorder struct {
	mock *MockItemRepoInterface
}

// NewMockItemRepoInterface creates a new mock instance.
func NewMockItemRepoInterface(ctrl *gomock.Controller) *MockItemRepoInterface {
	mock := &MockItemRepoInterface{ctrl: ctrl}
	mock.recorder = &MockItemRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepoInterface) EXPECT() *MockItemRepoInterfaceMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockItemRepoInterface) GetByID(id int) (entity.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(entity.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockItemRepoInterfaceMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockItemRepoInterface)(nil).GetByID), id)
}
