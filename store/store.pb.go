// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: store/store.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MethodSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Required string `protobuf:"bytes,1,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *MethodSignature) Reset() {
	*x = MethodSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodSignature) ProtoMessage() {}

func (x *MethodSignature) ProtoReflect() protoreflect.Message {
	mi := &file_store_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodSignature.ProtoReflect.Descriptor instead.
func (*MethodSignature) Descriptor() ([]byte, []int) {
	return file_store_store_proto_rawDescGZIP(), []int{0}
}

func (x *MethodSignature) GetRequired() string {
	if x != nil {
		return x.Required
	}
	return ""
}

var file_store_store_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         8873973,
		Name:          "store.method_signature",
		Tag:           "bytes,8873973,opt,name=method_signature",
		Filename:      "store/store.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional string method_signature = 8873973;
	E_MethodSignature = &file_store_store_proto_extTypes[0]
)

var File_store_store_proto protoreflect.FileDescriptor

var file_store_store_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x0f,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x3a, 0x4c, 0x0a, 0x10, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xf5, 0xcf, 0x9d, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x69, 0x6b, 0x72, 0x73, 0x6e, 0x61,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_store_proto_rawDescOnce sync.Once
	file_store_store_proto_rawDescData = file_store_store_proto_rawDesc
)

func file_store_store_proto_rawDescGZIP() []byte {
	file_store_store_proto_rawDescOnce.Do(func() {
		file_store_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_store_proto_rawDescData)
	})
	return file_store_store_proto_rawDescData
}

var file_store_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_store_proto_goTypes = []interface{}{
	(*MethodSignature)(nil),            // 0: store.MethodSignature
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_store_store_proto_depIdxs = []int32{
	1, // 0: store.method_signature:extendee -> google.protobuf.MethodOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_store_proto_init() }
func file_store_store_proto_init() {
	if File_store_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodSignature); i {
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
			RawDescriptor: file_store_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_store_store_proto_goTypes,
		DependencyIndexes: file_store_store_proto_depIdxs,
		MessageInfos:      file_store_store_proto_msgTypes,
		ExtensionInfos:    file_store_store_proto_extTypes,
	}.Build()
	File_store_store_proto = out.File
	file_store_store_proto_rawDesc = nil
	file_store_store_proto_goTypes = nil
	file_store_store_proto_depIdxs = nil
}
