// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	handlers "github.com/skip-mev/connect/v2/providers/base/websocket/handlers"

	types "github.com/skip-mev/connect/v2/providers/types"
)

// WebSocketDataHandler is an autogenerated mock type for the WebSocketDataHandler type
type WebSocketDataHandler[K types.ResponseKey, V types.ResponseValue] struct {
	mock.Mock
}

type WebSocketDataHandler_Expecter[K types.ResponseKey, V types.ResponseValue] struct {
	mock *mock.Mock
}

func (_m *WebSocketDataHandler[K, V]) EXPECT() *WebSocketDataHandler_Expecter[K, V] {
	return &WebSocketDataHandler_Expecter[K, V]{mock: &_m.Mock}
}

// Copy provides a mock function with given fields:
func (_m *WebSocketDataHandler[K, V]) Copy() handlers.WebSocketDataHandler[K, V] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Copy")
	}

	var r0 handlers.WebSocketDataHandler[K, V]
	if rf, ok := ret.Get(0).(func() handlers.WebSocketDataHandler[K, V]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(handlers.WebSocketDataHandler[K, V])
		}
	}

	return r0
}

// WebSocketDataHandler_Copy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Copy'
type WebSocketDataHandler_Copy_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// Copy is a helper method to define mock.On call
func (_e *WebSocketDataHandler_Expecter[K, V]) Copy() *WebSocketDataHandler_Copy_Call[K, V] {
	return &WebSocketDataHandler_Copy_Call[K, V]{Call: _e.mock.On("Copy")}
}

func (_c *WebSocketDataHandler_Copy_Call[K, V]) Run(run func()) *WebSocketDataHandler_Copy_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WebSocketDataHandler_Copy_Call[K, V]) Return(_a0 handlers.WebSocketDataHandler[K, V]) *WebSocketDataHandler_Copy_Call[K, V] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WebSocketDataHandler_Copy_Call[K, V]) RunAndReturn(run func() handlers.WebSocketDataHandler[K, V]) *WebSocketDataHandler_Copy_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// CreateMessages provides a mock function with given fields: ids
func (_m *WebSocketDataHandler[K, V]) CreateMessages(ids []K) ([]handlers.WebsocketEncodedMessage, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessages")
	}

	var r0 []handlers.WebsocketEncodedMessage
	var r1 error
	if rf, ok := ret.Get(0).(func([]K) ([]handlers.WebsocketEncodedMessage, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]K) []handlers.WebsocketEncodedMessage); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]handlers.WebsocketEncodedMessage)
		}
	}

	if rf, ok := ret.Get(1).(func([]K) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebSocketDataHandler_CreateMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMessages'
type WebSocketDataHandler_CreateMessages_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// CreateMessages is a helper method to define mock.On call
//   - ids []K
func (_e *WebSocketDataHandler_Expecter[K, V]) CreateMessages(ids interface{}) *WebSocketDataHandler_CreateMessages_Call[K, V] {
	return &WebSocketDataHandler_CreateMessages_Call[K, V]{Call: _e.mock.On("CreateMessages", ids)}
}

func (_c *WebSocketDataHandler_CreateMessages_Call[K, V]) Run(run func(ids []K)) *WebSocketDataHandler_CreateMessages_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]K))
	})
	return _c
}

func (_c *WebSocketDataHandler_CreateMessages_Call[K, V]) Return(_a0 []handlers.WebsocketEncodedMessage, _a1 error) *WebSocketDataHandler_CreateMessages_Call[K, V] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebSocketDataHandler_CreateMessages_Call[K, V]) RunAndReturn(run func([]K) ([]handlers.WebsocketEncodedMessage, error)) *WebSocketDataHandler_CreateMessages_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// HandleMessage provides a mock function with given fields: message
func (_m *WebSocketDataHandler[K, V]) HandleMessage(message []byte) (types.GetResponse[K, V], []handlers.WebsocketEncodedMessage, error) {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for HandleMessage")
	}

	var r0 types.GetResponse[K, V]
	var r1 []handlers.WebsocketEncodedMessage
	var r2 error
	if rf, ok := ret.Get(0).(func([]byte) (types.GetResponse[K, V], []handlers.WebsocketEncodedMessage, error)); ok {
		return rf(message)
	}
	if rf, ok := ret.Get(0).(func([]byte) types.GetResponse[K, V]); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Get(0).(types.GetResponse[K, V])
	}

	if rf, ok := ret.Get(1).(func([]byte) []handlers.WebsocketEncodedMessage); ok {
		r1 = rf(message)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]handlers.WebsocketEncodedMessage)
		}
	}

	if rf, ok := ret.Get(2).(func([]byte) error); ok {
		r2 = rf(message)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// WebSocketDataHandler_HandleMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HandleMessage'
type WebSocketDataHandler_HandleMessage_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// HandleMessage is a helper method to define mock.On call
//   - message []byte
func (_e *WebSocketDataHandler_Expecter[K, V]) HandleMessage(message interface{}) *WebSocketDataHandler_HandleMessage_Call[K, V] {
	return &WebSocketDataHandler_HandleMessage_Call[K, V]{Call: _e.mock.On("HandleMessage", message)}
}

func (_c *WebSocketDataHandler_HandleMessage_Call[K, V]) Run(run func(message []byte)) *WebSocketDataHandler_HandleMessage_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *WebSocketDataHandler_HandleMessage_Call[K, V]) Return(response types.GetResponse[K, V], updateMessages []handlers.WebsocketEncodedMessage, err error) *WebSocketDataHandler_HandleMessage_Call[K, V] {
	_c.Call.Return(response, updateMessages, err)
	return _c
}

func (_c *WebSocketDataHandler_HandleMessage_Call[K, V]) RunAndReturn(run func([]byte) (types.GetResponse[K, V], []handlers.WebsocketEncodedMessage, error)) *WebSocketDataHandler_HandleMessage_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// HeartBeatMessages provides a mock function with given fields:
func (_m *WebSocketDataHandler[K, V]) HeartBeatMessages() ([]handlers.WebsocketEncodedMessage, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for HeartBeatMessages")
	}

	var r0 []handlers.WebsocketEncodedMessage
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]handlers.WebsocketEncodedMessage, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []handlers.WebsocketEncodedMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]handlers.WebsocketEncodedMessage)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebSocketDataHandler_HeartBeatMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeartBeatMessages'
type WebSocketDataHandler_HeartBeatMessages_Call[K types.ResponseKey, V types.ResponseValue] struct {
	*mock.Call
}

// HeartBeatMessages is a helper method to define mock.On call
func (_e *WebSocketDataHandler_Expecter[K, V]) HeartBeatMessages() *WebSocketDataHandler_HeartBeatMessages_Call[K, V] {
	return &WebSocketDataHandler_HeartBeatMessages_Call[K, V]{Call: _e.mock.On("HeartBeatMessages")}
}

func (_c *WebSocketDataHandler_HeartBeatMessages_Call[K, V]) Run(run func()) *WebSocketDataHandler_HeartBeatMessages_Call[K, V] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WebSocketDataHandler_HeartBeatMessages_Call[K, V]) Return(_a0 []handlers.WebsocketEncodedMessage, _a1 error) *WebSocketDataHandler_HeartBeatMessages_Call[K, V] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *WebSocketDataHandler_HeartBeatMessages_Call[K, V]) RunAndReturn(run func() ([]handlers.WebsocketEncodedMessage, error)) *WebSocketDataHandler_HeartBeatMessages_Call[K, V] {
	_c.Call.Return(run)
	return _c
}

// NewWebSocketDataHandler creates a new instance of WebSocketDataHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWebSocketDataHandler[K types.ResponseKey, V types.ResponseValue](t interface {
	mock.TestingT
	Cleanup(func())
}) *WebSocketDataHandler[K, V] {
	mock := &WebSocketDataHandler[K, V]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}