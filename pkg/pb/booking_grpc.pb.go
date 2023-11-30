// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: booking.proto

package __

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	RegisterSelectSeat(ctx context.Context, in *SeatRequest, opts ...grpc.CallOption) (*SeatResponse, error)
	AddConfirmedSeats(ctx context.Context, in *ConfirmedSeatRequest, opts ...grpc.CallOption) (*ConfirmedSeatResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) RegisterSelectSeat(ctx context.Context, in *SeatRequest, opts ...grpc.CallOption) (*SeatResponse, error) {
	out := new(SeatResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/RegisterSelectSeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AddConfirmedSeats(ctx context.Context, in *ConfirmedSeatRequest, opts ...grpc.CallOption) (*ConfirmedSeatResponse, error) {
	out := new(ConfirmedSeatResponse)
	err := c.cc.Invoke(ctx, "/admin.AdminService/AddConfirmedSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	RegisterSelectSeat(context.Context, *SeatRequest) (*SeatResponse, error)
	AddConfirmedSeats(context.Context, *ConfirmedSeatRequest) (*ConfirmedSeatResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) RegisterSelectSeat(context.Context, *SeatRequest) (*SeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSelectSeat not implemented")
}
func (UnimplementedAdminServiceServer) AddConfirmedSeats(context.Context, *ConfirmedSeatRequest) (*ConfirmedSeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfirmedSeats not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_RegisterSelectSeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RegisterSelectSeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/RegisterSelectSeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RegisterSelectSeat(ctx, req.(*SeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AddConfirmedSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmedSeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddConfirmedSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.AdminService/AddConfirmedSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddConfirmedSeats(ctx, req.(*ConfirmedSeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterSelectSeat",
			Handler:    _AdminService_RegisterSelectSeat_Handler,
		},
		{
			MethodName: "AddConfirmedSeats",
			Handler:    _AdminService_AddConfirmedSeats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}
