// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// MeasurementServerInterface is an autogenerated mock type for the MeasurementServerInterface type
type MeasurementServerInterface struct {
	mock.Mock
}

type MeasurementServerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MeasurementServerInterface) EXPECT() *MeasurementServerInterface_Expecter {
	return &MeasurementServerInterface_Expecter{mock: &_m.Mock}
}

// AddDescription provides a mock function with given fields: description
func (_m *MeasurementServerInterface) AddDescription(description model.MeasurementDescriptionDataType) *model.MeasurementIdType {
	ret := _m.Called(description)

	if len(ret) == 0 {
		panic("no return value specified for AddDescription")
	}

	var r0 *model.MeasurementIdType
	if rf, ok := ret.Get(0).(func(model.MeasurementDescriptionDataType) *model.MeasurementIdType); ok {
		r0 = rf(description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MeasurementIdType)
		}
	}

	return r0
}

// MeasurementServerInterface_AddDescription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddDescription'
type MeasurementServerInterface_AddDescription_Call struct {
	*mock.Call
}

// AddDescription is a helper method to define mock.On call
//   - description model.MeasurementDescriptionDataType
func (_e *MeasurementServerInterface_Expecter) AddDescription(description interface{}) *MeasurementServerInterface_AddDescription_Call {
	return &MeasurementServerInterface_AddDescription_Call{Call: _e.mock.On("AddDescription", description)}
}

func (_c *MeasurementServerInterface_AddDescription_Call) Run(run func(description model.MeasurementDescriptionDataType)) *MeasurementServerInterface_AddDescription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementDescriptionDataType))
	})
	return _c
}

func (_c *MeasurementServerInterface_AddDescription_Call) Return(_a0 *model.MeasurementIdType) *MeasurementServerInterface_AddDescription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MeasurementServerInterface_AddDescription_Call) RunAndReturn(run func(model.MeasurementDescriptionDataType) *model.MeasurementIdType) *MeasurementServerInterface_AddDescription_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDataForFilter provides a mock function with given fields: data, deleteElements, filter
func (_m *MeasurementServerInterface) UpdateDataForFilter(data model.MeasurementDataType, deleteElements *model.MeasurementDataElementsType, filter model.MeasurementDescriptionDataType) error {
	ret := _m.Called(data, deleteElements, filter)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDataForFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.MeasurementDataType, *model.MeasurementDataElementsType, model.MeasurementDescriptionDataType) error); ok {
		r0 = rf(data, deleteElements, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MeasurementServerInterface_UpdateDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDataForFilter'
type MeasurementServerInterface_UpdateDataForFilter_Call struct {
	*mock.Call
}

// UpdateDataForFilter is a helper method to define mock.On call
//   - data model.MeasurementDataType
//   - deleteElements *model.MeasurementDataElementsType
//   - filter model.MeasurementDescriptionDataType
func (_e *MeasurementServerInterface_Expecter) UpdateDataForFilter(data interface{}, deleteElements interface{}, filter interface{}) *MeasurementServerInterface_UpdateDataForFilter_Call {
	return &MeasurementServerInterface_UpdateDataForFilter_Call{Call: _e.mock.On("UpdateDataForFilter", data, deleteElements, filter)}
}

func (_c *MeasurementServerInterface_UpdateDataForFilter_Call) Run(run func(data model.MeasurementDataType, deleteElements *model.MeasurementDataElementsType, filter model.MeasurementDescriptionDataType)) *MeasurementServerInterface_UpdateDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementDataType), args[1].(*model.MeasurementDataElementsType), args[2].(model.MeasurementDescriptionDataType))
	})
	return _c
}

func (_c *MeasurementServerInterface_UpdateDataForFilter_Call) Return(_a0 error) *MeasurementServerInterface_UpdateDataForFilter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MeasurementServerInterface_UpdateDataForFilter_Call) RunAndReturn(run func(model.MeasurementDataType, *model.MeasurementDataElementsType, model.MeasurementDescriptionDataType) error) *MeasurementServerInterface_UpdateDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateDataForId provides a mock function with given fields: data, deleteElements, measurementId
func (_m *MeasurementServerInterface) UpdateDataForId(data model.MeasurementDataType, deleteElements *model.MeasurementDataElementsType, measurementId model.MeasurementIdType) error {
	ret := _m.Called(data, deleteElements, measurementId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDataForId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.MeasurementDataType, *model.MeasurementDataElementsType, model.MeasurementIdType) error); ok {
		r0 = rf(data, deleteElements, measurementId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MeasurementServerInterface_UpdateDataForId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateDataForId'
type MeasurementServerInterface_UpdateDataForId_Call struct {
	*mock.Call
}

// UpdateDataForId is a helper method to define mock.On call
//   - data model.MeasurementDataType
//   - deleteElements *model.MeasurementDataElementsType
//   - measurementId model.MeasurementIdType
func (_e *MeasurementServerInterface_Expecter) UpdateDataForId(data interface{}, deleteElements interface{}, measurementId interface{}) *MeasurementServerInterface_UpdateDataForId_Call {
	return &MeasurementServerInterface_UpdateDataForId_Call{Call: _e.mock.On("UpdateDataForId", data, deleteElements, measurementId)}
}

func (_c *MeasurementServerInterface_UpdateDataForId_Call) Run(run func(data model.MeasurementDataType, deleteElements *model.MeasurementDataElementsType, measurementId model.MeasurementIdType)) *MeasurementServerInterface_UpdateDataForId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MeasurementDataType), args[1].(*model.MeasurementDataElementsType), args[2].(model.MeasurementIdType))
	})
	return _c
}

func (_c *MeasurementServerInterface_UpdateDataForId_Call) Return(_a0 error) *MeasurementServerInterface_UpdateDataForId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MeasurementServerInterface_UpdateDataForId_Call) RunAndReturn(run func(model.MeasurementDataType, *model.MeasurementDataElementsType, model.MeasurementIdType) error) *MeasurementServerInterface_UpdateDataForId_Call {
	_c.Call.Return(run)
	return _c
}

// NewMeasurementServerInterface creates a new instance of MeasurementServerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMeasurementServerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MeasurementServerInterface {
	mock := &MeasurementServerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
