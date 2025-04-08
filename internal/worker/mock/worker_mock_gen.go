// Code generated by MockGen. DO NOT EDIT.
// Source: internal/worker/worker.go
//
// Generated by this command:
//
//	mockgen -source=internal/worker/worker.go -destination=internal/worker/mock/worker_mock_gen.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	apigen "github.com/risingwavelabs/wavekit/internal/apigen"
	gomock "go.uber.org/mock/gomock"
)

// MockExecutorInterface is a mock of ExecutorInterface interface.
type MockExecutorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorInterfaceMockRecorder
	isgomock struct{}
}

// MockExecutorInterfaceMockRecorder is the mock recorder for MockExecutorInterface.
type MockExecutorInterfaceMockRecorder struct {
	mock *MockExecutorInterface
}

// NewMockExecutorInterface creates a new mock instance.
func NewMockExecutorInterface(ctrl *gomock.Controller) *MockExecutorInterface {
	mock := &MockExecutorInterface{ctrl: ctrl}
	mock.recorder = &MockExecutorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutorInterface) EXPECT() *MockExecutorInterfaceMockRecorder {
	return m.recorder
}

// ExecuteAutoBackup mocks base method.
func (m *MockExecutorInterface) ExecuteAutoBackup(ctx context.Context, spec apigen.TaskSpecAutoBackup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteAutoBackup", ctx, spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteAutoBackup indicates an expected call of ExecuteAutoBackup.
func (mr *MockExecutorInterfaceMockRecorder) ExecuteAutoBackup(ctx, spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAutoBackup", reflect.TypeOf((*MockExecutorInterface)(nil).ExecuteAutoBackup), ctx, spec)
}

// ExecuteAutoDiagnostic mocks base method.
func (m *MockExecutorInterface) ExecuteAutoDiagnostic(ctx context.Context, spec apigen.TaskSpecAutoDiagnostic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteAutoDiagnostic", ctx, spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteAutoDiagnostic indicates an expected call of ExecuteAutoDiagnostic.
func (mr *MockExecutorInterfaceMockRecorder) ExecuteAutoDiagnostic(ctx, spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteAutoDiagnostic", reflect.TypeOf((*MockExecutorInterface)(nil).ExecuteAutoDiagnostic), ctx, spec)
}

// ExecuteDeleteClusterDiagnostic mocks base method.
func (m *MockExecutorInterface) ExecuteDeleteClusterDiagnostic(ctx context.Context, spec apigen.TaskSpecDeleteClusterDiagnostic) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteDeleteClusterDiagnostic", ctx, spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteDeleteClusterDiagnostic indicates an expected call of ExecuteDeleteClusterDiagnostic.
func (mr *MockExecutorInterfaceMockRecorder) ExecuteDeleteClusterDiagnostic(ctx, spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteDeleteClusterDiagnostic", reflect.TypeOf((*MockExecutorInterface)(nil).ExecuteDeleteClusterDiagnostic), ctx, spec)
}

// ExecuteDeleteSnapshot mocks base method.
func (m *MockExecutorInterface) ExecuteDeleteSnapshot(ctx context.Context, spec apigen.TaskSpecDeleteSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteDeleteSnapshot", ctx, spec)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteDeleteSnapshot indicates an expected call of ExecuteDeleteSnapshot.
func (mr *MockExecutorInterfaceMockRecorder) ExecuteDeleteSnapshot(ctx, spec any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteDeleteSnapshot", reflect.TypeOf((*MockExecutorInterface)(nil).ExecuteDeleteSnapshot), ctx, spec)
}
