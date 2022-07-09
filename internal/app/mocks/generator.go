// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/belamov/ypgo-url-shortener/internal/app/services/generator (interfaces: URLGenerator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockURLGenerator is a mock of URLGenerator interface.
type MockURLGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockURLGeneratorMockRecorder
}

// MockURLGeneratorMockRecorder is the mock recorder for MockURLGenerator.
type MockURLGeneratorMockRecorder struct {
	mock *MockURLGenerator
}

// NewMockURLGenerator creates a new mock instance.
func NewMockURLGenerator(ctrl *gomock.Controller) *MockURLGenerator {
	mock := &MockURLGenerator{ctrl: ctrl}
	mock.recorder = &MockURLGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLGenerator) EXPECT() *MockURLGeneratorMockRecorder {
	return m.recorder
}

// GenerateIDFromString mocks base method.
func (m *MockURLGenerator) GenerateIDFromString(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateIDFromString", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateIDFromString indicates an expected call of GenerateIDFromString.
func (mr *MockURLGeneratorMockRecorder) GenerateIDFromString(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateIDFromString", reflect.TypeOf((*MockURLGenerator)(nil).GenerateIDFromString), arg0)
}
