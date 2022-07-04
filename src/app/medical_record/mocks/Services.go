// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	medicalrecord "clinic-api/src/app/medical_record"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// CreateMedicalRecord provides a mock function with given fields: domain
func (_m *Services) CreateMedicalRecord(domain medicalrecord.Domain) (string, error) {
	ret := _m.Called(domain)

	var r0 string
	if rf, ok := ret.Get(0).(func(medicalrecord.Domain) string); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(medicalrecord.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMedicalRecordByID provides a mock function with given fields: id
func (_m *Services) FindMedicalRecordByID(id string) (*medicalrecord.Domain, error) {
	ret := _m.Called(id)

	var r0 *medicalrecord.Domain
	if rf, ok := ret.Get(0).(func(string) *medicalrecord.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*medicalrecord.Domain)
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

// FindMedicalRecordByPatientNIK provides a mock function with given fields: nik
func (_m *Services) FindMedicalRecordByPatientNIK(nik string) ([]medicalrecord.Domain, error) {
	ret := _m.Called(nik)

	var r0 []medicalrecord.Domain
	if rf, ok := ret.Get(0).(func(string) []medicalrecord.Domain); ok {
		r0 = rf(nik)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]medicalrecord.Domain)
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

// NewServices creates a new instance of Services. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewServices(t testing.TB) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
