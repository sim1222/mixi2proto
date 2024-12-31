// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.2
// source: models/Post.proto

package models

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

type PostSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostUuid     string    `protobuf:"bytes,1,opt,name=post_uuid,json=postUuid,proto3" json:"post_uuid,omitempty"`
	UnixTimeUsec string    `protobuf:"bytes,2,opt,name=unix_time_usec,json=unixTimeUsec,proto3" json:"unix_time_usec,omitempty"`
	PostTime     *PostTime `protobuf:"bytes,3,opt,name=post_time,json=postTime,proto3" json:"post_time,omitempty"`
	UserUuid     string    `protobuf:"bytes,4,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Body         string    `protobuf:"bytes,12,opt,name=body,proto3" json:"body,omitempty"`
	// 81 80 04: hoppy
	// 82 80 04: shake
	// 83 80 04: super shake
	//
	// 01: big
	// 02: super big
	//
	// 81 80 08: angry
	// 82 80 08: down
	// 83 80 08: shine
	// 84 80 08: love
	// 85 80 08: heartbeat
	//
	// 01 81 80 04 81 80 08: hoppy big angly
	// concat all the above
	DecorationFlag []byte `protobuf:"bytes,28,opt,name=decoration_flag,json=decorationFlag,proto3" json:"decoration_flag,omitempty"`
}

func (x *PostSummary) Reset() {
	*x = PostSummary{}
	mi := &file_models_Post_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSummary) ProtoMessage() {}

func (x *PostSummary) ProtoReflect() protoreflect.Message {
	mi := &file_models_Post_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSummary.ProtoReflect.Descriptor instead.
func (*PostSummary) Descriptor() ([]byte, []int) {
	return file_models_Post_proto_rawDescGZIP(), []int{0}
}

func (x *PostSummary) GetPostUuid() string {
	if x != nil {
		return x.PostUuid
	}
	return ""
}

func (x *PostSummary) GetUnixTimeUsec() string {
	if x != nil {
		return x.UnixTimeUsec
	}
	return ""
}

func (x *PostSummary) GetPostTime() *PostTime {
	if x != nil {
		return x.PostTime
	}
	return nil
}

func (x *PostSummary) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *PostSummary) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *PostSummary) GetDecorationFlag() []byte {
	if x != nil {
		return x.DecorationFlag
	}
	return nil
}

type PostTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnixTimeSec uint32 `protobuf:"varint,1,opt,name=unix_time_sec,json=unixTimeSec,proto3" json:"unix_time_sec,omitempty"`
}

func (x *PostTime) Reset() {
	*x = PostTime{}
	mi := &file_models_Post_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTime) ProtoMessage() {}

func (x *PostTime) ProtoReflect() protoreflect.Message {
	mi := &file_models_Post_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTime.ProtoReflect.Descriptor instead.
func (*PostTime) Descriptor() ([]byte, []int) {
	return file_models_Post_proto_rawDescGZIP(), []int{1}
}

func (x *PostTime) GetUnixTimeSec() uint32 {
	if x != nil {
		return x.UnixTimeSec
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostSummary *PostSummary `protobuf:"bytes,1,opt,name=post_summary,json=postSummary,proto3" json:"post_summary,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_models_Post_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_models_Post_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_models_Post_proto_rawDescGZIP(), []int{2}
}

func (x *Post) GetPostSummary() *PostSummary {
	if x != nil {
		return x.PostSummary
	}
	return nil
}

var File_models_Post_proto protoreflect.FileDescriptor

var file_models_Post_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x50, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x27, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x73,
	0x69, 0x6d, 0x31, 0x32, 0x32, 0x32, 0x2e, 0x6d, 0x69, 0x78, 0x69, 0x32, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0xfa, 0x01, 0x0a,
	0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x6e, 0x69,
	0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63, 0x12,
	0x4e, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x31, 0x2e, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x73,
	0x69, 0x6d, 0x31, 0x32, 0x32, 0x32, 0x2e, 0x6d, 0x69, 0x78, 0x69, 0x32, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x6f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x22, 0x2e, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x6e,
	0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x22, 0x5f, 0x0a, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x12, 0x57, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x73, 0x69, 0x6d, 0x31, 0x32, 0x32, 0x32, 0x2e, 0x6d, 0x69, 0x78, 0x69,
	0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0b, 0x70,
	0x6f, 0x73, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x31, 0x32, 0x32, 0x32,
	0x2f, 0x6d, 0x69, 0x78, 0x69, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_Post_proto_rawDescOnce sync.Once
	file_models_Post_proto_rawDescData = file_models_Post_proto_rawDesc
)

func file_models_Post_proto_rawDescGZIP() []byte {
	file_models_Post_proto_rawDescOnce.Do(func() {
		file_models_Post_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_Post_proto_rawDescData)
	})
	return file_models_Post_proto_rawDescData
}

var file_models_Post_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_models_Post_proto_goTypes = []any{
	(*PostSummary)(nil), // 0: io.github.sim1222.mixi2proto.rpc.models.PostSummary
	(*PostTime)(nil),    // 1: io.github.sim1222.mixi2proto.rpc.models.PostTime
	(*Post)(nil),        // 2: io.github.sim1222.mixi2proto.rpc.models.Post
}
var file_models_Post_proto_depIdxs = []int32{
	1, // 0: io.github.sim1222.mixi2proto.rpc.models.PostSummary.post_time:type_name -> io.github.sim1222.mixi2proto.rpc.models.PostTime
	0, // 1: io.github.sim1222.mixi2proto.rpc.models.Post.post_summary:type_name -> io.github.sim1222.mixi2proto.rpc.models.PostSummary
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_Post_proto_init() }
func file_models_Post_proto_init() {
	if File_models_Post_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_Post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_Post_proto_goTypes,
		DependencyIndexes: file_models_Post_proto_depIdxs,
		MessageInfos:      file_models_Post_proto_msgTypes,
	}.Build()
	File_models_Post_proto = out.File
	file_models_Post_proto_rawDesc = nil
	file_models_Post_proto_goTypes = nil
	file_models_Post_proto_depIdxs = nil
}