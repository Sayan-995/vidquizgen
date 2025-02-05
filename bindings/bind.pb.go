// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: bindings/bind.proto

package bindings

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProblemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TitleSlug     string                 `protobuf:"bytes,1,opt,name=title_slug,json=titleSlug,proto3" json:"title_slug,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProblemRequest) Reset() {
	*x = ProblemRequest{}
	mi := &file_bindings_bind_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProblemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemRequest) ProtoMessage() {}

func (x *ProblemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bindings_bind_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemRequest.ProtoReflect.Descriptor instead.
func (*ProblemRequest) Descriptor() ([]byte, []int) {
	return file_bindings_bind_proto_rawDescGZIP(), []int{0}
}

func (x *ProblemRequest) GetTitleSlug() string {
	if x != nil {
		return x.TitleSlug
	}
	return ""
}

type ProblemStatement struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Statement     string                 `protobuf:"bytes,1,opt,name=statement,proto3" json:"statement,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProblemStatement) Reset() {
	*x = ProblemStatement{}
	mi := &file_bindings_bind_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProblemStatement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProblemStatement) ProtoMessage() {}

func (x *ProblemStatement) ProtoReflect() protoreflect.Message {
	mi := &file_bindings_bind_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProblemStatement.ProtoReflect.Descriptor instead.
func (*ProblemStatement) Descriptor() ([]byte, []int) {
	return file_bindings_bind_proto_rawDescGZIP(), []int{1}
}

func (x *ProblemStatement) GetStatement() string {
	if x != nil {
		return x.Statement
	}
	return ""
}

var File_bindings_bind_proto protoreflect.FileDescriptor

var file_bindings_bind_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x69, 0x6e, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x53, 0x6c, 0x75, 0x67, 0x22, 0x30, 0x0a, 0x10,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x52,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x14, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e,
	0x50, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_bindings_bind_proto_rawDescOnce sync.Once
	file_bindings_bind_proto_rawDescData []byte
)

func file_bindings_bind_proto_rawDescGZIP() []byte {
	file_bindings_bind_proto_rawDescOnce.Do(func() {
		file_bindings_bind_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bindings_bind_proto_rawDesc), len(file_bindings_bind_proto_rawDesc)))
	})
	return file_bindings_bind_proto_rawDescData
}

var file_bindings_bind_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bindings_bind_proto_goTypes = []any{
	(*ProblemRequest)(nil),   // 0: bind.ProblemRequest
	(*ProblemStatement)(nil), // 1: bind.ProblemStatement
}
var file_bindings_bind_proto_depIdxs = []int32{
	0, // 0: bind.StatementService.GetStatement:input_type -> bind.ProblemRequest
	1, // 1: bind.StatementService.GetStatement:output_type -> bind.ProblemStatement
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bindings_bind_proto_init() }
func file_bindings_bind_proto_init() {
	if File_bindings_bind_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bindings_bind_proto_rawDesc), len(file_bindings_bind_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bindings_bind_proto_goTypes,
		DependencyIndexes: file_bindings_bind_proto_depIdxs,
		MessageInfos:      file_bindings_bind_proto_msgTypes,
	}.Build()
	File_bindings_bind_proto = out.File
	file_bindings_bind_proto_goTypes = nil
	file_bindings_bind_proto_depIdxs = nil
}
