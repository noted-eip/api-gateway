package main

import (
	"context"
	"errors"
	"io"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	httpAuthorizationHeader = "Authorization"
	grpcAuthorizationHeader = "authorization"
)

var (
	ErrUnauthenticated = errors.New("unauthenticated")
)

var grpcCodeToHttpCode = map[codes.Code]int{
	codes.OK:                 http.StatusOK,
	codes.Canceled:           http.StatusInternalServerError,
	codes.Unknown:            http.StatusInternalServerError,
	codes.InvalidArgument:    http.StatusBadRequest,
	codes.DeadlineExceeded:   http.StatusRequestTimeout,
	codes.NotFound:           http.StatusNotFound,
	codes.AlreadyExists:      http.StatusConflict,
	codes.PermissionDenied:   http.StatusUnauthorized,
	codes.ResourceExhausted:  http.StatusServiceUnavailable,
	codes.FailedPrecondition: http.StatusFailedDependency,
	codes.Aborted:            http.StatusServiceUnavailable,
	codes.OutOfRange:         http.StatusRequestedRangeNotSatisfiable,
	codes.Unimplemented:      http.StatusNotImplemented,
	codes.Internal:           http.StatusInternalServerError,
	codes.Unavailable:        http.StatusServiceUnavailable,
	codes.DataLoss:           http.StatusInternalServerError,
	codes.Unauthenticated:    http.StatusUnauthorized,
}

// contextWithGrpcBearer returns a copy of the parent context with
// outgoing authorization metadata attached.
func contextWithGrpcBearer(parent context.Context, bearer string) context.Context {
	return metadata.AppendToOutgoingContext(parent, grpcAuthorizationHeader, bearer)
}

// writeError writes to the body of an http request the error passed as
// argument. If the error is a gRPC error, the status code written will
// depend on the gRPC status code.
func writeError(w http.ResponseWriter, code int, err error) {
	s, ok := status.FromError(err)
	if ok {
		translatedCode := grpcCodeToHttpCode[s.Code()]
		if translatedCode == 0 {
			translatedCode = http.StatusInternalServerError
		}
		http.Error(w, s.Message(), translatedCode)
	} else {
		http.Error(w, err.Error(), code)
	}
}

func authenticate(r *http.Request) (string, error) {
	bearer := r.Header.Get(httpAuthorizationHeader)

	if bearer == "" {
		return "", ErrUnauthenticated
	}
	return bearer, nil
}

func convertJsonToProto(body io.ReadCloser, message protoreflect.ProtoMessage) error {
	requestBody, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	err = protojson.Unmarshal(requestBody, message)
	if err != nil {
		return err
	}
	return nil
}
