// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/autotuning.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type AutotuneResult_FailureKind int32

const (
	AutotuneResult_UNKNOWN          AutotuneResult_FailureKind = 0
	AutotuneResult_REDZONE_MODIFIED AutotuneResult_FailureKind = 1
	AutotuneResult_WRONG_RESULT     AutotuneResult_FailureKind = 2
)

var AutotuneResult_FailureKind_name = map[int32]string{
	0: "UNKNOWN",
	1: "REDZONE_MODIFIED",
	2: "WRONG_RESULT",
}

var AutotuneResult_FailureKind_value = map[string]int32{
	"UNKNOWN":          0,
	"REDZONE_MODIFIED": 1,
	"WRONG_RESULT":     2,
}

func (x AutotuneResult_FailureKind) String() string {
	return proto.EnumName(AutotuneResult_FailureKind_name, int32(x))
}

func (AutotuneResult_FailureKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{2, 0}
}

type CudnnVersion struct {
	Major                int32    `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int32    `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                int32    `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CudnnVersion) Reset()         { *m = CudnnVersion{} }
func (m *CudnnVersion) String() string { return proto.CompactTextString(m) }
func (*CudnnVersion) ProtoMessage()    {}
func (*CudnnVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{0}
}

func (m *CudnnVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CudnnVersion.Unmarshal(m, b)
}
func (m *CudnnVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CudnnVersion.Marshal(b, m, deterministic)
}
func (m *CudnnVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CudnnVersion.Merge(m, src)
}
func (m *CudnnVersion) XXX_Size() int {
	return xxx_messageInfo_CudnnVersion.Size(m)
}
func (m *CudnnVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CudnnVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CudnnVersion proto.InternalMessageInfo

func (m *CudnnVersion) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *CudnnVersion) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *CudnnVersion) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

type ComputeCapability struct {
	Major                int32    `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int32    `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeCapability) Reset()         { *m = ComputeCapability{} }
func (m *ComputeCapability) String() string { return proto.CompactTextString(m) }
func (*ComputeCapability) ProtoMessage()    {}
func (*ComputeCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{1}
}

func (m *ComputeCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeCapability.Unmarshal(m, b)
}
func (m *ComputeCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeCapability.Marshal(b, m, deterministic)
}
func (m *ComputeCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeCapability.Merge(m, src)
}
func (m *ComputeCapability) XXX_Size() int {
	return xxx_messageInfo_ComputeCapability.Size(m)
}
func (m *ComputeCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeCapability.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeCapability proto.InternalMessageInfo

func (m *ComputeCapability) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *ComputeCapability) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

type AutotuneResult struct {
	ScratchBytes int64                         `protobuf:"varint,8,opt,name=scratch_bytes,json=scratchBytes,proto3" json:"scratch_bytes,omitempty"`
	RunTime      *duration.Duration            `protobuf:"bytes,9,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	Failure      *AutotuneResult_FailureResult `protobuf:"bytes,7,opt,name=failure,proto3" json:"failure,omitempty"`
	// Types that are valid to be assigned to Key:
	//	*AutotuneResult_Conv
	//	*AutotuneResult_Gemm
	Key                  isAutotuneResult_Key `protobuf_oneof:"key"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AutotuneResult) Reset()         { *m = AutotuneResult{} }
func (m *AutotuneResult) String() string { return proto.CompactTextString(m) }
func (*AutotuneResult) ProtoMessage()    {}
func (*AutotuneResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{2}
}

func (m *AutotuneResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutotuneResult.Unmarshal(m, b)
}
func (m *AutotuneResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutotuneResult.Marshal(b, m, deterministic)
}
func (m *AutotuneResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutotuneResult.Merge(m, src)
}
func (m *AutotuneResult) XXX_Size() int {
	return xxx_messageInfo_AutotuneResult.Size(m)
}
func (m *AutotuneResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AutotuneResult.DiscardUnknown(m)
}

var xxx_messageInfo_AutotuneResult proto.InternalMessageInfo

func (m *AutotuneResult) GetScratchBytes() int64 {
	if m != nil {
		return m.ScratchBytes
	}
	return 0
}

func (m *AutotuneResult) GetRunTime() *duration.Duration {
	if m != nil {
		return m.RunTime
	}
	return nil
}

func (m *AutotuneResult) GetFailure() *AutotuneResult_FailureResult {
	if m != nil {
		return m.Failure
	}
	return nil
}

type isAutotuneResult_Key interface {
	isAutotuneResult_Key()
}

type AutotuneResult_Conv struct {
	Conv *AutotuneResult_ConvKey `protobuf:"bytes,5,opt,name=conv,proto3,oneof"`
}

type AutotuneResult_Gemm struct {
	Gemm *AutotuneResult_GemmKey `protobuf:"bytes,6,opt,name=gemm,proto3,oneof"`
}

func (*AutotuneResult_Conv) isAutotuneResult_Key() {}

func (*AutotuneResult_Gemm) isAutotuneResult_Key() {}

func (m *AutotuneResult) GetKey() isAutotuneResult_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AutotuneResult) GetConv() *AutotuneResult_ConvKey {
	if x, ok := m.GetKey().(*AutotuneResult_Conv); ok {
		return x.Conv
	}
	return nil
}

func (m *AutotuneResult) GetGemm() *AutotuneResult_GemmKey {
	if x, ok := m.GetKey().(*AutotuneResult_Gemm); ok {
		return x.Gemm
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AutotuneResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AutotuneResult_Conv)(nil),
		(*AutotuneResult_Gemm)(nil),
	}
}

type AutotuneResult_FailureResult struct {
	Kind AutotuneResult_FailureKind `protobuf:"varint,1,opt,name=kind,proto3,enum=tensorflow.AutotuneResult_FailureKind" json:"kind,omitempty"`
	Msg  string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// For failure_kind == WRONG_RESULT, this field indicates the reference
	// configuration that we compared against.
	//
	// Note that the reference algorithm isn't always correct.  However,
	// empirically it's more correct, as it's "algo 0", less fancy than the
	// compared one.
	//
	// Types that are valid to be assigned to Key:
	//	*AutotuneResult_FailureResult_ReferenceConv
	//	*AutotuneResult_FailureResult_ReferenceGemm
	Key                  isAutotuneResult_FailureResult_Key `protobuf_oneof:"key"`
	BufferAddress        int64                              `protobuf:"varint,13,opt,name=buffer_address,json=bufferAddress,proto3" json:"buffer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *AutotuneResult_FailureResult) Reset()         { *m = AutotuneResult_FailureResult{} }
func (m *AutotuneResult_FailureResult) String() string { return proto.CompactTextString(m) }
func (*AutotuneResult_FailureResult) ProtoMessage()    {}
func (*AutotuneResult_FailureResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{2, 0}
}

func (m *AutotuneResult_FailureResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutotuneResult_FailureResult.Unmarshal(m, b)
}
func (m *AutotuneResult_FailureResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutotuneResult_FailureResult.Marshal(b, m, deterministic)
}
func (m *AutotuneResult_FailureResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutotuneResult_FailureResult.Merge(m, src)
}
func (m *AutotuneResult_FailureResult) XXX_Size() int {
	return xxx_messageInfo_AutotuneResult_FailureResult.Size(m)
}
func (m *AutotuneResult_FailureResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AutotuneResult_FailureResult.DiscardUnknown(m)
}

var xxx_messageInfo_AutotuneResult_FailureResult proto.InternalMessageInfo

func (m *AutotuneResult_FailureResult) GetKind() AutotuneResult_FailureKind {
	if m != nil {
		return m.Kind
	}
	return AutotuneResult_UNKNOWN
}

func (m *AutotuneResult_FailureResult) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type isAutotuneResult_FailureResult_Key interface {
	isAutotuneResult_FailureResult_Key()
}

type AutotuneResult_FailureResult_ReferenceConv struct {
	ReferenceConv *AutotuneResult_ConvKey `protobuf:"bytes,11,opt,name=reference_conv,json=referenceConv,proto3,oneof"`
}

type AutotuneResult_FailureResult_ReferenceGemm struct {
	ReferenceGemm *AutotuneResult_GemmKey `protobuf:"bytes,12,opt,name=reference_gemm,json=referenceGemm,proto3,oneof"`
}

func (*AutotuneResult_FailureResult_ReferenceConv) isAutotuneResult_FailureResult_Key() {}

func (*AutotuneResult_FailureResult_ReferenceGemm) isAutotuneResult_FailureResult_Key() {}

func (m *AutotuneResult_FailureResult) GetKey() isAutotuneResult_FailureResult_Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AutotuneResult_FailureResult) GetReferenceConv() *AutotuneResult_ConvKey {
	if x, ok := m.GetKey().(*AutotuneResult_FailureResult_ReferenceConv); ok {
		return x.ReferenceConv
	}
	return nil
}

func (m *AutotuneResult_FailureResult) GetReferenceGemm() *AutotuneResult_GemmKey {
	if x, ok := m.GetKey().(*AutotuneResult_FailureResult_ReferenceGemm); ok {
		return x.ReferenceGemm
	}
	return nil
}

func (m *AutotuneResult_FailureResult) GetBufferAddress() int64 {
	if m != nil {
		return m.BufferAddress
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AutotuneResult_FailureResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AutotuneResult_FailureResult_ReferenceConv)(nil),
		(*AutotuneResult_FailureResult_ReferenceGemm)(nil),
	}
}

type AutotuneResult_ConvKey struct {
	Algorithm            int64    `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	TensorOpsEnabled     bool     `protobuf:"varint,2,opt,name=tensor_ops_enabled,json=tensorOpsEnabled,proto3" json:"tensor_ops_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutotuneResult_ConvKey) Reset()         { *m = AutotuneResult_ConvKey{} }
func (m *AutotuneResult_ConvKey) String() string { return proto.CompactTextString(m) }
func (*AutotuneResult_ConvKey) ProtoMessage()    {}
func (*AutotuneResult_ConvKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{2, 1}
}

func (m *AutotuneResult_ConvKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutotuneResult_ConvKey.Unmarshal(m, b)
}
func (m *AutotuneResult_ConvKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutotuneResult_ConvKey.Marshal(b, m, deterministic)
}
func (m *AutotuneResult_ConvKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutotuneResult_ConvKey.Merge(m, src)
}
func (m *AutotuneResult_ConvKey) XXX_Size() int {
	return xxx_messageInfo_AutotuneResult_ConvKey.Size(m)
}
func (m *AutotuneResult_ConvKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AutotuneResult_ConvKey.DiscardUnknown(m)
}

var xxx_messageInfo_AutotuneResult_ConvKey proto.InternalMessageInfo

func (m *AutotuneResult_ConvKey) GetAlgorithm() int64 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

func (m *AutotuneResult_ConvKey) GetTensorOpsEnabled() bool {
	if m != nil {
		return m.TensorOpsEnabled
	}
	return false
}

type AutotuneResult_GemmKey struct {
	Algorithm            int64    `protobuf:"varint,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutotuneResult_GemmKey) Reset()         { *m = AutotuneResult_GemmKey{} }
func (m *AutotuneResult_GemmKey) String() string { return proto.CompactTextString(m) }
func (*AutotuneResult_GemmKey) ProtoMessage()    {}
func (*AutotuneResult_GemmKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{2, 2}
}

func (m *AutotuneResult_GemmKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutotuneResult_GemmKey.Unmarshal(m, b)
}
func (m *AutotuneResult_GemmKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutotuneResult_GemmKey.Marshal(b, m, deterministic)
}
func (m *AutotuneResult_GemmKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutotuneResult_GemmKey.Merge(m, src)
}
func (m *AutotuneResult_GemmKey) XXX_Size() int {
	return xxx_messageInfo_AutotuneResult_GemmKey.Size(m)
}
func (m *AutotuneResult_GemmKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AutotuneResult_GemmKey.DiscardUnknown(m)
}

var xxx_messageInfo_AutotuneResult_GemmKey proto.InternalMessageInfo

func (m *AutotuneResult_GemmKey) GetAlgorithm() int64 {
	if m != nil {
		return m.Algorithm
	}
	return 0
}

type AutotuningLog struct {
	Instr *any.Any `protobuf:"bytes,1,opt,name=instr,proto3" json:"instr,omitempty"`
	// Records all auto-tuning results per algorithm.
	Results           []*AutotuneResult  `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	CudnnVersion      *CudnnVersion      `protobuf:"bytes,3,opt,name=cudnn_version,json=cudnnVersion,proto3" json:"cudnn_version,omitempty"`
	ComputeCapability *ComputeCapability `protobuf:"bytes,4,opt,name=compute_capability,json=computeCapability,proto3" json:"compute_capability,omitempty"`
	// stream_executor::DeviceDescription::pci_bus_id.
	DevicePciBusId       string   `protobuf:"bytes,5,opt,name=device_pci_bus_id,json=devicePciBusId,proto3" json:"device_pci_bus_id,omitempty"`
	BlasVersion          string   `protobuf:"bytes,6,opt,name=blas_version,json=blasVersion,proto3" json:"blas_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutotuningLog) Reset()         { *m = AutotuningLog{} }
func (m *AutotuningLog) String() string { return proto.CompactTextString(m) }
func (*AutotuningLog) ProtoMessage()    {}
func (*AutotuningLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_f61248520e180396, []int{3}
}

func (m *AutotuningLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutotuningLog.Unmarshal(m, b)
}
func (m *AutotuningLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutotuningLog.Marshal(b, m, deterministic)
}
func (m *AutotuningLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutotuningLog.Merge(m, src)
}
func (m *AutotuningLog) XXX_Size() int {
	return xxx_messageInfo_AutotuningLog.Size(m)
}
func (m *AutotuningLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AutotuningLog.DiscardUnknown(m)
}

var xxx_messageInfo_AutotuningLog proto.InternalMessageInfo

func (m *AutotuningLog) GetInstr() *any.Any {
	if m != nil {
		return m.Instr
	}
	return nil
}

func (m *AutotuningLog) GetResults() []*AutotuneResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *AutotuningLog) GetCudnnVersion() *CudnnVersion {
	if m != nil {
		return m.CudnnVersion
	}
	return nil
}

func (m *AutotuningLog) GetComputeCapability() *ComputeCapability {
	if m != nil {
		return m.ComputeCapability
	}
	return nil
}

func (m *AutotuningLog) GetDevicePciBusId() string {
	if m != nil {
		return m.DevicePciBusId
	}
	return ""
}

func (m *AutotuningLog) GetBlasVersion() string {
	if m != nil {
		return m.BlasVersion
	}
	return ""
}

func init() {
	proto.RegisterEnum("tensorflow.AutotuneResult_FailureKind", AutotuneResult_FailureKind_name, AutotuneResult_FailureKind_value)
	proto.RegisterType((*CudnnVersion)(nil), "tensorflow.CudnnVersion")
	proto.RegisterType((*ComputeCapability)(nil), "tensorflow.ComputeCapability")
	proto.RegisterType((*AutotuneResult)(nil), "tensorflow.AutotuneResult")
	proto.RegisterType((*AutotuneResult_FailureResult)(nil), "tensorflow.AutotuneResult.FailureResult")
	proto.RegisterType((*AutotuneResult_ConvKey)(nil), "tensorflow.AutotuneResult.ConvKey")
	proto.RegisterType((*AutotuneResult_GemmKey)(nil), "tensorflow.AutotuneResult.GemmKey")
	proto.RegisterType((*AutotuningLog)(nil), "tensorflow.AutotuningLog")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/autotuning.proto", fileDescriptor_f61248520e180396)
}

var fileDescriptor_f61248520e180396 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0x1a, 0x3d,
	0x18, 0x0c, 0x10, 0x20, 0x7c, 0x0b, 0x88, 0x58, 0xb9, 0xd8, 0xa0, 0xff, 0xaf, 0x28, 0x55, 0x5b,
	0x52, 0x55, 0x44, 0xa2, 0xb9, 0xa8, 0x2a, 0x55, 0x15, 0xa7, 0xa4, 0x11, 0x29, 0x44, 0x6e, 0xd2,
	0x48, 0xbd, 0x59, 0xed, 0xc1, 0x10, 0x37, 0xbb, 0xf6, 0xca, 0xde, 0xa5, 0xe2, 0x45, 0xfa, 0x06,
	0x7d, 0x9b, 0x3e, 0x54, 0xb5, 0xf6, 0x86, 0x43, 0xa2, 0x26, 0xca, 0x1d, 0x1e, 0x7f, 0x33, 0xfa,
	0x66, 0x58, 0x0f, 0x1c, 0x44, 0x84, 0x49, 0x2e, 0xa6, 0x3e, 0xff, 0x79, 0xe8, 0x72, 0x41, 0x0e,
	0x43, 0xc1, 0x23, 0xee, 0xc4, 0xd3, 0x43, 0x3b, 0x8e, 0x78, 0x14, 0x33, 0xca, 0x66, 0x6d, 0x85,
	0x21, 0x58, 0x8d, 0xd6, 0xf7, 0x67, 0x9c, 0xcf, 0xfc, 0xf5, 0x69, 0xb6, 0xd0, 0x63, 0xf5, 0x67,
	0x77, 0xaf, 0xbc, 0x58, 0xd8, 0x11, 0xe5, 0x4c, 0xdf, 0x37, 0xcf, 0xa1, 0xdc, 0x8f, 0x3d, 0xc6,
	0xbe, 0x11, 0x21, 0x29, 0x67, 0x68, 0x0f, 0xf2, 0x81, 0xfd, 0x83, 0x0b, 0x33, 0xd3, 0xc8, 0xb4,
	0xf2, 0x58, 0x1f, 0x14, 0x4a, 0x19, 0x17, 0x66, 0x36, 0x45, 0x93, 0x43, 0x82, 0x86, 0x76, 0xe4,
	0x5e, 0x9b, 0x39, 0x8d, 0xaa, 0x43, 0xf3, 0x13, 0xec, 0xf6, 0x79, 0x10, 0xc6, 0x11, 0xe9, 0xdb,
	0xa1, 0xed, 0x50, 0x9f, 0x46, 0x8b, 0xa7, 0xc8, 0x36, 0x7f, 0x15, 0xa0, 0xda, 0xd5, 0x76, 0x09,
	0x26, 0x32, 0xf6, 0x23, 0xf4, 0x02, 0x2a, 0xd2, 0x15, 0x89, 0xbc, 0xe5, 0x2c, 0x22, 0x22, 0xcd,
	0x9d, 0x46, 0xa6, 0x95, 0xc3, 0xe5, 0x14, 0xec, 0x25, 0x18, 0x3a, 0x82, 0x1d, 0x11, 0x33, 0x2b,
	0xa2, 0x01, 0x31, 0x4b, 0x8d, 0x4c, 0xcb, 0xe8, 0xec, 0xb7, 0xb5, 0xfb, 0xf6, 0xad, 0xfb, 0xf6,
	0x20, 0x75, 0x8f, 0x8b, 0x22, 0x66, 0x17, 0x34, 0x20, 0xa8, 0x07, 0xc5, 0xa9, 0x4d, 0xfd, 0x58,
	0x10, 0xb3, 0xa8, 0x48, 0xad, 0xf6, 0x2a, 0xd9, 0xf6, 0xe6, 0x1e, 0xed, 0x63, 0x3d, 0xa9, 0x4f,
	0xf8, 0x96, 0x88, 0xde, 0xc3, 0xb6, 0xcb, 0xd9, 0xdc, 0xcc, 0x2b, 0x81, 0xe6, 0x03, 0x02, 0x7d,
	0xce, 0xe6, 0x23, 0xb2, 0xf8, 0xbc, 0x85, 0x15, 0x23, 0x61, 0xce, 0x48, 0x10, 0x98, 0x85, 0x47,
	0x99, 0x27, 0x24, 0x08, 0x52, 0x66, 0xc2, 0xa8, 0xff, 0xce, 0x42, 0x65, 0x63, 0x1d, 0xf4, 0x01,
	0xb6, 0x6f, 0x28, 0xf3, 0x54, 0xc4, 0xd5, 0xce, 0xab, 0xc7, 0x6d, 0x8c, 0x28, 0xf3, 0xb0, 0xe2,
	0xa0, 0x1a, 0xe4, 0x02, 0x39, 0x53, 0xff, 0x43, 0x09, 0x27, 0x3f, 0xd1, 0x08, 0xaa, 0x82, 0x4c,
	0x89, 0x20, 0xcc, 0x25, 0x96, 0x72, 0x67, 0x3c, 0xc1, 0x5d, 0x65, 0xc9, 0x4d, 0xb0, 0x4d, 0x31,
	0x65, 0xb8, 0xfc, 0x04, 0xc3, 0x2b, 0xb1, 0x04, 0x43, 0x2f, 0xa1, 0xea, 0xc4, 0xd3, 0x29, 0x11,
	0x96, 0xed, 0x79, 0x82, 0x48, 0x69, 0x56, 0xd4, 0xd7, 0x50, 0xd1, 0x68, 0x57, 0x83, 0xbd, 0x3c,
	0xe4, 0x6e, 0xc8, 0xa2, 0x7e, 0x09, 0xc5, 0x74, 0x2d, 0xf4, 0x1f, 0x94, 0x6c, 0x7f, 0xc6, 0x05,
	0x8d, 0xae, 0x03, 0x95, 0x52, 0x0e, 0xaf, 0x00, 0xf4, 0x16, 0x90, 0x5e, 0xc6, 0xe2, 0xa1, 0xb4,
	0x08, 0xb3, 0x1d, 0x9f, 0x78, 0x2a, 0x91, 0x1d, 0x5c, 0xd3, 0x37, 0x93, 0x50, 0x0e, 0x35, 0x5e,
	0x7f, 0x0d, 0xc5, 0x74, 0xc1, 0x87, 0x65, 0x9b, 0x3d, 0x30, 0xd6, 0xe2, 0x46, 0x06, 0x14, 0x2f,
	0xc7, 0xa3, 0xf1, 0xe4, 0x6a, 0x5c, 0xdb, 0x42, 0x7b, 0x50, 0xc3, 0xc3, 0xc1, 0xf7, 0xc9, 0x78,
	0x68, 0x7d, 0x99, 0x0c, 0x4e, 0x8f, 0x4f, 0x87, 0x83, 0x5a, 0x06, 0xd5, 0xa0, 0x7c, 0x85, 0x27,
	0xe3, 0x13, 0x0b, 0x0f, 0xbf, 0x5e, 0x9e, 0x5d, 0xd4, 0xb2, 0xa9, 0x95, 0xe6, 0x9f, 0x2c, 0x54,
	0xba, 0xcb, 0x1e, 0x38, 0xe3, 0x33, 0xf4, 0x06, 0xf2, 0x94, 0xc9, 0x48, 0x3f, 0x2b, 0xa3, 0xb3,
	0x77, 0xef, 0x7b, 0xef, 0xb2, 0x05, 0xd6, 0x23, 0xe8, 0x08, 0x8a, 0x42, 0x25, 0x2b, 0xcd, 0x6c,
	0x23, 0xd7, 0x32, 0x3a, 0xf5, 0x7f, 0x87, 0x8f, 0x6f, 0x47, 0xd1, 0x47, 0xa8, 0xb8, 0x49, 0x3f,
	0x58, 0x73, 0x5d, 0x10, 0xea, 0xad, 0x1b, 0x1d, 0x73, 0x9d, 0xbb, 0x5e, 0x20, 0xb8, 0xec, 0xae,
	0xd7, 0xc9, 0x19, 0x20, 0x57, 0x97, 0x81, 0xe5, 0x2e, 0xdb, 0xc0, 0xdc, 0x56, 0x1a, 0xff, 0x6f,
	0x68, 0xdc, 0xad, 0x0c, 0xbc, 0xeb, 0xde, 0x6b, 0x91, 0x03, 0xd8, 0xf5, 0xc8, 0x9c, 0xba, 0xc4,
	0x0a, 0x5d, 0x6a, 0x39, 0xb1, 0xb4, 0xa8, 0xa7, 0x1e, 0x5d, 0x09, 0x57, 0xf5, 0xc5, 0xb9, 0x4b,
	0x7b, 0xb1, 0x3c, 0xf5, 0xd0, 0x73, 0x28, 0x3b, 0xbe, 0x2d, 0x97, 0x6b, 0x17, 0xd4, 0x94, 0x91,
	0x60, 0xe9, 0x6e, 0x4e, 0x41, 0xa5, 0xf4, 0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x70,
	0x12, 0xaa, 0x75, 0x05, 0x00, 0x00,
}
