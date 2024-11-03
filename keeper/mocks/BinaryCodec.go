// Code generated by mockery v2.46.3. DO NOT EDIT.

package keeper_test

import (
	proto "github.com/cosmos/gogoproto/proto"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/codec/types"
)

// BinaryCodec is an autogenerated mock type for the BinaryCodec type
type BinaryCodec struct {
	mock.Mock
}

type BinaryCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *BinaryCodec) EXPECT() *BinaryCodec_Expecter {
	return &BinaryCodec_Expecter{mock: &_m.Mock}
}

// Marshal provides a mock function with given fields: o
func (_m *BinaryCodec) Marshal(o proto.Message) ([]byte, error) {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for Marshal")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(proto.Message) ([]byte, error)); ok {
		return rf(o)
	}
	if rf, ok := ret.Get(0).(func(proto.Message) []byte); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(proto.Message) error); ok {
		r1 = rf(o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BinaryCodec_Marshal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Marshal'
type BinaryCodec_Marshal_Call struct {
	*mock.Call
}

// Marshal is a helper method to define mock.On call
//   - o proto.Message
func (_e *BinaryCodec_Expecter) Marshal(o interface{}) *BinaryCodec_Marshal_Call {
	return &BinaryCodec_Marshal_Call{Call: _e.mock.On("Marshal", o)}
}

func (_c *BinaryCodec_Marshal_Call) Run(run func(o proto.Message)) *BinaryCodec_Marshal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_Marshal_Call) Return(_a0 []byte, _a1 error) *BinaryCodec_Marshal_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BinaryCodec_Marshal_Call) RunAndReturn(run func(proto.Message) ([]byte, error)) *BinaryCodec_Marshal_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalInterface provides a mock function with given fields: i
func (_m *BinaryCodec) MarshalInterface(i proto.Message) ([]byte, error) {
	ret := _m.Called(i)

	if len(ret) == 0 {
		panic("no return value specified for MarshalInterface")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(proto.Message) ([]byte, error)); ok {
		return rf(i)
	}
	if rf, ok := ret.Get(0).(func(proto.Message) []byte); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(proto.Message) error); ok {
		r1 = rf(i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BinaryCodec_MarshalInterface_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalInterface'
type BinaryCodec_MarshalInterface_Call struct {
	*mock.Call
}

// MarshalInterface is a helper method to define mock.On call
//   - i proto.Message
func (_e *BinaryCodec_Expecter) MarshalInterface(i interface{}) *BinaryCodec_MarshalInterface_Call {
	return &BinaryCodec_MarshalInterface_Call{Call: _e.mock.On("MarshalInterface", i)}
}

func (_c *BinaryCodec_MarshalInterface_Call) Run(run func(i proto.Message)) *BinaryCodec_MarshalInterface_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_MarshalInterface_Call) Return(_a0 []byte, _a1 error) *BinaryCodec_MarshalInterface_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BinaryCodec_MarshalInterface_Call) RunAndReturn(run func(proto.Message) ([]byte, error)) *BinaryCodec_MarshalInterface_Call {
	_c.Call.Return(run)
	return _c
}

// MarshalLengthPrefixed provides a mock function with given fields: o
func (_m *BinaryCodec) MarshalLengthPrefixed(o proto.Message) ([]byte, error) {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for MarshalLengthPrefixed")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(proto.Message) ([]byte, error)); ok {
		return rf(o)
	}
	if rf, ok := ret.Get(0).(func(proto.Message) []byte); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(proto.Message) error); ok {
		r1 = rf(o)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BinaryCodec_MarshalLengthPrefixed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarshalLengthPrefixed'
type BinaryCodec_MarshalLengthPrefixed_Call struct {
	*mock.Call
}

// MarshalLengthPrefixed is a helper method to define mock.On call
//   - o proto.Message
func (_e *BinaryCodec_Expecter) MarshalLengthPrefixed(o interface{}) *BinaryCodec_MarshalLengthPrefixed_Call {
	return &BinaryCodec_MarshalLengthPrefixed_Call{Call: _e.mock.On("MarshalLengthPrefixed", o)}
}

func (_c *BinaryCodec_MarshalLengthPrefixed_Call) Run(run func(o proto.Message)) *BinaryCodec_MarshalLengthPrefixed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_MarshalLengthPrefixed_Call) Return(_a0 []byte, _a1 error) *BinaryCodec_MarshalLengthPrefixed_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BinaryCodec_MarshalLengthPrefixed_Call) RunAndReturn(run func(proto.Message) ([]byte, error)) *BinaryCodec_MarshalLengthPrefixed_Call {
	_c.Call.Return(run)
	return _c
}

// MustMarshal provides a mock function with given fields: o
func (_m *BinaryCodec) MustMarshal(o proto.Message) []byte {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for MustMarshal")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func(proto.Message) []byte); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// BinaryCodec_MustMarshal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustMarshal'
type BinaryCodec_MustMarshal_Call struct {
	*mock.Call
}

// MustMarshal is a helper method to define mock.On call
//   - o proto.Message
func (_e *BinaryCodec_Expecter) MustMarshal(o interface{}) *BinaryCodec_MustMarshal_Call {
	return &BinaryCodec_MustMarshal_Call{Call: _e.mock.On("MustMarshal", o)}
}

func (_c *BinaryCodec_MustMarshal_Call) Run(run func(o proto.Message)) *BinaryCodec_MustMarshal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_MustMarshal_Call) Return(_a0 []byte) *BinaryCodec_MustMarshal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BinaryCodec_MustMarshal_Call) RunAndReturn(run func(proto.Message) []byte) *BinaryCodec_MustMarshal_Call {
	_c.Call.Return(run)
	return _c
}

// MustMarshalLengthPrefixed provides a mock function with given fields: o
func (_m *BinaryCodec) MustMarshalLengthPrefixed(o proto.Message) []byte {
	ret := _m.Called(o)

	if len(ret) == 0 {
		panic("no return value specified for MustMarshalLengthPrefixed")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func(proto.Message) []byte); ok {
		r0 = rf(o)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// BinaryCodec_MustMarshalLengthPrefixed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustMarshalLengthPrefixed'
type BinaryCodec_MustMarshalLengthPrefixed_Call struct {
	*mock.Call
}

// MustMarshalLengthPrefixed is a helper method to define mock.On call
//   - o proto.Message
func (_e *BinaryCodec_Expecter) MustMarshalLengthPrefixed(o interface{}) *BinaryCodec_MustMarshalLengthPrefixed_Call {
	return &BinaryCodec_MustMarshalLengthPrefixed_Call{Call: _e.mock.On("MustMarshalLengthPrefixed", o)}
}

func (_c *BinaryCodec_MustMarshalLengthPrefixed_Call) Run(run func(o proto.Message)) *BinaryCodec_MustMarshalLengthPrefixed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_MustMarshalLengthPrefixed_Call) Return(_a0 []byte) *BinaryCodec_MustMarshalLengthPrefixed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BinaryCodec_MustMarshalLengthPrefixed_Call) RunAndReturn(run func(proto.Message) []byte) *BinaryCodec_MustMarshalLengthPrefixed_Call {
	_c.Call.Return(run)
	return _c
}

// MustUnmarshal provides a mock function with given fields: bz, ptr
func (_m *BinaryCodec) MustUnmarshal(bz []byte, ptr proto.Message) {
	_m.Called(bz, ptr)
}

// BinaryCodec_MustUnmarshal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustUnmarshal'
type BinaryCodec_MustUnmarshal_Call struct {
	*mock.Call
}

// MustUnmarshal is a helper method to define mock.On call
//   - bz []byte
//   - ptr proto.Message
func (_e *BinaryCodec_Expecter) MustUnmarshal(bz interface{}, ptr interface{}) *BinaryCodec_MustUnmarshal_Call {
	return &BinaryCodec_MustUnmarshal_Call{Call: _e.mock.On("MustUnmarshal", bz, ptr)}
}

func (_c *BinaryCodec_MustUnmarshal_Call) Run(run func(bz []byte, ptr proto.Message)) *BinaryCodec_MustUnmarshal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_MustUnmarshal_Call) Return() *BinaryCodec_MustUnmarshal_Call {
	_c.Call.Return()
	return _c
}

func (_c *BinaryCodec_MustUnmarshal_Call) RunAndReturn(run func([]byte, proto.Message)) *BinaryCodec_MustUnmarshal_Call {
	_c.Call.Return(run)
	return _c
}

// MustUnmarshalLengthPrefixed provides a mock function with given fields: bz, ptr
func (_m *BinaryCodec) MustUnmarshalLengthPrefixed(bz []byte, ptr proto.Message) {
	_m.Called(bz, ptr)
}

// BinaryCodec_MustUnmarshalLengthPrefixed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustUnmarshalLengthPrefixed'
type BinaryCodec_MustUnmarshalLengthPrefixed_Call struct {
	*mock.Call
}

// MustUnmarshalLengthPrefixed is a helper method to define mock.On call
//   - bz []byte
//   - ptr proto.Message
func (_e *BinaryCodec_Expecter) MustUnmarshalLengthPrefixed(bz interface{}, ptr interface{}) *BinaryCodec_MustUnmarshalLengthPrefixed_Call {
	return &BinaryCodec_MustUnmarshalLengthPrefixed_Call{Call: _e.mock.On("MustUnmarshalLengthPrefixed", bz, ptr)}
}

func (_c *BinaryCodec_MustUnmarshalLengthPrefixed_Call) Run(run func(bz []byte, ptr proto.Message)) *BinaryCodec_MustUnmarshalLengthPrefixed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_MustUnmarshalLengthPrefixed_Call) Return() *BinaryCodec_MustUnmarshalLengthPrefixed_Call {
	_c.Call.Return()
	return _c
}

func (_c *BinaryCodec_MustUnmarshalLengthPrefixed_Call) RunAndReturn(run func([]byte, proto.Message)) *BinaryCodec_MustUnmarshalLengthPrefixed_Call {
	_c.Call.Return(run)
	return _c
}

// Unmarshal provides a mock function with given fields: bz, ptr
func (_m *BinaryCodec) Unmarshal(bz []byte, ptr proto.Message) error {
	ret := _m.Called(bz, ptr)

	if len(ret) == 0 {
		panic("no return value specified for Unmarshal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, proto.Message) error); ok {
		r0 = rf(bz, ptr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BinaryCodec_Unmarshal_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unmarshal'
type BinaryCodec_Unmarshal_Call struct {
	*mock.Call
}

// Unmarshal is a helper method to define mock.On call
//   - bz []byte
//   - ptr proto.Message
func (_e *BinaryCodec_Expecter) Unmarshal(bz interface{}, ptr interface{}) *BinaryCodec_Unmarshal_Call {
	return &BinaryCodec_Unmarshal_Call{Call: _e.mock.On("Unmarshal", bz, ptr)}
}

func (_c *BinaryCodec_Unmarshal_Call) Run(run func(bz []byte, ptr proto.Message)) *BinaryCodec_Unmarshal_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_Unmarshal_Call) Return(_a0 error) *BinaryCodec_Unmarshal_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BinaryCodec_Unmarshal_Call) RunAndReturn(run func([]byte, proto.Message) error) *BinaryCodec_Unmarshal_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalInterface provides a mock function with given fields: bz, ptr
func (_m *BinaryCodec) UnmarshalInterface(bz []byte, ptr interface{}) error {
	ret := _m.Called(bz, ptr)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalInterface")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(bz, ptr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BinaryCodec_UnmarshalInterface_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalInterface'
type BinaryCodec_UnmarshalInterface_Call struct {
	*mock.Call
}

// UnmarshalInterface is a helper method to define mock.On call
//   - bz []byte
//   - ptr interface{}
func (_e *BinaryCodec_Expecter) UnmarshalInterface(bz interface{}, ptr interface{}) *BinaryCodec_UnmarshalInterface_Call {
	return &BinaryCodec_UnmarshalInterface_Call{Call: _e.mock.On("UnmarshalInterface", bz, ptr)}
}

func (_c *BinaryCodec_UnmarshalInterface_Call) Run(run func(bz []byte, ptr interface{})) *BinaryCodec_UnmarshalInterface_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(interface{}))
	})
	return _c
}

func (_c *BinaryCodec_UnmarshalInterface_Call) Return(_a0 error) *BinaryCodec_UnmarshalInterface_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BinaryCodec_UnmarshalInterface_Call) RunAndReturn(run func([]byte, interface{}) error) *BinaryCodec_UnmarshalInterface_Call {
	_c.Call.Return(run)
	return _c
}

// UnmarshalLengthPrefixed provides a mock function with given fields: bz, ptr
func (_m *BinaryCodec) UnmarshalLengthPrefixed(bz []byte, ptr proto.Message) error {
	ret := _m.Called(bz, ptr)

	if len(ret) == 0 {
		panic("no return value specified for UnmarshalLengthPrefixed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, proto.Message) error); ok {
		r0 = rf(bz, ptr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BinaryCodec_UnmarshalLengthPrefixed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnmarshalLengthPrefixed'
type BinaryCodec_UnmarshalLengthPrefixed_Call struct {
	*mock.Call
}

// UnmarshalLengthPrefixed is a helper method to define mock.On call
//   - bz []byte
//   - ptr proto.Message
func (_e *BinaryCodec_Expecter) UnmarshalLengthPrefixed(bz interface{}, ptr interface{}) *BinaryCodec_UnmarshalLengthPrefixed_Call {
	return &BinaryCodec_UnmarshalLengthPrefixed_Call{Call: _e.mock.On("UnmarshalLengthPrefixed", bz, ptr)}
}

func (_c *BinaryCodec_UnmarshalLengthPrefixed_Call) Run(run func(bz []byte, ptr proto.Message)) *BinaryCodec_UnmarshalLengthPrefixed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(proto.Message))
	})
	return _c
}

func (_c *BinaryCodec_UnmarshalLengthPrefixed_Call) Return(_a0 error) *BinaryCodec_UnmarshalLengthPrefixed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BinaryCodec_UnmarshalLengthPrefixed_Call) RunAndReturn(run func([]byte, proto.Message) error) *BinaryCodec_UnmarshalLengthPrefixed_Call {
	_c.Call.Return(run)
	return _c
}

// UnpackAny provides a mock function with given fields: any, iface
func (_m *BinaryCodec) UnpackAny(any *types.Any, iface interface{}) error {
	ret := _m.Called(any, iface)

	if len(ret) == 0 {
		panic("no return value specified for UnpackAny")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Any, interface{}) error); ok {
		r0 = rf(any, iface)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BinaryCodec_UnpackAny_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnpackAny'
type BinaryCodec_UnpackAny_Call struct {
	*mock.Call
}

// UnpackAny is a helper method to define mock.On call
//   - any *types.Any
//   - iface interface{}
func (_e *BinaryCodec_Expecter) UnpackAny(any interface{}, iface interface{}) *BinaryCodec_UnpackAny_Call {
	return &BinaryCodec_UnpackAny_Call{Call: _e.mock.On("UnpackAny", any, iface)}
}

func (_c *BinaryCodec_UnpackAny_Call) Run(run func(any *types.Any, iface interface{})) *BinaryCodec_UnpackAny_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Any), args[1].(interface{}))
	})
	return _c
}

func (_c *BinaryCodec_UnpackAny_Call) Return(_a0 error) *BinaryCodec_UnpackAny_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BinaryCodec_UnpackAny_Call) RunAndReturn(run func(*types.Any, interface{}) error) *BinaryCodec_UnpackAny_Call {
	_c.Call.Return(run)
	return _c
}

// NewBinaryCodec creates a new instance of BinaryCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBinaryCodec(t interface {
	mock.TestingT
	Cleanup(func())
}) *BinaryCodec {
	mock := &BinaryCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
