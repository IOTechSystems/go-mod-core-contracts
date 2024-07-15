// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
)

// DeviceServiceCommandClient is an autogenerated mock type for the DeviceServiceCommandClient type
type DeviceServiceCommandClient struct {
	mock.Mock
}

// Discovery provides a mock function with given fields: ctx, baseUrl
func (_m *DeviceServiceCommandClient) Discovery(ctx context.Context, baseUrl string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, baseUrl)

	if len(ret) == 0 {
		panic("no return value specified for Discovery")
	}

	var r0 common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string) (common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, baseUrl)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, baseUrl)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, baseUrl)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// GetCommand provides a mock function with given fields: ctx, baseUrl, deviceName, commandName, queryParams
func (_m *DeviceServiceCommandClient) GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string) (*responses.EventResponse, errors.EdgeX) {
	ret := _m.Called(ctx, baseUrl, deviceName, commandName, queryParams)

	if len(ret) == 0 {
		panic("no return value specified for GetCommand")
	}

	var r0 *responses.EventResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) (*responses.EventResponse, errors.EdgeX)); ok {
		return rf(ctx, baseUrl, deviceName, commandName, queryParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) *responses.EventResponse); ok {
		r0 = rf(ctx, baseUrl, deviceName, commandName, queryParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*responses.EventResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) errors.EdgeX); ok {
		r1 = rf(ctx, baseUrl, deviceName, commandName, queryParams)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProfileScan provides a mock function with given fields: ctx, baseUrl, req
func (_m *DeviceServiceCommandClient) ProfileScan(ctx context.Context, baseUrl string, req requests.ProfileScanRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, baseUrl, req)

	if len(ret) == 0 {
		panic("no return value specified for ProfileScan")
	}

	var r0 common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, requests.ProfileScanRequest) (common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, baseUrl, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, requests.ProfileScanRequest) common.BaseResponse); ok {
		r0 = rf(ctx, baseUrl, req)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, requests.ProfileScanRequest) errors.EdgeX); ok {
		r1 = rf(ctx, baseUrl, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SetCommand provides a mock function with given fields: ctx, baseUrl, deviceName, commandName, queryParams, settings
func (_m *DeviceServiceCommandClient) SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, baseUrl, deviceName, commandName, queryParams, settings)

	if len(ret) == 0 {
		panic("no return value specified for SetCommand")
	}

	var r0 common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]string) (common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, baseUrl, deviceName, commandName, queryParams, settings)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]string) common.BaseResponse); ok {
		r0 = rf(ctx, baseUrl, deviceName, commandName, queryParams, settings)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, map[string]string) errors.EdgeX); ok {
		r1 = rf(ctx, baseUrl, deviceName, commandName, queryParams, settings)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// SetCommandWithObject provides a mock function with given fields: ctx, baseUrl, deviceName, commandName, queryParams, settings
func (_m *DeviceServiceCommandClient) SetCommandWithObject(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]interface{}) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, baseUrl, deviceName, commandName, queryParams, settings)

	if len(ret) == 0 {
		panic("no return value specified for SetCommandWithObject")
	}

	var r0 common.BaseResponse
	var r1 errors.EdgeX
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]interface{}) (common.BaseResponse, errors.EdgeX)); ok {
		return rf(ctx, baseUrl, deviceName, commandName, queryParams, settings)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, map[string]interface{}) common.BaseResponse); ok {
		r0 = rf(ctx, baseUrl, deviceName, commandName, queryParams, settings)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, map[string]interface{}) errors.EdgeX); ok {
		r1 = rf(ctx, baseUrl, deviceName, commandName, queryParams, settings)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// NewDeviceServiceCommandClient creates a new instance of DeviceServiceCommandClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeviceServiceCommandClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeviceServiceCommandClient {
	mock := &DeviceServiceCommandClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
