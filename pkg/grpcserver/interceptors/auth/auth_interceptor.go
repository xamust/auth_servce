package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ProductToken string

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Print("No metadata")
	}
	values := md.Get("Cookie")
	if len(values) == 0 {
		log.Print("No x-authorization header")
		return resp, nil
	}

	md.Append("x-header", values[0])

	s := ProductToken("token")
	ctx = context.WithValue(ctx, s, values[0])

	metadata.NewOutgoingContext(ctx, md)

	resp, err = handler(ctx, req)

	return resp, err
}
