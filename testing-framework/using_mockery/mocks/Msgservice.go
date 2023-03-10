// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Msgservice is an autogenerated mock type for the Msgservice type
type Msgservice struct {
	mock.Mock
}

// Sendmsg provides a mock function with given fields: msg
func (_m *Msgservice) Sendmsg(msg string) bool {
	ret := _m.Called(msg)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewMsgservice interface {
	mock.TestingT
	Cleanup(func())
}

// NewMsgservice creates a new instance of Msgservice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMsgservice(t mockConstructorTestingTNewMsgservice) *Msgservice {
	mock := &Msgservice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
