// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	users "github.com/absmach/magistrala/users"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ChangeStatus provides a mock function with given fields: ctx, user
func (_m *Repository) ChangeStatus(ctx context.Context, user users.User) (users.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for ChangeStatus")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.User) (users.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.User) users.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckSuperAdmin provides a mock function with given fields: ctx, adminID
func (_m *Repository) CheckSuperAdmin(ctx context.Context, adminID string) error {
	ret := _m.Called(ctx, adminID)

	if len(ret) == 0 {
		panic("no return value specified for CheckSuperAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, adminID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAll provides a mock function with given fields: ctx, pm
func (_m *Repository) RetrieveAll(ctx context.Context, pm users.Page) (users.UsersPage, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAll")
	}

	var r0 users.UsersPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.Page) (users.UsersPage, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.Page) users.UsersPage); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(users.UsersPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.Page) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAllByIDs provides a mock function with given fields: ctx, pm
func (_m *Repository) RetrieveAllByIDs(ctx context.Context, pm users.Page) (users.UsersPage, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllByIDs")
	}

	var r0 users.UsersPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.Page) (users.UsersPage, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.Page) users.UsersPage); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(users.UsersPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.Page) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByEmail provides a mock function with given fields: ctx, email
func (_m *Repository) RetrieveByEmail(ctx context.Context, email string) (users.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByEmail")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (users.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) users.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByID provides a mock function with given fields: ctx, id
func (_m *Repository) RetrieveByID(ctx context.Context, id string) (users.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByID")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (users.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) users.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByUserName provides a mock function with given fields: ctx, userName
func (_m *Repository) RetrieveByUserName(ctx context.Context, userName string) (users.User, error) {
	ret := _m.Called(ctx, userName)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByUserName")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (users.User, error)); ok {
		return rf(ctx, userName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) users.User); ok {
		r0 = rf(ctx, userName)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, user
func (_m *Repository) Save(ctx context.Context, user users.User) (users.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.User) (users.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.User) users.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUsers provides a mock function with given fields: ctx, pm
func (_m *Repository) SearchUsers(ctx context.Context, pm users.Page) (users.UsersPage, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for SearchUsers")
	}

	var r0 users.UsersPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.Page) (users.UsersPage, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.Page) users.UsersPage); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(users.UsersPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.Page) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, user
func (_m *Repository) Update(ctx context.Context, user users.User) (users.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.User) (users.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.User) users.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecret provides a mock function with given fields: ctx, user
func (_m *Repository) UpdateSecret(ctx context.Context, user users.User) (users.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSecret")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.User) (users.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.User) users.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserName provides a mock function with given fields: ctx, user
func (_m *Repository) UpdateUserName(ctx context.Context, user users.User) (users.User, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserName")
	}

	var r0 users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, users.User) (users.User, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, users.User) users.User); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(users.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, users.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
