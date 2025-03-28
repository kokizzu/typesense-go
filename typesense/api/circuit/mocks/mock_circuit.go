// Code generated by MockGen. DO NOT EDIT.
// Source: http_client.go
//
// Generated by this command:
//
//	mockgen -source=http_client.go -destination=mocks/mock_circuit.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHTTPRequestDoer is a mock of HTTPRequestDoer interface.
type MockHTTPRequestDoer struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRequestDoerMockRecorder
	isgomock struct{}
}

// MockHTTPRequestDoerMockRecorder is the mock recorder for MockHTTPRequestDoer.
type MockHTTPRequestDoerMockRecorder struct {
	mock *MockHTTPRequestDoer
}

// NewMockHTTPRequestDoer creates a new mock instance.
func NewMockHTTPRequestDoer(ctrl *gomock.Controller) *MockHTTPRequestDoer {
	mock := &MockHTTPRequestDoer{ctrl: ctrl}
	mock.recorder = &MockHTTPRequestDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPRequestDoer) EXPECT() *MockHTTPRequestDoerMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHTTPRequestDoer) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPRequestDoerMockRecorder) Do(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPRequestDoer)(nil).Do), req)
}

// MockBreaker is a mock of Breaker interface.
type MockBreaker struct {
	ctrl     *gomock.Controller
	recorder *MockBreakerMockRecorder
	isgomock struct{}
}

// MockBreakerMockRecorder is the mock recorder for MockBreaker.
type MockBreakerMockRecorder struct {
	mock *MockBreaker
}

// NewMockBreaker creates a new mock instance.
func NewMockBreaker(ctrl *gomock.Controller) *MockBreaker {
	mock := &MockBreaker{ctrl: ctrl}
	mock.recorder = &MockBreakerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBreaker) EXPECT() *MockBreakerMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockBreaker) Execute(req func() error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockBreakerMockRecorder) Execute(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockBreaker)(nil).Execute), req)
}
