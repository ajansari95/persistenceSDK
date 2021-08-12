// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/metaProperties.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_persistenceOne_persistenceSDK_schema_types "github.com/persistenceOne/persistenceSDK/schema/types"
	math "math"
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

type MetaProperties struct {
	MetaPropertyList []github_com_persistenceOne_persistenceSDK_schema_types.MetaProperty `protobuf:"bytes,1,rep,name=metaPropertyList,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaProperty" json:"metaPropertyList"`
}

func (m *MetaProperties) Reset()         { *m = MetaProperties{} }
func (m *MetaProperties) String() string { return proto.CompactTextString(m) }
func (*MetaProperties) ProtoMessage()    {}
func (*MetaProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fb9d7c35b622115, []int{0}
}
func (m *MetaProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaProperties.Unmarshal(m, b)
}
func (m *MetaProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaProperties.Marshal(b, m, deterministic)
}
func (m *MetaProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaProperties.Merge(m, src)
}
func (m *MetaProperties) XXX_Size() int {
	return xxx_messageInfo_MetaProperties.Size(m)
}
func (m *MetaProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MetaProperties proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MetaProperties)(nil), "schema.types.base.MetaProperties")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/metaProperties.proto", fileDescriptor_8fb9d7c35b622115)
}

var fileDescriptor_8fb9d7c35b622115 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2b, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4d,
	0x2d, 0x49, 0x0c, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x84, 0xa8, 0xd3, 0x03, 0xab, 0xd3, 0x03, 0xa9, 0x93, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83, 0x58, 0x10, 0x85, 0x4a, 0xd3, 0x18, 0xb9, 0xf8, 0x7c, 0x51,
	0x4c, 0x10, 0x2a, 0xe3, 0x12, 0x40, 0x32, 0xb3, 0xd2, 0x27, 0xb3, 0xb8, 0x44, 0x82, 0x51, 0x81,
	0x59, 0x83, 0xd3, 0xc9, 0xeb, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4, 0x9d, 0xd2, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x91, 0x1c, 0xe8, 0x9f, 0x97, 0x8a, 0xcc, 0x0d,
	0x76, 0xf1, 0x46, 0x71, 0xae, 0x1e, 0x92, 0x3d, 0x95, 0x41, 0x18, 0x76, 0x58, 0xf1, 0x74, 0x2c,
	0x94, 0x67, 0x98, 0xb0, 0x50, 0x9e, 0x61, 0xc1, 0x42, 0x79, 0x06, 0xa7, 0x90, 0x13, 0x8f, 0xe4,
	0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f,
	0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x22, 0xcb, 0x76, 0x70, 0x60, 0x25, 0xb1, 0x81,
	0x7d, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x58, 0x9a, 0x73, 0x3f, 0x58, 0x01, 0x00, 0x00,
}