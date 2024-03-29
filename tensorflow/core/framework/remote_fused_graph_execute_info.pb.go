// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/remote_fused_graph_execute_info.proto

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

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type RemoteFusedGraphExecuteInfo struct {
	// Definition of remote graph
	RemoteGraph *GraphDef `protobuf:"bytes,1,opt,name=remote_graph,json=remoteGraph,proto3" json:"remote_graph,omitempty"`
	// Remote fused graph input node name
	GraphInputNodeName []string `protobuf:"bytes,2,rep,name=graph_input_node_name,json=graphInputNodeName,proto3" json:"graph_input_node_name,omitempty"`
	// Remote fused graph output node name
	GraphOutputNodeName []string `protobuf:"bytes,3,rep,name=graph_output_node_name,json=graphOutputNodeName,proto3" json:"graph_output_node_name,omitempty"`
	// Executor's name
	ExecutorName string `protobuf:"bytes,4,opt,name=executor_name,json=executorName,proto3" json:"executor_name,omitempty"`
	// Optional: Parameters given to the executor
	SerializedExecutorParameters []byte `protobuf:"bytes,5,opt,name=serialized_executor_parameters,json=serializedExecutorParameters,proto3" json:"serialized_executor_parameters,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	DefaultGraphInputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,6,rep,name=default_graph_input_tensor_shape,json=defaultGraphInputTensorShape,proto3" json:"default_graph_input_tensor_shape,omitempty"`
	// Optional: Default graph input tensor shape used to allocate memory
	// before executing op
	// TODO(satok): Remote output tensor shape once shape information is stored
	// in NodeDef
	DefaultGraphOutputTensorShape []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto `protobuf:"bytes,7,rep,name=default_graph_output_tensor_shape,json=defaultGraphOutputTensorShape,proto3" json:"default_graph_output_tensor_shape,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                                            `json:"-"`
	XXX_unrecognized              []byte                                              `json:"-"`
	XXX_sizecache                 int32                                               `json:"-"`
}

func (m *RemoteFusedGraphExecuteInfo) Reset()         { *m = RemoteFusedGraphExecuteInfo{} }
func (m *RemoteFusedGraphExecuteInfo) String() string { return proto.CompactTextString(m) }
func (*RemoteFusedGraphExecuteInfo) ProtoMessage()    {}
func (*RemoteFusedGraphExecuteInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15f13da5b37f691, []int{0}
}

func (m *RemoteFusedGraphExecuteInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Unmarshal(m, b)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Marshal(b, m, deterministic)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo.Merge(m, src)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_Size() int {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo.Size(m)
}
func (m *RemoteFusedGraphExecuteInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteFusedGraphExecuteInfo proto.InternalMessageInfo

func (m *RemoteFusedGraphExecuteInfo) GetRemoteGraph() *GraphDef {
	if m != nil {
		return m.RemoteGraph
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphInputNodeName() []string {
	if m != nil {
		return m.GraphInputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetGraphOutputNodeName() []string {
	if m != nil {
		return m.GraphOutputNodeName
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetExecutorName() string {
	if m != nil {
		return m.ExecutorName
	}
	return ""
}

func (m *RemoteFusedGraphExecuteInfo) GetSerializedExecutorParameters() []byte {
	if m != nil {
		return m.SerializedExecutorParameters
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphInputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphInputTensorShape
	}
	return nil
}

func (m *RemoteFusedGraphExecuteInfo) GetDefaultGraphOutputTensorShape() []*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto {
	if m != nil {
		return m.DefaultGraphOutputTensorShape
	}
	return nil
}

type RemoteFusedGraphExecuteInfo_TensorShapeTypeProto struct {
	Dtype                DataType          `protobuf:"varint,1,opt,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	Shape                *TensorShapeProto `protobuf:"bytes,2,opt,name=shape,proto3" json:"shape,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Reset() {
	*m = RemoteFusedGraphExecuteInfo_TensorShapeTypeProto{}
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) String() string {
	return proto.CompactTextString(m)
}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) ProtoMessage() {}
func (*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_c15f13da5b37f691, []int{0, 0}
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Unmarshal(m, b)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Marshal(b, m, deterministic)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Merge(m, src)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_Size() int {
	return xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.Size(m)
}
func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteFusedGraphExecuteInfo_TensorShapeTypeProto proto.InternalMessageInfo

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func (m *RemoteFusedGraphExecuteInfo_TensorShapeTypeProto) GetShape() *TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteFusedGraphExecuteInfo)(nil), "tensorflow.RemoteFusedGraphExecuteInfo")
	proto.RegisterType((*RemoteFusedGraphExecuteInfo_TensorShapeTypeProto)(nil), "tensorflow.RemoteFusedGraphExecuteInfo.TensorShapeTypeProto")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/remote_fused_graph_execute_info.proto", fileDescriptor_c15f13da5b37f691)
}

var fileDescriptor_c15f13da5b37f691 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0xef, 0xd2, 0x30,
	0x18, 0xc6, 0xd3, 0x3f, 0x0e, 0x43, 0x41, 0x0f, 0x15, 0xcd, 0x82, 0x68, 0xa6, 0xc6, 0x64, 0x31,
	0x66, 0x44, 0x38, 0x78, 0x31, 0x31, 0x12, 0x90, 0x70, 0x41, 0x32, 0xb9, 0x37, 0x95, 0xbd, 0x83,
	0x45, 0xb6, 0x2e, 0x5d, 0x27, 0xe2, 0xd9, 0xf8, 0x7d, 0xfc, 0x76, 0x1e, 0x4d, 0xdb, 0x39, 0xba,
	0x44, 0x38, 0xfd, 0x8f, 0xf0, 0xfe, 0x9e, 0xb7, 0xcf, 0xfb, 0xec, 0xc1, 0xef, 0x25, 0x64, 0x05,
	0x17, 0xf1, 0x81, 0x1f, 0x47, 0x5b, 0x2e, 0x60, 0x14, 0x0b, 0x96, 0xc2, 0x91, 0x8b, 0xaf, 0x23,
	0x01, 0x29, 0x97, 0x40, 0xe3, 0xb2, 0x80, 0x88, 0xee, 0x04, 0xcb, 0xf7, 0x14, 0xbe, 0xc3, 0xb6,
	0x94, 0x40, 0x93, 0x2c, 0xe6, 0x41, 0x2e, 0xb8, 0xe4, 0x04, 0x9f, 0x17, 0x0c, 0x5e, 0x5e, 0x5e,
	0xa6, 0xf5, 0x46, 0x32, 0x78, 0x7d, 0x19, 0x33, 0x13, 0x5a, 0xec, 0x59, 0x0e, 0x15, 0x7d, 0x65,
	0xa9, 0x3c, 0xe5, 0x50, 0x18, 0xec, 0xf9, 0x6f, 0x07, 0x3f, 0x0e, 0xb5, 0xe3, 0x8f, 0xca, 0xf0,
	0x42, 0xbd, 0x37, 0x37, 0x76, 0x97, 0x59, 0xcc, 0xc9, 0x5b, 0xdc, 0xab, 0x0e, 0xd2, 0x56, 0x5c,
	0xe4, 0x21, 0xbf, 0x3b, 0xee, 0x07, 0xe7, 0xed, 0x81, 0xd6, 0xcc, 0x20, 0x0e, 0xbb, 0x86, 0xd4,
	0xbf, 0xc9, 0x1b, 0xfc, 0xd0, 0x1c, 0x9f, 0x64, 0x79, 0x29, 0x69, 0xc6, 0x23, 0xa0, 0x19, 0x4b,
	0xc1, 0xbd, 0xf1, 0x5a, 0x7e, 0x27, 0x24, 0x7a, 0xb8, 0x54, 0xb3, 0x15, 0x8f, 0x60, 0xc5, 0x52,
	0x20, 0x13, 0xfc, 0xc8, 0x48, 0x78, 0x29, 0x9b, 0x9a, 0x96, 0xd6, 0x3c, 0xd0, 0xd3, 0x4f, 0x7a,
	0x58, 0x8b, 0x5e, 0xe0, 0x7b, 0x26, 0x5e, 0x2e, 0x0c, 0x7b, 0xc7, 0x43, 0x7e, 0x27, 0xec, 0xfd,
	0xfb, 0x53, 0x43, 0x33, 0xfc, 0xb4, 0x00, 0x91, 0xb0, 0x43, 0xf2, 0x03, 0x22, 0x5a, 0xf3, 0x39,
	0x53, 0x99, 0x48, 0x10, 0x85, 0xeb, 0x78, 0xc8, 0xef, 0x85, 0xc3, 0x33, 0x35, 0xaf, 0xa0, 0x75,
	0xcd, 0x90, 0x9f, 0x08, 0x7b, 0x11, 0xc4, 0xac, 0x3c, 0x48, 0x6a, 0xdf, 0x66, 0xa7, 0xef, 0xb6,
	0xbd, 0x96, 0xdf, 0x1d, 0xbf, 0xb3, 0x03, 0xba, 0x92, 0x6f, 0xb0, 0xd1, 0xd8, 0x67, 0x25, 0xdd,
	0x9c, 0x72, 0x58, 0xab, 0x8f, 0x12, 0x0e, 0xab, 0x57, 0x16, 0x75, 0x46, 0x16, 0x46, 0x7e, 0x21,
	0xfc, 0xac, 0x69, 0xa3, 0xca, 0xab, 0xe1, 0xe3, 0xee, 0x2d, 0xf8, 0x78, 0x62, 0xfb, 0x30, 0xb9,
	0x5b, 0xdc, 0xe0, 0x1b, 0xee, 0xff, 0x4f, 0x46, 0x5e, 0x61, 0x27, 0x52, 0x1d, 0xd3, 0x65, 0xb9,
	0xdf, 0x2c, 0xcb, 0x8c, 0x49, 0xa6, 0xc8, 0xd0, 0x20, 0x64, 0x8c, 0x1d, 0xe3, 0xf7, 0x46, 0x17,
	0x6b, 0x68, 0xb3, 0xd6, 0x72, 0xe3, 0xc7, 0xa0, 0xd3, 0x0f, 0xd8, 0xe5, 0x62, 0x67, 0x93, 0x75,
	0xb7, 0xa7, 0xde, 0x95, 0x23, 0xf5, 0x92, 0x35, 0xfa, 0x83, 0xd0, 0x97, 0xb6, 0x6e, 0xff, 0xe4,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xdb, 0x57, 0x3e, 0xc8, 0x03, 0x00, 0x00,
}
