// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/modules/identities/internal/queries/identity/query.proto

package identity

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_helpers "github.com/persistenceOne/persistenceSDK/schema/helpers"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryRequest struct {
	IdentityID github_com_persistenceOne_persistenceSDK_schema_types.ID `protobuf:"bytes,1,opt,name=identity_i_d,json=identityID,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"identity_i_d" valid:"required~required AssetID missing"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d54734af8667f22, []int{0}
}
func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

type QueryResponse struct {
	Success bool                                                               `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   string                                                             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	List    []github_com_persistenceOne_persistenceSDK_schema_helpers.Mappable `protobuf:"bytes,3,rep,name=list,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/helpers.Mappable" json:"list"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d54734af8667f22, []int{1}
}
func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
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
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryRequest)(nil), "persistence_sdk.modules.identities.internal.queries.identity.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "persistence_sdk.modules.identities.internal.queries.identity.QueryResponse")
}

func init() {
	proto.RegisterFile("persistence_sdk/modules/identities/internal/queries/identity/query.proto", fileDescriptor_3d54734af8667f22)
}

var fileDescriptor_3d54734af8667f22 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x31, 0x8b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x77, 0x9e, 0x77, 0x37, 0x9c, 0x4d, 0xb8, 0x62, 0x59, 0x24, 0x09, 0xa9, 0xd6,
	0x66, 0x06, 0xb4, 0x91, 0xc3, 0x42, 0x97, 0x15, 0x6e, 0x3d, 0x44, 0x8c, 0xd8, 0x88, 0xb0, 0xcc,
	0x26, 0x8f, 0xec, 0x60, 0x32, 0x93, 0x9d, 0x37, 0x11, 0xb6, 0xb1, 0xb0, 0xb2, 0x53, 0xf0, 0x0b,
	0x5c, 0x27, 0xf8, 0x49, 0xae, 0xb0, 0x38, 0x10, 0x44, 0x2c, 0x16, 0xd9, 0xb5, 0xb0, 0xf6, 0x13,
	0x48, 0xb2, 0x09, 0xbb, 0x5a, 0x79, 0x78, 0xdd, 0x7b, 0x79, 0xe4, 0xc7, 0xef, 0x3f, 0xf3, 0x86,
	0x1e, 0x17, 0x60, 0x50, 0xa2, 0x05, 0x15, 0xc3, 0x08, 0x93, 0x17, 0x3c, 0xd7, 0x49, 0x99, 0x01,
	0x72, 0x99, 0x80, 0xb2, 0xd2, 0xca, 0xaa, 0x54, 0x16, 0x8c, 0x12, 0x19, 0x9f, 0x96, 0x60, 0xe4,
	0x7a, 0x36, 0xab, 0x3f, 0xcc, 0x58, 0x61, 0xb4, 0xd5, 0xee, 0x9d, 0xbf, 0x48, 0xac, 0x21, 0xb1,
	0x35, 0x89, 0xb5, 0x24, 0xd6, 0x90, 0xda, 0xd9, 0xac, 0x7b, 0x98, 0xea, 0x54, 0xd7, 0x20, 0x5e,
	0x55, 0x2b, 0x66, 0xf7, 0x7a, 0xaa, 0x75, 0x9a, 0x01, 0x17, 0x85, 0xe4, 0x42, 0x29, 0x6d, 0x85,
	0x95, 0x5a, 0xe1, 0x6a, 0x1a, 0x7e, 0x24, 0xf4, 0xe0, 0x71, 0x65, 0x10, 0xc1, 0xb4, 0x04, 0xb4,
	0xee, 0x5b, 0x42, 0x0f, 0x5a, 0xe2, 0x48, 0x8e, 0x92, 0x0e, 0x09, 0x48, 0x6f, 0xbf, 0x9f, 0x9f,
	0xcd, 0x7d, 0xe7, 0xdb, 0xdc, 0xbf, 0x9d, 0x4a, 0x3b, 0x29, 0xc7, 0x2c, 0xd6, 0x39, 0xdf, 0x90,
	0x7d, 0xa4, 0x60, 0xb3, 0x7d, 0x32, 0x38, 0xe1, 0x18, 0x4f, 0x20, 0x17, 0xdc, 0xce, 0x0a, 0x40,
	0x36, 0x1c, 0xfc, 0x9a, 0xfb, 0x37, 0x5e, 0x8a, 0x4c, 0x26, 0x47, 0xa1, 0x81, 0x69, 0x29, 0x0d,
	0x24, 0xaf, 0xda, 0x22, 0xb8, 0x87, 0x08, 0x76, 0x38, 0x08, 0x72, 0x89, 0x28, 0x55, 0x1a, 0x46,
	0xb4, 0x55, 0x18, 0x0e, 0x8e, 0xf6, 0xde, 0x9c, 0xfa, 0xce, 0xcf, 0x53, 0xdf, 0x09, 0x3f, 0x10,
	0x7a, 0xad, 0x91, 0xc5, 0x42, 0x2b, 0x04, 0xb7, 0x43, 0x77, 0xb1, 0x8c, 0x63, 0x40, 0xac, 0x3d,
	0xf7, 0xa2, 0xb6, 0x75, 0x0f, 0xe9, 0x0e, 0x18, 0xa3, 0x4d, 0x67, 0xab, 0xf2, 0x8f, 0x56, 0x8d,
	0xfb, 0x9c, 0x5e, 0xc9, 0x24, 0xda, 0xce, 0x76, 0xb0, 0xdd, 0xdb, 0xef, 0x1f, 0x37, 0xa1, 0xee,
	0x5e, 0x34, 0xd4, 0x04, 0xb2, 0x6a, 0xc0, 0x1e, 0x8a, 0xa2, 0x10, 0xe3, 0x0c, 0xa2, 0x9a, 0xba,
	0x36, 0xbd, 0xf9, 0x85, 0xd0, 0x9d, 0xda, 0xd4, 0xfd, 0x44, 0xe8, 0xee, 0x7d, 0x55, 0xe7, 0x74,
	0x1f, 0xb0, 0xff, 0xb9, 0x5f, 0xb6, 0x79, 0x4f, 0xdd, 0x93, 0x4b, 0x61, 0xad, 0x8e, 0x31, 0xec,
	0xbd, 0xfe, 0xfc, 0xe3, 0xfd, 0x56, 0xe8, 0x06, 0x9b, 0x79, 0xff, 0x58, 0xdf, 0xe6, 0xbf, 0xbe,
	0x3e, 0x5b, 0x78, 0xe4, 0x7c, 0xe1, 0x91, 0xef, 0x0b, 0x8f, 0xbc, 0x5b, 0x7a, 0xce, 0xf9, 0xd2,
	0x73, 0xbe, 0x2e, 0x3d, 0xe7, 0xd9, 0xd3, 0x7f, 0x3e, 0xc4, 0x8b, 0x3c, 0x8f, 0xf1, 0xd5, 0x7a,
	0x4f, 0x6f, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xb0, 0x57, 0x8a, 0x65, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	Enquire(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Enquire(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/persistence_sdk.modules.identities.internal.queries.identity.Query/Enquire", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	Enquire(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Enquire(ctx context.Context, req *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enquire not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Enquire_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Enquire(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/persistence_sdk.modules.identities.internal.queries.identity.Query/Enquire",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Enquire(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "persistence_sdk.modules.identities.internal.queries.identity.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enquire",
			Handler:    _Query_Enquire_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistence_sdk/modules/identities/internal/queries/identity/query.proto",
}

func (m *QueryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.IdentityID.Size()
		i -= size
		if _, err := m.IdentityID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0x12
	}
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

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.IdentityID.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IdentityID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_persistenceOne_persistenceSDK_schema_helpers.Mappable
			m.List = append(m.List, v)
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)