// Code generated by MockGen. DO NOT EDIT.
// Source: projects.blender.org/studio/flamenco/internal/manager/sleep_scheduler (interfaces: PersistenceService,ChangeBroadcaster)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	persistence "projects.blender.org/studio/flamenco/internal/manager/persistence"
	api "projects.blender.org/studio/flamenco/pkg/api"
)

// MockPersistenceService is a mock of PersistenceService interface.
type MockPersistenceService struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceServiceMockRecorder
}

// MockPersistenceServiceMockRecorder is the mock recorder for MockPersistenceService.
type MockPersistenceServiceMockRecorder struct {
	mock *MockPersistenceService
}

// NewMockPersistenceService creates a new mock instance.
func NewMockPersistenceService(ctrl *gomock.Controller) *MockPersistenceService {
	mock := &MockPersistenceService{ctrl: ctrl}
	mock.recorder = &MockPersistenceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersistenceService) EXPECT() *MockPersistenceServiceMockRecorder {
	return m.recorder
}

// FetchSleepScheduleWorker mocks base method.
func (m *MockPersistenceService) FetchSleepScheduleWorker(arg0 context.Context, arg1 *persistence.SleepSchedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSleepScheduleWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchSleepScheduleWorker indicates an expected call of FetchSleepScheduleWorker.
func (mr *MockPersistenceServiceMockRecorder) FetchSleepScheduleWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSleepScheduleWorker", reflect.TypeOf((*MockPersistenceService)(nil).FetchSleepScheduleWorker), arg0, arg1)
}

// FetchSleepSchedulesToCheck mocks base method.
func (m *MockPersistenceService) FetchSleepSchedulesToCheck(arg0 context.Context) ([]*persistence.SleepSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSleepSchedulesToCheck", arg0)
	ret0, _ := ret[0].([]*persistence.SleepSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSleepSchedulesToCheck indicates an expected call of FetchSleepSchedulesToCheck.
func (mr *MockPersistenceServiceMockRecorder) FetchSleepSchedulesToCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSleepSchedulesToCheck", reflect.TypeOf((*MockPersistenceService)(nil).FetchSleepSchedulesToCheck), arg0)
}

// FetchWorkerSleepSchedule mocks base method.
func (m *MockPersistenceService) FetchWorkerSleepSchedule(arg0 context.Context, arg1 string) (*persistence.SleepSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchWorkerSleepSchedule", arg0, arg1)
	ret0, _ := ret[0].(*persistence.SleepSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchWorkerSleepSchedule indicates an expected call of FetchWorkerSleepSchedule.
func (mr *MockPersistenceServiceMockRecorder) FetchWorkerSleepSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWorkerSleepSchedule", reflect.TypeOf((*MockPersistenceService)(nil).FetchWorkerSleepSchedule), arg0, arg1)
}

// SaveWorkerStatus mocks base method.
func (m *MockPersistenceService) SaveWorkerStatus(arg0 context.Context, arg1 *persistence.Worker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWorkerStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWorkerStatus indicates an expected call of SaveWorkerStatus.
func (mr *MockPersistenceServiceMockRecorder) SaveWorkerStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWorkerStatus", reflect.TypeOf((*MockPersistenceService)(nil).SaveWorkerStatus), arg0, arg1)
}

// SetWorkerSleepSchedule mocks base method.
func (m *MockPersistenceService) SetWorkerSleepSchedule(arg0 context.Context, arg1 string, arg2 *persistence.SleepSchedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkerSleepSchedule", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorkerSleepSchedule indicates an expected call of SetWorkerSleepSchedule.
func (mr *MockPersistenceServiceMockRecorder) SetWorkerSleepSchedule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkerSleepSchedule", reflect.TypeOf((*MockPersistenceService)(nil).SetWorkerSleepSchedule), arg0, arg1, arg2)
}

// SetWorkerSleepScheduleNextCheck mocks base method.
func (m *MockPersistenceService) SetWorkerSleepScheduleNextCheck(arg0 context.Context, arg1 *persistence.SleepSchedule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkerSleepScheduleNextCheck", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorkerSleepScheduleNextCheck indicates an expected call of SetWorkerSleepScheduleNextCheck.
func (mr *MockPersistenceServiceMockRecorder) SetWorkerSleepScheduleNextCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkerSleepScheduleNextCheck", reflect.TypeOf((*MockPersistenceService)(nil).SetWorkerSleepScheduleNextCheck), arg0, arg1)
}

// MockChangeBroadcaster is a mock of ChangeBroadcaster interface.
type MockChangeBroadcaster struct {
	ctrl     *gomock.Controller
	recorder *MockChangeBroadcasterMockRecorder
}

// MockChangeBroadcasterMockRecorder is the mock recorder for MockChangeBroadcaster.
type MockChangeBroadcasterMockRecorder struct {
	mock *MockChangeBroadcaster
}

// NewMockChangeBroadcaster creates a new mock instance.
func NewMockChangeBroadcaster(ctrl *gomock.Controller) *MockChangeBroadcaster {
	mock := &MockChangeBroadcaster{ctrl: ctrl}
	mock.recorder = &MockChangeBroadcasterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChangeBroadcaster) EXPECT() *MockChangeBroadcasterMockRecorder {
	return m.recorder
}

// BroadcastWorkerUpdate mocks base method.
func (m *MockChangeBroadcaster) BroadcastWorkerUpdate(arg0 api.SocketIOWorkerUpdate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastWorkerUpdate", arg0)
}

// BroadcastWorkerUpdate indicates an expected call of BroadcastWorkerUpdate.
func (mr *MockChangeBroadcasterMockRecorder) BroadcastWorkerUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastWorkerUpdate", reflect.TypeOf((*MockChangeBroadcaster)(nil).BroadcastWorkerUpdate), arg0)
}
