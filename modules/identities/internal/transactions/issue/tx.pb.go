// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/identities/transactions/issue/tx.proto

package issue

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	"github.com/persistenceOne/persistenceSDK/schema/proto/types"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type message struct {
	From                    github_com_cosmos_cosmos_sdk_types.AccAddress                        `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"from"  valid:"required~required field from missing"`
	To                      github_com_cosmos_cosmos_sdk_types.AccAddress                        `protobuf:"bytes,2,opt,name=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"fromID"  valid:"required~required field to missing, matches(^[a-z0-9]*$)~invalid field to"`
	FromID                  types.ID                                                             `protobuf:"bytes,3,opt,name=toID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"toID"  valid:"required~required field fromID missing"`
	ClassificationID        types.ID                                                             `protobuf:"bytes,4,opt,name=classificationID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"classificationID"  valid:"required~required field classificationID missing"`
	ImmutableMetaProperties github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties `protobuf:"bytes,5,opt,name=immutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"immutableMetaProperties"  valid:"required~required field immutableMetaProperties missing"`
	ImmutableProperties     github_com_persistenceOne_persistenceSDK_schema_types.Properties     `protobuf:"bytes,6,opt,name=immutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"immutableProperties"  valid:"required~required field immutableProperties missing"`
	MutableMetaProperties   github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties `protobuf:"bytes,7,opt,name=mutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"mutableMetaProperties"  valid:"required~required field mutableMetaProperties missing"`
	MutableProperties       github_com_persistenceOne_persistenceSDK_schema_types.Properties     `protobuf:"bytes,8,opt,name=mutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"mutableProperties"  valid:"required~required field mutableProperties missing"`
}

func (m message) Reset()         { m = message{} }
func (m message) String() string { return proto.CompactTextString(&m) }
func (message) ProtoMessage()    {}
func (*message) Descriptor() ([]byte, []int) {
	return fileDescriptor_181cf5541be00294, []int{0}
}
func (m *message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *message) XXX_Size() int {
	return m.Size()
}
func (m *message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func init() {
	proto.RegisterType((*message)(nil), "base.message")
}

func init() {
	proto.RegisterFile("proto/identities/transactions/issue/tx.proto", fileDescriptor_181cf5541be00294)
}

var fileDescriptor_181cf5541be00294 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xc0, 0x13, 0xbb, 0xee, 0xd6, 0xb9, 0xa8, 0x23, 0x62, 0xdc, 0x43, 0x56, 0x3c, 0x88, 0x07,
	0x9b, 0x60, 0xc5, 0x8b, 0x20, 0xd8, 0x12, 0xa5, 0x41, 0x17, 0x65, 0x3d, 0x29, 0x48, 0x99, 0x9d,
	0xbc, 0x66, 0x07, 0x37, 0x99, 0x74, 0xde, 0x8b, 0x58, 0x8f, 0x1e, 0xc4, 0x8b, 0xe0, 0x47, 0xe8,
	0xc7, 0xe9, 0xb1, 0x47, 0xf1, 0x50, 0x64, 0xf7, 0xe2, 0xc7, 0x90, 0xfc, 0xc1, 0x46, 0x36, 0xc2,
	0x2a, 0xdb, 0xd3, 0x24, 0xcc, 0xbc, 0xdf, 0xef, 0xcd, 0x1b, 0xde, 0x63, 0x77, 0x32, 0xa3, 0x49,
	0xfb, 0x2a, 0x82, 0x94, 0x14, 0x29, 0x40, 0x9f, 0x8c, 0x48, 0x51, 0x48, 0x52, 0x3a, 0x45, 0x5f,
	0x21, 0xe6, 0xe0, 0xd3, 0x7b, 0xaf, 0x3c, 0xc6, 0x3b, 0x63, 0x81, 0xd0, 0xbf, 0x45, 0x13, 0x65,
	0xa2, 0xdd, 0x4c, 0x18, 0x3a, 0xf0, 0xab, 0xf8, 0x58, 0xc7, 0xfa, 0xf4, 0xab, 0x3a, 0xdd, 0xbf,
	0xbb, 0x0c, 0xdb, 0xc0, 0x7e, 0x0e, 0x48, 0x75, 0xc8, 0xe6, 0x72, 0x21, 0x98, 0xe9, 0x14, 0xa1,
	0x8a, 0xb9, 0xf9, 0xa5, 0xc7, 0x7a, 0x09, 0x20, 0x8a, 0x18, 0x78, 0xc8, 0x3a, 0x7b, 0x46, 0x27,
	0x8e, 0x7d, 0xc3, 0xbe, 0x7d, 0x61, 0xfb, 0xfe, 0xd1, 0xc9, 0xc0, 0xfa, 0x7e, 0x32, 0xd8, 0x88,
	0x15, 0x4d, 0xf2, 0xb1, 0x27, 0x75, 0xe2, 0x4b, 0x8d, 0x89, 0xc6, 0x7a, 0xd9, 0xc0, 0xe8, 0xad,
	0x4f, 0x07, 0x19, 0xa0, 0xb7, 0x25, 0xe5, 0x56, 0x14, 0x19, 0x40, 0x1c, 0x95, 0x08, 0xfe, 0x86,
	0x75, 0x8b, 0x35, 0x0c, 0x9c, 0x73, 0x25, 0xec, 0x71, 0x0d, 0x7b, 0xd8, 0x80, 0x65, 0x60, 0x50,
	0x21, 0x41, 0x2a, 0xe1, 0x79, 0x0a, 0xcd, 0xdf, 0x97, 0xc1, 0x53, 0x1f, 0xe5, 0x04, 0x12, 0xe1,
	0x13, 0x20, 0xed, 0x56, 0x92, 0x30, 0x18, 0xd5, 0x50, 0xfe, 0x8a, 0x75, 0x48, 0x87, 0x81, 0xb3,
	0xb6, 0x4a, 0x78, 0x89, 0xe4, 0xfb, 0xec, 0x92, 0x9c, 0x0a, 0x44, 0xb5, 0xa7, 0xa4, 0x28, 0x0a,
	0x17, 0x06, 0x4e, 0x67, 0x95, 0x9a, 0x05, 0x3c, 0xff, 0x64, 0xb3, 0x6b, 0x2a, 0x49, 0x72, 0x12,
	0xe3, 0x29, 0x0c, 0x81, 0xc4, 0x0b, 0xa3, 0x33, 0x30, 0xc5, 0x0b, 0x3a, 0xe7, 0x4b, 0xf5, 0xb3,
	0x5a, 0x1d, 0xfc, 0xb3, 0xba, 0xb4, 0xfe, 0xc9, 0x1c, 0xfd, 0x4d, 0xc6, 0x3f, 0xb0, 0x2b, 0xbf,
	0xb7, 0x1a, 0x39, 0x74, 0xcb, 0x1c, 0x76, 0xea, 0x1c, 0x1e, 0xfd, 0x5f, 0x0e, 0x0d, 0x7f, 0x9b,
	0x84, 0x7f, 0xb4, 0xd9, 0xd5, 0xf6, 0x12, 0xf4, 0xce, 0xa0, 0x04, 0xed, 0x2a, 0xfe, 0x8e, 0x5d,
	0x5e, 0xbc, 0xfe, 0xfa, 0x8a, 0xaf, 0xbf, 0xa8, 0x78, 0xb0, 0xfe, 0xf9, 0x70, 0x60, 0xfd, 0x3c,
	0x1c, 0x58, 0x9b, 0x43, 0xb6, 0x36, 0xc4, 0x98, 0x3f, 0x61, 0x17, 0x4f, 0xdb, 0x38, 0x2c, 0x1a,
	0x97, 0x3b, 0x5e, 0x31, 0x3f, 0xbc, 0x46, 0x4b, 0x8f, 0xaa, 0xee, 0xef, 0x5f, 0x6f, 0xd9, 0xa9,
	0x9a, 0x7c, 0x7b, 0xe7, 0x68, 0xe6, 0xda, 0xc7, 0x33, 0xd7, 0xfe, 0x31, 0x73, 0xed, 0xaf, 0x73,
	0xd7, 0x3a, 0x9e, 0xbb, 0xd6, 0xb7, 0xb9, 0x6b, 0xbd, 0xf6, 0x12, 0x1d, 0xe5, 0x53, 0xc0, 0xe6,
	0xb8, 0x50, 0x29, 0x81, 0x49, 0xc5, 0xb4, 0x65, 0x6e, 0x8c, 0xbb, 0xe5, 0xbc, 0xb8, 0xf7, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x7c, 0x97, 0x45, 0xf4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	IdentitiesIssue(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IdentitiesIssue(ctx context.Context, in *transactionRequest, opts ...grpc.CallOption) (*transactionResponse, error) {
	out := new(transactionResponse)
	err := c.cc.Invoke(ctx, "/base.Msg/identitiesIssue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	IdentitiesIssue(context.Context, *transactionRequest) (*transactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) IdentitiesIssue(ctx context.Context, req *transactionRequest) (*transactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IdentitiesIssue not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_IdentitiesIssue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(transactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IdentitiesIssue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.Msg/IdentitiesIssue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IdentitiesIssue(ctx, req.(*transactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "identitiesIssue",
			Handler:    _Msg_IdentitiesIssue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/identities/transactions/issue/tx.proto",
}

func (m *message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MutableProperties.Size()
		i -= size
		if _, err := m.MutableProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.MutableMetaProperties.Size()
		i -= size
		if _, err := m.MutableMetaProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.ImmutableProperties.Size()
		i -= size
		if _, err := m.ImmutableProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ImmutableMetaProperties.Size()
		i -= size
		if _, err := m.ImmutableMetaProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.ClassificationID.Size()
		i -= size
		if _, err := m.ClassificationID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := len(m.To)
		i -= size
		if _, err := m.Error.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := len(m.From)
		i -= size
		if _, err := m.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	n += 1 + l + sovTx(uint64(l))
	l = len(m.To)
	n += 1 + l + sovTx(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.ClassificationID.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.ImmutableMetaProperties.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.ImmutableProperties.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.MutableMetaProperties.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.MutableProperties.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *message) Unmarshal(dAtA []byte) error {
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
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			if err := m.From.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
			if err := m.To.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassificationID", wireType)
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
			if err := m.ClassificationID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableMetaProperties", wireType)
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
			if err := m.ImmutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImmutableProperties", wireType)
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
			if err := m.ImmutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
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
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
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
			if err := m.MutableProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
