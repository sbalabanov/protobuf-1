// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merge.proto

package merge

import (
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type A struct {
	B                    B        `protobuf:"bytes,1,opt,name=B,proto3" json:"B"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A) Reset()         { *m = A{} }
func (m *A) String() string { return proto.CompactTextString(m) }
func (*A) ProtoMessage()    {}
func (*A) Descriptor() ([]byte, []int) {
	return fileDescriptor_82caea6f5430298b, []int{0}
}
func (m *A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A.Unmarshal(m, b)
}
func (m *A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A.Marshal(b, m, deterministic)
}
func (m *A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A.Merge(m, src)
}
func (m *A) XXX_Size() int {
	return xxx_messageInfo_A.Size(m)
}
func (m *A) XXX_DiscardUnknown() {
	xxx_messageInfo_A.DiscardUnknown(m)
}

var xxx_messageInfo_A proto.InternalMessageInfo

func (m *A) GetB() B {
	if m != nil {
		return m.B
	}
	return B{}
}

type B struct {
	C                    int64    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *B) Reset()         { *m = B{} }
func (m *B) String() string { return proto.CompactTextString(m) }
func (*B) ProtoMessage()    {}
func (*B) Descriptor() ([]byte, []int) {
	return fileDescriptor_82caea6f5430298b, []int{1}
}
func (m *B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_B.Unmarshal(m, b)
}
func (m *B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_B.Marshal(b, m, deterministic)
}
func (m *B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_B.Merge(m, src)
}
func (m *B) XXX_Size() int {
	return xxx_messageInfo_B.Size(m)
}
func (m *B) XXX_DiscardUnknown() {
	xxx_messageInfo_B.DiscardUnknown(m)
}

var xxx_messageInfo_B proto.InternalMessageInfo

func (m *B) GetC() int64 {
	if m != nil {
		return m.C
	}
	return 0
}

func init() {
	proto.RegisterType((*A)(nil), "merge.A")
	proto.RegisterType((*B)(nil), "merge.B")
}

func init() { proto.RegisterFile("merge.proto", fileDescriptor_82caea6f5430298b) }

var fileDescriptor_82caea6f5430298b = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x4d, 0x2d, 0x4a,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x74, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xb2, 0x49,
	0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x74, 0x29, 0x29, 0x72, 0x31, 0x3a, 0x0a, 0xc9,
	0x70, 0x31, 0x3a, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xe8, 0x41, 0xcc, 0x74, 0x72,
	0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x88, 0xd1, 0x49, 0x49, 0x90, 0x8b, 0xd1, 0x49, 0x88, 0x87,
	0x8b, 0x31, 0x19, 0xac, 0x84, 0x39, 0x88, 0x31, 0x39, 0x89, 0x0d, 0xac, 0xd9, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x86, 0x7b, 0xb3, 0xe7, 0x81, 0x00, 0x00, 0x00,
}
