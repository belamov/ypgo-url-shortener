// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/belamov/ypgo-url-shortener/internal/app/services (interfaces: ShortenerInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/belamov/ypgo-url-shortener/internal/app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockShortenerInterface is a mock of ShortenerInterface interface.
type MockShortenerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockShortenerInterfaceMockRecorder
}

// MockShortenerInterfaceMockRecorder is the mock recorder for MockShortenerInterface.
type MockShortenerInterfaceMockRecorder struct {
	mock *MockShortenerInterface
}

// NewMockShortenerInterface creates a new mock instance.
func NewMockShortenerInterface(ctrl *gomock.Controller) *MockShortenerInterface {
	mock := &MockShortenerInterface{ctrl: ctrl}
	mock.recorder = &MockShortenerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortenerInterface) EXPECT() *MockShortenerInterfaceMockRecorder {
	return m.recorder
}

// DeleteUrls mocks base method.
func (m *MockShortenerInterface) DeleteUrls(arg0 context.Context, arg1 []string, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUrls", arg0, arg1, arg2)
}

// DeleteUrls indicates an expected call of DeleteUrls.
func (mr *MockShortenerInterfaceMockRecorder) DeleteUrls(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUrls", reflect.TypeOf((*MockShortenerInterface)(nil).DeleteUrls), arg0, arg1, arg2)
}

// Expand mocks base method.
func (m *MockShortenerInterface) Expand(arg0 context.Context, arg1 string) (models.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expand", arg0, arg1)
	ret0, _ := ret[0].(models.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expand indicates an expected call of Expand.
func (mr *MockShortenerInterfaceMockRecorder) Expand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expand", reflect.TypeOf((*MockShortenerInterface)(nil).Expand), arg0, arg1)
}

// FormatShortURL mocks base method.
func (m *MockShortenerInterface) FormatShortURL(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatShortURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// FormatShortURL indicates an expected call of FormatShortURL.
func (mr *MockShortenerInterfaceMockRecorder) FormatShortURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatShortURL", reflect.TypeOf((*MockShortenerInterface)(nil).FormatShortURL), arg0)
}

// GenerateNewUserID mocks base method.
func (m *MockShortenerInterface) GenerateNewUserID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateNewUserID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateNewUserID indicates an expected call of GenerateNewUserID.
func (mr *MockShortenerInterfaceMockRecorder) GenerateNewUserID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateNewUserID", reflect.TypeOf((*MockShortenerInterface)(nil).GenerateNewUserID))
}

// GetStats mocks base method.
func (m *MockShortenerInterface) GetStats(arg0 context.Context) (models.Stats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", arg0)
	ret0, _ := ret[0].(models.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockShortenerInterfaceMockRecorder) GetStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockShortenerInterface)(nil).GetStats), arg0)
}

// GetUrlsCreatedBy mocks base method.
func (m *MockShortenerInterface) GetUrlsCreatedBy(arg0 context.Context, arg1 string) ([]models.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUrlsCreatedBy", arg0, arg1)
	ret0, _ := ret[0].([]models.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUrlsCreatedBy indicates an expected call of GetUrlsCreatedBy.
func (mr *MockShortenerInterfaceMockRecorder) GetUrlsCreatedBy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUrlsCreatedBy", reflect.TypeOf((*MockShortenerInterface)(nil).GetUrlsCreatedBy), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockShortenerInterface) HealthCheck(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockShortenerInterfaceMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockShortenerInterface)(nil).HealthCheck), arg0)
}

// Shorten mocks base method.
func (m *MockShortenerInterface) Shorten(arg0 context.Context, arg1, arg2 string) (models.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shorten", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shorten indicates an expected call of Shorten.
func (mr *MockShortenerInterfaceMockRecorder) Shorten(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shorten", reflect.TypeOf((*MockShortenerInterface)(nil).Shorten), arg0, arg1, arg2)
}

// ShortenBatch mocks base method.
func (m *MockShortenerInterface) ShortenBatch(arg0 context.Context, arg1 []models.ShortURL, arg2 string) ([]models.ShortURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShortenBatch", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.ShortURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShortenBatch indicates an expected call of ShortenBatch.
func (mr *MockShortenerInterfaceMockRecorder) ShortenBatch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShortenBatch", reflect.TypeOf((*MockShortenerInterface)(nil).ShortenBatch), arg0, arg1, arg2)
}