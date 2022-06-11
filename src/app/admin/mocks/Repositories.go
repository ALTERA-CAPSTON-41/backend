// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	admin "clinic-api/src/app/admin"

	mock "github.com/stretchr/testify/mock"
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

// InsertData provides a mock function with given fields: data
func (_m *Repositories) InsertData(data admin.Domain) (string, error) {
	ret := _m.Called(data)

	var r0 string
	if rf, ok := ret.Get(0).(func(admin.Domain) string); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(admin.Domain) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAllData provides a mock function with given fields:
func (_m *Repositories) SelectAllData() ([]admin.Domain, error) {
	ret := _m.Called()

	var r0 []admin.Domain
	if rf, ok := ret.Get(0).(func() []admin.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admin.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDataByID provides a mock function with given fields: id
func (_m *Repositories) SelectDataByID(id string) (*admin.Domain, error) {
	ret := _m.Called(id)

	var r0 *admin.Domain
	if rf, ok := ret.Get(0).(func(string) *admin.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Domain)
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

// UpdateByID provides a mock function with given fields: id, data
func (_m *Repositories) UpdateByID(id string, data admin.Domain) error {
	ret := _m.Called(id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, admin.Domain) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
