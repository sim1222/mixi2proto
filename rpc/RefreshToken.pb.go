// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.2
// source: RefreshToken.proto

package rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_RefreshToken_proto protoreflect.FileDescriptor

var file_RefreshToken_proto_rawDesc = []byte{
	0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x69, 0x6f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x73, 0x69, 0x6d, 0x31, 0x32, 0x32, 0x32, 0x2e, 0x6d, 0x69, 0x78, 0x69, 0x32, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x70, 0x63, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x6d, 0x31, 0x32, 0x32, 0x32, 0x2f, 0x6d, 0x69, 0x78,
	0x69, 0x32, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_RefreshToken_proto_goTypes = []any{}
var file_RefreshToken_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RefreshToken_proto_init() }
func file_RefreshToken_proto_init() {
	if File_RefreshToken_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RefreshToken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RefreshToken_proto_goTypes,
		DependencyIndexes: file_RefreshToken_proto_depIdxs,
	}.Build()
	File_RefreshToken_proto = out.File
	file_RefreshToken_proto_rawDesc = nil
	file_RefreshToken_proto_goTypes = nil
	file_RefreshToken_proto_depIdxs = nil
}
