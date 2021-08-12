// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/metaProperty.proto

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

type MetaProperty struct {
	Id       github_com_persistenceOne_persistenceSDK_schema_types.ID       `protobuf:"bytes,1,opt,name=id,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.ID" json:"id"`
	MetaFact github_com_persistenceOne_persistenceSDK_schema_types.MetaFact `protobuf:"bytes,2,opt,name=metaFact,proto3,customtype=github.com/persistenceOne/persistenceSDK/schema/types.MetaFact" json:"metaFact"`
}

func (m *MetaProperty) Reset()         { *m = MetaProperty{} }
func (m *MetaProperty) String() string { return proto.CompactTextString(m) }
func (*MetaProperty) ProtoMessage()    {}
func (*MetaProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e00c620db06edd3e, []int{0}
}
func (m *MetaProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaProperty.Unmarshal(m, b)
}
func (m *MetaProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaProperty.Marshal(b, m, deterministic)
}
func (m *MetaProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaProperty.Merge(m, src)
}
func (m *MetaProperty) XXX_Size() int {
	return xxx_messageInfo_MetaProperty.Size(m)
}
func (m *MetaProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaProperty.DiscardUnknown(m)
}

var xxx_messageInfo_MetaProperty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MetaProperty)(nil), "schema.types.base.MetaProperty")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/metaProperty.proto", fileDescriptor_e00c620db06edd3e)
}

var fileDescriptor_e00c620db06edd3e = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4d,
	0x2d, 0x49, 0x0c, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x84, 0xa8, 0xd2, 0x03, 0xab, 0xd2, 0x03, 0xa9, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf,
	0x07, 0xcb, 0xea, 0x83, 0x58, 0x10, 0x85, 0x4a, 0x97, 0x18, 0xb9, 0x78, 0x7c, 0x91, 0xf4, 0x0b,
	0x05, 0x70, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0x39, 0x9c, 0xb8, 0x27,
	0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x45, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x3e, 0x92, 0x73, 0xfc, 0xf3, 0x52, 0x91, 0xb9, 0xc1, 0x2e, 0xde, 0x28, 0x8e, 0xd3, 0xf3, 0x74,
	0x09, 0x62, 0xca, 0x4c, 0x11, 0x4a, 0xe2, 0xe2, 0x00, 0xb9, 0xd0, 0x2d, 0x31, 0xb9, 0x44, 0x82,
	0x09, 0x6c, 0xae, 0x1b, 0xd4, 0x5c, 0x3b, 0xf2, 0xcc, 0xf5, 0x85, 0x9a, 0x16, 0x04, 0x37, 0xd7,
	0x8a, 0xa7, 0x63, 0xa1, 0x3c, 0xc3, 0x84, 0x85, 0xf2, 0x0c, 0x0b, 0x16, 0xca, 0x33, 0x38, 0x85,
	0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x15, 0x59, 0x36, 0x82, 0x83,
	0x39, 0x89, 0x0d, 0x1c, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xad, 0x70, 0x3f,
	0x92, 0x01, 0x00, 0x00,
}