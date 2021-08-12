// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: persistence_sdk/schema/types/base/listData.proto

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

type ListData struct {
	Value sortedDataList `protobuf:"bytes,1,opt,name=value,proto3,customtype=sortedDataList" json:"value"`
}

func (m *ListData) Reset()      { *m = ListData{} }
func (*ListData) ProtoMessage() {}
func (*ListData) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f65daeca57ea5b6, []int{0}
}
func (m *ListData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListData.Unmarshal(m, b)
}
func (m *ListData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListData.Marshal(b, m, deterministic)
}
func (m *ListData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListData.Merge(m, src)
}
func (m *ListData) XXX_Size() int {
	return xxx_messageInfo_ListData.Size(m)
}
func (m *ListData) XXX_DiscardUnknown() {
	xxx_messageInfo_ListData.DiscardUnknown(m)
}

var xxx_messageInfo_ListData proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListData)(nil), "schema.types.base.ListData")
}

func init() {
	proto.RegisterFile("persistence_sdk/schema/types/base/listData.proto", fileDescriptor_3f65daeca57ea5b6)
}

var fileDescriptor_3f65daeca57ea5b6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x28, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x4e, 0x8d, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x4e, 0xce, 0x48,
	0xcd, 0x4d, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0xcf, 0xc9,
	0x2c, 0x2e, 0x71, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0xa8,
	0xd0, 0x03, 0xab, 0xd0, 0x03, 0xa9, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83,
	0x58, 0x10, 0x85, 0x4a, 0x5e, 0x5c, 0x1c, 0x3e, 0x50, 0xad, 0x42, 0x3a, 0x5c, 0xac, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0x62, 0x27, 0xee, 0xc9, 0x33, 0xdc,
	0xba, 0x27, 0xcf, 0x57, 0x9c, 0x5f, 0x54, 0x92, 0x9a, 0x02, 0x52, 0x02, 0x52, 0x1a, 0x04, 0x51,
	0x64, 0x25, 0x30, 0x63, 0x81, 0x3c, 0x43, 0xc7, 0x42, 0x79, 0x86, 0x09, 0x0b, 0xe5, 0x19, 0x16,
	0x2c, 0x94, 0x67, 0x70, 0x0a, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0xab, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x24, 0xbf, 0xf8, 0xe7,
	0xa5, 0x22, 0x73, 0x83, 0x5d, 0xbc, 0x31, 0x7d, 0x96, 0xc4, 0x06, 0x76, 0xa8, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0xec, 0x33, 0x5c, 0x05, 0x01, 0x00, 0x00,
}