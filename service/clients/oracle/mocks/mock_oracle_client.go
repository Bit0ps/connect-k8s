// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	types "github.com/skip-mev/connect/v2/service/servers/oracle/types"
)

// OracleClient is an autogenerated mock type for the OracleClient type
type OracleClient struct {
	mock.Mock
}

type OracleClient_Expecter struct {
	mock *mock.Mock
}

func (_m *OracleClient) EXPECT() *OracleClient_Expecter {
	return &OracleClient_Expecter{mock: &_m.Mock}
}

// MarketMap provides a mock function with given fields: ctx, in, opts
func (_m *OracleClient) MarketMap(ctx context.Context, in *types.QueryMarketMapRequest, opts ...grpc.CallOption) (*types.QueryMarketMapResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MarketMap")
	}

	var r0 *types.QueryMarketMapResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketMapRequest, ...grpc.CallOption) (*types.QueryMarketMapResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryMarketMapRequest, ...grpc.CallOption) *types.QueryMarketMapResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryMarketMapResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryMarketMapRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OracleClient_MarketMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarketMap'
type OracleClient_MarketMap_Call struct {
	*mock.Call
}

// MarketMap is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryMarketMapRequest
//   - opts ...grpc.CallOption
func (_e *OracleClient_Expecter) MarketMap(ctx interface{}, in interface{}, opts ...interface{}) *OracleClient_MarketMap_Call {
	return &OracleClient_MarketMap_Call{Call: _e.mock.On("MarketMap",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OracleClient_MarketMap_Call) Run(run func(ctx context.Context, in *types.QueryMarketMapRequest, opts ...grpc.CallOption)) *OracleClient_MarketMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryMarketMapRequest), variadicArgs...)
	})
	return _c
}

func (_c *OracleClient_MarketMap_Call) Return(_a0 *types.QueryMarketMapResponse, _a1 error) *OracleClient_MarketMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleClient_MarketMap_Call) RunAndReturn(run func(context.Context, *types.QueryMarketMapRequest, ...grpc.CallOption) (*types.QueryMarketMapResponse, error)) *OracleClient_MarketMap_Call {
	_c.Call.Return(run)
	return _c
}

// Prices provides a mock function with given fields: ctx, in, opts
func (_m *OracleClient) Prices(ctx context.Context, in *types.QueryPricesRequest, opts ...grpc.CallOption) (*types.QueryPricesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Prices")
	}

	var r0 *types.QueryPricesResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryPricesRequest, ...grpc.CallOption) (*types.QueryPricesResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryPricesRequest, ...grpc.CallOption) *types.QueryPricesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryPricesResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryPricesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OracleClient_Prices_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prices'
type OracleClient_Prices_Call struct {
	*mock.Call
}

// Prices is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryPricesRequest
//   - opts ...grpc.CallOption
func (_e *OracleClient_Expecter) Prices(ctx interface{}, in interface{}, opts ...interface{}) *OracleClient_Prices_Call {
	return &OracleClient_Prices_Call{Call: _e.mock.On("Prices",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OracleClient_Prices_Call) Run(run func(ctx context.Context, in *types.QueryPricesRequest, opts ...grpc.CallOption)) *OracleClient_Prices_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryPricesRequest), variadicArgs...)
	})
	return _c
}

func (_c *OracleClient_Prices_Call) Return(_a0 *types.QueryPricesResponse, _a1 error) *OracleClient_Prices_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleClient_Prices_Call) RunAndReturn(run func(context.Context, *types.QueryPricesRequest, ...grpc.CallOption) (*types.QueryPricesResponse, error)) *OracleClient_Prices_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0
func (_m *OracleClient) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OracleClient_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type OracleClient_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *OracleClient_Expecter) Start(_a0 interface{}) *OracleClient_Start_Call {
	return &OracleClient_Start_Call{Call: _e.mock.On("Start", _a0)}
}

func (_c *OracleClient_Start_Call) Run(run func(_a0 context.Context)) *OracleClient_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OracleClient_Start_Call) Return(_a0 error) *OracleClient_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OracleClient_Start_Call) RunAndReturn(run func(context.Context) error) *OracleClient_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields:
func (_m *OracleClient) Stop() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Stop")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OracleClient_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type OracleClient_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
func (_e *OracleClient_Expecter) Stop() *OracleClient_Stop_Call {
	return &OracleClient_Stop_Call{Call: _e.mock.On("Stop")}
}

func (_c *OracleClient_Stop_Call) Run(run func()) *OracleClient_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *OracleClient_Stop_Call) Return(_a0 error) *OracleClient_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OracleClient_Stop_Call) RunAndReturn(run func() error) *OracleClient_Stop_Call {
	_c.Call.Return(run)
	return _c
}

// Version provides a mock function with given fields: ctx, in, opts
func (_m *OracleClient) Version(ctx context.Context, in *types.QueryVersionRequest, opts ...grpc.CallOption) (*types.QueryVersionResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Version")
	}

	var r0 *types.QueryVersionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryVersionRequest, ...grpc.CallOption) (*types.QueryVersionResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.QueryVersionRequest, ...grpc.CallOption) *types.QueryVersionResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.QueryVersionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.QueryVersionRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OracleClient_Version_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Version'
type OracleClient_Version_Call struct {
	*mock.Call
}

// Version is a helper method to define mock.On call
//   - ctx context.Context
//   - in *types.QueryVersionRequest
//   - opts ...grpc.CallOption
func (_e *OracleClient_Expecter) Version(ctx interface{}, in interface{}, opts ...interface{}) *OracleClient_Version_Call {
	return &OracleClient_Version_Call{Call: _e.mock.On("Version",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *OracleClient_Version_Call) Run(run func(ctx context.Context, in *types.QueryVersionRequest, opts ...grpc.CallOption)) *OracleClient_Version_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*types.QueryVersionRequest), variadicArgs...)
	})
	return _c
}

func (_c *OracleClient_Version_Call) Return(_a0 *types.QueryVersionResponse, _a1 error) *OracleClient_Version_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OracleClient_Version_Call) RunAndReturn(run func(context.Context, *types.QueryVersionRequest, ...grpc.CallOption) (*types.QueryVersionResponse, error)) *OracleClient_Version_Call {
	_c.Call.Return(run)
	return _c
}

// NewOracleClient creates a new instance of OracleClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOracleClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *OracleClient {
	mock := &OracleClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}