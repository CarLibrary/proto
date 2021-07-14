// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package carlibrary

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

// CarLibraryServiceClient is the client API for CarLibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarLibraryServiceClient interface {
	//查看所有品牌
	FindALLCarBand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CarLibraryService_FindALLCarBandClient, error)
	//查看某品牌的全部车系
	FindAllCarSeries(ctx context.Context, in *CarSeriesRequest, opts ...grpc.CallOption) (CarLibraryService_FindAllCarSeriesClient, error)
	//查看某品牌的某车系的全部车型
	FindAllCarModel(ctx context.Context, in *CarModelRequest, opts ...grpc.CallOption) (CarLibraryService_FindAllCarModelClient, error)
}

type carLibraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarLibraryServiceClient(cc grpc.ClientConnInterface) CarLibraryServiceClient {
	return &carLibraryServiceClient{cc}
}

func (c *carLibraryServiceClient) FindALLCarBand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CarLibraryService_FindALLCarBandClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarLibraryService_ServiceDesc.Streams[0], "/carlibrary.CarLibraryService/FindALLCarBand", opts...)
	if err != nil {
		return nil, err
	}
	x := &carLibraryServiceFindALLCarBandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarLibraryService_FindALLCarBandClient interface {
	Recv() (*CarBand, error)
	grpc.ClientStream
}

type carLibraryServiceFindALLCarBandClient struct {
	grpc.ClientStream
}

func (x *carLibraryServiceFindALLCarBandClient) Recv() (*CarBand, error) {
	m := new(CarBand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *carLibraryServiceClient) FindAllCarSeries(ctx context.Context, in *CarSeriesRequest, opts ...grpc.CallOption) (CarLibraryService_FindAllCarSeriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarLibraryService_ServiceDesc.Streams[1], "/carlibrary.CarLibraryService/FindAllCarSeries", opts...)
	if err != nil {
		return nil, err
	}
	x := &carLibraryServiceFindAllCarSeriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarLibraryService_FindAllCarSeriesClient interface {
	Recv() (*CarSeries, error)
	grpc.ClientStream
}

type carLibraryServiceFindAllCarSeriesClient struct {
	grpc.ClientStream
}

func (x *carLibraryServiceFindAllCarSeriesClient) Recv() (*CarSeries, error) {
	m := new(CarSeries)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *carLibraryServiceClient) FindAllCarModel(ctx context.Context, in *CarModelRequest, opts ...grpc.CallOption) (CarLibraryService_FindAllCarModelClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarLibraryService_ServiceDesc.Streams[2], "/carlibrary.CarLibraryService/FindAllCarModel", opts...)
	if err != nil {
		return nil, err
	}
	x := &carLibraryServiceFindAllCarModelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarLibraryService_FindAllCarModelClient interface {
	Recv() (*CarModel, error)
	grpc.ClientStream
}

type carLibraryServiceFindAllCarModelClient struct {
	grpc.ClientStream
}

func (x *carLibraryServiceFindAllCarModelClient) Recv() (*CarModel, error) {
	m := new(CarModel)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CarLibraryServiceServer is the server API for CarLibraryService service.
// All implementations must embed UnimplementedCarLibraryServiceServer
// for forward compatibility
type CarLibraryServiceServer interface {
	//查看所有品牌
	FindALLCarBand(*Empty, CarLibraryService_FindALLCarBandServer) error
	//查看某品牌的全部车系
	FindAllCarSeries(*CarSeriesRequest, CarLibraryService_FindAllCarSeriesServer) error
	//查看某品牌的某车系的全部车型
	FindAllCarModel(*CarModelRequest, CarLibraryService_FindAllCarModelServer) error
	mustEmbedUnimplementedCarLibraryServiceServer()
}

// UnimplementedCarLibraryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarLibraryServiceServer struct {
}

func (UnimplementedCarLibraryServiceServer) FindALLCarBand(*Empty, CarLibraryService_FindALLCarBandServer) error {
	return status.Errorf(codes.Unimplemented, "method FindALLCarBand not implemented")
}
func (UnimplementedCarLibraryServiceServer) FindAllCarSeries(*CarSeriesRequest, CarLibraryService_FindAllCarSeriesServer) error {
	return status.Errorf(codes.Unimplemented, "method FindAllCarSeries not implemented")
}
func (UnimplementedCarLibraryServiceServer) FindAllCarModel(*CarModelRequest, CarLibraryService_FindAllCarModelServer) error {
	return status.Errorf(codes.Unimplemented, "method FindAllCarModel not implemented")
}
func (UnimplementedCarLibraryServiceServer) mustEmbedUnimplementedCarLibraryServiceServer() {}

// UnsafeCarLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarLibraryServiceServer will
// result in compilation errors.
type UnsafeCarLibraryServiceServer interface {
	mustEmbedUnimplementedCarLibraryServiceServer()
}

func RegisterCarLibraryServiceServer(s grpc.ServiceRegistrar, srv CarLibraryServiceServer) {
	s.RegisterService(&CarLibraryService_ServiceDesc, srv)
}

func _CarLibraryService_FindALLCarBand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarLibraryServiceServer).FindALLCarBand(m, &carLibraryServiceFindALLCarBandServer{stream})
}

type CarLibraryService_FindALLCarBandServer interface {
	Send(*CarBand) error
	grpc.ServerStream
}

type carLibraryServiceFindALLCarBandServer struct {
	grpc.ServerStream
}

func (x *carLibraryServiceFindALLCarBandServer) Send(m *CarBand) error {
	return x.ServerStream.SendMsg(m)
}

func _CarLibraryService_FindAllCarSeries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CarSeriesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarLibraryServiceServer).FindAllCarSeries(m, &carLibraryServiceFindAllCarSeriesServer{stream})
}

type CarLibraryService_FindAllCarSeriesServer interface {
	Send(*CarSeries) error
	grpc.ServerStream
}

type carLibraryServiceFindAllCarSeriesServer struct {
	grpc.ServerStream
}

func (x *carLibraryServiceFindAllCarSeriesServer) Send(m *CarSeries) error {
	return x.ServerStream.SendMsg(m)
}

func _CarLibraryService_FindAllCarModel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CarModelRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarLibraryServiceServer).FindAllCarModel(m, &carLibraryServiceFindAllCarModelServer{stream})
}

type CarLibraryService_FindAllCarModelServer interface {
	Send(*CarModel) error
	grpc.ServerStream
}

type carLibraryServiceFindAllCarModelServer struct {
	grpc.ServerStream
}

func (x *carLibraryServiceFindAllCarModelServer) Send(m *CarModel) error {
	return x.ServerStream.SendMsg(m)
}

// CarLibraryService_ServiceDesc is the grpc.ServiceDesc for CarLibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarLibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "carlibrary.CarLibraryService",
	HandlerType: (*CarLibraryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindALLCarBand",
			Handler:       _CarLibraryService_FindALLCarBand_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindAllCarSeries",
			Handler:       _CarLibraryService_FindAllCarSeries_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindAllCarModel",
			Handler:       _CarLibraryService_FindAllCarModel_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "carlibrary.proto",
}
