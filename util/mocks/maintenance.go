// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/maintenance (interfaces: Maintenance)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockMaintenance is a mock of Maintenance interface
type MockMaintenance struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceMockRecorder
}

// MockMaintenanceMockRecorder is the mock recorder for MockMaintenance
type MockMaintenanceMockRecorder struct {
	mock *MockMaintenance
}

// NewMockMaintenance creates a new mock instance
func NewMockMaintenance(ctrl *gomock.Controller) *MockMaintenance {
	mock := &MockMaintenance{ctrl: ctrl}
	mock.recorder = &MockMaintenanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMaintenance) EXPECT() *MockMaintenanceMockRecorder {
	return m.recorder
}

// End mocks base method
func (m *MockMaintenance) End() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "End")
	ret0, _ := ret[0].(error)
	return ret0
}

// End indicates an expected call of End
func (mr *MockMaintenanceMockRecorder) End() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "End", reflect.TypeOf((*MockMaintenance)(nil).End))
}

// StartControlPlane mocks base method
func (m *MockMaintenance) StartControlPlane(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartControlPlane", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartControlPlane indicates an expected call of StartControlPlane
func (mr *MockMaintenanceMockRecorder) StartControlPlane(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartControlPlane", reflect.TypeOf((*MockMaintenance)(nil).StartControlPlane), arg0)
}

// StartWorker mocks base method
func (m *MockMaintenance) StartWorker(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartWorker", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartWorker indicates an expected call of StartWorker
func (mr *MockMaintenanceMockRecorder) StartWorker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorker", reflect.TypeOf((*MockMaintenance)(nil).StartWorker), arg0)
}
