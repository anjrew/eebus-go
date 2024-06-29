// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	eebus_goapi "github.com/enbility/eebus-go/api"
	api "github.com/enbility/eebus-go/usecases/api"

	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"

	spine_goapi "github.com/enbility/spine-go/api"

	time "time"
)

// CsLPPInterface is an autogenerated mock type for the CsLPPInterface type
type CsLPPInterface struct {
	mock.Mock
}

type CsLPPInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *CsLPPInterface) EXPECT() *CsLPPInterface_Expecter {
	return &CsLPPInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *CsLPPInterface) AddFeatures() {
	_m.Called()
}

// CsLPPInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type CsLPPInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) AddFeatures() *CsLPPInterface_AddFeatures_Call {
	return &CsLPPInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *CsLPPInterface_AddFeatures_Call) Run(run func()) *CsLPPInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_AddFeatures_Call) Return() *CsLPPInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPPInterface_AddFeatures_Call) RunAndReturn(run func()) *CsLPPInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *CsLPPInterface) AddUseCase() {
	_m.Called()
}

// CsLPPInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type CsLPPInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) AddUseCase() *CsLPPInterface_AddUseCase_Call {
	return &CsLPPInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *CsLPPInterface_AddUseCase_Call) Run(run func()) *CsLPPInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_AddUseCase_Call) Return() *CsLPPInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPPInterface_AddUseCase_Call) RunAndReturn(run func()) *CsLPPInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// ApproveOrDenyProductionLimit provides a mock function with given fields: msgCounter, approve, reason
func (_m *CsLPPInterface) ApproveOrDenyProductionLimit(msgCounter model.MsgCounterType, approve bool, reason string) {
	_m.Called(msgCounter, approve, reason)
}

// CsLPPInterface_ApproveOrDenyProductionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApproveOrDenyProductionLimit'
type CsLPPInterface_ApproveOrDenyProductionLimit_Call struct {
	*mock.Call
}

// ApproveOrDenyProductionLimit is a helper method to define mock.On call
//   - msgCounter model.MsgCounterType
//   - approve bool
//   - reason string
func (_e *CsLPPInterface_Expecter) ApproveOrDenyProductionLimit(msgCounter interface{}, approve interface{}, reason interface{}) *CsLPPInterface_ApproveOrDenyProductionLimit_Call {
	return &CsLPPInterface_ApproveOrDenyProductionLimit_Call{Call: _e.mock.On("ApproveOrDenyProductionLimit", msgCounter, approve, reason)}
}

func (_c *CsLPPInterface_ApproveOrDenyProductionLimit_Call) Run(run func(msgCounter model.MsgCounterType, approve bool, reason string)) *CsLPPInterface_ApproveOrDenyProductionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.MsgCounterType), args[1].(bool), args[2].(string))
	})
	return _c
}

func (_c *CsLPPInterface_ApproveOrDenyProductionLimit_Call) Return() *CsLPPInterface_ApproveOrDenyProductionLimit_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPPInterface_ApproveOrDenyProductionLimit_Call) RunAndReturn(run func(model.MsgCounterType, bool, string)) *CsLPPInterface_ApproveOrDenyProductionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// AvailableScenariosForEntity provides a mock function with given fields: entity
func (_m *CsLPPInterface) AvailableScenariosForEntity(entity spine_goapi.EntityRemoteInterface) []uint {
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

// CsLPPInterface_AvailableScenariosForEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AvailableScenariosForEntity'
type CsLPPInterface_AvailableScenariosForEntity_Call struct {
	*mock.Call
}

// AvailableScenariosForEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CsLPPInterface_Expecter) AvailableScenariosForEntity(entity interface{}) *CsLPPInterface_AvailableScenariosForEntity_Call {
	return &CsLPPInterface_AvailableScenariosForEntity_Call{Call: _e.mock.On("AvailableScenariosForEntity", entity)}
}

func (_c *CsLPPInterface_AvailableScenariosForEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CsLPPInterface_AvailableScenariosForEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CsLPPInterface_AvailableScenariosForEntity_Call) Return(_a0 []uint) *CsLPPInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPPInterface_AvailableScenariosForEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) []uint) *CsLPPInterface_AvailableScenariosForEntity_Call {
	_c.Call.Return(run)
	return _c
}

// ContractualProductionNominalMax provides a mock function with given fields:
func (_m *CsLPPInterface) ContractualProductionNominalMax() (float64, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ContractualProductionNominalMax")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func() (float64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CsLPPInterface_ContractualProductionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContractualProductionNominalMax'
type CsLPPInterface_ContractualProductionNominalMax_Call struct {
	*mock.Call
}

// ContractualProductionNominalMax is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) ContractualProductionNominalMax() *CsLPPInterface_ContractualProductionNominalMax_Call {
	return &CsLPPInterface_ContractualProductionNominalMax_Call{Call: _e.mock.On("ContractualProductionNominalMax")}
}

func (_c *CsLPPInterface_ContractualProductionNominalMax_Call) Run(run func()) *CsLPPInterface_ContractualProductionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_ContractualProductionNominalMax_Call) Return(_a0 float64, _a1 error) *CsLPPInterface_ContractualProductionNominalMax_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CsLPPInterface_ContractualProductionNominalMax_Call) RunAndReturn(run func() (float64, error)) *CsLPPInterface_ContractualProductionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeDurationMinimum provides a mock function with given fields:
func (_m *CsLPPInterface) FailsafeDurationMinimum() (time.Duration, bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FailsafeDurationMinimum")
	}

	var r0 time.Duration
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func() (time.Duration, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() time.Duration); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CsLPPInterface_FailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeDurationMinimum'
type CsLPPInterface_FailsafeDurationMinimum_Call struct {
	*mock.Call
}

// FailsafeDurationMinimum is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) FailsafeDurationMinimum() *CsLPPInterface_FailsafeDurationMinimum_Call {
	return &CsLPPInterface_FailsafeDurationMinimum_Call{Call: _e.mock.On("FailsafeDurationMinimum")}
}

func (_c *CsLPPInterface_FailsafeDurationMinimum_Call) Run(run func()) *CsLPPInterface_FailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_FailsafeDurationMinimum_Call) Return(duration time.Duration, isChangeable bool, resultErr error) *CsLPPInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(duration, isChangeable, resultErr)
	return _c
}

func (_c *CsLPPInterface_FailsafeDurationMinimum_Call) RunAndReturn(run func() (time.Duration, bool, error)) *CsLPPInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeProductionActivePowerLimit provides a mock function with given fields:
func (_m *CsLPPInterface) FailsafeProductionActivePowerLimit() (float64, bool, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FailsafeProductionActivePowerLimit")
	}

	var r0 float64
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func() (float64, bool, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CsLPPInterface_FailsafeProductionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeProductionActivePowerLimit'
type CsLPPInterface_FailsafeProductionActivePowerLimit_Call struct {
	*mock.Call
}

// FailsafeProductionActivePowerLimit is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) FailsafeProductionActivePowerLimit() *CsLPPInterface_FailsafeProductionActivePowerLimit_Call {
	return &CsLPPInterface_FailsafeProductionActivePowerLimit_Call{Call: _e.mock.On("FailsafeProductionActivePowerLimit")}
}

func (_c *CsLPPInterface_FailsafeProductionActivePowerLimit_Call) Run(run func()) *CsLPPInterface_FailsafeProductionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_FailsafeProductionActivePowerLimit_Call) Return(value float64, isChangeable bool, resultErr error) *CsLPPInterface_FailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(value, isChangeable, resultErr)
	return _c
}

func (_c *CsLPPInterface_FailsafeProductionActivePowerLimit_Call) RunAndReturn(run func() (float64, bool, error)) *CsLPPInterface_FailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntityType provides a mock function with given fields: entity
func (_m *CsLPPInterface) IsCompatibleEntityType(entity spine_goapi.EntityRemoteInterface) bool {
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

// CsLPPInterface_IsCompatibleEntityType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntityType'
type CsLPPInterface_IsCompatibleEntityType_Call struct {
	*mock.Call
}

// IsCompatibleEntityType is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *CsLPPInterface_Expecter) IsCompatibleEntityType(entity interface{}) *CsLPPInterface_IsCompatibleEntityType_Call {
	return &CsLPPInterface_IsCompatibleEntityType_Call{Call: _e.mock.On("IsCompatibleEntityType", entity)}
}

func (_c *CsLPPInterface_IsCompatibleEntityType_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *CsLPPInterface_IsCompatibleEntityType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *CsLPPInterface_IsCompatibleEntityType_Call) Return(_a0 bool) *CsLPPInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPPInterface_IsCompatibleEntityType_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *CsLPPInterface_IsCompatibleEntityType_Call {
	_c.Call.Return(run)
	return _c
}

// IsHeartbeatWithinDuration provides a mock function with given fields:
func (_m *CsLPPInterface) IsHeartbeatWithinDuration() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsHeartbeatWithinDuration")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// CsLPPInterface_IsHeartbeatWithinDuration_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsHeartbeatWithinDuration'
type CsLPPInterface_IsHeartbeatWithinDuration_Call struct {
	*mock.Call
}

// IsHeartbeatWithinDuration is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) IsHeartbeatWithinDuration() *CsLPPInterface_IsHeartbeatWithinDuration_Call {
	return &CsLPPInterface_IsHeartbeatWithinDuration_Call{Call: _e.mock.On("IsHeartbeatWithinDuration")}
}

func (_c *CsLPPInterface_IsHeartbeatWithinDuration_Call) Run(run func()) *CsLPPInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_IsHeartbeatWithinDuration_Call) Return(_a0 bool) *CsLPPInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPPInterface_IsHeartbeatWithinDuration_Call) RunAndReturn(run func() bool) *CsLPPInterface_IsHeartbeatWithinDuration_Call {
	_c.Call.Return(run)
	return _c
}

// IsScenarioAvailableAtEntity provides a mock function with given fields: entity, scenario
func (_m *CsLPPInterface) IsScenarioAvailableAtEntity(entity spine_goapi.EntityRemoteInterface, scenario uint) bool {
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

// CsLPPInterface_IsScenarioAvailableAtEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsScenarioAvailableAtEntity'
type CsLPPInterface_IsScenarioAvailableAtEntity_Call struct {
	*mock.Call
}

// IsScenarioAvailableAtEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - scenario uint
func (_e *CsLPPInterface_Expecter) IsScenarioAvailableAtEntity(entity interface{}, scenario interface{}) *CsLPPInterface_IsScenarioAvailableAtEntity_Call {
	return &CsLPPInterface_IsScenarioAvailableAtEntity_Call{Call: _e.mock.On("IsScenarioAvailableAtEntity", entity, scenario)}
}

func (_c *CsLPPInterface_IsScenarioAvailableAtEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, scenario uint)) *CsLPPInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(uint))
	})
	return _c
}

func (_c *CsLPPInterface_IsScenarioAvailableAtEntity_Call) Return(_a0 bool) *CsLPPInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPPInterface_IsScenarioAvailableAtEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, uint) bool) *CsLPPInterface_IsScenarioAvailableAtEntity_Call {
	_c.Call.Return(run)
	return _c
}

// PendingProductionLimits provides a mock function with given fields:
func (_m *CsLPPInterface) PendingProductionLimits() map[model.MsgCounterType]api.LoadLimit {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PendingProductionLimits")
	}

	var r0 map[model.MsgCounterType]api.LoadLimit
	if rf, ok := ret.Get(0).(func() map[model.MsgCounterType]api.LoadLimit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.MsgCounterType]api.LoadLimit)
		}
	}

	return r0
}

// CsLPPInterface_PendingProductionLimits_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingProductionLimits'
type CsLPPInterface_PendingProductionLimits_Call struct {
	*mock.Call
}

// PendingProductionLimits is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) PendingProductionLimits() *CsLPPInterface_PendingProductionLimits_Call {
	return &CsLPPInterface_PendingProductionLimits_Call{Call: _e.mock.On("PendingProductionLimits")}
}

func (_c *CsLPPInterface_PendingProductionLimits_Call) Run(run func()) *CsLPPInterface_PendingProductionLimits_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_PendingProductionLimits_Call) Return(_a0 map[model.MsgCounterType]api.LoadLimit) *CsLPPInterface_PendingProductionLimits_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPPInterface_PendingProductionLimits_Call) RunAndReturn(run func() map[model.MsgCounterType]api.LoadLimit) *CsLPPInterface_PendingProductionLimits_Call {
	_c.Call.Return(run)
	return _c
}

// ProductionLimit provides a mock function with given fields:
func (_m *CsLPPInterface) ProductionLimit() (api.LoadLimit, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProductionLimit")
	}

	var r0 api.LoadLimit
	var r1 error
	if rf, ok := ret.Get(0).(func() (api.LoadLimit, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() api.LoadLimit); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.LoadLimit)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CsLPPInterface_ProductionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProductionLimit'
type CsLPPInterface_ProductionLimit_Call struct {
	*mock.Call
}

// ProductionLimit is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) ProductionLimit() *CsLPPInterface_ProductionLimit_Call {
	return &CsLPPInterface_ProductionLimit_Call{Call: _e.mock.On("ProductionLimit")}
}

func (_c *CsLPPInterface_ProductionLimit_Call) Run(run func()) *CsLPPInterface_ProductionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_ProductionLimit_Call) Return(_a0 api.LoadLimit, _a1 error) *CsLPPInterface_ProductionLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CsLPPInterface_ProductionLimit_Call) RunAndReturn(run func() (api.LoadLimit, error)) *CsLPPInterface_ProductionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// RemoteEntitiesScenarios provides a mock function with given fields:
func (_m *CsLPPInterface) RemoteEntitiesScenarios() []eebus_goapi.RemoteEntityScenarios {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for RemoteEntitiesScenarios")
	}

	var r0 []eebus_goapi.RemoteEntityScenarios
	if rf, ok := ret.Get(0).(func() []eebus_goapi.RemoteEntityScenarios); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]eebus_goapi.RemoteEntityScenarios)
		}
	}

	return r0
}

// CsLPPInterface_RemoteEntitiesScenarios_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoteEntitiesScenarios'
type CsLPPInterface_RemoteEntitiesScenarios_Call struct {
	*mock.Call
}

// RemoteEntitiesScenarios is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) RemoteEntitiesScenarios() *CsLPPInterface_RemoteEntitiesScenarios_Call {
	return &CsLPPInterface_RemoteEntitiesScenarios_Call{Call: _e.mock.On("RemoteEntitiesScenarios")}
}

func (_c *CsLPPInterface_RemoteEntitiesScenarios_Call) Run(run func()) *CsLPPInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_RemoteEntitiesScenarios_Call) Return(_a0 []eebus_goapi.RemoteEntityScenarios) *CsLPPInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CsLPPInterface_RemoteEntitiesScenarios_Call) RunAndReturn(run func() []eebus_goapi.RemoteEntityScenarios) *CsLPPInterface_RemoteEntitiesScenarios_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveUseCase provides a mock function with given fields:
func (_m *CsLPPInterface) RemoveUseCase() {
	_m.Called()
}

// CsLPPInterface_RemoveUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveUseCase'
type CsLPPInterface_RemoveUseCase_Call struct {
	*mock.Call
}

// RemoveUseCase is a helper method to define mock.On call
func (_e *CsLPPInterface_Expecter) RemoveUseCase() *CsLPPInterface_RemoveUseCase_Call {
	return &CsLPPInterface_RemoveUseCase_Call{Call: _e.mock.On("RemoveUseCase")}
}

func (_c *CsLPPInterface_RemoveUseCase_Call) Run(run func()) *CsLPPInterface_RemoveUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CsLPPInterface_RemoveUseCase_Call) Return() *CsLPPInterface_RemoveUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPPInterface_RemoveUseCase_Call) RunAndReturn(run func()) *CsLPPInterface_RemoveUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// SetContractualProductionNominalMax provides a mock function with given fields: value
func (_m *CsLPPInterface) SetContractualProductionNominalMax(value float64) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for SetContractualProductionNominalMax")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(float64) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPPInterface_SetContractualProductionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetContractualProductionNominalMax'
type CsLPPInterface_SetContractualProductionNominalMax_Call struct {
	*mock.Call
}

// SetContractualProductionNominalMax is a helper method to define mock.On call
//   - value float64
func (_e *CsLPPInterface_Expecter) SetContractualProductionNominalMax(value interface{}) *CsLPPInterface_SetContractualProductionNominalMax_Call {
	return &CsLPPInterface_SetContractualProductionNominalMax_Call{Call: _e.mock.On("SetContractualProductionNominalMax", value)}
}

func (_c *CsLPPInterface_SetContractualProductionNominalMax_Call) Run(run func(value float64)) *CsLPPInterface_SetContractualProductionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *CsLPPInterface_SetContractualProductionNominalMax_Call) Return(resultErr error) *CsLPPInterface_SetContractualProductionNominalMax_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPPInterface_SetContractualProductionNominalMax_Call) RunAndReturn(run func(float64) error) *CsLPPInterface_SetContractualProductionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// SetFailsafeDurationMinimum provides a mock function with given fields: duration, changeable
func (_m *CsLPPInterface) SetFailsafeDurationMinimum(duration time.Duration, changeable bool) error {
	ret := _m.Called(duration, changeable)

	if len(ret) == 0 {
		panic("no return value specified for SetFailsafeDurationMinimum")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Duration, bool) error); ok {
		r0 = rf(duration, changeable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPPInterface_SetFailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFailsafeDurationMinimum'
type CsLPPInterface_SetFailsafeDurationMinimum_Call struct {
	*mock.Call
}

// SetFailsafeDurationMinimum is a helper method to define mock.On call
//   - duration time.Duration
//   - changeable bool
func (_e *CsLPPInterface_Expecter) SetFailsafeDurationMinimum(duration interface{}, changeable interface{}) *CsLPPInterface_SetFailsafeDurationMinimum_Call {
	return &CsLPPInterface_SetFailsafeDurationMinimum_Call{Call: _e.mock.On("SetFailsafeDurationMinimum", duration, changeable)}
}

func (_c *CsLPPInterface_SetFailsafeDurationMinimum_Call) Run(run func(duration time.Duration, changeable bool)) *CsLPPInterface_SetFailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(time.Duration), args[1].(bool))
	})
	return _c
}

func (_c *CsLPPInterface_SetFailsafeDurationMinimum_Call) Return(resultErr error) *CsLPPInterface_SetFailsafeDurationMinimum_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPPInterface_SetFailsafeDurationMinimum_Call) RunAndReturn(run func(time.Duration, bool) error) *CsLPPInterface_SetFailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// SetFailsafeProductionActivePowerLimit provides a mock function with given fields: value, changeable
func (_m *CsLPPInterface) SetFailsafeProductionActivePowerLimit(value float64, changeable bool) error {
	ret := _m.Called(value, changeable)

	if len(ret) == 0 {
		panic("no return value specified for SetFailsafeProductionActivePowerLimit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(float64, bool) error); ok {
		r0 = rf(value, changeable)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetFailsafeProductionActivePowerLimit'
type CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call struct {
	*mock.Call
}

// SetFailsafeProductionActivePowerLimit is a helper method to define mock.On call
//   - value float64
//   - changeable bool
func (_e *CsLPPInterface_Expecter) SetFailsafeProductionActivePowerLimit(value interface{}, changeable interface{}) *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call {
	return &CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call{Call: _e.mock.On("SetFailsafeProductionActivePowerLimit", value, changeable)}
}

func (_c *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call) Run(run func(value float64, changeable bool)) *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(bool))
	})
	return _c
}

func (_c *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call) Return(resultErr error) *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call) RunAndReturn(run func(float64, bool) error) *CsLPPInterface_SetFailsafeProductionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// SetProductionLimit provides a mock function with given fields: limit
func (_m *CsLPPInterface) SetProductionLimit(limit api.LoadLimit) error {
	ret := _m.Called(limit)

	if len(ret) == 0 {
		panic("no return value specified for SetProductionLimit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(api.LoadLimit) error); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CsLPPInterface_SetProductionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetProductionLimit'
type CsLPPInterface_SetProductionLimit_Call struct {
	*mock.Call
}

// SetProductionLimit is a helper method to define mock.On call
//   - limit api.LoadLimit
func (_e *CsLPPInterface_Expecter) SetProductionLimit(limit interface{}) *CsLPPInterface_SetProductionLimit_Call {
	return &CsLPPInterface_SetProductionLimit_Call{Call: _e.mock.On("SetProductionLimit", limit)}
}

func (_c *CsLPPInterface_SetProductionLimit_Call) Run(run func(limit api.LoadLimit)) *CsLPPInterface_SetProductionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.LoadLimit))
	})
	return _c
}

func (_c *CsLPPInterface_SetProductionLimit_Call) Return(resultErr error) *CsLPPInterface_SetProductionLimit_Call {
	_c.Call.Return(resultErr)
	return _c
}

func (_c *CsLPPInterface_SetProductionLimit_Call) RunAndReturn(run func(api.LoadLimit) error) *CsLPPInterface_SetProductionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *CsLPPInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// CsLPPInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type CsLPPInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *CsLPPInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *CsLPPInterface_UpdateUseCaseAvailability_Call {
	return &CsLPPInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *CsLPPInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *CsLPPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *CsLPPInterface_UpdateUseCaseAvailability_Call) Return() *CsLPPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *CsLPPInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *CsLPPInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// NewCsLPPInterface creates a new instance of CsLPPInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCsLPPInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CsLPPInterface {
	mock := &CsLPPInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
