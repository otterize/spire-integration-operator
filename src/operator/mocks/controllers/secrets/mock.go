// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/otterize/spire-integration-operator/src/controllers/secrets (interfaces: Manager)

// Package mock_secrets is a generated GoMock package.
package mock_secrets

import (
	context "context"
	"github.com/otterize/spire-integration-operator/src/controllers/secrets/types"
	corev1 "k8s.io/api/core/v1"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// EnsureTLSSecret mocks base method.
func (m *MockManager) EnsureTLSSecret(arg0 context.Context, arg1 secretstypes.SecretConfig, arg2 *corev1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureTLSSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureTLSSecret indicates an expected call of EnsureTLSSecret.
func (mr *MockManagerMockRecorder) EnsureTLSSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureTLSSecret", reflect.TypeOf((*MockManager)(nil).EnsureTLSSecret), arg0, arg1, arg2)
}

// RefreshTLSSecrets mocks base method.
func (m *MockManager) RefreshTLSSecrets(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTLSSecrets", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshTLSSecrets indicates an expected call of RefreshTLSSecrets.
func (mr *MockManagerMockRecorder) RefreshTLSSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTLSSecrets", reflect.TypeOf((*MockManager)(nil).RefreshTLSSecrets), arg0)
}
