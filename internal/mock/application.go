// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	dto "github.com/quizardhq/golang-ddd-boilerplate/internal/entrypoint/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockPersonApplication is a mock of PersonApplication interface.
type MockPersonApplication struct {
	ctrl     *gomock.Controller
	recorder *MockPersonApplicationMockRecorder
}

// MockPersonApplicationMockRecorder is the mock recorder for MockPersonApplication.
type MockPersonApplicationMockRecorder struct {
	mock *MockPersonApplication
}

// NewMockPersonApplication creates a new mock instance.
func NewMockPersonApplication(ctrl *gomock.Controller) *MockPersonApplication {
	mock := &MockPersonApplication{ctrl: ctrl}
	mock.recorder = &MockPersonApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonApplication) EXPECT() *MockPersonApplicationMockRecorder {
	return m.recorder
}

// CreatePerson mocks base method.
func (m *MockPersonApplication) CreatePerson(arg0 dto.PersonDTO) dto.PersonDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerson", arg0)
	ret0, _ := ret[0].(dto.PersonDTO)
	return ret0
}

// CreatePerson indicates an expected call of CreatePerson.
func (mr *MockPersonApplicationMockRecorder) CreatePerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerson", reflect.TypeOf((*MockPersonApplication)(nil).CreatePerson), arg0)
}

// DeletePerson mocks base method.
func (m *MockPersonApplication) DeletePerson(arg0 int32) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePerson", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeletePerson indicates an expected call of DeletePerson.
func (mr *MockPersonApplicationMockRecorder) DeletePerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePerson", reflect.TypeOf((*MockPersonApplication)(nil).DeletePerson), arg0)
}

// FindAll mocks base method.
func (m *MockPersonApplication) FindAll() []dto.PersonDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]dto.PersonDTO)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockPersonApplicationMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockPersonApplication)(nil).FindAll))
}

// FindOne mocks base method.
func (m *MockPersonApplication) FindOne(arg0 int32) dto.PersonDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", arg0)
	ret0, _ := ret[0].(dto.PersonDTO)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockPersonApplicationMockRecorder) FindOne(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockPersonApplication)(nil).FindOne), arg0)
}

// UpdatePerson mocks base method.
func (m *MockPersonApplication) UpdatePerson(arg0 dto.PersonDTO) dto.PersonDTO {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePerson", arg0)
	ret0, _ := ret[0].(dto.PersonDTO)
	return ret0
}

// UpdatePerson indicates an expected call of UpdatePerson.
func (mr *MockPersonApplicationMockRecorder) UpdatePerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePerson", reflect.TypeOf((*MockPersonApplication)(nil).UpdatePerson), arg0)
}
