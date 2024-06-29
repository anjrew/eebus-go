// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// TimeSeriesCommonInterface is an autogenerated mock type for the TimeSeriesCommonInterface type
type TimeSeriesCommonInterface struct {
	mock.Mock
}

type TimeSeriesCommonInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *TimeSeriesCommonInterface) EXPECT() *TimeSeriesCommonInterface_Expecter {
	return &TimeSeriesCommonInterface_Expecter{mock: &_m.Mock}
}

// GetConstraints provides a mock function with given fields:
func (_m *TimeSeriesCommonInterface) GetConstraints() ([]model.TimeSeriesConstraintsDataType, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConstraints")
	}

	var r0 []model.TimeSeriesConstraintsDataType
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.TimeSeriesConstraintsDataType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.TimeSeriesConstraintsDataType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TimeSeriesConstraintsDataType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesCommonInterface_GetConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConstraints'
type TimeSeriesCommonInterface_GetConstraints_Call struct {
	*mock.Call
}

// GetConstraints is a helper method to define mock.On call
func (_e *TimeSeriesCommonInterface_Expecter) GetConstraints() *TimeSeriesCommonInterface_GetConstraints_Call {
	return &TimeSeriesCommonInterface_GetConstraints_Call{Call: _e.mock.On("GetConstraints")}
}

func (_c *TimeSeriesCommonInterface_GetConstraints_Call) Run(run func()) *TimeSeriesCommonInterface_GetConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TimeSeriesCommonInterface_GetConstraints_Call) Return(_a0 []model.TimeSeriesConstraintsDataType, _a1 error) *TimeSeriesCommonInterface_GetConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesCommonInterface_GetConstraints_Call) RunAndReturn(run func() ([]model.TimeSeriesConstraintsDataType, error)) *TimeSeriesCommonInterface_GetConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataForFilter provides a mock function with given fields: filter
func (_m *TimeSeriesCommonInterface) GetDataForFilter(filter model.TimeSeriesDescriptionDataType) ([]model.TimeSeriesDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetDataForFilter")
	}

	var r0 []model.TimeSeriesDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.TimeSeriesDescriptionDataType) ([]model.TimeSeriesDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.TimeSeriesDescriptionDataType) []model.TimeSeriesDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TimeSeriesDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.TimeSeriesDescriptionDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesCommonInterface_GetDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataForFilter'
type TimeSeriesCommonInterface_GetDataForFilter_Call struct {
	*mock.Call
}

// GetDataForFilter is a helper method to define mock.On call
//   - filter model.TimeSeriesDescriptionDataType
func (_e *TimeSeriesCommonInterface_Expecter) GetDataForFilter(filter interface{}) *TimeSeriesCommonInterface_GetDataForFilter_Call {
	return &TimeSeriesCommonInterface_GetDataForFilter_Call{Call: _e.mock.On("GetDataForFilter", filter)}
}

func (_c *TimeSeriesCommonInterface_GetDataForFilter_Call) Run(run func(filter model.TimeSeriesDescriptionDataType)) *TimeSeriesCommonInterface_GetDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.TimeSeriesDescriptionDataType))
	})
	return _c
}

func (_c *TimeSeriesCommonInterface_GetDataForFilter_Call) Return(_a0 []model.TimeSeriesDataType, _a1 error) *TimeSeriesCommonInterface_GetDataForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesCommonInterface_GetDataForFilter_Call) RunAndReturn(run func(model.TimeSeriesDescriptionDataType) ([]model.TimeSeriesDataType, error)) *TimeSeriesCommonInterface_GetDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetDescriptionsForFilter provides a mock function with given fields: filter
func (_m *TimeSeriesCommonInterface) GetDescriptionsForFilter(filter model.TimeSeriesDescriptionDataType) ([]model.TimeSeriesDescriptionDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetDescriptionsForFilter")
	}

	var r0 []model.TimeSeriesDescriptionDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.TimeSeriesDescriptionDataType) ([]model.TimeSeriesDescriptionDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.TimeSeriesDescriptionDataType) []model.TimeSeriesDescriptionDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.TimeSeriesDescriptionDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.TimeSeriesDescriptionDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TimeSeriesCommonInterface_GetDescriptionsForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDescriptionsForFilter'
type TimeSeriesCommonInterface_GetDescriptionsForFilter_Call struct {
	*mock.Call
}

// GetDescriptionsForFilter is a helper method to define mock.On call
//   - filter model.TimeSeriesDescriptionDataType
func (_e *TimeSeriesCommonInterface_Expecter) GetDescriptionsForFilter(filter interface{}) *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call {
	return &TimeSeriesCommonInterface_GetDescriptionsForFilter_Call{Call: _e.mock.On("GetDescriptionsForFilter", filter)}
}

func (_c *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call) Run(run func(filter model.TimeSeriesDescriptionDataType)) *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.TimeSeriesDescriptionDataType))
	})
	return _c
}

func (_c *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call) Return(_a0 []model.TimeSeriesDescriptionDataType, _a1 error) *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call) RunAndReturn(run func(model.TimeSeriesDescriptionDataType) ([]model.TimeSeriesDescriptionDataType, error)) *TimeSeriesCommonInterface_GetDescriptionsForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// NewTimeSeriesCommonInterface creates a new instance of TimeSeriesCommonInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTimeSeriesCommonInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TimeSeriesCommonInterface {
	mock := &TimeSeriesCommonInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
