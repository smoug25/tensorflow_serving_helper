// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework/op_def.proto

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

// Defines an operation. A NodeDef in a GraphDef specifies an Op by
// using the "op" field which should match the name of a OpDef.
// LINT.IfChange
type OpDef struct {
	// Op names starting with an underscore are reserved for internal use.
	// Names should be CamelCase and match the regexp "[A-Z][a-zA-Z0-9>_]*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the input(s).
	InputArg []*OpDef_ArgDef `protobuf:"bytes,2,rep,name=input_arg,json=inputArg,proto3" json:"input_arg,omitempty"`
	// Description of the output(s).
	OutputArg []*OpDef_ArgDef `protobuf:"bytes,3,rep,name=output_arg,json=outputArg,proto3" json:"output_arg,omitempty"`
	// Named control outputs for this operation. Useful only for composite
	// operations (i.e. functions) which want to name different control outputs.
	ControlOutput []string         `protobuf:"bytes,20,rep,name=control_output,json=controlOutput,proto3" json:"control_output,omitempty"`
	Attr          []*OpDef_AttrDef `protobuf:"bytes,4,rep,name=attr,proto3" json:"attr,omitempty"`
	// Optional deprecation based on GraphDef versions.
	Deprecation *OpDeprecation `protobuf:"bytes,8,opt,name=deprecation,proto3" json:"deprecation,omitempty"`
	// One-line human-readable description of what the Op does.
	Summary string `protobuf:"bytes,5,opt,name=summary,proto3" json:"summary,omitempty"`
	// Additional, longer human-readable description of what the Op does.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// True if the operation is commutative ("op(a,b) == op(b,a)" for all inputs)
	IsCommutative bool `protobuf:"varint,18,opt,name=is_commutative,json=isCommutative,proto3" json:"is_commutative,omitempty"`
	// If is_aggregate is true, then this operation accepts N >= 2
	// inputs and produces 1 output all of the same type.  Should be
	// associative and commutative, and produce output with the same
	// shape as the input.  The optimizer may replace an aggregate op
	// taking input from multiple devices with a tree of aggregate ops
	// that aggregate locally within each device (and possibly within
	// groups of nearby devices) before communicating.
	// TODO(josh11b): Implement that optimization.
	IsAggregate bool `protobuf:"varint,16,opt,name=is_aggregate,json=isAggregate,proto3" json:"is_aggregate,omitempty"`
	// Ops are marked as stateful if their behavior depends on some state beyond
	// their input tensors (e.g. variable reading op) or if they have
	// a side-effect (e.g. printing or asserting ops). Equivalently, stateless ops
	// must always produce the same output for the same input and have
	// no side-effects.
	//
	// By default Ops may be moved between devices.  Stateful ops should
	// either not be moved, or should only be moved if that state can also
	// be moved (e.g. via some sort of save / restore).
	// Stateful ops are guaranteed to never be optimized away by Common
	// Subexpression Elimination (CSE).
	IsStateful bool `protobuf:"varint,17,opt,name=is_stateful,json=isStateful,proto3" json:"is_stateful,omitempty"`
	// By default, all inputs to an Op must be initialized Tensors.  Ops
	// that may initialize tensors for the first time should set this
	// field to true, to allow the Op to take an uninitialized Tensor as
	// input.
	AllowsUninitializedInput bool     `protobuf:"varint,19,opt,name=allows_uninitialized_input,json=allowsUninitializedInput,proto3" json:"allows_uninitialized_input,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *OpDef) Reset()         { *m = OpDef{} }
func (m *OpDef) String() string { return proto.CompactTextString(m) }
func (*OpDef) ProtoMessage()    {}
func (*OpDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_705e263509f7bee8, []int{0}
}

func (m *OpDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpDef.Unmarshal(m, b)
}
func (m *OpDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpDef.Marshal(b, m, deterministic)
}
func (m *OpDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpDef.Merge(m, src)
}
func (m *OpDef) XXX_Size() int {
	return xxx_messageInfo_OpDef.Size(m)
}
func (m *OpDef) XXX_DiscardUnknown() {
	xxx_messageInfo_OpDef.DiscardUnknown(m)
}

var xxx_messageInfo_OpDef proto.InternalMessageInfo

func (m *OpDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OpDef) GetInputArg() []*OpDef_ArgDef {
	if m != nil {
		return m.InputArg
	}
	return nil
}

func (m *OpDef) GetOutputArg() []*OpDef_ArgDef {
	if m != nil {
		return m.OutputArg
	}
	return nil
}

func (m *OpDef) GetControlOutput() []string {
	if m != nil {
		return m.ControlOutput
	}
	return nil
}

func (m *OpDef) GetAttr() []*OpDef_AttrDef {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *OpDef) GetDeprecation() *OpDeprecation {
	if m != nil {
		return m.Deprecation
	}
	return nil
}

func (m *OpDef) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *OpDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OpDef) GetIsCommutative() bool {
	if m != nil {
		return m.IsCommutative
	}
	return false
}

func (m *OpDef) GetIsAggregate() bool {
	if m != nil {
		return m.IsAggregate
	}
	return false
}

func (m *OpDef) GetIsStateful() bool {
	if m != nil {
		return m.IsStateful
	}
	return false
}

func (m *OpDef) GetAllowsUninitializedInput() bool {
	if m != nil {
		return m.AllowsUninitializedInput
	}
	return false
}

// For describing inputs and outputs.
type OpDef_ArgDef struct {
	// Name for the input/output.  Should match the regexp "[a-z][a-z0-9_]*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Human readable description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Describes the type of one or more tensors that are accepted/produced
	// by this input/output arg.  The only legal combinations are:
	// * For a single tensor: either the "type" field is set or the
	//   "type_attr" field is set to the name of an attr with type "type".
	// * For a sequence of tensors with the same type: the "number_attr"
	//   field will be set to the name of an attr with type "int", and
	//   either the "type" or "type_attr" field will be set as for
	//   single tensors.
	// * For a sequence of tensors, the "type_list_attr" field will be set
	//   to the name of an attr with type "list(type)".
	Type       DataType `protobuf:"varint,3,opt,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	TypeAttr   string   `protobuf:"bytes,4,opt,name=type_attr,json=typeAttr,proto3" json:"type_attr,omitempty"`
	NumberAttr string   `protobuf:"bytes,5,opt,name=number_attr,json=numberAttr,proto3" json:"number_attr,omitempty"`
	// If specified, attr must have type "list(type)", and none of
	// type, type_attr, and number_attr may be specified.
	TypeListAttr string `protobuf:"bytes,6,opt,name=type_list_attr,json=typeListAttr,proto3" json:"type_list_attr,omitempty"`
	// For inputs: if true, the inputs are required to be refs.
	//   By default, inputs can be either refs or non-refs.
	// For outputs: if true, outputs are refs, otherwise they are not.
	IsRef                bool     `protobuf:"varint,16,opt,name=is_ref,json=isRef,proto3" json:"is_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpDef_ArgDef) Reset()         { *m = OpDef_ArgDef{} }
func (m *OpDef_ArgDef) String() string { return proto.CompactTextString(m) }
func (*OpDef_ArgDef) ProtoMessage()    {}
func (*OpDef_ArgDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_705e263509f7bee8, []int{0, 0}
}

func (m *OpDef_ArgDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpDef_ArgDef.Unmarshal(m, b)
}
func (m *OpDef_ArgDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpDef_ArgDef.Marshal(b, m, deterministic)
}
func (m *OpDef_ArgDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpDef_ArgDef.Merge(m, src)
}
func (m *OpDef_ArgDef) XXX_Size() int {
	return xxx_messageInfo_OpDef_ArgDef.Size(m)
}
func (m *OpDef_ArgDef) XXX_DiscardUnknown() {
	xxx_messageInfo_OpDef_ArgDef.DiscardUnknown(m)
}

var xxx_messageInfo_OpDef_ArgDef proto.InternalMessageInfo

func (m *OpDef_ArgDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OpDef_ArgDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OpDef_ArgDef) GetType() DataType {
	if m != nil {
		return m.Type
	}
	return DataType_DT_INVALID
}

func (m *OpDef_ArgDef) GetTypeAttr() string {
	if m != nil {
		return m.TypeAttr
	}
	return ""
}

func (m *OpDef_ArgDef) GetNumberAttr() string {
	if m != nil {
		return m.NumberAttr
	}
	return ""
}

func (m *OpDef_ArgDef) GetTypeListAttr() string {
	if m != nil {
		return m.TypeListAttr
	}
	return ""
}

func (m *OpDef_ArgDef) GetIsRef() bool {
	if m != nil {
		return m.IsRef
	}
	return false
}

// Description of the graph-construction-time configuration of this
// Op.  That is to say, this describes the attr fields that will
// be specified in the NodeDef.
type OpDef_AttrDef struct {
	// A descriptive name for the argument.  May be used, e.g. by the
	// Python client, as a keyword argument name, and so should match
	// the regexp "[a-z][a-z0-9_]+".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// One of the type names from attr_value.proto ("string", "list(string)",
	// "int", etc.).
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// A reasonable default for this attribute if the user does not supply
	// a value.  If not specified, the user must supply a value.
	DefaultValue *AttrValue `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Human-readable description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// For type == "int", this is a minimum value.  For "list(___)"
	// types, this is the minimum length.
	HasMinimum bool  `protobuf:"varint,5,opt,name=has_minimum,json=hasMinimum,proto3" json:"has_minimum,omitempty"`
	Minimum    int64 `protobuf:"varint,6,opt,name=minimum,proto3" json:"minimum,omitempty"`
	// The set of allowed values.  Has type that is the "list" version
	// of the "type" field above (uses the "list" field of AttrValue).
	// If type == "type" or "list(type)" above, then the "type" field
	// of "allowed_values.list" has the set of allowed DataTypes.
	// If type == "string" or "list(string)", then the "s" field of
	// "allowed_values.list" has the set of allowed strings.
	AllowedValues        *AttrValue `protobuf:"bytes,7,opt,name=allowed_values,json=allowedValues,proto3" json:"allowed_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OpDef_AttrDef) Reset()         { *m = OpDef_AttrDef{} }
func (m *OpDef_AttrDef) String() string { return proto.CompactTextString(m) }
func (*OpDef_AttrDef) ProtoMessage()    {}
func (*OpDef_AttrDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_705e263509f7bee8, []int{0, 1}
}

func (m *OpDef_AttrDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpDef_AttrDef.Unmarshal(m, b)
}
func (m *OpDef_AttrDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpDef_AttrDef.Marshal(b, m, deterministic)
}
func (m *OpDef_AttrDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpDef_AttrDef.Merge(m, src)
}
func (m *OpDef_AttrDef) XXX_Size() int {
	return xxx_messageInfo_OpDef_AttrDef.Size(m)
}
func (m *OpDef_AttrDef) XXX_DiscardUnknown() {
	xxx_messageInfo_OpDef_AttrDef.DiscardUnknown(m)
}

var xxx_messageInfo_OpDef_AttrDef proto.InternalMessageInfo

func (m *OpDef_AttrDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OpDef_AttrDef) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OpDef_AttrDef) GetDefaultValue() *AttrValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *OpDef_AttrDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OpDef_AttrDef) GetHasMinimum() bool {
	if m != nil {
		return m.HasMinimum
	}
	return false
}

func (m *OpDef_AttrDef) GetMinimum() int64 {
	if m != nil {
		return m.Minimum
	}
	return 0
}

func (m *OpDef_AttrDef) GetAllowedValues() *AttrValue {
	if m != nil {
		return m.AllowedValues
	}
	return nil
}

// Information about version-dependent deprecation of an op
type OpDeprecation struct {
	// First GraphDef version at which the op is disallowed.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Explanation of why it was deprecated and what to use instead.
	Explanation          string   `protobuf:"bytes,2,opt,name=explanation,proto3" json:"explanation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpDeprecation) Reset()         { *m = OpDeprecation{} }
func (m *OpDeprecation) String() string { return proto.CompactTextString(m) }
func (*OpDeprecation) ProtoMessage()    {}
func (*OpDeprecation) Descriptor() ([]byte, []int) {
	return fileDescriptor_705e263509f7bee8, []int{1}
}

func (m *OpDeprecation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpDeprecation.Unmarshal(m, b)
}
func (m *OpDeprecation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpDeprecation.Marshal(b, m, deterministic)
}
func (m *OpDeprecation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpDeprecation.Merge(m, src)
}
func (m *OpDeprecation) XXX_Size() int {
	return xxx_messageInfo_OpDeprecation.Size(m)
}
func (m *OpDeprecation) XXX_DiscardUnknown() {
	xxx_messageInfo_OpDeprecation.DiscardUnknown(m)
}

var xxx_messageInfo_OpDeprecation proto.InternalMessageInfo

func (m *OpDeprecation) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *OpDeprecation) GetExplanation() string {
	if m != nil {
		return m.Explanation
	}
	return ""
}

// A collection of OpDefs
type OpList struct {
	Op                   []*OpDef `protobuf:"bytes,1,rep,name=op,proto3" json:"op,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpList) Reset()         { *m = OpList{} }
func (m *OpList) String() string { return proto.CompactTextString(m) }
func (*OpList) ProtoMessage()    {}
func (*OpList) Descriptor() ([]byte, []int) {
	return fileDescriptor_705e263509f7bee8, []int{2}
}

func (m *OpList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpList.Unmarshal(m, b)
}
func (m *OpList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpList.Marshal(b, m, deterministic)
}
func (m *OpList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpList.Merge(m, src)
}
func (m *OpList) XXX_Size() int {
	return xxx_messageInfo_OpList.Size(m)
}
func (m *OpList) XXX_DiscardUnknown() {
	xxx_messageInfo_OpList.DiscardUnknown(m)
}

var xxx_messageInfo_OpList proto.InternalMessageInfo

func (m *OpList) GetOp() []*OpDef {
	if m != nil {
		return m.Op
	}
	return nil
}

func init() {
	proto.RegisterType((*OpDef)(nil), "tensorflow.OpDef")
	proto.RegisterType((*OpDef_ArgDef)(nil), "tensorflow.OpDef.ArgDef")
	proto.RegisterType((*OpDef_AttrDef)(nil), "tensorflow.OpDef.AttrDef")
	proto.RegisterType((*OpDeprecation)(nil), "tensorflow.OpDeprecation")
	proto.RegisterType((*OpList)(nil), "tensorflow.OpList")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework/op_def.proto", fileDescriptor_705e263509f7bee8)
}

var fileDescriptor_705e263509f7bee8 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0x13, 0x3b,
	0x14, 0xd5, 0xe4, 0x3b, 0x37, 0x4d, 0xf4, 0xea, 0xd7, 0x4a, 0x7e, 0x79, 0x8b, 0xa6, 0xd5, 0x7b,
	0x52, 0x24, 0x20, 0x91, 0x82, 0x2a, 0x24, 0xe8, 0xa6, 0xa5, 0x1b, 0x04, 0xa8, 0xd5, 0x94, 0x8f,
	0xa5, 0xe5, 0x26, 0x77, 0x26, 0x16, 0x33, 0xe3, 0x91, 0xed, 0x49, 0x29, 0xbf, 0x80, 0xff, 0xc9,
	0x5f, 0x60, 0xc1, 0x12, 0xd9, 0x9e, 0xb4, 0x43, 0x1b, 0xc1, 0xa6, 0xab, 0xc4, 0xe7, 0x9e, 0x7b,
	0xc6, 0xe7, 0x5c, 0xdb, 0x70, 0x11, 0x0b, 0xb3, 0x2c, 0x2e, 0x27, 0x73, 0x99, 0x4e, 0x75, 0x2a,
	0x8b, 0x78, 0x76, 0x38, 0x35, 0x98, 0x69, 0xa9, 0xa2, 0x44, 0x5e, 0x31, 0x8d, 0x6a, 0x25, 0xb2,
	0x98, 0x2d, 0x31, 0xc9, 0x51, 0x55, 0x2a, 0xd3, 0xb9, 0x54, 0x38, 0x8d, 0x14, 0x4f, 0xf1, 0x4a,
	0xaa, 0x4f, 0x53, 0x99, 0xb3, 0x05, 0x46, 0x93, 0x5c, 0x49, 0x23, 0x09, 0xdc, 0xf2, 0x86, 0x1f,
	0x1f, 0xe8, 0x03, 0xdc, 0x18, 0xc5, 0x56, 0x3c, 0x29, 0xd0, 0x7f, 0x64, 0x18, 0x3e, 0x90, 0xb0,
	0xb9, 0xce, 0x51, 0x7b, 0xcd, 0x83, 0xef, 0x6d, 0x68, 0x9e, 0xe5, 0xa7, 0x18, 0x11, 0x02, 0x8d,
	0x8c, 0xa7, 0x48, 0x83, 0x51, 0x30, 0xee, 0x86, 0xee, 0x3f, 0x39, 0x84, 0xae, 0xc8, 0xf2, 0xc2,
	0x30, 0xae, 0x62, 0x5a, 0x1b, 0xd5, 0xc7, 0xbd, 0x19, 0x9d, 0xdc, 0x0a, 0x4f, 0x5c, 0xe7, 0xe4,
	0x58, 0xc5, 0xa7, 0x18, 0x85, 0x1d, 0x47, 0x3d, 0x56, 0x31, 0x79, 0x06, 0x20, 0x0b, 0xb3, 0xee,
	0xab, 0xff, 0xa1, 0xaf, 0xeb, 0xb9, 0xb6, 0xf1, 0x7f, 0x18, 0xcc, 0x65, 0x66, 0x94, 0x4c, 0x98,
	0x07, 0xe9, 0xce, 0xa8, 0x3e, 0xee, 0x86, 0xfd, 0x12, 0x3d, 0x73, 0x20, 0x79, 0x02, 0x0d, 0x1b,
	0x0e, 0x6d, 0x38, 0xe5, 0x7f, 0x36, 0x28, 0x1b, 0xa3, 0xac, 0xb4, 0xa3, 0x91, 0x17, 0xd0, 0x5b,
	0x60, 0xae, 0x70, 0xce, 0x8d, 0x90, 0x19, 0xed, 0x8c, 0x82, 0x4d, 0x5d, 0x37, 0x84, 0xb0, 0xca,
	0x26, 0x14, 0xda, 0xba, 0x48, 0x53, 0xae, 0xae, 0x69, 0xd3, 0x25, 0xb3, 0x5e, 0x92, 0x91, 0x95,
	0xd5, 0x73, 0x25, 0x72, 0x27, 0xdb, 0x72, 0xd5, 0x2a, 0x64, 0xed, 0x08, 0xcd, 0xe6, 0x32, 0x4d,
	0x0b, 0xc3, 0x8d, 0x58, 0x21, 0x25, 0xa3, 0x60, 0xdc, 0x09, 0xfb, 0x42, 0xbf, 0xbc, 0x05, 0xc9,
	0x3e, 0x6c, 0x09, 0xcd, 0x78, 0x1c, 0x2b, 0x8c, 0xb9, 0x41, 0xfa, 0x97, 0x23, 0xf5, 0x84, 0x3e,
	0x5e, 0x43, 0x64, 0x0f, 0x7a, 0x42, 0x33, 0x6d, 0xb8, 0xc1, 0xa8, 0x48, 0xe8, 0xb6, 0x63, 0x80,
	0xd0, 0x17, 0x25, 0x42, 0x8e, 0x60, 0xc8, 0x93, 0x44, 0x5e, 0x69, 0x56, 0x64, 0x22, 0x13, 0x46,
	0xf0, 0x44, 0x7c, 0xc1, 0x05, 0x73, 0x33, 0xa1, 0x7f, 0x3b, 0x3e, 0xf5, 0x8c, 0xf7, 0x55, 0xc2,
	0x2b, 0x5b, 0x1f, 0x7e, 0x0b, 0xa0, 0xe5, 0xa7, 0xb1, 0xf1, 0x18, 0xdc, 0x71, 0x5a, 0xbb, 0xef,
	0x74, 0x0c, 0x0d, 0x7b, 0xaa, 0x68, 0x7d, 0x14, 0x8c, 0x07, 0xb3, 0x9d, 0x6a, 0xb6, 0xa7, 0xdc,
	0xf0, 0x77, 0xd7, 0x39, 0x86, 0x8e, 0x41, 0xfe, 0x85, 0xae, 0xfd, 0x65, 0xe5, 0x00, 0xad, 0x52,
	0xc7, 0x02, 0x76, 0x64, 0xd6, 0x66, 0x56, 0xa4, 0x97, 0xa8, 0x7c, 0xd9, 0x07, 0x0e, 0x1e, 0x72,
	0x84, 0xff, 0x60, 0xe0, 0xba, 0x13, 0xa1, 0x8d, 0xe7, 0xf8, 0xd8, 0xb7, 0x2c, 0xfa, 0x46, 0x68,
	0xe3, 0x58, 0xbb, 0xd0, 0x12, 0x9a, 0x29, 0x8c, 0xca, 0x28, 0x9b, 0x42, 0x87, 0x18, 0x0d, 0xbf,
	0xd6, 0xa0, 0x5d, 0x9e, 0x8c, 0x8d, 0x36, 0x49, 0x69, 0xc2, 0xfb, 0xf3, 0xdb, 0x7d, 0x0e, 0xfd,
	0x05, 0x46, 0xbc, 0x48, 0x8c, 0xbf, 0x8a, 0xce, 0x61, 0x6f, 0xb6, 0x5b, 0x75, 0x68, 0x35, 0x3f,
	0xd8, 0x62, 0xb8, 0x55, 0x72, 0xdd, 0xea, 0x6e, 0x6c, 0x8d, 0xfb, 0xb1, 0xed, 0x41, 0x6f, 0xc9,
	0x35, 0x4b, 0x45, 0x26, 0xd2, 0x22, 0x75, 0x7e, 0x3b, 0x21, 0x2c, 0xb9, 0x7e, 0xeb, 0x11, 0x7b,
	0xfa, 0xd6, 0x45, 0x6b, 0xb4, 0x1e, 0xae, 0x97, 0xe4, 0x08, 0x06, 0x6e, 0x9c, 0xb8, 0xf0, 0x1b,
	0xd3, 0xb4, 0xfd, 0xbb, 0x9d, 0xf5, 0x4b, 0xb2, 0x5b, 0xe9, 0x83, 0xd7, 0xd0, 0xff, 0xe5, 0xcc,
	0xdb, 0x0f, 0xad, 0x50, 0x69, 0xbb, 0x4f, 0x1b, 0x49, 0x33, 0x5c, 0x2f, 0xad, 0x0b, 0xfc, 0x9c,
	0x27, 0x3c, 0xe3, 0xd5, 0xe1, 0x57, 0xa0, 0x83, 0x47, 0xd0, 0x3a, 0xcb, 0x6d, 0xf8, 0x64, 0x1f,
	0x6a, 0x32, 0xa7, 0x81, 0xbb, 0x96, 0xdb, 0xf7, 0xae, 0x65, 0x58, 0x93, 0xf9, 0xc9, 0x63, 0xa0,
	0x52, 0xc5, 0xd5, 0xda, 0xcd, 0xc3, 0x74, 0xd2, 0x73, 0xb4, 0x73, 0xfb, 0x30, 0xe9, 0xf3, 0xe0,
	0x47, 0x10, 0x5c, 0xb6, 0xdc, 0x2b, 0xf5, 0xf4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xf1,
	0x60, 0x72, 0xb5, 0x05, 0x00, 0x00,
}
