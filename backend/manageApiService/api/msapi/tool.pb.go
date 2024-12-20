// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: msapi/tool.proto

package msapi

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type BatchCreateAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BatchCreateAccountsRequest) Reset() {
	*x = BatchCreateAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_tool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateAccountsRequest) ProtoMessage() {}

func (x *BatchCreateAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_tool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateAccountsRequest.ProtoReflect.Descriptor instead.
func (*BatchCreateAccountsRequest) Descriptor() ([]byte, []int) {
	return file_msapi_tool_proto_rawDescGZIP(), []int{0}
}

type BatchCreateAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BatchCreateAccountsResponse) Reset() {
	*x = BatchCreateAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msapi_tool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateAccountsResponse) ProtoMessage() {}

func (x *BatchCreateAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_msapi_tool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateAccountsResponse.ProtoReflect.Descriptor instead.
func (*BatchCreateAccountsResponse) Descriptor() ([]byte, []int) {
	return file_msapi_tool_proto_rawDescGZIP(), []int{1}
}

var File_msapi_tool_proto protoreflect.FileDescriptor

var file_msapi_tool_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x87, 0x01, 0x0a, 0x0b, 0x54, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x78, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x3a, 0x01, 0x2a, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x7a,
	0x65, 0x6e, 0x69, 0x74, 0x68, 0x2f, 0x44, 0x6f, 0x75, 0x54, 0x6f, 0x6b, 0x2f, 0x2e, 0x2e, 0x2e,
	0x3b, 0x6d, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msapi_tool_proto_rawDescOnce sync.Once
	file_msapi_tool_proto_rawDescData = file_msapi_tool_proto_rawDesc
)

func file_msapi_tool_proto_rawDescGZIP() []byte {
	file_msapi_tool_proto_rawDescOnce.Do(func() {
		file_msapi_tool_proto_rawDescData = protoimpl.X.CompressGZIP(file_msapi_tool_proto_rawDescData)
	})
	return file_msapi_tool_proto_rawDescData
}

var file_msapi_tool_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msapi_tool_proto_goTypes = []interface{}{
	(*BatchCreateAccountsRequest)(nil),  // 0: api.BatchCreateAccountsRequest
	(*BatchCreateAccountsResponse)(nil), // 1: api.BatchCreateAccountsResponse
}
var file_msapi_tool_proto_depIdxs = []int32{
	0, // 0: api.ToolService.BatchCreateAccounts:input_type -> api.BatchCreateAccountsRequest
	1, // 1: api.ToolService.BatchCreateAccounts:output_type -> api.BatchCreateAccountsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msapi_tool_proto_init() }
func file_msapi_tool_proto_init() {
	if File_msapi_tool_proto != nil {
		return
	}
	file_msapi_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_msapi_tool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateAccountsRequest); i {
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
		file_msapi_tool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreateAccountsResponse); i {
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
			RawDescriptor: file_msapi_tool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msapi_tool_proto_goTypes,
		DependencyIndexes: file_msapi_tool_proto_depIdxs,
		MessageInfos:      file_msapi_tool_proto_msgTypes,
	}.Build()
	File_msapi_tool_proto = out.File
	file_msapi_tool_proto_rawDesc = nil
	file_msapi_tool_proto_goTypes = nil
	file_msapi_tool_proto_depIdxs = nil
}
