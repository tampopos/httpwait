// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/client/client.go

// Package mock_client is a generated GoMock package.
package httpwaittest

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/tampopos/httpwait/src/client"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetStatusCode mocks base method
func (m *MockClient) GetStatusCode(request *client.Request) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatusCode", request)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatusCode indicates an expected call of GetStatusCode
func (mr *MockClientMockRecorder) GetStatusCode(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatusCode", reflect.TypeOf((*MockClient)(nil).GetStatusCode), request)
}

// GetBody mocks base method
func (m *MockClient) GetBody(request *client.Request) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBody", request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBody indicates an expected call of GetBody
func (mr *MockClientMockRecorder) GetBody(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBody", reflect.TypeOf((*MockClient)(nil).GetBody), request)
}
