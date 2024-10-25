// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
)

// ProvisionWatcherClient is an autogenerated mock type for the ProvisionWatcherClient type
type ProvisionWatcherClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, reqs
func (_m *ProvisionWatcherClient) Add(ctx context.Context, reqs []requests.AddProvisionWatcherRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.AddProvisionWatcherRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.AddProvisionWatcherRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllProvisionWatchers provides a mock function with given fields: ctx, labels, offset, limit
func (_m *ProvisionWatcherClient) AllProvisionWatchers(ctx context.Context, labels []string, offset int, limit int) (responses.MultiProvisionWatchersResponse, errors.EdgeX) {
	ret := _m.Called(ctx, labels, offset, limit)

	var r0 responses.MultiProvisionWatchersResponse
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) responses.MultiProvisionWatchersResponse); ok {
		r0 = rf(ctx, labels, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiProvisionWatchersResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, labels, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteProvisionWatcherByName provides a mock function with given fields: ctx, name
func (_m *ProvisionWatcherClient) DeleteProvisionWatcherByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProvisionWatcherByName provides a mock function with given fields: ctx, name
func (_m *ProvisionWatcherClient) ProvisionWatcherByName(ctx context.Context, name string) (responses.ProvisionWatcherResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 responses.ProvisionWatcherResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.ProvisionWatcherResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(responses.ProvisionWatcherResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProvisionWatchersByProfileName provides a mock function with given fields: ctx, name, offset, limit
func (_m *ProvisionWatcherClient) ProvisionWatchersByProfileName(ctx context.Context, name string, offset int, limit int) (responses.MultiProvisionWatchersResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	var r0 responses.MultiProvisionWatchersResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiProvisionWatchersResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiProvisionWatchersResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ProvisionWatchersByServiceName provides a mock function with given fields: ctx, name, offset, limit
func (_m *ProvisionWatcherClient) ProvisionWatchersByServiceName(ctx context.Context, name string, offset int, limit int) (responses.MultiProvisionWatchersResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	var r0 responses.MultiProvisionWatchersResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiProvisionWatchersResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiProvisionWatchersResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, reqs
func (_m *ProvisionWatcherClient) Update(ctx context.Context, reqs []requests.UpdateProvisionWatcherRequest) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.UpdateProvisionWatcherRequest) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.UpdateProvisionWatcherRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewProvisionWatcherClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewProvisionWatcherClient creates a new instance of ProvisionWatcherClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProvisionWatcherClient(t mockConstructorTestingTNewProvisionWatcherClient) *ProvisionWatcherClient {
	mock := &ProvisionWatcherClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
