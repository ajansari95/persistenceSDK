// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/identities/transactions/unprovision/request.proto

package unprovision

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
	BaseReq    protoTypes.BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq,proto3" json:"base_req"`
	To         string             `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"  valid:"required~required field to missing, matches(^[a-z0-9]+$)~invalid field to"`
	IdentityID string             `protobuf:"bytes,3,opt,name=identityID,proto3" json:"identityID,omitempty"  valid:"required~required field identityID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field identityID"`
}

func (m transactionRequest) Reset()         { m = transactionRequest{} }
func (m transactionRequest) String() string { return proto.CompactTextString(&m) }
func (transactionRequest) ProtoMessage()    {}
func (*transactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cea53f85f61c4a5, []int{0}
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



func (m *transactionRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *transactionRequest) GetIdentityID() string {
	if m != nil {
		return m.IdentityID
	}
	return ""
}

func init() {
	proto.RegisterType((*transactionRequest)(nil), "base.transactionRequest")
}

func init() {
	proto.RegisterFile("proto/identities/transactions/unprovision/request.proto", fileDescriptor_4cea53f85f61c4a5)
}

var fileDescriptor_4cea53f85f61c4a5 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x2f, 0x29, 0x4a,
	0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x2b, 0xd6, 0x2f, 0xcd, 0x2b, 0x28, 0xca, 0x2f,
	0xcb, 0x2c, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x03, 0xeb,
	0x10, 0x62, 0x49, 0x4a, 0x2c, 0x4e, 0x95, 0x52, 0x2b, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48,
	0x2c, 0x2a, 0xa9, 0xd4, 0x87, 0x18, 0x95, 0x9e, 0x9f, 0x9e, 0x8f, 0x60, 0x41, 0x54, 0x4b, 0x69,
	0x63, 0xaa, 0x4b, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x47, 0xe6, 0x40, 0x15, 0xab, 0x23, 0x8b,
	0xe9, 0x83, 0xec, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0x04, 0x73, 0x8a, 0x52, 0x0b,
	0x21, 0x0a, 0x95, 0xe6, 0x33, 0x72, 0x09, 0x21, 0x39, 0x37, 0x08, 0xe2, 0x40, 0x21, 0x5b, 0x2e,
	0x0e, 0x90, 0xba, 0xf8, 0xa2, 0xd4, 0x42, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x19, 0x3d,
	0xa8, 0x05, 0x20, 0x71, 0x3d, 0xa8, 0x61, 0x7a, 0x4e, 0x89, 0xc5, 0xa9, 0x41, 0xa9, 0x85, 0x4e,
	0x2c, 0x27, 0xee, 0xc9, 0x33, 0x04, 0xb1, 0x27, 0x41, 0xb8, 0x42, 0x7c, 0x5c, 0x4c, 0x25, 0xf9,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x25, 0xf9, 0x42, 0x72, 0x5c, 0x5c, 0xd0, 0xe0,
	0xa9, 0xf4, 0x74, 0x91, 0x60, 0x06, 0x8b, 0x23, 0x89, 0x58, 0x49, 0xbd, 0x58, 0x20, 0xcf, 0x70,
	0x69, 0x8b, 0xae, 0x50, 0x08, 0x86, 0x53, 0x9c, 0xfc, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1,
	0x58, 0x8e, 0x21, 0xca, 0x24, 0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x18, 0x39, 0xe8, 0x33, 0xf3,
	0x4a, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0x70, 0xc6, 0x41, 0x12, 0x1b, 0xd8, 0xe3, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x81, 0x99, 0x50, 0x90, 0xb7, 0x01, 0x00, 0x00,
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
	if len(m.IdentityID) > 0 {
		i -= len(m.IdentityID)
		copy(dAtA[i:], m.IdentityID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.IdentityID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.To)))
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
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	l = len(m.IdentityID)
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
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
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
			m.IdentityID = string(dAtA[iNdEx:postIndex])
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