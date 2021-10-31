// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UmsClient is the client API for Ums service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UmsClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
	CreateAccountUser(ctx context.Context, in *CreateAccountUserReq, opts ...grpc.CallOption) (*CreateAccountUserResp, error)
	UpdateAccountUser(ctx context.Context, in *UpdateAccountUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAccountUser(ctx context.Context, in *DeleteAccountUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAccountUser(ctx context.Context, in *GetAccountUserReq, opts ...grpc.CallOption) (*GetAccountUserResp, error)
	ListAccountUser(ctx context.Context, in *ListAccountUserReq, opts ...grpc.CallOption) (*ListAccountUserResp, error)
}

type umsClient struct {
	cc grpc.ClientConnInterface
}

func NewUmsClient(cc grpc.ClientConnInterface) UmsClient {
	return &umsClient{cc}
}

func (c *umsClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	out := new(LogoutResp)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) CreateAccountUser(ctx context.Context, in *CreateAccountUserReq, opts ...grpc.CallOption) (*CreateAccountUserResp, error) {
	out := new(CreateAccountUserResp)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/CreateAccountUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) UpdateAccountUser(ctx context.Context, in *UpdateAccountUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/UpdateAccountUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) DeleteAccountUser(ctx context.Context, in *DeleteAccountUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/DeleteAccountUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) GetAccountUser(ctx context.Context, in *GetAccountUserReq, opts ...grpc.CallOption) (*GetAccountUserResp, error) {
	out := new(GetAccountUserResp)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/GetAccountUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *umsClient) ListAccountUser(ctx context.Context, in *ListAccountUserReq, opts ...grpc.CallOption) (*ListAccountUserResp, error) {
	out := new(ListAccountUserResp)
	err := c.cc.Invoke(ctx, "/api.ums.service.v1.Ums/ListAccountUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UmsServer is the server API for Ums service.
// All implementations must embed UnimplementedUmsServer
// for forward compatibility
type UmsServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Logout(context.Context, *LogoutReq) (*LogoutResp, error)
	CreateAccountUser(context.Context, *CreateAccountUserReq) (*CreateAccountUserResp, error)
	UpdateAccountUser(context.Context, *UpdateAccountUserReq) (*emptypb.Empty, error)
	DeleteAccountUser(context.Context, *DeleteAccountUserReq) (*emptypb.Empty, error)
	GetAccountUser(context.Context, *GetAccountUserReq) (*GetAccountUserResp, error)
	ListAccountUser(context.Context, *ListAccountUserReq) (*ListAccountUserResp, error)
	mustEmbedUnimplementedUmsServer()
}

// UnimplementedUmsServer must be embedded to have forward compatible implementations.
type UnimplementedUmsServer struct {
}

func (UnimplementedUmsServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUmsServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUmsServer) Logout(context.Context, *LogoutReq) (*LogoutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUmsServer) CreateAccountUser(context.Context, *CreateAccountUserReq) (*CreateAccountUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccountUser not implemented")
}
func (UnimplementedUmsServer) UpdateAccountUser(context.Context, *UpdateAccountUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountUser not implemented")
}
func (UnimplementedUmsServer) DeleteAccountUser(context.Context, *DeleteAccountUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccountUser not implemented")
}
func (UnimplementedUmsServer) GetAccountUser(context.Context, *GetAccountUserReq) (*GetAccountUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountUser not implemented")
}
func (UnimplementedUmsServer) ListAccountUser(context.Context, *ListAccountUserReq) (*ListAccountUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccountUser not implemented")
}
func (UnimplementedUmsServer) mustEmbedUnimplementedUmsServer() {}

// UnsafeUmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UmsServer will
// result in compilation errors.
type UnsafeUmsServer interface {
	mustEmbedUnimplementedUmsServer()
}

func RegisterUmsServer(s grpc.ServiceRegistrar, srv UmsServer) {
	s.RegisterService(&Ums_ServiceDesc, srv)
}

func _Ums_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_CreateAccountUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).CreateAccountUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/CreateAccountUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).CreateAccountUser(ctx, req.(*CreateAccountUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_UpdateAccountUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).UpdateAccountUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/UpdateAccountUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).UpdateAccountUser(ctx, req.(*UpdateAccountUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_DeleteAccountUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).DeleteAccountUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/DeleteAccountUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).DeleteAccountUser(ctx, req.(*DeleteAccountUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_GetAccountUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).GetAccountUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/GetAccountUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).GetAccountUser(ctx, req.(*GetAccountUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ums_ListAccountUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAccountUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UmsServer).ListAccountUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ums.service.v1.Ums/ListAccountUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UmsServer).ListAccountUser(ctx, req.(*ListAccountUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ums_ServiceDesc is the grpc.ServiceDesc for Ums service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ums_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.ums.service.v1.Ums",
	HandlerType: (*UmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Ums_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Ums_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Ums_Logout_Handler,
		},
		{
			MethodName: "CreateAccountUser",
			Handler:    _Ums_CreateAccountUser_Handler,
		},
		{
			MethodName: "UpdateAccountUser",
			Handler:    _Ums_UpdateAccountUser_Handler,
		},
		{
			MethodName: "DeleteAccountUser",
			Handler:    _Ums_DeleteAccountUser_Handler,
		},
		{
			MethodName: "GetAccountUser",
			Handler:    _Ums_GetAccountUser_Handler,
		},
		{
			MethodName: "ListAccountUser",
			Handler:    _Ums_ListAccountUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/ums.proto",
}
