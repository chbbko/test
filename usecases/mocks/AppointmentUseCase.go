// Code generated by mockery v2.6.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "github.com/test/entities"
)

// AppointmentUseCase is an autogenerated mock type for the AppointmentUseCase type
type AppointmentUseCase struct {
	mock.Mock
}

// Cancel provides a mock function with given fields: _a0
func (_m *AppointmentUseCase) Cancel(_a0 entities.AppointmentDelete) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.AppointmentDelete) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0
func (_m *AppointmentUseCase) Create(_a0 entities.AppointmentCreate) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(entities.AppointmentCreate) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: _a0
func (_m *AppointmentUseCase) List(_a0 entities.AppointmentListParams) ([]entities.AppointmentList, error) {
	ret := _m.Called(_a0)

	var r0 []entities.AppointmentList
	if rf, ok := ret.Get(0).(func(entities.AppointmentListParams) []entities.AppointmentList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.AppointmentList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.AppointmentListParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}