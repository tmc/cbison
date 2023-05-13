// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package generativelanguage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	generativelanguagepb "google.golang.org/genproto/googleapis/ai/generativelanguage/v1beta2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newDiscussClientHook clientHook

// DiscussCallOptions contains the retry settings for each method of DiscussClient.
type DiscussCallOptions struct {
	GenerateMessage    []gax.CallOption
	CountMessageTokens []gax.CallOption
}

func defaultDiscussGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("generativelanguage.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("generativelanguage.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultDiscussCallOptions() *DiscussCallOptions {
	return &DiscussCallOptions{
		GenerateMessage:    []gax.CallOption{},
		CountMessageTokens: []gax.CallOption{},
	}
}

func defaultDiscussRESTCallOptions() *DiscussCallOptions {
	return &DiscussCallOptions{
		GenerateMessage:    []gax.CallOption{},
		CountMessageTokens: []gax.CallOption{},
	}
}

// internalDiscussClient is an interface that defines the methods available from Generative Language API.
type internalDiscussClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateMessage(context.Context, *generativelanguagepb.GenerateMessageRequest, ...gax.CallOption) (*generativelanguagepb.GenerateMessageResponse, error)
	CountMessageTokens(context.Context, *generativelanguagepb.CountMessageTokensRequest, ...gax.CallOption) (*generativelanguagepb.CountMessageTokensResponse, error)
}

// DiscussClient is a client for interacting with Generative Language API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// An API for using Generative Language Models (GLMs) in dialog applications.
//
// Also known as large language models (LLMs), this API provides models that
// are trained for multi-turn dialog.
type DiscussClient struct {
	// The internal transport-dependent client.
	internalClient internalDiscussClient

	// The call options for this service.
	CallOptions *DiscussCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *DiscussClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *DiscussClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *DiscussClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateMessage generates a response from the model given an input MessagePrompt.
func (c *DiscussClient) GenerateMessage(ctx context.Context, req *generativelanguagepb.GenerateMessageRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateMessageResponse, error) {
	return c.internalClient.GenerateMessage(ctx, req, opts...)
}

// CountMessageTokens runs a model’s tokenizer on a string and returns the token count.
func (c *DiscussClient) CountMessageTokens(ctx context.Context, req *generativelanguagepb.CountMessageTokensRequest, opts ...gax.CallOption) (*generativelanguagepb.CountMessageTokensResponse, error) {
	return c.internalClient.CountMessageTokens(ctx, req, opts...)
}

// discussGRPCClient is a client for interacting with Generative Language API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type discussGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing DiscussClient
	CallOptions **DiscussCallOptions

	// The gRPC API client.
	discussClient generativelanguagepb.DiscussServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewDiscussClient creates a new discuss service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// An API for using Generative Language Models (GLMs) in dialog applications.
//
// Also known as large language models (LLMs), this API provides models that
// are trained for multi-turn dialog.
func NewDiscussClient(ctx context.Context, opts ...option.ClientOption) (*DiscussClient, error) {
	clientOpts := defaultDiscussGRPCClientOptions()
	if newDiscussClientHook != nil {
		hookOpts, err := newDiscussClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := DiscussClient{CallOptions: defaultDiscussCallOptions()}

	c := &discussGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		discussClient:    generativelanguagepb.NewDiscussServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *discussGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *discussGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *discussGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type discussRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing DiscussClient
	CallOptions **DiscussCallOptions
}

// NewDiscussRESTClient creates a new discuss service rest client.
//
// An API for using Generative Language Models (GLMs) in dialog applications.
//
// Also known as large language models (LLMs), this API provides models that
// are trained for multi-turn dialog.
func NewDiscussRESTClient(ctx context.Context, opts ...option.ClientOption) (*DiscussClient, error) {
	clientOpts := append(defaultDiscussRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultDiscussRESTCallOptions()
	c := &discussRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &DiscussClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultDiscussRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://generativelanguage.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://generativelanguage.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *discussRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *discussRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *discussRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *discussGRPCClient) GenerateMessage(ctx context.Context, req *generativelanguagepb.GenerateMessageRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateMessageResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateMessage[0:len((*c.CallOptions).GenerateMessage):len((*c.CallOptions).GenerateMessage)], opts...)
	var resp *generativelanguagepb.GenerateMessageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.discussClient.GenerateMessage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *discussGRPCClient) CountMessageTokens(ctx context.Context, req *generativelanguagepb.CountMessageTokensRequest, opts ...gax.CallOption) (*generativelanguagepb.CountMessageTokensResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CountMessageTokens[0:len((*c.CallOptions).CountMessageTokens):len((*c.CallOptions).CountMessageTokens)], opts...)
	var resp *generativelanguagepb.CountMessageTokensResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.discussClient.CountMessageTokens(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GenerateMessage generates a response from the model given an input MessagePrompt.
func (c *discussRESTClient) GenerateMessage(ctx context.Context, req *generativelanguagepb.GenerateMessageRequest, opts ...gax.CallOption) (*generativelanguagepb.GenerateMessageResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v:generateMessage", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GenerateMessage[0:len((*c.CallOptions).GenerateMessage):len((*c.CallOptions).GenerateMessage)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.GenerateMessageResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// CountMessageTokens runs a model’s tokenizer on a string and returns the token count.
func (c *discussRESTClient) CountMessageTokens(ctx context.Context, req *generativelanguagepb.CountMessageTokensRequest, opts ...gax.CallOption) (*generativelanguagepb.CountMessageTokensResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta2/%v:countMessageTokens", req.GetModel())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "model", url.QueryEscape(req.GetModel())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CountMessageTokens[0:len((*c.CallOptions).CountMessageTokens):len((*c.CallOptions).CountMessageTokens)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.CountMessageTokensResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}