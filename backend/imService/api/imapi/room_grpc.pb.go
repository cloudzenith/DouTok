// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: imapi/room.proto

package imapi

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

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	// 创建聊天室
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// 删除聊天室
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
	// 将account加入聊天室
	AddAccount2Room(ctx context.Context, in *AddAccount2RoomRequest, opts ...grpc.CallOption) (*AddAccount2RoomResponse, error)
	// 将account从聊天室中移除
	RmAccountFromRoom(ctx context.Context, in *RmAccountFromRoomRequest, opts ...grpc.CallOption) (*RmAccountFromRoomResponse, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.RoomService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	out := new(RemoveResponse)
	err := c.cc.Invoke(ctx, "/api.RoomService/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) AddAccount2Room(ctx context.Context, in *AddAccount2RoomRequest, opts ...grpc.CallOption) (*AddAccount2RoomResponse, error) {
	out := new(AddAccount2RoomResponse)
	err := c.cc.Invoke(ctx, "/api.RoomService/AddAccount2Room", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) RmAccountFromRoom(ctx context.Context, in *RmAccountFromRoomRequest, opts ...grpc.CallOption) (*RmAccountFromRoomResponse, error) {
	out := new(RmAccountFromRoomResponse)
	err := c.cc.Invoke(ctx, "/api.RoomService/RmAccountFromRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations should embed UnimplementedRoomServiceServer
// for forward compatibility
type RoomServiceServer interface {
	// 创建聊天室
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// 删除聊天室
	Remove(context.Context, *RemoveRequest) (*RemoveResponse, error)
	// 将account加入聊天室
	AddAccount2Room(context.Context, *AddAccount2RoomRequest) (*AddAccount2RoomResponse, error)
	// 将account从聊天室中移除
	RmAccountFromRoom(context.Context, *RmAccountFromRoomRequest) (*RmAccountFromRoomResponse, error)
}

// UnimplementedRoomServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (UnimplementedRoomServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRoomServiceServer) Remove(context.Context, *RemoveRequest) (*RemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedRoomServiceServer) AddAccount2Room(context.Context, *AddAccount2RoomRequest) (*AddAccount2RoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccount2Room not implemented")
}
func (UnimplementedRoomServiceServer) RmAccountFromRoom(context.Context, *RmAccountFromRoomRequest) (*RmAccountFromRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RmAccountFromRoom not implemented")
}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s grpc.ServiceRegistrar, srv RoomServiceServer) {
	s.RegisterService(&RoomService_ServiceDesc, srv)
}

func _RoomService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoomService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoomService/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_AddAccount2Room_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccount2RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).AddAccount2Room(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoomService/AddAccount2Room",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).AddAccount2Room(ctx, req.(*AddAccount2RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_RmAccountFromRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RmAccountFromRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).RmAccountFromRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RoomService/RmAccountFromRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).RmAccountFromRoom(ctx, req.(*RmAccountFromRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomService_ServiceDesc is the grpc.ServiceDesc for RoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RoomService_Create_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _RoomService_Remove_Handler,
		},
		{
			MethodName: "AddAccount2Room",
			Handler:    _RoomService_AddAccount2Room_Handler,
		},
		{
			MethodName: "RmAccountFromRoom",
			Handler:    _RoomService_RmAccountFromRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "imapi/room.proto",
}
