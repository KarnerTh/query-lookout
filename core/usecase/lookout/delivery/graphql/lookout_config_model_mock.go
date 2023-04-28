// Code generated by mockery v2.21.4. DO NOT EDIT.

package graphql

import (
	deliverygraphql "github.com/KarnerTh/query-lookout/usecase/review/delivery/graphql"
	mock "github.com/stretchr/testify/mock"
)

// mockLookoutConfigModel is an autogenerated mock type for the lookoutConfigModel type
type mockLookoutConfigModel struct {
	mock.Mock
}

// Cron provides a mock function with given fields:
func (_m *mockLookoutConfigModel) Cron() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Id provides a mock function with given fields:
func (_m *mockLookoutConfigModel) Id() int32 {
	ret := _m.Called()

	var r0 int32
	if rf, ok := ret.Get(0).(func() int32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int32)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *mockLookoutConfigModel) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NotifyLocal provides a mock function with given fields:
func (_m *mockLookoutConfigModel) NotifyLocal() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NotifyMail provides a mock function with given fields:
func (_m *mockLookoutConfigModel) NotifyMail() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Query provides a mock function with given fields:
func (_m *mockLookoutConfigModel) Query() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Rules provides a mock function with given fields:
func (_m *mockLookoutConfigModel) Rules() ([]deliverygraphql.ReviewRuleModel, error) {
	ret := _m.Called()

	var r0 []deliverygraphql.ReviewRuleModel
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]deliverygraphql.ReviewRuleModel, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []deliverygraphql.ReviewRuleModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]deliverygraphql.ReviewRuleModel)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTnewMockLookoutConfigModel interface {
	mock.TestingT
	Cleanup(func())
}

// newMockLookoutConfigModel creates a new instance of mockLookoutConfigModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockLookoutConfigModel(t mockConstructorTestingTnewMockLookoutConfigModel) *mockLookoutConfigModel {
	mock := &mockLookoutConfigModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
