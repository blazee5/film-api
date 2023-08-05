package auth

import (
	"context"
	"fmt"
	"github.com/blazee5/film-api/lib/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func ServerInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	if info.FullMethod == "/films.UserService/GetUser" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	h, err := handler(ctx, req)

	return h, err
}

func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	bearerToken := authHeader[0]

	token := strings.Fields(bearerToken)[1]

	_, err := auth.ParseToken(token)

	fmt.Println(ctx)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, err.Error())
	}

	return nil
}
