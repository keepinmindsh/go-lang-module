// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/cqrs_sample/cqrs_sample.proto

package cqrs_sample

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SaveMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SaveMessageRequest) Reset() {
	*x = SaveMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMessageRequest) ProtoMessage() {}

func (x *SaveMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMessageRequest.ProtoReflect.Descriptor instead.
func (*SaveMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{0}
}

func (x *SaveMessageRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SaveMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessagesRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageModels []*MessageModel `protobuf:"bytes,1,rep,name=message_models,json=messageModels,proto3" json:"message_models,omitempty"`
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessagesResponse) GetMessageModels() []*MessageModel {
	if x != nil {
		return x.MessageModels
	}
	return nil
}

type MessageModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *MessageModel) Reset() {
	*x = MessageModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageModel) ProtoMessage() {}

func (x *MessageModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageModel.ProtoReflect.Descriptor instead.
func (*MessageModel) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{3}
}

func (x *MessageModel) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MessageModel) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Key     string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessageRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *GetMessageRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{5}
}

func (x *GetMessageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type SaveMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SaveMessageResponse) Reset() {
	*x = SaveMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveMessageResponse) ProtoMessage() {}

func (x *SaveMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveMessageResponse.ProtoReflect.Descriptor instead.
func (*SaveMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP(), []int{6}
}

func (x *SaveMessageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_cqrs_sample_cqrs_sample_proto protoreflect.FileDescriptor

var file_proto_cqrs_sample_cqrs_sample_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x63, 0x71, 0x72,
	0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x5d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e,
	0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x3a, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xe7, 0x02, 0x0a, 0x0b, 0x43, 0x71, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x74, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x25, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e,
	0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x76, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x6e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x63, 0x71,
	0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x72, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x63,
	0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f,
	0x63, 0x71, 0x72, 0x73, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_cqrs_sample_cqrs_sample_proto_rawDescOnce sync.Once
	file_proto_cqrs_sample_cqrs_sample_proto_rawDescData = file_proto_cqrs_sample_cqrs_sample_proto_rawDesc
)

func file_proto_cqrs_sample_cqrs_sample_proto_rawDescGZIP() []byte {
	file_proto_cqrs_sample_cqrs_sample_proto_rawDescOnce.Do(func() {
		file_proto_cqrs_sample_cqrs_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cqrs_sample_cqrs_sample_proto_rawDescData)
	})
	return file_proto_cqrs_sample_cqrs_sample_proto_rawDescData
}

var file_proto_cqrs_sample_cqrs_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_cqrs_sample_cqrs_sample_proto_goTypes = []interface{}{
	(*SaveMessageRequest)(nil),  // 0: lines.cqrs_sample.SaveMessageRequest
	(*GetMessagesRequest)(nil),  // 1: lines.cqrs_sample.GetMessagesRequest
	(*GetMessagesResponse)(nil), // 2: lines.cqrs_sample.GetMessagesResponse
	(*MessageModel)(nil),        // 3: lines.cqrs_sample.MessageModel
	(*GetMessageRequest)(nil),   // 4: lines.cqrs_sample.GetMessageRequest
	(*GetMessageResponse)(nil),  // 5: lines.cqrs_sample.GetMessageResponse
	(*SaveMessageResponse)(nil), // 6: lines.cqrs_sample.SaveMessageResponse
}
var file_proto_cqrs_sample_cqrs_sample_proto_depIdxs = []int32{
	3, // 0: lines.cqrs_sample.GetMessagesResponse.message_models:type_name -> lines.cqrs_sample.MessageModel
	0, // 1: lines.cqrs_sample.CqrsService.SaveMessage:input_type -> lines.cqrs_sample.SaveMessageRequest
	4, // 2: lines.cqrs_sample.CqrsService.GetMessage:input_type -> lines.cqrs_sample.GetMessageRequest
	1, // 3: lines.cqrs_sample.CqrsService.GetMessages:input_type -> lines.cqrs_sample.GetMessagesRequest
	6, // 4: lines.cqrs_sample.CqrsService.SaveMessage:output_type -> lines.cqrs_sample.SaveMessageResponse
	5, // 5: lines.cqrs_sample.CqrsService.GetMessage:output_type -> lines.cqrs_sample.GetMessageResponse
	2, // 6: lines.cqrs_sample.CqrsService.GetMessages:output_type -> lines.cqrs_sample.GetMessagesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cqrs_sample_cqrs_sample_proto_init() }
func file_proto_cqrs_sample_cqrs_sample_proto_init() {
	if File_proto_cqrs_sample_cqrs_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMessageRequest); i {
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
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesRequest); i {
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
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesResponse); i {
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
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageModel); i {
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
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
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
		file_proto_cqrs_sample_cqrs_sample_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveMessageResponse); i {
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
			RawDescriptor: file_proto_cqrs_sample_cqrs_sample_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cqrs_sample_cqrs_sample_proto_goTypes,
		DependencyIndexes: file_proto_cqrs_sample_cqrs_sample_proto_depIdxs,
		MessageInfos:      file_proto_cqrs_sample_cqrs_sample_proto_msgTypes,
	}.Build()
	File_proto_cqrs_sample_cqrs_sample_proto = out.File
	file_proto_cqrs_sample_cqrs_sample_proto_rawDesc = nil
	file_proto_cqrs_sample_cqrs_sample_proto_goTypes = nil
	file_proto_cqrs_sample_cqrs_sample_proto_depIdxs = nil
}