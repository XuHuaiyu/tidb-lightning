// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pingcap/tidb-lightning/lightning/backend (interfaces: AbstractBackend,Encoder,Rows,Row)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	mysql "github.com/pingcap/parser/mysql"
	backend "github.com/pingcap/tidb-lightning/lightning/backend"
	log "github.com/pingcap/tidb-lightning/lightning/log"
	verification "github.com/pingcap/tidb-lightning/lightning/verification"
	table "github.com/pingcap/tidb/table"
	types "github.com/pingcap/tidb/types"
	go_uuid "github.com/satori/go.uuid"
	reflect "reflect"
	time "time"
)

// MockAbstractBackend is a mock of AbstractBackend interface
type MockAbstractBackend struct {
	ctrl     *gomock.Controller
	recorder *MockAbstractBackendMockRecorder
}

// MockAbstractBackendMockRecorder is the mock recorder for MockAbstractBackend
type MockAbstractBackendMockRecorder struct {
	mock *MockAbstractBackend
}

// NewMockAbstractBackend creates a new mock instance
func NewMockAbstractBackend(ctrl *gomock.Controller) *MockAbstractBackend {
	mock := &MockAbstractBackend{ctrl: ctrl}
	mock.recorder = &MockAbstractBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAbstractBackend) EXPECT() *MockAbstractBackendMockRecorder {
	return m.recorder
}

// CleanupEngine mocks base method
func (m *MockAbstractBackend) CleanupEngine(arg0 context.Context, arg1 go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupEngine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupEngine indicates an expected call of CleanupEngine
func (mr *MockAbstractBackendMockRecorder) CleanupEngine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupEngine", reflect.TypeOf((*MockAbstractBackend)(nil).CleanupEngine), arg0, arg1)
}

// Close mocks base method
func (m *MockAbstractBackend) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockAbstractBackendMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAbstractBackend)(nil).Close))
}

// CloseEngine mocks base method
func (m *MockAbstractBackend) CloseEngine(arg0 context.Context, arg1 go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseEngine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseEngine indicates an expected call of CloseEngine
func (mr *MockAbstractBackendMockRecorder) CloseEngine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseEngine", reflect.TypeOf((*MockAbstractBackend)(nil).CloseEngine), arg0, arg1)
}

// ImportEngine mocks base method
func (m *MockAbstractBackend) ImportEngine(arg0 context.Context, arg1 go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportEngine", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportEngine indicates an expected call of ImportEngine
func (mr *MockAbstractBackendMockRecorder) ImportEngine(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportEngine", reflect.TypeOf((*MockAbstractBackend)(nil).ImportEngine), arg0, arg1)
}

// MakeEmptyRows mocks base method
func (m *MockAbstractBackend) MakeEmptyRows() backend.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeEmptyRows")
	ret0, _ := ret[0].(backend.Rows)
	return ret0
}

// MakeEmptyRows indicates an expected call of MakeEmptyRows
func (mr *MockAbstractBackendMockRecorder) MakeEmptyRows() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeEmptyRows", reflect.TypeOf((*MockAbstractBackend)(nil).MakeEmptyRows))
}

// MaxChunkSize mocks base method
func (m *MockAbstractBackend) MaxChunkSize() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxChunkSize")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxChunkSize indicates an expected call of MaxChunkSize
func (mr *MockAbstractBackendMockRecorder) MaxChunkSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxChunkSize", reflect.TypeOf((*MockAbstractBackend)(nil).MaxChunkSize))
}

// NewEncoder mocks base method
func (m *MockAbstractBackend) NewEncoder(arg0 table.Table, arg1 mysql.SQLMode) backend.Encoder {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEncoder", arg0, arg1)
	ret0, _ := ret[0].(backend.Encoder)
	return ret0
}

// NewEncoder indicates an expected call of NewEncoder
func (mr *MockAbstractBackendMockRecorder) NewEncoder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEncoder", reflect.TypeOf((*MockAbstractBackend)(nil).NewEncoder), arg0, arg1)
}

// OpenEngine mocks base method
func (m *MockAbstractBackend) OpenEngine(arg0 context.Context, arg1 go_uuid.UUID, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenEngine", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenEngine indicates an expected call of OpenEngine
func (mr *MockAbstractBackendMockRecorder) OpenEngine(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenEngine", reflect.TypeOf((*MockAbstractBackend)(nil).OpenEngine), arg0, arg1, arg2)
}

// RetryImportDelay mocks base method
func (m *MockAbstractBackend) RetryImportDelay() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetryImportDelay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// RetryImportDelay indicates an expected call of RetryImportDelay
func (mr *MockAbstractBackendMockRecorder) RetryImportDelay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetryImportDelay", reflect.TypeOf((*MockAbstractBackend)(nil).RetryImportDelay))
}

// ShouldPostProcess mocks base method
func (m *MockAbstractBackend) ShouldPostProcess() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldPostProcess")
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldPostProcess indicates an expected call of ShouldPostProcess
func (mr *MockAbstractBackendMockRecorder) ShouldPostProcess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldPostProcess", reflect.TypeOf((*MockAbstractBackend)(nil).ShouldPostProcess))
}

// WriteRows mocks base method
func (m *MockAbstractBackend) WriteRows(arg0 context.Context, arg1 go_uuid.UUID, arg2 string, arg3 []string, arg4 uint64, arg5 backend.Rows) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteRows", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteRows indicates an expected call of WriteRows
func (mr *MockAbstractBackendMockRecorder) WriteRows(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteRows", reflect.TypeOf((*MockAbstractBackend)(nil).WriteRows), arg0, arg1, arg2, arg3, arg4, arg5)
}

// MockEncoder is a mock of Encoder interface
type MockEncoder struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder
type MockEncoderMockRecorder struct {
	mock *MockEncoder
}

// NewMockEncoder creates a new mock instance
func NewMockEncoder(ctrl *gomock.Controller) *MockEncoder {
	mock := &MockEncoder{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEncoder) EXPECT() *MockEncoderMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockEncoder) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockEncoderMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEncoder)(nil).Close))
}

// Encode mocks base method
func (m *MockEncoder) Encode(arg0 log.Logger, arg1 []types.Datum, arg2 int64, arg3 []int) (backend.Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(backend.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockEncoderMockRecorder) Encode(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder)(nil).Encode), arg0, arg1, arg2, arg3)
}

// MockRows is a mock of Rows interface
type MockRows struct {
	ctrl     *gomock.Controller
	recorder *MockRowsMockRecorder
}

// MockRowsMockRecorder is the mock recorder for MockRows
type MockRowsMockRecorder struct {
	mock *MockRows
}

// NewMockRows creates a new mock instance
func NewMockRows(ctrl *gomock.Controller) *MockRows {
	mock := &MockRows{ctrl: ctrl}
	mock.recorder = &MockRowsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRows) EXPECT() *MockRowsMockRecorder {
	return m.recorder
}

// Clear mocks base method
func (m *MockRows) Clear() backend.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(backend.Rows)
	return ret0
}

// Clear indicates an expected call of Clear
func (mr *MockRowsMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockRows)(nil).Clear))
}

// SplitIntoChunks mocks base method
func (m *MockRows) SplitIntoChunks(arg0 int, arg1 bool) []backend.Rows {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SplitIntoChunks", arg0, arg1)
	ret0, _ := ret[0].([]backend.Rows)
	return ret0
}

// SplitIntoChunks indicates an expected call of SplitIntoChunks
func (mr *MockRowsMockRecorder) SplitIntoChunks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SplitIntoChunks", reflect.TypeOf((*MockRows)(nil).SplitIntoChunks), arg0, arg1)
}

// MockRow is a mock of Row interface
type MockRow struct {
	ctrl     *gomock.Controller
	recorder *MockRowMockRecorder
}

// MockRowMockRecorder is the mock recorder for MockRow
type MockRowMockRecorder struct {
	mock *MockRow
}

// NewMockRow creates a new mock instance
func NewMockRow(ctrl *gomock.Controller) *MockRow {
	mock := &MockRow{ctrl: ctrl}
	mock.recorder = &MockRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRow) EXPECT() *MockRowMockRecorder {
	return m.recorder
}

// ClassifyAndAppend mocks base method
func (m *MockRow) ClassifyAndAppend(arg0 *backend.Rows, arg1 *verification.KVChecksum, arg2 *backend.Rows, arg3 *verification.KVChecksum) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClassifyAndAppend", arg0, arg1, arg2, arg3)
}

// ClassifyAndAppend indicates an expected call of ClassifyAndAppend
func (mr *MockRowMockRecorder) ClassifyAndAppend(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassifyAndAppend", reflect.TypeOf((*MockRow)(nil).ClassifyAndAppend), arg0, arg1, arg2, arg3)
}
