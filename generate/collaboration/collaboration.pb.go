// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: collaboration.proto

package collaboration

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

type CollaborationRole int32

const (
	CollaborationRole_OWNER        CollaborationRole = 0
	CollaborationRole_COLLABORATOR CollaborationRole = 1
	CollaborationRole_VIEWER       CollaborationRole = 2
)

// Enum value maps for CollaborationRole.
var (
	CollaborationRole_name = map[int32]string{
		0: "OWNER",
		1: "COLLABORATOR",
		2: "VIEWER",
	}
	CollaborationRole_value = map[string]int32{
		"OWNER":        0,
		"COLLABORATOR": 1,
		"VIEWER":       2,
	}
)

func (x CollaborationRole) Enum() *CollaborationRole {
	p := new(CollaborationRole)
	*p = x
	return p
}

func (x CollaborationRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CollaborationRole) Descriptor() protoreflect.EnumDescriptor {
	return file_collaboration_proto_enumTypes[0].Descriptor()
}

func (CollaborationRole) Type() protoreflect.EnumType {
	return &file_collaboration_proto_enumTypes[0]
}

func (x CollaborationRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CollaborationRole.Descriptor instead.
func (CollaborationRole) EnumDescriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{0}
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaboration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_collaboration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{0}
}

type CollaborationID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollaborationId string `protobuf:"bytes,1,opt,name=collaboration_id,json=collaborationId,proto3" json:"collaboration_id,omitempty"`
}

func (x *CollaborationID) Reset() {
	*x = CollaborationID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaboration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollaborationID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollaborationID) ProtoMessage() {}

func (x *CollaborationID) ProtoReflect() protoreflect.Message {
	mi := &file_collaboration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollaborationID.ProtoReflect.Descriptor instead.
func (*CollaborationID) Descriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{1}
}

func (x *CollaborationID) GetCollaborationId() string {
	if x != nil {
		return x.CollaborationId
	}
	return ""
}

type Collaboration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompositionId string            `protobuf:"bytes,1,opt,name=composition_id,json=compositionId,proto3" json:"composition_id,omitempty"`
	UserId        string            `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          CollaborationRole `protobuf:"varint,3,opt,name=role,proto3,enum=collaboration.CollaborationRole" json:"role,omitempty"`
	JoinedAt      string            `protobuf:"bytes,4,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
}

func (x *Collaboration) Reset() {
	*x = Collaboration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaboration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collaboration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collaboration) ProtoMessage() {}

func (x *Collaboration) ProtoReflect() protoreflect.Message {
	mi := &file_collaboration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collaboration.ProtoReflect.Descriptor instead.
func (*Collaboration) Descriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{2}
}

func (x *Collaboration) GetCompositionId() string {
	if x != nil {
		return x.CompositionId
	}
	return ""
}

func (x *Collaboration) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Collaboration) GetRole() CollaborationRole {
	if x != nil {
		return x.Role
	}
	return CollaborationRole_OWNER
}

func (x *Collaboration) GetJoinedAt() string {
	if x != nil {
		return x.JoinedAt
	}
	return ""
}

type CompositionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompositionId string `protobuf:"bytes,1,opt,name=composition_id,json=compositionId,proto3" json:"composition_id,omitempty"`
}

func (x *CompositionID) Reset() {
	*x = CompositionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaboration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositionID) ProtoMessage() {}

func (x *CompositionID) ProtoReflect() protoreflect.Message {
	mi := &file_collaboration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositionID.ProtoReflect.Descriptor instead.
func (*CompositionID) Descriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{3}
}

func (x *CompositionID) GetCompositionId() string {
	if x != nil {
		return x.CompositionId
	}
	return ""
}

type CollaborationRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompositionId string            `protobuf:"bytes,2,opt,name=composition_id,json=compositionId,proto3" json:"composition_id,omitempty"`
	UserId        string            `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          CollaborationRole `protobuf:"varint,4,opt,name=role,proto3,enum=collaboration.CollaborationRole" json:"role,omitempty"`
	JoinedAt      string            `protobuf:"bytes,5,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
}

func (x *CollaborationRes) Reset() {
	*x = CollaborationRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaboration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollaborationRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollaborationRes) ProtoMessage() {}

func (x *CollaborationRes) ProtoReflect() protoreflect.Message {
	mi := &file_collaboration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollaborationRes.ProtoReflect.Descriptor instead.
func (*CollaborationRes) Descriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{4}
}

func (x *CollaborationRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CollaborationRes) GetCompositionId() string {
	if x != nil {
		return x.CompositionId
	}
	return ""
}

func (x *CollaborationRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CollaborationRes) GetRole() CollaborationRole {
	if x != nil {
		return x.Role
	}
	return CollaborationRole_OWNER
}

func (x *CollaborationRes) GetJoinedAt() string {
	if x != nil {
		return x.JoinedAt
	}
	return ""
}

type CollaborationResList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collaborations []*CollaborationRes `protobuf:"bytes,1,rep,name=collaborations,proto3" json:"collaborations,omitempty"`
}

func (x *CollaborationResList) Reset() {
	*x = CollaborationResList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collaboration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollaborationResList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollaborationResList) ProtoMessage() {}

func (x *CollaborationResList) ProtoReflect() protoreflect.Message {
	mi := &file_collaboration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollaborationResList.ProtoReflect.Descriptor instead.
func (*CollaborationResList) Descriptor() ([]byte, []int) {
	return file_collaboration_proto_rawDescGZIP(), []int{5}
}

func (x *CollaborationResList) GetCollaborations() []*CollaborationRes {
	if x != nil {
		return x.Collaborations
	}
	return nil
}

var File_collaboration_proto protoreflect.FileDescriptor

var file_collaboration_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0f,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x0d, 0x43,
	0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x36, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb5, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x5f, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x61,
	0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2a, 0x3c, 0x0a, 0x11, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4c, 0x4c, 0x41, 0x42, 0x4f, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x02, 0x32, 0xb7,
	0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1e, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a,
	0x1f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collaboration_proto_rawDescOnce sync.Once
	file_collaboration_proto_rawDescData = file_collaboration_proto_rawDesc
)

func file_collaboration_proto_rawDescGZIP() []byte {
	file_collaboration_proto_rawDescOnce.Do(func() {
		file_collaboration_proto_rawDescData = protoimpl.X.CompressGZIP(file_collaboration_proto_rawDescData)
	})
	return file_collaboration_proto_rawDescData
}

var file_collaboration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_collaboration_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_collaboration_proto_goTypes = []interface{}{
	(CollaborationRole)(0),       // 0: collaboration.CollaborationRole
	(*Void)(nil),                 // 1: collaboration.Void
	(*CollaborationID)(nil),      // 2: collaboration.CollaborationID
	(*Collaboration)(nil),        // 3: collaboration.Collaboration
	(*CompositionID)(nil),        // 4: collaboration.CompositionID
	(*CollaborationRes)(nil),     // 5: collaboration.CollaborationRes
	(*CollaborationResList)(nil), // 6: collaboration.CollaborationResList
}
var file_collaboration_proto_depIdxs = []int32{
	0, // 0: collaboration.Collaboration.role:type_name -> collaboration.CollaborationRole
	0, // 1: collaboration.CollaborationRes.role:type_name -> collaboration.CollaborationRole
	5, // 2: collaboration.CollaborationResList.collaborations:type_name -> collaboration.CollaborationRes
	3, // 3: collaboration.CollaborationService.Create:input_type -> collaboration.Collaboration
	2, // 4: collaboration.CollaborationService.GetById:input_type -> collaboration.CollaborationID
	4, // 5: collaboration.CollaborationService.GetByCompositionId:input_type -> collaboration.CompositionID
	2, // 6: collaboration.CollaborationService.Delete:input_type -> collaboration.CollaborationID
	1, // 7: collaboration.CollaborationService.Create:output_type -> collaboration.Void
	5, // 8: collaboration.CollaborationService.GetById:output_type -> collaboration.CollaborationRes
	6, // 9: collaboration.CollaborationService.GetByCompositionId:output_type -> collaboration.CollaborationResList
	1, // 10: collaboration.CollaborationService.Delete:output_type -> collaboration.Void
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_collaboration_proto_init() }
func file_collaboration_proto_init() {
	if File_collaboration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collaboration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_collaboration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollaborationID); i {
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
		file_collaboration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collaboration); i {
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
		file_collaboration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositionID); i {
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
		file_collaboration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollaborationRes); i {
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
		file_collaboration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollaborationResList); i {
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
			RawDescriptor: file_collaboration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collaboration_proto_goTypes,
		DependencyIndexes: file_collaboration_proto_depIdxs,
		EnumInfos:         file_collaboration_proto_enumTypes,
		MessageInfos:      file_collaboration_proto_msgTypes,
	}.Build()
	File_collaboration_proto = out.File
	file_collaboration_proto_rawDesc = nil
	file_collaboration_proto_goTypes = nil
	file_collaboration_proto_depIdxs = nil
}
