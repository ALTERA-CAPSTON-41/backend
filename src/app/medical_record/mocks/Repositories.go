// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	medicalrecord "clinic-api/src/app/medical_record"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Repositories is an autogenerated mock type for the Repositories type
type Repositories struct {
	mock.Mock
}

// InsertData provides a mock function with given fields: domain
func (_m *Repositories) InsertData(domain medicalrecord.Domain) (string, error) {
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

// LookupICD10Data provides a mock function with given fields: icd10Code
func (_m *Repositories) LookupICD10Data(icd10Code string) (string, error) {
	ret := _m.Called(icd10Code)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(icd10Code)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(icd10Code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectDataByID provides a mock function with given fields: id
func (_m *Repositories) SelectDataByID(id string) (*medicalrecord.Domain, error) {
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

// SelectDataByPatientID provides a mock function with given fields: id
func (_m *Repositories) SelectDataByPatientID(id string) ([]medicalrecord.Domain, error) {
	ret := _m.Called(id)

	var r0 []medicalrecord.Domain
	if rf, ok := ret.Get(0).(func(string) []medicalrecord.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]medicalrecord.Domain)
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

// SelectPatientIDByNIK provides a mock function with given fields: nik
func (_m *Repositories) SelectPatientIDByNIK(nik string) (string, error) {
	ret := _m.Called(nik)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(nik)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(nik)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepositories creates a new instance of Repositories. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepositories(t testing.TB) *Repositories {
	mock := &Repositories{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
