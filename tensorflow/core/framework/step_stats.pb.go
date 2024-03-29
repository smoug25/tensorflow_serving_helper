// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/step_stats.proto

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
	return fileDescriptor_1e915309f7ed52e5, []int{0}
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
	return fileDescriptor_1e915309f7ed52e5, []int{1}
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
	return fileDescriptor_1e915309f7ed52e5, []int{2}
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
	return fileDescriptor_1e915309f7ed52e5, []int{3}
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
	return fileDescriptor_1e915309f7ed52e5, []int{4}
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
	return fileDescriptor_1e915309f7ed52e5, []int{5}
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
	return fileDescriptor_1e915309f7ed52e5, []int{6}
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
	proto.RegisterFile("tensorflow/core/framework/step_stats.proto", fileDescriptor_1e915309f7ed52e5)
}

var fileDescriptor_1e915309f7ed52e5 = []byte{
	// 932 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x5d, 0x4f, 0xe3, 0x46,
	0x14, 0x95, 0x13, 0x48, 0xc9, 0x0d, 0x21, 0xce, 0x80, 0x58, 0x17, 0x4a, 0x97, 0x8d, 0xd4, 0x6e,
	0xe8, 0x47, 0xa8, 0xd8, 0xaa, 0xa5, 0x2b, 0xb5, 0x52, 0xd1, 0xe6, 0x01, 0x75, 0x37, 0x8b, 0x0c,
	0xdb, 0x57, 0xcb, 0x64, 0x2e, 0x5d, 0x0b, 0xc7, 0x63, 0xcd, 0x4c, 0xb2, 0x85, 0xbf, 0xd1, 0xe7,
	0xfe, 0xcb, 0x3e, 0xf4, 0xa9, 0xaa, 0xe6, 0xce, 0xc4, 0x76, 0x12, 0xd8, 0x37, 0xfb, 0xcc, 0xb9,
	0xd7, 0xf7, 0xe3, 0x9c, 0x31, 0x7c, 0xa5, 0x31, 0x53, 0x42, 0xde, 0xa4, 0xe2, 0xc3, 0xf1, 0x58,
	0x48, 0x3c, 0xbe, 0x91, 0xf1, 0x04, 0x3f, 0x08, 0x79, 0x7b, 0xac, 0x34, 0xe6, 0x91, 0xd2, 0xb1,
	0x56, 0x83, 0x5c, 0x0a, 0x2d, 0x18, 0x94, 0xdc, 0xbd, 0x1f, 0x1e, 0x8f, 0x8b, 0xd3, 0x54, 0x8c,
	0x63, 0x9d, 0x88, 0x2c, 0xe2, 0xa8, 0xc6, 0x32, 0xc9, 0xcd, 0xb3, 0xcd, 0xb1, 0x77, 0xf2, 0x78,
	0x9c, 0x3d, 0x59, 0x8d, 0xe9, 0xfd, 0x0e, 0xfe, 0xaf, 0x45, 0xce, 0x10, 0xc7, 0x42, 0x72, 0xf6,
	0x0c, 0x36, 0xe9, 0x3b, 0xd1, 0x24, 0x19, 0x4b, 0xa1, 0x02, 0xef, 0xd0, 0xeb, 0xd7, 0xc3, 0x16,
	0x61, 0x6f, 0x08, 0x62, 0x4f, 0xc1, 0xbe, 0x46, 0xd7, 0x77, 0x1a, 0x55, 0x50, 0x23, 0x06, 0x10,
	0x74, 0x66, 0x90, 0xde, 0xdf, 0x35, 0xd8, 0x76, 0x89, 0x85, 0x7c, 0x83, 0x13, 0x21, 0xef, 0xde,
	0x29, 0xe4, 0xec, 0x0b, 0xd8, 0x8a, 0xe7, 0x70, 0x94, 0xc5, 0x13, 0xa4, 0xec, 0xcd, 0xb0, 0x5d,
	0xa0, 0xa3, 0x78, 0x82, 0x26, 0xbf, 0x16, 0x3a, 0x4e, 0x17, 0xf3, 0x13, 0x44, 0xf9, 0xd9, 0x01,
	0x40, 0x8e, 0xf1, 0xad, 0x3b, 0xaf, 0xd3, 0x79, 0xd3, 0x20, 0xc5, 0x71, 0x9a, 0xcc, 0xd0, 0x1d,
	0xaf, 0xd9, 0x63, 0x83, 0xd8, 0xe3, 0xdf, 0x80, 0x55, 0x26, 0x29, 0xa9, 0x6d, 0x15, 0x34, 0x0e,
	0xeb, 0xfd, 0xd6, 0xc9, 0x67, 0x83, 0x72, 0x8c, 0x83, 0xe5, 0xd9, 0x84, 0xdd, 0x78, 0x09, 0x51,
	0xec, 0x05, 0xec, 0x96, 0x2d, 0xd1, 0x07, 0xa3, 0x24, 0x8b, 0xa6, 0x0a, 0x83, 0x75, 0xfa, 0xee,
	0x76, 0x71, 0x4a, 0x1f, 0x3f, 0xcf, 0xde, 0x29, 0xec, 0x65, 0x00, 0x23, 0xc1, 0xf1, 0xed, 0x54,
	0xe7, 0x53, 0xcd, 0x18, 0xac, 0xa9, 0x54, 0x68, 0x9a, 0xc5, 0x7a, 0x48, 0xcf, 0xec, 0x35, 0xb0,
	0xd5, 0xad, 0x51, 0xa7, 0xad, 0x93, 0x83, 0x6a, 0x8d, 0x57, 0xf4, 0xf8, 0xaa, 0x24, 0x85, 0x5d,
	0xbd, 0x0c, 0xf5, 0xfe, 0xab, 0x41, 0xcb, 0xae, 0xe1, 0xd2, 0xa8, 0x8e, 0xf5, 0xc1, 0xd7, 0x38,
	0xc9, 0xa3, 0x09, 0x61, 0x91, 0x4a, 0xee, 0xd1, 0xed, 0x79, 0xcb, 0xe0, 0x8e, 0x9a, 0xdc, 0x23,
	0xfb, 0x1e, 0x76, 0x73, 0x94, 0x2a, 0x51, 0x1a, 0x33, 0xbd, 0xc0, 0xb7, 0x53, 0xdf, 0x29, 0x4f,
	0x2b, 0x51, 0x3f, 0xc3, 0x7e, 0x25, 0xca, 0x35, 0x62, 0x25, 0x93, 0x70, 0x15, 0xac, 0x1f, 0xd6,
	0xfb, 0xf5, 0x30, 0x28, 0x29, 0xb6, 0x09, 0x1a, 0xf7, 0x39, 0x57, 0xec, 0x27, 0x78, 0xc2, 0x71,
	0x96, 0x8c, 0x31, 0x5a, 0xa9, 0x92, 0xb4, 0x70, 0x56, 0x0b, 0xbc, 0x70, 0xc7, 0x52, 0xae, 0x16,
	0xeb, 0x1d, 0xc2, 0x81, 0x0b, 0x7d, 0xa4, 0xec, 0xb5, 0x22, 0xc1, 0x9e, 0x25, 0x5e, 0x3c, 0xd4,
	0xc0, 0x08, 0x7a, 0xab, 0x69, 0x56, 0xfa, 0x30, 0x92, 0xb1, 0xb9, 0x3e, 0x5f, 0xce, 0xb5, 0xd8,
	0x51, 0xef, 0xaf, 0x06, 0xb4, 0xcd, 0xc6, 0x87, 0x7f, 0xe2, 0xd8, 0xae, 0x60, 0x1f, 0x9a, 0x99,
	0xe0, 0x58, 0x75, 0xc1, 0x86, 0x01, 0xc8, 0x00, 0x7d, 0xf0, 0xe3, 0x34, 0x35, 0x57, 0x84, 0xd4,
	0x73, 0x1f, 0x5a, 0x17, 0x18, 0xff, 0x5c, 0x1a, 0xd8, 0x59, 0xf1, 0x5b, 0xd8, 0x16, 0xb9, 0x23,
	0x4a, 0x4c, 0xe7, 0x64, 0xbb, 0x1c, 0x5f, 0xe4, 0xc4, 0x0d, 0x31, 0x75, 0xf4, 0x23, 0xe8, 0x8a,
	0x3c, 0xc2, 0x8c, 0x57, 0xc9, 0xd6, 0x20, 0x5b, 0x22, 0x1f, 0x66, 0xbc, 0xa4, 0x7e, 0x4d, 0x2e,
	0x59, 0xe6, 0x5a, 0x51, 0x77, 0xe2, 0x34, 0x5d, 0x20, 0xff, 0x08, 0x0d, 0x3b, 0x64, 0x67, 0xa3,
	0xa7, 0x0f, 0xd8, 0xa8, 0x7a, 0x13, 0x84, 0x8e, 0xce, 0x06, 0xd0, 0x10, 0xe4, 0x82, 0xe0, 0x13,
	0x0a, 0xdc, 0xad, 0x06, 0x96, 0x1e, 0x09, 0x1d, 0xcb, 0xdc, 0x20, 0x3a, 0x99, 0x60, 0x9a, 0x64,
	0x18, 0xa5, 0xf1, 0x35, 0xa6, 0xc1, 0x86, 0xbd, 0x41, 0xe6, 0xe8, 0x6b, 0x03, 0xb2, 0x23, 0xf0,
	0xd5, 0xf8, 0x3d, 0xf2, 0x69, 0x8a, 0x7c, 0x5e, 0x7a, 0xd3, 0x96, 0x5e, 0xe0, 0xae, 0xf4, 0x7d,
	0x68, 0xea, 0xf7, 0x12, 0x63, 0x1e, 0x25, 0x3c, 0x80, 0x43, 0xaf, 0xdf, 0x0e, 0x37, 0x2c, 0x70,
	0xce, 0xd9, 0x08, 0xba, 0x12, 0x6f, 0x50, 0x62, 0x36, 0x46, 0xee, 0x04, 0x10, 0xb4, 0xa8, 0xd2,
	0x67, 0x0f, 0xdf, 0x14, 0x55, 0x27, 0xfa, 0x65, 0xac, 0xd5, 0x03, 0x7b, 0x09, 0x9b, 0x73, 0x31,
	0x1a, 0x15, 0x04, 0x9b, 0x64, 0xe8, 0x27, 0xd5, 0x54, 0x15, 0x9f, 0x86, 0xad, 0x49, 0xc5, 0xb4,
	0x5f, 0x42, 0xa7, 0x14, 0x45, 0x16, 0x67, 0x42, 0x05, 0x6d, 0x6a, 0xa9, 0x3d, 0xd7, 0xc4, 0xc8,
	0x80, 0x66, 0x71, 0x0b, 0x92, 0xb0, 0xd4, 0x2d, 0xdb, 0x7d, 0xa9, 0x08, 0x4b, 0x7e, 0x0e, 0x7e,
	0x45, 0x10, 0x96, 0xda, 0xb1, 0x59, 0xe7, 0x7a, 0xb0, 0xc4, 0x23, 0xe8, 0x56, 0xe5, 0x60, 0x99,
	0x7e, 0xa1, 0xc9, 0x2a, 0xf5, 0x39, 0x94, 0x43, 0x76, 0xc4, 0xae, 0x25, 0x16, 0x30, 0x11, 0x7b,
	0xff, 0x78, 0xd0, 0x79, 0x45, 0xc6, 0xb9, 0xd4, 0x98, 0xdb, 0x2e, 0x77, 0xa1, 0x61, 0xbd, 0xe4,
	0x4c, 0xe1, 0xde, 0xd8, 0x29, 0x00, 0xf9, 0xc5, 0xce, 0xad, 0x46, 0x2b, 0xf8, 0x74, 0x59, 0x2c,
	0x85, 0xbd, 0x42, 0x32, 0x97, 0xcd, 0xf8, 0x16, 0x36, 0xdd, 0x82, 0x8d, 0xd7, 0x8c, 0x37, 0x4c,
	0xec, 0x37, 0xd5, 0xd8, 0xa5, 0x22, 0x06, 0x57, 0xc4, 0x37, 0x56, 0x54, 0xc3, 0x4c, 0xcb, 0xbb,
	0xb0, 0xa5, 0x4b, 0x64, 0xef, 0x17, 0xf0, 0x97, 0x09, 0xcc, 0x87, 0xfa, 0x2d, 0xde, 0x51, 0xcd,
	0xed, 0xd0, 0x3c, 0xb2, 0x1d, 0x58, 0x9f, 0xc5, 0xe9, 0xd4, 0x5e, 0x59, 0xcd, 0xd0, 0xbe, 0xbc,
	0xac, 0x9d, 0x7a, 0xbd, 0x21, 0x34, 0xcb, 0x7e, 0x4f, 0xa1, 0xc9, 0x71, 0xe6, 0xda, 0xf2, 0xa8,
	0xb4, 0xfd, 0x8f, 0x94, 0x16, 0x6e, 0x70, 0x9c, 0xd1, 0xd3, 0xd9, 0x77, 0x10, 0x08, 0xf9, 0x47,
	0x95, 0x5b, 0xfc, 0xf1, 0xcf, 0x3a, 0x45, 0xc0, 0x85, 0xf9, 0xd1, 0xab, 0x0b, 0xef, 0x5f, 0xcf,
	0xbb, 0x6e, 0xd0, 0x5f, 0xff, 0xc5, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xde, 0x78, 0x3e,
	0x9b, 0x08, 0x00, 0x00,
}
