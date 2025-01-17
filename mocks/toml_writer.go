// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TOMLWriter is an autogenerated mock type for the TOMLWriter type
type TOMLWriter struct {
	mock.Mock
}

// Write provides a mock function with given fields: path, value
func (_m *TOMLWriter) Write(path string, value interface{}) error {
	ret := _m.Called(path, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(path, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
