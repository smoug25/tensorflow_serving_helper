// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/config/monitoring_config.proto

package tensorflow_serving

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

// Configuration for Prometheus monitoring.
type PrometheusConfig struct {
	// Whether to expose Prometheus metrics.
	Enable bool `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	// The endpoint to expose Prometheus metrics.
	// If not specified, PrometheusExporter::kPrometheusPath value is used.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrometheusConfig) Reset()         { *m = PrometheusConfig{} }
func (m *PrometheusConfig) String() string { return proto.CompactTextString(m) }
func (*PrometheusConfig) ProtoMessage()    {}
func (*PrometheusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_75e3f536986affeb, []int{0}
}

func (m *PrometheusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrometheusConfig.Unmarshal(m, b)
}
func (m *PrometheusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrometheusConfig.Marshal(b, m, deterministic)
}
func (m *PrometheusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusConfig.Merge(m, src)
}
func (m *PrometheusConfig) XXX_Size() int {
	return xxx_messageInfo_PrometheusConfig.Size(m)
}
func (m *PrometheusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusConfig proto.InternalMessageInfo

func (m *PrometheusConfig) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *PrometheusConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Configuration for monitoring.
type MonitoringConfig struct {
	PrometheusConfig     *PrometheusConfig `protobuf:"bytes,1,opt,name=prometheus_config,json=prometheusConfig,proto3" json:"prometheus_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MonitoringConfig) Reset()         { *m = MonitoringConfig{} }
func (m *MonitoringConfig) String() string { return proto.CompactTextString(m) }
func (*MonitoringConfig) ProtoMessage()    {}
func (*MonitoringConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_75e3f536986affeb, []int{1}
}

func (m *MonitoringConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitoringConfig.Unmarshal(m, b)
}
func (m *MonitoringConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitoringConfig.Marshal(b, m, deterministic)
}
func (m *MonitoringConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringConfig.Merge(m, src)
}
func (m *MonitoringConfig) XXX_Size() int {
	return xxx_messageInfo_MonitoringConfig.Size(m)
}
func (m *MonitoringConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringConfig proto.InternalMessageInfo

func (m *MonitoringConfig) GetPrometheusConfig() *PrometheusConfig {
	if m != nil {
		return m.PrometheusConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*PrometheusConfig)(nil), "tensorflow.serving.PrometheusConfig")
	proto.RegisterType((*MonitoringConfig)(nil), "tensorflow.serving.MonitoringConfig")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/config/monitoring_config.proto", fileDescriptor_75e3f536986affeb)
}

var fileDescriptor_75e3f536986affeb = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x8a, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcd, 0x2f, 0x4d, 0x37, 0x32, 0xd5, 0x2f,
	0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc,
	0x4b, 0x8f, 0xcf, 0x48, 0xcd, 0x29, 0x48, 0x2d, 0xc2, 0x22, 0xa3, 0x9f, 0x9c, 0x9f, 0x97, 0x96,
	0x99, 0xae, 0x9f, 0x9b, 0x9f, 0x97, 0x59, 0x92, 0x5f, 0x04, 0x52, 0x0b, 0x11, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x68, 0xd1, 0x83, 0x6a, 0x51, 0xb2, 0xe3, 0x12, 0x08, 0x28,
	0xca, 0xcf, 0x4d, 0x2d, 0xc9, 0x48, 0x2d, 0x2d, 0x76, 0x06, 0xab, 0x16, 0x12, 0xe3, 0x62, 0x4b,
	0xcd, 0x4b, 0x4c, 0xca, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0xf2, 0x84, 0x84,
	0xb8, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa5,
	0x54, 0x2e, 0x01, 0x5f, 0xb8, 0x75, 0x50, 0xfd, 0x81, 0x5c, 0x82, 0x05, 0x70, 0x33, 0xa1, 0x4e,
	0x00, 0x1b, 0xc5, 0x6d, 0xa4, 0xa2, 0x87, 0xe9, 0x06, 0x3d, 0x74, 0x07, 0x04, 0x09, 0x14, 0xa0,
	0x89, 0x38, 0x31, 0xff, 0x60, 0x64, 0x4c, 0x62, 0x03, 0x7b, 0xc3, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0x69, 0x7c, 0xfc, 0x28, 0x01, 0x00, 0x00,
}
