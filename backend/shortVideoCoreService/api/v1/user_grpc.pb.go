// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	GetUserByIdList(ctx context.Context, in *GetUserByIdListRequest, opts ...grpc.CallOption) (*GetUserByIdListResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	out := new(GetUserInfoResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.UserService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.UserService/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByIdList(ctx context.Context, in *GetUserByIdListRequest, opts ...grpc.CallOption) (*GetUserByIdListResponse, error) {
	out := new(GetUserByIdListResponse)
	err := c.cc.Invoke(ctx, "/shortVideoCoreService.api.v1.UserService/GetUserByIdList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	GetUserByIdList(context.Context, *GetUserByIdListRequest) (*GetUserByIdListResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserByIdList(context.Context, *GetUserByIdListRequest) (*GetUserByIdListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByIdList not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.UserService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GetUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.UserService/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortVideoCoreService.api.v1.UserService/GetUserByIdList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByIdList(ctx, req.(*GetUserByIdListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortVideoCoreService.api.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "GetUserByIdList",
			Handler:    _UserService_GetUserByIdList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/user.proto",
}
