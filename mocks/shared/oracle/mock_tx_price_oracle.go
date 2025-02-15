// Code generated by mockery v2.46.2. DO NOT EDIT.

package oracle

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// MockTxPriceOracle is an autogenerated mock type for the TxPriceOracle type
type MockTxPriceOracle struct {
	mock.Mock
}

type MockTxPriceOracle_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTxPriceOracle) EXPECT() *MockTxPriceOracle_Expecter {
	return &MockTxPriceOracle_Expecter{mock: &_m.Mock}
}

// GasCostUUSDC provides a mock function with given fields: ctx, txFee, chainID
func (_m *MockTxPriceOracle) GasCostUUSDC(ctx context.Context, txFee *big.Int, chainID string) (*big.Int, error) {
	ret := _m.Called(ctx, txFee, chainID)

	if len(ret) == 0 {
		panic("no return value specified for GasCostUUSDC")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, string) (*big.Int, error)); ok {
		return rf(ctx, txFee, chainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int, string) *big.Int); ok {
		r0 = rf(ctx, txFee, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int, string) error); ok {
		r1 = rf(ctx, txFee, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTxPriceOracle_GasCostUUSDC_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GasCostUUSDC'
type MockTxPriceOracle_GasCostUUSDC_Call struct {
	*mock.Call
}

// GasCostUUSDC is a helper method to define mock.On call
//   - ctx context.Context
//   - txFee *big.Int
//   - chainID string
func (_e *MockTxPriceOracle_Expecter) GasCostUUSDC(ctx interface{}, txFee interface{}, chainID interface{}) *MockTxPriceOracle_GasCostUUSDC_Call {
	return &MockTxPriceOracle_GasCostUUSDC_Call{Call: _e.mock.On("GasCostUUSDC", ctx, txFee, chainID)}
}

func (_c *MockTxPriceOracle_GasCostUUSDC_Call) Run(run func(ctx context.Context, txFee *big.Int, chainID string)) *MockTxPriceOracle_GasCostUUSDC_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int), args[2].(string))
	})
	return _c
}

func (_c *MockTxPriceOracle_GasCostUUSDC_Call) Return(_a0 *big.Int, _a1 error) *MockTxPriceOracle_GasCostUUSDC_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTxPriceOracle_GasCostUUSDC_Call) RunAndReturn(run func(context.Context, *big.Int, string) (*big.Int, error)) *MockTxPriceOracle_GasCostUUSDC_Call {
	_c.Call.Return(run)
	return _c
}

// TxFeeUUSDC provides a mock function with given fields: ctx, tx
func (_m *MockTxPriceOracle) TxFeeUUSDC(ctx context.Context, tx *types.Transaction) (*big.Int, error) {
	ret := _m.Called(ctx, tx)

	if len(ret) == 0 {
		panic("no return value specified for TxFeeUUSDC")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) (*big.Int, error)); ok {
		return rf(ctx, tx)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) *big.Int); ok {
		r0 = rf(ctx, tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Transaction) error); ok {
		r1 = rf(ctx, tx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTxPriceOracle_TxFeeUUSDC_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxFeeUUSDC'
type MockTxPriceOracle_TxFeeUUSDC_Call struct {
	*mock.Call
}

// TxFeeUUSDC is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
func (_e *MockTxPriceOracle_Expecter) TxFeeUUSDC(ctx interface{}, tx interface{}) *MockTxPriceOracle_TxFeeUUSDC_Call {
	return &MockTxPriceOracle_TxFeeUUSDC_Call{Call: _e.mock.On("TxFeeUUSDC", ctx, tx)}
}

func (_c *MockTxPriceOracle_TxFeeUUSDC_Call) Run(run func(ctx context.Context, tx *types.Transaction)) *MockTxPriceOracle_TxFeeUUSDC_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction))
	})
	return _c
}

func (_c *MockTxPriceOracle_TxFeeUUSDC_Call) Return(_a0 *big.Int, _a1 error) *MockTxPriceOracle_TxFeeUUSDC_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTxPriceOracle_TxFeeUUSDC_Call) RunAndReturn(run func(context.Context, *types.Transaction) (*big.Int, error)) *MockTxPriceOracle_TxFeeUUSDC_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTxPriceOracle creates a new instance of MockTxPriceOracle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTxPriceOracle(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTxPriceOracle {
	mock := &MockTxPriceOracle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
