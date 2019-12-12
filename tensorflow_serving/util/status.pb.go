// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/util/status.proto

package util

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework"
	_ "github.com/smoug25/tensorflow_serving_helper/tensorflow/core/lib/core"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Status that corresponds to Status in
// third_party/tensorflow/core/lib/core/status.h.
type StatusProto struct {
	// Error code.
	ErrorCode framework.Code `protobuf:"varint,1,opt,name=error_code,proto3,enum=tensorflow.error.Code" json:"error_code,omitempty"`
	// Error message. Will only be set if an error was encountered.
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusProto) Reset()         { *m = StatusProto{} }
func (m *StatusProto) String() string { return proto.CompactTextString(m) }
func (*StatusProto) ProtoMessage()    {}
func (*StatusProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_e97f03f1515e99b0, []int{0}
}

func (m *StatusProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusProto.Unmarshal(m, b)
}
func (m *StatusProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusProto.Marshal(b, m, deterministic)
}
func (m *StatusProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusProto.Merge(m, src)
}
func (m *StatusProto) XXX_Size() int {
	return xxx_messageInfo_StatusProto.Size(m)
}
func (m *StatusProto) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusProto.DiscardUnknown(m)
}

var xxx_messageInfo_StatusProto proto.InternalMessageInfo

func (m *StatusProto) GetErrorCode() framework.Code {
	if m != nil {
		return m.ErrorCode
	}
	return framework.Code_OK
}

func (m *StatusProto) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusProto)(nil), "tensorflow.serving.StatusProto")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/util/status.proto", fileDescriptor_e97f03f1515e99b0)
}

var fileDescriptor_e97f03f1515e99b0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcd, 0x2f, 0x4d, 0x37, 0x32, 0xd5, 0x2f,
	0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc,
	0x4b, 0x8f, 0xcf, 0x48, 0xcd, 0x29, 0x48, 0x2d, 0xc2, 0x22, 0xa3, 0x5f, 0x5a, 0x92, 0x99, 0xa3,
	0x5f, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0xac, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x84, 0x50,
	0xa5, 0x07, 0x55, 0x25, 0x15, 0x4e, 0xa6, 0x35, 0xfa, 0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x39, 0x99,
	0x49, 0x10, 0x46, 0x6a, 0x51, 0x51, 0x7e, 0x51, 0x7c, 0x72, 0x7e, 0x4a, 0x2a, 0xd4, 0x32, 0xa5,
	0x6c, 0x2e, 0xee, 0x60, 0xb0, 0xe5, 0x01, 0x60, 0xbb, 0xcd, 0xb8, 0xb8, 0x10, 0x6a, 0x24, 0x18,
	0x15, 0x18, 0x35, 0xf8, 0x8c, 0xc4, 0xf4, 0x90, 0x1c, 0x04, 0x96, 0xd5, 0x73, 0xce, 0x4f, 0x49,
	0x0d, 0x42, 0x52, 0x29, 0xa4, 0xc2, 0xc5, 0x0b, 0xe1, 0xe5, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7,
	0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x3a, 0x79, 0x46, 0xb9, 0x50, 0x23, 0xb8,
	0x7e, 0x30, 0x32, 0x26, 0xb1, 0x81, 0x9d, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x79,
	0x75, 0xa5, 0x80, 0x01, 0x00, 0x00,
}