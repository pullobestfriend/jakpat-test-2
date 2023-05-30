// Code generated by MockGen. DO NOT EDIT.
// Source: G:\code\jakpat-test-2\pkg\service\service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	entity "jakpat-test-2/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUser) CreateUser(input entity.Users) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserMockRecorder) CreateUser(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUser)(nil).CreateUser), input)
}

// GetUserByIdAndStatus mocks base method.
func (m *MockUser) GetUserByIdAndStatus(id, status int) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByIdAndStatus", id, status)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByIdAndStatus indicates an expected call of GetUserByIdAndStatus.
func (mr *MockUserMockRecorder) GetUserByIdAndStatus(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByIdAndStatus", reflect.TypeOf((*MockUser)(nil).GetUserByIdAndStatus), id, status)
}

// GetUserByNameAndPassword mocks base method.
func (m *MockUser) GetUserByNameAndPassword(name, password string) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByNameAndPassword", name, password)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByNameAndPassword indicates an expected call of GetUserByNameAndPassword.
func (mr *MockUserMockRecorder) GetUserByNameAndPassword(name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByNameAndPassword", reflect.TypeOf((*MockUser)(nil).GetUserByNameAndPassword), name, password)
}

// MockItem is a mock of Item interface.
type MockItem struct {
	ctrl     *gomock.Controller
	recorder *MockItemMockRecorder
}

// MockItemMockRecorder is the mock recorder for MockItem.
type MockItemMockRecorder struct {
	mock *MockItem
}

// NewMockItem creates a new mock instance.
func NewMockItem(ctrl *gomock.Controller) *MockItem {
	mock := &MockItem{ctrl: ctrl}
	mock.recorder = &MockItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItem) EXPECT() *MockItemMockRecorder {
	return m.recorder
}

// AddItem mocks base method.
func (m *MockItem) AddItem(input entity.Items) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddItem indicates an expected call of AddItem.
func (mr *MockItemMockRecorder) AddItem(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockItem)(nil).AddItem), input)
}

// CreateOrder mocks base method.
func (m *MockItem) CreateOrder(input entity.Oders) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockItemMockRecorder) CreateOrder(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockItem)(nil).CreateOrder), input)
}

// GetItemByIdAndStatus mocks base method.
func (m *MockItem) GetItemByIdAndStatus(id, status int) (entity.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemByIdAndStatus", id, status)
	ret0, _ := ret[0].(entity.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemByIdAndStatus indicates an expected call of GetItemByIdAndStatus.
func (mr *MockItemMockRecorder) GetItemByIdAndStatus(id, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemByIdAndStatus", reflect.TypeOf((*MockItem)(nil).GetItemByIdAndStatus), id, status)
}

// GetItemsBySellerIdAndStatus mocks base method.
func (m *MockItem) GetItemsBySellerIdAndStatus(sellerID, status int) ([]entity.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemsBySellerIdAndStatus", sellerID, status)
	ret0, _ := ret[0].([]entity.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemsBySellerIdAndStatus indicates an expected call of GetItemsBySellerIdAndStatus.
func (mr *MockItemMockRecorder) GetItemsBySellerIdAndStatus(sellerID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemsBySellerIdAndStatus", reflect.TypeOf((*MockItem)(nil).GetItemsBySellerIdAndStatus), sellerID, status)
}

// GetOrderById mocks base method.
func (m *MockItem) GetOrderById(id int) (entity.Oders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderById", id)
	ret0, _ := ret[0].(entity.Oders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderById indicates an expected call of GetOrderById.
func (mr *MockItemMockRecorder) GetOrderById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderById", reflect.TypeOf((*MockItem)(nil).GetOrderById), id)
}

// UpdateItemById mocks base method.
func (m *MockItem) UpdateItemById(id int, input entity.Items) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemById", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItemById indicates an expected call of UpdateItemById.
func (mr *MockItemMockRecorder) UpdateItemById(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemById", reflect.TypeOf((*MockItem)(nil).UpdateItemById), id, input)
}

// UpdateOrderById mocks base method.
func (m *MockItem) UpdateOrderById(id int, input entity.Oders) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderById", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderById indicates an expected call of UpdateOrderById.
func (mr *MockItemMockRecorder) UpdateOrderById(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderById", reflect.TypeOf((*MockItem)(nil).UpdateOrderById), id, input)
}
