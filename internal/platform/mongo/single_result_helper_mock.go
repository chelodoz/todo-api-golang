// Code generated by mockery v2.16.0. DO NOT EDIT.

package mongo

import mock "github.com/stretchr/testify/mock"

// MockSingleResultHelper is an autogenerated mock type for the SingleResultHelper type
type MockSingleResultHelper struct {
	mock.Mock
}

// Decode provides a mock function with given fields: v
func (_m *MockSingleResultHelper) Decode(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockSingleResultHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockSingleResultHelper creates a new instance of MockSingleResultHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockSingleResultHelper(t mockConstructorTestingTNewMockSingleResultHelper) *MockSingleResultHelper {
	mock := &MockSingleResultHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
