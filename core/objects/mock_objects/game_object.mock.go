// Code generated by MockGen. DO NOT EDIT.
// Source: ./game_object.go

// Package mock_objects is a generated GoMock package.
package mock_objects

import (
	gomock "github.com/golang/mock/gomock"
	objects "github.com/xosmig/roguelike/core/objects"
	reflect "reflect"
)

// MockGameObject is a mock of GameObject interface
type MockGameObject struct {
	ctrl     *gomock.Controller
	recorder *MockGameObjectMockRecorder
}

// MockGameObjectMockRecorder is the mock recorder for MockGameObject
type MockGameObjectMockRecorder struct {
	mock *MockGameObject
}

// NewMockGameObject creates a new mock instance
func NewMockGameObject(ctrl *gomock.Controller) *MockGameObject {
	mock := &MockGameObject{ctrl: ctrl}
	mock.recorder = &MockGameObjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGameObject) EXPECT() *MockGameObjectMockRecorder {
	return m.recorder
}

// Interact mocks base method
func (m *MockGameObject) Interact(other objects.GameObject) {
	m.ctrl.Call(m, "Interact", other)
}

// Interact indicates an expected call of Interact
func (mr *MockGameObjectMockRecorder) Interact(other interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Interact", reflect.TypeOf((*MockGameObject)(nil).Interact), other)
}

// Response mocks base method
func (m *MockGameObject) Response(other objects.GameObject) {
	m.ctrl.Call(m, "Response", other)
}

// Response indicates an expected call of Response
func (mr *MockGameObjectMockRecorder) Response(other interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockGameObject)(nil).Response), other)
}

// ModelName mocks base method
func (m *MockGameObject) ModelName() string {
	ret := m.ctrl.Call(m, "ModelName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ModelName indicates an expected call of ModelName
func (mr *MockGameObjectMockRecorder) ModelName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelName", reflect.TypeOf((*MockGameObject)(nil).ModelName))
}

// MockHasPosition is a mock of HasPosition interface
type MockHasPosition struct {
	ctrl     *gomock.Controller
	recorder *MockHasPositionMockRecorder
}

// MockHasPositionMockRecorder is the mock recorder for MockHasPosition
type MockHasPositionMockRecorder struct {
	mock *MockHasPosition
}

// NewMockHasPosition creates a new mock instance
func NewMockHasPosition(ctrl *gomock.Controller) *MockHasPosition {
	mock := &MockHasPosition{ctrl: ctrl}
	mock.recorder = &MockHasPositionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHasPosition) EXPECT() *MockHasPositionMockRecorder {
	return m.recorder
}

// SetPosition mocks base method
func (m *MockHasPosition) SetPosition(pos objects.Location) {
	m.ctrl.Call(m, "SetPosition", pos)
}

// SetPosition indicates an expected call of SetPosition
func (mr *MockHasPositionMockRecorder) SetPosition(pos interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPosition", reflect.TypeOf((*MockHasPosition)(nil).SetPosition), pos)
}

// GetPosition mocks base method
func (m *MockHasPosition) GetPosition() objects.Location {
	ret := m.ctrl.Call(m, "GetPosition")
	ret0, _ := ret[0].(objects.Location)
	return ret0
}

// GetPosition indicates an expected call of GetPosition
func (mr *MockHasPositionMockRecorder) GetPosition() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosition", reflect.TypeOf((*MockHasPosition)(nil).GetPosition))
}