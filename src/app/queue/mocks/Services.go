// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	queue "clinic-api/src/app/queue"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AmendQueueByID provides a mock function with given fields: id, _a1
func (_m *Services) AmendQueueByID(id string, _a1 queue.Domain) error {
	ret := _m.Called(id, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, queue.Domain) error); ok {
		r0 = rf(id, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateQueue provides a mock function with given fields: _a0
func (_m *Services) CreateQueue(_a0 queue.Domain) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(queue.Domain) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(queue.Domain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllQueues provides a mock function with given fields: fromDate, polyclinic, page
func (_m *Services) GetAllQueues(fromDate string, polyclinic int, page int) ([]queue.Domain, error) {
	ret := _m.Called(fromDate, polyclinic, page)

	var r0 []queue.Domain
	if rf, ok := ret.Get(0).(func(string, int, int) []queue.Domain); ok {
		r0 = rf(fromDate, polyclinic, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]queue.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(fromDate, polyclinic, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveQueueByID provides a mock function with given fields: id
func (_m *Services) RemoveQueueByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewServices creates a new instance of Services. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewServices(t testing.TB) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
