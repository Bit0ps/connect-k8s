// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/x/marketmap/types"
)

// MarketMapKeeper is an autogenerated mock type for the MarketMapKeeper type
type MarketMapKeeper struct {
	mock.Mock
}

type MarketMapKeeper_Expecter struct {
	mock *mock.Mock
}

func (_m *MarketMapKeeper) EXPECT() *MarketMapKeeper_Expecter {
	return &MarketMapKeeper_Expecter{mock: &_m.Mock}
}

// GetMarket provides a mock function with given fields: ctx, tickerStr
func (_m *MarketMapKeeper) GetMarket(ctx context.Context, tickerStr string) (types.Market, error) {
	ret := _m.Called(ctx, tickerStr)

	if len(ret) == 0 {
		panic("no return value specified for GetMarket")
	}

	var r0 types.Market
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (types.Market, error)); ok {
		return rf(ctx, tickerStr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) types.Market); ok {
		r0 = rf(ctx, tickerStr)
	} else {
		r0 = ret.Get(0).(types.Market)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tickerStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarketMapKeeper_GetMarket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMarket'
type MarketMapKeeper_GetMarket_Call struct {
	*mock.Call
}

// GetMarket is a helper method to define mock.On call
//   - ctx context.Context
//   - tickerStr string
func (_e *MarketMapKeeper_Expecter) GetMarket(ctx interface{}, tickerStr interface{}) *MarketMapKeeper_GetMarket_Call {
	return &MarketMapKeeper_GetMarket_Call{Call: _e.mock.On("GetMarket", ctx, tickerStr)}
}

func (_c *MarketMapKeeper_GetMarket_Call) Run(run func(ctx context.Context, tickerStr string)) *MarketMapKeeper_GetMarket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MarketMapKeeper_GetMarket_Call) Return(_a0 types.Market, _a1 error) *MarketMapKeeper_GetMarket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MarketMapKeeper_GetMarket_Call) RunAndReturn(run func(context.Context, string) (types.Market, error)) *MarketMapKeeper_GetMarket_Call {
	_c.Call.Return(run)
	return _c
}

// NewMarketMapKeeper creates a new instance of MarketMapKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMarketMapKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *MarketMapKeeper {
	mock := &MarketMapKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}