// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/id.proto

package base

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ID struct {
	IdString string `protobuf:"bytes,1,opt,name=id_string,json=idString,proto3" json:"id_string,omitempty"`
}

func (m *ID) Reset()      { *m = ID{} }
func (*ID) ProtoMessage() {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_31e3a90fdeb3b61b, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetIdString() string {
	if m != nil {
		return m.IdString
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "schema.types.base.ID")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/id.proto", fileDescriptor_31e3a90fdeb3b61b)
}

var fileDescriptor_31e3a90fdeb3b61b = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xc8, 0xe9, 0x81, 0xe5, 0xf4, 0x40, 0x72,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x59, 0x7d, 0x10, 0x0b, 0xa2, 0x50, 0xc9, 0x98, 0x8b,
	0xc9, 0xd3, 0x45, 0x48, 0x9a, 0x8b, 0x33, 0x33, 0x25, 0xbe, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x5d,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0x23, 0x33, 0x25, 0x18, 0xcc, 0xb7, 0x12, 0x98, 0xb1,
	0x40, 0x9e, 0xa1, 0x63, 0xa1, 0x3c, 0xc3, 0x84, 0x85, 0xf2, 0x0c, 0x0b, 0x16, 0xca, 0x33, 0x38,
	0x85, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x55, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x92, 0x73, 0xfd, 0xf3, 0x52, 0x91, 0xb9, 0xc1,
	0x2e, 0xde, 0x98, 0x8e, 0x4f, 0x62, 0x03, 0xbb, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0x4a, 0x9b, 0x26, 0xe8, 0x00, 0x00, 0x00,
}