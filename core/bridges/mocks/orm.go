// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	auth "github.com/smartcontractkit/chainlink/core/auth"
	bridges "github.com/smartcontractkit/chainlink/core/bridges"

	mock "github.com/stretchr/testify/mock"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// BridgeTypes provides a mock function with given fields: offset, limit
func (_m *ORM) BridgeTypes(offset int, limit int) ([]bridges.BridgeType, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []bridges.BridgeType
	if rf, ok := ret.Get(0).(func(int, int) []bridges.BridgeType); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bridges.BridgeType)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CreateBridgeType provides a mock function with given fields: bt
func (_m *ORM) CreateBridgeType(bt *bridges.BridgeType) error {
	ret := _m.Called(bt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bridges.BridgeType) error); ok {
		r0 = rf(bt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateExternalInitiator provides a mock function with given fields: externalInitiator
func (_m *ORM) CreateExternalInitiator(externalInitiator *bridges.ExternalInitiator) error {
	ret := _m.Called(externalInitiator)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bridges.ExternalInitiator) error); ok {
		r0 = rf(externalInitiator)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBridgeType provides a mock function with given fields: bt
func (_m *ORM) DeleteBridgeType(bt *bridges.BridgeType) error {
	ret := _m.Called(bt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bridges.BridgeType) error); ok {
		r0 = rf(bt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteExternalInitiator provides a mock function with given fields: name
func (_m *ORM) DeleteExternalInitiator(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExternalInitiators provides a mock function with given fields: offset, limit
func (_m *ORM) ExternalInitiators(offset int, limit int) ([]bridges.ExternalInitiator, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []bridges.ExternalInitiator
	if rf, ok := ret.Get(0).(func(int, int) []bridges.ExternalInitiator); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bridges.ExternalInitiator)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindBridge provides a mock function with given fields: name
func (_m *ORM) FindBridge(name bridges.BridgeName) (bridges.BridgeType, error) {
	ret := _m.Called(name)

	var r0 bridges.BridgeType
	if rf, ok := ret.Get(0).(func(bridges.BridgeName) bridges.BridgeType); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bridges.BridgeType)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bridges.BridgeName) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBridges provides a mock function with given fields: name
func (_m *ORM) FindBridges(name []bridges.BridgeName) ([]bridges.BridgeType, error) {
	ret := _m.Called(name)

	var r0 []bridges.BridgeType
	if rf, ok := ret.Get(0).(func([]bridges.BridgeName) []bridges.BridgeType); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bridges.BridgeType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]bridges.BridgeName) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindExternalInitiator provides a mock function with given fields: eia
func (_m *ORM) FindExternalInitiator(eia *auth.Token) (*bridges.ExternalInitiator, error) {
	ret := _m.Called(eia)

	var r0 *bridges.ExternalInitiator
	if rf, ok := ret.Get(0).(func(*auth.Token) *bridges.ExternalInitiator); ok {
		r0 = rf(eia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridges.ExternalInitiator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*auth.Token) error); ok {
		r1 = rf(eia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindExternalInitiatorByName provides a mock function with given fields: iname
func (_m *ORM) FindExternalInitiatorByName(iname string) (bridges.ExternalInitiator, error) {
	ret := _m.Called(iname)

	var r0 bridges.ExternalInitiator
	if rf, ok := ret.Get(0).(func(string) bridges.ExternalInitiator); ok {
		r0 = rf(iname)
	} else {
		r0 = ret.Get(0).(bridges.ExternalInitiator)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(iname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBridgeType provides a mock function with given fields: bt, btr
func (_m *ORM) UpdateBridgeType(bt *bridges.BridgeType, btr *bridges.BridgeTypeRequest) error {
	ret := _m.Called(bt, btr)

	var r0 error
	if rf, ok := ret.Get(0).(func(*bridges.BridgeType, *bridges.BridgeTypeRequest) error); ok {
		r0 = rf(bt, btr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
