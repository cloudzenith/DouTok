// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: svapi/collection.proto

package svapi

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCollectionServiceAddVideo2Collection = "/svapi.CollectionService/AddVideo2Collection"
const OperationCollectionServiceCreateCollection = "/svapi.CollectionService/CreateCollection"
const OperationCollectionServiceListCollection = "/svapi.CollectionService/ListCollection"
const OperationCollectionServiceListVideo4Collection = "/svapi.CollectionService/ListVideo4Collection"
const OperationCollectionServiceRemoveCollection = "/svapi.CollectionService/RemoveCollection"
const OperationCollectionServiceRemoveVideoFromCollection = "/svapi.CollectionService/RemoveVideoFromCollection"
const OperationCollectionServiceUpdateCollection = "/svapi.CollectionService/UpdateCollection"

type CollectionServiceHTTPServer interface {
	AddVideo2Collection(context.Context, *AddVideo2CollectionRequest) (*AddVideo2CollectionResponse, error)
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	ListCollection(context.Context, *ListCollectionRequest) (*ListCollectionResponse, error)
	ListVideo4Collection(context.Context, *ListVideo4CollectionRequest) (*ListVideo4CollectionResponse, error)
	RemoveCollection(context.Context, *RemoveCollectionRequest) (*RemoveCollectionResponse, error)
	RemoveVideoFromCollection(context.Context, *RemoveVideoFromCollectionRequest) (*RemoveVideoFromCollectionResponse, error)
	UpdateCollection(context.Context, *UpdateCollectionRequest) (*UpdateCollectionResponse, error)
}

func RegisterCollectionServiceHTTPServer(s *http.Server, srv CollectionServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/collection", _CollectionService_CreateCollection0_HTTP_Handler(srv))
	r.DELETE("/collection", _CollectionService_RemoveCollection0_HTTP_Handler(srv))
	r.GET("/collection", _CollectionService_ListCollection0_HTTP_Handler(srv))
	r.PUT("/collection", _CollectionService_UpdateCollection0_HTTP_Handler(srv))
	r.POST("/collection/video", _CollectionService_AddVideo2Collection0_HTTP_Handler(srv))
	r.DELETE("/collection/video", _CollectionService_RemoveVideoFromCollection0_HTTP_Handler(srv))
	r.GET("/collection/video", _CollectionService_ListVideo4Collection0_HTTP_Handler(srv))
}

func _CollectionService_CreateCollection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCollectionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceCreateCollection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCollection(ctx, req.(*CreateCollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _CollectionService_RemoveCollection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveCollectionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceRemoveCollection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveCollection(ctx, req.(*RemoveCollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _CollectionService_ListCollection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListCollectionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceListCollection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListCollection(ctx, req.(*ListCollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _CollectionService_UpdateCollection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateCollectionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceUpdateCollection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateCollection(ctx, req.(*UpdateCollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _CollectionService_AddVideo2Collection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddVideo2CollectionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceAddVideo2Collection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddVideo2Collection(ctx, req.(*AddVideo2CollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _CollectionService_RemoveVideoFromCollection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RemoveVideoFromCollectionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceRemoveVideoFromCollection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RemoveVideoFromCollection(ctx, req.(*RemoveVideoFromCollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

func _CollectionService_ListVideo4Collection0_HTTP_Handler(srv CollectionServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListVideo4CollectionRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCollectionServiceListVideo4Collection)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListVideo4Collection(ctx, req.(*ListVideo4CollectionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		return ctx.Result(200, out)
	}
}

type CollectionServiceHTTPClient interface {
	AddVideo2Collection(ctx context.Context, req *AddVideo2CollectionRequest, opts ...http.CallOption) (rsp *AddVideo2CollectionResponse, err error)
	CreateCollection(ctx context.Context, req *CreateCollectionRequest, opts ...http.CallOption) (rsp *CreateCollectionResponse, err error)
	ListCollection(ctx context.Context, req *ListCollectionRequest, opts ...http.CallOption) (rsp *ListCollectionResponse, err error)
	ListVideo4Collection(ctx context.Context, req *ListVideo4CollectionRequest, opts ...http.CallOption) (rsp *ListVideo4CollectionResponse, err error)
	RemoveCollection(ctx context.Context, req *RemoveCollectionRequest, opts ...http.CallOption) (rsp *RemoveCollectionResponse, err error)
	RemoveVideoFromCollection(ctx context.Context, req *RemoveVideoFromCollectionRequest, opts ...http.CallOption) (rsp *RemoveVideoFromCollectionResponse, err error)
	UpdateCollection(ctx context.Context, req *UpdateCollectionRequest, opts ...http.CallOption) (rsp *UpdateCollectionResponse, err error)
}

type CollectionServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewCollectionServiceHTTPClient(client *http.Client) CollectionServiceHTTPClient {
	return &CollectionServiceHTTPClientImpl{client}
}

func (c *CollectionServiceHTTPClientImpl) AddVideo2Collection(ctx context.Context, in *AddVideo2CollectionRequest, opts ...http.CallOption) (*AddVideo2CollectionResponse, error) {
	var out AddVideo2CollectionResponse
	pattern := "/collection/video"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionServiceAddVideo2Collection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CollectionServiceHTTPClientImpl) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...http.CallOption) (*CreateCollectionResponse, error) {
	var out CreateCollectionResponse
	pattern := "/collection"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionServiceCreateCollection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CollectionServiceHTTPClientImpl) ListCollection(ctx context.Context, in *ListCollectionRequest, opts ...http.CallOption) (*ListCollectionResponse, error) {
	var out ListCollectionResponse
	pattern := "/collection"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCollectionServiceListCollection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CollectionServiceHTTPClientImpl) ListVideo4Collection(ctx context.Context, in *ListVideo4CollectionRequest, opts ...http.CallOption) (*ListVideo4CollectionResponse, error) {
	var out ListVideo4CollectionResponse
	pattern := "/collection/video"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCollectionServiceListVideo4Collection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CollectionServiceHTTPClientImpl) RemoveCollection(ctx context.Context, in *RemoveCollectionRequest, opts ...http.CallOption) (*RemoveCollectionResponse, error) {
	var out RemoveCollectionResponse
	pattern := "/collection"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionServiceRemoveCollection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CollectionServiceHTTPClientImpl) RemoveVideoFromCollection(ctx context.Context, in *RemoveVideoFromCollectionRequest, opts ...http.CallOption) (*RemoveVideoFromCollectionResponse, error) {
	var out RemoveVideoFromCollectionResponse
	pattern := "/collection/video"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionServiceRemoveVideoFromCollection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *CollectionServiceHTTPClientImpl) UpdateCollection(ctx context.Context, in *UpdateCollectionRequest, opts ...http.CallOption) (*UpdateCollectionResponse, error) {
	var out UpdateCollectionResponse
	pattern := "/collection"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCollectionServiceUpdateCollection))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}