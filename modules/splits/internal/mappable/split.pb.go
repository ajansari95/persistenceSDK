// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/splits/internal/mappable/split.proto

package mappable

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type Split struct {
	ID    base.ID                                `protobuf:"bytes,1,opt,name=i_d,json=iD,proto3" json:"i_d"`
	Value github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=value,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"value"`
}

func (m *Split) Reset()         { *m = Split{} }
func (m *Split) String() string { return proto.CompactTextString(m) }
func (*Split) ProtoMessage()    {}
func (*Split) Descriptor() ([]byte, []int) {
	return fileDescriptor_03f63749cdee6543, []int{0}
}
func (m *Split) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Split) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Split.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Split) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Split.Merge(m, src)
}
func (m *Split) XXX_Size() int {
	return m.Size()
}
func (m *Split) XXX_DiscardUnknown() {
	xxx_messageInfo_Split.DiscardUnknown(m)
}

var xxx_messageInfo_Split proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Split)(nil), "persistence_sdk.modules.splits.internal.mappable.Split")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/splits/internal/mappable/split.proto", fileDescriptor_03f63749cdee6543)
}

var fileDescriptor_03f63749cdee6543 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0xcf, 0xcd, 0x4f, 0x29,
	0xcd, 0x49, 0x2d, 0xd6, 0x2f, 0x2e, 0xc8, 0xc9, 0x2c, 0x29, 0xd6, 0xcf, 0xcc, 0x2b, 0x49, 0x2d,
	0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0x4d, 0x2c, 0x28, 0x48, 0x4c, 0xca, 0x49, 0x85, 0x48, 0xe8, 0x15,
	0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x19, 0xa0, 0xe9, 0xd6, 0x83, 0xea, 0xd6, 0x83, 0xe8, 0xd6, 0x83,
	0xe9, 0xd6, 0x83, 0xe9, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x6b, 0xd6, 0x07, 0xb1, 0x20,
	0xe6, 0x48, 0x69, 0xa1, 0xbb, 0xa2, 0x38, 0x39, 0x23, 0x35, 0x37, 0x51, 0xbf, 0xa4, 0xb2, 0x20,
	0xb5, 0x58, 0x3f, 0x29, 0xb1, 0x38, 0x55, 0x3f, 0x33, 0x05, 0xa2, 0x56, 0x69, 0x22, 0x23, 0x17,
	0x6b, 0x30, 0xc8, 0x78, 0x21, 0x1b, 0x2e, 0xe6, 0xcc, 0xf8, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x55, 0x3d, 0x74, 0xb7, 0x40, 0xcc, 0xd0, 0x03, 0x9b, 0xa1, 0x07, 0x32, 0x43, 0xcf,
	0xd3, 0xc5, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0xa6, 0x4c, 0x17, 0x21, 0x17, 0x2e, 0xd6,
	0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0x3d, 0x90, 0xc4, 0xad,
	0x7b, 0xf2, 0x6a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc9, 0xf9,
	0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0x17, 0xe4, 0x2e, 0x88, 0x61, 0x2e, 0xa9, 0xc9, 0x41, 0x10,
	0xcd, 0x56, 0x2c, 0x1d, 0x0b, 0xe4, 0x19, 0x9c, 0x92, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1,
	0x58, 0x8e, 0x21, 0xca, 0x13, 0xc9, 0x38, 0x24, 0x07, 0xfa, 0xe7, 0xa5, 0x22, 0x73, 0x83, 0x5d,
	0xbc, 0x09, 0x06, 0x7c, 0x12, 0x1b, 0xd8, 0xff, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45,
	0x20, 0xa1, 0xc9, 0xb3, 0x01, 0x00, 0x00,
}

func (m *Split) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Split) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Split) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintSplit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintSplit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintSplit(dAtA []byte, offset int, v uint64) int {
	offset -= sovSplit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Split) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovSplit(uint64(l))
	l = m.Value.Size()
	n += 1 + l + sovSplit(uint64(l))
	return n
}

func sovSplit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSplit(x uint64) (n int) {
	return sovSplit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Split) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSplit
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
			return fmt.Errorf("proto: Split: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Split: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplit
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
				return ErrInvalidLengthSplit
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSplit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSplit
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
				return ErrInvalidLengthSplit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSplit
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
			skippy, err := skipSplit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSplit
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
func skipSplit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSplit
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
					return 0, ErrIntOverflowSplit
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
					return 0, ErrIntOverflowSplit
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
				return 0, ErrInvalidLengthSplit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSplit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSplit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSplit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSplit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSplit = fmt.Errorf("proto: unexpected end of group")
)