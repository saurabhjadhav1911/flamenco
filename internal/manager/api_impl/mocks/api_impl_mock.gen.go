// Code generated by MockGen. DO NOT EDIT.
// Source: git.blender.org/flamenco/internal/manager/api_impl (interfaces: PersistenceService,ChangeBroadcaster,JobCompiler,LogStorage,ConfigService,TaskStateMachine,Shaman)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	config "git.blender.org/flamenco/internal/manager/config"
	job_compilers "git.blender.org/flamenco/internal/manager/job_compilers"
	persistence "git.blender.org/flamenco/internal/manager/persistence"
	api "git.blender.org/flamenco/pkg/api"
	gomock "github.com/golang/mock/gomock"
	zerolog "github.com/rs/zerolog"
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

// CreateWorker mocks base method.
func (m *MockPersistenceService) CreateWorker(arg0 context.Context, arg1 *persistence.Worker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWorker indicates an expected call of CreateWorker.
func (mr *MockPersistenceServiceMockRecorder) CreateWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorker", reflect.TypeOf((*MockPersistenceService)(nil).CreateWorker), arg0, arg1)
}

// FetchJob mocks base method.
func (m *MockPersistenceService) FetchJob(arg0 context.Context, arg1 string) (*persistence.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJob", arg0, arg1)
	ret0, _ := ret[0].(*persistence.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJob indicates an expected call of FetchJob.
func (mr *MockPersistenceServiceMockRecorder) FetchJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJob", reflect.TypeOf((*MockPersistenceService)(nil).FetchJob), arg0, arg1)
}

// FetchTask mocks base method.
func (m *MockPersistenceService) FetchTask(arg0 context.Context, arg1 string) (*persistence.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchTask", arg0, arg1)
	ret0, _ := ret[0].(*persistence.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchTask indicates an expected call of FetchTask.
func (mr *MockPersistenceServiceMockRecorder) FetchTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchTask", reflect.TypeOf((*MockPersistenceService)(nil).FetchTask), arg0, arg1)
}

// FetchWorker mocks base method.
func (m *MockPersistenceService) FetchWorker(arg0 context.Context, arg1 string) (*persistence.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchWorker", arg0, arg1)
	ret0, _ := ret[0].(*persistence.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchWorker indicates an expected call of FetchWorker.
func (mr *MockPersistenceServiceMockRecorder) FetchWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWorker", reflect.TypeOf((*MockPersistenceService)(nil).FetchWorker), arg0, arg1)
}

// FetchWorkers mocks base method.
func (m *MockPersistenceService) FetchWorkers(arg0 context.Context) ([]*persistence.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchWorkers", arg0)
	ret0, _ := ret[0].([]*persistence.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchWorkers indicates an expected call of FetchWorkers.
func (mr *MockPersistenceServiceMockRecorder) FetchWorkers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchWorkers", reflect.TypeOf((*MockPersistenceService)(nil).FetchWorkers), arg0)
}

// QueryJobTaskSummaries mocks base method.
func (m *MockPersistenceService) QueryJobTaskSummaries(arg0 context.Context, arg1 string) ([]*persistence.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryJobTaskSummaries", arg0, arg1)
	ret0, _ := ret[0].([]*persistence.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryJobTaskSummaries indicates an expected call of QueryJobTaskSummaries.
func (mr *MockPersistenceServiceMockRecorder) QueryJobTaskSummaries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryJobTaskSummaries", reflect.TypeOf((*MockPersistenceService)(nil).QueryJobTaskSummaries), arg0, arg1)
}

// QueryJobs mocks base method.
func (m *MockPersistenceService) QueryJobs(arg0 context.Context, arg1 api.JobsQuery) ([]*persistence.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryJobs", arg0, arg1)
	ret0, _ := ret[0].([]*persistence.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryJobs indicates an expected call of QueryJobs.
func (mr *MockPersistenceServiceMockRecorder) QueryJobs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryJobs", reflect.TypeOf((*MockPersistenceService)(nil).QueryJobs), arg0, arg1)
}

// SaveTask mocks base method.
func (m *MockPersistenceService) SaveTask(arg0 context.Context, arg1 *persistence.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTask indicates an expected call of SaveTask.
func (mr *MockPersistenceServiceMockRecorder) SaveTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTask", reflect.TypeOf((*MockPersistenceService)(nil).SaveTask), arg0, arg1)
}

// SaveTaskActivity mocks base method.
func (m *MockPersistenceService) SaveTaskActivity(arg0 context.Context, arg1 *persistence.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTaskActivity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTaskActivity indicates an expected call of SaveTaskActivity.
func (mr *MockPersistenceServiceMockRecorder) SaveTaskActivity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTaskActivity", reflect.TypeOf((*MockPersistenceService)(nil).SaveTaskActivity), arg0, arg1)
}

// SaveWorker mocks base method.
func (m *MockPersistenceService) SaveWorker(arg0 context.Context, arg1 *persistence.Worker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWorker indicates an expected call of SaveWorker.
func (mr *MockPersistenceServiceMockRecorder) SaveWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWorker", reflect.TypeOf((*MockPersistenceService)(nil).SaveWorker), arg0, arg1)
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

// ScheduleTask mocks base method.
func (m *MockPersistenceService) ScheduleTask(arg0 context.Context, arg1 *persistence.Worker) (*persistence.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleTask", arg0, arg1)
	ret0, _ := ret[0].(*persistence.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleTask indicates an expected call of ScheduleTask.
func (mr *MockPersistenceServiceMockRecorder) ScheduleTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleTask", reflect.TypeOf((*MockPersistenceService)(nil).ScheduleTask), arg0, arg1)
}

// StoreAuthoredJob mocks base method.
func (m *MockPersistenceService) StoreAuthoredJob(arg0 context.Context, arg1 job_compilers.AuthoredJob) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAuthoredJob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreAuthoredJob indicates an expected call of StoreAuthoredJob.
func (mr *MockPersistenceServiceMockRecorder) StoreAuthoredJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAuthoredJob", reflect.TypeOf((*MockPersistenceService)(nil).StoreAuthoredJob), arg0, arg1)
}

// TaskTouchedByWorker mocks base method.
func (m *MockPersistenceService) TaskTouchedByWorker(arg0 context.Context, arg1 *persistence.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskTouchedByWorker", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TaskTouchedByWorker indicates an expected call of TaskTouchedByWorker.
func (mr *MockPersistenceServiceMockRecorder) TaskTouchedByWorker(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskTouchedByWorker", reflect.TypeOf((*MockPersistenceService)(nil).TaskTouchedByWorker), arg0, arg1)
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

// BroadcastNewJob mocks base method.
func (m *MockChangeBroadcaster) BroadcastNewJob(arg0 api.SocketIOJobUpdate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastNewJob", arg0)
}

// BroadcastNewJob indicates an expected call of BroadcastNewJob.
func (mr *MockChangeBroadcasterMockRecorder) BroadcastNewJob(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastNewJob", reflect.TypeOf((*MockChangeBroadcaster)(nil).BroadcastNewJob), arg0)
}

// BroadcastNewWorker mocks base method.
func (m *MockChangeBroadcaster) BroadcastNewWorker(arg0 api.SocketIOWorkerUpdate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastNewWorker", arg0)
}

// BroadcastNewWorker indicates an expected call of BroadcastNewWorker.
func (mr *MockChangeBroadcasterMockRecorder) BroadcastNewWorker(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastNewWorker", reflect.TypeOf((*MockChangeBroadcaster)(nil).BroadcastNewWorker), arg0)
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

// MockJobCompiler is a mock of JobCompiler interface.
type MockJobCompiler struct {
	ctrl     *gomock.Controller
	recorder *MockJobCompilerMockRecorder
}

// MockJobCompilerMockRecorder is the mock recorder for MockJobCompiler.
type MockJobCompilerMockRecorder struct {
	mock *MockJobCompiler
}

// NewMockJobCompiler creates a new mock instance.
func NewMockJobCompiler(ctrl *gomock.Controller) *MockJobCompiler {
	mock := &MockJobCompiler{ctrl: ctrl}
	mock.recorder = &MockJobCompilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobCompiler) EXPECT() *MockJobCompilerMockRecorder {
	return m.recorder
}

// Compile mocks base method.
func (m *MockJobCompiler) Compile(arg0 context.Context, arg1 api.SubmittedJob) (*job_compilers.AuthoredJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compile", arg0, arg1)
	ret0, _ := ret[0].(*job_compilers.AuthoredJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compile indicates an expected call of Compile.
func (mr *MockJobCompilerMockRecorder) Compile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compile", reflect.TypeOf((*MockJobCompiler)(nil).Compile), arg0, arg1)
}

// GetJobType mocks base method.
func (m *MockJobCompiler) GetJobType(arg0 string) (api.AvailableJobType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobType", arg0)
	ret0, _ := ret[0].(api.AvailableJobType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobType indicates an expected call of GetJobType.
func (mr *MockJobCompilerMockRecorder) GetJobType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobType", reflect.TypeOf((*MockJobCompiler)(nil).GetJobType), arg0)
}

// ListJobTypes mocks base method.
func (m *MockJobCompiler) ListJobTypes() api.AvailableJobTypes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJobTypes")
	ret0, _ := ret[0].(api.AvailableJobTypes)
	return ret0
}

// ListJobTypes indicates an expected call of ListJobTypes.
func (mr *MockJobCompilerMockRecorder) ListJobTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobTypes", reflect.TypeOf((*MockJobCompiler)(nil).ListJobTypes))
}

// MockLogStorage is a mock of LogStorage interface.
type MockLogStorage struct {
	ctrl     *gomock.Controller
	recorder *MockLogStorageMockRecorder
}

// MockLogStorageMockRecorder is the mock recorder for MockLogStorage.
type MockLogStorageMockRecorder struct {
	mock *MockLogStorage
}

// NewMockLogStorage creates a new mock instance.
func NewMockLogStorage(ctrl *gomock.Controller) *MockLogStorage {
	mock := &MockLogStorage{ctrl: ctrl}
	mock.recorder = &MockLogStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogStorage) EXPECT() *MockLogStorageMockRecorder {
	return m.recorder
}

// RotateFile mocks base method.
func (m *MockLogStorage) RotateFile(arg0 zerolog.Logger, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RotateFile", arg0, arg1, arg2)
}

// RotateFile indicates an expected call of RotateFile.
func (mr *MockLogStorageMockRecorder) RotateFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RotateFile", reflect.TypeOf((*MockLogStorage)(nil).RotateFile), arg0, arg1, arg2)
}

// Tail mocks base method.
func (m *MockLogStorage) Tail(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tail", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tail indicates an expected call of Tail.
func (mr *MockLogStorageMockRecorder) Tail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tail", reflect.TypeOf((*MockLogStorage)(nil).Tail), arg0, arg1)
}

// Write mocks base method.
func (m *MockLogStorage) Write(arg0 zerolog.Logger, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockLogStorageMockRecorder) Write(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockLogStorage)(nil).Write), arg0, arg1, arg2, arg3)
}

// WriteTimestamped mocks base method.
func (m *MockLogStorage) WriteTimestamped(arg0 zerolog.Logger, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTimestamped", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTimestamped indicates an expected call of WriteTimestamped.
func (mr *MockLogStorageMockRecorder) WriteTimestamped(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTimestamped", reflect.TypeOf((*MockLogStorage)(nil).WriteTimestamped), arg0, arg1, arg2, arg3)
}

// MockConfigService is a mock of ConfigService interface.
type MockConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockConfigServiceMockRecorder
}

// MockConfigServiceMockRecorder is the mock recorder for MockConfigService.
type MockConfigServiceMockRecorder struct {
	mock *MockConfigService
}

// NewMockConfigService creates a new mock instance.
func NewMockConfigService(ctrl *gomock.Controller) *MockConfigService {
	mock := &MockConfigService{ctrl: ctrl}
	mock.recorder = &MockConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigService) EXPECT() *MockConfigServiceMockRecorder {
	return m.recorder
}

// EffectiveStoragePath mocks base method.
func (m *MockConfigService) EffectiveStoragePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EffectiveStoragePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// EffectiveStoragePath indicates an expected call of EffectiveStoragePath.
func (mr *MockConfigServiceMockRecorder) EffectiveStoragePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EffectiveStoragePath", reflect.TypeOf((*MockConfigService)(nil).EffectiveStoragePath))
}

// ExpandVariables mocks base method.
func (m *MockConfigService) ExpandVariables(arg0 string, arg1 config.VariableAudience, arg2 config.VariablePlatform) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpandVariables", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	return ret0
}

// ExpandVariables indicates an expected call of ExpandVariables.
func (mr *MockConfigServiceMockRecorder) ExpandVariables(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandVariables", reflect.TypeOf((*MockConfigService)(nil).ExpandVariables), arg0, arg1, arg2)
}

// MockTaskStateMachine is a mock of TaskStateMachine interface.
type MockTaskStateMachine struct {
	ctrl     *gomock.Controller
	recorder *MockTaskStateMachineMockRecorder
}

// MockTaskStateMachineMockRecorder is the mock recorder for MockTaskStateMachine.
type MockTaskStateMachineMockRecorder struct {
	mock *MockTaskStateMachine
}

// NewMockTaskStateMachine creates a new mock instance.
func NewMockTaskStateMachine(ctrl *gomock.Controller) *MockTaskStateMachine {
	mock := &MockTaskStateMachine{ctrl: ctrl}
	mock.recorder = &MockTaskStateMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskStateMachine) EXPECT() *MockTaskStateMachineMockRecorder {
	return m.recorder
}

// JobStatusChange mocks base method.
func (m *MockTaskStateMachine) JobStatusChange(arg0 context.Context, arg1 *persistence.Job, arg2 api.JobStatus, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JobStatusChange", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// JobStatusChange indicates an expected call of JobStatusChange.
func (mr *MockTaskStateMachineMockRecorder) JobStatusChange(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JobStatusChange", reflect.TypeOf((*MockTaskStateMachine)(nil).JobStatusChange), arg0, arg1, arg2, arg3)
}

// RequeueTasksOfWorker mocks base method.
func (m *MockTaskStateMachine) RequeueTasksOfWorker(arg0 context.Context, arg1 *persistence.Worker, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequeueTasksOfWorker", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequeueTasksOfWorker indicates an expected call of RequeueTasksOfWorker.
func (mr *MockTaskStateMachineMockRecorder) RequeueTasksOfWorker(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequeueTasksOfWorker", reflect.TypeOf((*MockTaskStateMachine)(nil).RequeueTasksOfWorker), arg0, arg1, arg2)
}

// TaskStatusChange mocks base method.
func (m *MockTaskStateMachine) TaskStatusChange(arg0 context.Context, arg1 *persistence.Task, arg2 api.TaskStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskStatusChange", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// TaskStatusChange indicates an expected call of TaskStatusChange.
func (mr *MockTaskStateMachineMockRecorder) TaskStatusChange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskStatusChange", reflect.TypeOf((*MockTaskStateMachine)(nil).TaskStatusChange), arg0, arg1, arg2)
}

// MockShaman is a mock of Shaman interface.
type MockShaman struct {
	ctrl     *gomock.Controller
	recorder *MockShamanMockRecorder
}

// MockShamanMockRecorder is the mock recorder for MockShaman.
type MockShamanMockRecorder struct {
	mock *MockShaman
}

// NewMockShaman creates a new mock instance.
func NewMockShaman(ctrl *gomock.Controller) *MockShaman {
	mock := &MockShaman{ctrl: ctrl}
	mock.recorder = &MockShamanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShaman) EXPECT() *MockShamanMockRecorder {
	return m.recorder
}

// Checkout mocks base method.
func (m *MockShaman) Checkout(arg0 context.Context, arg1 api.ShamanCheckout) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkout", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkout indicates an expected call of Checkout.
func (mr *MockShamanMockRecorder) Checkout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkout", reflect.TypeOf((*MockShaman)(nil).Checkout), arg0, arg1)
}

// FileStore mocks base method.
func (m *MockShaman) FileStore(arg0 context.Context, arg1 io.ReadCloser, arg2 string, arg3 int64, arg4 bool, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileStore", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// FileStore indicates an expected call of FileStore.
func (mr *MockShamanMockRecorder) FileStore(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileStore", reflect.TypeOf((*MockShaman)(nil).FileStore), arg0, arg1, arg2, arg3, arg4, arg5)
}

// FileStoreCheck mocks base method.
func (m *MockShaman) FileStoreCheck(arg0 context.Context, arg1 string, arg2 int64) api.ShamanFileStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileStoreCheck", arg0, arg1, arg2)
	ret0, _ := ret[0].(api.ShamanFileStatus)
	return ret0
}

// FileStoreCheck indicates an expected call of FileStoreCheck.
func (mr *MockShamanMockRecorder) FileStoreCheck(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileStoreCheck", reflect.TypeOf((*MockShaman)(nil).FileStoreCheck), arg0, arg1, arg2)
}

// IsEnabled mocks base method.
func (m *MockShaman) IsEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEnabled indicates an expected call of IsEnabled.
func (mr *MockShamanMockRecorder) IsEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEnabled", reflect.TypeOf((*MockShaman)(nil).IsEnabled))
}

// Requirements mocks base method.
func (m *MockShaman) Requirements(arg0 context.Context, arg1 api.ShamanRequirementsRequest) (api.ShamanRequirementsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Requirements", arg0, arg1)
	ret0, _ := ret[0].(api.ShamanRequirementsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Requirements indicates an expected call of Requirements.
func (mr *MockShamanMockRecorder) Requirements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Requirements", reflect.TypeOf((*MockShaman)(nil).Requirements), arg0, arg1)
}
