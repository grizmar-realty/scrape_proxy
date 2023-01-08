// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: scrape.proto

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

// ScrapeProxyClient is the client API for ScrapeProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScrapeProxyClient interface {
	SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type scrapeProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewScrapeProxyClient(cc grpc.ClientConnInterface) ScrapeProxyClient {
	return &scrapeProxyClient{cc}
}

func (c *scrapeProxyClient) SendRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.ScrapeProxy/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScrapeProxyServer is the server API for ScrapeProxy service.
// All implementations must embed UnimplementedScrapeProxyServer
// for forward compatibility
type ScrapeProxyServer interface {
	SendRequest(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedScrapeProxyServer()
}

// UnimplementedScrapeProxyServer must be embedded to have forward compatible implementations.
type UnimplementedScrapeProxyServer struct {
}

func (UnimplementedScrapeProxyServer) SendRequest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedScrapeProxyServer) mustEmbedUnimplementedScrapeProxyServer() {}

// UnsafeScrapeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScrapeProxyServer will
// result in compilation errors.
type UnsafeScrapeProxyServer interface {
	mustEmbedUnimplementedScrapeProxyServer()
}

func RegisterScrapeProxyServer(s grpc.ServiceRegistrar, srv ScrapeProxyServer) {
	s.RegisterService(&ScrapeProxy_ServiceDesc, srv)
}

func _ScrapeProxy_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScrapeProxyServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ScrapeProxy/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScrapeProxyServer).SendRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ScrapeProxy_ServiceDesc is the grpc.ServiceDesc for ScrapeProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScrapeProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ScrapeProxy",
	HandlerType: (*ScrapeProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _ScrapeProxy_SendRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrape.proto",
}
