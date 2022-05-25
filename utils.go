package main

import (
	"context"

	"google.golang.org/grpc/metadata"
)

var (
	httpAuthorizationHeader = "Authorization"
	grpcAuthorizationHeader = "authorization"
)

// contextWithGrpcBearer returns a copy of the parent context with
// outgoing authorization metadata attached.
func contextWithGrpcBearer(parent context.Context, bearer string) context.Context {
	return metadata.AppendToOutgoingContext(parent, grpcAuthorizationHeader, bearer)
}
