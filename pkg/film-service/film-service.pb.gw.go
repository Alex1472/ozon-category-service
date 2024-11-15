// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/film-service/film-service.proto

/*
Package film_service is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package film_service

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

func request_FilmService_GetCategoryById_0(ctx context.Context, marshaler runtime.Marshaler, client FilmServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetCategoryByIdRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetCategoryById(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_FilmService_GetCategoryById_0(ctx context.Context, marshaler runtime.Marshaler, server FilmServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetCategoryByIdRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetCategoryById(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterFilmServiceHandlerServer registers the http handlers for service FilmService to "mux".
// UnaryRPC     :call FilmServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterFilmServiceHandlerFromEndpoint instead.
func RegisterFilmServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server FilmServiceServer) error {

	mux.Handle("POST", pattern_FilmService_GetCategoryById_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/alex1472.ozon_film_service.film_service.v1.FilmService/GetCategoryById", runtime.WithHTTPPathPattern("/alex1472.ozon_film_service.film_service.v1.FilmService/GetCategoryById"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_FilmService_GetCategoryById_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FilmService_GetCategoryById_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterFilmServiceHandlerFromEndpoint is same as RegisterFilmServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterFilmServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
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

	return RegisterFilmServiceHandler(ctx, mux, conn)
}

// RegisterFilmServiceHandler registers the http handlers for service FilmService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterFilmServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterFilmServiceHandlerClient(ctx, mux, NewFilmServiceClient(conn))
}

// RegisterFilmServiceHandlerClient registers the http handlers for service FilmService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "FilmServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "FilmServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "FilmServiceClient" to call the correct interceptors.
func RegisterFilmServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client FilmServiceClient) error {

	mux.Handle("POST", pattern_FilmService_GetCategoryById_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/alex1472.ozon_film_service.film_service.v1.FilmService/GetCategoryById", runtime.WithHTTPPathPattern("/alex1472.ozon_film_service.film_service.v1.FilmService/GetCategoryById"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_FilmService_GetCategoryById_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_FilmService_GetCategoryById_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_FilmService_GetCategoryById_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"alex1472.ozon_film_service.film_service.v1.FilmService", "GetCategoryById"}, ""))
)

var (
	forward_FilmService_GetCategoryById_0 = runtime.ForwardResponseMessage
)
