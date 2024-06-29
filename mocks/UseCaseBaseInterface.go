// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/eebus-go/api"
	mock "github.com/stretchr/testify/mock"

	spine_goapi "github.com/enbility/spine-go/api"
)

// UseCaseBaseInterface is an autogenerated mock type for the UseCaseBaseInterface type
type UseCaseBaseInterface struct {
	mock.Mock
}

type UseCaseBaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *UseCaseBaseInterface) EXPECT() *UseCaseBaseInterface_Expecter {
	return &UseCaseBaseInterface_Expecter{mock: &_m.Mock}
}

// AddUseCase provides a mock function with given fields:
func (_m *UseCaseBaseInterface) AddUseCase() {
	_m.Called()
}

// UseCaseBaseInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type UseCaseBaseInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *UseCaseBaseInterface_Expecter) AddUseCase() *UseCaseBaseInterface_AddUseCase_Call {
	return &UseCaseBaseInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *UseCaseBaseInterface_AddUseCase_Call) Run(run func()) *UseCaseBaseInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UseCaseBaseInterface_AddUseCase_Call) Return() *UseCaseBaseInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *UseCaseBaseInterface_AddUseCase_Call) RunAndReturn(run func()) *UseCaseBaseInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *UseCaseBaseInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for AvailableScenariosForEntity")
	}

	var r0 []uint
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) []uint); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	return r0
}

// UseCaseBaseInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type UseCaseBaseInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *UseCaseBaseInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *UseCaseBaseInterface_AvailableScenariosForEntity_Call {
	return &UseCaseBaseInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *UseCaseBaseInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *UseCaseBaseInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *UseCaseBaseInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *UseCaseBaseInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseBaseInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *UseCaseBaseInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *UseCaseBaseInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsCompatibleEntityType")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UseCaseBaseInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type UseCaseBaseInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *UseCaseBaseInterface_Expecter) IsCompatibleEntityType(entity interface{}) *UseCaseBaseInterface_IsCompatibleEntityType_Call {
	return &UseCaseBaseInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *UseCaseBaseInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *UseCaseBaseInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *UseCaseBaseInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *UseCaseBaseInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseBaseInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *UseCaseBaseInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *UseCaseBaseInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
	ret := _m.Called(entity, scenario)

	if len(ret) == 0 {
		panic("no return value specified for IsScenarioAvailableAtEntity")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, uint) bool); ok {
		r0 = rf(entity, scenario)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *UseCaseBaseInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call {
	return &UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *UseCaseBaseInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *UseCaseBaseInterface) RemoteEntitiesScenarios() []api.RemoteEntityScenarios {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteEntitiesScenarios")
	}

	var r0 []api.RemoteEntityScenarios
	if rf, ok := ret.Get(0).(func() []api.RemoteEntityScenarios); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]api.RemoteEntityScenarios)
		}
	}

	return r0
}

// UseCaseBaseInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type UseCaseBaseInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *UseCaseBaseInterface_Expecter) RemoteEntitiesScenarios() *UseCaseBaseInterface_RemoteEntitiesScenarios_Call {
	return &UseCaseBaseInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *UseCaseBaseInterface_RemoteEntitiesScenarios_Call) Run(run func()) *UseCaseBaseInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UseCaseBaseInterface_RemoteEntitiesScenarios_Call) Return(_a0 []api.RemoteEntityScenarios) *UseCaseBaseInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UseCaseBaseInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []api.RemoteEntityScenarios) *UseCaseBaseInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *UseCaseBaseInterface) RemoveUseCase() {
	_m.Called()
}

// UseCaseBaseInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type UseCaseBaseInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *UseCaseBaseInterface_Expecter) RemoveUseCase() *UseCaseBaseInterface_RemoveUseCase_Call {
	return &UseCaseBaseInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *UseCaseBaseInterface_RemoveUseCase_Call) Run(run func()) *UseCaseBaseInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UseCaseBaseInterface_RemoveUseCase_Call) Return() *UseCaseBaseInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *UseCaseBaseInterface_RemoveUseCase_Call) RunAndReturn(run func()) *UseCaseBaseInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *UseCaseBaseInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// UseCaseBaseInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type UseCaseBaseInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *UseCaseBaseInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *UseCaseBaseInterface_UpdateUseCaseAvailability_Call {
	return &UseCaseBaseInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *UseCaseBaseInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *UseCaseBaseInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *UseCaseBaseInterface_UpdateUseCaseAvailability_Call) Return() *UseCaseBaseInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *UseCaseBaseInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *UseCaseBaseInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewUseCaseBaseInterface creates a new instance of UseCaseBaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUseCaseBaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UseCaseBaseInterface {
	mock := &UseCaseBaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
