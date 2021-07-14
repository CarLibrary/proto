// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package score

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

// ScoreServiceClient is the client API for ScoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoreServiceClient interface {
	MakeScore(ctx context.Context, in *Score, opts ...grpc.CallOption) (*Score, error)
}

type scoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScoreServiceClient(cc grpc.ClientConnInterface) ScoreServiceClient {
	return &scoreServiceClient{cc}
}

func (c *scoreServiceClient) MakeScore(ctx context.Context, in *Score, opts ...grpc.CallOption) (*Score, error) {
	out := new(Score)
	err := c.cc.Invoke(ctx, "/score.ScoreService/MakeScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoreServiceServer is the server API for ScoreService service.
// All implementations must embed UnimplementedScoreServiceServer
// for forward compatibility
type ScoreServiceServer interface {
	MakeScore(context.Context, *Score) (*Score, error)
	mustEmbedUnimplementedScoreServiceServer()
}

// UnimplementedScoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScoreServiceServer struct {
}

func (UnimplementedScoreServiceServer) MakeScore(context.Context, *Score) (*Score, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeScore not implemented")
}
func (UnimplementedScoreServiceServer) mustEmbedUnimplementedScoreServiceServer() {}

// UnsafeScoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoreServiceServer will
// result in compilation errors.
type UnsafeScoreServiceServer interface {
	mustEmbedUnimplementedScoreServiceServer()
}

func RegisterScoreServiceServer(s grpc.ServiceRegistrar, srv ScoreServiceServer) {
	s.RegisterService(&ScoreService_ServiceDesc, srv)
}

func _ScoreService_MakeScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Score)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoreServiceServer).MakeScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/score.ScoreService/MakeScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoreServiceServer).MakeScore(ctx, req.(*Score))
	}
	return interceptor(ctx, in, info, handler)
}

// ScoreService_ServiceDesc is the grpc.ServiceDesc for ScoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "score.ScoreService",
	HandlerType: (*ScoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeScore",
			Handler:    _ScoreService_MakeScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "score.proto",
}
