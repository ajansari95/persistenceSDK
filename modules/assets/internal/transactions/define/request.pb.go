// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/assets/transactions/define/request.proto

package define

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	_ "github.com/regen-network/cosmos-proto"
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

type transactionRequest struct {
	BaseReq                 protoTypes.BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq,proto3" json:"base_req" `
	FromID                  string             `protobuf:"bytes,2,opt,name=from_iD,json=fromID,proto3" json:"from_iD,omitempty" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ImmutableMetaProperties string             `protobuf:"bytes,3,opt,name=immutableMetaProperties,proto3" json:"immutableMetaProperties,omitempty" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string             `protobuf:"bytes,4,opt,name=immutableProperties,proto3" json:"immutableProperties,omitempty" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaProperties   string             `protobuf:"bytes,5,opt,name=mutableMetaProperties,proto3" json:"mutableMetaProperties,omitempty" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties       string             `protobuf:"bytes,6,opt,name=mutableProperties,proto3" json:"mutableProperties,omitempty" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

func (m *transactionRequest) Reset()         { *m = transactionRequest{} }
func (m *transactionRequest) String() string { return proto.CompactTextString(m) }
func (*transactionRequest) ProtoMessage()    {}
func (*transactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f778ce26a88af345, []int{0}
}
func (m *transactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *transactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *transactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionRequest.Merge(m, src)
}
func (m *transactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *transactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionRequest proto.InternalMessageInfo


func (m *transactionRequest) GetFromID() string {
	if m != nil {
		return m.FromID
	}
	return ""
}

func (m *transactionRequest) GetImmutableMetaProperties() string {
	if m != nil {
		return m.ImmutableMetaProperties
	}
	return ""
}

func (m *transactionRequest) GetImmutableProperties() string {
	if m != nil {
		return m.ImmutableProperties
	}
	return ""
}

func (m *transactionRequest) GetMutableMetaProperties() string {
	if m != nil {
		return m.MutableMetaProperties
	}
	return ""
}

func (m *transactionRequest) GetMutableProperties() string {
	if m != nil {
		return m.MutableProperties
	}
	return ""
}

func init() {
	proto.RegisterType((*transactionRequest)(nil), "base.transactionRequest")
}

func init() {
	proto.RegisterFile("proto/assets/transactions/define/request.proto", fileDescriptor_f778ce26a88af345)
}

var fileDescriptor_f778ce26a88af345 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x5b, 0x2e, 0x17, 0xee, 0x1d, 0x57, 0x0e, 0x31, 0x10, 0x34, 0x85, 0xb8, 0x50, 0x12,
	0xb4, 0x23, 0xea, 0xc2, 0x98, 0xb8, 0x21, 0xb8, 0x70, 0x61, 0x62, 0x1a, 0x57, 0x6e, 0x9a, 0x29,
	0x1c, 0x70, 0x12, 0xda, 0x69, 0xe7, 0x0c, 0x26, 0xbe, 0x85, 0x8f, 0xe0, 0x43, 0xf8, 0x10, 0x2c,
	0x89, 0x2b, 0xdd, 0x18, 0x03, 0x1b, 0x1f, 0xc3, 0xb4, 0x45, 0x6d, 0x42, 0xdd, 0xf5, 0x9f, 0xff,
	0xfb, 0xe7, 0x3f, 0xe9, 0x1c, 0x62, 0x87, 0x4a, 0x6a, 0xc9, 0x38, 0x22, 0x68, 0x64, 0x5a, 0xf1,
	0x00, 0x79, 0x5f, 0x0b, 0x19, 0x20, 0x1b, 0xc0, 0x50, 0x04, 0xc0, 0x14, 0x44, 0x13, 0x40, 0x9d,
	0x82, 0xb4, 0xe8, 0x71, 0x84, 0xfa, 0x8e, 0xbe, 0x15, 0x6a, 0xe0, 0x86, 0x5c, 0xe9, 0x7b, 0x96,
	0xde, 0x30, 0x92, 0x23, 0xf9, 0xf3, 0x95, 0xd2, 0xf5, 0xf6, 0x2a, 0xd7, 0x97, 0xe8, 0x4b, 0x74,
	0xb3, 0x62, 0x09, 0xef, 0x66, 0xcf, 0x58, 0xdc, 0xc3, 0xee, 0x3a, 0x1e, 0x68, 0xde, 0x49, 0x84,
	0x82, 0x28, 0x05, 0xb7, 0x5f, 0x0b, 0x84, 0x66, 0x26, 0x75, 0xd2, 0x01, 0xe9, 0x19, 0xf9, 0x17,
	0x73, 0xae, 0x82, 0xa8, 0x66, 0x36, 0xcd, 0xd6, 0xda, 0xe1, 0x96, 0xbd, 0x2c, 0x88, 0xcf, 0xed,
	0xe5, 0x65, 0x76, 0x97, 0x23, 0x38, 0x10, 0x75, 0x8b, 0xd3, 0xb7, 0x86, 0xe1, 0x94, 0xbd, 0x54,
	0xd2, 0x2a, 0x29, 0x0f, 0x95, 0xf4, 0x5d, 0xd1, 0xab, 0x15, 0x9a, 0x66, 0xeb, 0xbf, 0x53, 0x8a,
	0xe5, 0x45, 0x8f, 0x9e, 0x90, 0xaa, 0xf0, 0xfd, 0x89, 0xe6, 0xde, 0x18, 0x2e, 0x41, 0xf3, 0x2b,
	0x25, 0x43, 0x50, 0x5a, 0x00, 0xd6, 0xfe, 0x24, 0xe0, 0x6f, 0x36, 0x3d, 0x20, 0x95, 0x6f, 0x2b,
	0x93, 0x2a, 0x26, 0xa9, 0x3c, 0x8b, 0x1e, 0x93, 0x8d, 0xfc, 0xa6, 0xbf, 0x49, 0x26, 0xdf, 0xa4,
	0x7b, 0x64, 0x7d, 0xb5, 0xa5, 0x94, 0x24, 0x56, 0x8d, 0xd3, 0xcd, 0x8f, 0xc7, 0x86, 0xf1, 0xfc,
	0xb4, 0x5f, 0xb9, 0xce, 0xfe, 0x43, 0x0c, 0x65, 0x80, 0xd0, 0x3d, 0x9f, 0xce, 0x2d, 0x73, 0x36,
	0xb7, 0xcc, 0xf7, 0xb9, 0x65, 0x3e, 0x2c, 0x2c, 0x63, 0xb6, 0xb0, 0x8c, 0x97, 0x85, 0x65, 0xdc,
	0xb4, 0x7d, 0x39, 0x98, 0x8c, 0x01, 0xbf, 0x76, 0x45, 0x04, 0x1a, 0x54, 0xc0, 0xc7, 0x79, 0x4b,
	0xe3, 0x95, 0x92, 0x97, 0x3a, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x25, 0xb8, 0x66, 0x5f,
	0x02, 0x00, 0x00,
}

func (m *transactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *transactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *transactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MutableProperties) > 0 {
		i -= len(m.MutableProperties)
		copy(dAtA[i:], m.MutableProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MutableProperties)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MutableMetaProperties) > 0 {
		i -= len(m.MutableMetaProperties)
		copy(dAtA[i:], m.MutableMetaProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.MutableMetaProperties)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ImmutableProperties) > 0 {
		i -= len(m.ImmutableProperties)
		copy(dAtA[i:], m.ImmutableProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ImmutableProperties)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ImmutableMetaProperties) > 0 {
		i -= len(m.ImmutableMetaProperties)
		copy(dAtA[i:], m.ImmutableMetaProperties)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ImmutableMetaProperties)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FromID) > 0 {
		i -= len(m.FromID)
		copy(dAtA[i:], m.FromID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.FromID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.BaseReq.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRequest(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *transactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseReq.Size()
	n += 1 + l + sovRequest(uint64(l))
	l = len(m.FromID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.ImmutableMetaProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.ImmutableProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.MutableMetaProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.MutableProperties)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *transactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: transactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: transactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseReq", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseReq.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImmutableMetaProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImmutableProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MutableMetaProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MutableProperties = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)