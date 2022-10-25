// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/app/proto/shortener.proto

package pb

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

// ShortenerClient is the client API for Shortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenerClient interface {
	Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShorteningResponse, error)
	DeleteUrls(ctx context.Context, in *DeleteUrlsRequest, opts ...grpc.CallOption) (*Empty, error)
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error)
	ShortenBatch(ctx context.Context, in *ShortenBatchRequest, opts ...grpc.CallOption) (*ShortenBatchResponse, error)
}

type shortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenerClient(cc grpc.ClientConnInterface) ShortenerClient {
	return &shortenerClient{cc}
}

func (c *shortenerClient) Shorten(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShorteningResponse, error) {
	out := new(ShorteningResponse)
	err := c.cc.Invoke(ctx, "/shortener.Shortener/Shorten", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) DeleteUrls(ctx context.Context, in *DeleteUrlsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/shortener.Shortener/DeleteUrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*ExpandResponse, error) {
	out := new(ExpandResponse)
	err := c.cc.Invoke(ctx, "/shortener.Shortener/Expand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenerClient) ShortenBatch(ctx context.Context, in *ShortenBatchRequest, opts ...grpc.CallOption) (*ShortenBatchResponse, error) {
	out := new(ShortenBatchResponse)
	err := c.cc.Invoke(ctx, "/shortener.Shortener/ShortenBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenerServer is the server API for Shortener service.
// All implementations must embed UnimplementedShortenerServer
// for forward compatibility
type ShortenerServer interface {
	Shorten(context.Context, *ShortenRequest) (*ShorteningResponse, error)
	DeleteUrls(context.Context, *DeleteUrlsRequest) (*Empty, error)
	Expand(context.Context, *ExpandRequest) (*ExpandResponse, error)
	ShortenBatch(context.Context, *ShortenBatchRequest) (*ShortenBatchResponse, error)
	mustEmbedUnimplementedShortenerServer()
}

// UnimplementedShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedShortenerServer struct {
}

func (UnimplementedShortenerServer) Shorten(context.Context, *ShortenRequest) (*ShorteningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shorten not implemented")
}
func (UnimplementedShortenerServer) DeleteUrls(context.Context, *DeleteUrlsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUrls not implemented")
}
func (UnimplementedShortenerServer) Expand(context.Context, *ExpandRequest) (*ExpandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedShortenerServer) ShortenBatch(context.Context, *ShortenBatchRequest) (*ShortenBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenBatch not implemented")
}
func (UnimplementedShortenerServer) mustEmbedUnimplementedShortenerServer() {}

// UnsafeShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenerServer will
// result in compilation errors.
type UnsafeShortenerServer interface {
	mustEmbedUnimplementedShortenerServer()
}

func RegisterShortenerServer(s grpc.ServiceRegistrar, srv ShortenerServer) {
	s.RegisterService(&Shortener_ServiceDesc, srv)
}

func _Shortener_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Shortener/Shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Shorten(ctx, req.(*ShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_DeleteUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).DeleteUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Shortener/DeleteUrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).DeleteUrls(ctx, req.(*DeleteUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Shortener/Expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).Expand(ctx, req.(*ExpandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortener_ShortenBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenerServer).ShortenBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shortener.Shortener/ShortenBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenerServer).ShortenBatch(ctx, req.(*ShortenBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shortener_ServiceDesc is the grpc.ServiceDesc for Shortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shortener.Shortener",
	HandlerType: (*ShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shorten",
			Handler:    _Shortener_Shorten_Handler,
		},
		{
			MethodName: "DeleteUrls",
			Handler:    _Shortener_DeleteUrls_Handler,
		},
		{
			MethodName: "Expand",
			Handler:    _Shortener_Expand_Handler,
		},
		{
			MethodName: "ShortenBatch",
			Handler:    _Shortener_ShortenBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/app/proto/shortener.proto",
}
