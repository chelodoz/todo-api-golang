// Code generated by mockery v2.16.0. DO NOT EDIT.

package note

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// Create provides a mock function with given fields: note, ctx
func (_m *MockService) Create(note *Note, ctx context.Context) (*Note, error) {
	ret := _m.Called(note, ctx)

	var r0 *Note
	if rf, ok := ret.Get(0).(func(*Note, context.Context) *Note); ok {
		r0 = rf(note, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Note)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Note, context.Context) error); ok {
		r1 = rf(note, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockService) GetAll(ctx context.Context) ([]Note, error) {
	ret := _m.Called(ctx)

	var r0 []Note
	if rf, ok := ret.Get(0).(func(context.Context) []Note); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Note)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id, ctx
func (_m *MockService) GetById(id uuid.UUID, ctx context.Context) (*Note, error) {
	ret := _m.Called(id, ctx)

	var r0 *Note
	if rf, ok := ret.Get(0).(func(uuid.UUID, context.Context) *Note); ok {
		r0 = rf(id, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Note)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, context.Context) error); ok {
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: note, ctx
func (_m *MockService) Update(note *Note, ctx context.Context) (*Note, error) {
	ret := _m.Called(note, ctx)

	var r0 *Note
	if rf, ok := ret.Get(0).(func(*Note, context.Context) *Note); ok {
		r0 = rf(note, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Note)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*Note, context.Context) error); ok {
		r1 = rf(note, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockService(t mockConstructorTestingTNewMockService) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
