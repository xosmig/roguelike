// Code generated by MockGen. DO NOT EDIT.
// Source: ./game_map.go

// Package mock_gamemap is a generated GoMock package.
package mock_gamemap

import (
	gomock "github.com/golang/mock/gomock"
	gamemap "github.com/xosmig/roguelike/core/gamemap"
	geom "github.com/xosmig/roguelike/core/geom"
	reflect "reflect"
)

// MockGameMap is a mock of GameMap interface
type MockGameMap struct {
	ctrl     *gomock.Controller
	recorder *MockGameMapMockRecorder
}

// MockGameMapMockRecorder is the mock recorder for MockGameMap
type MockGameMapMockRecorder struct {
	mock *MockGameMap
}

// NewMockGameMap creates a new mock instance
func NewMockGameMap(ctrl *gomock.Controller) *MockGameMap {
	mock := &MockGameMap{ctrl: ctrl}
	mock.recorder = &MockGameMapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGameMap) EXPECT() *MockGameMapMockRecorder {
	return m.recorder
}

// GetHeight mocks base method
func (m *MockGameMap) GetHeight() int {
	ret := m.ctrl.Call(m, "GetHeight")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetHeight indicates an expected call of GetHeight
func (mr *MockGameMapMockRecorder) GetHeight() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeight", reflect.TypeOf((*MockGameMap)(nil).GetHeight))
}

// GetWidth mocks base method
func (m *MockGameMap) GetWidth() int {
	ret := m.ctrl.Call(m, "GetWidth")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetWidth indicates an expected call of GetWidth
func (mr *MockGameMapMockRecorder) GetWidth() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWidth", reflect.TypeOf((*MockGameMap)(nil).GetWidth))
}

// Get mocks base method
func (m *MockGameMap) Get(arg0 geom.Location) *gamemap.Cell {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*gamemap.Cell)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockGameMapMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockGameMap)(nil).Get), arg0)
}
