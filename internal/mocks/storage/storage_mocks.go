// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	diary "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/diary"
)

// MockNotesDB is a mock of NotesDB interface.
type MockNotesDB struct {
	ctrl     *gomock.Controller
	recorder *MockNotesDBMockRecorder
}

// MockNotesDBMockRecorder is the mock recorder for MockNotesDB.
type MockNotesDBMockRecorder struct {
	mock *MockNotesDB
}

// NewMockNotesDB creates a new mock instance.
func NewMockNotesDB(ctrl *gomock.Controller) *MockNotesDB {
	mock := &MockNotesDB{ctrl: ctrl}
	mock.recorder = &MockNotesDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotesDB) EXPECT() *MockNotesDBMockRecorder {
	return m.recorder
}

// AddNote mocks base method.
func (m *MockNotesDB) AddNote(id int64, date string, note diary.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", id, date, note)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockNotesDBMockRecorder) AddNote(id, date, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockNotesDB)(nil).AddNote), id, date, note)
}

// GetNote mocks base method.
func (m *MockNotesDB) GetNote(id int64, date string) ([]diary.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", id, date)
	ret0, _ := ret[0].([]diary.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockNotesDBMockRecorder) GetNote(id, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockNotesDB)(nil).GetNote), id, date)
}

// MockRatesDB is a mock of RatesDB interface.
type MockRatesDB struct {
	ctrl     *gomock.Controller
	recorder *MockRatesDBMockRecorder
}

// MockRatesDBMockRecorder is the mock recorder for MockRatesDB.
type MockRatesDBMockRecorder struct {
	mock *MockRatesDB
}

// NewMockRatesDB creates a new mock instance.
func NewMockRatesDB(ctrl *gomock.Controller) *MockRatesDB {
	mock := &MockRatesDB{ctrl: ctrl}
	mock.recorder = &MockRatesDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRatesDB) EXPECT() *MockRatesDBMockRecorder {
	return m.recorder
}

// AddRate mocks base method.
func (m *MockRatesDB) AddRate(valute diary.Valute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRate", valute)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRate indicates an expected call of AddRate.
func (mr *MockRatesDBMockRecorder) AddRate(valute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRate", reflect.TypeOf((*MockRatesDB)(nil).AddRate), valute)
}

// GetRate mocks base method.
func (m *MockRatesDB) GetRate(abbreviation string) (*diary.Valute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate", abbreviation)
	ret0, _ := ret[0].(*diary.Valute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockRatesDBMockRecorder) GetRate(abbreviation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockRatesDB)(nil).GetRate), abbreviation)
}

// SetDefaultCurrency mocks base method.
func (m *MockRatesDB) SetDefaultCurrency() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefaultCurrency")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultCurrency indicates an expected call of SetDefaultCurrency.
func (mr *MockRatesDBMockRecorder) SetDefaultCurrency() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultCurrency", reflect.TypeOf((*MockRatesDB)(nil).SetDefaultCurrency))
}

// MockUsersDB is a mock of UsersDB interface.
type MockUsersDB struct {
	ctrl     *gomock.Controller
	recorder *MockUsersDBMockRecorder
}

// MockUsersDBMockRecorder is the mock recorder for MockUsersDB.
type MockUsersDBMockRecorder struct {
	mock *MockUsersDB
}

// NewMockUsersDB creates a new mock instance.
func NewMockUsersDB(ctrl *gomock.Controller) *MockUsersDB {
	mock := &MockUsersDB{ctrl: ctrl}
	mock.recorder = &MockUsersDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersDB) EXPECT() *MockUsersDBMockRecorder {
	return m.recorder
}

// GetUserAbbValute mocks base method.
func (m *MockUsersDB) GetUserAbbValute(userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAbbValute", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAbbValute indicates an expected call of GetUserAbbValute.
func (mr *MockUsersDBMockRecorder) GetUserAbbValute(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAbbValute", reflect.TypeOf((*MockUsersDB)(nil).GetUserAbbValute), userID)
}

// SetUserAbbValute mocks base method.
func (m *MockUsersDB) SetUserAbbValute(userID int64, abbreviation string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserAbbValute", userID, abbreviation)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserAbbValute indicates an expected call of SetUserAbbValute.
func (mr *MockUsersDBMockRecorder) SetUserAbbValute(userID, abbreviation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAbbValute", reflect.TypeOf((*MockUsersDB)(nil).SetUserAbbValute), userID, abbreviation)
}

// MockBudgetsDB is a mock of BudgetsDB interface.
type MockBudgetsDB struct {
	ctrl     *gomock.Controller
	recorder *MockBudgetsDBMockRecorder
}

// MockBudgetsDBMockRecorder is the mock recorder for MockBudgetsDB.
type MockBudgetsDBMockRecorder struct {
	mock *MockBudgetsDB
}

// NewMockBudgetsDB creates a new mock instance.
func NewMockBudgetsDB(ctrl *gomock.Controller) *MockBudgetsDB {
	mock := &MockBudgetsDB{ctrl: ctrl}
	mock.recorder = &MockBudgetsDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBudgetsDB) EXPECT() *MockBudgetsDBMockRecorder {
	return m.recorder
}

// AddMonthlyBudget mocks base method.
func (m *MockBudgetsDB) AddMonthlyBudget(userID int64, monthlyBudget diary.Budget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMonthlyBudget", userID, monthlyBudget)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMonthlyBudget indicates an expected call of AddMonthlyBudget.
func (mr *MockBudgetsDBMockRecorder) AddMonthlyBudget(userID, monthlyBudget interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMonthlyBudget", reflect.TypeOf((*MockBudgetsDB)(nil).AddMonthlyBudget), userID, monthlyBudget)
}

// GetMonthlyBudget mocks base method.
func (m *MockBudgetsDB) GetMonthlyBudget(userID int64, date string) (*diary.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonthlyBudget", userID, date)
	ret0, _ := ret[0].(*diary.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMonthlyBudget indicates an expected call of GetMonthlyBudget.
func (mr *MockBudgetsDBMockRecorder) GetMonthlyBudget(userID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonthlyBudget", reflect.TypeOf((*MockBudgetsDB)(nil).GetMonthlyBudget), userID, date)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddMonthlyBudget mocks base method.
func (m *MockStorage) AddMonthlyBudget(userID int64, monthlyBudget diary.Budget) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMonthlyBudget", userID, monthlyBudget)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMonthlyBudget indicates an expected call of AddMonthlyBudget.
func (mr *MockStorageMockRecorder) AddMonthlyBudget(userID, monthlyBudget interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMonthlyBudget", reflect.TypeOf((*MockStorage)(nil).AddMonthlyBudget), userID, monthlyBudget)
}

// AddNote mocks base method.
func (m *MockStorage) AddNote(id int64, date string, note diary.Note) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNote", id, date, note)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNote indicates an expected call of AddNote.
func (mr *MockStorageMockRecorder) AddNote(id, date, note interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNote", reflect.TypeOf((*MockStorage)(nil).AddNote), id, date, note)
}

// AddRate mocks base method.
func (m *MockStorage) AddRate(valute diary.Valute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRate", valute)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRate indicates an expected call of AddRate.
func (mr *MockStorageMockRecorder) AddRate(valute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRate", reflect.TypeOf((*MockStorage)(nil).AddRate), valute)
}

// GetMonthlyBudget mocks base method.
func (m *MockStorage) GetMonthlyBudget(userID int64, date string) (*diary.Budget, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMonthlyBudget", userID, date)
	ret0, _ := ret[0].(*diary.Budget)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMonthlyBudget indicates an expected call of GetMonthlyBudget.
func (mr *MockStorageMockRecorder) GetMonthlyBudget(userID, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMonthlyBudget", reflect.TypeOf((*MockStorage)(nil).GetMonthlyBudget), userID, date)
}

// GetNote mocks base method.
func (m *MockStorage) GetNote(id int64, date string) ([]diary.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNote", id, date)
	ret0, _ := ret[0].([]diary.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNote indicates an expected call of GetNote.
func (mr *MockStorageMockRecorder) GetNote(id, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNote", reflect.TypeOf((*MockStorage)(nil).GetNote), id, date)
}

// GetRate mocks base method.
func (m *MockStorage) GetRate(abbreviation string) (*diary.Valute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRate", abbreviation)
	ret0, _ := ret[0].(*diary.Valute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRate indicates an expected call of GetRate.
func (mr *MockStorageMockRecorder) GetRate(abbreviation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRate", reflect.TypeOf((*MockStorage)(nil).GetRate), abbreviation)
}

// GetUserAbbValute mocks base method.
func (m *MockStorage) GetUserAbbValute(userID int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAbbValute", userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAbbValute indicates an expected call of GetUserAbbValute.
func (mr *MockStorageMockRecorder) GetUserAbbValute(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAbbValute", reflect.TypeOf((*MockStorage)(nil).GetUserAbbValute), userID)
}

// SetDefaultCurrency mocks base method.
func (m *MockStorage) SetDefaultCurrency() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDefaultCurrency")
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDefaultCurrency indicates an expected call of SetDefaultCurrency.
func (mr *MockStorageMockRecorder) SetDefaultCurrency() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDefaultCurrency", reflect.TypeOf((*MockStorage)(nil).SetDefaultCurrency))
}

// SetUserAbbValute mocks base method.
func (m *MockStorage) SetUserAbbValute(userID int64, abbreviation string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUserAbbValute", userID, abbreviation)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUserAbbValute indicates an expected call of SetUserAbbValute.
func (mr *MockStorageMockRecorder) SetUserAbbValute(userID, abbreviation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserAbbValute", reflect.TypeOf((*MockStorage)(nil).SetUserAbbValute), userID, abbreviation)
}