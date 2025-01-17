// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/sync/http/http_sync.go

// Package syncmock is a generated GoMock package.
package syncmock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
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

// Do mocks base method.
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPClientMockRecorder) Do(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), req)
}

// MockCron is a mock of Cron interface.
type MockCron struct {
	ctrl     *gomock.Controller
	recorder *MockCronMockRecorder
}

// MockCronMockRecorder is the mock recorder for MockCron.
type MockCronMockRecorder struct {
	mock *MockCron
}

// NewMockCron creates a new mock instance.
func NewMockCron(ctrl *gomock.Controller) *MockCron {
	mock := &MockCron{ctrl: ctrl}
	mock.recorder = &MockCronMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCron) EXPECT() *MockCronMockRecorder {
	return m.recorder
}

// AddFunc mocks base method.
func (m *MockCron) AddFunc(spec string, cmd func()) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFunc", spec, cmd)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFunc indicates an expected call of AddFunc.
func (mr *MockCronMockRecorder) AddFunc(spec, cmd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFunc", reflect.TypeOf((*MockCron)(nil).AddFunc), spec, cmd)
}

// Start mocks base method.
func (m *MockCron) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockCronMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCron)(nil).Start))
}

// Stop mocks base method.
func (m *MockCron) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockCronMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCron)(nil).Stop))
}
