// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rpc/kubernetes.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for KubernetesDeploy service

func NewKubernetesDeployEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for KubernetesDeploy service

type KubernetesDeployService interface {
	CheckStatus(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetKubernetesConfig(ctx context.Context, in *ConfigRequest, opts ...client.CallOption) (*ConfigResponse, error)
}

type kubernetesDeployService struct {
	c    client.Client
	name string
}

func NewKubernetesDeployService(name string, c client.Client) KubernetesDeployService {
	return &kubernetesDeployService{
		c:    c,
		name: name,
	}
}

func (c *kubernetesDeployService) CheckStatus(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "KubernetesDeploy.CheckStatus", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesDeployService) GetKubernetesConfig(ctx context.Context, in *ConfigRequest, opts ...client.CallOption) (*ConfigResponse, error) {
	req := c.c.NewRequest(c.name, "KubernetesDeploy.GetKubernetesConfig", in)
	out := new(ConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KubernetesDeploy service

type KubernetesDeployHandler interface {
	CheckStatus(context.Context, *Request, *Response) error
	GetKubernetesConfig(context.Context, *ConfigRequest, *ConfigResponse) error
}

func RegisterKubernetesDeployHandler(s server.Server, hdlr KubernetesDeployHandler, opts ...server.HandlerOption) error {
	type kubernetesDeploy interface {
		CheckStatus(ctx context.Context, in *Request, out *Response) error
		GetKubernetesConfig(ctx context.Context, in *ConfigRequest, out *ConfigResponse) error
	}
	type KubernetesDeploy struct {
		kubernetesDeploy
	}
	h := &kubernetesDeployHandler{hdlr}
	return s.Handle(s.NewHandler(&KubernetesDeploy{h}, opts...))
}

type kubernetesDeployHandler struct {
	KubernetesDeployHandler
}

func (h *kubernetesDeployHandler) CheckStatus(ctx context.Context, in *Request, out *Response) error {
	return h.KubernetesDeployHandler.CheckStatus(ctx, in, out)
}

func (h *kubernetesDeployHandler) GetKubernetesConfig(ctx context.Context, in *ConfigRequest, out *ConfigResponse) error {
	return h.KubernetesDeployHandler.GetKubernetesConfig(ctx, in, out)
}