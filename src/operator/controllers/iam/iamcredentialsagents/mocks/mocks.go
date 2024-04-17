// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock_iamcredentialsagents is a generated GoMock package.
package mock_iamcredentialsagents

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockIAMCredentialsAgent is a mock of IAMCredentialsAgent interface.
type MockIAMCredentialsAgent struct {
	ctrl     *gomock.Controller
	recorder *MockIAMCredentialsAgentMockRecorder
}

// MockIAMCredentialsAgentMockRecorder is the mock recorder for MockIAMCredentialsAgent.
type MockIAMCredentialsAgentMockRecorder struct {
	mock *MockIAMCredentialsAgent
}

// NewMockIAMCredentialsAgent creates a new mock instance.
func NewMockIAMCredentialsAgent(ctrl *gomock.Controller) *MockIAMCredentialsAgent {
	mock := &MockIAMCredentialsAgent{ctrl: ctrl}
	mock.recorder = &MockIAMCredentialsAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAMCredentialsAgent) EXPECT() *MockIAMCredentialsAgentMockRecorder {
	return m.recorder
}

// AppliesOnPod mocks base method.
func (m *MockIAMCredentialsAgent) AppliesOnPod(pod *v1.Pod) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppliesOnPod", pod)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AppliesOnPod indicates an expected call of AppliesOnPod.
func (mr *MockIAMCredentialsAgentMockRecorder) AppliesOnPod(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppliesOnPod", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).AppliesOnPod), pod)
}

// FinalizerName mocks base method.
func (m *MockIAMCredentialsAgent) FinalizerName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizerName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FinalizerName indicates an expected call of FinalizerName.
func (mr *MockIAMCredentialsAgentMockRecorder) FinalizerName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizerName", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).FinalizerName))
}

// OnPodAdmission mocks base method.
func (m *MockIAMCredentialsAgent) OnPodAdmission(ctx context.Context, pod *v1.Pod, serviceAccount *v1.ServiceAccount, dryRun bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnPodAdmission", ctx, pod, serviceAccount, dryRun)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnPodAdmission indicates an expected call of OnPodAdmission.
func (mr *MockIAMCredentialsAgentMockRecorder) OnPodAdmission(ctx, pod, serviceAccount, dryRun interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPodAdmission", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).OnPodAdmission), ctx, pod, serviceAccount, dryRun)
}

// OnPodUpdate mocks base method.
func (m *MockIAMCredentialsAgent) OnPodUpdate(ctx context.Context, pod *v1.Pod, serviceAccount *v1.ServiceAccount) (bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnPodUpdate", ctx, pod, serviceAccount)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OnPodUpdate indicates an expected call of OnPodUpdate.
func (mr *MockIAMCredentialsAgentMockRecorder) OnPodUpdate(ctx, pod, serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnPodUpdate", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).OnPodUpdate), ctx, pod, serviceAccount)
}

// OnServiceAccountTermination mocks base method.
func (m *MockIAMCredentialsAgent) OnServiceAccountTermination(ctx context.Context, serviceAccount *v1.ServiceAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnServiceAccountTermination", ctx, serviceAccount)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnServiceAccountTermination indicates an expected call of OnServiceAccountTermination.
func (mr *MockIAMCredentialsAgentMockRecorder) OnServiceAccountTermination(ctx, serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnServiceAccountTermination", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).OnServiceAccountTermination), ctx, serviceAccount)
}

// OnServiceAccountUpdate mocks base method.
func (m *MockIAMCredentialsAgent) OnServiceAccountUpdate(ctx context.Context, serviceAccount *v1.ServiceAccount) (bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnServiceAccountUpdate", ctx, serviceAccount)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// OnServiceAccountUpdate indicates an expected call of OnServiceAccountUpdate.
func (mr *MockIAMCredentialsAgentMockRecorder) OnServiceAccountUpdate(ctx, serviceAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnServiceAccountUpdate", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).OnServiceAccountUpdate), ctx, serviceAccount)
}

// ServiceAccountLabel mocks base method.
func (m *MockIAMCredentialsAgent) ServiceAccountLabel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceAccountLabel")
	ret0, _ := ret[0].(string)
	return ret0
}

// ServiceAccountLabel indicates an expected call of ServiceAccountLabel.
func (mr *MockIAMCredentialsAgentMockRecorder) ServiceAccountLabel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceAccountLabel", reflect.TypeOf((*MockIAMCredentialsAgent)(nil).ServiceAccountLabel))
}