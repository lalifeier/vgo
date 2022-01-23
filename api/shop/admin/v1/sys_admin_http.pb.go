// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type SysHTTPServer interface {
	CreateDict(context.Context, *CreateDictReq) (*CreateDictResp, error)
	DeleteDict(context.Context, *DeleteDictReq) (*emptypb.Empty, error)
	GetDict(context.Context, *GetDictReq) (*GetDictResp, error)
	ListDict(context.Context, *ListDictReq) (*ListDictResp, error)
	PageListDict(context.Context, *PageListDictReq) (*PageListDictResp, error)
	UpdateDict(context.Context, *UpdateDictReq) (*emptypb.Empty, error)
}

func RegisterSysHTTPServer(s *http.Server, srv SysHTTPServer) {
	r := s.Route("/")
	r.POST("/api/dict", _Sys_CreateDict0_HTTP_Handler(srv))
	r.PUT("/api/dict/{id}", _Sys_UpdateDict0_HTTP_Handler(srv))
	r.DELETE("/api/dict/{id}", _Sys_DeleteDict0_HTTP_Handler(srv))
	r.GET("/api/dicts", _Sys_ListDict0_HTTP_Handler(srv))
	r.GET("/api/dict", _Sys_PageListDict0_HTTP_Handler(srv))
	r.GET("/api/dict/{id}", _Sys_GetDict0_HTTP_Handler(srv))
}

func _Sys_CreateDict0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDictReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.admin.v1.Sys/CreateDict")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDict(ctx, req.(*CreateDictReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDictResp)
		return ctx.Result(200, reply)
	}
}

func _Sys_UpdateDict0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.admin.v1.Sys/UpdateDict")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDict(ctx, req.(*UpdateDictReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Sys_DeleteDict0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDictReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.admin.v1.Sys/DeleteDict")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDict(ctx, req.(*DeleteDictReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Sys_ListDict0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDictReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.admin.v1.Sys/ListDict")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDict(ctx, req.(*ListDictReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDictResp)
		return ctx.Result(200, reply)
	}
}

func _Sys_PageListDict0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageListDictReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.admin.v1.Sys/PageListDict")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PageListDict(ctx, req.(*PageListDictReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageListDictResp)
		return ctx.Result(200, reply)
	}
}

func _Sys_GetDict0_HTTP_Handler(srv SysHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.shop.admin.v1.Sys/GetDict")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDict(ctx, req.(*GetDictReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDictResp)
		return ctx.Result(200, reply)
	}
}

type SysHTTPClient interface {
	CreateDict(ctx context.Context, req *CreateDictReq, opts ...http.CallOption) (rsp *CreateDictResp, err error)
	DeleteDict(ctx context.Context, req *DeleteDictReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetDict(ctx context.Context, req *GetDictReq, opts ...http.CallOption) (rsp *GetDictResp, err error)
	ListDict(ctx context.Context, req *ListDictReq, opts ...http.CallOption) (rsp *ListDictResp, err error)
	PageListDict(ctx context.Context, req *PageListDictReq, opts ...http.CallOption) (rsp *PageListDictResp, err error)
	UpdateDict(ctx context.Context, req *UpdateDictReq, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type SysHTTPClientImpl struct {
	cc *http.Client
}

func NewSysHTTPClient(client *http.Client) SysHTTPClient {
	return &SysHTTPClientImpl{client}
}

func (c *SysHTTPClientImpl) CreateDict(ctx context.Context, in *CreateDictReq, opts ...http.CallOption) (*CreateDictResp, error) {
	var out CreateDictResp
	pattern := "/api/dict"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.shop.admin.v1.Sys/CreateDict"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysHTTPClientImpl) DeleteDict(ctx context.Context, in *DeleteDictReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/api/dict/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.admin.v1.Sys/DeleteDict"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysHTTPClientImpl) GetDict(ctx context.Context, in *GetDictReq, opts ...http.CallOption) (*GetDictResp, error) {
	var out GetDictResp
	pattern := "/api/dict/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.admin.v1.Sys/GetDict"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysHTTPClientImpl) ListDict(ctx context.Context, in *ListDictReq, opts ...http.CallOption) (*ListDictResp, error) {
	var out ListDictResp
	pattern := "/api/dicts"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.admin.v1.Sys/ListDict"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysHTTPClientImpl) PageListDict(ctx context.Context, in *PageListDictReq, opts ...http.CallOption) (*PageListDictResp, error) {
	var out PageListDictResp
	pattern := "/api/dict"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.shop.admin.v1.Sys/PageListDict"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *SysHTTPClientImpl) UpdateDict(ctx context.Context, in *UpdateDictReq, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/api/dict/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.shop.admin.v1.Sys/UpdateDict"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}