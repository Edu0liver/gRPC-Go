// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: temperature.proto

package pb

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
	TemperatureService_RecordTemperatura_FullMethodName = "/pb.TemperatureService/RecordTemperatura"
)

// TemperatureServiceClient is the client API for TemperatureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemperatureServiceClient interface {
	RecordTemperatura(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[TemperatureRequest, TemperatureResponse], error)
}

type temperatureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTemperatureServiceClient(cc grpc.ClientConnInterface) TemperatureServiceClient {
	return &temperatureServiceClient{cc}
}

func (c *temperatureServiceClient) RecordTemperatura(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[TemperatureRequest, TemperatureResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TemperatureService_ServiceDesc.Streams[0], TemperatureService_RecordTemperatura_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[TemperatureRequest, TemperatureResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TemperatureService_RecordTemperaturaClient = grpc.ClientStreamingClient[TemperatureRequest, TemperatureResponse]

// TemperatureServiceServer is the server API for TemperatureService service.
// All implementations must embed UnimplementedTemperatureServiceServer
// for forward compatibility.
type TemperatureServiceServer interface {
	RecordTemperatura(grpc.ClientStreamingServer[TemperatureRequest, TemperatureResponse]) error
	mustEmbedUnimplementedTemperatureServiceServer()
}

// UnimplementedTemperatureServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTemperatureServiceServer struct{}

func (UnimplementedTemperatureServiceServer) RecordTemperatura(grpc.ClientStreamingServer[TemperatureRequest, TemperatureResponse]) error {
	return status.Errorf(codes.Unimplemented, "method RecordTemperatura not implemented")
}
func (UnimplementedTemperatureServiceServer) mustEmbedUnimplementedTemperatureServiceServer() {}
func (UnimplementedTemperatureServiceServer) testEmbeddedByValue()                            {}

// UnsafeTemperatureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemperatureServiceServer will
// result in compilation errors.
type UnsafeTemperatureServiceServer interface {
	mustEmbedUnimplementedTemperatureServiceServer()
}

func RegisterTemperatureServiceServer(s grpc.ServiceRegistrar, srv TemperatureServiceServer) {
	// If the following call pancis, it indicates UnimplementedTemperatureServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TemperatureService_ServiceDesc, srv)
}

func _TemperatureService_RecordTemperatura_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TemperatureServiceServer).RecordTemperatura(&grpc.GenericServerStream[TemperatureRequest, TemperatureResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TemperatureService_RecordTemperaturaServer = grpc.ClientStreamingServer[TemperatureRequest, TemperatureResponse]

// TemperatureService_ServiceDesc is the grpc.ServiceDesc for TemperatureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TemperatureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TemperatureService",
	HandlerType: (*TemperatureServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordTemperatura",
			Handler:       _TemperatureService_RecordTemperatura_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "temperature.proto",
}
