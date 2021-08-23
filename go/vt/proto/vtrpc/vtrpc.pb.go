//
//Copyright 2019 The Vitess Authors.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// This file contains useful data structures for RPCs in Vitess.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: vtrpc.proto

package vtrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Code represents canonical error codes. The names, numbers and comments
// must match the ones defined by grpc:
// https://godoc.org/google.golang.org/grpc/codes.
type Code int32

const (
	// OK is returned on success.
	Code_OK Code = 0
	// CANCELED indicates the operation was cancelled (typically by the caller).
	Code_CANCELED Code = 1
	// UNKNOWN error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Code_UNKNOWN Code = 2
	// INVALID_ARGUMENT indicates client specified an invalid argument.
	// Note that this differs from FAILED_PRECONDITION. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Code_INVALID_ARGUMENT Code = 3
	// DEADLINE_EXCEEDED means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	Code_DEADLINE_EXCEEDED Code = 4
	// NOT_FOUND means some requested entity (e.g., file or directory) was
	// not found.
	Code_NOT_FOUND Code = 5
	// ALREADY_EXISTS means an attempt to create an entity failed because one
	// already exists.
	Code_ALREADY_EXISTS Code = 6
	// PERMISSION_DENIED indicates the caller does not have permission to
	// execute the specified operation. It must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors).  It must not be
	// used if the caller cannot be identified (use Unauthenticated
	// instead for those errors).
	Code_PERMISSION_DENIED Code = 7
	// UNAUTHENTICATED indicates the request does not have valid
	// authentication credentials for the operation.
	Code_UNAUTHENTICATED Code = 16
	// RESOURCE_EXHAUSTED indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	Code_RESOURCE_EXHAUSTED Code = 8
	// FAILED_PRECONDITION indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
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
	// ABORTED indicates the operation was aborted, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_ABORTED Code = 10
	// OUT_OF_RANGE means operation was attempted past the valid range.
	// E.g., seeking or reading past end of file.
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
	// UNIMPLEMENTED indicates operation is not implemented or not
	// supported/enabled in this service.
	Code_UNIMPLEMENTED Code = 12
	// INTERNAL errors. Means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	Code_INTERNAL Code = 13
	// UNAVAILABLE indicates the service is currently unavailable.
	// This is a most likely a transient condition and may be corrected
	// by retrying with a backoff.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_UNAVAILABLE Code = 14
	// DATA_LOSS indicates unrecoverable data loss or corruption.
	Code_DATA_LOSS Code = 15
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "OK",
		1:  "CANCELED",
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
	}
	Code_value = map[string]int32{
		"OK":                  0,
		"CANCELED":            1,
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
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_vtrpc_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_vtrpc_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Code.Descriptor instead.
func (Code) EnumDescriptor() ([]byte, []int) {
	return file_vtrpc_proto_rawDescGZIP(), []int{0}
}

// LegacyErrorCode is the enum values for Errors. This type is deprecated.
// Use Code instead. Background: In the initial design, we thought
// that we may end up with a different list of canonical error codes
// than the ones defined by grpc. In hindsight, we realize that
// the grpc error codes are fairly generic and mostly sufficient.
// In order to avoid confusion, this type will be deprecated in
// favor of the new Code that matches exactly what grpc defines.
// Some names below have a _LEGACY suffix. This is to prevent
// name collisions with Code.
type LegacyErrorCode int32

const (
	// SUCCESS_LEGACY is returned from a successful call.
	LegacyErrorCode_SUCCESS_LEGACY LegacyErrorCode = 0
	// CANCELLED_LEGACY means that the context was cancelled (and noticed in the app layer,
	// as opposed to the RPC layer).
	LegacyErrorCode_CANCELLED_LEGACY LegacyErrorCode = 1
	// UNKNOWN_ERROR_LEGACY includes:
	// 1. MySQL error codes that we don't explicitly handle.
	// 2. MySQL response that wasn't as expected. For example, we might expect a MySQL
	//  timestamp to be returned in a particular way, but it wasn't.
	// 3. Anything else that doesn't fall into a different bucket.
	LegacyErrorCode_UNKNOWN_ERROR_LEGACY LegacyErrorCode = 2
	// BAD_INPUT_LEGACY is returned when an end-user either sends SQL that couldn't be parsed correctly,
	// or tries a query that isn't supported by Vitess.
	LegacyErrorCode_BAD_INPUT_LEGACY LegacyErrorCode = 3
	// DEADLINE_EXCEEDED_LEGACY is returned when an action is taking longer than a given timeout.
	LegacyErrorCode_DEADLINE_EXCEEDED_LEGACY LegacyErrorCode = 4
	// INTEGRITY_ERROR_LEGACY is returned on integrity error from MySQL, usually due to
	// duplicate primary keys.
	LegacyErrorCode_INTEGRITY_ERROR_LEGACY LegacyErrorCode = 5
	// PERMISSION_DENIED_LEGACY errors are returned when a user requests access to something
	// that they don't have permissions for.
	LegacyErrorCode_PERMISSION_DENIED_LEGACY LegacyErrorCode = 6
	// RESOURCE_EXHAUSTED_LEGACY is returned when a query exceeds its quota in some dimension
	// and can't be completed due to that. Queries that return RESOURCE_EXHAUSTED
	// should not be retried, as it could be detrimental to the server's health.
	// Examples of errors that will cause the RESOURCE_EXHAUSTED code:
	// 1. TxPoolFull: this is retried server-side, and is only returned as an error
	//  if the server-side retries failed.
	// 2. Query is killed due to it taking too long.
	LegacyErrorCode_RESOURCE_EXHAUSTED_LEGACY LegacyErrorCode = 7
	// QUERY_NOT_SERVED_LEGACY means that a query could not be served right now.
	// Client can interpret it as: "the tablet that you sent this query to cannot
	// serve the query right now, try a different tablet or try again later."
	// This could be due to various reasons: QueryService is not serving, should
	// not be serving, wrong shard, wrong tablet type, table is part of the denylist, etc.
	// Clients that receive this error should usually retry the query, but after taking
	// the appropriate steps to make sure that the query will get sent to the correct
	// tablet.
	LegacyErrorCode_QUERY_NOT_SERVED_LEGACY LegacyErrorCode = 8
	// NOT_IN_TX_LEGACY means that we're not currently in a transaction, but we should be.
	LegacyErrorCode_NOT_IN_TX_LEGACY LegacyErrorCode = 9
	// INTERNAL_ERROR_LEGACY means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	LegacyErrorCode_INTERNAL_ERROR_LEGACY LegacyErrorCode = 10
	// TRANSIENT_ERROR_LEGACY is used for when there is some error that we expect we can
	// recover from automatically - often due to a resource limit temporarily being
	// reached. Retrying this error, with an exponential backoff, should succeed.
	// Clients should be able to successfully retry the query on the same backends.
	// Examples of things that can trigger this error:
	// 1. Query has been throttled
	// 2. VtGate could have request backlog
	LegacyErrorCode_TRANSIENT_ERROR_LEGACY LegacyErrorCode = 11
	// UNAUTHENTICATED_LEGACY errors are returned when a user requests access to something,
	// and we're unable to verify the user's authentication.
	LegacyErrorCode_UNAUTHENTICATED_LEGACY LegacyErrorCode = 12
)

// Enum value maps for LegacyErrorCode.
var (
	LegacyErrorCode_name = map[int32]string{
		0:  "SUCCESS_LEGACY",
		1:  "CANCELLED_LEGACY",
		2:  "UNKNOWN_ERROR_LEGACY",
		3:  "BAD_INPUT_LEGACY",
		4:  "DEADLINE_EXCEEDED_LEGACY",
		5:  "INTEGRITY_ERROR_LEGACY",
		6:  "PERMISSION_DENIED_LEGACY",
		7:  "RESOURCE_EXHAUSTED_LEGACY",
		8:  "QUERY_NOT_SERVED_LEGACY",
		9:  "NOT_IN_TX_LEGACY",
		10: "INTERNAL_ERROR_LEGACY",
		11: "TRANSIENT_ERROR_LEGACY",
		12: "UNAUTHENTICATED_LEGACY",
	}
	LegacyErrorCode_value = map[string]int32{
		"SUCCESS_LEGACY":            0,
		"CANCELLED_LEGACY":          1,
		"UNKNOWN_ERROR_LEGACY":      2,
		"BAD_INPUT_LEGACY":          3,
		"DEADLINE_EXCEEDED_LEGACY":  4,
		"INTEGRITY_ERROR_LEGACY":    5,
		"PERMISSION_DENIED_LEGACY":  6,
		"RESOURCE_EXHAUSTED_LEGACY": 7,
		"QUERY_NOT_SERVED_LEGACY":   8,
		"NOT_IN_TX_LEGACY":          9,
		"INTERNAL_ERROR_LEGACY":     10,
		"TRANSIENT_ERROR_LEGACY":    11,
		"UNAUTHENTICATED_LEGACY":    12,
	}
)

func (x LegacyErrorCode) Enum() *LegacyErrorCode {
	p := new(LegacyErrorCode)
	*p = x
	return p
}

func (x LegacyErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LegacyErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_vtrpc_proto_enumTypes[1].Descriptor()
}

func (LegacyErrorCode) Type() protoreflect.EnumType {
	return &file_vtrpc_proto_enumTypes[1]
}

func (x LegacyErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LegacyErrorCode.Descriptor instead.
func (LegacyErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_vtrpc_proto_rawDescGZIP(), []int{1}
}

// CallerID is passed along RPCs to identify the originating client
// for a request. It is not meant to be secure, but only
// informational.  The client can put whatever info they want in these
// fields, and they will be trusted by the servers. The fields will
// just be used for logging purposes, and to easily find a client.
// VtGate propagates it to VtTablet, and VtTablet may use this
// information for monitoring purposes, to display on dashboards, or
// for denying access to tables during a migration.
type CallerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// principal is the effective user identifier. It is usually filled in
	// with whoever made the request to the appserver, if the request
	// came from an automated job or another system component.
	// If the request comes directly from the Internet, or if the Vitess client
	// takes action on its own accord, it is okay for this field to be absent.
	Principal string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	// component describes the running process of the effective caller.
	// It can for instance be the hostname:port of the servlet initiating the
	// database call, or the container engine ID used by the servlet.
	Component string `protobuf:"bytes,2,opt,name=component,proto3" json:"component,omitempty"`
	// subcomponent describes a component inisde the immediate caller which
	// is responsible for generating is request. Suggested values are a
	// servlet name or an API endpoint name.
	Subcomponent string `protobuf:"bytes,3,opt,name=subcomponent,proto3" json:"subcomponent,omitempty"`
}

func (x *CallerID) Reset() {
	*x = CallerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vtrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallerID) ProtoMessage() {}

func (x *CallerID) ProtoReflect() protoreflect.Message {
	mi := &file_vtrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallerID.ProtoReflect.Descriptor instead.
func (*CallerID) Descriptor() ([]byte, []int) {
	return file_vtrpc_proto_rawDescGZIP(), []int{0}
}

func (x *CallerID) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *CallerID) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *CallerID) GetSubcomponent() string {
	if x != nil {
		return x.Subcomponent
	}
	return ""
}

// RPCError is an application-level error structure returned by
// VtTablet (and passed along by VtGate if appropriate).
// We use this so the clients don't have to parse the error messages,
// but instead can depend on the value of the code.
type RPCError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LegacyCode LegacyErrorCode `protobuf:"varint,1,opt,name=legacy_code,json=legacyCode,proto3,enum=vtrpc.LegacyErrorCode" json:"legacy_code,omitempty"`
	Message    string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code       Code            `protobuf:"varint,3,opt,name=code,proto3,enum=vtrpc.Code" json:"code,omitempty"`
}

func (x *RPCError) Reset() {
	*x = RPCError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vtrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RPCError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RPCError) ProtoMessage() {}

func (x *RPCError) ProtoReflect() protoreflect.Message {
	mi := &file_vtrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RPCError.ProtoReflect.Descriptor instead.
func (*RPCError) Descriptor() ([]byte, []int) {
	return file_vtrpc_proto_rawDescGZIP(), []int{1}
}

func (x *RPCError) GetLegacyCode() LegacyErrorCode {
	if x != nil {
		return x.LegacyCode
	}
	return LegacyErrorCode_SUCCESS_LEGACY
}

func (x *RPCError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *RPCError) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_OK
}

var File_vtrpc_proto protoreflect.FileDescriptor

var file_vtrpc_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x74, 0x72, 0x70, 0x63, 0x22, 0x6a, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x22, 0x7e, 0x0a, 0x08, 0x52, 0x50, 0x43, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x0b,
	0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x76, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x76, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x2a, 0xb6, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45,
	0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45,
	0x44, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54,
	0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x10, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x08,
	0x12, 0x17, 0x0a, 0x13, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x50, 0x52, 0x45, 0x43, 0x4f,
	0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x42, 0x4f,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46,
	0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x49, 0x4d,
	0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x41,
	0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x41,
	0x54, 0x41, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x0f, 0x2a, 0xe8, 0x02, 0x0a, 0x0f, 0x4c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x4c,
	0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x41, 0x44, 0x5f, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x4c,
	0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x41, 0x44, 0x4c,
	0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x5f, 0x4c, 0x45, 0x47,
	0x41, 0x43, 0x59, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10,
	0x05, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x06, 0x12,
	0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x45, 0x58, 0x48, 0x41,
	0x55, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x07, 0x12, 0x1b,
	0x0a, 0x17, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x44, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x54, 0x58, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10,
	0x09, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x0b, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x4e, 0x41, 0x55,
	0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x4c, 0x45, 0x47, 0x41,
	0x43, 0x59, 0x10, 0x0c, 0x42, 0x35, 0x0a, 0x0f, 0x69, 0x6f, 0x2e, 0x76, 0x69, 0x74, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2e,
	0x69, 0x6f, 0x2f, 0x76, 0x69, 0x74, 0x65, 0x73, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x74, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vtrpc_proto_rawDescOnce sync.Once
	file_vtrpc_proto_rawDescData = file_vtrpc_proto_rawDesc
)

func file_vtrpc_proto_rawDescGZIP() []byte {
	file_vtrpc_proto_rawDescOnce.Do(func() {
		file_vtrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_vtrpc_proto_rawDescData)
	})
	return file_vtrpc_proto_rawDescData
}

var file_vtrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vtrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vtrpc_proto_goTypes = []interface{}{
	(Code)(0),            // 0: vtrpc.Code
	(LegacyErrorCode)(0), // 1: vtrpc.LegacyErrorCode
	(*CallerID)(nil),     // 2: vtrpc.CallerID
	(*RPCError)(nil),     // 3: vtrpc.RPCError
}
var file_vtrpc_proto_depIdxs = []int32{
	1, // 0: vtrpc.RPCError.legacy_code:type_name -> vtrpc.LegacyErrorCode
	0, // 1: vtrpc.RPCError.code:type_name -> vtrpc.Code
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vtrpc_proto_init() }
func file_vtrpc_proto_init() {
	if File_vtrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vtrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallerID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vtrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RPCError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vtrpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vtrpc_proto_goTypes,
		DependencyIndexes: file_vtrpc_proto_depIdxs,
		EnumInfos:         file_vtrpc_proto_enumTypes,
		MessageInfos:      file_vtrpc_proto_msgTypes,
	}.Build()
	File_vtrpc_proto = out.File
	file_vtrpc_proto_rawDesc = nil
	file_vtrpc_proto_goTypes = nil
	file_vtrpc_proto_depIdxs = nil
}
