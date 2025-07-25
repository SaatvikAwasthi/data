// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go
//
// Generated by this command:
//
//	mockgen -package=mockContract -source=contract.go -destination=../../mock/app/contract/contract.go Contract
//

// Package mockContract is a generated GoMock package.
package mockContract

import (
	context "context"
	reflect "reflect"
	http "tester/crosscutting/http"

	gomock "go.uber.org/mock/gomock"
)

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
	isgomock struct{}
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockHTTPClient) Get(ctx context.Context, path string, queryParams, headers map[string]string, result any) http.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, path, queryParams, headers, result)
	ret0, _ := ret[0].(http.Response)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockHTTPClientMockRecorder) Get(ctx, path, queryParams, headers, result any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClient)(nil).Get), ctx, path, queryParams, headers, result)
}
