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

// MockTaskHandler is a mock of TaskHandler interface.
type MockTaskHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTaskHandlerMockRecorder
	isgomock struct{}
}

// MockTaskHandlerMockRecorder is the mock recorder for MockTaskHandler.
type MockTaskHandlerMockRecorder struct {
	mock *MockTaskHandler
}

// NewMockTaskHandler creates a new mock instance.
func NewMockTaskHandler(ctrl *gomock.Controller) *MockTaskHandler {
	mock := &MockTaskHandler{ctrl: ctrl}
	mock.recorder = &MockTaskHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskHandler) EXPECT() *MockTaskHandlerMockRecorder {
	return m.recorder
}

// HandleTask mocks base method.
func (m *MockTaskHandler) HandleTask(ctx context.Context, task apigen.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleTask", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleTask indicates an expected call of HandleTask.
func (mr *MockTaskHandlerMockRecorder) HandleTask(ctx, task any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTask", reflect.TypeOf((*MockTaskHandler)(nil).HandleTask), ctx, task)
}
