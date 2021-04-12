// Code generated by MockGen. DO NOT EDIT.
// Source: gateway.go

// Package simulationsmocks is a generated GoMock package.
package simulationsmocks

import (
	context "context"
	reflect "reflect"

	simulations "github.com/emstoppel/microservices-arch/simulations"
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

// GetSimulation mocks base method.
func (m *MockGateway) GetSimulation(ctx context.Context, simulationID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetSimulation", ctx, simulationID)
}

// GetSimulation indicates an expected call of GetSimulation.
func (mr *MockGatewayMockRecorder) GetSimulation(ctx, simulationID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSimulation", reflect.TypeOf((*MockGateway)(nil).GetSimulation), ctx, simulationID)
}

// SaveSimulation mocks base method.
func (m *MockGateway) SaveSimulation(ctx context.Context, request simulations.SimulationSetup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SaveSimulation", ctx, request)
}

// SaveSimulation indicates an expected call of SaveSimulation.
func (mr *MockGatewayMockRecorder) SaveSimulation(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSimulation", reflect.TypeOf((*MockGateway)(nil).SaveSimulation), ctx, request)
}
