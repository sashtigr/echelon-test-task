// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: api/proto/ytpreview.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	YouTubeService_GetThumbnail_FullMethodName = "/youtube.YouTubeService/GetThumbnail"
)

// YouTubeServiceClient is the client API for YouTubeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YouTubeServiceClient interface {
	GetThumbnail(ctx context.Context, in *ThumbnailRequest, opts ...grpc.CallOption) (*ThumbnailResponse, error)
}

type youTubeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYouTubeServiceClient(cc grpc.ClientConnInterface) YouTubeServiceClient {
	return &youTubeServiceClient{cc}
}

func (c *youTubeServiceClient) GetThumbnail(ctx context.Context, in *ThumbnailRequest, opts ...grpc.CallOption) (*ThumbnailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ThumbnailResponse)
	err := c.cc.Invoke(ctx, YouTubeService_GetThumbnail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YouTubeServiceServer is the server API for YouTubeService service.
// All implementations must embed UnimplementedYouTubeServiceServer
// for forward compatibility.
type YouTubeServiceServer interface {
	GetThumbnail(context.Context, *ThumbnailRequest) (*ThumbnailResponse, error)
	mustEmbedUnimplementedYouTubeServiceServer()
}

// UnimplementedYouTubeServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedYouTubeServiceServer struct{}

func (UnimplementedYouTubeServiceServer) GetThumbnail(context.Context, *ThumbnailRequest) (*ThumbnailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThumbnail not implemented")
}
func (UnimplementedYouTubeServiceServer) mustEmbedUnimplementedYouTubeServiceServer() {}
func (UnimplementedYouTubeServiceServer) testEmbeddedByValue()                        {}

// UnsafeYouTubeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YouTubeServiceServer will
// result in compilation errors.
type UnsafeYouTubeServiceServer interface {
	mustEmbedUnimplementedYouTubeServiceServer()
}

func RegisterYouTubeServiceServer(s grpc.ServiceRegistrar, srv YouTubeServiceServer) {
	// If the following call pancis, it indicates UnimplementedYouTubeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&YouTubeService_ServiceDesc, srv)
}

func _YouTubeService_GetThumbnail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbnailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YouTubeServiceServer).GetThumbnail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YouTubeService_GetThumbnail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YouTubeServiceServer).GetThumbnail(ctx, req.(*ThumbnailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YouTubeService_ServiceDesc is the grpc.ServiceDesc for YouTubeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YouTubeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "youtube.YouTubeService",
	HandlerType: (*YouTubeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThumbnail",
			Handler:    _YouTubeService_GetThumbnail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ytpreview.proto",
}
