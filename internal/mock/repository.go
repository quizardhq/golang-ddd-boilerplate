// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/repository/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	entity "github.com/quizardhq/golang-ddd-boilerplate/internal/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockPersonRepository is a mock of PersonRepository interface.
type MockPersonRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPersonRepositoryMockRecorder
}

// MockPersonRepositoryMockRecorder is the mock recorder for MockPersonRepository.
type MockPersonRepositoryMockRecorder struct {
	mock *MockPersonRepository
}

// NewMockPersonRepository creates a new mock instance.
func NewMockPersonRepository(ctrl *gomock.Controller) *MockPersonRepository {
	mock := &MockPersonRepository{ctrl: ctrl}
	mock.recorder = &MockPersonRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonRepository) EXPECT() *MockPersonRepositoryMockRecorder {
	return m.recorder
}

// CreatePerson mocks base method.
func (m *MockPersonRepository) CreatePerson(arg0 entity.Person) entity.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerson", arg0)
	ret0, _ := ret[0].(entity.Person)
	return ret0
}

// CreatePerson indicates an expected call of CreatePerson.
func (mr *MockPersonRepositoryMockRecorder) CreatePerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerson", reflect.TypeOf((*MockPersonRepository)(nil).CreatePerson), arg0)
}

// DeletePerson mocks base method.
func (m *MockPersonRepository) DeletePerson(arg0 int32) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePerson", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeletePerson indicates an expected call of DeletePerson.
func (mr *MockPersonRepositoryMockRecorder) DeletePerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePerson", reflect.TypeOf((*MockPersonRepository)(nil).DeletePerson), arg0)
}

// GetPeople mocks base method.
func (m *MockPersonRepository) GetPeople() []entity.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeople")
	ret0, _ := ret[0].([]entity.Person)
	return ret0
}

// GetPeople indicates an expected call of GetPeople.
func (mr *MockPersonRepositoryMockRecorder) GetPeople() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeople", reflect.TypeOf((*MockPersonRepository)(nil).GetPeople))
}

// GetPerson mocks base method.
func (m *MockPersonRepository) GetPerson(arg0 int32) entity.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPerson", arg0)
	ret0, _ := ret[0].(entity.Person)
	return ret0
}

// GetPerson indicates an expected call of GetPerson.
func (mr *MockPersonRepositoryMockRecorder) GetPerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerson", reflect.TypeOf((*MockPersonRepository)(nil).GetPerson), arg0)
}

// UpdatePerson mocks base method.
func (m *MockPersonRepository) UpdatePerson(arg0 int32, arg1 entity.Person) entity.Person {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePerson", arg0, arg1)
	ret0, _ := ret[0].(entity.Person)
	return ret0
}

// UpdatePerson indicates an expected call of UpdatePerson.
func (mr *MockPersonRepositoryMockRecorder) UpdatePerson(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePerson", reflect.TypeOf((*MockPersonRepository)(nil).UpdatePerson), arg0, arg1)
}
