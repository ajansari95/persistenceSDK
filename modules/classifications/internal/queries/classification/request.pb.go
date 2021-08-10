// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/classifications/queries/v1beta1/request.proto

package classification

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/proto/types"
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

type queryRequest struct {
	ClassificationID types.ID `protobuf:"bytes,1,opt,name=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classificationID"valid:"required~required field classificationID missing"`
}

func (m *queryRequest) Reset()         { *m = queryRequest{} }
func (m *queryRequest) String() string { return proto.CompactTextString(m) }
func (*queryRequest) ProtoMessage()    {}
func (*queryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af9f6a9ff4e7b68f, []int{0}
}
func (m *queryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *queryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *queryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *queryRequest) XXX_Size() int {
	return m.Size()
}
func (m *queryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*queryRequest)(nil), "base.queryRequest")
}

func init() {
	proto.RegisterFile("proto/classifications/queries/v1beta1/request.proto", fileDescriptor_af9f6a9ff4e7b68f)
}

var fileDescriptor_af9f6a9ff4e7b68f = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0x49, 0x2c, 0x2e, 0xce, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf,
	0x2b, 0xd6, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x03, 0xab, 0x16, 0x62, 0x49, 0x4a,
	0x2c, 0x4e, 0x95, 0x52, 0x2b, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0xd4,
	0x87, 0x18, 0x93, 0x9e, 0x9f, 0x9e, 0x8f, 0x60, 0x41, 0x54, 0x4b, 0x69, 0x63, 0xaa, 0x4b, 0xce,
	0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x47, 0xe6, 0x40, 0x14, 0x2b, 0x4d, 0x67, 0xe4, 0xe2, 0x01, 0x59,
	0x5e, 0x19, 0x04, 0xb1, 0x51, 0xa8, 0x90, 0x4b, 0x00, 0xd5, 0x71, 0x9e, 0x2e, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x4e, 0xae, 0x27, 0xee, 0xc9, 0x33, 0xdc, 0xba, 0x27, 0x6f, 0x9b, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x90, 0x5a, 0x54, 0x9c, 0x59, 0x5c, 0x92,
	0x9a, 0x97, 0x9c, 0xea, 0x9f, 0x97, 0x8a, 0xcc, 0x0d, 0x76, 0xf1, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0xf3, 0x74,
	0x09, 0xc2, 0x30, 0xde, 0x4a, 0xe4, 0xc5, 0x02, 0x79, 0x86, 0x4b, 0x5b, 0x74, 0x79, 0x02, 0x91,
	0x1c, 0xe2, 0x14, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xe6, 0xb9,
	0xf9, 0x29, 0xa5, 0x39, 0xa9, 0xc5, 0x18, 0xa1, 0x98, 0x99, 0x57, 0x92, 0x5a, 0x94, 0x97, 0x98,
	0x03, 0x0f, 0x4e, 0x54, 0x05, 0x49, 0x6c, 0x60, 0x3f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x52, 0x57, 0xe6, 0xf3, 0x85, 0x01, 0x00, 0x00,
}

func (m *queryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *queryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *queryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClassificationID.Size()
		i -= size
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
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
func (m *queryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ClassificationID.Size()
	n += 1 + l + sovRequest(uint64(l))
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *queryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: queryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: queryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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