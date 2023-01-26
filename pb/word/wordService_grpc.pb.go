// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: wordService.proto

package word

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

// WordServiceClient is the client API for WordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WordServiceClient interface {
	CreateWord(ctx context.Context, in *AddWordReq, opts ...grpc.CallOption) (*WordRes, error)
	GetRandomWord(ctx context.Context, in *GetRandomWordReq, opts ...grpc.CallOption) (*WordRes, error)
	SearchWord(ctx context.Context, in *SearchWordReq, opts ...grpc.CallOption) (*SearchWordRes, error)
}

type wordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWordServiceClient(cc grpc.ClientConnInterface) WordServiceClient {
	return &wordServiceClient{cc}
}

func (c *wordServiceClient) CreateWord(ctx context.Context, in *AddWordReq, opts ...grpc.CallOption) (*WordRes, error) {
	out := new(WordRes)
	err := c.cc.Invoke(ctx, "/word.WordService/CreateWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordServiceClient) GetRandomWord(ctx context.Context, in *GetRandomWordReq, opts ...grpc.CallOption) (*WordRes, error) {
	out := new(WordRes)
	err := c.cc.Invoke(ctx, "/word.WordService/GetRandomWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wordServiceClient) SearchWord(ctx context.Context, in *SearchWordReq, opts ...grpc.CallOption) (*SearchWordRes, error) {
	out := new(SearchWordRes)
	err := c.cc.Invoke(ctx, "/word.WordService/SearchWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordServiceServer is the server API for WordService service.
// All implementations must embed UnimplementedWordServiceServer
// for forward compatibility
type WordServiceServer interface {
	CreateWord(context.Context, *AddWordReq) (*WordRes, error)
	GetRandomWord(context.Context, *GetRandomWordReq) (*WordRes, error)
	SearchWord(context.Context, *SearchWordReq) (*SearchWordRes, error)
	mustEmbedUnimplementedWordServiceServer()
}

// UnimplementedWordServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWordServiceServer struct {
}

func (UnimplementedWordServiceServer) CreateWord(context.Context, *AddWordReq) (*WordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWord not implemented")
}
func (UnimplementedWordServiceServer) GetRandomWord(context.Context, *GetRandomWordReq) (*WordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandomWord not implemented")
}
func (UnimplementedWordServiceServer) SearchWord(context.Context, *SearchWordReq) (*SearchWordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchWord not implemented")
}
func (UnimplementedWordServiceServer) mustEmbedUnimplementedWordServiceServer() {}

// UnsafeWordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WordServiceServer will
// result in compilation errors.
type UnsafeWordServiceServer interface {
	mustEmbedUnimplementedWordServiceServer()
}

func RegisterWordServiceServer(s grpc.ServiceRegistrar, srv WordServiceServer) {
	s.RegisterService(&WordService_ServiceDesc, srv)
}

func _WordService_CreateWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).CreateWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/CreateWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).CreateWord(ctx, req.(*AddWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordService_GetRandomWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).GetRandomWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/GetRandomWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).GetRandomWord(ctx, req.(*GetRandomWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WordService_SearchWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordServiceServer).SearchWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/word.WordService/SearchWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordServiceServer).SearchWord(ctx, req.(*SearchWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// WordService_ServiceDesc is the grpc.ServiceDesc for WordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "word.WordService",
	HandlerType: (*WordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWord",
			Handler:    _WordService_CreateWord_Handler,
		},
		{
			MethodName: "GetRandomWord",
			Handler:    _WordService_GetRandomWord_Handler,
		},
		{
			MethodName: "SearchWord",
			Handler:    _WordService_SearchWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordService.proto",
}