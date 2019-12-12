// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/control_flow.proto

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

// Protocol buffer representing the values in ControlFlowContext.
type ValuesDef struct {
	// Value names that have been seen in this context.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// Value names referenced by but external to this context.
	ExternalValues       map[string]string `protobuf:"bytes,2,rep,name=external_values,json=externalValues,proto3" json:"external_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ValuesDef) Reset()         { *m = ValuesDef{} }
func (m *ValuesDef) String() string { return proto.CompactTextString(m) }
func (*ValuesDef) ProtoMessage()    {}
func (*ValuesDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_64affc5a646d7df1, []int{0}
}

func (m *ValuesDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValuesDef.Unmarshal(m, b)
}
func (m *ValuesDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValuesDef.Marshal(b, m, deterministic)
}
func (m *ValuesDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValuesDef.Merge(m, src)
}
func (m *ValuesDef) XXX_Size() int {
	return xxx_messageInfo_ValuesDef.Size(m)
}
func (m *ValuesDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ValuesDef.DiscardUnknown(m)
}

var xxx_messageInfo_ValuesDef proto.InternalMessageInfo

func (m *ValuesDef) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ValuesDef) GetExternalValues() map[string]string {
	if m != nil {
		return m.ExternalValues
	}
	return nil
}

// Container for any kind of control flow context. Any other control flow
// contexts that are added below should also be added here.
type ControlFlowContextDef struct {
	// Types that are valid to be assigned to Ctxt:
	//	*ControlFlowContextDef_CondCtxt
	//	*ControlFlowContextDef_WhileCtxt
	Ctxt                 isControlFlowContextDef_Ctxt `protobuf_oneof:"ctxt"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ControlFlowContextDef) Reset()         { *m = ControlFlowContextDef{} }
func (m *ControlFlowContextDef) String() string { return proto.CompactTextString(m) }
func (*ControlFlowContextDef) ProtoMessage()    {}
func (*ControlFlowContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_64affc5a646d7df1, []int{1}
}

func (m *ControlFlowContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlFlowContextDef.Unmarshal(m, b)
}
func (m *ControlFlowContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlFlowContextDef.Marshal(b, m, deterministic)
}
func (m *ControlFlowContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlFlowContextDef.Merge(m, src)
}
func (m *ControlFlowContextDef) XXX_Size() int {
	return xxx_messageInfo_ControlFlowContextDef.Size(m)
}
func (m *ControlFlowContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlFlowContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_ControlFlowContextDef proto.InternalMessageInfo

type isControlFlowContextDef_Ctxt interface {
	isControlFlowContextDef_Ctxt()
}

type ControlFlowContextDef_CondCtxt struct {
	CondCtxt *CondContextDef `protobuf:"bytes,1,opt,name=cond_ctxt,json=condCtxt,proto3,oneof"`
}

type ControlFlowContextDef_WhileCtxt struct {
	WhileCtxt *WhileContextDef `protobuf:"bytes,2,opt,name=while_ctxt,json=whileCtxt,proto3,oneof"`
}

func (*ControlFlowContextDef_CondCtxt) isControlFlowContextDef_Ctxt() {}

func (*ControlFlowContextDef_WhileCtxt) isControlFlowContextDef_Ctxt() {}

func (m *ControlFlowContextDef) GetCtxt() isControlFlowContextDef_Ctxt {
	if m != nil {
		return m.Ctxt
	}
	return nil
}

func (m *ControlFlowContextDef) GetCondCtxt() *CondContextDef {
	if x, ok := m.GetCtxt().(*ControlFlowContextDef_CondCtxt); ok {
		return x.CondCtxt
	}
	return nil
}

func (m *ControlFlowContextDef) GetWhileCtxt() *WhileContextDef {
	if x, ok := m.GetCtxt().(*ControlFlowContextDef_WhileCtxt); ok {
		return x.WhileCtxt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ControlFlowContextDef) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ControlFlowContextDef_CondCtxt)(nil),
		(*ControlFlowContextDef_WhileCtxt)(nil),
	}
}

// Protocol buffer representing a CondContext object.
type CondContextDef struct {
	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// Name of the pred tensor.
	PredName string `protobuf:"bytes,2,opt,name=pred_name,json=predName,proto3" json:"pred_name,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,3,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Branch prediction. 0 or 1.
	Branch int32 `protobuf:"varint,4,opt,name=branch,proto3" json:"branch,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,5,opt,name=values_def,json=valuesDef,proto3" json:"values_def,omitempty"`
	// Contexts contained inside this context (e.g. nested conds).
	NestedContexts       []*ControlFlowContextDef `protobuf:"bytes,6,rep,name=nested_contexts,json=nestedContexts,proto3" json:"nested_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CondContextDef) Reset()         { *m = CondContextDef{} }
func (m *CondContextDef) String() string { return proto.CompactTextString(m) }
func (*CondContextDef) ProtoMessage()    {}
func (*CondContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_64affc5a646d7df1, []int{2}
}

func (m *CondContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CondContextDef.Unmarshal(m, b)
}
func (m *CondContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CondContextDef.Marshal(b, m, deterministic)
}
func (m *CondContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CondContextDef.Merge(m, src)
}
func (m *CondContextDef) XXX_Size() int {
	return xxx_messageInfo_CondContextDef.Size(m)
}
func (m *CondContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_CondContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_CondContextDef proto.InternalMessageInfo

func (m *CondContextDef) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

func (m *CondContextDef) GetPredName() string {
	if m != nil {
		return m.PredName
	}
	return ""
}

func (m *CondContextDef) GetPivotName() string {
	if m != nil {
		return m.PivotName
	}
	return ""
}

func (m *CondContextDef) GetBranch() int32 {
	if m != nil {
		return m.Branch
	}
	return 0
}

func (m *CondContextDef) GetValuesDef() *ValuesDef {
	if m != nil {
		return m.ValuesDef
	}
	return nil
}

func (m *CondContextDef) GetNestedContexts() []*ControlFlowContextDef {
	if m != nil {
		return m.NestedContexts
	}
	return nil
}

// Protocol buffer representing a WhileContext object.
type WhileContextDef struct {
	// Name of the context.
	ContextName string `protobuf:"bytes,1,opt,name=context_name,json=contextName,proto3" json:"context_name,omitempty"`
	// The number of iterations allowed to run in parallel.
	ParallelIterations int32 `protobuf:"varint,2,opt,name=parallel_iterations,json=parallelIterations,proto3" json:"parallel_iterations,omitempty"`
	// Whether backprop is enabled for this while loop.
	BackProp bool `protobuf:"varint,3,opt,name=back_prop,json=backProp,proto3" json:"back_prop,omitempty"`
	// Whether GPU-CPU memory swap is enabled for this loop.
	SwapMemory bool `protobuf:"varint,4,opt,name=swap_memory,json=swapMemory,proto3" json:"swap_memory,omitempty"`
	// Name of the pivot tensor.
	PivotName string `protobuf:"bytes,5,opt,name=pivot_name,json=pivotName,proto3" json:"pivot_name,omitempty"`
	// Name of the pivot_for_pred tensor.
	PivotForPredName string `protobuf:"bytes,6,opt,name=pivot_for_pred_name,json=pivotForPredName,proto3" json:"pivot_for_pred_name,omitempty"`
	// Name of the pivot_for_body tensor.
	PivotForBodyName string `protobuf:"bytes,7,opt,name=pivot_for_body_name,json=pivotForBodyName,proto3" json:"pivot_for_body_name,omitempty"`
	// List of names for exit tensors.
	LoopExitNames []string `protobuf:"bytes,8,rep,name=loop_exit_names,json=loopExitNames,proto3" json:"loop_exit_names,omitempty"`
	// List of names for enter tensors.
	LoopEnterNames []string `protobuf:"bytes,10,rep,name=loop_enter_names,json=loopEnterNames,proto3" json:"loop_enter_names,omitempty"`
	// Values and external values in control flow context.
	ValuesDef *ValuesDef `protobuf:"bytes,9,opt,name=values_def,json=valuesDef,proto3" json:"values_def,omitempty"`
	// Optional name of the maximum_iterations tensor.
	MaximumIterationsName string `protobuf:"bytes,11,opt,name=maximum_iterations_name,json=maximumIterationsName,proto3" json:"maximum_iterations_name,omitempty"`
	// Contexts contained inside this context (e.g. nested whiles).
	NestedContexts       []*ControlFlowContextDef `protobuf:"bytes,12,rep,name=nested_contexts,json=nestedContexts,proto3" json:"nested_contexts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WhileContextDef) Reset()         { *m = WhileContextDef{} }
func (m *WhileContextDef) String() string { return proto.CompactTextString(m) }
func (*WhileContextDef) ProtoMessage()    {}
func (*WhileContextDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_64affc5a646d7df1, []int{3}
}

func (m *WhileContextDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhileContextDef.Unmarshal(m, b)
}
func (m *WhileContextDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhileContextDef.Marshal(b, m, deterministic)
}
func (m *WhileContextDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhileContextDef.Merge(m, src)
}
func (m *WhileContextDef) XXX_Size() int {
	return xxx_messageInfo_WhileContextDef.Size(m)
}
func (m *WhileContextDef) XXX_DiscardUnknown() {
	xxx_messageInfo_WhileContextDef.DiscardUnknown(m)
}

var xxx_messageInfo_WhileContextDef proto.InternalMessageInfo

func (m *WhileContextDef) GetContextName() string {
	if m != nil {
		return m.ContextName
	}
	return ""
}

func (m *WhileContextDef) GetParallelIterations() int32 {
	if m != nil {
		return m.ParallelIterations
	}
	return 0
}

func (m *WhileContextDef) GetBackProp() bool {
	if m != nil {
		return m.BackProp
	}
	return false
}

func (m *WhileContextDef) GetSwapMemory() bool {
	if m != nil {
		return m.SwapMemory
	}
	return false
}

func (m *WhileContextDef) GetPivotName() string {
	if m != nil {
		return m.PivotName
	}
	return ""
}

func (m *WhileContextDef) GetPivotForPredName() string {
	if m != nil {
		return m.PivotForPredName
	}
	return ""
}

func (m *WhileContextDef) GetPivotForBodyName() string {
	if m != nil {
		return m.PivotForBodyName
	}
	return ""
}

func (m *WhileContextDef) GetLoopExitNames() []string {
	if m != nil {
		return m.LoopExitNames
	}
	return nil
}

func (m *WhileContextDef) GetLoopEnterNames() []string {
	if m != nil {
		return m.LoopEnterNames
	}
	return nil
}

func (m *WhileContextDef) GetValuesDef() *ValuesDef {
	if m != nil {
		return m.ValuesDef
	}
	return nil
}

func (m *WhileContextDef) GetMaximumIterationsName() string {
	if m != nil {
		return m.MaximumIterationsName
	}
	return ""
}

func (m *WhileContextDef) GetNestedContexts() []*ControlFlowContextDef {
	if m != nil {
		return m.NestedContexts
	}
	return nil
}

func init() {
	proto.RegisterType((*ValuesDef)(nil), "tensorflow.ValuesDef")
	proto.RegisterMapType((map[string]string)(nil), "tensorflow.ValuesDef.ExternalValuesEntry")
	proto.RegisterType((*ControlFlowContextDef)(nil), "tensorflow.ControlFlowContextDef")
	proto.RegisterType((*CondContextDef)(nil), "tensorflow.CondContextDef")
	proto.RegisterType((*WhileContextDef)(nil), "tensorflow.WhileContextDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/control_flow.proto", fileDescriptor_64affc5a646d7df1)
}

var fileDescriptor_64affc5a646d7df1 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xff, 0x9c, 0x34, 0xfe, 0xec, 0x49, 0x49, 0xca, 0x96, 0x16, 0xab, 0x15, 0x22, 0xcd, 0x01,
	0x05, 0x21, 0x12, 0xa9, 0x20, 0x04, 0x88, 0x0b, 0x09, 0xa9, 0x00, 0x09, 0x14, 0xf9, 0x00, 0x47,
	0x6b, 0x63, 0xaf, 0xa9, 0x15, 0xdb, 0x6b, 0xad, 0x37, 0x89, 0xf3, 0x08, 0xbc, 0x01, 0x0f, 0xc1,
	0x13, 0xf0, 0x64, 0x1c, 0xd1, 0xce, 0xba, 0x49, 0x1c, 0x72, 0xa8, 0xb8, 0x79, 0x7e, 0x7f, 0xd6,
	0x33, 0xbf, 0x59, 0x1b, 0x9e, 0x48, 0x96, 0xe6, 0x5c, 0x84, 0x31, 0x5f, 0x0e, 0x7c, 0x2e, 0xd8,
	0x20, 0x13, 0x5c, 0xf2, 0xe9, 0x3c, 0x1c, 0xf8, 0x3c, 0x95, 0x82, 0xc7, 0x9e, 0xa2, 0xfa, 0x88,
	0x12, 0xd8, 0x88, 0xbb, 0xbf, 0x0c, 0xb0, 0xbf, 0xd0, 0x78, 0xce, 0xf2, 0x77, 0x2c, 0x24, 0xa7,
	0x60, 0x2e, 0xb0, 0x70, 0x8c, 0x4e, 0xbd, 0x67, 0xbb, 0x65, 0x45, 0x5c, 0x68, 0xb3, 0x42, 0x32,
	0x91, 0xd2, 0xd8, 0x2b, 0x05, 0xb5, 0x4e, 0xbd, 0xd7, 0xbc, 0x7c, 0xdc, 0xdf, 0x9c, 0xd5, 0x5f,
	0x9f, 0xd3, 0x1f, 0x97, 0x62, 0x8d, 0x8c, 0x53, 0x29, 0x56, 0x6e, 0x8b, 0x55, 0xc0, 0xb3, 0xb7,
	0x70, 0xbc, 0x47, 0x46, 0x8e, 0xa0, 0x3e, 0x63, 0x2b, 0xc7, 0xe8, 0x18, 0x3d, 0xdb, 0x55, 0x8f,
	0xe4, 0x1e, 0x34, 0xf0, 0x9d, 0x4e, 0x0d, 0x31, 0x5d, 0xbc, 0xae, 0xbd, 0x34, 0xba, 0x3f, 0x0c,
	0x38, 0x19, 0xe9, 0xf9, 0xae, 0x62, 0xbe, 0x54, 0x8f, 0xac, 0x90, 0x6a, 0x90, 0x57, 0x60, 0xfb,
	0x3c, 0x0d, 0x3c, 0x5f, 0x16, 0x12, 0xcf, 0x6a, 0x5e, 0x9e, 0x6d, 0xb7, 0x3a, 0xe2, 0x69, 0xb0,
	0x91, 0xbf, 0xff, 0xcf, 0xb5, 0x94, 0x7c, 0x24, 0x0b, 0x49, 0xde, 0x00, 0x2c, 0xaf, 0xa3, 0x98,
	0x69, 0x6f, 0x0d, 0xbd, 0xe7, 0xdb, 0xde, 0xaf, 0x8a, 0xad, 0x98, 0x6d, 0x34, 0x28, 0xf7, 0xd0,
	0x84, 0x03, 0xe5, 0xeb, 0x7e, 0xaf, 0x41, 0xab, 0xfa, 0x12, 0x72, 0x01, 0x87, 0xbe, 0xae, 0xbc,
	0x94, 0x26, 0xac, 0x1c, 0xb1, 0x59, 0x62, 0x9f, 0x69, 0xc2, 0xc8, 0x39, 0xd8, 0x99, 0x60, 0x81,
	0xe6, 0xf5, 0xb8, 0x96, 0x02, 0x90, 0x7c, 0x00, 0x90, 0x45, 0x0b, 0x5e, 0xba, 0xeb, 0xc8, 0xda,
	0x88, 0x20, 0x7d, 0x0a, 0xe6, 0x54, 0xd0, 0xd4, 0xbf, 0x76, 0x0e, 0x3a, 0x46, 0xaf, 0xe1, 0x96,
	0x15, 0x79, 0x0e, 0xa0, 0x57, 0xe6, 0x05, 0x2c, 0x74, 0x1a, 0x38, 0xcf, 0xc9, 0xde, 0xb5, 0xb9,
	0xf6, 0x62, 0x7d, 0x13, 0x3e, 0x42, 0x3b, 0x65, 0xb9, 0x64, 0x81, 0x57, 0xf6, 0x97, 0x3b, 0x26,
	0x6e, 0xfc, 0x62, 0x27, 0xc6, 0xbf, 0xc3, 0x77, 0x5b, 0xda, 0x59, 0x22, 0x79, 0xf7, 0xe7, 0x01,
	0xb4, 0x77, 0x42, 0xbb, 0x4d, 0x18, 0x03, 0x38, 0xce, 0xa8, 0xa0, 0x71, 0xcc, 0x62, 0x2f, 0x92,
	0x4c, 0x50, 0x19, 0xf1, 0x34, 0xc7, 0x58, 0x1a, 0x2e, 0xb9, 0xa1, 0x3e, 0xac, 0x19, 0x95, 0xde,
	0x94, 0xfa, 0x33, 0x2f, 0x13, 0x3c, 0xc3, 0x7c, 0x2c, 0xd7, 0x52, 0xc0, 0x44, 0xf0, 0x8c, 0x3c,
	0x84, 0x66, 0xbe, 0xa4, 0x99, 0x97, 0xb0, 0x84, 0x8b, 0x15, 0x66, 0x64, 0xb9, 0xa0, 0xa0, 0x4f,
	0x88, 0xec, 0xc4, 0xdb, 0xd8, 0x8d, 0xf7, 0x29, 0x1c, 0x6b, 0x3a, 0xe4, 0xc2, 0xdb, 0x2c, 0xc9,
	0x44, 0xdd, 0x11, 0x52, 0x57, 0x5c, 0x4c, 0x6e, 0x96, 0x55, 0x91, 0x4f, 0x79, 0xb0, 0xd2, 0xf2,
	0xff, 0xab, 0xf2, 0x21, 0x0f, 0x56, 0x28, 0x7f, 0x04, 0xed, 0x98, 0xf3, 0xcc, 0x63, 0x45, 0xa4,
	0x1b, 0xc8, 0x1d, 0x0b, 0xbf, 0xc0, 0x3b, 0x0a, 0x1e, 0x17, 0x11, 0x36, 0x91, 0x93, 0x1e, 0x1c,
	0x69, 0x5d, 0x2a, 0x99, 0x28, 0x85, 0x80, 0xc2, 0x16, 0x0a, 0x15, 0xac, 0x95, 0xd5, 0xb5, 0xdb,
	0xb7, 0x5c, 0xfb, 0x0b, 0xb8, 0x9f, 0xd0, 0x22, 0x4a, 0xe6, 0xc9, 0x56, 0xe4, 0xba, 0xf5, 0x26,
	0xb6, 0x7e, 0x52, 0xd2, 0x9b, 0xd8, 0xb1, 0xff, 0x3d, 0xd7, 0xe5, 0xf0, 0x1f, 0xaf, 0xcb, 0xf0,
	0x12, 0x1c, 0x2e, 0xbe, 0x6d, 0xfb, 0x42, 0x41, 0x13, 0xb6, 0xe4, 0x62, 0x36, 0xbc, 0xbb, 0x75,
	0xc4, 0x44, 0xfd, 0xcc, 0xf2, 0x89, 0xf1, 0xdb, 0x30, 0xa6, 0x26, 0xfe, 0xd9, 0x9e, 0xfd, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x12, 0x90, 0x35, 0x08, 0x05, 0x00, 0x00,
}
