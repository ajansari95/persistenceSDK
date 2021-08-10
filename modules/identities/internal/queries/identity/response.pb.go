// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/identities/queries/v1beta1/response.proto

package identity

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	helpers "github.com/persistenceOne/persistenceSDK/schema/helpers"
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

type queryResponse struct {
	Success bool               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   protoTypes.Error   `protobuf:"bytes,2,opt,name=error,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/proto/types.Error" json:"error"`
	List    []helpers.Mappable `protobuf:"bytes,3,rep,name=list,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/helpers.Mappable" json:"list"`
}

func (m queryResponse) Reset()         { m = queryResponse{} }
func (m queryResponse) String() string { return proto.CompactTextString(&m) }
func (queryResponse) ProtoMessage()    {}
func (*queryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06901d5d4cf72311, []int{0}
}
func (m *queryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *queryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *queryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *queryResponse) XXX_Size() int {
	return m.Size()
}
func (m *queryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *queryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*queryResponse)(nil), "identities.queries.v1beta1.queryResponse")
}

func init() {
	proto.RegisterFile("proto/identities/queries/v1beta1/response.proto", fileDescriptor_06901d5d4cf72311)
}

var fileDescriptor_06901d5d4cf72311 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xb3, 0xd6, 0x7f, 0x0d, 0xf4, 0x12, 0x10, 0x42, 0x0f, 0x69, 0xf1, 0x20, 0x05, 0x35,
	0xa1, 0x78, 0xf3, 0xa4, 0x45, 0x45, 0x11, 0x11, 0xe3, 0x4d, 0x84, 0x92, 0x6c, 0x87, 0x66, 0x21,
	0xd9, 0x5d, 0x77, 0x36, 0x42, 0xde, 0xc2, 0x47, 0xf0, 0x21, 0x7c, 0x88, 0x1e, 0x8b, 0x27, 0xf1,
	0x50, 0xa4, 0xbd, 0xf8, 0x16, 0x4a, 0xda, 0xd5, 0x06, 0x3c, 0x79, 0xda, 0xfd, 0x31, 0x1f, 0xdf,
	0xcc, 0x30, 0x76, 0x20, 0x95, 0xd0, 0x22, 0x60, 0x03, 0xe0, 0x9a, 0x69, 0x06, 0x18, 0x3c, 0xe4,
	0xa0, 0xca, 0xf7, 0xb1, 0x1b, 0x83, 0x8e, 0xba, 0x81, 0x02, 0x94, 0x82, 0x23, 0xf8, 0x73, 0xd2,
	0x69, 0x2e, 0x51, 0xdf, 0xa0, 0xbe, 0x41, 0x9b, 0x3b, 0x3a, 0x61, 0x6a, 0xd0, 0x97, 0x91, 0xd2,
	0x85, 0x11, 0x0f, 0xc5, 0x50, 0x2c, 0x7f, 0x0b, 0x47, 0x73, 0xf7, 0x2f, 0x47, 0x05, 0x66, 0x02,
	0xfb, 0xd5, 0xb0, 0x80, 0xb7, 0xbf, 0x88, 0xdd, 0x28, 0x1b, 0x15, 0xa1, 0x19, 0xc4, 0x71, 0xed,
	0x0d, 0xcc, 0x29, 0x05, 0x44, 0x97, 0xb4, 0x49, 0x67, 0x33, 0xfc, 0x89, 0x4e, 0xdf, 0x5e, 0x03,
	0xa5, 0x84, 0x72, 0x57, 0xda, 0xa4, 0x53, 0xef, 0x5d, 0x8c, 0x26, 0x2d, 0xeb, 0x7d, 0xd2, 0x3a,
	0x1e, 0x32, 0x9d, 0xe4, 0xb1, 0x4f, 0x45, 0x16, 0x48, 0x50, 0xc8, 0x50, 0x03, 0xa7, 0x70, 0xcd,
	0xa1, 0x1a, 0x6f, 0x4f, 0x2e, 0x03, 0xa4, 0x09, 0x64, 0x91, 0x99, 0x49, 0x17, 0x12, 0xd0, 0x3f,
	0x2d, 0x85, 0xe1, 0xc2, 0xeb, 0xdc, 0xdb, 0xab, 0x29, 0x43, 0xed, 0xd6, 0xda, 0xb5, 0x4e, 0xbd,
	0x77, 0x6e, 0xfc, 0x47, 0xff, 0xf5, 0x27, 0x90, 0x96, 0x05, 0xff, 0x2a, 0x92, 0x32, 0x8a, 0x53,
	0x08, 0xe7, 0xd6, 0xc3, 0xad, 0xcf, 0xe7, 0x96, 0xf5, 0xfa, 0xb2, 0xdf, 0xb8, 0xa9, 0xee, 0xdb,
	0x3b, 0x1b, 0x4d, 0x3d, 0x32, 0x9e, 0x7a, 0xe4, 0x63, 0xea, 0x91, 0xa7, 0x99, 0x67, 0x8d, 0x67,
	0x9e, 0xf5, 0x36, 0xf3, 0xac, 0xbb, 0xbd, 0x4c, 0x0c, 0xf2, 0x14, 0xb0, 0x7a, 0x3f, 0xc6, 0x35,
	0x28, 0x1e, 0xa5, 0xbf, 0x87, 0x34, 0xb5, 0x22, 0x5e, 0x9f, 0xaf, 0x75, 0xf0, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0x0e, 0xaf, 0x12, 0xf4, 0x01, 0x00, 0x00,
}

func (m *queryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *queryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *queryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.List[iNdEx].Size()
				i -= size
				if _, err := m.List[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintResponse(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.Error.Size()
		i -= size
		if _, err := m.Error.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintResponse(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *queryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = m.Error.Size()
	n += 1 + l + sovResponse(uint64(l))
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovResponse(uint64(l))
		}
	}
	return n
}

func sovResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResponse(x uint64) (n int) {
	return sovResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *queryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResponse
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
			return fmt.Errorf("proto: queryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: queryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
			m.Success = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResponse
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
				return ErrInvalidLengthResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v helpers.Mappable
			m.List = append(m.List, v)
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResponse
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
func skipResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
					return 0, ErrIntOverflowResponse
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
				return 0, ErrInvalidLengthResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResponse = fmt.Errorf("proto: unexpected end of group")
)
