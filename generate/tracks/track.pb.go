// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: track.proto

package tracks

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

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[0]
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
	return file_track_proto_rawDescGZIP(), []int{0}
}

type TracksRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracks []*TrackRes `protobuf:"bytes,1,rep,name=Tracks,proto3" json:"Tracks,omitempty"`
}

func (x *TracksRes) Reset() {
	*x = TracksRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TracksRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TracksRes) ProtoMessage() {}

func (x *TracksRes) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TracksRes.ProtoReflect.Descriptor instead.
func (*TracksRes) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{1}
}

func (x *TracksRes) GetTracks() []*TrackRes {
	if x != nil {
		return x.Tracks
	}
	return nil
}

type Track struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompositionID string `protobuf:"bytes,1,opt,name=compositionID,proto3" json:"compositionID,omitempty"`
	UserID        string `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	FileUrl       string `protobuf:"bytes,4,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
}

func (x *Track) Reset() {
	*x = Track{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Track) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Track) ProtoMessage() {}

func (x *Track) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Track.ProtoReflect.Descriptor instead.
func (*Track) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{2}
}

func (x *Track) GetCompositionID() string {
	if x != nil {
		return x.CompositionID
	}
	return ""
}

func (x *Track) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Track) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Track) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type TrackId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackId string `protobuf:"bytes,1,opt,name=trackId,proto3" json:"trackId,omitempty"`
}

func (x *TrackId) Reset() {
	*x = TrackId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackId) ProtoMessage() {}

func (x *TrackId) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackId.ProtoReflect.Descriptor instead.
func (*TrackId) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{3}
}

func (x *TrackId) GetTrackId() string {
	if x != nil {
		return x.TrackId
	}
	return ""
}

type TrackRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TracId        string `protobuf:"bytes,1,opt,name=tracId,proto3" json:"tracId,omitempty"`
	CompositionID string `protobuf:"bytes,2,opt,name=compositionID,proto3" json:"compositionID,omitempty"`
	UserID        string `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	Title         string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	FileUrl       string `protobuf:"bytes,5,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
	CreateAt      string `protobuf:"bytes,6,opt,name=createAt,proto3" json:"createAt,omitempty"`
}

func (x *TrackRes) Reset() {
	*x = TrackRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrackRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackRes) ProtoMessage() {}

func (x *TrackRes) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackRes.ProtoReflect.Descriptor instead.
func (*TrackRes) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{4}
}

func (x *TrackRes) GetTracId() string {
	if x != nil {
		return x.TracId
	}
	return ""
}

func (x *TrackRes) GetCompositionID() string {
	if x != nil {
		return x.CompositionID
	}
	return ""
}

func (x *TrackRes) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *TrackRes) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TrackRes) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *TrackRes) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

type CompositionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompositionID string `protobuf:"bytes,1,opt,name=compositionID,proto3" json:"compositionID,omitempty"`
}

func (x *CompositionID) Reset() {
	*x = CompositionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositionID) ProtoMessage() {}

func (x *CompositionID) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[5]
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
	return file_track_proto_rawDescGZIP(), []int{5}
}

func (x *CompositionID) GetCompositionID() string {
	if x != nil {
		return x.CompositionID
	}
	return ""
}

var File_track_proto protoreflect.FileDescriptor

var file_track_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x09,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x52, 0x06, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x73, 0x22, 0x75, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0xac,
	0x01, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x72, 0x61, 0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61,
	0x63, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x35, 0x0a,
	0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x32, 0xec, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x0c, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x1a, 0x0b, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x75, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x1a, 0x10, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f,
	0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x1a,
	0x0b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_track_proto_rawDescOnce sync.Once
	file_track_proto_rawDescData = file_track_proto_rawDesc
)

func file_track_proto_rawDescGZIP() []byte {
	file_track_proto_rawDescOnce.Do(func() {
		file_track_proto_rawDescData = protoimpl.X.CompressGZIP(file_track_proto_rawDescData)
	})
	return file_track_proto_rawDescData
}

var file_track_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_track_proto_goTypes = []interface{}{
	(*Void)(nil),          // 0: track.Void
	(*TracksRes)(nil),     // 1: track.TracksRes
	(*Track)(nil),         // 2: track.Track
	(*TrackId)(nil),       // 3: track.TrackId
	(*TrackRes)(nil),      // 4: track.TrackRes
	(*CompositionID)(nil), // 5: track.CompositionID
}
var file_track_proto_depIdxs = []int32{
	4, // 0: track.TracksRes.Tracks:type_name -> track.TrackRes
	2, // 1: track.TrackService.Create:input_type -> track.Track
	3, // 2: track.TrackService.GetByID:input_type -> track.TrackId
	5, // 3: track.TrackService.GetBuCompositionId:input_type -> track.CompositionID
	4, // 4: track.TrackService.Update:input_type -> track.TrackRes
	3, // 5: track.TrackService.Delete:input_type -> track.TrackId
	0, // 6: track.TrackService.Create:output_type -> track.Void
	4, // 7: track.TrackService.GetByID:output_type -> track.TrackRes
	1, // 8: track.TrackService.GetBuCompositionId:output_type -> track.TracksRes
	0, // 9: track.TrackService.Update:output_type -> track.Void
	0, // 10: track.TrackService.Delete:output_type -> track.Void
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_track_proto_init() }
func file_track_proto_init() {
	if File_track_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_track_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_track_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TracksRes); i {
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
		file_track_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Track); i {
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
		file_track_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackId); i {
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
		file_track_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrackRes); i {
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
		file_track_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_track_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_track_proto_goTypes,
		DependencyIndexes: file_track_proto_depIdxs,
		MessageInfos:      file_track_proto_msgTypes,
	}.Build()
	File_track_proto = out.File
	file_track_proto_rawDesc = nil
	file_track_proto_goTypes = nil
	file_track_proto_depIdxs = nil
}
