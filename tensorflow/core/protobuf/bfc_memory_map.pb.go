// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/bfc_memory_map.proto

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

// Some of the data from AllocatorStats
type MemAllocatorStats struct {
	NumAllocs            int64    `protobuf:"varint,1,opt,name=num_allocs,json=numAllocs,proto3" json:"num_allocs,omitempty"`
	BytesInUse           int64    `protobuf:"varint,2,opt,name=bytes_in_use,json=bytesInUse,proto3" json:"bytes_in_use,omitempty"`
	PeakBytesInUse       int64    `protobuf:"varint,3,opt,name=peak_bytes_in_use,json=peakBytesInUse,proto3" json:"peak_bytes_in_use,omitempty"`
	LargestAllocSize     int64    `protobuf:"varint,4,opt,name=largest_alloc_size,json=largestAllocSize,proto3" json:"largest_alloc_size,omitempty"`
	FragmentationMetric  float32  `protobuf:"fixed32,5,opt,name=fragmentation_metric,json=fragmentationMetric,proto3" json:"fragmentation_metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemAllocatorStats) Reset()         { *m = MemAllocatorStats{} }
func (m *MemAllocatorStats) String() string { return proto.CompactTextString(m) }
func (*MemAllocatorStats) ProtoMessage()    {}
func (*MemAllocatorStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{0}
}

func (m *MemAllocatorStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemAllocatorStats.Unmarshal(m, b)
}
func (m *MemAllocatorStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemAllocatorStats.Marshal(b, m, deterministic)
}
func (m *MemAllocatorStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemAllocatorStats.Merge(m, src)
}
func (m *MemAllocatorStats) XXX_Size() int {
	return xxx_messageInfo_MemAllocatorStats.Size(m)
}
func (m *MemAllocatorStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MemAllocatorStats.DiscardUnknown(m)
}

var xxx_messageInfo_MemAllocatorStats proto.InternalMessageInfo

func (m *MemAllocatorStats) GetNumAllocs() int64 {
	if m != nil {
		return m.NumAllocs
	}
	return 0
}

func (m *MemAllocatorStats) GetBytesInUse() int64 {
	if m != nil {
		return m.BytesInUse
	}
	return 0
}

func (m *MemAllocatorStats) GetPeakBytesInUse() int64 {
	if m != nil {
		return m.PeakBytesInUse
	}
	return 0
}

func (m *MemAllocatorStats) GetLargestAllocSize() int64 {
	if m != nil {
		return m.LargestAllocSize
	}
	return 0
}

func (m *MemAllocatorStats) GetFragmentationMetric() float32 {
	if m != nil {
		return m.FragmentationMetric
	}
	return 0
}

type MemChunk struct {
	Address              uint64   `protobuf:"varint,1,opt,name=address,proto3" json:"address,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	RequestedSize        int64    `protobuf:"varint,3,opt,name=requested_size,json=requestedSize,proto3" json:"requested_size,omitempty"`
	Bin                  int32    `protobuf:"varint,4,opt,name=bin,proto3" json:"bin,omitempty"`
	OpName               string   `protobuf:"bytes,5,opt,name=op_name,json=opName,proto3" json:"op_name,omitempty"`
	FreedAtCount         uint64   `protobuf:"varint,6,opt,name=freed_at_count,json=freedAtCount,proto3" json:"freed_at_count,omitempty"`
	ActionCount          uint64   `protobuf:"varint,7,opt,name=action_count,json=actionCount,proto3" json:"action_count,omitempty"`
	InUse                bool     `protobuf:"varint,8,opt,name=in_use,json=inUse,proto3" json:"in_use,omitempty"`
	StepId               uint64   `protobuf:"varint,9,opt,name=step_id,json=stepId,proto3" json:"step_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemChunk) Reset()         { *m = MemChunk{} }
func (m *MemChunk) String() string { return proto.CompactTextString(m) }
func (*MemChunk) ProtoMessage()    {}
func (*MemChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{1}
}

func (m *MemChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemChunk.Unmarshal(m, b)
}
func (m *MemChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemChunk.Marshal(b, m, deterministic)
}
func (m *MemChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemChunk.Merge(m, src)
}
func (m *MemChunk) XXX_Size() int {
	return xxx_messageInfo_MemChunk.Size(m)
}
func (m *MemChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_MemChunk.DiscardUnknown(m)
}

var xxx_messageInfo_MemChunk proto.InternalMessageInfo

func (m *MemChunk) GetAddress() uint64 {
	if m != nil {
		return m.Address
	}
	return 0
}

func (m *MemChunk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MemChunk) GetRequestedSize() int64 {
	if m != nil {
		return m.RequestedSize
	}
	return 0
}

func (m *MemChunk) GetBin() int32 {
	if m != nil {
		return m.Bin
	}
	return 0
}

func (m *MemChunk) GetOpName() string {
	if m != nil {
		return m.OpName
	}
	return ""
}

func (m *MemChunk) GetFreedAtCount() uint64 {
	if m != nil {
		return m.FreedAtCount
	}
	return 0
}

func (m *MemChunk) GetActionCount() uint64 {
	if m != nil {
		return m.ActionCount
	}
	return 0
}

func (m *MemChunk) GetInUse() bool {
	if m != nil {
		return m.InUse
	}
	return false
}

func (m *MemChunk) GetStepId() uint64 {
	if m != nil {
		return m.StepId
	}
	return 0
}

type BinSummary struct {
	Bin                  int32    `protobuf:"varint,1,opt,name=bin,proto3" json:"bin,omitempty"`
	TotalBytesInUse      int64    `protobuf:"varint,2,opt,name=total_bytes_in_use,json=totalBytesInUse,proto3" json:"total_bytes_in_use,omitempty"`
	TotalBytesInBin      int64    `protobuf:"varint,3,opt,name=total_bytes_in_bin,json=totalBytesInBin,proto3" json:"total_bytes_in_bin,omitempty"`
	TotalChunksInUse     int64    `protobuf:"varint,4,opt,name=total_chunks_in_use,json=totalChunksInUse,proto3" json:"total_chunks_in_use,omitempty"`
	TotalChunksInBin     int64    `protobuf:"varint,5,opt,name=total_chunks_in_bin,json=totalChunksInBin,proto3" json:"total_chunks_in_bin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BinSummary) Reset()         { *m = BinSummary{} }
func (m *BinSummary) String() string { return proto.CompactTextString(m) }
func (*BinSummary) ProtoMessage()    {}
func (*BinSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{2}
}

func (m *BinSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BinSummary.Unmarshal(m, b)
}
func (m *BinSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BinSummary.Marshal(b, m, deterministic)
}
func (m *BinSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BinSummary.Merge(m, src)
}
func (m *BinSummary) XXX_Size() int {
	return xxx_messageInfo_BinSummary.Size(m)
}
func (m *BinSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BinSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BinSummary proto.InternalMessageInfo

func (m *BinSummary) GetBin() int32 {
	if m != nil {
		return m.Bin
	}
	return 0
}

func (m *BinSummary) GetTotalBytesInUse() int64 {
	if m != nil {
		return m.TotalBytesInUse
	}
	return 0
}

func (m *BinSummary) GetTotalBytesInBin() int64 {
	if m != nil {
		return m.TotalBytesInBin
	}
	return 0
}

func (m *BinSummary) GetTotalChunksInUse() int64 {
	if m != nil {
		return m.TotalChunksInUse
	}
	return 0
}

func (m *BinSummary) GetTotalChunksInBin() int64 {
	if m != nil {
		return m.TotalChunksInBin
	}
	return 0
}

type SnapShot struct {
	ActionCount          uint64   `protobuf:"varint,1,opt,name=action_count,json=actionCount,proto3" json:"action_count,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnapShot) Reset()         { *m = SnapShot{} }
func (m *SnapShot) String() string { return proto.CompactTextString(m) }
func (*SnapShot) ProtoMessage()    {}
func (*SnapShot) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{3}
}

func (m *SnapShot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapShot.Unmarshal(m, b)
}
func (m *SnapShot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapShot.Marshal(b, m, deterministic)
}
func (m *SnapShot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapShot.Merge(m, src)
}
func (m *SnapShot) XXX_Size() int {
	return xxx_messageInfo_SnapShot.Size(m)
}
func (m *SnapShot) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapShot.DiscardUnknown(m)
}

var xxx_messageInfo_SnapShot proto.InternalMessageInfo

func (m *SnapShot) GetActionCount() uint64 {
	if m != nil {
		return m.ActionCount
	}
	return 0
}

func (m *SnapShot) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MemoryDump struct {
	AllocatorName        string             `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	BinSummary           []*BinSummary      `protobuf:"bytes,2,rep,name=bin_summary,json=binSummary,proto3" json:"bin_summary,omitempty"`
	Chunk                []*MemChunk        `protobuf:"bytes,3,rep,name=chunk,proto3" json:"chunk,omitempty"`
	SnapShot             []*SnapShot        `protobuf:"bytes,4,rep,name=snap_shot,json=snapShot,proto3" json:"snap_shot,omitempty"`
	Stats                *MemAllocatorStats `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MemoryDump) Reset()         { *m = MemoryDump{} }
func (m *MemoryDump) String() string { return proto.CompactTextString(m) }
func (*MemoryDump) ProtoMessage()    {}
func (*MemoryDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdf22777007c1f3d, []int{4}
}

func (m *MemoryDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryDump.Unmarshal(m, b)
}
func (m *MemoryDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryDump.Marshal(b, m, deterministic)
}
func (m *MemoryDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryDump.Merge(m, src)
}
func (m *MemoryDump) XXX_Size() int {
	return xxx_messageInfo_MemoryDump.Size(m)
}
func (m *MemoryDump) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryDump.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryDump proto.InternalMessageInfo

func (m *MemoryDump) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *MemoryDump) GetBinSummary() []*BinSummary {
	if m != nil {
		return m.BinSummary
	}
	return nil
}

func (m *MemoryDump) GetChunk() []*MemChunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

func (m *MemoryDump) GetSnapShot() []*SnapShot {
	if m != nil {
		return m.SnapShot
	}
	return nil
}

func (m *MemoryDump) GetStats() *MemAllocatorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*MemAllocatorStats)(nil), "tensorflow.MemAllocatorStats")
	proto.RegisterType((*MemChunk)(nil), "tensorflow.MemChunk")
	proto.RegisterType((*BinSummary)(nil), "tensorflow.BinSummary")
	proto.RegisterType((*SnapShot)(nil), "tensorflow.SnapShot")
	proto.RegisterType((*MemoryDump)(nil), "tensorflow.MemoryDump")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/bfc_memory_map.proto", fileDescriptor_fdf22777007c1f3d)
}

var fileDescriptor_fdf22777007c1f3d = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x3d, 0x6f, 0xdb, 0x30,
	0x10, 0x86, 0x21, 0xcb, 0xf2, 0xc7, 0x39, 0x49, 0x13, 0x26, 0x6d, 0xb4, 0x04, 0x70, 0x8d, 0x16,
	0x70, 0x3f, 0x92, 0x20, 0xc9, 0xd0, 0x39, 0x4e, 0x97, 0x0c, 0xee, 0x20, 0xa3, 0x33, 0x41, 0x49,
	0x74, 0x42, 0xc4, 0x24, 0x55, 0x92, 0x42, 0xe1, 0x4c, 0x05, 0xfa, 0x1f, 0x3b, 0xf7, 0xa7, 0x14,
	0x3a, 0x5a, 0x56, 0x5c, 0x7b, 0x13, 0xef, 0x1e, 0x1e, 0xef, 0xde, 0xd7, 0x67, 0x38, 0x77, 0x5c,
	0x59, 0x6d, 0xe6, 0x0b, 0xfd, 0xf3, 0x32, 0xd3, 0x86, 0x5f, 0x16, 0x46, 0x3b, 0x9d, 0x96, 0xf3,
	0xcb, 0x74, 0x9e, 0x51, 0xc9, 0xa5, 0x36, 0x4b, 0x2a, 0x59, 0x71, 0x81, 0x71, 0x02, 0x0d, 0x3e,
	0xfa, 0x1b, 0xc0, 0xd1, 0x94, 0xcb, 0xdb, 0xc5, 0x42, 0x67, 0xcc, 0x69, 0x33, 0x73, 0xcc, 0x59,
	0x72, 0x06, 0xa0, 0x4a, 0x49, 0x59, 0x15, 0xb5, 0x71, 0x30, 0x0c, 0xc6, 0x61, 0xd2, 0x57, 0xa5,
	0xc7, 0x2c, 0x19, 0xc2, 0x5e, 0xba, 0x74, 0xdc, 0x52, 0xa1, 0x68, 0x69, 0x79, 0xdc, 0x42, 0x00,
	0x30, 0x76, 0xaf, 0xbe, 0x5b, 0x4e, 0x3e, 0xc0, 0x51, 0xc1, 0xd9, 0x13, 0xdd, 0xc0, 0x42, 0xc4,
	0x0e, 0xaa, 0xc4, 0xa4, 0x41, 0x3f, 0x03, 0x59, 0x30, 0xf3, 0xc0, 0xad, 0xf3, 0xef, 0x51, 0x2b,
	0x9e, 0x79, 0xdc, 0x46, 0xf6, 0x70, 0x95, 0xc1, 0x77, 0x67, 0xe2, 0x99, 0x93, 0x2b, 0x38, 0x99,
	0x1b, 0xf6, 0x20, 0xb9, 0x72, 0xcc, 0x09, 0xad, 0xa8, 0xe4, 0xce, 0x88, 0x2c, 0x8e, 0x86, 0xc1,
	0xb8, 0x95, 0x1c, 0x6f, 0xe4, 0xa6, 0x98, 0x1a, 0xfd, 0x6e, 0x41, 0x6f, 0xca, 0xe5, 0xdd, 0x63,
	0xa9, 0x9e, 0x48, 0x0c, 0x5d, 0x96, 0xe7, 0x86, 0x5b, 0x3f, 0x56, 0x3b, 0xa9, 0x8f, 0x84, 0x40,
	0x1b, 0x5f, 0xf6, 0xc3, 0xe0, 0x37, 0x79, 0x0f, 0x07, 0x86, 0xff, 0x28, 0xb9, 0x75, 0x3c, 0xf7,
	0x7d, 0xf9, 0x19, 0xf6, 0xd7, 0x51, 0x6c, 0xea, 0x10, 0xc2, 0x54, 0x28, 0xec, 0x39, 0x4a, 0xaa,
	0x4f, 0x72, 0x0a, 0x5d, 0x5d, 0x50, 0xc5, 0x24, 0xc7, 0xce, 0xfa, 0x49, 0x47, 0x17, 0xdf, 0x98,
	0xe4, 0xe4, 0x1d, 0x1c, 0xcc, 0x0d, 0xe7, 0x39, 0x65, 0x8e, 0x66, 0xba, 0x54, 0x2e, 0xee, 0x60,
	0x1b, 0x7b, 0x18, 0xbd, 0x75, 0x77, 0x55, 0x8c, 0xbc, 0x85, 0x3d, 0x96, 0xe1, 0x78, 0x9e, 0xe9,
	0x22, 0x33, 0xf0, 0x31, 0x8f, 0xbc, 0x86, 0xce, 0x4a, 0xd6, 0xde, 0x30, 0x18, 0xf7, 0x92, 0x48,
	0xa0, 0x9a, 0xa7, 0xd0, 0xb5, 0x8e, 0x17, 0x54, 0xe4, 0x71, 0x1f, 0x2f, 0x75, 0xaa, 0xe3, 0x7d,
	0x3e, 0xfa, 0x13, 0x00, 0x4c, 0x84, 0x9a, 0x95, 0x52, 0x32, 0xb3, 0xac, 0x5b, 0x0e, 0x9a, 0x96,
	0x3f, 0x01, 0x71, 0xda, 0xb1, 0x05, 0xdd, 0x61, 0xed, 0x2b, 0xcc, 0xbc, 0x30, 0x6d, 0x1b, 0xae,
	0xaa, 0x85, 0xdb, 0xf0, 0x44, 0x28, 0x72, 0x0e, 0xc7, 0x1e, 0xce, 0x2a, 0x0b, 0xd6, 0xa5, 0x57,
	0x16, 0x63, 0x0a, 0xcd, 0x59, 0xd5, 0xde, 0x81, 0x57, 0xc5, 0xa3, 0x1d, 0xf8, 0x44, 0xa8, 0xd1,
	0x2d, 0xf4, 0x66, 0x8a, 0x15, 0xb3, 0x47, 0xbd, 0xad, 0x5b, 0xb0, 0xad, 0xdb, 0x0e, 0x9b, 0x47,
	0xbf, 0x5a, 0x00, 0x53, 0xdc, 0x92, 0xaf, 0xa5, 0x2c, 0x2a, 0xd7, 0x59, 0xbd, 0x0f, 0xde, 0xc3,
	0x00, 0x3d, 0xdc, 0x5f, 0x47, 0xd1, 0xca, 0x2f, 0x30, 0x48, 0x85, 0xa2, 0xd6, 0x2b, 0x1a, 0xb7,
	0x86, 0xe1, 0x78, 0x70, 0xfd, 0xe6, 0xa2, 0x59, 0xae, 0x8b, 0x46, 0xef, 0x04, 0xd2, 0x46, 0xfb,
	0x8f, 0x10, 0xe1, 0x68, 0x71, 0x88, 0x57, 0x4e, 0x5e, 0x5e, 0xa9, 0x7f, 0xa8, 0x89, 0x47, 0xc8,
	0x15, 0xf4, 0xad, 0x62, 0x05, 0xb5, 0x8f, 0xda, 0xc5, 0xed, 0x6d, 0xbe, 0x1e, 0x3d, 0xe9, 0xd9,
	0x5a, 0x84, 0x1b, 0x88, 0x6c, 0xb5, 0xc5, 0xa8, 0xd8, 0xe0, 0xfa, 0xec, 0xbf, 0xf2, 0x9b, 0xab,
	0x9e, 0x78, 0x36, 0xed, 0xe0, 0x5f, 0xc3, 0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x7a,
	0x14, 0x28, 0x4b, 0x04, 0x00, 0x00,
}
