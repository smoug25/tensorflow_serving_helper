// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/critical_section.proto

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

// Protocol buffer representing a CriticalSection.
type CriticalSectionDef struct {
	// Name of the critical section handle.
	CriticalSectionName  string   `protobuf:"bytes,1,opt,name=critical_section_name,json=criticalSectionName,proto3" json:"critical_section_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CriticalSectionDef) Reset()         { *m = CriticalSectionDef{} }
func (m *CriticalSectionDef) String() string { return proto.CompactTextString(m) }
func (*CriticalSectionDef) ProtoMessage()    {}
func (*CriticalSectionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c10c289f67aabd, []int{0}
}

func (m *CriticalSectionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalSectionDef.Unmarshal(m, b)
}
func (m *CriticalSectionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalSectionDef.Marshal(b, m, deterministic)
}
func (m *CriticalSectionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalSectionDef.Merge(m, src)
}
func (m *CriticalSectionDef) XXX_Size() int {
	return xxx_messageInfo_CriticalSectionDef.Size(m)
}
func (m *CriticalSectionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalSectionDef.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalSectionDef proto.InternalMessageInfo

func (m *CriticalSectionDef) GetCriticalSectionName() string {
	if m != nil {
		return m.CriticalSectionName
	}
	return ""
}

// Protocol buffer representing a CriticalSection execution.
type CriticalSectionExecutionDef struct {
	// Name of the critical section handle.
	ExecuteInCriticalSectionName string `protobuf:"bytes,1,opt,name=execute_in_critical_section_name,json=executeInCriticalSectionName,proto3" json:"execute_in_critical_section_name,omitempty"`
	// Whether this operation requires exclusive access to its resources,
	// (i.e., no other CriticalSections may request the same resources).
	ExclusiveResourceAccess bool     `protobuf:"varint,2,opt,name=exclusive_resource_access,json=exclusiveResourceAccess,proto3" json:"exclusive_resource_access,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CriticalSectionExecutionDef) Reset()         { *m = CriticalSectionExecutionDef{} }
func (m *CriticalSectionExecutionDef) String() string { return proto.CompactTextString(m) }
func (*CriticalSectionExecutionDef) ProtoMessage()    {}
func (*CriticalSectionExecutionDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c10c289f67aabd, []int{1}
}

func (m *CriticalSectionExecutionDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CriticalSectionExecutionDef.Unmarshal(m, b)
}
func (m *CriticalSectionExecutionDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CriticalSectionExecutionDef.Marshal(b, m, deterministic)
}
func (m *CriticalSectionExecutionDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CriticalSectionExecutionDef.Merge(m, src)
}
func (m *CriticalSectionExecutionDef) XXX_Size() int {
	return xxx_messageInfo_CriticalSectionExecutionDef.Size(m)
}
func (m *CriticalSectionExecutionDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CriticalSectionExecutionDef.DiscardUnknown(m)
}

var xxx_messageInfo_CriticalSectionExecutionDef proto.InternalMessageInfo

func (m *CriticalSectionExecutionDef) GetExecuteInCriticalSectionName() string {
	if m != nil {
		return m.ExecuteInCriticalSectionName
	}
	return ""
}

func (m *CriticalSectionExecutionDef) GetExclusiveResourceAccess() bool {
	if m != nil {
		return m.ExclusiveResourceAccess
	}
	return false
}

func init() {
	proto.RegisterType((*CriticalSectionDef)(nil), "tensorflow.CriticalSectionDef")
	proto.RegisterType((*CriticalSectionExecutionDef)(nil), "tensorflow.CriticalSectionExecutionDef")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/critical_section.proto", fileDescriptor_f4c10c289f67aabd)
}

var fileDescriptor_f4c10c289f67aabd = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0x07, 0xd1, 0x1c, 0x2b, 0xc3, 0x8a, 0x1e, 0xca, 0x4e, 0x3b, 0xb5, 0x30, 0xd1,
	0x83, 0x37, 0x37, 0x15, 0xbd, 0xc8, 0xa8, 0x57, 0x21, 0xa4, 0x8f, 0xd7, 0x2e, 0xd8, 0xe4, 0x8d,
	0x97, 0x64, 0xdb, 0xbf, 0xe3, 0x7f, 0xe9, 0x51, 0xac, 0xd5, 0x69, 0x05, 0x8f, 0xe1, 0xf7, 0xfb,
	0xbe, 0x47, 0x3e, 0xf9, 0xdc, 0x98, 0xb0, 0x8c, 0x55, 0x0e, 0x64, 0x0b, 0x6f, 0x29, 0x36, 0xd3,
	0x8b, 0x22, 0xa0, 0xf3, 0xc4, 0x75, 0x4b, 0x1b, 0xe5, 0x91, 0xd7, 0xc6, 0x35, 0x6a, 0x89, 0xed,
	0x0a, 0xf9, 0x07, 0x29, 0x80, 0x18, 0x8b, 0x15, 0x53, 0xa0, 0x2a, 0xd6, 0x05, 0xb0, 0x09, 0x06,
	0x74, 0xab, 0x3c, 0x42, 0x30, 0xe4, 0xf2, 0x8e, 0x24, 0x72, 0x17, 0x18, 0xdf, 0xcb, 0x64, 0xde,
	0x5b, 0x4f, 0x9f, 0xd2, 0x0d, 0xd6, 0xc9, 0x54, 0x8e, 0x86, 0x59, 0xe5, 0xb4, 0xc5, 0x54, 0x64,
	0x62, 0x72, 0x58, 0x1e, 0xc1, 0xef, 0xc8, 0xa3, 0xb6, 0x38, 0x7e, 0x15, 0xf2, 0x74, 0x50, 0x75,
	0xbb, 0x45, 0x88, 0x5f, 0x9d, 0x77, 0x32, 0xc3, 0xee, 0x8d, 0xca, 0x38, 0xf5, 0x5f, 0xfd, 0x59,
	0xef, 0x3d, 0xb8, 0xf9, 0xdf, 0x3b, 0xc9, 0x95, 0x3c, 0xc1, 0x2d, 0xb4, 0xd1, 0x9b, 0x35, 0x2a,
	0x46, 0x4f, 0x91, 0x01, 0x95, 0x06, 0x40, 0xef, 0xd3, 0xbd, 0x4c, 0x4c, 0x0e, 0xca, 0xe3, 0x6f,
	0xa1, 0xec, 0xf9, 0x75, 0x87, 0x67, 0x97, 0x32, 0x25, 0x6e, 0xf2, 0xdd, 0xff, 0xf3, 0x9a, 0xb5,
	0xc5, 0x0d, 0xf1, 0xcb, 0x6c, 0x34, 0x38, 0xb6, 0xf8, 0xd8, 0xca, 0x2f, 0xc4, 0x9b, 0x10, 0xd5,
	0x7e, 0x37, 0xdc, 0xf9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0x32, 0x0c, 0x77, 0x98, 0x01,
	0x00, 0x00,
}
