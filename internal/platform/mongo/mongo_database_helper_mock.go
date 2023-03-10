// Code generated by mockery v2.16.0. DO NOT EDIT.

package mongo

import mock "github.com/stretchr/testify/mock"

// MockDatabaseHelper is an autogenerated mock type for the DatabaseHelper type
type MockDatabaseHelper struct {
	mock.Mock
}

// Client provides a mock function with given fields:
func (_m *MockDatabaseHelper) Client() ClientHelper {
	ret := _m.Called()

	var r0 ClientHelper
	if rf, ok := ret.Get(0).(func() ClientHelper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ClientHelper)
		}
	}

	return r0
}

// Collection provides a mock function with given fields: name
func (_m *MockDatabaseHelper) Collection(name string) CollectionHelper {
	ret := _m.Called(name)

	var r0 CollectionHelper
	if rf, ok := ret.Get(0).(func(string) CollectionHelper); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(CollectionHelper)
		}
	}

	return r0
}

type mockConstructorTestingTNewMockDatabaseHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDatabaseHelper creates a new instance of MockDatabaseHelper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDatabaseHelper(t mockConstructorTestingTNewMockDatabaseHelper) *MockDatabaseHelper {
	mock := &MockDatabaseHelper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
