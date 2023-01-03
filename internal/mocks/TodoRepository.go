// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "todo-api-golang/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: _a0, ctx
func (_m *TodoRepository) CreateTodo(_a0 *entity.Todo, ctx context.Context) (*entity.Todo, error) {
	ret := _m.Called(_a0, ctx)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(*entity.Todo, context.Context) *entity.Todo); ok {
		r0 = rf(_a0, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Todo, context.Context) error); ok {
		r1 = rf(_a0, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodoById provides a mock function with given fields: id, ctx
func (_m *TodoRepository) GetTodoById(id uint, ctx context.Context) (*entity.Todo, error) {
	ret := _m.Called(id, ctx)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(uint, context.Context) *entity.Todo); ok {
		r0 = rf(id, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context) error); ok {
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTodos provides a mock function with given fields: ctx
func (_m *TodoRepository) GetTodos(ctx context.Context) ([]entity.Todo, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Todo); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Todo)
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

// UpdateTodo provides a mock function with given fields: _a0, ctx
func (_m *TodoRepository) UpdateTodo(_a0 *entity.Todo, ctx context.Context) (*entity.Todo, error) {
	ret := _m.Called(_a0, ctx)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(*entity.Todo, context.Context) *entity.Todo); ok {
		r0 = rf(_a0, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Todo, context.Context) error); ok {
		r1 = rf(_a0, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTodoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTodoRepository(t mockConstructorTestingTNewTodoRepository) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
