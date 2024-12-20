// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: v1/follow.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type FollowType int32

const (
	FollowType_FOLLOWING FollowType = 0
	FollowType_FOLLOWER  FollowType = 1
	FollowType_BOTH      FollowType = 2
)

// Enum value maps for FollowType.
var (
	FollowType_name = map[int32]string{
		0: "FOLLOWING",
		1: "FOLLOWER",
		2: "BOTH",
	}
	FollowType_value = map[string]int32{
		"FOLLOWING": 0,
		"FOLLOWER":  1,
		"BOTH":      2,
	}
)

func (x FollowType) Enum() *FollowType {
	p := new(FollowType)
	*p = x
	return p
}

func (x FollowType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FollowType) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_follow_proto_enumTypes[0].Descriptor()
}

func (FollowType) Type() protoreflect.EnumType {
	return &file_v1_follow_proto_enumTypes[0]
}

func (x FollowType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FollowType.Descriptor instead.
func (FollowType) EnumDescriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{0}
}

type AddFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TargetUserId int64 `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
}

func (x *AddFollowRequest) Reset() {
	*x = AddFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFollowRequest) ProtoMessage() {}

func (x *AddFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFollowRequest.ProtoReflect.Descriptor instead.
func (*AddFollowRequest) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{0}
}

func (x *AddFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddFollowRequest) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type AddFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *AddFollowResponse) Reset() {
	*x = AddFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFollowResponse) ProtoMessage() {}

func (x *AddFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFollowResponse.ProtoReflect.Descriptor instead.
func (*AddFollowResponse) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{1}
}

func (x *AddFollowResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type RemoveFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TargetUserId int64 `protobuf:"varint,2,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
}

func (x *RemoveFollowRequest) Reset() {
	*x = RemoveFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFollowRequest) ProtoMessage() {}

func (x *RemoveFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFollowRequest.ProtoReflect.Descriptor instead.
func (*RemoveFollowRequest) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RemoveFollowRequest) GetTargetUserId() int64 {
	if x != nil {
		return x.TargetUserId
	}
	return 0
}

type RemoveFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *RemoveFollowResponse) Reset() {
	*x = RemoveFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFollowResponse) ProtoMessage() {}

func (x *RemoveFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFollowResponse.ProtoReflect.Descriptor instead.
func (*RemoveFollowResponse) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveFollowResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ListFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64              `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FollowType FollowType         `protobuf:"varint,2,opt,name=follow_type,json=followType,proto3,enum=shortVideoCoreService.api.v1.FollowType" json:"follow_type,omitempty"`
	Pagination *PaginationRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListFollowingRequest) Reset() {
	*x = ListFollowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingRequest) ProtoMessage() {}

func (x *ListFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingRequest.ProtoReflect.Descriptor instead.
func (*ListFollowingRequest) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{4}
}

func (x *ListFollowingRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListFollowingRequest) GetFollowType() FollowType {
	if x != nil {
		return x.FollowType
	}
	return FollowType_FOLLOWING
}

func (x *ListFollowingRequest) GetPagination() *PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Metadata           `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	UserIdList []int64             `protobuf:"varint,2,rep,packed,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"`
	Pagination *PaginationResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListFollowingResponse) Reset() {
	*x = ListFollowingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingResponse) ProtoMessage() {}

func (x *ListFollowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingResponse.ProtoReflect.Descriptor instead.
func (*ListFollowingResponse) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{5}
}

func (x *ListFollowingResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ListFollowingResponse) GetUserIdList() []int64 {
	if x != nil {
		return x.UserIdList
	}
	return nil
}

func (x *ListFollowingResponse) GetPagination() *PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type IsFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TargetUserIdList []int64 `protobuf:"varint,2,rep,packed,name=target_user_id_list,json=targetUserIdList,proto3" json:"target_user_id_list,omitempty"`
}

func (x *IsFollowingRequest) Reset() {
	*x = IsFollowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFollowingRequest) ProtoMessage() {}

func (x *IsFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFollowingRequest.ProtoReflect.Descriptor instead.
func (*IsFollowingRequest) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{6}
}

func (x *IsFollowingRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *IsFollowingRequest) GetTargetUserIdList() []int64 {
	if x != nil {
		return x.TargetUserIdList
	}
	return nil
}

type IsFollowingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta          *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	FollowingList []int64   `protobuf:"varint,2,rep,packed,name=following_list,json=followingList,proto3" json:"following_list,omitempty"`
}

func (x *IsFollowingResponse) Reset() {
	*x = IsFollowingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsFollowingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsFollowingResponse) ProtoMessage() {}

func (x *IsFollowingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsFollowingResponse.ProtoReflect.Descriptor instead.
func (*IsFollowingResponse) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{7}
}

func (x *IsFollowingResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *IsFollowingResponse) GetFollowingList() []int64 {
	if x != nil {
		return x.FollowingList
	}
	return nil
}

type CountFollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CountFollowRequest) Reset() {
	*x = CountFollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFollowRequest) ProtoMessage() {}

func (x *CountFollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFollowRequest.ProtoReflect.Descriptor instead.
func (*CountFollowRequest) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{8}
}

func (x *CountFollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CountFollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta           *Metadata `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	FollowingCount int64     `protobuf:"varint,2,opt,name=following_count,json=followingCount,proto3" json:"following_count,omitempty"`
	FollowerCount  int64     `protobuf:"varint,3,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
}

func (x *CountFollowResponse) Reset() {
	*x = CountFollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_follow_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountFollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountFollowResponse) ProtoMessage() {}

func (x *CountFollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_follow_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountFollowResponse.ProtoReflect.Descriptor instead.
func (*CountFollowResponse) Descriptor() ([]byte, []int) {
	return file_v1_follow_proto_rawDescGZIP(), []int{9}
}

func (x *CountFollowResponse) GetMeta() *Metadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CountFollowResponse) GetFollowingCount() int64 {
	if x != nil {
		return x.FollowingCount
	}
	return 0
}

func (x *CountFollowResponse) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

var File_v1_follow_proto protoreflect.FileDescriptor

var file_v1_follow_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x13, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x52, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x22, 0xcb, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x4f, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x12, 0x49,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x13, 0x49, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x33, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x54, 0x48, 0x10, 0x02, 0x32, 0xd6, 0x04, 0x0a, 0x0d,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a,
	0x09, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x2e, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x0c, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x31, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x78, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x12, 0x32, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0b,
	0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x2e, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x72, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12,
	0x30, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x2f, 0x44,
	0x6f, 0x75, 0x54, 0x6f, 0x6b, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_follow_proto_rawDescOnce sync.Once
	file_v1_follow_proto_rawDescData = file_v1_follow_proto_rawDesc
)

func file_v1_follow_proto_rawDescGZIP() []byte {
	file_v1_follow_proto_rawDescOnce.Do(func() {
		file_v1_follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_follow_proto_rawDescData)
	})
	return file_v1_follow_proto_rawDescData
}

var file_v1_follow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_follow_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_follow_proto_goTypes = []interface{}{
	(FollowType)(0),               // 0: shortVideoCoreService.api.v1.FollowType
	(*AddFollowRequest)(nil),      // 1: shortVideoCoreService.api.v1.AddFollowRequest
	(*AddFollowResponse)(nil),     // 2: shortVideoCoreService.api.v1.AddFollowResponse
	(*RemoveFollowRequest)(nil),   // 3: shortVideoCoreService.api.v1.RemoveFollowRequest
	(*RemoveFollowResponse)(nil),  // 4: shortVideoCoreService.api.v1.RemoveFollowResponse
	(*ListFollowingRequest)(nil),  // 5: shortVideoCoreService.api.v1.ListFollowingRequest
	(*ListFollowingResponse)(nil), // 6: shortVideoCoreService.api.v1.ListFollowingResponse
	(*IsFollowingRequest)(nil),    // 7: shortVideoCoreService.api.v1.IsFollowingRequest
	(*IsFollowingResponse)(nil),   // 8: shortVideoCoreService.api.v1.IsFollowingResponse
	(*CountFollowRequest)(nil),    // 9: shortVideoCoreService.api.v1.CountFollowRequest
	(*CountFollowResponse)(nil),   // 10: shortVideoCoreService.api.v1.CountFollowResponse
	(*Metadata)(nil),              // 11: shortVideoCoreService.api.v1.Metadata
	(*PaginationRequest)(nil),     // 12: shortVideoCoreService.api.v1.PaginationRequest
	(*PaginationResponse)(nil),    // 13: shortVideoCoreService.api.v1.PaginationResponse
}
var file_v1_follow_proto_depIdxs = []int32{
	11, // 0: shortVideoCoreService.api.v1.AddFollowResponse.meta:type_name -> shortVideoCoreService.api.v1.Metadata
	11, // 1: shortVideoCoreService.api.v1.RemoveFollowResponse.meta:type_name -> shortVideoCoreService.api.v1.Metadata
	0,  // 2: shortVideoCoreService.api.v1.ListFollowingRequest.follow_type:type_name -> shortVideoCoreService.api.v1.FollowType
	12, // 3: shortVideoCoreService.api.v1.ListFollowingRequest.pagination:type_name -> shortVideoCoreService.api.v1.PaginationRequest
	11, // 4: shortVideoCoreService.api.v1.ListFollowingResponse.meta:type_name -> shortVideoCoreService.api.v1.Metadata
	13, // 5: shortVideoCoreService.api.v1.ListFollowingResponse.pagination:type_name -> shortVideoCoreService.api.v1.PaginationResponse
	11, // 6: shortVideoCoreService.api.v1.IsFollowingResponse.meta:type_name -> shortVideoCoreService.api.v1.Metadata
	11, // 7: shortVideoCoreService.api.v1.CountFollowResponse.meta:type_name -> shortVideoCoreService.api.v1.Metadata
	1,  // 8: shortVideoCoreService.api.v1.FollowService.AddFollow:input_type -> shortVideoCoreService.api.v1.AddFollowRequest
	3,  // 9: shortVideoCoreService.api.v1.FollowService.RemoveFollow:input_type -> shortVideoCoreService.api.v1.RemoveFollowRequest
	5,  // 10: shortVideoCoreService.api.v1.FollowService.ListFollowing:input_type -> shortVideoCoreService.api.v1.ListFollowingRequest
	7,  // 11: shortVideoCoreService.api.v1.FollowService.IsFollowing:input_type -> shortVideoCoreService.api.v1.IsFollowingRequest
	9,  // 12: shortVideoCoreService.api.v1.FollowService.CountFollow:input_type -> shortVideoCoreService.api.v1.CountFollowRequest
	2,  // 13: shortVideoCoreService.api.v1.FollowService.AddFollow:output_type -> shortVideoCoreService.api.v1.AddFollowResponse
	4,  // 14: shortVideoCoreService.api.v1.FollowService.RemoveFollow:output_type -> shortVideoCoreService.api.v1.RemoveFollowResponse
	6,  // 15: shortVideoCoreService.api.v1.FollowService.ListFollowing:output_type -> shortVideoCoreService.api.v1.ListFollowingResponse
	8,  // 16: shortVideoCoreService.api.v1.FollowService.IsFollowing:output_type -> shortVideoCoreService.api.v1.IsFollowingResponse
	10, // 17: shortVideoCoreService.api.v1.FollowService.CountFollow:output_type -> shortVideoCoreService.api.v1.CountFollowResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_v1_follow_proto_init() }
func file_v1_follow_proto_init() {
	if File_v1_follow_proto != nil {
		return
	}
	file_v1_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFollowRequest); i {
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
		file_v1_follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFollowResponse); i {
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
		file_v1_follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFollowRequest); i {
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
		file_v1_follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFollowResponse); i {
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
		file_v1_follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingRequest); i {
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
		file_v1_follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingResponse); i {
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
		file_v1_follow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFollowingRequest); i {
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
		file_v1_follow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsFollowingResponse); i {
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
		file_v1_follow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFollowRequest); i {
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
		file_v1_follow_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountFollowResponse); i {
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
			RawDescriptor: file_v1_follow_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_follow_proto_goTypes,
		DependencyIndexes: file_v1_follow_proto_depIdxs,
		EnumInfos:         file_v1_follow_proto_enumTypes,
		MessageInfos:      file_v1_follow_proto_msgTypes,
	}.Build()
	File_v1_follow_proto = out.File
	file_v1_follow_proto_rawDesc = nil
	file_v1_follow_proto_goTypes = nil
	file_v1_follow_proto_depIdxs = nil
}
