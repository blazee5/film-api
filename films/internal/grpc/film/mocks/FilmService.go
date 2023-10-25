// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	films_api "github.com/blazee5/film-api/api/proto/v1"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"
)

// FilmService is an autogenerated mock type for the FilmService type
type FilmService struct {
	mock.Mock
}

// CreateFilm provides a mock function with given fields: db, in
func (_m *FilmService) CreateFilm(db *gorm.DB, in *films_api.Film) (int64, error) {
	ret := _m.Called(db, in)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.Film) (int64, error)); ok {
		return rf(db, in)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.Film) int64); ok {
		r0 = rf(db, in)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *films_api.Film) error); ok {
		r1 = rf(db, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFilm provides a mock function with given fields: db, in
func (_m *FilmService) DeleteFilm(db *gorm.DB, in *films_api.FilmRequest) error {
	ret := _m.Called(db, in)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.FilmRequest) error); ok {
		r0 = rf(db, in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFilm provides a mock function with given fields: db, in
func (_m *FilmService) GetFilm(db *gorm.DB, in *films_api.FilmRequest) (*films_api.Film, error) {
	ret := _m.Called(db, in)

	var r0 *films_api.Film
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.FilmRequest) (*films_api.Film, error)); ok {
		return rf(db, in)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.FilmRequest) *films_api.Film); ok {
		r0 = rf(db, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*films_api.Film)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *films_api.FilmRequest) error); ok {
		r1 = rf(db, in)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateFilm provides a mock function with given fields: db, in
func (_m *FilmService) UpdateFilm(db *gorm.DB, in *films_api.Film) (*films_api.Film, error) {
	ret := _m.Called(db, in)

	var r0 *films_api.Film
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.Film) (*films_api.Film, error)); ok {
		return rf(db, in)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, *films_api.Film) *films_api.Film); ok {
		r0 = rf(db, in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*films_api.Film)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, *films_api.Film) error); ok {
		r1 = rf(db, in)
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
