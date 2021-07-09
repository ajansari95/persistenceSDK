// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/splits/trnasactions/send/v1beta1/tx.proto

package send

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_test_types "github.com/persistenceOne/persistenceSDK/schema/test_types"
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

type Message struct {
	From      github_com_cosmos_cosmos_sdk_types.AccAddress                                                      `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	FromID    github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,2,opt,name=from_iD,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"from_iD"`
	ToID      github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,3,opt,name=to_iD,json=toID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"to_iD"`
	OwnableID github_com_persistenceOne_persistenceSDK_schema_test_types.ID `protobuf:"bytes,4,opt,name=ownable_iD,json=ownableID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/test_types.ID" json:"ownable_iD"`
	Value     github_com_cosmos_cosmos_sdk_types.Dec                        `protobuf:"bytes,5,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"value"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8382cab0e70961f8, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "base.message")
}

func init() {
	proto.RegisterFile("proto/splits/trnasactions/send/v1beta1/tx.proto", fileDescriptor_8382cab0e70961f8)
}

var fileDescriptor_8382cab0e70961f8 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x6e, 0xe2, 0x40,
	0x10, 0x86, 0xed, 0xc3, 0xc0, 0xb1, 0xa5, 0x2b, 0x74, 0x85, 0x7d, 0xba, 0x02, 0xdd, 0x15, 0xe7,
	0x15, 0xba, 0xee, 0xa4, 0x6b, 0xb8, 0x4d, 0x61, 0xa5, 0x88, 0x44, 0x3a, 0x8a, 0xa0, 0xb5, 0x3d,
	0x31, 0x56, 0xec, 0x5d, 0x6b, 0x67, 0x20, 0xe1, 0x0d, 0xd2, 0x25, 0x8f, 0xc0, 0xe3, 0x50, 0x52,
	0x46, 0x29, 0x50, 0x04, 0x4d, 0x1e, 0x23, 0xb2, 0x71, 0x14, 0xd2, 0x53, 0xcd, 0xec, 0xea, 0xdf,
	0xef, 0x2b, 0xf6, 0x67, 0xbc, 0x34, 0x9a, 0x34, 0xc7, 0x32, 0xcf, 0x08, 0x39, 0x19, 0x25, 0x51,
	0xc6, 0x94, 0x69, 0x85, 0x1c, 0x41, 0x25, 0x7c, 0x31, 0x8c, 0x80, 0xe4, 0x90, 0xd3, 0x5d, 0x50,
	0x27, 0x5d, 0x27, 0x92, 0x08, 0xdf, 0x06, 0x34, 0xcb, 0x4c, 0x32, 0x2d, 0xa5, 0xa1, 0x65, 0x83,
	0x48, 0x75, 0xaa, 0x3f, 0xb6, 0x43, 0xfa, 0xc7, 0x43, 0x8b, 0x75, 0x0b, 0x40, 0x94, 0x29, 0xb8,
	0x2e, 0x73, 0xae, 0x8d, 0x2e, 0xfa, 0xf6, 0x77, 0xfb, 0x67, 0x6f, 0x5c, 0xef, 0xee, 0x15, 0xeb,
	0x56, 0x73, 0x9a, 0x89, 0xfe, 0x97, 0xea, 0x7a, 0x74, 0xb6, 0xde, 0xfa, 0xd6, 0xf3, 0xd6, 0xff,
	0x97, 0x66, 0x34, 0x9b, 0x47, 0x41, 0xac, 0x0b, 0x5e, 0x82, 0xc1, 0x0c, 0x09, 0x54, 0x0c, 0x17,
	0x0a, 0x8e, 0x8f, 0x97, 0xe2, 0x9c, 0x63, 0x3c, 0x83, 0x42, 0x72, 0x02, 0xa4, 0x29, 0x2d, 0x4b,
	0xc0, 0x20, 0x14, 0xe3, 0x4e, 0x45, 0x0d, 0x85, 0x3b, 0x61, 0x6d, 0xd2, 0x15, 0xbd, 0x75, 0x4a,
	0xba, 0x43, 0x3a, 0x14, 0x6e, 0xc2, 0x98, 0xbe, 0x55, 0x32, 0xca, 0xa1, 0x12, 0x38, 0xa7, 0x14,
	0xf4, 0x1a, 0x70, 0x28, 0x5c, 0xc1, 0xda, 0x0b, 0x99, 0xcf, 0xa1, 0xdf, 0xae, 0x05, 0x41, 0x23,
	0x18, 0x1c, 0x09, 0x62, 0x8d, 0x85, 0xc6, 0x66, 0xfc, 0xc6, 0xe4, 0x86, 0x1f, 0x40, 0x02, 0xe2,
	0xf1, 0xe1, 0xf1, 0xdf, 0xaf, 0xf7, 0x2b, 0xdf, 0x7a, 0x5d, 0xf9, 0xd6, 0xe8, 0xff, 0x7a, 0xe7,
	0xd9, 0x9b, 0x9d, 0x67, 0xbf, 0xec, 0x3c, 0xfb, 0x71, 0xef, 0x59, 0x9b, 0xbd, 0x67, 0x3d, 0xed,
	0x3d, 0x6b, 0xf2, 0xab, 0xd0, 0xc9, 0x3c, 0x07, 0x7c, 0x2f, 0x43, 0xa6, 0x08, 0x8c, 0x92, 0x39,
	0x27, 0x23, 0xd5, 0xa7, 0x56, 0x44, 0x9d, 0xfa, 0x77, 0xff, 0xbc, 0x05, 0x00, 0x00, 0xff, 0xff,
	0xb4, 0xf8, 0x53, 0xba, 0x3e, 0x02, 0x00, 0x00,
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Value.Size()
		i -= size
		if _, err := m.Value.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.OwnableID.Size()
		i -= size
		if _, err := m.OwnableID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ToID.Size()
		i -= size
		if _, err := m.ToID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.FromID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.ToID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.OwnableID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Value.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = dAtA[iNdEx:postIndex]
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnableID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OwnableID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
