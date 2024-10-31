// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	aggregator "github.com/skip-mev/connect/v2/abci/strategies/aggregator"

	mock "github.com/stretchr/testify/mock"

	pkgtypes "github.com/skip-mev/connect/v2/pkg/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// VoteAggregator is an autogenerated mock type for the VoteAggregator type
type VoteAggregator struct {
	mock.Mock
}

type VoteAggregator_Expecter struct {
	mock *mock.Mock
}

func (_m *VoteAggregator) EXPECT() *VoteAggregator_Expecter {
	return &VoteAggregator_Expecter{mock: &_m.Mock}
}

// AggregateOracleVotes provides a mock function with given fields: ctx, votes
func (_m *VoteAggregator) AggregateOracleVotes(ctx types.Context, votes []aggregator.Vote) (map[pkgtypes.CurrencyPair]*big.Int, error) {
	ret := _m.Called(ctx, votes)

	if len(ret) == 0 {
		panic("no return value specified for AggregateOracleVotes")
	}

	var r0 map[pkgtypes.CurrencyPair]*big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, []aggregator.Vote) (map[pkgtypes.CurrencyPair]*big.Int, error)); ok {
		return rf(ctx, votes)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []aggregator.Vote) map[pkgtypes.CurrencyPair]*big.Int); ok {
		r0 = rf(ctx, votes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[pkgtypes.CurrencyPair]*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, []aggregator.Vote) error); ok {
		r1 = rf(ctx, votes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VoteAggregator_AggregateOracleVotes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AggregateOracleVotes'
type VoteAggregator_AggregateOracleVotes_Call struct {
	*mock.Call
}

// AggregateOracleVotes is a helper method to define mock.On call
//   - ctx types.Context
//   - votes []aggregator.Vote
func (_e *VoteAggregator_Expecter) AggregateOracleVotes(ctx interface{}, votes interface{}) *VoteAggregator_AggregateOracleVotes_Call {
	return &VoteAggregator_AggregateOracleVotes_Call{Call: _e.mock.On("AggregateOracleVotes", ctx, votes)}
}

func (_c *VoteAggregator_AggregateOracleVotes_Call) Run(run func(ctx types.Context, votes []aggregator.Vote)) *VoteAggregator_AggregateOracleVotes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Context), args[1].([]aggregator.Vote))
	})
	return _c
}

func (_c *VoteAggregator_AggregateOracleVotes_Call) Return(_a0 map[pkgtypes.CurrencyPair]*big.Int, _a1 error) *VoteAggregator_AggregateOracleVotes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VoteAggregator_AggregateOracleVotes_Call) RunAndReturn(run func(types.Context, []aggregator.Vote) (map[pkgtypes.CurrencyPair]*big.Int, error)) *VoteAggregator_AggregateOracleVotes_Call {
	_c.Call.Return(run)
	return _c
}

// GetPriceForValidator provides a mock function with given fields: validator
func (_m *VoteAggregator) GetPriceForValidator(validator types.ConsAddress) map[pkgtypes.CurrencyPair]*big.Int {
	ret := _m.Called(validator)

	if len(ret) == 0 {
		panic("no return value specified for GetPriceForValidator")
	}

	var r0 map[pkgtypes.CurrencyPair]*big.Int
	if rf, ok := ret.Get(0).(func(types.ConsAddress) map[pkgtypes.CurrencyPair]*big.Int); ok {
		r0 = rf(validator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[pkgtypes.CurrencyPair]*big.Int)
		}
	}

	return r0
}

// VoteAggregator_GetPriceForValidator_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPriceForValidator'
type VoteAggregator_GetPriceForValidator_Call struct {
	*mock.Call
}

// GetPriceForValidator is a helper method to define mock.On call
//   - validator types.ConsAddress
func (_e *VoteAggregator_Expecter) GetPriceForValidator(validator interface{}) *VoteAggregator_GetPriceForValidator_Call {
	return &VoteAggregator_GetPriceForValidator_Call{Call: _e.mock.On("GetPriceForValidator", validator)}
}

func (_c *VoteAggregator_GetPriceForValidator_Call) Run(run func(validator types.ConsAddress)) *VoteAggregator_GetPriceForValidator_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.ConsAddress))
	})
	return _c
}

func (_c *VoteAggregator_GetPriceForValidator_Call) Return(_a0 map[pkgtypes.CurrencyPair]*big.Int) *VoteAggregator_GetPriceForValidator_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VoteAggregator_GetPriceForValidator_Call) RunAndReturn(run func(types.ConsAddress) map[pkgtypes.CurrencyPair]*big.Int) *VoteAggregator_GetPriceForValidator_Call {
	_c.Call.Return(run)
	return _c
}

// NewVoteAggregator creates a new instance of VoteAggregator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVoteAggregator(t interface {
	mock.TestingT
	Cleanup(func())
}) *VoteAggregator {
	mock := &VoteAggregator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}