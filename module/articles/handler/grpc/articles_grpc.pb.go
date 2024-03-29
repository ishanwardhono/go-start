// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: module/articles/handler/grpc/articles.proto

package grpc

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

// ArticlesGrpcClient is the client API for ArticlesGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesGrpcClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type articlesGrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesGrpcClient(cc grpc.ClientConnInterface) ArticlesGrpcClient {
	return &articlesGrpcClient{cc}
}

func (c *articlesGrpcClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/ArticlesGrpc/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesGrpcServer is the server API for ArticlesGrpc service.
// All implementations must embed UnimplementedArticlesGrpcServer
// for forward compatibility
type ArticlesGrpcServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	mustEmbedUnimplementedArticlesGrpcServer()
}

// UnimplementedArticlesGrpcServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesGrpcServer struct {
}

func (UnimplementedArticlesGrpcServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedArticlesGrpcServer) mustEmbedUnimplementedArticlesGrpcServer() {}

// UnsafeArticlesGrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesGrpcServer will
// result in compilation errors.
type UnsafeArticlesGrpcServer interface {
	mustEmbedUnimplementedArticlesGrpcServer()
}

func RegisterArticlesGrpcServer(s grpc.ServiceRegistrar, srv ArticlesGrpcServer) {
	s.RegisterService(&ArticlesGrpc_ServiceDesc, srv)
}

func _ArticlesGrpc_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesGrpcServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ArticlesGrpc/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesGrpcServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArticlesGrpc_ServiceDesc is the grpc.ServiceDesc for ArticlesGrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArticlesGrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ArticlesGrpc",
	HandlerType: (*ArticlesGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _ArticlesGrpc_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module/articles/handler/grpc/articles.proto",
}
