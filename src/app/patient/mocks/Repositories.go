// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	patient "clinic-api/src/app/patient"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Repositories is an autogenerated mock type for the Repositories type
type Repositories struct {
	mock.Mock
}

// DeleteByID provides a mock function with given fields: id
func (_m *Repositories) DeleteByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertData provides a mock function with given fields: domain
func (_m *Repositories) InsertData(domain patient.Domain) (string, error) {
	ret := _m.Called(domain)

	var r0 string
	if rf, ok := ret.Get(0).(func(patient.Domain) string); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(patient.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchDataByNIKParam provides a mock function with given fields: nik
func (_m *Repositories) SearchDataByNIKParam(nik string) ([]patient.Domain, error) {
	ret := _m.Called(nik)

	var r0 []patient.Domain
	if rf, ok := ret.Get(0).(func(string) []patient.Domain); ok {
		r0 = rf(nik)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patient.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nik)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchDataByNameParam provides a mock function with given fields: name, offset
func (_m *Repositories) SearchDataByNameParam(name string, offset int) ([]patient.Domain, error) {
	ret := _m.Called(name, offset)

	var r0 []patient.Domain
	if rf, ok := ret.Get(0).(func(string, int) []patient.Domain); ok {
		r0 = rf(name, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patient.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int) error); ok {
		r1 = rf(name, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllData provides a mock function with given fields: offset
func (_m *Repositories) SelectAllData(offset int) ([]patient.Domain, error) {
	ret := _m.Called(offset)

	var r0 []patient.Domain
	if rf, ok := ret.Get(0).(func(int) []patient.Domain); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patient.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDataByID provides a mock function with given fields: id
func (_m *Repositories) SelectDataByID(id string) (*patient.Domain, error) {
	ret := _m.Called(id)

	var r0 *patient.Domain
	if rf, ok := ret.Get(0).(func(string) *patient.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*patient.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: id, domain
func (_m *Repositories) UpdateByID(id string, domain patient.Domain) error {
	ret := _m.Called(id, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, patient.Domain) error); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepositories creates a new instance of Repositories. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositories(t testing.TB) *Repositories {
	mock := &Repositories{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
