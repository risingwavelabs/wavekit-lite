// Code generated by MockGen. DO NOT EDIT.
// Source: internal/macaroons/store/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=internal/macaroons/store/interfaces.go -destination=internal/macaroons/store/mock/mock_gen.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockKeyStore is a mock of KeyStore interface.
type MockKeyStore struct {
	ctrl     *gomock.Controller
	recorder *MockKeyStoreMockRecorder
	isgomock struct{}
}

// MockKeyStoreMockRecorder is the mock recorder for MockKeyStore.
type MockKeyStoreMockRecorder struct {
	mock *MockKeyStore
}

// NewMockKeyStore creates a new mock instance.
func NewMockKeyStore(ctrl *gomock.Controller) *MockKeyStore {
	mock := &MockKeyStore{ctrl: ctrl}
	mock.recorder = &MockKeyStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyStore) EXPECT() *MockKeyStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockKeyStore) Create(ctx context.Context, userID int32, key []byte, ttl time.Duration) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID, key, ttl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockKeyStoreMockRecorder) Create(ctx, userID, key, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockKeyStore)(nil).Create), ctx, userID, key, ttl)
}

// Delete mocks base method.
func (m *MockKeyStore) Delete(ctx context.Context, keyID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, keyID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockKeyStoreMockRecorder) Delete(ctx, keyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKeyStore)(nil).Delete), ctx, keyID)
}

// DeleteUserKeys mocks base method.
func (m *MockKeyStore) DeleteUserKeys(ctx context.Context, userID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserKeys", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserKeys indicates an expected call of DeleteUserKeys.
func (mr *MockKeyStoreMockRecorder) DeleteUserKeys(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserKeys", reflect.TypeOf((*MockKeyStore)(nil).DeleteUserKeys), ctx, userID)
}

// Get mocks base method.
func (m *MockKeyStore) Get(ctx context.Context, keyID int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, keyID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeyStoreMockRecorder) Get(ctx, keyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyStore)(nil).Get), ctx, keyID)
}
