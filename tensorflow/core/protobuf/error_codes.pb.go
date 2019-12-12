// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/error_codes.proto

package tensorflow_error

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

// The canonical error codes for TensorFlow APIs.
//
// Warnings:
//
// -   Do not change any numeric assignments.
// -   Changes to this list should only be made if there is a compelling
//     need that can't be satisfied in another way.  Such changes
//     must be approved by at least two OWNERS.
// -   These error codes must match gRPC and protobuf error codes (except for
//     DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_).
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// OUT_OF_RANGE over FAILED_PRECONDITION if both codes apply.
// Similarly prefer NOT_FOUND or ALREADY_EXISTS over FAILED_PRECONDITION.
type Code int32

const (
	// Not an error; returned on success
	Code_OK Code = 0
	// The operation was cancelled (typically by the caller).
	Code_CANCELLED Code = 1
	// Unknown error.  An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Code_UNKNOWN Code = 2
	// Client specified an invalid argument.  Note that this differs
	// from FAILED_PRECONDITION.  INVALID_ARGUMENT indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Code_INVALID_ARGUMENT Code = 3
	// Deadline expired before operation could complete.  For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	Code_DEADLINE_EXCEEDED Code = 4
	// Some requested entity (e.g., file or directory) was not found.
	// For privacy reasons, this code *may* be returned when the client
	// does not have the access right to the entity.
	Code_NOT_FOUND Code = 5
	// Some entity that we attempted to create (e.g., file or directory)
	// already exists.
	Code_ALREADY_EXISTS Code = 6
	// The caller does not have permission to execute the specified
	// operation.  PERMISSION_DENIED must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors).  PERMISSION_DENIED must not be
	// used if the caller can not be identified (use UNAUTHENTICATED
	// instead for those errors).
	Code_PERMISSION_DENIED Code = 7
	// The request does not have valid authentication credentials for the
	// operation.
	Code_UNAUTHENTICATED Code = 16
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	Code_RESOURCE_EXHAUSTED Code = 8
	// Operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, directory
	// to be deleted may be non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE:
	//  (a) Use UNAVAILABLE if the client can retry just the failing call.
	//  (b) Use ABORTED if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FAILED_PRECONDITION if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, FAILED_PRECONDITION
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FAILED_PRECONDITION if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	Code_FAILED_PRECONDITION Code = 9
	// The operation was aborted, typically due to a concurrency issue
	// like sequencer check failures, transaction aborts, etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_ABORTED Code = 10
	// Operation tried to iterate past the valid input range.  E.g., seeking or
	// reading past end of file.
	//
	// Unlike INVALID_ARGUMENT, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate INVALID_ARGUMENT if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OUT_OF_RANGE if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FAILED_PRECONDITION and
	// OUT_OF_RANGE.  We recommend using OUT_OF_RANGE (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OUT_OF_RANGE error to detect when
	// they are done.
	Code_OUT_OF_RANGE Code = 11
	// Operation is not implemented or not supported/enabled in this service.
	Code_UNIMPLEMENTED Code = 12
	// Internal errors.  Means some invariant expected by the underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	Code_INTERNAL Code = 13
	// The service is currently unavailable.  This is a most likely a
	// transient condition and may be corrected by retrying with
	// a backoff.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_UNAVAILABLE Code = 14
	// Unrecoverable data loss or corruption.
	Code_DATA_LOSS Code = 15
	// An extra enum entry to prevent people from writing code that
	// fails to compile when a new code is added.
	//
	// Nobody should ever reference this enumeration entry. In particular,
	// if you write C++ code that switches on this enumeration, add a default:
	// case instead of a case that mentions this enumeration entry.
	//
	// Nobody should rely on the value (currently 20) listed here.  It
	// may change in the future.
	Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_ Code = 20
)

var Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	16: "UNAUTHENTICATED",
	8:  "RESOURCE_EXHAUSTED",
	9:  "FAILED_PRECONDITION",
	10: "ABORTED",
	11: "OUT_OF_RANGE",
	12: "UNIMPLEMENTED",
	13: "INTERNAL",
	14: "UNAVAILABLE",
	15: "DATA_LOSS",
	20: "DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_",
}

var Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     16,
	"RESOURCE_EXHAUSTED":  8,
	"FAILED_PRECONDITION": 9,
	"ABORTED":             10,
	"OUT_OF_RANGE":        11,
	"UNIMPLEMENTED":       12,
	"INTERNAL":            13,
	"UNAVAILABLE":         14,
	"DATA_LOSS":           15,
	"DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_": 20,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fe19ebdfbdec439e, []int{0}
}

func init() {
	proto.RegisterEnum("tensorflow.error.Code", Code_name, Code_value)
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/error_codes.proto", fileDescriptor_fe19ebdfbdec439e)
}

var fileDescriptor_fe19ebdfbdec439e = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcb, 0x72, 0x1a, 0x31,
	0x10, 0x45, 0x83, 0xed, 0x60, 0x5b, 0x80, 0x69, 0xcb, 0xce, 0xe3, 0x1b, 0xbc, 0x80, 0x4a, 0xe5,
	0x0b, 0x9a, 0xe9, 0x1e, 0x50, 0x21, 0x5a, 0x53, 0x7a, 0x60, 0x67, 0xa5, 0x8a, 0xed, 0x21, 0x8b,
	0x3c, 0x26, 0x35, 0x38, 0xc5, 0x0f, 0xe4, 0xa3, 0xb3, 0x4c, 0x89, 0x4d, 0xb2, 0xbd, 0xba, 0x92,
	0xee, 0xa9, 0xa3, 0xee, 0x5e, 0xda, 0x1f, 0xfb, 0xae, 0xdf, 0x7d, 0xeb, 0x0e, 0xf3, 0xa7, 0xae,
	0x6f, 0xe7, 0x3f, 0xfb, 0xee, 0xa5, 0x7b, 0xfc, 0xb5, 0x9b, 0xb7, 0x7d, 0xdf, 0xf5, 0xf9, 0xa9,
	0x7b, 0x6e, 0xf7, 0xb3, 0x63, 0xa8, 0xe1, 0x5f, 0x77, 0x76, 0x3c, 0xbd, 0xfb, 0x7d, 0xaa, 0xce,
	0xaa, 0xee, 0xb9, 0xd5, 0x43, 0x75, 0xe2, 0xd6, 0xf0, 0x4a, 0x4f, 0xd4, 0x65, 0x85, 0x52, 0xb1,
	0xb5, 0x4c, 0x30, 0xd0, 0x23, 0x75, 0x9e, 0x64, 0x2d, 0xee, 0x5e, 0xe0, 0x44, 0xdf, 0x2a, 0x30,
	0xb2, 0x45, 0x6b, 0x28, 0xa3, 0x5f, 0xa6, 0x0d, 0x4b, 0x84, 0x53, 0xfd, 0x46, 0x5d, 0x13, 0x23,
	0x59, 0x23, 0x9c, 0xf9, 0xa1, 0x62, 0x26, 0x26, 0x38, 0x2b, 0x0f, 0x89, 0x8b, 0xb9, 0x76, 0x49,
	0x08, 0x5e, 0x6b, 0xad, 0xae, 0xd0, 0x7a, 0x46, 0xfa, 0x94, 0xf9, 0xc1, 0x84, 0x18, 0x60, 0x58,
	0x6e, 0x36, 0xec, 0x37, 0x26, 0x04, 0xe3, 0x24, 0x13, 0x8b, 0x61, 0x82, 0x73, 0x7d, 0xa3, 0xa6,
	0x49, 0x30, 0xc5, 0x15, 0x4b, 0x34, 0x15, 0x46, 0x26, 0x00, 0xfd, 0x56, 0x69, 0xcf, 0xc1, 0x25,
	0x5f, 0x95, 0x5f, 0x56, 0x98, 0x42, 0xc9, 0x2f, 0xf4, 0x3b, 0x75, 0x53, 0xa3, 0xb1, 0x4c, 0xb9,
	0xf1, 0x5c, 0x39, 0x21, 0x13, 0x8d, 0x13, 0xb8, 0x2c, 0xcb, 0x71, 0xe1, 0x7c, 0x69, 0x29, 0x0d,
	0x6a, 0xec, 0x52, 0xcc, 0xae, 0xce, 0x1e, 0x65, 0xc9, 0x30, 0xd2, 0xd7, 0x6a, 0x92, 0xc4, 0x6c,
	0x1a, 0xcb, 0x05, 0x83, 0x09, 0xc6, 0x7a, 0xac, 0x2e, 0x8c, 0x44, 0xf6, 0x82, 0x16, 0x26, 0x7a,
	0xaa, 0x46, 0x49, 0x70, 0x8b, 0xc6, 0xe2, 0xc2, 0x32, 0x5c, 0x15, 0x20, 0xc2, 0x88, 0xd9, 0xba,
	0x10, 0x60, 0xaa, 0xd7, 0x6a, 0x49, 0x2e, 0x17, 0xc4, 0x14, 0x38, 0x7b, 0x0e, 0xec, 0xb7, 0x4c,
	0xb9, 0x76, 0x3e, 0xd7, 0x29, 0x26, 0x5f, 0x66, 0x36, 0x28, 0x47, 0xb4, 0xd2, 0x20, 0xae, 0x31,
	0xd9, 0x98, 0x8d, 0xe4, 0x70, 0x6f, 0x62, 0xb5, 0xca, 0x46, 0x42, 0x64, 0xa4, 0x0c, 0xb7, 0x8b,
	0x0f, 0xea, 0x7d, 0xd7, 0x7f, 0x99, 0xfd, 0xa7, 0x67, 0xd7, 0x7f, 0xfe, 0xde, 0x1e, 0xba, 0xfe,
	0xeb, 0x02, 0xb8, 0x98, 0x2a, 0x92, 0xf6, 0x4d, 0xb1, 0xb8, 0x6f, 0x06, 0x7f, 0x06, 0x83, 0xc7,
	0xe1, 0x51, 0xe9, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0xb0, 0x97, 0xc4, 0x00, 0x02,
	0x00, 0x00,
}
