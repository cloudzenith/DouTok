// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: promotion.proto

package api

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

// PromotionServiceClient is the client API for PromotionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionServiceClient interface {
	PromotionCreate(ctx context.Context, in *PromotionCreateRequest, opts ...grpc.CallOption) (*PromotionCreateResponse, error)
	PromotionQuery(ctx context.Context, in *PromotionQueryRequest, opts ...grpc.CallOption) (*PromotionQueryResponse, error)
	PromotionUpdate(ctx context.Context, in *PromotionUpdateRequest, opts ...grpc.CallOption) (*PromotionUpdateResponse, error)
	PromotionDelete(ctx context.Context, in *PromotionDeleteRequest, opts ...grpc.CallOption) (*PromotionDeleteResponse, error)
	QuerySpecificPromotion(ctx context.Context, in *QuerySpecificPromotionRequest, opts ...grpc.CallOption) (*QuerySpecificPromotionResponse, error)
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
}

type promotionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionServiceClient(cc grpc.ClientConnInterface) PromotionServiceClient {
	return &promotionServiceClient{cc}
}

func (c *promotionServiceClient) PromotionCreate(ctx context.Context, in *PromotionCreateRequest, opts ...grpc.CallOption) (*PromotionCreateResponse, error) {
	out := new(PromotionCreateResponse)
	err := c.cc.Invoke(ctx, "/api.PromotionService/PromotionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) PromotionQuery(ctx context.Context, in *PromotionQueryRequest, opts ...grpc.CallOption) (*PromotionQueryResponse, error) {
	out := new(PromotionQueryResponse)
	err := c.cc.Invoke(ctx, "/api.PromotionService/PromotionQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) PromotionUpdate(ctx context.Context, in *PromotionUpdateRequest, opts ...grpc.CallOption) (*PromotionUpdateResponse, error) {
	out := new(PromotionUpdateResponse)
	err := c.cc.Invoke(ctx, "/api.PromotionService/PromotionUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) PromotionDelete(ctx context.Context, in *PromotionDeleteRequest, opts ...grpc.CallOption) (*PromotionDeleteResponse, error) {
	out := new(PromotionDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.PromotionService/PromotionDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) QuerySpecificPromotion(ctx context.Context, in *QuerySpecificPromotionRequest, opts ...grpc.CallOption) (*QuerySpecificPromotionResponse, error) {
	out := new(QuerySpecificPromotionResponse)
	err := c.cc.Invoke(ctx, "/api.PromotionService/QuerySpecificPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/api.PromotionService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServiceServer is the server API for PromotionService service.
// All implementations should embed UnimplementedPromotionServiceServer
// for forward compatibility
type PromotionServiceServer interface {
	PromotionCreate(context.Context, *PromotionCreateRequest) (*PromotionCreateResponse, error)
	PromotionQuery(context.Context, *PromotionQueryRequest) (*PromotionQueryResponse, error)
	PromotionUpdate(context.Context, *PromotionUpdateRequest) (*PromotionUpdateResponse, error)
	PromotionDelete(context.Context, *PromotionDeleteRequest) (*PromotionDeleteResponse, error)
	QuerySpecificPromotion(context.Context, *QuerySpecificPromotionRequest) (*QuerySpecificPromotionResponse, error)
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
}

// UnimplementedPromotionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPromotionServiceServer struct {
}

func (UnimplementedPromotionServiceServer) PromotionCreate(context.Context, *PromotionCreateRequest) (*PromotionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionCreate not implemented")
}
func (UnimplementedPromotionServiceServer) PromotionQuery(context.Context, *PromotionQueryRequest) (*PromotionQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionQuery not implemented")
}
func (UnimplementedPromotionServiceServer) PromotionUpdate(context.Context, *PromotionUpdateRequest) (*PromotionUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionUpdate not implemented")
}
func (UnimplementedPromotionServiceServer) PromotionDelete(context.Context, *PromotionDeleteRequest) (*PromotionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PromotionDelete not implemented")
}
func (UnimplementedPromotionServiceServer) QuerySpecificPromotion(context.Context, *QuerySpecificPromotionRequest) (*QuerySpecificPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySpecificPromotion not implemented")
}
func (UnimplementedPromotionServiceServer) Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}

// UnsafePromotionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromotionServiceServer will
// result in compilation errors.
type UnsafePromotionServiceServer interface {
	mustEmbedUnimplementedPromotionServiceServer()
}

func RegisterPromotionServiceServer(s grpc.ServiceRegistrar, srv PromotionServiceServer) {
	s.RegisterService(&PromotionService_ServiceDesc, srv)
}

func _PromotionService_PromotionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PromotionService/PromotionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionCreate(ctx, req.(*PromotionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_PromotionQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PromotionService/PromotionQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionQuery(ctx, req.(*PromotionQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_PromotionUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PromotionService/PromotionUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionUpdate(ctx, req.(*PromotionUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_PromotionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromotionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).PromotionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PromotionService/PromotionDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).PromotionDelete(ctx, req.(*PromotionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_QuerySpecificPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySpecificPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).QuerySpecificPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PromotionService/QuerySpecificPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).QuerySpecificPromotion(ctx, req.(*QuerySpecificPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PromotionService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromotionService_ServiceDesc is the grpc.ServiceDesc for PromotionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromotionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.PromotionService",
	HandlerType: (*PromotionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PromotionCreate",
			Handler:    _PromotionService_PromotionCreate_Handler,
		},
		{
			MethodName: "PromotionQuery",
			Handler:    _PromotionService_PromotionQuery_Handler,
		},
		{
			MethodName: "PromotionUpdate",
			Handler:    _PromotionService_PromotionUpdate_Handler,
		},
		{
			MethodName: "PromotionDelete",
			Handler:    _PromotionService_PromotionDelete_Handler,
		},
		{
			MethodName: "QuerySpecificPromotion",
			Handler:    _PromotionService_QuerySpecificPromotion_Handler,
		},
		{
			MethodName: "Calculate",
			Handler:    _PromotionService_Calculate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion.proto",
}
