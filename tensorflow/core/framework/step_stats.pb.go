// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework/step_stats.proto

package framework

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

// An allocation/de-allocation operation performed by the allocator.
type AllocationRecord struct {
	// The timestamp of the operation.
	AllocMicros int64 `protobuf:"varint,1,opt,name=alloc_micros,json=allocMicros,proto3" json:"alloc_micros,omitempty"`
	// Number of bytes allocated, or de-allocated if negative.
	AllocBytes           int64    `protobuf:"varint,2,opt,name=alloc_bytes,json=allocBytes,proto3" json:"alloc_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationRecord) Reset()         { *m = AllocationRecord{} }
func (m *AllocationRecord) String() string { return proto.CompactTextString(m) }
func (*AllocationRecord) ProtoMessage()    {}
func (*AllocationRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{0}
}

func (m *AllocationRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationRecord.Unmarshal(m, b)
}
func (m *AllocationRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationRecord.Marshal(b, m, deterministic)
}
func (m *AllocationRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationRecord.Merge(m, src)
}
func (m *AllocationRecord) XXX_Size() int {
	return xxx_messageInfo_AllocationRecord.Size(m)
}
func (m *AllocationRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationRecord.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationRecord proto.InternalMessageInfo

func (m *AllocationRecord) GetAllocMicros() int64 {
	if m != nil {
		return m.AllocMicros
	}
	return 0
}

func (m *AllocationRecord) GetAllocBytes() int64 {
	if m != nil {
		return m.AllocBytes
	}
	return 0
}

type AllocatorMemoryUsed struct {
	AllocatorName string `protobuf:"bytes,1,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// These are per-node allocator memory stats.
	TotalBytes int64 `protobuf:"varint,2,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	PeakBytes  int64 `protobuf:"varint,3,opt,name=peak_bytes,json=peakBytes,proto3" json:"peak_bytes,omitempty"`
	// The bytes that are not deallocated.
	LiveBytes int64 `protobuf:"varint,4,opt,name=live_bytes,json=liveBytes,proto3" json:"live_bytes,omitempty"`
	// The allocation and deallocation timeline.
	AllocationRecords []*AllocationRecord `protobuf:"bytes,6,rep,name=allocation_records,json=allocationRecords,proto3" json:"allocation_records,omitempty"`
	// These are snapshots of the overall allocator memory stats.
	// The number of live bytes currently allocated by the allocator.
	AllocatorBytesInUse  int64    `protobuf:"varint,5,opt,name=allocator_bytes_in_use,json=allocatorBytesInUse,proto3" json:"allocator_bytes_in_use,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocatorMemoryUsed) Reset()         { *m = AllocatorMemoryUsed{} }
func (m *AllocatorMemoryUsed) String() string { return proto.CompactTextString(m) }
func (*AllocatorMemoryUsed) ProtoMessage()    {}
func (*AllocatorMemoryUsed) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{1}
}

func (m *AllocatorMemoryUsed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocatorMemoryUsed.Unmarshal(m, b)
}
func (m *AllocatorMemoryUsed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocatorMemoryUsed.Marshal(b, m, deterministic)
}
func (m *AllocatorMemoryUsed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocatorMemoryUsed.Merge(m, src)
}
func (m *AllocatorMemoryUsed) XXX_Size() int {
	return xxx_messageInfo_AllocatorMemoryUsed.Size(m)
}
func (m *AllocatorMemoryUsed) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocatorMemoryUsed.DiscardUnknown(m)
}

var xxx_messageInfo_AllocatorMemoryUsed proto.InternalMessageInfo

func (m *AllocatorMemoryUsed) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocatorMemoryUsed) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetPeakBytes() int64 {
	if m != nil {
		return m.PeakBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetLiveBytes() int64 {
	if m != nil {
		return m.LiveBytes
	}
	return 0
}

func (m *AllocatorMemoryUsed) GetAllocationRecords() []*AllocationRecord {
	if m != nil {
		return m.AllocationRecords
	}
	return nil
}

func (m *AllocatorMemoryUsed) GetAllocatorBytesInUse() int64 {
	if m != nil {
		return m.AllocatorBytesInUse
	}
	return 0
}

// Output sizes recorded for a single execution of a graph node.
type NodeOutput struct {
	Slot                 int32              `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	TensorDescription    *TensorDescription `protobuf:"bytes,3,opt,name=tensor_description,json=tensorDescription,proto3" json:"tensor_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NodeOutput) Reset()         { *m = NodeOutput{} }
func (m *NodeOutput) String() string { return proto.CompactTextString(m) }
func (*NodeOutput) ProtoMessage()    {}
func (*NodeOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{2}
}

func (m *NodeOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeOutput.Unmarshal(m, b)
}
func (m *NodeOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeOutput.Marshal(b, m, deterministic)
}
func (m *NodeOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeOutput.Merge(m, src)
}
func (m *NodeOutput) XXX_Size() int {
	return xxx_messageInfo_NodeOutput.Size(m)
}
func (m *NodeOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeOutput.DiscardUnknown(m)
}

var xxx_messageInfo_NodeOutput proto.InternalMessageInfo

func (m *NodeOutput) GetSlot() int32 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *NodeOutput) GetTensorDescription() *TensorDescription {
	if m != nil {
		return m.TensorDescription
	}
	return nil
}

// For memory tracking.
type MemoryStats struct {
	TempMemorySize                 int64    `protobuf:"varint,1,opt,name=temp_memory_size,json=tempMemorySize,proto3" json:"temp_memory_size,omitempty"`
	PersistentMemorySize           int64    `protobuf:"varint,3,opt,name=persistent_memory_size,json=persistentMemorySize,proto3" json:"persistent_memory_size,omitempty"`
	PersistentTensorAllocIds       []int64  `protobuf:"varint,5,rep,packed,name=persistent_tensor_alloc_ids,json=persistentTensorAllocIds,proto3" json:"persistent_tensor_alloc_ids,omitempty"`
	DeviceTempMemorySize           int64    `protobuf:"varint,2,opt,name=device_temp_memory_size,json=deviceTempMemorySize,proto3" json:"device_temp_memory_size,omitempty"`                                        // Deprecated: Do not use.
	DevicePersistentMemorySize     int64    `protobuf:"varint,4,opt,name=device_persistent_memory_size,json=devicePersistentMemorySize,proto3" json:"device_persistent_memory_size,omitempty"`                      // Deprecated: Do not use.
	DevicePersistentTensorAllocIds []int64  `protobuf:"varint,6,rep,packed,name=device_persistent_tensor_alloc_ids,json=devicePersistentTensorAllocIds,proto3" json:"device_persistent_tensor_alloc_ids,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *MemoryStats) Reset()         { *m = MemoryStats{} }
func (m *MemoryStats) String() string { return proto.CompactTextString(m) }
func (*MemoryStats) ProtoMessage()    {}
func (*MemoryStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{3}
}

func (m *MemoryStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryStats.Unmarshal(m, b)
}
func (m *MemoryStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryStats.Marshal(b, m, deterministic)
}
func (m *MemoryStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryStats.Merge(m, src)
}
func (m *MemoryStats) XXX_Size() int {
	return xxx_messageInfo_MemoryStats.Size(m)
}
func (m *MemoryStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryStats.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryStats proto.InternalMessageInfo

func (m *MemoryStats) GetTempMemorySize() int64 {
	if m != nil {
		return m.TempMemorySize
	}
	return 0
}

func (m *MemoryStats) GetPersistentMemorySize() int64 {
	if m != nil {
		return m.PersistentMemorySize
	}
	return 0
}

func (m *MemoryStats) GetPersistentTensorAllocIds() []int64 {
	if m != nil {
		return m.PersistentTensorAllocIds
	}
	return nil
}

// Deprecated: Do not use.
func (m *MemoryStats) GetDeviceTempMemorySize() int64 {
	if m != nil {
		return m.DeviceTempMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *MemoryStats) GetDevicePersistentMemorySize() int64 {
	if m != nil {
		return m.DevicePersistentMemorySize
	}
	return 0
}

// Deprecated: Do not use.
func (m *MemoryStats) GetDevicePersistentTensorAllocIds() []int64 {
	if m != nil {
		return m.DevicePersistentTensorAllocIds
	}
	return nil
}

// Time/size stats recorded for a single execution of a graph node.
type NodeExecStats struct {
	// TODO(tucker): Use some more compact form of node identity than
	// the full string name.  Either all processes should agree on a
	// global id (cost_id?) for each node, or we should use a hash of
	// the name.
	NodeName             string                   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AllStartMicros       int64                    `protobuf:"varint,2,opt,name=all_start_micros,json=allStartMicros,proto3" json:"all_start_micros,omitempty"`
	OpStartRelMicros     int64                    `protobuf:"varint,3,opt,name=op_start_rel_micros,json=opStartRelMicros,proto3" json:"op_start_rel_micros,omitempty"`
	OpEndRelMicros       int64                    `protobuf:"varint,4,opt,name=op_end_rel_micros,json=opEndRelMicros,proto3" json:"op_end_rel_micros,omitempty"`
	AllEndRelMicros      int64                    `protobuf:"varint,5,opt,name=all_end_rel_micros,json=allEndRelMicros,proto3" json:"all_end_rel_micros,omitempty"`
	Memory               []*AllocatorMemoryUsed   `protobuf:"bytes,6,rep,name=memory,proto3" json:"memory,omitempty"`
	Output               []*NodeOutput            `protobuf:"bytes,7,rep,name=output,proto3" json:"output,omitempty"`
	TimelineLabel        string                   `protobuf:"bytes,8,opt,name=timeline_label,json=timelineLabel,proto3" json:"timeline_label,omitempty"`
	ScheduledMicros      int64                    `protobuf:"varint,9,opt,name=scheduled_micros,json=scheduledMicros,proto3" json:"scheduled_micros,omitempty"`
	ThreadId             uint32                   `protobuf:"varint,10,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	ReferencedTensor     []*AllocationDescription `protobuf:"bytes,11,rep,name=referenced_tensor,json=referencedTensor,proto3" json:"referenced_tensor,omitempty"`
	MemoryStats          *MemoryStats             `protobuf:"bytes,12,opt,name=memory_stats,json=memoryStats,proto3" json:"memory_stats,omitempty"`
	AllStartNanos        int64                    `protobuf:"varint,13,opt,name=all_start_nanos,json=allStartNanos,proto3" json:"all_start_nanos,omitempty"`
	OpStartRelNanos      int64                    `protobuf:"varint,14,opt,name=op_start_rel_nanos,json=opStartRelNanos,proto3" json:"op_start_rel_nanos,omitempty"`
	OpEndRelNanos        int64                    `protobuf:"varint,15,opt,name=op_end_rel_nanos,json=opEndRelNanos,proto3" json:"op_end_rel_nanos,omitempty"`
	AllEndRelNanos       int64                    `protobuf:"varint,16,opt,name=all_end_rel_nanos,json=allEndRelNanos,proto3" json:"all_end_rel_nanos,omitempty"`
	ScheduledNanos       int64                    `protobuf:"varint,17,opt,name=scheduled_nanos,json=scheduledNanos,proto3" json:"scheduled_nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NodeExecStats) Reset()         { *m = NodeExecStats{} }
func (m *NodeExecStats) String() string { return proto.CompactTextString(m) }
func (*NodeExecStats) ProtoMessage()    {}
func (*NodeExecStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{4}
}

func (m *NodeExecStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecStats.Unmarshal(m, b)
}
func (m *NodeExecStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecStats.Marshal(b, m, deterministic)
}
func (m *NodeExecStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecStats.Merge(m, src)
}
func (m *NodeExecStats) XXX_Size() int {
	return xxx_messageInfo_NodeExecStats.Size(m)
}
func (m *NodeExecStats) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecStats.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecStats proto.InternalMessageInfo

func (m *NodeExecStats) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *NodeExecStats) GetAllStartMicros() int64 {
	if m != nil {
		return m.AllStartMicros
	}
	return 0
}

func (m *NodeExecStats) GetOpStartRelMicros() int64 {
	if m != nil {
		return m.OpStartRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetOpEndRelMicros() int64 {
	if m != nil {
		return m.OpEndRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetAllEndRelMicros() int64 {
	if m != nil {
		return m.AllEndRelMicros
	}
	return 0
}

func (m *NodeExecStats) GetMemory() []*AllocatorMemoryUsed {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *NodeExecStats) GetOutput() []*NodeOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *NodeExecStats) GetTimelineLabel() string {
	if m != nil {
		return m.TimelineLabel
	}
	return ""
}

func (m *NodeExecStats) GetScheduledMicros() int64 {
	if m != nil {
		return m.ScheduledMicros
	}
	return 0
}

func (m *NodeExecStats) GetThreadId() uint32 {
	if m != nil {
		return m.ThreadId
	}
	return 0
}

func (m *NodeExecStats) GetReferencedTensor() []*AllocationDescription {
	if m != nil {
		return m.ReferencedTensor
	}
	return nil
}

func (m *NodeExecStats) GetMemoryStats() *MemoryStats {
	if m != nil {
		return m.MemoryStats
	}
	return nil
}

func (m *NodeExecStats) GetAllStartNanos() int64 {
	if m != nil {
		return m.AllStartNanos
	}
	return 0
}

func (m *NodeExecStats) GetOpStartRelNanos() int64 {
	if m != nil {
		return m.OpStartRelNanos
	}
	return 0
}

func (m *NodeExecStats) GetOpEndRelNanos() int64 {
	if m != nil {
		return m.OpEndRelNanos
	}
	return 0
}

func (m *NodeExecStats) GetAllEndRelNanos() int64 {
	if m != nil {
		return m.AllEndRelNanos
	}
	return 0
}

func (m *NodeExecStats) GetScheduledNanos() int64 {
	if m != nil {
		return m.ScheduledNanos
	}
	return 0
}

type DeviceStepStats struct {
	Device    string           `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	NodeStats []*NodeExecStats `protobuf:"bytes,2,rep,name=node_stats,json=nodeStats,proto3" json:"node_stats,omitempty"`
	// Its key is thread id.
	ThreadNames          map[uint32]string `protobuf:"bytes,3,rep,name=thread_names,json=threadNames,proto3" json:"thread_names,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeviceStepStats) Reset()         { *m = DeviceStepStats{} }
func (m *DeviceStepStats) String() string { return proto.CompactTextString(m) }
func (*DeviceStepStats) ProtoMessage()    {}
func (*DeviceStepStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{5}
}

func (m *DeviceStepStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStepStats.Unmarshal(m, b)
}
func (m *DeviceStepStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStepStats.Marshal(b, m, deterministic)
}
func (m *DeviceStepStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStepStats.Merge(m, src)
}
func (m *DeviceStepStats) XXX_Size() int {
	return xxx_messageInfo_DeviceStepStats.Size(m)
}
func (m *DeviceStepStats) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStepStats.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStepStats proto.InternalMessageInfo

func (m *DeviceStepStats) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *DeviceStepStats) GetNodeStats() []*NodeExecStats {
	if m != nil {
		return m.NodeStats
	}
	return nil
}

func (m *DeviceStepStats) GetThreadNames() map[uint32]string {
	if m != nil {
		return m.ThreadNames
	}
	return nil
}

type StepStats struct {
	DevStats             []*DeviceStepStats `protobuf:"bytes,1,rep,name=dev_stats,json=devStats,proto3" json:"dev_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StepStats) Reset()         { *m = StepStats{} }
func (m *StepStats) String() string { return proto.CompactTextString(m) }
func (*StepStats) ProtoMessage()    {}
func (*StepStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_664c6e253478b684, []int{6}
}

func (m *StepStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepStats.Unmarshal(m, b)
}
func (m *StepStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepStats.Marshal(b, m, deterministic)
}
func (m *StepStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepStats.Merge(m, src)
}
func (m *StepStats) XXX_Size() int {
	return xxx_messageInfo_StepStats.Size(m)
}
func (m *StepStats) XXX_DiscardUnknown() {
	xxx_messageInfo_StepStats.DiscardUnknown(m)
}

var xxx_messageInfo_StepStats proto.InternalMessageInfo

func (m *StepStats) GetDevStats() []*DeviceStepStats {
	if m != nil {
		return m.DevStats
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationRecord)(nil), "tensorflow.AllocationRecord")
	proto.RegisterType((*AllocatorMemoryUsed)(nil), "tensorflow.AllocatorMemoryUsed")
	proto.RegisterType((*NodeOutput)(nil), "tensorflow.NodeOutput")
	proto.RegisterType((*MemoryStats)(nil), "tensorflow.MemoryStats")
	proto.RegisterType((*NodeExecStats)(nil), "tensorflow.NodeExecStats")
	proto.RegisterType((*DeviceStepStats)(nil), "tensorflow.DeviceStepStats")
	proto.RegisterMapType((map[uint32]string)(nil), "tensorflow.DeviceStepStats.ThreadNamesEntry")
	proto.RegisterType((*StepStats)(nil), "tensorflow.StepStats")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework/step_stats.proto", fileDescriptor_664c6e253478b684)
}

var fileDescriptor_664c6e253478b684 = []byte{
	// 974 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xd1, 0x4f, 0xe3, 0xc6,
	0x13, 0x96, 0x13, 0xc8, 0x8f, 0x4c, 0x08, 0x38, 0x0b, 0xe2, 0xfc, 0x83, 0xd2, 0xe3, 0x22, 0xb5,
	0x17, 0xd4, 0x36, 0x48, 0x5c, 0xab, 0xd2, 0x93, 0x5a, 0xa9, 0xe8, 0x52, 0x09, 0xf5, 0x2e, 0x87,
	0x0c, 0xd7, 0x4a, 0x7d, 0xb1, 0x8c, 0x3d, 0x80, 0x85, 0xed, 0xb5, 0x76, 0x37, 0xb9, 0xe3, 0xfe,
	0x8d, 0x3e, 0xf7, 0xbf, 0xec, 0x43, 0x9f, 0xaa, 0x6a, 0x67, 0x37, 0xb6, 0x93, 0xd0, 0x3e, 0xf1,
	0xb6, 0xfe, 0xe6, 0xdb, 0xd9, 0x99, 0x9d, 0xef, 0xdb, 0x04, 0x7e, 0xbd, 0x49, 0xd4, 0xed, 0xe4,
	0x6a, 0x18, 0xf1, 0xec, 0x48, 0x66, 0x7c, 0x72, 0x73, 0xfc, 0xcd, 0x91, 0xc2, 0x5c, 0x72, 0x71,
	0x9d, 0xf2, 0xf7, 0x81, 0x44, 0x31, 0x4d, 0xf2, 0x9b, 0xe0, 0x16, 0xd3, 0x02, 0x45, 0x2d, 0x72,
	0x14, 0x71, 0x81, 0x47, 0xd7, 0x22, 0xcc, 0xf0, 0x3d, 0x17, 0x77, 0x47, 0x52, 0x61, 0x11, 0x48,
	0x15, 0x2a, 0x39, 0x2c, 0x04, 0x57, 0x9c, 0x41, 0xc5, 0xdd, 0x8d, 0x1e, 0xe9, 0x90, 0x30, 0x4d,
	0x79, 0x14, 0xaa, 0x84, 0xe7, 0x41, 0x8c, 0x32, 0x12, 0x49, 0xa1, 0xd7, 0xe6, 0xc0, 0xdd, 0xe0,
	0x91, 0x0e, 0x31, 0x91, 0xe5, 0x03, 0xfa, 0xbf, 0x80, 0xfb, 0x63, 0x59, 0x80, 0x8f, 0x11, 0x17,
	0x31, 0x7b, 0x06, 0xeb, 0x54, 0x54, 0x90, 0x25, 0x91, 0xe0, 0xd2, 0x73, 0x0e, 0x9c, 0x41, 0xd3,
	0xef, 0x10, 0xf6, 0x86, 0x20, 0xf6, 0x14, 0xcc, 0x67, 0x70, 0x75, 0xaf, 0x50, 0x7a, 0x0d, 0x62,
	0x00, 0x41, 0xa7, 0x1a, 0xe9, 0xff, 0xd1, 0x80, 0x2d, 0x9b, 0x98, 0x8b, 0x37, 0x98, 0x71, 0x71,
	0xff, 0x4e, 0x62, 0xcc, 0x3e, 0x83, 0x8d, 0x70, 0x06, 0x07, 0x79, 0x98, 0x21, 0x65, 0x6f, 0xfb,
	0xdd, 0x12, 0x1d, 0x87, 0x19, 0xea, 0xfc, 0x8a, 0xab, 0x30, 0x9d, 0xcf, 0x4f, 0x10, 0xe5, 0x67,
	0xfb, 0x00, 0x05, 0x86, 0x77, 0x36, 0xde, 0xa4, 0x78, 0x5b, 0x23, 0x65, 0x38, 0x4d, 0xa6, 0x68,
	0xc3, 0x2b, 0x26, 0xac, 0x11, 0x13, 0xfe, 0x19, 0x58, 0xed, 0xda, 0x05, 0xb5, 0x2d, 0xbd, 0xd6,
	0x41, 0x73, 0xd0, 0x39, 0xfe, 0x64, 0x58, 0x5d, 0xe3, 0x70, 0xf1, 0x6e, 0xfc, 0x5e, 0xb8, 0x80,
	0x48, 0xf6, 0x02, 0x76, 0xaa, 0x96, 0xe8, 0xc0, 0x20, 0xc9, 0x83, 0x89, 0x44, 0x6f, 0x95, 0xce,
	0xdd, 0x2a, 0xa3, 0x74, 0xf8, 0x59, 0xfe, 0x4e, 0x62, 0x3f, 0x07, 0x18, 0xf3, 0x18, 0xdf, 0x4e,
	0x54, 0x31, 0x51, 0x8c, 0xc1, 0x8a, 0x4c, 0xb9, 0xa2, 0xbb, 0x58, 0xf5, 0x69, 0xcd, 0x5e, 0x03,
	0x5b, 0x9e, 0x1a, 0x75, 0xda, 0x39, 0xde, 0xaf, 0xd7, 0x78, 0x49, 0xcb, 0x57, 0x15, 0xc9, 0xef,
	0xa9, 0x45, 0xa8, 0xff, 0x77, 0x03, 0x3a, 0x66, 0x0c, 0x17, 0x5a, 0xcf, 0x6c, 0x00, 0xae, 0xc2,
	0xac, 0x08, 0x32, 0xc2, 0x02, 0x99, 0x7c, 0x44, 0x3b, 0xe7, 0x0d, 0x8d, 0x5b, 0x6a, 0xf2, 0x11,
	0xd9, 0xd7, 0xb0, 0x53, 0xa0, 0x90, 0x89, 0x54, 0x98, 0xab, 0x39, 0xbe, 0xb9, 0xf5, 0xed, 0x2a,
	0x5a, 0xdb, 0xf5, 0x3d, 0xec, 0xd5, 0x76, 0xd9, 0x46, 0x8c, 0x64, 0x92, 0x58, 0x7a, 0xab, 0x07,
	0xcd, 0x41, 0xd3, 0xf7, 0x2a, 0x8a, 0x69, 0x82, 0xae, 0xfb, 0x2c, 0x96, 0xec, 0x3b, 0x78, 0x12,
	0xe3, 0x34, 0x89, 0x30, 0x58, 0xaa, 0x92, 0xb4, 0x70, 0xda, 0xf0, 0x1c, 0x7f, 0xdb, 0x50, 0x2e,
	0xe7, 0xeb, 0x1d, 0xc1, 0xbe, 0xdd, 0xfa, 0x2f, 0x65, 0xaf, 0x94, 0x09, 0x76, 0x0d, 0xf1, 0xfc,
	0xa1, 0x06, 0xc6, 0xd0, 0x5f, 0x4e, 0xb3, 0xd4, 0x87, 0x96, 0x8c, 0xc9, 0xf5, 0xe9, 0x62, 0xae,
	0xf9, 0x8e, 0xfa, 0xbf, 0xb7, 0xa0, 0xab, 0x27, 0x3e, 0xfa, 0x80, 0x91, 0x19, 0xc1, 0x1e, 0xb4,
	0x73, 0x1e, 0x63, 0xdd, 0x05, 0x6b, 0x1a, 0x20, 0x03, 0x0c, 0xc0, 0x0d, 0xd3, 0x54, 0x3f, 0x3e,
	0x42, 0xcd, 0x7c, 0x68, 0x5c, 0xa0, 0xfd, 0x73, 0xa1, 0x61, 0x6b, 0xc5, 0xaf, 0x60, 0x8b, 0x17,
	0x96, 0x28, 0x30, 0x9d, 0x91, 0xcd, 0x70, 0x5c, 0x5e, 0x10, 0xd7, 0xc7, 0xd4, 0xd2, 0x0f, 0xa1,
	0xc7, 0x8b, 0x00, 0xf3, 0xb8, 0x4e, 0x36, 0x06, 0xd9, 0xe0, 0xc5, 0x28, 0x8f, 0x2b, 0xea, 0x17,
	0xe4, 0x92, 0x45, 0xae, 0x11, 0xf5, 0x66, 0x98, 0xa6, 0x73, 0xe4, 0x6f, 0xa1, 0x65, 0x2e, 0xd9,
	0xda, 0xe8, 0xe9, 0x03, 0x36, 0xaa, 0xbf, 0x04, 0xbe, 0xa5, 0xb3, 0x21, 0xb4, 0x38, 0xb9, 0xc0,
	0xfb, 0x1f, 0x6d, 0xdc, 0xa9, 0x6f, 0xac, 0x3c, 0xe2, 0x5b, 0x96, 0x7e, 0x41, 0x54, 0x92, 0x61,
	0x9a, 0xe4, 0x18, 0xa4, 0xe1, 0x15, 0xa6, 0xde, 0x9a, 0x79, 0x41, 0x66, 0xe8, 0x6b, 0x0d, 0xb2,
	0x43, 0x70, 0x65, 0x74, 0x8b, 0xf1, 0x24, 0xc5, 0x78, 0x56, 0x7a, 0xdb, 0x94, 0x5e, 0xe2, 0xb6,
	0xf4, 0x3d, 0x68, 0xab, 0x5b, 0x81, 0x61, 0x1c, 0x24, 0xb1, 0x07, 0x07, 0xce, 0xa0, 0xeb, 0xaf,
	0x19, 0xe0, 0x2c, 0x66, 0x63, 0xe8, 0x09, 0xbc, 0x46, 0x81, 0x79, 0x84, 0xb1, 0x15, 0x80, 0xd7,
	0xa1, 0x4a, 0x9f, 0x3d, 0xfc, 0x52, 0xd4, 0x9d, 0xe8, 0x56, 0x7b, 0x8d, 0x1e, 0xd8, 0x4b, 0x58,
	0x9f, 0x89, 0x51, 0xab, 0xc0, 0x5b, 0x27, 0x43, 0x3f, 0xa9, 0xa7, 0xaa, 0xf9, 0xd4, 0xef, 0x64,
	0x35, 0xd3, 0x7e, 0x0e, 0x9b, 0x95, 0x28, 0xf2, 0x30, 0xe7, 0xd2, 0xeb, 0x52, 0x4b, 0xdd, 0x99,
	0x26, 0xc6, 0x1a, 0xd4, 0x83, 0x9b, 0x93, 0x84, 0xa1, 0x6e, 0x98, 0xee, 0x2b, 0x45, 0x18, 0xf2,
	0x73, 0x70, 0x6b, 0x82, 0x30, 0xd4, 0x4d, 0x93, 0x75, 0xa6, 0x07, 0x43, 0x3c, 0x84, 0x5e, 0x5d,
	0x0e, 0x86, 0xe9, 0x96, 0x9a, 0xac, 0x53, 0x9f, 0x43, 0x75, 0xc9, 0x96, 0xd8, 0x33, 0xc4, 0x12,
	0x26, 0x62, 0xff, 0x4f, 0x07, 0x36, 0x5f, 0x91, 0x71, 0x2e, 0x14, 0x16, 0xa6, 0xcb, 0x1d, 0x68,
	0x19, 0x2f, 0x59, 0x53, 0xd8, 0x2f, 0x76, 0x02, 0x40, 0x7e, 0x31, 0xf7, 0xd6, 0xa0, 0x11, 0xfc,
	0x7f, 0x51, 0x2c, 0xa5, 0xbd, 0x7c, 0x32, 0x97, 0xc9, 0xf8, 0x16, 0xd6, 0xed, 0x80, 0xb5, 0xd7,
	0xb4, 0x37, 0xf4, 0xde, 0x2f, 0xeb, 0x7b, 0x17, 0x8a, 0x18, 0x5e, 0x12, 0x5f, 0x5b, 0x51, 0x8e,
	0x72, 0x25, 0xee, 0xfd, 0x8e, 0xaa, 0x90, 0xdd, 0x1f, 0xc0, 0x5d, 0x24, 0x30, 0x17, 0x9a, 0x77,
	0x78, 0x4f, 0x35, 0x77, 0x7d, 0xbd, 0x64, 0xdb, 0xb0, 0x3a, 0x0d, 0xd3, 0x89, 0x79, 0xb2, 0xda,
	0xbe, 0xf9, 0x78, 0xd9, 0x38, 0x71, 0xfa, 0x23, 0x68, 0x57, 0xfd, 0x9e, 0x40, 0x3b, 0xc6, 0xa9,
	0x6d, 0xcb, 0xa1, 0xd2, 0xf6, 0xfe, 0xa3, 0x34, 0x7f, 0x2d, 0xc6, 0x29, 0xad, 0x4e, 0x3f, 0x80,
	0xc7, 0xc5, 0x4d, 0x9d, 0x5b, 0xfe, 0xe2, 0x9f, 0x6e, 0x96, 0x1b, 0xce, 0xf5, 0x0f, 0xbd, 0x3c,
	0x77, 0x7e, 0xfb, 0xe9, 0x71, 0xfe, 0x4c, 0xfc, 0xe5, 0x38, 0x57, 0x2d, 0xfa, 0xf7, 0xf0, 0xe2,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x07, 0xab, 0x31, 0x6a, 0x09, 0x00, 0x00,
}