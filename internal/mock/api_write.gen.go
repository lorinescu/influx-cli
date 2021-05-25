// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/api (interfaces: WriteApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/api"
)

// MockWriteApi is a mock of WriteApi interface.
type MockWriteApi struct {
	ctrl     *gomock.Controller
	recorder *MockWriteApiMockRecorder
}

// MockWriteApiMockRecorder is the mock recorder for MockWriteApi.
type MockWriteApiMockRecorder struct {
	mock *MockWriteApi
}

// NewMockWriteApi creates a new mock instance.
func NewMockWriteApi(ctrl *gomock.Controller) *MockWriteApi {
	mock := &MockWriteApi{ctrl: ctrl}
	mock.recorder = &MockWriteApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWriteApi) EXPECT() *MockWriteApiMockRecorder {
	return m.recorder
}

// PostWrite mocks base method.
func (m *MockWriteApi) PostWrite(arg0 context.Context) api.ApiPostWriteRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostWrite", arg0)
	ret0, _ := ret[0].(api.ApiPostWriteRequest)
	return ret0
}

// PostWrite indicates an expected call of PostWrite.
func (mr *MockWriteApiMockRecorder) PostWrite(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostWrite", reflect.TypeOf((*MockWriteApi)(nil).PostWrite), arg0)
}

// PostWriteExecute mocks base method.
func (m *MockWriteApi) PostWriteExecute(arg0 api.ApiPostWriteRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostWriteExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostWriteExecute indicates an expected call of PostWriteExecute.
func (mr *MockWriteApiMockRecorder) PostWriteExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostWriteExecute", reflect.TypeOf((*MockWriteApi)(nil).PostWriteExecute), arg0)
}
