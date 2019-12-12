// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/get_model_status.proto

package tensorflow_serving

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	util "github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/util"
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

// States that map to ManagerState enum in
// tensorflow_serving/core/servable_state.h
type ModelVersionStatus_State int32

const (
	// Default value.
	ModelVersionStatus_UNKNOWN ModelVersionStatus_State = 0
	// The manager is tracking this servable, but has not initiated any action
	// pertaining to it.
	ModelVersionStatus_START ModelVersionStatus_State = 10
	// The manager has decided to load this servable. In particular, checks
	// around resource availability and other aspects have passed, and the
	// manager is about to invoke the loader's Load() method.
	ModelVersionStatus_LOADING ModelVersionStatus_State = 20
	// The manager has successfully loaded this servable and made it available
	// for serving (i.e. GetServableHandle(id) will succeed). To avoid races,
	// this state is not reported until *after* the servable is made
	// available.
	ModelVersionStatus_AVAILABLE ModelVersionStatus_State = 30
	// The manager has decided to make this servable unavailable, and unload
	// it. To avoid races, this state is reported *before* the servable is
	// made unavailable.
	ModelVersionStatus_UNLOADING ModelVersionStatus_State = 40
	// This servable has reached the end of its journey in the manager. Either
	// it loaded and ultimately unloaded successfully, or it hit an error at
	// some point in its lifecycle.
	ModelVersionStatus_END ModelVersionStatus_State = 50
)

var ModelVersionStatus_State_name = map[int32]string{
	0:  "UNKNOWN",
	10: "START",
	20: "LOADING",
	30: "AVAILABLE",
	40: "UNLOADING",
	50: "END",
}

var ModelVersionStatus_State_value = map[string]int32{
	"UNKNOWN":   0,
	"START":     10,
	"LOADING":   20,
	"AVAILABLE": 30,
	"UNLOADING": 40,
	"END":       50,
}

func (x ModelVersionStatus_State) String() string {
	return proto.EnumName(ModelVersionStatus_State_name, int32(x))
}

func (ModelVersionStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2623c4c1b86d7c87, []int{1, 0}
}

// GetModelStatusRequest contains a ModelSpec indicating the model for which
// to get status.
type GetModelStatusRequest struct {
	// Model Specification. If version is not specified, information about all
	// versions of the model will be returned. If a version is specified, the
	// status of only that version will be returned.
	ModelSpec            *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetModelStatusRequest) Reset()         { *m = GetModelStatusRequest{} }
func (m *GetModelStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetModelStatusRequest) ProtoMessage()    {}
func (*GetModelStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2623c4c1b86d7c87, []int{0}
}

func (m *GetModelStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelStatusRequest.Unmarshal(m, b)
}
func (m *GetModelStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelStatusRequest.Marshal(b, m, deterministic)
}
func (m *GetModelStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelStatusRequest.Merge(m, src)
}
func (m *GetModelStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetModelStatusRequest.Size(m)
}
func (m *GetModelStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelStatusRequest proto.InternalMessageInfo

func (m *GetModelStatusRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

// Version number, state, and status for a single version of a model.
type ModelVersionStatus struct {
	// Model version.
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Model state.
	State ModelVersionStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=tensorflow.serving.ModelVersionStatus_State" json:"state,omitempty"`
	// Model status.
	Status               *util.StatusProto `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ModelVersionStatus) Reset()         { *m = ModelVersionStatus{} }
func (m *ModelVersionStatus) String() string { return proto.CompactTextString(m) }
func (*ModelVersionStatus) ProtoMessage()    {}
func (*ModelVersionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_2623c4c1b86d7c87, []int{1}
}

func (m *ModelVersionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelVersionStatus.Unmarshal(m, b)
}
func (m *ModelVersionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelVersionStatus.Marshal(b, m, deterministic)
}
func (m *ModelVersionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelVersionStatus.Merge(m, src)
}
func (m *ModelVersionStatus) XXX_Size() int {
	return xxx_messageInfo_ModelVersionStatus.Size(m)
}
func (m *ModelVersionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelVersionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ModelVersionStatus proto.InternalMessageInfo

func (m *ModelVersionStatus) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ModelVersionStatus) GetState() ModelVersionStatus_State {
	if m != nil {
		return m.State
	}
	return ModelVersionStatus_UNKNOWN
}

func (m *ModelVersionStatus) GetStatus() *util.StatusProto {
	if m != nil {
		return m.Status
	}
	return nil
}

// Response for ModelStatusRequest on successful run.
type GetModelStatusResponse struct {
	// Version number and status information for applicable model version(s).
	ModelVersionStatus   []*ModelVersionStatus `protobuf:"bytes,1,rep,name=model_version_status,proto3" json:"model_version_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetModelStatusResponse) Reset()         { *m = GetModelStatusResponse{} }
func (m *GetModelStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetModelStatusResponse) ProtoMessage()    {}
func (*GetModelStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2623c4c1b86d7c87, []int{2}
}

func (m *GetModelStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetModelStatusResponse.Unmarshal(m, b)
}
func (m *GetModelStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetModelStatusResponse.Marshal(b, m, deterministic)
}
func (m *GetModelStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetModelStatusResponse.Merge(m, src)
}
func (m *GetModelStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetModelStatusResponse.Size(m)
}
func (m *GetModelStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetModelStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetModelStatusResponse proto.InternalMessageInfo

func (m *GetModelStatusResponse) GetModelVersionStatus() []*ModelVersionStatus {
	if m != nil {
		return m.ModelVersionStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("tensorflow.serving.ModelVersionStatus_State", ModelVersionStatus_State_name, ModelVersionStatus_State_value)
	proto.RegisterType((*GetModelStatusRequest)(nil), "tensorflow.serving.GetModelStatusRequest")
	proto.RegisterType((*ModelVersionStatus)(nil), "tensorflow.serving.ModelVersionStatus")
	proto.RegisterType((*GetModelStatusResponse)(nil), "tensorflow.serving.GetModelStatusResponse")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow_serving/apis/get_model_status.proto", fileDescriptor_2623c4c1b86d7c87)
}

var fileDescriptor_2623c4c1b86d7c87 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcb, 0x6b, 0xa3, 0x50,
	0x14, 0xc6, 0xc7, 0x48, 0x12, 0x72, 0xc2, 0x0c, 0x72, 0xc9, 0x0c, 0x12, 0x98, 0x99, 0xe0, 0xa2,
	0xb8, 0x28, 0x0a, 0x96, 0xd2, 0x4d, 0x37, 0x86, 0x84, 0x10, 0x9a, 0x9a, 0xd4, 0x3c, 0x0a, 0xed,
	0x42, 0x92, 0xf4, 0xd4, 0x08, 0xea, 0xb5, 0xde, 0x6b, 0xba, 0xed, 0x9f, 0xdd, 0x65, 0xd1, 0x6b,
	0xe8, 0x2b, 0x85, 0x42, 0xbb, 0xd2, 0xf3, 0xfa, 0x7d, 0xe7, 0x3b, 0x0a, 0xd7, 0x7e, 0xc0, 0x37,
	0xd9, 0xca, 0x58, 0xd3, 0xc8, 0x64, 0x11, 0xcd, 0x7c, 0xeb, 0xd8, 0xe4, 0x18, 0x33, 0x9a, 0xde,
	0x86, 0xf4, 0xde, 0x63, 0x98, 0x6e, 0x83, 0xd8, 0xf7, 0x36, 0x18, 0x26, 0x98, 0xee, 0xa9, 0x98,
	0xcb, 0x24, 0x60, 0xa6, 0x8f, 0xdc, 0x8b, 0xe8, 0x0d, 0x86, 0x1e, 0xe3, 0x4b, 0x9e, 0x31, 0x23,
	0x49, 0x29, 0xa7, 0x84, 0x3c, 0xf7, 0x1b, 0x65, 0x7f, 0x7b, 0xf2, 0x2d, 0x82, 0x85, 0x98, 0x50,
	0x69, 0x5f, 0x7c, 0x95, 0x98, 0xf1, 0x20, 0x34, 0x5f, 0x2e, 0xae, 0xcd, 0xe1, 0xf7, 0x00, 0xf9,
	0x79, 0x2e, 0x32, 0x2d, 0xf2, 0x2e, 0xde, 0x65, 0xc8, 0x38, 0x39, 0x05, 0x28, 0x7d, 0x26, 0xb8,
	0x56, 0xa5, 0x8e, 0xa4, 0x37, 0xad, 0xbf, 0xc6, 0x7b, 0x9b, 0x86, 0x98, 0x4d, 0x70, 0xed, 0x36,
	0xa2, 0xdd, 0xab, 0xf6, 0x50, 0x01, 0x52, 0x14, 0x16, 0x98, 0xb2, 0x80, 0xc6, 0x82, 0x4d, 0x54,
	0xa8, 0x6f, 0x45, 0xa2, 0x20, 0xca, 0xee, 0x2e, 0x24, 0x5d, 0xa8, 0xe6, 0x7b, 0xa1, 0x5a, 0xe9,
	0x48, 0xfa, 0x2f, 0xeb, 0xf0, 0x43, 0xa5, 0x57, 0x40, 0x23, 0x7f, 0xa0, 0x2b, 0x46, 0xc9, 0x09,
	0xd4, 0x84, 0x37, 0x55, 0x2e, 0xd6, 0xfd, 0xbf, 0x0f, 0x22, 0x06, 0x27, 0xb9, 0x79, 0xb7, 0x6c,
	0xd7, 0xa6, 0x50, 0x2d, 0x40, 0xa4, 0x09, 0xf5, 0xb9, 0x73, 0xe6, 0x8c, 0x2f, 0x1d, 0xe5, 0x07,
	0x69, 0x40, 0x75, 0x3a, 0xb3, 0xdd, 0x99, 0x02, 0x79, 0x7e, 0x34, 0xb6, 0x7b, 0x43, 0x67, 0xa0,
	0xb4, 0xc8, 0x4f, 0x68, 0xd8, 0x0b, 0x7b, 0x38, 0xb2, 0xbb, 0xa3, 0xbe, 0xf2, 0x2f, 0x0f, 0xe7,
	0xce, 0xae, 0xaa, 0x93, 0x3a, 0xc8, 0x7d, 0xa7, 0xa7, 0x58, 0x1a, 0x87, 0x3f, 0x6f, 0x2f, 0xcb,
	0x12, 0x1a, 0x33, 0x24, 0x57, 0xd0, 0x12, 0xa7, 0x2d, 0xcd, 0x97, 0xbf, 0x92, 0x2a, 0x75, 0x64,
	0xbd, 0x69, 0x1d, 0x7c, 0xce, 0xba, 0xbb, 0x97, 0xd1, 0x95, 0x1f, 0x25, 0x69, 0x55, 0x2b, 0xbe,
	0xed, 0xd1, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x0f, 0xaf, 0x06, 0xf3, 0x02, 0x00, 0x00,
}
