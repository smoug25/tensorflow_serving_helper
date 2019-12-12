// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/config/model_server_config.proto

package tensorflow_serving

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	storage_path "github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/sources/storage_path"
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

// The type of model.
// TODO(b/31336131): DEPRECATED.
type ModelType int32

const (
	ModelType_MODEL_TYPE_UNSPECIFIED ModelType = 0 // Deprecated: Do not use.
	ModelType_TENSORFLOW             ModelType = 1 // Deprecated: Do not use.
	ModelType_OTHER                  ModelType = 2 // Deprecated: Do not use.
)

var ModelType_name = map[int32]string{
	0: "MODEL_TYPE_UNSPECIFIED",
	1: "TENSORFLOW",
	2: "OTHER",
}

var ModelType_value = map[string]int32{
	"MODEL_TYPE_UNSPECIFIED": 0,
	"TENSORFLOW":             1,
	"OTHER":                  2,
}

func (x ModelType) String() string {
	return proto.EnumName(ModelType_name, int32(x))
}

func (ModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e880a30d4350b3e, []int{0}
}

// Common configuration for loading a model being served.
type ModelConfig struct {
	// Name of the model.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Base path to the model, excluding the version directory.
	// E.g> for a model at /foo/bar/my_model/123, where 123 is the version, the
	// base path is /foo/bar/my_model.
	//
	// (This can be changed once a model is in serving, *if* the underlying data
	// remains the same. Otherwise there are no guarantees about whether the old
	// or new data will be used for model versions currently loaded.)
	BasePath string `protobuf:"bytes,2,opt,name=base_path,json=basePath,proto3" json:"base_path,omitempty"`
	// Type of model.
	// TODO(b/31336131): DEPRECATED. Please use 'model_platform' instead.
	ModelType ModelType `protobuf:"varint,3,opt,name=model_type,json=modelType,proto3,enum=tensorflow.serving.ModelType" json:"model_type,omitempty"` // Deprecated: Do not use.
	// Type of model (e.g. "tensorflow").
	//
	// (This cannot be changed once a model is in serving.)
	ModelPlatform string `protobuf:"bytes,4,opt,name=model_platform,json=modelPlatform,proto3" json:"model_platform,omitempty"`
	// Version policy for the model indicating which version(s) of the model to
	// load and make available for serving simultaneously.
	// The default option is to serve only the latest version of the model.
	//
	// (This can be changed once a model is in serving.)
	ModelVersionPolicy *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy `protobuf:"bytes,7,opt,name=model_version_policy,json=modelVersionPolicy,proto3" json:"model_version_policy,omitempty"`
	// String labels to associate with versions of the model, allowing inference
	// queries to refer to versions by label instead of number. Multiple labels
	// can map to the same version, but not vice-versa.
	//
	// An envisioned use-case for these labels is canarying tentative versions.
	// For example, one can assign labels "stable" and "canary" to two specific
	// versions. Perhaps initially "stable" is assigned to version 0 and "canary"
	// to version 1. Once version 1 passes canary, one can shift the "stable"
	// label to refer to version 1 (at that point both labels map to the same
	// version -- version 1 -- which is fine). Later once version 2 is ready to
	// canary one can move the "canary" label to version 2. And so on.
	VersionLabels map[string]int64 `protobuf:"bytes,8,rep,name=version_labels,json=versionLabels,proto3" json:"version_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Configures logging requests and responses, to the model.
	//
	// (This can be changed once a model is in serving.)
	LoggingConfig        *LoggingConfig `protobuf:"bytes,6,opt,name=logging_config,json=loggingConfig,proto3" json:"logging_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ModelConfig) Reset()         { *m = ModelConfig{} }
func (m *ModelConfig) String() string { return proto.CompactTextString(m) }
func (*ModelConfig) ProtoMessage()    {}
func (*ModelConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e880a30d4350b3e, []int{0}
}

func (m *ModelConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelConfig.Unmarshal(m, b)
}
func (m *ModelConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelConfig.Marshal(b, m, deterministic)
}
func (m *ModelConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelConfig.Merge(m, src)
}
func (m *ModelConfig) XXX_Size() int {
	return xxx_messageInfo_ModelConfig.Size(m)
}
func (m *ModelConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModelConfig proto.InternalMessageInfo

func (m *ModelConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelConfig) GetBasePath() string {
	if m != nil {
		return m.BasePath
	}
	return ""
}

// Deprecated: Do not use.
func (m *ModelConfig) GetModelType() ModelType {
	if m != nil {
		return m.ModelType
	}
	return ModelType_MODEL_TYPE_UNSPECIFIED
}

func (m *ModelConfig) GetModelPlatform() string {
	if m != nil {
		return m.ModelPlatform
	}
	return ""
}

func (m *ModelConfig) GetModelVersionPolicy() *storage_path.FileSystemStoragePathSourceConfig_ServableVersionPolicy {
	if m != nil {
		return m.ModelVersionPolicy
	}
	return nil
}

func (m *ModelConfig) GetVersionLabels() map[string]int64 {
	if m != nil {
		return m.VersionLabels
	}
	return nil
}

func (m *ModelConfig) GetLoggingConfig() *LoggingConfig {
	if m != nil {
		return m.LoggingConfig
	}
	return nil
}

// Static list of models to be loaded for serving.
type ModelConfigList struct {
	Config               []*ModelConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ModelConfigList) Reset()         { *m = ModelConfigList{} }
func (m *ModelConfigList) String() string { return proto.CompactTextString(m) }
func (*ModelConfigList) ProtoMessage()    {}
func (*ModelConfigList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e880a30d4350b3e, []int{1}
}

func (m *ModelConfigList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelConfigList.Unmarshal(m, b)
}
func (m *ModelConfigList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelConfigList.Marshal(b, m, deterministic)
}
func (m *ModelConfigList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelConfigList.Merge(m, src)
}
func (m *ModelConfigList) XXX_Size() int {
	return xxx_messageInfo_ModelConfigList.Size(m)
}
func (m *ModelConfigList) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelConfigList.DiscardUnknown(m)
}

var xxx_messageInfo_ModelConfigList proto.InternalMessageInfo

func (m *ModelConfigList) GetConfig() []*ModelConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

// ModelServer config.
type ModelServerConfig struct {
	// ModelServer takes either a static file-based model config list or an Any
	// proto representing custom model config that is fetched dynamically at
	// runtime (through network RPC, custom service, etc.).
	//
	// Types that are valid to be assigned to Config:
	//	*ModelServerConfig_ModelConfigList
	//	*ModelServerConfig_CustomModelConfig
	Config               isModelServerConfig_Config `protobuf_oneof:"config"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ModelServerConfig) Reset()         { *m = ModelServerConfig{} }
func (m *ModelServerConfig) String() string { return proto.CompactTextString(m) }
func (*ModelServerConfig) ProtoMessage()    {}
func (*ModelServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e880a30d4350b3e, []int{2}
}

func (m *ModelServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelServerConfig.Unmarshal(m, b)
}
func (m *ModelServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelServerConfig.Marshal(b, m, deterministic)
}
func (m *ModelServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelServerConfig.Merge(m, src)
}
func (m *ModelServerConfig) XXX_Size() int {
	return xxx_messageInfo_ModelServerConfig.Size(m)
}
func (m *ModelServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ModelServerConfig proto.InternalMessageInfo

type isModelServerConfig_Config interface {
	isModelServerConfig_Config()
}

type ModelServerConfig_ModelConfigList struct {
	ModelConfigList *ModelConfigList `protobuf:"bytes,1,opt,name=model_config_list,json=modelConfigList,proto3,oneof"`
}

type ModelServerConfig_CustomModelConfig struct {
	CustomModelConfig *any.Any `protobuf:"bytes,2,opt,name=custom_model_config,json=customModelConfig,proto3,oneof"`
}

func (*ModelServerConfig_ModelConfigList) isModelServerConfig_Config() {}

func (*ModelServerConfig_CustomModelConfig) isModelServerConfig_Config() {}

func (m *ModelServerConfig) GetConfig() isModelServerConfig_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ModelServerConfig) GetModelConfigList() *ModelConfigList {
	if x, ok := m.GetConfig().(*ModelServerConfig_ModelConfigList); ok {
		return x.ModelConfigList
	}
	return nil
}

func (m *ModelServerConfig) GetCustomModelConfig() *any.Any {
	if x, ok := m.GetConfig().(*ModelServerConfig_CustomModelConfig); ok {
		return x.CustomModelConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ModelServerConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ModelServerConfig_ModelConfigList)(nil),
		(*ModelServerConfig_CustomModelConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("tensorflow.serving.ModelType", ModelType_name, ModelType_value)
	proto.RegisterType((*ModelConfig)(nil), "tensorflow.serving.ModelConfig")
	proto.RegisterMapType((map[string]int64)(nil), "tensorflow.serving.ModelConfig.VersionLabelsEntry")
	proto.RegisterType((*ModelConfigList)(nil), "tensorflow.serving.ModelConfigList")
	proto.RegisterType((*ModelServerConfig)(nil), "tensorflow.serving.ModelServerConfig")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/config/model_server_config.proto", fileDescriptor_0e880a30d4350b3e)
}

var fileDescriptor_0e880a30d4350b3e = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x9f, 0xdb, 0xae, 0xb4, 0xae, 0xda, 0x75, 0x66, 0x42, 0xa1, 0x08, 0x28, 0x43, 0x48, 0x15,
	0x87, 0x44, 0x0a, 0x42, 0x20, 0x4e, 0x63, 0x5b, 0xaa, 0x6e, 0x74, 0x6b, 0x49, 0x0a, 0x68, 0xe2,
	0x60, 0x92, 0xe2, 0xa6, 0x11, 0x4e, 0x1c, 0xc5, 0x4e, 0x51, 0x90, 0x78, 0x0a, 0x5e, 0x86, 0x47,
	0xe3, 0x88, 0x62, 0xa7, 0x90, 0xb1, 0xa2, 0x1d, 0xe0, 0x66, 0xff, 0xfc, 0x7d, 0xbf, 0x3f, 0x9f,
	0x6d, 0xf8, 0xc1, 0x0f, 0xc4, 0x32, 0xf5, 0xf4, 0x39, 0x0b, 0x0d, 0x1e, 0xb2, 0xd4, 0x37, 0x9f,
	0x1a, 0x82, 0x44, 0x9c, 0x25, 0x0b, 0xca, 0x3e, 0x63, 0x4e, 0x92, 0x55, 0x10, 0xf9, 0x78, 0x49,
	0x68, 0x4c, 0x92, 0x0d, 0x27, 0xc6, 0x9c, 0x45, 0x8b, 0xc0, 0x37, 0x42, 0xf6, 0x91, 0x50, 0x09,
	0x92, 0x04, 0x2b, 0x4c, 0x8f, 0x13, 0x26, 0x18, 0x42, 0xbf, 0x9b, 0xf4, 0xa2, 0xa9, 0x77, 0xdb,
	0x67, 0xcc, 0xa7, 0xc4, 0x90, 0x15, 0x5e, 0xba, 0x30, 0xdc, 0x28, 0x53, 0xe5, 0xbd, 0xf7, 0xff,
	0xc9, 0x10, 0x65, 0xbe, 0x9f, 0x17, 0x96, 0xbd, 0xf4, 0xbe, 0xfc, 0x2b, 0x39, 0x67, 0x69, 0x32,
	0x27, 0xdc, 0xe0, 0x82, 0x25, 0xae, 0x4f, 0x70, 0xec, 0x8a, 0xa5, 0xb1, 0x08, 0x28, 0xc1, 0x3c,
	0xe3, 0x82, 0x84, 0xb8, 0x7c, 0x80, 0x55, 0xb5, 0xd2, 0xde, 0xff, 0x56, 0x83, 0xad, 0xb3, 0x7c,
	0x4a, 0x47, 0xd2, 0x11, 0x42, 0xb0, 0x16, 0xb9, 0x21, 0xd1, 0x40, 0x1f, 0x0c, 0x9a, 0xb6, 0x5c,
	0xa3, 0x3b, 0xb0, 0xe9, 0xb9, 0x5c, 0x75, 0x6b, 0x15, 0x79, 0xd0, 0xc8, 0x81, 0xa9, 0x2b, 0x96,
	0xe8, 0x00, 0x42, 0x35, 0x65, 0x91, 0xc5, 0x44, 0xab, 0xf6, 0xc1, 0xa0, 0x63, 0xde, 0xd5, 0xaf,
	0x4e, 0x57, 0x97, 0x2a, 0xb3, 0x2c, 0x26, 0x87, 0x15, 0x0d, 0xd8, 0xcd, 0x70, 0xbd, 0x45, 0x8f,
	0x60, 0x47, 0x31, 0xc4, 0xd4, 0x15, 0x0b, 0x96, 0x84, 0x5a, 0x4d, 0x6a, 0xb4, 0x25, 0x3a, 0x2d,
	0x40, 0xf4, 0x15, 0xee, 0xa9, 0xb2, 0x15, 0x49, 0x78, 0xc0, 0x22, 0x1c, 0x33, 0x1a, 0xcc, 0x33,
	0xed, 0x46, 0x1f, 0x0c, 0x5a, 0xe6, 0xab, 0x4d, 0x92, 0xc3, 0x80, 0x12, 0x47, 0x4e, 0xc0, 0x51,
	0x03, 0xc8, 0x1d, 0x3b, 0x32, 0xbe, 0x8a, 0xab, 0x3b, 0x24, 0x59, 0xb9, 0x1e, 0x25, 0x6f, 0x15,
	0xe7, 0x54, 0x52, 0xda, 0x48, 0x0a, 0x5d, 0xc2, 0xd0, 0x05, 0xec, 0xac, 0x85, 0xa9, 0xeb, 0x11,
	0xca, 0xb5, 0x46, 0xbf, 0x3a, 0x68, 0x99, 0xe6, 0x5f, 0xb3, 0x16, 0x12, 0x05, 0xcd, 0x58, 0x36,
	0x59, 0x91, 0x48, 0x32, 0xbb, 0xbd, 0x2a, 0x63, 0x68, 0x04, 0x3b, 0x97, 0xdf, 0x85, 0x56, 0x97,
	0x99, 0x1e, 0x6c, 0xa2, 0x1e, 0xab, 0x4a, 0x45, 0x6e, 0xb7, 0x69, 0x79, 0xdb, 0x3b, 0x80, 0xe8,
	0xaa, 0x1c, 0xea, 0xc2, 0xea, 0x27, 0x92, 0x15, 0x57, 0x9a, 0x2f, 0xd1, 0x1e, 0xdc, 0x5e, 0xb9,
	0x34, 0x25, 0xf2, 0x36, 0xab, 0xb6, 0xda, 0xbc, 0xa8, 0x3c, 0x07, 0xa7, 0xb5, 0xc6, 0x76, 0xb7,
	0xbe, 0x7f, 0x0a, 0x77, 0x4a, 0x11, 0xc6, 0x01, 0x17, 0xe8, 0x19, 0xac, 0x17, 0xe6, 0x80, 0xcc,
	0x7d, 0xff, 0x9a, 0xdc, 0x76, 0x51, 0xbe, 0xff, 0x1d, 0xc0, 0x5d, 0x89, 0x3b, 0xf2, 0x1b, 0x16,
	0xef, 0xec, 0x35, 0xdc, 0x55, 0xb7, 0xa9, 0xaa, 0x30, 0x0d, 0xb8, 0x90, 0x0e, 0x5b, 0xe6, 0xc3,
	0x6b, 0x98, 0x73, 0x3b, 0xa3, 0x2d, 0x7b, 0x27, 0xfc, 0xc3, 0xe1, 0x10, 0xde, 0x9c, 0xa7, 0x5c,
	0xb0, 0x10, 0x97, 0x99, 0x65, 0xc4, 0x96, 0xb9, 0xa7, 0xab, 0xcf, 0xad, 0xaf, 0x3f, 0xb7, 0xfe,
	0x32, 0xca, 0x46, 0x5b, 0xf6, 0xae, 0x6a, 0x29, 0xd1, 0x1f, 0x36, 0xd6, 0x49, 0x1f, 0x9f, 0xc3,
	0xe6, 0xaf, 0x57, 0x8b, 0xee, 0xc1, 0x5b, 0x67, 0x93, 0x63, 0x6b, 0x8c, 0x67, 0x17, 0x53, 0x0b,
	0xbf, 0x39, 0x77, 0xa6, 0xd6, 0xd1, 0xc9, 0xf0, 0xc4, 0x3a, 0xee, 0x6e, 0xf5, 0x2a, 0x0d, 0x80,
	0x10, 0x84, 0x33, 0xeb, 0xdc, 0x99, 0xd8, 0xc3, 0xf1, 0xe4, 0x5d, 0x17, 0x48, 0xac, 0x0d, 0xb7,
	0x27, 0xb3, 0x91, 0x65, 0x77, 0x2b, 0xf9, 0xf6, 0xb0, 0xfa, 0x03, 0x00, 0xaf, 0x2e, 0x1d, 0x3c,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x33, 0x73, 0xdd, 0x58, 0xe4, 0x04, 0x00, 0x00,
}
