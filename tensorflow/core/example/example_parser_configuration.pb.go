// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/example/example_parser_configuration.proto

package tensorflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	framework "github.com/smoug25/tensorflow_serving_helper/tensorflow/core/framework"
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

type VarLenFeatureProto struct {
	Dtype                   framework.DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	ValuesOutputTensorName  string             `protobuf:"bytes,2,opt,name=values_output_tensor_name,json=valuesOutputTensorName,proto3" json:"values_output_tensor_name,omitempty"`
	IndicesOutputTensorName string             `protobuf:"bytes,3,opt,name=indices_output_tensor_name,json=indicesOutputTensorName,proto3" json:"indices_output_tensor_name,omitempty"`
	ShapesOutputTensorName  string             `protobuf:"bytes,4,opt,name=shapes_output_tensor_name,json=shapesOutputTensorName,proto3" json:"shapes_output_tensor_name,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}           `json:"-"`
	XXX_unrecognized        []byte             `json:"-"`
	XXX_sizecache           int32              `json:"-"`
}

func (m *VarLenFeatureProto) Reset()         { *m = VarLenFeatureProto{} }
func (m *VarLenFeatureProto) String() string { return proto.CompactTextString(m) }
func (*VarLenFeatureProto) ProtoMessage()    {}
func (*VarLenFeatureProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8ee005f857d2cfe, []int{0}
}

func (m *VarLenFeatureProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VarLenFeatureProto.Unmarshal(m, b)
}
func (m *VarLenFeatureProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VarLenFeatureProto.Marshal(b, m, deterministic)
}
func (m *VarLenFeatureProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VarLenFeatureProto.Merge(m, src)
}
func (m *VarLenFeatureProto) XXX_Size() int {
	return xxx_messageInfo_VarLenFeatureProto.Size(m)
}
func (m *VarLenFeatureProto) XXX_DiscardUnknown() {
	xxx_messageInfo_VarLenFeatureProto.DiscardUnknown(m)
}

var xxx_messageInfo_VarLenFeatureProto proto.InternalMessageInfo

func (m *VarLenFeatureProto) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *VarLenFeatureProto) GetValuesOutputTensorName() string {
	if m != nil {
		return m.ValuesOutputTensorName
	}
	return ""
}

func (m *VarLenFeatureProto) GetIndicesOutputTensorName() string {
	if m != nil {
		return m.IndicesOutputTensorName
	}
	return ""
}

func (m *VarLenFeatureProto) GetShapesOutputTensorName() string {
	if m != nil {
		return m.ShapesOutputTensorName
	}
	return ""
}

type FixedLenFeatureProto struct {
	Dtype                  framework.DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                  *framework.TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	DefaultValue           *framework.TensorProto      `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	ValuesOutputTensorName string                      `protobuf:"bytes,4,opt,name=values_output_tensor_name,json=valuesOutputTensorName,proto3" json:"values_output_tensor_name,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                    `json:"-"`
	XXX_unrecognized       []byte                      `json:"-"`
	XXX_sizecache          int32                       `json:"-"`
}

func (m *FixedLenFeatureProto) Reset()         { *m = FixedLenFeatureProto{} }
func (m *FixedLenFeatureProto) String() string { return proto.CompactTextString(m) }
func (*FixedLenFeatureProto) ProtoMessage()    {}
func (*FixedLenFeatureProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8ee005f857d2cfe, []int{1}
}

func (m *FixedLenFeatureProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedLenFeatureProto.Unmarshal(m, b)
}
func (m *FixedLenFeatureProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedLenFeatureProto.Marshal(b, m, deterministic)
}
func (m *FixedLenFeatureProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedLenFeatureProto.Merge(m, src)
}
func (m *FixedLenFeatureProto) XXX_Size() int {
	return xxx_messageInfo_FixedLenFeatureProto.Size(m)
}
func (m *FixedLenFeatureProto) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedLenFeatureProto.DiscardUnknown(m)
}

var xxx_messageInfo_FixedLenFeatureProto proto.InternalMessageInfo

func (m *FixedLenFeatureProto) GetDtype() framework.DataType {
	if m != nil {
		return m.Dtype
	}
	return framework.DataType_DT_INVALID
}

func (m *FixedLenFeatureProto) GetShape() *framework.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *FixedLenFeatureProto) GetDefaultValue() *framework.TensorProto {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *FixedLenFeatureProto) GetValuesOutputTensorName() string {
	if m != nil {
		return m.ValuesOutputTensorName
	}
	return ""
}

type FeatureConfiguration struct {
	// Types that are valid to be assigned to Config:
	//	*FeatureConfiguration_FixedLenFeature
	//	*FeatureConfiguration_VarLenFeature
	Config               isFeatureConfiguration_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FeatureConfiguration) Reset()         { *m = FeatureConfiguration{} }
func (m *FeatureConfiguration) String() string { return proto.CompactTextString(m) }
func (*FeatureConfiguration) ProtoMessage()    {}
func (*FeatureConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8ee005f857d2cfe, []int{2}
}

func (m *FeatureConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureConfiguration.Unmarshal(m, b)
}
func (m *FeatureConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureConfiguration.Marshal(b, m, deterministic)
}
func (m *FeatureConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureConfiguration.Merge(m, src)
}
func (m *FeatureConfiguration) XXX_Size() int {
	return xxx_messageInfo_FeatureConfiguration.Size(m)
}
func (m *FeatureConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureConfiguration proto.InternalMessageInfo

type isFeatureConfiguration_Config interface {
	isFeatureConfiguration_Config()
}

type FeatureConfiguration_FixedLenFeature struct {
	FixedLenFeature *FixedLenFeatureProto `protobuf:"bytes,1,opt,name=fixed_len_feature,json=fixedLenFeature,proto3,oneof"`
}

type FeatureConfiguration_VarLenFeature struct {
	VarLenFeature *VarLenFeatureProto `protobuf:"bytes,2,opt,name=var_len_feature,json=varLenFeature,proto3,oneof"`
}

func (*FeatureConfiguration_FixedLenFeature) isFeatureConfiguration_Config() {}

func (*FeatureConfiguration_VarLenFeature) isFeatureConfiguration_Config() {}

func (m *FeatureConfiguration) GetConfig() isFeatureConfiguration_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *FeatureConfiguration) GetFixedLenFeature() *FixedLenFeatureProto {
	if x, ok := m.GetConfig().(*FeatureConfiguration_FixedLenFeature); ok {
		return x.FixedLenFeature
	}
	return nil
}

func (m *FeatureConfiguration) GetVarLenFeature() *VarLenFeatureProto {
	if x, ok := m.GetConfig().(*FeatureConfiguration_VarLenFeature); ok {
		return x.VarLenFeature
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeatureConfiguration) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeatureConfiguration_FixedLenFeature)(nil),
		(*FeatureConfiguration_VarLenFeature)(nil),
	}
}

type ExampleParserConfiguration struct {
	FeatureMap           map[string]*FeatureConfiguration `protobuf:"bytes,1,rep,name=feature_map,json=featureMap,proto3" json:"feature_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ExampleParserConfiguration) Reset()         { *m = ExampleParserConfiguration{} }
func (m *ExampleParserConfiguration) String() string { return proto.CompactTextString(m) }
func (*ExampleParserConfiguration) ProtoMessage()    {}
func (*ExampleParserConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8ee005f857d2cfe, []int{3}
}

func (m *ExampleParserConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExampleParserConfiguration.Unmarshal(m, b)
}
func (m *ExampleParserConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExampleParserConfiguration.Marshal(b, m, deterministic)
}
func (m *ExampleParserConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExampleParserConfiguration.Merge(m, src)
}
func (m *ExampleParserConfiguration) XXX_Size() int {
	return xxx_messageInfo_ExampleParserConfiguration.Size(m)
}
func (m *ExampleParserConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_ExampleParserConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_ExampleParserConfiguration proto.InternalMessageInfo

func (m *ExampleParserConfiguration) GetFeatureMap() map[string]*FeatureConfiguration {
	if m != nil {
		return m.FeatureMap
	}
	return nil
}

func init() {
	proto.RegisterType((*VarLenFeatureProto)(nil), "tensorflow.VarLenFeatureProto")
	proto.RegisterType((*FixedLenFeatureProto)(nil), "tensorflow.FixedLenFeatureProto")
	proto.RegisterType((*FeatureConfiguration)(nil), "tensorflow.FeatureConfiguration")
	proto.RegisterType((*ExampleParserConfiguration)(nil), "tensorflow.ExampleParserConfiguration")
	proto.RegisterMapType((map[string]*FeatureConfiguration)(nil), "tensorflow.ExampleParserConfiguration.FeatureMapEntry")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/example/example_parser_configuration.proto", fileDescriptor_c8ee005f857d2cfe)
}

var fileDescriptor_c8ee005f857d2cfe = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x9f, 0xdb, 0x75, 0x62, 0xaf, 0x8c, 0x42, 0x34, 0x6d, 0x5d, 0x84, 0x50, 0xd4, 0x53, 0xc5,
	0x21, 0x95, 0x82, 0x98, 0x18, 0x20, 0x21, 0x15, 0x36, 0xf5, 0x00, 0xa3, 0xca, 0xa6, 0x21, 0x4e,
	0x96, 0xd7, 0xbe, 0xa4, 0xd1, 0x92, 0xd8, 0x72, 0x9c, 0x6e, 0xfd, 0x6a, 0x7c, 0x10, 0xbe, 0x06,
	0x07, 0x2e, 0x1c, 0x51, 0xec, 0x68, 0x4b, 0xd7, 0x6c, 0x97, 0xed, 0xd4, 0xca, 0xef, 0xf7, 0xc7,
	0xbf, 0x5f, 0xf2, 0x02, 0x61, 0x18, 0xa9, 0x59, 0x7e, 0xee, 0x4e, 0x78, 0x32, 0xc8, 0x12, 0x9e,
	0x87, 0xde, 0xdb, 0x81, 0xc2, 0x34, 0xe3, 0x32, 0x88, 0xf9, 0x25, 0xcd, 0x50, 0xce, 0xa3, 0x34,
	0xa4, 0x33, 0x8c, 0x05, 0xca, 0xca, 0x64, 0x30, 0xe1, 0x12, 0x07, 0x78, 0xc5, 0x12, 0x11, 0x5f,
	0xff, 0x52, 0xc1, 0x64, 0x86, 0x92, 0x4e, 0x78, 0x1a, 0x44, 0x61, 0x2e, 0x99, 0x8a, 0x78, 0xea,
	0x0a, 0xc9, 0x15, 0xb7, 0xe0, 0x86, 0x6b, 0xff, 0x7c, 0x90, 0x69, 0x20, 0x59, 0x82, 0x97, 0x5c,
	0x5e, 0x94, 0x13, 0x9a, 0xcd, 0x98, 0x40, 0x63, 0x63, 0x9f, 0x3c, 0xaa, 0x74, 0x29, 0xea, 0x3f,
	0x96, 0xe8, 0x42, 0x60, 0x66, 0x34, 0x7b, 0x7f, 0x09, 0x58, 0x67, 0x4c, 0x7e, 0xc5, 0xf4, 0x08,
	0x99, 0xca, 0x25, 0x8e, 0x75, 0x4d, 0xaf, 0xa1, 0x35, 0x2d, 0x60, 0x5d, 0xe2, 0x90, 0xfe, 0x33,
	0x6f, 0xdb, 0xbd, 0x51, 0x73, 0xbf, 0x30, 0xc5, 0x4e, 0x17, 0x02, 0x7d, 0x03, 0xb1, 0x0e, 0x60,
	0x6f, 0xce, 0xe2, 0x1c, 0x33, 0xca, 0x73, 0x25, 0x72, 0x45, 0xcb, 0x3e, 0x52, 0x96, 0x60, 0xb7,
	0xe1, 0x90, 0xfe, 0xa6, 0xbf, 0x63, 0x00, 0xdf, 0xf5, 0xfc, 0x54, 0x8f, 0x8f, 0x59, 0x82, 0xd6,
	0x07, 0xb0, 0xa3, 0x74, 0x1a, 0x4d, 0xea, 0xb9, 0x4d, 0xcd, 0xdd, 0x2d, 0x11, 0x2b, 0xe4, 0x03,
	0xd8, 0xd3, 0x95, 0xd7, 0x72, 0xd7, 0x8d, 0xaf, 0x01, 0xdc, 0xa6, 0xf6, 0xfe, 0x10, 0xd8, 0x3e,
	0x8a, 0xae, 0x70, 0xfa, 0x90, 0xdc, 0x1e, 0xb4, 0xb4, 0xbc, 0xce, 0xd8, 0xf6, 0x5e, 0x56, 0xb1,
	0xc6, 0xeb, 0xa4, 0x18, 0x6b, 0x61, 0xdf, 0x40, 0xad, 0x8f, 0xb0, 0x35, 0xc5, 0x80, 0xe5, 0xb1,
	0xa2, 0xba, 0x12, 0x9d, 0xb1, 0xed, 0xed, 0xae, 0x72, 0x0d, 0xed, 0x69, 0x89, 0x3e, 0x2b, 0xc0,
	0xf7, 0x37, 0xbd, 0x7e, 0x5f, 0xd3, 0xbd, 0x5f, 0x45, 0x62, 0x93, 0xf4, 0x73, 0x75, 0x2d, 0xac,
	0x63, 0x78, 0x11, 0x14, 0x4d, 0xd0, 0x18, 0x53, 0x1a, 0x18, 0x84, 0x4e, 0xdf, 0xf6, 0x9c, 0xea,
	0xad, 0xea, 0xea, 0x1a, 0xad, 0xf9, 0x9d, 0x60, 0xf9, 0xdc, 0x1a, 0x41, 0x67, 0xce, 0xe4, 0x92,
	0x9a, 0xe9, 0xe7, 0x55, 0x55, 0x6d, 0xf5, 0x95, 0x1b, 0xad, 0xf9, 0x5b, 0xf3, 0xea, 0xe9, 0xf0,
	0x09, 0x6c, 0x98, 0x0d, 0xee, 0xfd, 0x26, 0x60, 0x1f, 0x9a, 0xdd, 0x1e, 0xeb, 0xd5, 0x5e, 0x8e,
	0xf0, 0x03, 0xda, 0xa5, 0x15, 0x4d, 0x98, 0xe8, 0x12, 0xa7, 0xd9, 0x6f, 0x7b, 0xfb, 0x55, 0xbb,
	0xbb, 0xc9, 0x6e, 0xe9, 0xf6, 0x8d, 0x89, 0xc3, 0x54, 0xc9, 0x85, 0x0f, 0xc1, 0xf5, 0x81, 0x4d,
	0xa1, 0x73, 0x6b, 0x6c, 0x3d, 0x87, 0xe6, 0x05, 0x2e, 0x74, 0x41, 0x9b, 0x7e, 0xf1, 0xd7, 0xda,
	0x87, 0x96, 0x79, 0x94, 0x8d, 0x9a, 0xd2, 0x6a, 0x1a, 0xf7, 0x0d, 0xfc, 0x7d, 0xe3, 0x1d, 0x19,
	0x7e, 0x82, 0x1d, 0x2e, 0xc3, 0x2a, 0xa3, 0xfc, 0x84, 0x0d, 0x9d, 0xbb, 0xaf, 0xac, 0xeb, 0xca,
	0xc6, 0xe4, 0x1f, 0x21, 0xe7, 0x1b, 0x7a, 0x8b, 0xdf, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x81,
	0xf0, 0xc0, 0x22, 0x40, 0x05, 0x00, 0x00,
}
