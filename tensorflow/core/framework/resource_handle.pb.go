// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/resource_handle.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type ResourceHandleProto struct {
	// Unique name for the device containing the resource.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// Container in which this resource is placed.
	Container string `protobuf:"bytes,2,opt,name=container,proto3" json:"container,omitempty"`
	// Unique name of this resource.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Hash code for the type of the resource. Is only valid in the same device
	// and in the same execution.
	HashCode uint64 `protobuf:"varint,4,opt,name=hash_code,json=hashCode,proto3" json:"hash_code,omitempty"`
	// For debug-only, the name of the type pointed to by this handle, if
	// available.
	MaybeTypeName string `protobuf:"bytes,5,opt,name=maybe_type_name,json=maybeTypeName,proto3" json:"maybe_type_name,omitempty"`
	// Data types and shapes for the underlying resource.
	DtypesAndShapes      []*ResourceHandleProto_DtypeAndShape `protobuf:"bytes,6,rep,name=dtypes_and_shapes,json=dtypesAndShapes,proto3" json:"dtypes_and_shapes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ResourceHandleProto) Reset()         { *m = ResourceHandleProto{} }
func (m *ResourceHandleProto) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto) ProtoMessage()    {}
func (*ResourceHandleProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36024d2bd9a2afd, []int{0}
}

func (m *ResourceHandleProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceHandleProto.Unmarshal(m, b)
}
func (m *ResourceHandleProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceHandleProto.Marshal(b, m, deterministic)
}
func (m *ResourceHandleProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandleProto.Merge(m, src)
}
func (m *ResourceHandleProto) XXX_Size() int {
	return xxx_messageInfo_ResourceHandleProto.Size(m)
}
func (m *ResourceHandleProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandleProto.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandleProto proto.InternalMessageInfo

func (m *ResourceHandleProto) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *ResourceHandleProto) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ResourceHandleProto) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceHandleProto) GetHashCode() uint64 {
	if m != nil {
		return m.HashCode
	}
	return 0
}

func (m *ResourceHandleProto) GetMaybeTypeName() string {
	if m != nil {
		return m.MaybeTypeName
	}
	return ""
}

func (m *ResourceHandleProto) GetDtypesAndShapes() []*ResourceHandleProto_DtypeAndShape {
	if m != nil {
		return m.DtypesAndShapes
	}
	return nil
}

// Protocol buffer representing a pair of (data type, tensor shape).
type ResourceHandleProto_DtypeAndShape struct {
	Dtype                DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResourceHandleProto_DtypeAndShape) Reset()         { *m = ResourceHandleProto_DtypeAndShape{} }
func (m *ResourceHandleProto_DtypeAndShape) String() string { return proto.CompactTextString(m) }
func (*ResourceHandleProto_DtypeAndShape) ProtoMessage()    {}
func (*ResourceHandleProto_DtypeAndShape) Descriptor() ([]byte, []int) {
	return fileDescriptor_a36024d2bd9a2afd, []int{0, 0}
}

func (m *ResourceHandleProto_DtypeAndShape) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Unmarshal(m, b)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Marshal(b, m, deterministic)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Merge(m, src)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_Size() int {
	return xxx_messageInfo_ResourceHandleProto_DtypeAndShape.Size(m)
}
func (m *ResourceHandleProto_DtypeAndShape) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceHandleProto_DtypeAndShape.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceHandleProto_DtypeAndShape proto.InternalMessageInfo

func (m *ResourceHandleProto_DtypeAndShape) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *ResourceHandleProto_DtypeAndShape) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceHandleProto)(nil), "tensorflow.ResourceHandleProto")
	proto.RegisterType((*ResourceHandleProto_DtypeAndShape)(nil), "tensorflow.ResourceHandleProto.DtypeAndShape")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/resource_handle.proto", fileDescriptor_a36024d2bd9a2afd)
}

var fileDescriptor_a36024d2bd9a2afd = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x53, 0xfe, 0x45, 0x86, 0x00, 0x71, 0x35, 0xa6, 0x41, 0x0e, 0xc4, 0x44, 0x43, 0x8c,
	0xb6, 0x49, 0xfd, 0x04, 0x22, 0x07, 0x4f, 0xc6, 0x54, 0x2e, 0x9e, 0x9a, 0xa5, 0x3b, 0x58, 0x22,
	0xec, 0x34, 0xbb, 0x55, 0xc2, 0xb7, 0xf6, 0xe8, 0xd1, 0x74, 0x16, 0x05, 0x8d, 0x7a, 0xdb, 0x9d,
	0xf9, 0xbd, 0x99, 0xf7, 0x76, 0x21, 0x2c, 0x50, 0x5b, 0x32, 0xb3, 0x05, 0xad, 0xc2, 0x94, 0x0c,
	0x86, 0x33, 0x23, 0x97, 0xb8, 0x22, 0xf3, 0x1c, 0x1a, 0xb4, 0xf4, 0x62, 0x52, 0x4c, 0x32, 0xa9,
	0xd5, 0x02, 0x83, 0xdc, 0x50, 0x41, 0x02, 0xb6, 0x82, 0xde, 0xc5, 0xdf, 0x62, 0xd7, 0x49, 0x6c,
	0x26, 0xf3, 0x8d, 0xb2, 0x77, 0xfa, 0x0f, 0xbd, 0xce, 0xd1, 0x3a, 0xec, 0xe4, 0xad, 0x02, 0x07,
	0xf1, 0x66, 0xf5, 0x2d, 0x6f, 0xbe, 0xe7, 0xc5, 0x47, 0xd0, 0x50, 0xf8, 0x3a, 0x4f, 0xd1, 0xf7,
	0x06, 0xde, 0xb0, 0x19, 0x6f, 0x6e, 0xa2, 0x0f, 0xcd, 0x94, 0x74, 0x21, 0xe7, 0x1a, 0x8d, 0x5f,
	0xe1, 0xd6, 0xb6, 0x20, 0x04, 0xd4, 0xb4, 0x5c, 0xa2, 0x5f, 0xe5, 0x06, 0x9f, 0xc5, 0x31, 0x34,
	0x33, 0x69, 0xb3, 0x24, 0x25, 0x85, 0x7e, 0x6d, 0xe0, 0x0d, 0x6b, 0xf1, 0x5e, 0x59, 0xb8, 0x21,
	0x85, 0xe2, 0x0c, 0xba, 0x4b, 0xb9, 0x9e, 0x62, 0x52, 0x7a, 0x4a, 0x58, 0x5b, 0x67, 0x6d, 0x9b,
	0xcb, 0x93, 0x75, 0x8e, 0x77, 0xe5, 0x90, 0x47, 0xd8, 0x57, 0x6c, 0x3b, 0x91, 0x5a, 0xb9, 0x9c,
	0xd6, 0x6f, 0x0c, 0xaa, 0xc3, 0x56, 0x74, 0x19, 0x6c, 0x93, 0x06, 0xbf, 0x44, 0x09, 0xc6, 0xa5,
	0xf0, 0x5a, 0xab, 0x87, 0x52, 0x15, 0x77, 0xdd, 0x9c, 0xcf, 0xbb, 0xed, 0x11, 0xb4, 0xbf, 0x11,
	0xe2, 0x1c, 0xea, 0xcc, 0x70, 0xf2, 0x4e, 0x74, 0xb8, 0x3b, 0x7f, 0x2c, 0x0b, 0x59, 0x9a, 0x8a,
	0x1d, 0x22, 0x22, 0xa8, 0xb3, 0x19, 0x7e, 0x8a, 0x56, 0xd4, 0xdf, 0x65, 0x27, 0x7c, 0xe4, 0x99,
	0x6c, 0x24, 0x76, 0xe8, 0x28, 0x04, 0x9f, 0xcc, 0xd3, 0x2e, 0xf9, 0xf5, 0x35, 0xa3, 0xce, 0x8f,
	0x00, 0xde, 0xbb, 0xe7, 0x4d, 0x1b, 0xfc, 0x55, 0x57, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x75,
	0x87, 0x3c, 0x31, 0x3e, 0x02, 0x00, 0x00,
}