// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	patient "clinic-api/src/app/patient"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AmendPatientByID provides a mock function with given fields: id, domain
func (_m *Services) AmendPatientByID(id string, domain patient.Domain) error {
	ret := _m.Called(id, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, patient.Domain) error); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePatient provides a mock function with given fields: domain
func (_m *Services) CreatePatient(domain patient.Domain) (string, error) {
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

// GetPatientByID provides a mock function with given fields: id
func (_m *Services) GetPatientByID(id string) (*patient.Domain, error) {
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

// HuntPatientByNameOrNIKOrAll provides a mock function with given fields: domain
func (_m *Services) HuntPatientByNameOrNIKOrAll(domain patient.Domain) ([]patient.Domain, error) {
	ret := _m.Called(domain)

	var r0 []patient.Domain
	if rf, ok := ret.Get(0).(func(patient.Domain) []patient.Domain); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patient.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(patient.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePatientByID provides a mock function with given fields: id
func (_m *Services) RemovePatientByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
