// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/user_sample/user.proto

package v1

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	InsertUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusResponse, error)
	SelectUser(ctx context.Context, in *SelectUserRequest, opts ...grpc.CallOption) (*SelectUserResponse, error)
	SelectListUser(ctx context.Context, in *SelectListUserRequest, opts ...grpc.CallOption) (*SelectListUserResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) InsertUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error) {
	out := new(InsertUserResponse)
	err := c.cc.Invoke(ctx, "/lines.grpc.user.v1.User/InsertUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/lines.grpc.user.v1.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUserStatus(ctx context.Context, in *UpdateUserStatusRequest, opts ...grpc.CallOption) (*UpdateUserStatusResponse, error) {
	out := new(UpdateUserStatusResponse)
	err := c.cc.Invoke(ctx, "/lines.grpc.user.v1.User/UpdateUserStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SelectUser(ctx context.Context, in *SelectUserRequest, opts ...grpc.CallOption) (*SelectUserResponse, error) {
	out := new(SelectUserResponse)
	err := c.cc.Invoke(ctx, "/lines.grpc.user.v1.User/SelectUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SelectListUser(ctx context.Context, in *SelectListUserRequest, opts ...grpc.CallOption) (*SelectListUserResponse, error) {
	out := new(SelectListUserResponse)
	err := c.cc.Invoke(ctx, "/lines.grpc.user.v1.User/SelectListUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	InsertUser(context.Context, *InsertUserRequest) (*InsertUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error)
	SelectUser(context.Context, *SelectUserRequest) (*SelectUserResponse, error)
	SelectListUser(context.Context, *SelectListUserRequest) (*SelectListUserResponse, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) InsertUser(context.Context, *InsertUserRequest) (*InsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertUser not implemented")
}
func (UnimplementedUserServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServer) UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*UpdateUserStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStatus not implemented")
}
func (UnimplementedUserServer) SelectUser(context.Context, *SelectUserRequest) (*SelectUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectUser not implemented")
}
func (UnimplementedUserServer) SelectListUser(context.Context, *SelectListUserRequest) (*SelectListUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectListUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_InsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).InsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.grpc.user.v1.User/InsertUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).InsertUser(ctx, req.(*InsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.grpc.user.v1.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.grpc.user.v1.User/UpdateUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUserStatus(ctx, req.(*UpdateUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SelectUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SelectUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.grpc.user.v1.User/SelectUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SelectUser(ctx, req.(*SelectUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SelectListUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SelectListUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lines.grpc.user.v1.User/SelectListUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SelectListUser(ctx, req.(*SelectListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lines.grpc.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertUser",
			Handler:    _User_InsertUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
		{
			MethodName: "UpdateUserStatus",
			Handler:    _User_UpdateUserStatus_Handler,
		},
		{
			MethodName: "SelectUser",
			Handler:    _User_SelectUser_Handler,
		},
		{
			MethodName: "SelectListUser",
			Handler:    _User_SelectListUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user_sample/user.proto",
}
