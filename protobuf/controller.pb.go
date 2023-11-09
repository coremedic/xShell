// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobuf/controller.proto

package protobuf

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

// Request to list all shells
type ListShellsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListShellsRequest) Reset() {
	*x = ListShellsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShellsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShellsRequest) ProtoMessage() {}

func (x *ListShellsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShellsRequest.ProtoReflect.Descriptor instead.
func (*ListShellsRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{0}
}

// Request to generate a new client
type NewClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *NewClientRequest) Reset() {
	*x = NewClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewClientRequest) ProtoMessage() {}

func (x *NewClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewClientRequest.ProtoReflect.Descriptor instead.
func (*NewClientRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{1}
}

func (x *NewClientRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// Request for shells log data
type ShellLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shell string `protobuf:"bytes,1,opt,name=shell,proto3" json:"shell,omitempty"`
}

func (x *ShellLogRequest) Reset() {
	*x = ShellLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellLogRequest) ProtoMessage() {}

func (x *ShellLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellLogRequest.ProtoReflect.Descriptor instead.
func (*ShellLogRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{2}
}

func (x *ShellLogRequest) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

// Request to stream shells log data
type StreamShellLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shell string `protobuf:"bytes,1,opt,name=shell,proto3" json:"shell,omitempty"`
}

func (x *StreamShellLogRequest) Reset() {
	*x = StreamShellLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamShellLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamShellLogRequest) ProtoMessage() {}

func (x *StreamShellLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamShellLogRequest.ProtoReflect.Descriptor instead.
func (*StreamShellLogRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{3}
}

func (x *StreamShellLogRequest) GetShell() string {
	if x != nil {
		return x.Shell
	}
	return ""
}

// Request to execute command on shell host
type ShellExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
}

func (x *ShellExecRequest) Reset() {
	*x = ShellExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellExecRequest) ProtoMessage() {}

func (x *ShellExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellExecRequest.ProtoReflect.Descriptor instead.
func (*ShellExecRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{4}
}

func (x *ShellExecRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

// Response with list of all active shells
type ListShellsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shells []*ShellInfo `protobuf:"bytes,1,rep,name=shells,proto3" json:"shells,omitempty"`
}

func (x *ListShellsResponse) Reset() {
	*x = ListShellsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShellsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShellsResponse) ProtoMessage() {}

func (x *ListShellsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShellsResponse.ProtoReflect.Descriptor instead.
func (*ListShellsResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{5}
}

func (x *ListShellsResponse) GetShells() []*ShellInfo {
	if x != nil {
		return x.Shells
	}
	return nil
}

// Response for a NewClientRequest, returns client cert as byte array
type NewClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert []byte `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *NewClientResponse) Reset() {
	*x = NewClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewClientResponse) ProtoMessage() {}

func (x *NewClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewClientResponse.ProtoReflect.Descriptor instead.
func (*NewClientResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{6}
}

func (x *NewClientResponse) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

// Response for shell log
type ShellLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShellLog []byte `protobuf:"bytes,1,opt,name=shellLog,proto3" json:"shellLog,omitempty"`
}

func (x *ShellLogResponse) Reset() {
	*x = ShellLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellLogResponse) ProtoMessage() {}

func (x *ShellLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellLogResponse.ProtoReflect.Descriptor instead.
func (*ShellLogResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{7}
}

func (x *ShellLogResponse) GetShellLog() []byte {
	if x != nil {
		return x.ShellLog
	}
	return nil
}

// Response to stream shells log
type StreamShellLogResposne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Log string `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
}

func (x *StreamShellLogResposne) Reset() {
	*x = StreamShellLogResposne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamShellLogResposne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamShellLogResposne) ProtoMessage() {}

func (x *StreamShellLogResposne) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamShellLogResposne.ProtoReflect.Descriptor instead.
func (*StreamShellLogResposne) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{8}
}

func (x *StreamShellLogResposne) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

// Shell info containing ID, IP, and last call back time
type ShellInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	LastCall int64  `protobuf:"varint,3,opt,name=last_call,json=lastCall,proto3" json:"last_call,omitempty"`
}

func (x *ShellInfo) Reset() {
	*x = ShellInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShellInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShellInfo) ProtoMessage() {}

func (x *ShellInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShellInfo.ProtoReflect.Descriptor instead.
func (*ShellInfo) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{9}
}

func (x *ShellInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShellInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *ShellInfo) GetLastCall() int64 {
	if x != nil {
		return x.LastCall
	}
	return 0
}

// Empty message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_controller_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_controller_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protobuf_controller_proto_rawDescGZIP(), []int{10}
}

var File_protobuf_controller_proto protoreflect.FileDescriptor

var file_protobuf_controller_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65,
	0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x10, 0x4e, 0x65,
	0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x53, 0x68,
	0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68,
	0x65, 0x6c, 0x6c, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x65,
	0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x65, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x65,
	0x6c, 0x6c, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x22, 0x41, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x73, 0x68, 0x65,
	0x6c, 0x6c, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x22, 0x2e, 0x0a, 0x10,
	0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x22, 0x2a, 0x0a, 0x16,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x73, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x48, 0x0a, 0x09, 0x53, 0x68, 0x65, 0x6c,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x61,
	0x6c, 0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xf6, 0x02, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x73, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x68, 0x65, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x65, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x53, 0x68,
	0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x68, 0x65,
	0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x12,
	0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73,
	0x6e, 0x65, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x09, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x45, 0x78, 0x65,
	0x63, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x68, 0x65,
	0x6c, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44,
	0x0a, 0x09, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x78, 0x53, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_controller_proto_rawDescOnce sync.Once
	file_protobuf_controller_proto_rawDescData = file_protobuf_controller_proto_rawDesc
)

func file_protobuf_controller_proto_rawDescGZIP() []byte {
	file_protobuf_controller_proto_rawDescOnce.Do(func() {
		file_protobuf_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_controller_proto_rawDescData)
	})
	return file_protobuf_controller_proto_rawDescData
}

var file_protobuf_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_protobuf_controller_proto_goTypes = []interface{}{
	(*ListShellsRequest)(nil),      // 0: protobuf.ListShellsRequest
	(*NewClientRequest)(nil),       // 1: protobuf.NewClientRequest
	(*ShellLogRequest)(nil),        // 2: protobuf.ShellLogRequest
	(*StreamShellLogRequest)(nil),  // 3: protobuf.StreamShellLogRequest
	(*ShellExecRequest)(nil),       // 4: protobuf.ShellExecRequest
	(*ListShellsResponse)(nil),     // 5: protobuf.ListShellsResponse
	(*NewClientResponse)(nil),      // 6: protobuf.NewClientResponse
	(*ShellLogResponse)(nil),       // 7: protobuf.ShellLogResponse
	(*StreamShellLogResposne)(nil), // 8: protobuf.StreamShellLogResposne
	(*ShellInfo)(nil),              // 9: protobuf.ShellInfo
	(*Empty)(nil),                  // 10: protobuf.Empty
}
var file_protobuf_controller_proto_depIdxs = []int32{
	9,  // 0: protobuf.ListShellsResponse.shells:type_name -> protobuf.ShellInfo
	0,  // 1: protobuf.ControllerService.ListShells:input_type -> protobuf.ListShellsRequest
	2,  // 2: protobuf.ControllerService.ShellLog:input_type -> protobuf.ShellLogRequest
	3,  // 3: protobuf.ControllerService.StreamShellLog:input_type -> protobuf.StreamShellLogRequest
	4,  // 4: protobuf.ControllerService.ShellExec:input_type -> protobuf.ShellExecRequest
	1,  // 5: protobuf.ControllerService.NewClient:input_type -> protobuf.NewClientRequest
	5,  // 6: protobuf.ControllerService.ListShells:output_type -> protobuf.ListShellsResponse
	7,  // 7: protobuf.ControllerService.ShellLog:output_type -> protobuf.ShellLogResponse
	8,  // 8: protobuf.ControllerService.StreamShellLog:output_type -> protobuf.StreamShellLogResposne
	10, // 9: protobuf.ControllerService.ShellExec:output_type -> protobuf.Empty
	6,  // 10: protobuf.ControllerService.NewClient:output_type -> protobuf.NewClientResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_controller_proto_init() }
func file_protobuf_controller_proto_init() {
	if File_protobuf_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShellsRequest); i {
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
		file_protobuf_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewClientRequest); i {
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
		file_protobuf_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellLogRequest); i {
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
		file_protobuf_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamShellLogRequest); i {
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
		file_protobuf_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellExecRequest); i {
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
		file_protobuf_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShellsResponse); i {
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
		file_protobuf_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewClientResponse); i {
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
		file_protobuf_controller_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellLogResponse); i {
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
		file_protobuf_controller_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamShellLogResposne); i {
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
		file_protobuf_controller_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShellInfo); i {
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
		file_protobuf_controller_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_protobuf_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_controller_proto_goTypes,
		DependencyIndexes: file_protobuf_controller_proto_depIdxs,
		MessageInfos:      file_protobuf_controller_proto_msgTypes,
	}.Build()
	File_protobuf_controller_proto = out.File
	file_protobuf_controller_proto_rawDesc = nil
	file_protobuf_controller_proto_goTypes = nil
	file_protobuf_controller_proto_depIdxs = nil
}
