// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	doctor "clinic-api/src/app/doctor"

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

// DeleteUserByID provides a mock function with given fields: id
func (_m *Repositories) DeleteUserByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertData provides a mock function with given fields: data
func (_m *Repositories) InsertData(data doctor.Domain) (string, error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(doctor.Domain) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(doctor.Domain) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupDataByEmail provides a mock function with given fields: email
func (_m *Repositories) LookupDataByEmail(email string) (string, error) {
	ret := _m.Called(email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllData provides a mock function with given fields: polyclinic, offset
func (_m *Repositories) SelectAllData(polyclinic int, offset int) ([]doctor.Domain, error) {
	ret := _m.Called(polyclinic, offset)

	var r0 []doctor.Domain
	if rf, ok := ret.Get(0).(func(int, int) []doctor.Domain); ok {
		r0 = rf(polyclinic, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctor.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(polyclinic, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDataByID provides a mock function with given fields: id
func (_m *Repositories) SelectDataByID(id string) (*doctor.Domain, error) {
	ret := _m.Called(id)

	var r0 *doctor.Domain
	if rf, ok := ret.Get(0).(func(string) *doctor.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*doctor.Domain)
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
func (_m *Repositories) UpdateByID(id string, domain doctor.Domain) error {
	ret := _m.Called(id, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, doctor.Domain) error); ok {
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
