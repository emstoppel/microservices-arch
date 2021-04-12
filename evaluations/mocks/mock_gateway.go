// Code generated by MockGen. DO NOT EDIT.
// Source: gateway.go

// Package evaluationsmocks is a generated GoMock package.
package evaluationsmocks

import (
	context "context"
	reflect "reflect"

	evaluations "github.com/emstoppel/microservices-arch/evaluations"
	gomock "github.com/golang/mock/gomock"
)

// MockGateway is a mock of Gateway interface.
type MockGateway struct {
	ctrl     *gomock.Controller
	recorder *MockGatewayMockRecorder
}

// MockGatewayMockRecorder is the mock recorder for MockGateway.
type MockGatewayMockRecorder struct {
	mock *MockGateway
}

// NewMockGateway creates a new mock instance.
func NewMockGateway(ctrl *gomock.Controller) *MockGateway {
	mock := &MockGateway{ctrl: ctrl}
	mock.recorder = &MockGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGateway) EXPECT() *MockGatewayMockRecorder {
	return m.recorder
}

// CreateEvaluation mocks base method.
func (m *MockGateway) CreateEvaluation(ctx context.Context, evaluation evaluations.EvaluationData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateEvaluation", ctx, evaluation)
}

// CreateEvaluation indicates an expected call of CreateEvaluation.
func (mr *MockGatewayMockRecorder) CreateEvaluation(ctx, evaluation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvaluation", reflect.TypeOf((*MockGateway)(nil).CreateEvaluation), ctx, evaluation)
}

// GetEvaluations mocks base method.
func (m *MockGateway) GetEvaluations(ctx context.Context, tag string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetEvaluations", ctx, tag)
}

// GetEvaluations indicates an expected call of GetEvaluations.
func (mr *MockGatewayMockRecorder) GetEvaluations(ctx, tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvaluations", reflect.TypeOf((*MockGateway)(nil).GetEvaluations), ctx, tag)
}
