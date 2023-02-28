// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             (unknown)
// source: shop/interface/v1/shop_interface.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationShopInterfaceGetUserInfo = "/api.sms.interface.v1.ShopInterface/GetUserInfo"
const OperationShopInterfaceLogin = "/api.sms.interface.v1.ShopInterface/Login"
const OperationShopInterfaceRegister = "/api.sms.interface.v1.ShopInterface/Register"

type ShopInterfaceHTTPServer interface {
	// GetUserInfo rpc QuickLogin(LoginReq) returns (LoginResp) {
	//   option (google.api.http) = {
	//     post : "/api/quickLogin"
	//     body : "*"
	//   };
	// }
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
}

func RegisterShopInterfaceHTTPServer(s *http.Server, srv ShopInterfaceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/register", _ShopInterface_Register0_HTTP_Handler(srv))
	r.POST("/api/login", _ShopInterface_Login0_HTTP_Handler(srv))
	r.POST("/api/user/info", _ShopInterface_GetUserInfo0_HTTP_Handler(srv))
}

func _ShopInterface_Register0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterResp)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_Login0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResp)
		return ctx.Result(200, reply)
	}
}

func _ShopInterface_GetUserInfo0_HTTP_Handler(srv ShopInterfaceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShopInterfaceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResp)
		return ctx.Result(200, reply)
	}
}

type ShopInterfaceHTTPClient interface {
	GetUserInfo(ctx context.Context, req *GetUserInfoReq, opts ...http.CallOption) (rsp *GetUserInfoResp, err error)
	Login(ctx context.Context, req *LoginReq, opts ...http.CallOption) (rsp *LoginResp, err error)
	Register(ctx context.Context, req *RegisterReq, opts ...http.CallOption) (rsp *RegisterResp, err error)
}

type ShopInterfaceHTTPClientImpl struct {
	cc *http.Client
}

func NewShopInterfaceHTTPClient(client *http.Client) ShopInterfaceHTTPClient {
	return &ShopInterfaceHTTPClientImpl{client}
}

func (c *ShopInterfaceHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...http.CallOption) (*GetUserInfoResp, error) {
	var out GetUserInfoResp
	pattern := "/api/user/info"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) Login(ctx context.Context, in *LoginReq, opts ...http.CallOption) (*LoginResp, error) {
	var out LoginResp
	pattern := "/api/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShopInterfaceHTTPClientImpl) Register(ctx context.Context, in *RegisterReq, opts ...http.CallOption) (*RegisterResp, error) {
	var out RegisterResp
	pattern := "/api/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShopInterfaceRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
