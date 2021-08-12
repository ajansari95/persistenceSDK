// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/stringData.proto

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

type StringData struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *StringData) Reset()      { *m = StringData{} }
func (*StringData) ProtoMessage() {}
func (*StringData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9921211417c39aeb, []int{0}
}
func (m *StringData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringData.Unmarshal(m, b)
}
func (m *StringData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringData.Marshal(b, m, deterministic)
}
func (m *StringData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringData.Merge(m, src)
}
func (m *StringData) XXX_Size() int {
	return xxx_messageInfo_StringData.Size(m)
}
func (m *StringData) XXX_DiscardUnknown() {
	xxx_messageInfo_StringData.DiscardUnknown(m)
}

var xxx_messageInfo_StringData proto.InternalMessageInfo

func (m *StringData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringData)(nil), "schema.types.base.StringData")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/stringData.proto", fileDescriptor_9921211417c39aeb)
}

var fileDescriptor_9921211417c39aeb = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2a, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x2f, 0x2e,
	0x29, 0xca, 0xcc, 0x4b, 0x77, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x84, 0xa8, 0xd1, 0x03, 0xab, 0xd1, 0x03, 0xa9, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb,
	0xea, 0x83, 0x58, 0x10, 0x85, 0x4a, 0x26, 0x5c, 0x5c, 0xc1, 0x70, 0xcd, 0x42, 0x22, 0x5c, 0xac,
	0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x95, 0xc0,
	0x8c, 0x05, 0xf2, 0x0c, 0x1d, 0x0b, 0xe5, 0x19, 0x26, 0x2c, 0x94, 0x67, 0x58, 0xb0, 0x50, 0x9e,
	0xc1, 0x29, 0xe4, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xac, 0xd2, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x91, 0xdc, 0xed, 0x9f, 0x97, 0x8a, 0xcc,
	0x0d, 0x76, 0xf1, 0xc6, 0xf4, 0x45, 0x12, 0x1b, 0xd8, 0x49, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0xce, 0x59, 0x51, 0xf1, 0x00, 0x00, 0x00,
}