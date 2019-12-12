// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/queue_runner.proto

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

// Protocol buffer representing a QueueRunner.
type QueueRunnerDef struct {
	// Queue name.
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	// A list of enqueue operations.
	EnqueueOpName []string `protobuf:"bytes,2,rep,name=enqueue_op_name,json=enqueueOpName,proto3" json:"enqueue_op_name,omitempty"`
	// The operation to run to close the queue.
	CloseOpName string `protobuf:"bytes,3,opt,name=close_op_name,json=closeOpName,proto3" json:"close_op_name,omitempty"`
	// The operation to run to cancel the queue.
	CancelOpName string `protobuf:"bytes,4,opt,name=cancel_op_name,json=cancelOpName,proto3" json:"cancel_op_name,omitempty"`
	// A list of exception types considered to signal a safely closed queue
	// if raised during enqueue operations.
	QueueClosedExceptionTypes []Code   `protobuf:"varint,5,rep,packed,name=queue_closed_exception_types,json=queueClosedExceptionTypes,proto3,enum=tensorflow.error.Code" json:"queue_closed_exception_types,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *QueueRunnerDef) Reset()         { *m = QueueRunnerDef{} }
func (m *QueueRunnerDef) String() string { return proto.CompactTextString(m) }
func (*QueueRunnerDef) ProtoMessage()    {}
func (*QueueRunnerDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28b395249d07e49, []int{0}
}

func (m *QueueRunnerDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueRunnerDef.Unmarshal(m, b)
}
func (m *QueueRunnerDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueRunnerDef.Marshal(b, m, deterministic)
}
func (m *QueueRunnerDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueRunnerDef.Merge(m, src)
}
func (m *QueueRunnerDef) XXX_Size() int {
	return xxx_messageInfo_QueueRunnerDef.Size(m)
}
func (m *QueueRunnerDef) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueRunnerDef.DiscardUnknown(m)
}

var xxx_messageInfo_QueueRunnerDef proto.InternalMessageInfo

func (m *QueueRunnerDef) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRunnerDef) GetEnqueueOpName() []string {
	if m != nil {
		return m.EnqueueOpName
	}
	return nil
}

func (m *QueueRunnerDef) GetCloseOpName() string {
	if m != nil {
		return m.CloseOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetCancelOpName() string {
	if m != nil {
		return m.CancelOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetQueueClosedExceptionTypes() []Code {
	if m != nil {
		return m.QueueClosedExceptionTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueRunnerDef)(nil), "tensorflow.QueueRunnerDef")
}

func init() {
	proto.RegisterFile("github.com/smoug25/tensorflow_serving_helper/tensorflow/core/protobuf/queue_runner.proto", fileDescriptor_d28b395249d07e49)
}

var fileDescriptor_d28b395249d07e49 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x99, 0xbf, 0xbf, 0x42, 0x47, 0x5b, 0xb1, 0x0b, 0xa9, 0xa2, 0x50, 0x8a, 0x48, 0x57,
	0x09, 0x54, 0x7c, 0x81, 0x56, 0xb7, 0x5a, 0x83, 0x50, 0x77, 0x43, 0x3a, 0xbd, 0x49, 0x8b, 0xc9,
	0xdc, 0xf1, 0x4e, 0xc6, 0xea, 0x9b, 0x8b, 0x2b, 0xe9, 0x9d, 0x6a, 0xf2, 0x00, 0x6e, 0x4f, 0xbe,
	0x2f, 0x67, 0xee, 0x91, 0xcf, 0xf9, 0xba, 0x5a, 0xf9, 0x45, 0xa4, 0xb1, 0x8c, 0x5d, 0x89, 0x3e,
	0x1f, 0xdf, 0xc4, 0x15, 0x18, 0x87, 0x94, 0x15, 0xb8, 0x51, 0x0e, 0xe8, 0x6d, 0x6d, 0x72, 0xb5,
	0x82, 0xc2, 0x02, 0x35, 0xbe, 0xc4, 0x1a, 0x09, 0x62, 0x4b, 0x58, 0xe1, 0xc2, 0x67, 0xf1, 0xab,
	0x07, 0x0f, 0x8a, 0xbc, 0x31, 0x40, 0x11, 0xa7, 0x3d, 0x59, 0xc3, 0x67, 0xf3, 0xbf, 0x69, 0x01,
	0x22, 0x24, 0xa5, 0x71, 0x09, 0x2e, 0x94, 0x0c, 0xbf, 0x84, 0xec, 0x3e, 0x6e, 0xbb, 0x13, 0xae,
	0xbe, 0x85, 0xac, 0x77, 0x21, 0x65, 0x78, 0x8d, 0x49, 0x4b, 0xe8, 0x8b, 0x81, 0x18, 0xb5, 0x93,
	0x36, 0x27, 0xf7, 0x69, 0x09, 0xbd, 0x2b, 0x79, 0x04, 0x26, 0x00, 0x68, 0x03, 0xf3, 0x6f, 0xd0,
	0x1a, 0xb5, 0x93, 0xce, 0x2e, 0x7e, 0xb0, 0xcc, 0x0d, 0x65, 0x47, 0x17, 0xe8, 0x6a, 0xaa, 0xc5,
	0x7f, 0x3a, 0xe0, 0x70, 0xc7, 0x5c, 0xca, 0xae, 0x4e, 0x8d, 0x86, 0xe2, 0x17, 0xfa, 0xcf, 0xd0,
	0x61, 0x48, 0x77, 0xd4, 0x5c, 0x9e, 0x87, 0x3e, 0x56, 0x97, 0x0a, 0xde, 0x35, 0xd8, 0x6a, 0x8d,
	0x46, 0x55, 0x1f, 0x16, 0x5c, 0x7f, 0x6f, 0xd0, 0x1a, 0x75, 0xc7, 0x27, 0x51, 0x7d, 0x76, 0xc4,
	0x87, 0x46, 0x53, 0x5c, 0x42, 0x72, 0xca, 0xee, 0x94, 0xd5, 0xbb, 0x1f, 0xf3, 0x69, 0x2b, 0x4e,
	0xc6, 0xb2, 0x8f, 0x94, 0x37, 0xbd, 0x8c, 0xd2, 0x12, 0x36, 0x48, 0x2f, 0x93, 0xe3, 0xc6, 0x2a,
	0xb3, 0xed, 0x54, 0x6e, 0x26, 0x3e, 0x85, 0x58, 0xec, 0xf3, 0x6e, 0xd7, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x22, 0x39, 0x58, 0xf8, 0x01, 0x00, 0x00,
}
