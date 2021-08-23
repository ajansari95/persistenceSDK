// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/assets/internal/transactions/mutate/message.proto

package mutate

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	github_com_persistenceOne_persistenceSDK_schema_types_base "github.com/persistenceOne/persistenceSDK/schema/types/base"
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

type Message struct {
	From                  github_com_persistenceOne_persistenceSDK_schema_types_base.AccAddress `protobuf:"bytes,1,opt,name=from,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types/base.AccAddress" json:"from" valid:"required~required field From missing"`
	FromID                github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,2,opt,name=from_i_d,json=fromID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"from_i_d" valid:"required~required field FromID missing"`
	AssetID               github_com_persistenceOne_persistenceSDK_schema_types.ID              `protobuf:"bytes,3,opt,name=asset_i_d,json=assetID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"asset_i_d" valid:"required~required field AssetID missing"`
	MutableMetaProperties github_com_persistenceOne_persistenceSDK_schema_types.MetaProperties  `protobuf:"bytes,4,opt,name=mutable_meta_properties,json=mutableMetaProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperties" json:"mutable_meta_properties" valid:"required~required field MutableMetaProperties missing"`
	MutableProperties     github_com_persistenceOne_persistenceSDK_schema_types.Properties      `protobuf:"bytes,5,opt,name=mutable_properties,json=mutableProperties,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.Properties" json:"mutable_properties" valid:"required~required field MutableProperties missing"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_595c2732186bdf0e, []int{0}
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

type TransactionResponse struct {
}

func (m *TransactionResponse) Reset()         { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()    {}
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_595c2732186bdf0e, []int{1}
}
func (m *TransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResponse.Merge(m, src)
}
func (m *TransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *TransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "persistence_sdk.modules.assets.internal.transactions.mutate.Message")
	proto.RegisterType((*TransactionResponse)(nil), "persistence_sdk.modules.assets.internal.transactions.mutate.TransactionResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/assets/internal/transactions/mutate/message.proto", fileDescriptor_595c2732186bdf0e)
}

var fileDescriptor_595c2732186bdf0e = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0xd4, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x07, 0xf0, 0x1d, 0x5b, 0xd3, 0x76, 0x6e, 0xae, 0x16, 0x43, 0x0f, 0xbb, 0x21, 0x27, 0x0f,
	0x32, 0x03, 0x7a, 0xd1, 0x8a, 0x60, 0x4a, 0x14, 0x82, 0x04, 0x4b, 0xed, 0xc9, 0x4b, 0x98, 0xec,
	0xbe, 0x6c, 0x07, 0x33, 0x3b, 0xeb, 0xbc, 0x89, 0xe0, 0x41, 0x6f, 0x82, 0x07, 0x45, 0xfd, 0x06,
	0x39, 0xfa, 0x05, 0xfc, 0x0a, 0x52, 0x6f, 0x3d, 0x8a, 0x87, 0x20, 0xc9, 0xc5, 0x73, 0x3e, 0x81,
	0x64, 0x36, 0x31, 0xdb, 0x12, 0x68, 0x4c, 0xbd, 0xbd, 0xcc, 0x90, 0xff, 0xfc, 0xde, 0x2c, 0x6f,
	0x68, 0x23, 0x03, 0x83, 0x12, 0x2d, 0xa4, 0x11, 0xb4, 0x30, 0x7e, 0xce, 0x95, 0x8e, 0x7b, 0x5d,
	0x40, 0x2e, 0x10, 0xc1, 0x22, 0x97, 0xa9, 0x05, 0x93, 0x8a, 0x2e, 0xb7, 0x46, 0xa4, 0x28, 0x22,
	0x2b, 0x75, 0x8a, 0x5c, 0xf5, 0xac, 0xb0, 0xc0, 0x15, 0x20, 0x8a, 0x04, 0x58, 0x66, 0xb4, 0xd5,
	0xfe, 0xbd, 0x33, 0x51, 0x6c, 0x1a, 0xc5, 0xf2, 0x28, 0x36, 0x8b, 0x62, 0xc5, 0x28, 0x96, 0x47,
	0xed, 0x5c, 0x4b, 0x74, 0xa2, 0x5d, 0x0e, 0x9f, 0x54, 0x79, 0x64, 0xf5, 0x7b, 0x89, 0x6e, 0x34,
	0xf3, 0x43, 0xfc, 0xcf, 0x84, 0xae, 0x77, 0x8c, 0x56, 0x65, 0x52, 0x21, 0x37, 0xb6, 0xf6, 0x5e,
	0x1f, 0x0f, 0x42, 0xef, 0xe7, 0x20, 0x7c, 0x98, 0x48, 0x7b, 0xd4, 0x6b, 0xb3, 0x48, 0x2b, 0x5e,
	0x00, 0x3c, 0x49, 0xa1, 0xf8, 0xf3, 0x69, 0xfd, 0x31, 0xc7, 0xe8, 0x08, 0x94, 0xe0, 0xf6, 0x55,
	0x06, 0xc8, 0xdb, 0x02, 0x81, 0xd5, 0xa2, 0xa8, 0x16, 0xc7, 0x06, 0x10, 0xc7, 0x83, 0xf0, 0xe6,
	0x4b, 0xd1, 0x95, 0xf1, 0x6e, 0xd5, 0xc0, 0x8b, 0x9e, 0x34, 0x10, 0xbf, 0x99, 0x15, 0x95, 0x8e,
	0x84, 0x6e, 0x5c, 0x79, 0x64, 0xb4, 0xaa, 0x28, 0x89, 0x28, 0xd3, 0xa4, 0x7a, 0xe0, 0x28, 0xfe,
	0x7b, 0x42, 0x37, 0x27, 0x45, 0x4b, 0xb6, 0xe2, 0xf2, 0x25, 0xe7, 0x32, 0x53, 0xd7, 0x9d, 0x95,
	0x5c, 0xac, 0x51, 0x1f, 0x0f, 0x42, 0xb6, 0x04, 0xa5, 0x51, 0x9f, 0x63, 0x4a, 0x1d, 0xb7, 0xe0,
	0x7f, 0x24, 0x74, 0xcb, 0x5d, 0xb6, 0xf3, 0xac, 0x39, 0x0f, 0xfe, 0x07, 0x0f, 0x3f, 0xc7, 0x53,
	0x9b, 0x9c, 0x59, 0x04, 0x6d, 0x88, 0x7c, 0xc5, 0xff, 0x46, 0xe8, 0xf5, 0xc9, 0x17, 0x6e, 0x77,
	0xa1, 0xa5, 0xc0, 0x8a, 0x56, 0x66, 0x74, 0x06, 0xc6, 0x4a, 0xc0, 0xf2, 0xba, 0xf3, 0x7d, 0x20,
	0x53, 0x60, 0x7d, 0x35, 0x60, 0x13, 0xac, 0xd8, 0xff, 0x1b, 0x3a, 0x1e, 0x84, 0xf7, 0xcf, 0xc1,
	0x36, 0x73, 0xce, 0xe9, 0xff, 0xcd, 0xe9, 0xdb, 0x6a, 0xd1, 0xbe, 0xff, 0x95, 0x50, 0x7f, 0xd6,
	0x48, 0xa1, 0x87, 0xcb, 0xae, 0x87, 0xb7, 0xb3, 0x1e, 0x1e, 0xac, 0xd6, 0xc3, 0x29, 0xff, 0xdd,
	0xe5, 0xfc, 0x8b, 0xec, 0x57, 0xd4, 0xd9, 0xbd, 0xdd, 0xcd, 0x77, 0xfd, 0xd0, 0xfb, 0xdd, 0x0f,
	0xbd, 0xea, 0x36, 0xbd, 0x7a, 0x38, 0x1f, 0xbc, 0x03, 0xc0, 0x4c, 0xa7, 0x08, 0xb7, 0xbe, 0x10,
	0xba, 0xd6, 0xc4, 0xc4, 0xef, 0x13, 0x5a, 0x6a, 0xba, 0x59, 0xf4, 0xeb, 0xec, 0x02, 0x93, 0xcc,
	0xa6, 0xf3, 0xba, 0xb3, 0x7f, 0xa1, 0x94, 0x05, 0xd4, 0xbd, 0xf4, 0x78, 0x18, 0x90, 0x93, 0x61,
	0x40, 0x7e, 0x0d, 0x03, 0xf2, 0x69, 0x14, 0x78, 0x27, 0xa3, 0xc0, 0xfb, 0x31, 0x0a, 0xbc, 0x67,
	0x87, 0x4b, 0xdf, 0xfb, 0x3f, 0x3c, 0x6f, 0xed, 0x92, 0x7b, 0x84, 0x6e, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0xa7, 0x8e, 0x15, 0x24, 0x05, 0x00, 0x00,
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
	Mutate(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Mutate(ctx context.Context, in *Message, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.assets.internal.transactions.mutate.Msg/Mutate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Mutate(context.Context, *Message) (*TransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Mutate(ctx context.Context, req *Message) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mutate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Mutate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mutate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.assets.internal.transactions.mutate.Msg/Mutate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mutate(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.assets.internal.transactions.mutate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mutate",
			Handler:    _Msg_Mutate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/assets/internal/transactions/mutate/message.proto",
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
		size := m.MutableProperties.Size()
		i -= size
		if _, err := m.MutableProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MutableMetaProperties.Size()
		i -= size
		if _, err := m.MutableMetaProperties.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.AssetID.Size()
		i -= size
		if _, err := m.AssetID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.FromID.Size()
		i -= size
		if _, err := m.FromID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.From.Size()
		i -= size
		if _, err := m.From.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMessage(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
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
	l = m.From.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.FromID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.AssetID.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.MutableMetaProperties.Size()
	n += 1 + l + sovMessage(uint64(l))
	l = m.MutableProperties.Size()
	n += 1 + l + sovMessage(uint64(l))
	return n
}

func (m *TransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field FromID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field AssetID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AssetID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableMetaProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MutableMetaProperties.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutableProperties", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
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
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func (m *TransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: TransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)