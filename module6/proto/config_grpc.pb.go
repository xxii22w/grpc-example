// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0--rc2
// source: proto/config.proto

package proto

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
	ConfigService_LongRunning_FullMethodName      = "/config.ConfigService/LongRunning"
	ConfigService_Flaky_FullMethodName            = "/config.ConfigService/Flaky"
	ConfigService_GetServerAddress_FullMethodName = "/config.ConfigService/GetServerAddress"
)

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	LongRunning(ctx context.Context, in *LongRunningRequest, opts ...grpc.CallOption) (*LongRunningResponse, error)
	Flaky(ctx context.Context, in *FlakyRequest, opts ...grpc.CallOption) (*FlakyResponse, error)
	GetServerAddress(ctx context.Context, in *GetServerAddressRequest, opts ...grpc.CallOption) (*GetServerAddressResponse, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) LongRunning(ctx context.Context, in *LongRunningRequest, opts ...grpc.CallOption) (*LongRunningResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LongRunningResponse)
	err := c.cc.Invoke(ctx, ConfigService_LongRunning_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) Flaky(ctx context.Context, in *FlakyRequest, opts ...grpc.CallOption) (*FlakyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlakyResponse)
	err := c.cc.Invoke(ctx, ConfigService_Flaky_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetServerAddress(ctx context.Context, in *GetServerAddressRequest, opts ...grpc.CallOption) (*GetServerAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServerAddressResponse)
	err := c.cc.Invoke(ctx, ConfigService_GetServerAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility.
type ConfigServiceServer interface {
	LongRunning(context.Context, *LongRunningRequest) (*LongRunningResponse, error)
	Flaky(context.Context, *FlakyRequest) (*FlakyResponse, error)
	GetServerAddress(context.Context, *GetServerAddressRequest) (*GetServerAddressResponse, error)
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConfigServiceServer struct{}

func (UnimplementedConfigServiceServer) LongRunning(context.Context, *LongRunningRequest) (*LongRunningResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LongRunning not implemented")
}
func (UnimplementedConfigServiceServer) Flaky(context.Context, *FlakyRequest) (*FlakyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flaky not implemented")
}
func (UnimplementedConfigServiceServer) GetServerAddress(context.Context, *GetServerAddressRequest) (*GetServerAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerAddress not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}
func (UnimplementedConfigServiceServer) testEmbeddedByValue()                       {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	// If the following call pancis, it indicates UnimplementedConfigServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_LongRunning_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongRunningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).LongRunning(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_LongRunning_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).LongRunning(ctx, req.(*LongRunningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_Flaky_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlakyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Flaky(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_Flaky_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Flaky(ctx, req.(*FlakyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetServerAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetServerAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_GetServerAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetServerAddress(ctx, req.(*GetServerAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LongRunning",
			Handler:    _ConfigService_LongRunning_Handler,
		},
		{
			MethodName: "Flaky",
			Handler:    _ConfigService_Flaky_Handler,
		},
		{
			MethodName: "GetServerAddress",
			Handler:    _ConfigService_GetServerAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/config.proto",
}