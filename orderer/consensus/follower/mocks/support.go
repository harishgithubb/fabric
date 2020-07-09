// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	channelconfig "github.com/hyperledger/fabric/common/channelconfig"
	blockcutter "github.com/hyperledger/fabric/orderer/common/blockcutter"

	common "github.com/hyperledger/fabric-protos-go/common"

	mock "github.com/stretchr/testify/mock"

	msgprocessor "github.com/hyperledger/fabric/orderer/common/msgprocessor"

	protoutil "github.com/hyperledger/fabric/protoutil"
)

// Support is an autogenerated mock type for the Support type
type Support struct {
	mock.Mock
}

// Append provides a mock function with given fields: block
func (_m *Support) Append(block *common.Block) error {
	ret := _m.Called(block)

	var r0 error
	if rf, ok := ret.Get(0).(func(*common.Block) error); ok {
		r0 = rf(block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Block provides a mock function with given fields: number
func (_m *Support) Block(number uint64) *common.Block {
	ret := _m.Called(number)

	var r0 *common.Block
	if rf, ok := ret.Get(0).(func(uint64) *common.Block); ok {
		r0 = rf(number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Block)
		}
	}

	return r0
}

// BlockCutter provides a mock function with given fields:
func (_m *Support) BlockCutter() blockcutter.Receiver {
	ret := _m.Called()

	var r0 blockcutter.Receiver
	if rf, ok := ret.Get(0).(func() blockcutter.Receiver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blockcutter.Receiver)
		}
	}

	return r0
}

// ChannelConfig provides a mock function with given fields:
func (_m *Support) ChannelConfig() channelconfig.Channel {
	ret := _m.Called()

	var r0 channelconfig.Channel
	if rf, ok := ret.Get(0).(func() channelconfig.Channel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(channelconfig.Channel)
		}
	}

	return r0
}

// ChannelID provides a mock function with given fields:
func (_m *Support) ChannelID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ClassifyMsg provides a mock function with given fields: chdr
func (_m *Support) ClassifyMsg(chdr *common.ChannelHeader) msgprocessor.Classification {
	ret := _m.Called(chdr)

	var r0 msgprocessor.Classification
	if rf, ok := ret.Get(0).(func(*common.ChannelHeader) msgprocessor.Classification); ok {
		r0 = rf(chdr)
	} else {
		r0 = ret.Get(0).(msgprocessor.Classification)
	}

	return r0
}

// CreateNextBlock provides a mock function with given fields: messages
func (_m *Support) CreateNextBlock(messages []*common.Envelope) *common.Block {
	ret := _m.Called(messages)

	var r0 *common.Block
	if rf, ok := ret.Get(0).(func([]*common.Envelope) *common.Block); ok {
		r0 = rf(messages)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Block)
		}
	}

	return r0
}

// Height provides a mock function with given fields:
func (_m *Support) Height() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// ProcessConfigMsg provides a mock function with given fields: env
func (_m *Support) ProcessConfigMsg(env *common.Envelope) (*common.Envelope, uint64, error) {
	ret := _m.Called(env)

	var r0 *common.Envelope
	if rf, ok := ret.Get(0).(func(*common.Envelope) *common.Envelope); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Envelope)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*common.Envelope) uint64); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*common.Envelope) error); ok {
		r2 = rf(env)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProcessConfigUpdateMsg provides a mock function with given fields: env
func (_m *Support) ProcessConfigUpdateMsg(env *common.Envelope) (*common.Envelope, uint64, error) {
	ret := _m.Called(env)

	var r0 *common.Envelope
	if rf, ok := ret.Get(0).(func(*common.Envelope) *common.Envelope); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Envelope)
		}
	}

	var r1 uint64
	if rf, ok := ret.Get(1).(func(*common.Envelope) uint64); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*common.Envelope) error); ok {
		r2 = rf(env)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProcessNormalMsg provides a mock function with given fields: env
func (_m *Support) ProcessNormalMsg(env *common.Envelope) (uint64, error) {
	ret := _m.Called(env)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*common.Envelope) uint64); ok {
		r0 = rf(env)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*common.Envelope) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sequence provides a mock function with given fields:
func (_m *Support) Sequence() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Serialize provides a mock function with given fields:
func (_m *Support) Serialize() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SharedConfig provides a mock function with given fields:
func (_m *Support) SharedConfig() channelconfig.Orderer {
	ret := _m.Called()

	var r0 channelconfig.Orderer
	if rf, ok := ret.Get(0).(func() channelconfig.Orderer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(channelconfig.Orderer)
		}
	}

	return r0
}

// Sign provides a mock function with given fields: message
func (_m *Support) Sign(message []byte) ([]byte, error) {
	ret := _m.Called(message)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyBlockSignature provides a mock function with given fields: _a0, _a1
func (_m *Support) VerifyBlockSignature(_a0 []*protoutil.SignedData, _a1 *common.ConfigEnvelope) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*protoutil.SignedData, *common.ConfigEnvelope) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteBlock provides a mock function with given fields: block, encodedMetadataValue
func (_m *Support) WriteBlock(block *common.Block, encodedMetadataValue []byte) {
	_m.Called(block, encodedMetadataValue)
}

// WriteConfigBlock provides a mock function with given fields: block, encodedMetadataValue
func (_m *Support) WriteConfigBlock(block *common.Block, encodedMetadataValue []byte) {
	_m.Called(block, encodedMetadataValue)
}