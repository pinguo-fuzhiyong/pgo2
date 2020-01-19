// Code generated by MockGen. DO NOT EDIT.
// Source: ./logs/interface.go

// Package mock_logs is a generated GoMock package.
package mock_logs

import (
	gomock "github.com/golang/mock/gomock"
	logs "github.com/pinguo/pgo2/logs"
	reflect "reflect"
)

// MockIFormatter is a mock of IFormatter interface
type MockIFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockIFormatterMockRecorder
}

// MockIFormatterMockRecorder is the mock recorder for MockIFormatter
type MockIFormatterMockRecorder struct {
	mock *MockIFormatter
}

// NewMockIFormatter creates a new mock instance
func NewMockIFormatter(ctrl *gomock.Controller) *MockIFormatter {
	mock := &MockIFormatter{ctrl: ctrl}
	mock.recorder = &MockIFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIFormatter) EXPECT() *MockIFormatterMockRecorder {
	return m.recorder
}

// Format mocks base method
func (m *MockIFormatter) Format(item *logs.LogItem) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Format", item)
	ret0, _ := ret[0].(string)
	return ret0
}

// Format indicates an expected call of Format
func (mr *MockIFormatterMockRecorder) Format(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Format", reflect.TypeOf((*MockIFormatter)(nil).Format), item)
}

// MockITarget is a mock of ITarget interface
type MockITarget struct {
	ctrl     *gomock.Controller
	recorder *MockITargetMockRecorder
}

// MockITargetMockRecorder is the mock recorder for MockITarget
type MockITargetMockRecorder struct {
	mock *MockITarget
}

// NewMockITarget creates a new mock instance
func NewMockITarget(ctrl *gomock.Controller) *MockITarget {
	mock := &MockITarget{ctrl: ctrl}
	mock.recorder = &MockITargetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITarget) EXPECT() *MockITargetMockRecorder {
	return m.recorder
}

// Process mocks base method
func (m *MockITarget) Process(item *logs.LogItem) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Process", item)
}

// Process indicates an expected call of Process
func (mr *MockITargetMockRecorder) Process(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockITarget)(nil).Process), item)
}

// Flush mocks base method
func (m *MockITarget) Flush(final bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush", final)
}

// Flush indicates an expected call of Flush
func (mr *MockITargetMockRecorder) Flush(final interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockITarget)(nil).Flush), final)
}

// MockILogger is a mock of ILogger interface
type MockILogger struct {
	ctrl     *gomock.Controller
	recorder *MockILoggerMockRecorder
}

// MockILoggerMockRecorder is the mock recorder for MockILogger
type MockILoggerMockRecorder struct {
	mock *MockILogger
}

// NewMockILogger creates a new mock instance
func NewMockILogger(ctrl *gomock.Controller) *MockILogger {
	mock := &MockILogger{ctrl: ctrl}
	mock.recorder = &MockILoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockILogger) EXPECT() *MockILoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *MockILogger) Debug(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockILoggerMockRecorder) Debug(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockILogger)(nil).Debug), varargs...)
}

// Info mocks base method
func (m *MockILogger) Info(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockILoggerMockRecorder) Info(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockILogger)(nil).Info), varargs...)
}

// Notice mocks base method
func (m *MockILogger) Notice(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Notice", varargs...)
}

// Notice indicates an expected call of Notice
func (mr *MockILoggerMockRecorder) Notice(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Notice", reflect.TypeOf((*MockILogger)(nil).Notice), varargs...)
}

// Warn mocks base method
func (m *MockILogger) Warn(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn
func (mr *MockILoggerMockRecorder) Warn(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockILogger)(nil).Warn), varargs...)
}

// Error mocks base method
func (m *MockILogger) Error(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockILoggerMockRecorder) Error(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockILogger)(nil).Error), varargs...)
}

// Fatal mocks base method
func (m *MockILogger) Fatal(format string, v ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{format}
	for _, a := range v {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal
func (mr *MockILoggerMockRecorder) Fatal(format interface{}, v ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{format}, v...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockILogger)(nil).Fatal), varargs...)
}
