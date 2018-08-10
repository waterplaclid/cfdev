// Code generated by MockGen. DO NOT EDIT.
// Source: code.cloudfoundry.org/cfdev/cmd/stop (interfaces: VpnKit)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVpnKit is a mock of VpnKit interface
type MockVpnKit struct {
	ctrl     *gomock.Controller
	recorder *MockVpnKitMockRecorder
}

// MockVpnKitMockRecorder is the mock recorder for MockVpnKit
type MockVpnKitMockRecorder struct {
	mock *MockVpnKit
}

// NewMockVpnKit creates a new mock instance
func NewMockVpnKit(ctrl *gomock.Controller) *MockVpnKit {
	mock := &MockVpnKit{ctrl: ctrl}
	mock.recorder = &MockVpnKitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVpnKit) EXPECT() *MockVpnKitMockRecorder {
	return m.recorder
}

// Destroy mocks base method
func (m *MockVpnKit) Destroy() error {
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockVpnKitMockRecorder) Destroy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockVpnKit)(nil).Destroy))
}

// Stop mocks base method
func (m *MockVpnKit) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockVpnKitMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockVpnKit)(nil).Stop))
}