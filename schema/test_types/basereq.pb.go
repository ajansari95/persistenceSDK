// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/comos/base/v1beta1/basereq.proto

package test_types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BaseReq struct {
	From          string           `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Memo          string           `protobuf:"bytes,2,opt,name=memo,proto3" json:"memo,omitempty"`
	ChainId       string           `protobuf:"bytes,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AccountNumber uint64           `protobuf:"varint,4,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
	Sequence      uint64           `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	TimeoutHeight uint64           `protobuf:"varint,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	Fees          []*types.Coin    `protobuf:"bytes,7,rep,name=fees,proto3" json:"fees,omitempty"`
	GasPrices     []*types.DecCoin `protobuf:"bytes,8,rep,name=gas_prices,json=gasPrices,proto3" json:"gas_prices,omitempty"`
	Gas           string           `protobuf:"bytes,9,opt,name=gas,proto3" json:"gas,omitempty"`
	GasAdjustment string           `protobuf:"bytes,10,opt,name=gas_adjustment,json=gasAdjustment,proto3" json:"gas_adjustment,omitempty"`
	Simulate      bool             `protobuf:"varint,11,opt,name=simulate,proto3" json:"simulate,omitempty"`
}

func (m *BaseReq) Reset()         { *m = BaseReq{} }
func (m *BaseReq) String() string { return proto.CompactTextString(m) }
func (*BaseReq) ProtoMessage()    {}
func (*BaseReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_57919bc98f071adc, []int{0}
}
func (m *BaseReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseReq.Merge(m, src)
}
func (m *BaseReq) XXX_Size() int {
	return m.Size()
}
func (m *BaseReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseReq.DiscardUnknown(m)
}

var xxx_messageInfo_BaseReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseReq)(nil), "cosmos.base.v1beta1.BaseReq")
}

func init() {
	proto.RegisterFile("proto/comos/base/v1beta1/basereq.proto", fileDescriptor_57919bc98f071adc)
}

var fileDescriptor_57919bc98f071adc = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x8f, 0x94, 0x40,
	0x1c, 0xc5, 0x41, 0xf0, 0x96, 0x9d, 0xcb, 0x19, 0x1d, 0x9b, 0xb9, 0x8d, 0xe1, 0x88, 0xc6, 0x0b,
	0x89, 0x11, 0x72, 0xda, 0x69, 0xe5, 0x69, 0xa1, 0x8d, 0x31, 0x94, 0x36, 0x64, 0x18, 0xfe, 0x07,
	0x63, 0x1c, 0x86, 0xe5, 0x3f, 0x98, 0x5c, 0x69, 0x67, 0xe9, 0x47, 0xb8, 0x8f, 0x63, 0x79, 0xa5,
	0xa5, 0xd9, 0x6d, 0xfc, 0x18, 0x86, 0x3f, 0x9b, 0xb5, 0xb9, 0xeb, 0xde, 0xfb, 0xf1, 0xde, 0x83,
	0x61, 0xd8, 0x69, 0x3f, 0x58, 0x67, 0x73, 0x65, 0x8d, 0xc5, 0xbc, 0x92, 0x08, 0xf9, 0xb7, 0xb3,
	0x0a, 0x9c, 0x3c, 0x23, 0x33, 0xc0, 0x3a, 0xa3, 0x00, 0x7f, 0xa8, 0x2c, 0x1a, 0x8b, 0xd9, 0x44,
	0xb3, 0x5d, 0x64, 0x75, 0xea, 0x5a, 0x3d, 0xd4, 0x65, 0x2f, 0x07, 0x77, 0x99, 0xcf, 0x43, 0x8d,
	0x6d, 0xec, 0x7f, 0x35, 0x97, 0x57, 0x4f, 0x6e, 0x7d, 0x89, 0xb2, 0xba, 0x9b, 0x43, 0x8f, 0xbf,
	0x07, 0x6c, 0x71, 0x2e, 0x11, 0x0a, 0x58, 0x73, 0xce, 0xc2, 0x8b, 0xc1, 0x1a, 0xe1, 0x27, 0x7e,
	0xba, 0x2c, 0x48, 0x4f, 0xcc, 0x80, 0xb1, 0xe2, 0xce, 0xcc, 0x26, 0xcd, 0x8f, 0x59, 0xa4, 0x5a,
	0xa9, 0xbb, 0x52, 0xd7, 0x22, 0x20, 0xbe, 0x20, 0xff, 0xa1, 0xe6, 0x4f, 0xd9, 0x3d, 0xa9, 0x94,
	0x1d, 0x3b, 0x57, 0x76, 0xa3, 0xa9, 0x60, 0x10, 0x61, 0xe2, 0xa7, 0x61, 0x71, 0xb4, 0xa3, 0x1f,
	0x09, 0xf2, 0x15, 0x8b, 0x10, 0xd6, 0x23, 0x74, 0x0a, 0xc4, 0x5d, 0x0a, 0xec, 0xfd, 0x34, 0xe1,
	0xb4, 0x01, 0x3b, 0xba, 0xb2, 0x05, 0xdd, 0xb4, 0x4e, 0x1c, 0xcc, 0x13, 0x3b, 0xfa, 0x9e, 0x20,
	0x7f, 0xce, 0xc2, 0x0b, 0x00, 0x14, 0x8b, 0x24, 0x48, 0x0f, 0x5f, 0x1c, 0x67, 0x37, 0xfc, 0xa9,
	0xec, 0xad, 0xd5, 0x5d, 0x41, 0x31, 0xfe, 0x9a, 0xb1, 0x46, 0x62, 0xd9, 0x0f, 0x5a, 0x01, 0x8a,
	0x88, 0x4a, 0x8f, 0x6e, 0x2c, 0xbd, 0x03, 0x45, 0xbd, 0x65, 0x23, 0xf1, 0x13, 0xc5, 0xf9, 0x7d,
	0x16, 0x34, 0x12, 0xc5, 0x92, 0xce, 0x3a, 0xc9, 0xe9, 0x23, 0xa7, 0x39, 0x59, 0x7f, 0x19, 0xd1,
	0x19, 0xe8, 0x9c, 0x60, 0xf4, 0xf0, 0xa8, 0x91, 0xf8, 0x66, 0x0f, 0xe9, 0x9c, 0xda, 0x8c, 0x5f,
	0xa5, 0x03, 0x71, 0x98, 0xf8, 0x69, 0x54, 0xec, 0xfd, 0xab, 0xe8, 0xc7, 0xd5, 0x89, 0xf7, 0xf7,
	0xea, 0xc4, 0x3b, 0x7f, 0xf6, 0x6b, 0x13, 0xfb, 0xd7, 0x9b, 0xd8, 0xff, 0xb3, 0x89, 0xfd, 0x9f,
	0xdb, 0xd8, 0xbb, 0xde, 0xc6, 0xde, 0xef, 0x6d, 0xec, 0x7d, 0x7e, 0x80, 0xaa, 0x05, 0x23, 0x73,
	0x07, 0xe8, 0x4a, 0x77, 0xd9, 0x03, 0x56, 0x07, 0x74, 0x6f, 0x2f, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0xb4, 0x9d, 0x60, 0x43, 0x02, 0x00, 0x00,
}

func (m *BaseReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Simulate {
		i--
		if m.Simulate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.GasAdjustment) > 0 {
		i -= len(m.GasAdjustment)
		copy(dAtA[i:], m.GasAdjustment)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.GasAdjustment)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Gas) > 0 {
		i -= len(m.Gas)
		copy(dAtA[i:], m.Gas)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.Gas)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.GasPrices) > 0 {
		for iNdEx := len(m.GasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBasereq(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBasereq(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintBasereq(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.Sequence != 0 {
		i = encodeVarintBasereq(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x28
	}
	if m.AccountNumber != 0 {
		i = encodeVarintBasereq(dAtA, i, uint64(m.AccountNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintBasereq(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBasereq(dAtA []byte, offset int, v uint64) int {
	offset -= sovBasereq(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	if m.AccountNumber != 0 {
		n += 1 + sovBasereq(uint64(m.AccountNumber))
	}
	if m.Sequence != 0 {
		n += 1 + sovBasereq(uint64(m.Sequence))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovBasereq(uint64(m.TimeoutHeight))
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovBasereq(uint64(l))
		}
	}
	if len(m.GasPrices) > 0 {
		for _, e := range m.GasPrices {
			l = e.Size()
			n += 1 + l + sovBasereq(uint64(l))
		}
	}
	l = len(m.Gas)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	l = len(m.GasAdjustment)
	if l > 0 {
		n += 1 + l + sovBasereq(uint64(l))
	}
	if m.Simulate {
		n += 2
	}
	return n
}

func sovBasereq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBasereq(x uint64) (n int) {
	return sovBasereq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBasereq
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BaseReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountNumber", wireType)
			}
			m.AccountNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AccountNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, &types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrices = append(m.GasPrices, &types.DecCoin{})
			if err := m.GasPrices[len(m.GasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gas = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasAdjustment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBasereq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBasereq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasAdjustment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Simulate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Simulate = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBasereq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBasereq
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBasereq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBasereq
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBasereq
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBasereq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBasereq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBasereq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBasereq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBasereq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBasereq = fmt.Errorf("proto: unexpected end of group")
)
