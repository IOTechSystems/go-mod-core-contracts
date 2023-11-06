// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
)

// EventClient is an autogenerated mock type for the EventClient type
type EventClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, req
func (_m *EventClient) Add(ctx context.Context, req requests.AddEventRequest) (common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, req)

	var r0 common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, requests.AddEventRequest) common.BaseWithIdResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(common.BaseWithIdResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, requests.AddEventRequest) errors.EdgeX); ok {
		r1 = rf(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllEvents provides a mock function with given fields: ctx, offset, limit
func (_m *EventClient) AllEvents(ctx context.Context, offset int, limit int) (responses.MultiEventsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, offset, limit)

	var r0 responses.MultiEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int) responses.MultiEventsResponse); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiEventsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteByAge provides a mock function with given fields: ctx, age
func (_m *EventClient) DeleteByAge(ctx context.Context, age int) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, age)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, int) common.BaseResponse); ok {
		r0 = rf(ctx, age)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int) errors.EdgeX); ok {
		r1 = rf(ctx, age)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteByDeviceName provides a mock function with given fields: ctx, name
func (_m *EventClient) DeleteByDeviceName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
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

// EventCount provides a mock function with given fields: ctx
func (_m *EventClient) EventCount(ctx context.Context) (common.CountResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.CountResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.CountResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.CountResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// EventCountByDeviceName provides a mock function with given fields: ctx, name
func (_m *EventClient) EventCountByDeviceName(ctx context.Context, name string) (common.CountResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.CountResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.CountResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.CountResponse)
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

// EventsByDeviceName provides a mock function with given fields: ctx, name, offset, limit
func (_m *EventClient) EventsByDeviceName(ctx context.Context, name string, offset int, limit int) (responses.MultiEventsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	var r0 responses.MultiEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiEventsResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiEventsResponse)
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

// EventsByTimeRange provides a mock function with given fields: ctx, start, end, offset, limit
func (_m *EventClient) EventsByTimeRange(ctx context.Context, start int64, end int64, offset int, limit int) (responses.MultiEventsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, start, end, offset, limit)

	var r0 responses.MultiEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, int, int) responses.MultiEventsResponse); ok {
		r0 = rf(ctx, start, end, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiEventsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewEventClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventClient creates a new instance of EventClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventClient(t mockConstructorTestingTNewEventClient) *EventClient {
	mock := &EventClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
