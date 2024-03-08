// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: clicks/clicks.proto

/*
Package clicks is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package clicks

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_SaveClicks_CreateUrlClick_0 = &utilities.DoubleArray{Encoding: map[string]int{"shortUrl": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_SaveClicks_CreateUrlClick_0(ctx context.Context, marshaler runtime.Marshaler, client SaveClicksClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UrlClickRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["shortUrl"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "shortUrl")
	}

	protoReq.ShortUrl, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "shortUrl", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SaveClicks_CreateUrlClick_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateUrlClick(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_SaveClicks_CreateUrlClick_0(ctx context.Context, marshaler runtime.Marshaler, server SaveClicksServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq UrlClickRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["shortUrl"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "shortUrl")
	}

	protoReq.ShortUrl, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "shortUrl", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_SaveClicks_CreateUrlClick_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateUrlClick(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterSaveClicksHandlerServer registers the http handlers for service SaveClicks to "mux".
// UnaryRPC     :call SaveClicksServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterSaveClicksHandlerFromEndpoint instead.
func RegisterSaveClicksHandlerServer(ctx context.Context, mux *runtime.ServeMux, server SaveClicksServer) error {

	mux.Handle("GET", pattern_SaveClicks_CreateUrlClick_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/.SaveClicks/CreateUrlClick", runtime.WithHTTPPathPattern("/api/v1/url/get/{shortUrl}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_SaveClicks_CreateUrlClick_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SaveClicks_CreateUrlClick_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterSaveClicksHandlerFromEndpoint is same as RegisterSaveClicksHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSaveClicksHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSaveClicksHandler(ctx, mux, conn)
}

// RegisterSaveClicksHandler registers the http handlers for service SaveClicks to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSaveClicksHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterSaveClicksHandlerClient(ctx, mux, NewSaveClicksClient(conn))
}

// RegisterSaveClicksHandlerClient registers the http handlers for service SaveClicks
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "SaveClicksClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "SaveClicksClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "SaveClicksClient" to call the correct interceptors.
func RegisterSaveClicksHandlerClient(ctx context.Context, mux *runtime.ServeMux, client SaveClicksClient) error {

	mux.Handle("GET", pattern_SaveClicks_CreateUrlClick_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/.SaveClicks/CreateUrlClick", runtime.WithHTTPPathPattern("/api/v1/url/get/{shortUrl}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_SaveClicks_CreateUrlClick_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_SaveClicks_CreateUrlClick_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_SaveClicks_CreateUrlClick_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "v1", "url", "get", "shortUrl"}, ""))
)

var (
	forward_SaveClicks_CreateUrlClick_0 = runtime.ForwardResponseMessage
)
