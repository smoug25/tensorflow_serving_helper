// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/device_properties.proto

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

type DeviceProperties struct {
	// Device type (CPU, GPU, ...)
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Vendor (Intel, nvidia, ...)
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Model (Haswell, K40, ...)
	Model string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// Core Frequency in Mhz
	Frequency int64 `protobuf:"varint,4,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Number of cores
	NumCores int64 `protobuf:"varint,5,opt,name=num_cores,json=numCores,proto3" json:"num_cores,omitempty"`
	// Version of the tools and libraries used with this device (e.g. gcc 4.9,
	// cudnn 5.1)
	Environment map[string]string `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Number of registers per core.
	NumRegisters int64 `protobuf:"varint,7,opt,name=num_registers,json=numRegisters,proto3" json:"num_registers,omitempty"`
	// L1 cache size in bytes
	L1CacheSize int64 `protobuf:"varint,8,opt,name=l1_cache_size,json=l1CacheSize,proto3" json:"l1_cache_size,omitempty"`
	// L2 cache size in bytes
	L2CacheSize int64 `protobuf:"varint,9,opt,name=l2_cache_size,json=l2CacheSize,proto3" json:"l2_cache_size,omitempty"`
	// L3 cache size in bytes
	L3CacheSize int64 `protobuf:"varint,10,opt,name=l3_cache_size,json=l3CacheSize,proto3" json:"l3_cache_size,omitempty"`
	// Shared memory size per multiprocessor in bytes. This field is
	// applicable to GPUs only.
	SharedMemorySizePerMultiprocessor int64 `protobuf:"varint,11,opt,name=shared_memory_size_per_multiprocessor,json=sharedMemorySizePerMultiprocessor,proto3" json:"shared_memory_size_per_multiprocessor,omitempty"`
	// Memory size in bytes
	MemorySize int64 `protobuf:"varint,12,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	// Memory bandwidth in KB/s
	Bandwidth            int64    `protobuf:"varint,13,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProperties) Reset()         { *m = DeviceProperties{} }
func (m *DeviceProperties) String() string { return proto.CompactTextString(m) }
func (*DeviceProperties) ProtoMessage()    {}
func (*DeviceProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c4fdb3c62f9732, []int{0}
}

func (m *DeviceProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProperties.Unmarshal(m, b)
}
func (m *DeviceProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProperties.Marshal(b, m, deterministic)
}
func (m *DeviceProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProperties.Merge(m, src)
}
func (m *DeviceProperties) XXX_Size() int {
	return xxx_messageInfo_DeviceProperties.Size(m)
}
func (m *DeviceProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProperties.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProperties proto.InternalMessageInfo

func (m *DeviceProperties) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DeviceProperties) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *DeviceProperties) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *DeviceProperties) GetFrequency() int64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DeviceProperties) GetNumCores() int64 {
	if m != nil {
		return m.NumCores
	}
	return 0
}

func (m *DeviceProperties) GetEnvironment() map[string]string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *DeviceProperties) GetNumRegisters() int64 {
	if m != nil {
		return m.NumRegisters
	}
	return 0
}

func (m *DeviceProperties) GetL1CacheSize() int64 {
	if m != nil {
		return m.L1CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL2CacheSize() int64 {
	if m != nil {
		return m.L2CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetL3CacheSize() int64 {
	if m != nil {
		return m.L3CacheSize
	}
	return 0
}

func (m *DeviceProperties) GetSharedMemorySizePerMultiprocessor() int64 {
	if m != nil {
		return m.SharedMemorySizePerMultiprocessor
	}
	return 0
}

func (m *DeviceProperties) GetMemorySize() int64 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *DeviceProperties) GetBandwidth() int64 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

type NamedDevice struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Properties           *DeviceProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NamedDevice) Reset()         { *m = NamedDevice{} }
func (m *NamedDevice) String() string { return proto.CompactTextString(m) }
func (*NamedDevice) ProtoMessage()    {}
func (*NamedDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_07c4fdb3c62f9732, []int{1}
}

func (m *NamedDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedDevice.Unmarshal(m, b)
}
func (m *NamedDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedDevice.Marshal(b, m, deterministic)
}
func (m *NamedDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedDevice.Merge(m, src)
}
func (m *NamedDevice) XXX_Size() int {
	return xxx_messageInfo_NamedDevice.Size(m)
}
func (m *NamedDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedDevice.DiscardUnknown(m)
}

var xxx_messageInfo_NamedDevice proto.InternalMessageInfo

func (m *NamedDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedDevice) GetProperties() *DeviceProperties {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceProperties)(nil), "tensorflow.DeviceProperties")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.DeviceProperties.EnvironmentEntry")
	proto.RegisterType((*NamedDevice)(nil), "tensorflow.NamedDevice")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/device_properties.proto", fileDescriptor_07c4fdb3c62f9732)
}

var fileDescriptor_07c4fdb3c62f9732 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0x26, 0x0d, 0xcd, 0xb8, 0x91, 0xa2, 0x15, 0xaa, 0x56, 0xb4, 0x12, 0x21, 0x08,
	0x29, 0x17, 0x12, 0x9a, 0x5c, 0x10, 0x42, 0x1c, 0x5a, 0x7a, 0x2c, 0x44, 0xe6, 0x01, 0x2c, 0xc7,
	0x9e, 0x90, 0x15, 0xde, 0x5d, 0x33, 0xbb, 0x4e, 0xe5, 0x3e, 0x1e, 0x4f, 0xc5, 0x11, 0xed, 0xda,
	0xc4, 0x6e, 0x0e, 0xbd, 0xcd, 0xfc, 0xf3, 0xcd, 0xec, 0x68, 0xe7, 0x87, 0x0f, 0x16, 0x95, 0xd1,
	0xb4, 0xcd, 0xf5, 0xc3, 0x22, 0xd5, 0x84, 0x8b, 0x82, 0xb4, 0xd5, 0x9b, 0x72, 0xbb, 0xc8, 0x70,
	0x2f, 0x52, 0x8c, 0x0b, 0xd2, 0x05, 0x92, 0x15, 0x68, 0xe6, 0xbe, 0xc4, 0xa0, 0xed, 0x98, 0xfe,
	0xe9, 0xc3, 0xf8, 0xab, 0xe7, 0xd6, 0x07, 0x8c, 0x31, 0xe8, 0xdb, 0xaa, 0x40, 0x1e, 0x4c, 0x82,
	0xd9, 0x30, 0xf2, 0x31, 0xbb, 0x80, 0xc1, 0x1e, 0x55, 0xa6, 0x89, 0x9f, 0x78, 0xb5, 0xc9, 0xd8,
	0x4b, 0x38, 0x95, 0x3a, 0xc3, 0x9c, 0xf7, 0xbc, 0x5c, 0x27, 0xec, 0x0a, 0x86, 0x5b, 0xc2, 0xdf,
	0x25, 0xaa, 0xb4, 0xe2, 0xfd, 0x49, 0x30, 0xeb, 0x45, 0xad, 0xc0, 0x2e, 0x61, 0xa8, 0x4a, 0x19,
	0xbb, 0x6d, 0x0d, 0x3f, 0xf5, 0xd5, 0x33, 0x55, 0xca, 0x5b, 0x97, 0xb3, 0xef, 0x10, 0xa2, 0xda,
	0x0b, 0xd2, 0x4a, 0xa2, 0xb2, 0x7c, 0x30, 0xe9, 0xcd, 0xc2, 0xe5, 0xfb, 0x79, 0xbb, 0xf3, 0xfc,
	0x78, 0xdf, 0xf9, 0x5d, 0xcb, 0xdf, 0x29, 0x4b, 0x55, 0xd4, 0x9d, 0xc0, 0xde, 0xc2, 0xc8, 0xbd,
	0x46, 0xf8, 0x53, 0x18, 0x8b, 0x64, 0xf8, 0x0b, 0xff, 0xe2, 0xb9, 0x2a, 0x65, 0xf4, 0x5f, 0x63,
	0x53, 0x18, 0xe5, 0xd7, 0x71, 0x9a, 0xa4, 0x3b, 0x8c, 0x8d, 0x78, 0x44, 0x7e, 0xe6, 0xa1, 0x30,
	0xbf, 0xbe, 0x75, 0xda, 0x0f, 0xf1, 0x88, 0x9e, 0x59, 0x76, 0x99, 0x61, 0xc3, 0x2c, 0x9f, 0x32,
	0xab, 0x2e, 0x03, 0x0d, 0xb3, 0x6a, 0x99, 0x35, 0xbc, 0x33, 0xbb, 0x84, 0x30, 0x8b, 0x25, 0x4a,
	0x4d, 0x95, 0x07, 0xe3, 0x02, 0x29, 0x96, 0x65, 0x6e, 0x45, 0x41, 0x3a, 0x45, 0x63, 0x34, 0xf1,
	0xd0, 0xf7, 0xbe, 0xa9, 0xe1, 0x7b, 0xcf, 0xba, 0x01, 0x6b, 0xa4, 0xfb, 0x27, 0x20, 0x7b, 0x0d,
	0x61, 0x67, 0x14, 0x3f, 0xf7, 0x7d, 0x20, 0x0f, 0x1d, 0xee, 0x1e, 0x9b, 0x44, 0x65, 0x0f, 0x22,
	0xb3, 0x3b, 0x3e, 0xaa, 0xef, 0x71, 0x10, 0x5e, 0x7d, 0x81, 0xf1, 0xf1, 0x17, 0xb2, 0x31, 0xf4,
	0x7e, 0x61, 0xd5, 0x58, 0xc0, 0x85, 0xee, 0xd2, 0xfb, 0x24, 0x2f, 0xb1, 0x31, 0x40, 0x9d, 0x7c,
	0x3a, 0xf9, 0x18, 0x4c, 0x63, 0x08, 0xbf, 0x25, 0x12, 0xb3, 0xfa, 0x30, 0xce, 0x3e, 0x2a, 0x91,
	0x07, 0xfb, 0xb8, 0x98, 0x7d, 0x06, 0x68, 0x7d, 0xe8, 0x27, 0x84, 0xcb, 0xab, 0xe7, 0x8e, 0x1a,
	0x75, 0xf8, 0x9b, 0xcb, 0x9b, 0x8b, 0xe3, 0xfa, 0xda, 0x59, 0xd9, 0xfc, 0x0d, 0x82, 0xcd, 0xc0,
	0xbb, 0x7a, 0xf5, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xaa, 0x8c, 0x04, 0x09, 0x03, 0x00, 0x00,
}
