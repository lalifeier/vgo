// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

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

type PaserHTTPServer interface {
	Paser(context.Context, *PaserReq) (*PaserReply, error)
}

func RegisterPaserHTTPServer(s *http.Server, srv PaserHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/paser", _Paser_Paser0_HTTP_Handler(srv))
}

func _Paser_Paser0_HTTP_Handler(srv PaserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PaserReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.paser.service.v1.Paser/Paser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Paser(ctx, req.(*PaserReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PaserReply)
		return ctx.Result(200, reply)
	}
}

type PaserHTTPClient interface {
	Paser(ctx context.Context, req *PaserReq, opts ...http.CallOption) (rsp *PaserReply, err error)
}

type PaserHTTPClientImpl struct {
	cc *http.Client
}

func NewPaserHTTPClient(client *http.Client) PaserHTTPClient {
	return &PaserHTTPClientImpl{client}
}

func (c *PaserHTTPClientImpl) Paser(ctx context.Context, in *PaserReq, opts ...http.CallOption) (*PaserReply, error) {
	var out PaserReply
	pattern := "/v1/paser"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.paser.service.v1.Paser/Paser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
