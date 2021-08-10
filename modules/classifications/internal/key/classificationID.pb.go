// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/classifications/key/v1beta1/classificationID.proto

package key

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
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

type classificationID struct {
	ChainID protoTypes.ID `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID" valid:"required~required field chainID missing"`
	HashID  protoTypes.ID `protobuf:"bytes,2,opt,name=hashID,proto3" json:"hashID" valid:"required~required field hashID missing"`
}

func (m *classificationID) Reset()         { *m = classificationID{} }


func (*classificationID) ProtoMessage()    {}
func (*classificationID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8743d8ff5ca95ad, []int{0}
}
func (m *classificationID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *classificationID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClassificationID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *classificationID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationID.Merge(m, src)
}
func (m *classificationID) XXX_Size() int {
	return m.Size()
}
func (m *classificationID) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationID.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationID proto.InternalMessageInfo

func init() {
	proto.RegisterType((*classificationID)(nil), "classifications.key.v1beta1.classificationID")
}

func init() {
	proto.RegisterFile("proto/classifications/key/v1beta1/classificationID.proto", fileDescriptor_e8743d8ff5ca95ad)
}

var fileDescriptor_e8743d8ff5ca95ad = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0xd6, 0xcf, 0x4e, 0xad, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0x44, 0x93, 0xf3,
	0x74, 0xd1, 0x03, 0x6b, 0x11, 0x92, 0x46, 0xd3, 0xa3, 0x97, 0x9d, 0x5a, 0xa9, 0x07, 0xd5, 0x23,
	0xa5, 0x56, 0x92, 0x91, 0x59, 0x94, 0x12, 0x5f, 0x90, 0x58, 0x54, 0x52, 0xa9, 0x0f, 0xb1, 0x22,
	0x3d, 0x3f, 0x3d, 0x1f, 0xc1, 0x82, 0x18, 0x22, 0xa5, 0x04, 0x11, 0x29, 0x4e, 0xce, 0x48, 0xcd,
	0x4d, 0xd4, 0x4f, 0x4a, 0x2c, 0x4e, 0x85, 0x5b, 0x9b, 0x99, 0x02, 0x51, 0xa3, 0xd4, 0xc3, 0xc8,
	0x25, 0x80, 0xee, 0x06, 0x21, 0x73, 0x2e, 0xf6, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x4f, 0x17, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x71, 0x3d, 0x88, 0x21, 0x7a, 0x20, 0x43, 0x60, 0xee, 0xd0,
	0xf3, 0x74, 0x71, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xa6, 0x5a, 0xc8, 0x94, 0x8b, 0x2d,
	0x23, 0xb1, 0x38, 0xc3, 0xd3, 0x45, 0x82, 0x89, 0x18, 0x7d, 0x50, 0xc5, 0x56, 0x1c, 0x1d, 0x0b,
	0xe4, 0x19, 0x5e, 0x2c, 0x90, 0x67, 0x70, 0xb2, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0x95, 0xdc, 0xfc, 0x94, 0xd2, 0x9c, 0xd4, 0x62, 0x8c, 0xd0, 0xcc, 0xcc, 0x2b,
	0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x01, 0x05, 0x6b, 0x12, 0x1b, 0xd8, 0x57, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0xc5, 0xc3, 0x1e, 0x7a, 0x01, 0x00, 0x00,
}

func (m *classificationID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m classificationID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m classificationID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HashID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClassificationID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ChainID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintClassificationID(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintClassificationID(dAtA []byte, offset int, v uint64) int {
	offset -= sovClassificationID(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m classificationID) Size() (n int) {

	var l int
	_ = l
	l = m.ChainID.Size()
	n += 1 + l + sovClassificationID(uint64(l))
	l = m.HashID.Size()
	n += 1 + l + sovClassificationID(uint64(l))
	return n
}

func sovClassificationID(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClassificationID(x uint64) (n int) {
	return sovClassificationID(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m classificationID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClassificationID
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
			return fmt.Errorf("proto: classificationID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: classificationID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassificationID
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
				return ErrInvalidLengthClassificationID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassificationID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClassificationID
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
				return ErrInvalidLengthClassificationID
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClassificationID
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HashID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClassificationID(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClassificationID
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
func skipClassificationID(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClassificationID
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
					return 0, ErrIntOverflowClassificationID
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
					return 0, ErrIntOverflowClassificationID
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
				return 0, ErrInvalidLengthClassificationID
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClassificationID
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClassificationID
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClassificationID        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClassificationID          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClassificationID = fmt.Errorf("proto: unexpected end of group")
)