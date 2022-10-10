// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/messages/model.go

// Package mock_messages is a generated GoMock package.
package mock_messages

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	messages "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/model/messages"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockClient) SendMessage(msg messages.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockClientMockRecorder) SendMessage(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockClient)(nil).SendMessage), msg)
}

// SetupCurrencyKeyboard mocks base method.
func (m *MockClient) SetupCurrencyKeyboard(msg *messages.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetupCurrencyKeyboard", msg)
}

// SetupCurrencyKeyboard indicates an expected call of SetupCurrencyKeyboard.
func (mr *MockClientMockRecorder) SetupCurrencyKeyboard(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupCurrencyKeyboard", reflect.TypeOf((*MockClient)(nil).SetupCurrencyKeyboard), msg)
}

// MockMessanger is a mock of Messanger interface.
type MockMessanger struct {
	ctrl     *gomock.Controller
	recorder *MockMessangerMockRecorder
}

// MockMessangerMockRecorder is the mock recorder for MockMessanger.
type MockMessangerMockRecorder struct {
	mock *MockMessanger
}

// NewMockMessanger creates a new mock instance.
func NewMockMessanger(ctrl *gomock.Controller) *MockMessanger {
	mock := &MockMessanger{ctrl: ctrl}
	mock.recorder = &MockMessangerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessanger) EXPECT() *MockMessangerMockRecorder {
	return m.recorder
}

// IsCurrency mocks base method.
func (m *MockMessanger) IsCurrency(text string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCurrency", text)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCurrency indicates an expected call of IsCurrency.
func (mr *MockMessangerMockRecorder) IsCurrency(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCurrency", reflect.TypeOf((*MockMessanger)(nil).IsCurrency), text)
}

// MessageDefault mocks base method.
func (m *MockMessanger) MessageDefault(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageDefault", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageDefault indicates an expected call of MessageDefault.
func (mr *MockMessangerMockRecorder) MessageDefault(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageDefault", reflect.TypeOf((*MockMessanger)(nil).MessageDefault), msg)
}

// MessageSetCurrency mocks base method.
func (m *MockMessanger) MessageSetCurrency(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageSetCurrency", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageSetCurrency indicates an expected call of MessageSetCurrency.
func (mr *MockMessangerMockRecorder) MessageSetCurrency(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageSetCurrency", reflect.TypeOf((*MockMessanger)(nil).MessageSetCurrency), msg)
}

// MockCommander is a mock of Commander interface.
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander.
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance.
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// CommandDefault mocks base method.
func (m *MockCommander) CommandDefault(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandDefault", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandDefault indicates an expected call of CommandDefault.
func (mr *MockCommanderMockRecorder) CommandDefault(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandDefault", reflect.TypeOf((*MockCommander)(nil).CommandDefault), msg)
}

// CommandGetStatistic mocks base method.
func (m *MockCommander) CommandGetStatistic(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandGetStatistic", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandGetStatistic indicates an expected call of CommandGetStatistic.
func (mr *MockCommanderMockRecorder) CommandGetStatistic(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandGetStatistic", reflect.TypeOf((*MockCommander)(nil).CommandGetStatistic), msg)
}

// CommandHelp mocks base method.
func (m *MockCommander) CommandHelp(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandHelp", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandHelp indicates an expected call of CommandHelp.
func (mr *MockCommanderMockRecorder) CommandHelp(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandHelp", reflect.TypeOf((*MockCommander)(nil).CommandHelp), msg)
}

// CommandSetNote mocks base method.
func (m *MockCommander) CommandSetNote(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandSetNote", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandSetNote indicates an expected call of CommandSetNote.
func (mr *MockCommanderMockRecorder) CommandSetNote(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandSetNote", reflect.TypeOf((*MockCommander)(nil).CommandSetNote), msg)
}

// CommandStart mocks base method.
func (m *MockCommander) CommandStart(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandStart", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandStart indicates an expected call of CommandStart.
func (mr *MockCommanderMockRecorder) CommandStart(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandStart", reflect.TypeOf((*MockCommander)(nil).CommandStart), msg)
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// CommandDefault mocks base method.
func (m *MockServer) CommandDefault(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandDefault", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandDefault indicates an expected call of CommandDefault.
func (mr *MockServerMockRecorder) CommandDefault(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandDefault", reflect.TypeOf((*MockServer)(nil).CommandDefault), msg)
}

// CommandGetStatistic mocks base method.
func (m *MockServer) CommandGetStatistic(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandGetStatistic", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandGetStatistic indicates an expected call of CommandGetStatistic.
func (mr *MockServerMockRecorder) CommandGetStatistic(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandGetStatistic", reflect.TypeOf((*MockServer)(nil).CommandGetStatistic), msg)
}

// CommandHelp mocks base method.
func (m *MockServer) CommandHelp(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandHelp", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandHelp indicates an expected call of CommandHelp.
func (mr *MockServerMockRecorder) CommandHelp(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandHelp", reflect.TypeOf((*MockServer)(nil).CommandHelp), msg)
}

// CommandSetNote mocks base method.
func (m *MockServer) CommandSetNote(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandSetNote", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandSetNote indicates an expected call of CommandSetNote.
func (mr *MockServerMockRecorder) CommandSetNote(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandSetNote", reflect.TypeOf((*MockServer)(nil).CommandSetNote), msg)
}

// CommandStart mocks base method.
func (m *MockServer) CommandStart(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommandStart", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommandStart indicates an expected call of CommandStart.
func (mr *MockServerMockRecorder) CommandStart(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandStart", reflect.TypeOf((*MockServer)(nil).CommandStart), msg)
}

// IsCurrency mocks base method.
func (m *MockServer) IsCurrency(text string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCurrency", text)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCurrency indicates an expected call of IsCurrency.
func (mr *MockServerMockRecorder) IsCurrency(text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCurrency", reflect.TypeOf((*MockServer)(nil).IsCurrency), text)
}

// MessageDefault mocks base method.
func (m *MockServer) MessageDefault(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageDefault", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageDefault indicates an expected call of MessageDefault.
func (mr *MockServerMockRecorder) MessageDefault(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageDefault", reflect.TypeOf((*MockServer)(nil).MessageDefault), msg)
}

// MessageSetCurrency mocks base method.
func (m *MockServer) MessageSetCurrency(msg *messages.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageSetCurrency", msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MessageSetCurrency indicates an expected call of MessageSetCurrency.
func (mr *MockServerMockRecorder) MessageSetCurrency(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageSetCurrency", reflect.TypeOf((*MockServer)(nil).MessageSetCurrency), msg)
}
