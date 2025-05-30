// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/conn/http/http.go
//
// Generated by this command:
//
//	mockgen -source pkg/conn/http/http.go -destination pkg/conn/http/mock/http_mock_gen.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockMetaHttpManagerInterface is a mock of MetaHttpManagerInterface interface.
type MockMetaHttpManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMetaHttpManagerInterfaceMockRecorder
	isgomock struct{}
}

// MockMetaHttpManagerInterfaceMockRecorder is the mock recorder for MockMetaHttpManagerInterface.
type MockMetaHttpManagerInterfaceMockRecorder struct {
	mock *MockMetaHttpManagerInterface
}

// NewMockMetaHttpManagerInterface creates a new mock instance.
func NewMockMetaHttpManagerInterface(ctrl *gomock.Controller) *MockMetaHttpManagerInterface {
	mock := &MockMetaHttpManagerInterface{ctrl: ctrl}
	mock.recorder = &MockMetaHttpManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetaHttpManagerInterface) EXPECT() *MockMetaHttpManagerInterfaceMockRecorder {
	return m.recorder
}

// GetDiagnose mocks base method.
func (m *MockMetaHttpManagerInterface) GetDiagnose(ctx context.Context, endpoint string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiagnose", ctx, endpoint)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiagnose indicates an expected call of GetDiagnose.
func (mr *MockMetaHttpManagerInterfaceMockRecorder) GetDiagnose(ctx, endpoint any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiagnose", reflect.TypeOf((*MockMetaHttpManagerInterface)(nil).GetDiagnose), ctx, endpoint)
}
