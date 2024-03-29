// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/verifier_config.proto

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

type VerifierConfig_Toggle int32

const (
	VerifierConfig_DEFAULT VerifierConfig_Toggle = 0
	VerifierConfig_ON      VerifierConfig_Toggle = 1
	VerifierConfig_OFF     VerifierConfig_Toggle = 2
)

var VerifierConfig_Toggle_name = map[int32]string{
	0: "DEFAULT",
	1: "ON",
	2: "OFF",
}

var VerifierConfig_Toggle_value = map[string]int32{
	"DEFAULT": 0,
	"ON":      1,
	"OFF":     2,
}

func (x VerifierConfig_Toggle) String() string {
	return proto.EnumName(VerifierConfig_Toggle_name, int32(x))
}

func (VerifierConfig_Toggle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5049fcf5d8bb3c3c, []int{0, 0}
}

// The config for graph verifiers.
type VerifierConfig struct {
	// Deadline for completion of all verification i.e. all the Toggle ON
	// verifiers must complete execution within this time.
	VerificationTimeoutInMs int64 `protobuf:"varint,1,opt,name=verification_timeout_in_ms,json=verificationTimeoutInMs,proto3" json:"verification_timeout_in_ms,omitempty"`
	// Perform structural validation on a tensorflow graph. Default is OFF.
	StructureVerifier    VerifierConfig_Toggle `protobuf:"varint,2,opt,name=structure_verifier,json=structureVerifier,proto3,enum=tensorflow.VerifierConfig_Toggle" json:"structure_verifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *VerifierConfig) Reset()         { *m = VerifierConfig{} }
func (m *VerifierConfig) String() string { return proto.CompactTextString(m) }
func (*VerifierConfig) ProtoMessage()    {}
func (*VerifierConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5049fcf5d8bb3c3c, []int{0}
}

func (m *VerifierConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifierConfig.Unmarshal(m, b)
}
func (m *VerifierConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifierConfig.Marshal(b, m, deterministic)
}
func (m *VerifierConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifierConfig.Merge(m, src)
}
func (m *VerifierConfig) XXX_Size() int {
	return xxx_messageInfo_VerifierConfig.Size(m)
}
func (m *VerifierConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifierConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VerifierConfig proto.InternalMessageInfo

func (m *VerifierConfig) GetVerificationTimeoutInMs() int64 {
	if m != nil {
		return m.VerificationTimeoutInMs
	}
	return 0
}

func (m *VerifierConfig) GetStructureVerifier() VerifierConfig_Toggle {
	if m != nil {
		return m.StructureVerifier
	}
	return VerifierConfig_DEFAULT
}

func init() {
	proto.RegisterEnum("tensorflow.VerifierConfig_Toggle", VerifierConfig_Toggle_name, VerifierConfig_Toggle_value)
	proto.RegisterType((*VerifierConfig)(nil), "tensorflow.VerifierConfig")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/verifier_config.proto", fileDescriptor_5049fcf5d8bb3c3c)
}

var fileDescriptor_5049fcf5d8bb3c3c = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x46, 0xdd, 0x14, 0x52, 0x18, 0xa1, 0xc4, 0x45, 0x30, 0x78, 0xaa, 0x3d, 0x48, 0x4f, 0x09,
	0x28, 0x9e, 0x3c, 0x59, 0x35, 0x20, 0xa8, 0x0d, 0x25, 0x7a, 0x5d, 0xda, 0x30, 0x1b, 0x16, 0x9b,
	0x1d, 0x99, 0xdd, 0xd8, 0x9f, 0xe8, 0x5f, 0xf2, 0x28, 0x26, 0x56, 0x93, 0xf3, 0x7b, 0xc3, 0x1b,
	0x3e, 0x48, 0x3c, 0x5a, 0x47, 0xac, 0xb7, 0xb4, 0x4b, 0x4b, 0x62, 0x4c, 0xdf, 0x99, 0x3c, 0x6d,
	0x1a, 0x9d, 0x7e, 0x20, 0x1b, 0x6d, 0x90, 0x55, 0x49, 0x56, 0x9b, 0x2a, 0x69, 0x81, 0x84, 0x7f,
	0x7f, 0xf6, 0x29, 0x60, 0xf2, 0xfa, 0x6b, 0xdd, 0xb6, 0x92, 0xbc, 0x86, 0xd3, 0xee, 0xae, 0x5c,
	0x7b, 0x43, 0x56, 0x79, 0x53, 0x23, 0x35, 0x5e, 0x19, 0xab, 0x6a, 0x17, 0x8b, 0xa9, 0x98, 0x8f,
	0x56, 0x27, 0x7d, 0xa3, 0xe8, 0x84, 0x07, 0xfb, 0xe4, 0x64, 0x0e, 0xd2, 0x79, 0x6e, 0x4a, 0xdf,
	0x30, 0xaa, 0x7d, 0x3e, 0x0e, 0xa6, 0x62, 0x3e, 0xb9, 0x38, 0xeb, 0x3d, 0x9a, 0x0c, 0xa3, 0x49,
	0x41, 0x55, 0xb5, 0xc5, 0xd5, 0xd1, 0xdf, 0xf1, 0x9e, 0xcf, 0xce, 0x21, 0xec, 0xa0, 0x3c, 0x84,
	0xf1, 0xdd, 0x7d, 0x76, 0xf3, 0xf2, 0x58, 0x44, 0x07, 0x32, 0x84, 0x60, 0xf9, 0x1c, 0x09, 0x39,
	0x86, 0xd1, 0x32, 0xcb, 0xa2, 0x60, 0x71, 0x05, 0x31, 0x71, 0xd5, 0x4f, 0x68, 0x5e, 0xd7, 0xb8,
	0x23, 0x7e, 0x5b, 0x1c, 0x0f, 0x6b, 0xf9, 0xcf, 0x0c, 0x2e, 0x17, 0x5f, 0x42, 0x6c, 0xc2, 0x76,
	0x93, 0xcb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x96, 0xd9, 0x05, 0x45, 0x01, 0x00, 0x00,
}
