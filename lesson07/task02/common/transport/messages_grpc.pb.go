// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transport

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

// UserRegisterClient is the client API for UserRegister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRegisterClient interface {
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	UsersList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UsersListResponse, error)
}

type userRegisterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegisterClient(cc grpc.ClientConnInterface) UserRegisterClient {
	return &userRegisterClient{cc}
}

func (c *userRegisterClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, "/UserRegister/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegisterClient) UsersList(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*UsersListResponse, error) {
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, "/UserRegister/UsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegisterServer is the server API for UserRegister service.
// All implementations must embed UnimplementedUserRegisterServer
// for forward compatibility
type UserRegisterServer interface {
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	UsersList(context.Context, *EmptyRequest) (*UsersListResponse, error)
	mustEmbedUnimplementedUserRegisterServer()
}

// UnimplementedUserRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedUserRegisterServer struct {
}

func (UnimplementedUserRegisterServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserRegisterServer) UsersList(context.Context, *EmptyRequest) (*UsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsersList not implemented")
}
func (UnimplementedUserRegisterServer) mustEmbedUnimplementedUserRegisterServer() {}

// UnsafeUserRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRegisterServer will
// result in compilation errors.
type UnsafeUserRegisterServer interface {
	mustEmbedUnimplementedUserRegisterServer()
}

func RegisterUserRegisterServer(s grpc.ServiceRegistrar, srv UserRegisterServer) {
	s.RegisterService(&UserRegister_ServiceDesc, srv)
}

func _UserRegister_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserRegister/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegister_UsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegisterServer).UsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserRegister/UsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegisterServer).UsersList(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRegister_ServiceDesc is the grpc.ServiceDesc for UserRegister service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRegister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserRegister",
	HandlerType: (*UserRegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _UserRegister_UserRegister_Handler,
		},
		{
			MethodName: "UsersList",
			Handler:    _UserRegister_UsersList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "messages.proto",
}