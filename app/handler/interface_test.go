// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rompi/tax-calc/app/svc (interfaces: Bill)

// Package handler is a generated GoMock package.
package handler

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/rompi/tax-calc/app/model"
	reflect "reflect"
)

// MockBill is a mock of Bill interface
type MockBill struct {
	ctrl     *gomock.Controller
	recorder *MockBillMockRecorder
}

// MockBillMockRecorder is the mock recorder for MockBill
type MockBillMockRecorder struct {
	mock *MockBill
}

// NewMockBill creates a new mock instance
func NewMockBill(ctrl *gomock.Controller) *MockBill {
	mock := &MockBill{ctrl: ctrl}
	mock.recorder = &MockBillMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBill) EXPECT() *MockBillMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockBill) Create(arg0 *model.Object) (*model.Object, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*model.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBillMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBill)(nil).Create), arg0)
}

// Read mocks base method
func (m *MockBill) Read() ([]*model.Object, int, error) {
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]*model.Object)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Read indicates an expected call of Read
func (mr *MockBillMockRecorder) Read() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockBill)(nil).Read))
}
