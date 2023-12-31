// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/blazee5/film-api/films/api/proto/v1"
)

// FilmService is an autogenerated mock type for the FilmService type
type FilmService struct {
	mock.Mock
}

// CreateFilm provides a mock function with given fields: ctx, db, in
func (_m *FilmService) CreateFilm(ctx context.Context, db *gorm.DB, in *v1.Film) (int64, error) {
	ret := _m.Called(ctx, db, in)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.Film) (int64, error)); ok {
		return rf(ctx, db, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.Film) int64); ok {
		r0 = rf(ctx, db, in)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *v1.Film) error); ok {
		r1 = rf(ctx, db, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFilm provides a mock function with given fields: ctx, db, in
func (_m *FilmService) DeleteFilm(ctx context.Context, db *gorm.DB, in *v1.FilmRequest) error {
	ret := _m.Called(ctx, db, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.FilmRequest) error); ok {
		r0 = rf(ctx, db, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFilm provides a mock function with given fields: ctx, db, in
func (_m *FilmService) GetFilm(ctx context.Context, db *gorm.DB, in *v1.FilmRequest) (*v1.Film, error) {
	ret := _m.Called(ctx, db, in)

	var r0 *v1.Film
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.FilmRequest) (*v1.Film, error)); ok {
		return rf(ctx, db, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.FilmRequest) *v1.Film); ok {
		r0 = rf(ctx, db, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Film)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *v1.FilmRequest) error); ok {
		r1 = rf(ctx, db, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFilm provides a mock function with given fields: ctx, db, in
func (_m *FilmService) UpdateFilm(ctx context.Context, db *gorm.DB, in *v1.Film) (*v1.Film, error) {
	ret := _m.Called(ctx, db, in)

	var r0 *v1.Film
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.Film) (*v1.Film, error)); ok {
		return rf(ctx, db, in)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *gorm.DB, *v1.Film) *v1.Film); ok {
		r0 = rf(ctx, db, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Film)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *gorm.DB, *v1.Film) error); ok {
		r1 = rf(ctx, db, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFilmService creates a new instance of FilmService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFilmService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FilmService {
	mock := &FilmService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
