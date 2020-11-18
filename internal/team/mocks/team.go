// Code generated by MockGen. DO NOT EDIT.
// Source: internal/team/team.go

// Package mock_team is a generated GoMock package.
package mock_team

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	entity "soccer-api/internal/entity"
)

// MockTeam is a mock of Team interface
type MockTeam struct {
	ctrl     *gomock.Controller
	recorder *MockTeamMockRecorder
}

// MockTeamMockRecorder is the mock recorder for MockTeam
type MockTeamMockRecorder struct {
	mock *MockTeam
}

// NewMockTeam creates a new mock instance
func NewMockTeam(ctrl *gomock.Controller) *MockTeam {
	mock := &MockTeam{ctrl: ctrl}
	mock.recorder = &MockTeamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTeam) EXPECT() *MockTeamMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockTeam) Fetch(ctx context.Context, filter entity.QueryFilter) ([]entity.Team, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, filter)
	ret0, _ := ret[0].([]entity.Team)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Fetch indicates an expected call of Fetch
func (mr *MockTeamMockRecorder) Fetch(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockTeam)(nil).Fetch), ctx, filter)
}

// Get mocks base method
func (m *MockTeam) Get(ctx context.Context, ID int64) (entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, ID)
	ret0, _ := ret[0].(entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTeamMockRecorder) Get(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTeam)(nil).Get), ctx, ID)
}

// Insert mocks base method
func (m *MockTeam) Insert(ctx context.Context, t entity.Team) (entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, t)
	ret0, _ := ret[0].(entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockTeamMockRecorder) Insert(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockTeam)(nil).Insert), ctx, t)
}
