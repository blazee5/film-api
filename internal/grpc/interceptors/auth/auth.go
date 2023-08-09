package auth

import (
	"context"
	"github.com/blazee5/film-api/lib/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func AuthInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	if info.FullMethod == "/films.UserService/GetUser" {
		newCtx, err := authorize(ctx)

		ctx = newCtx

		if err != nil {
			return nil, err
		}
	}

	h, err := handler(ctx, req)

	return h, err
}

func authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	bearerToken := authHeader[0]

	token := strings.Fields(bearerToken)[1]

	id, err := auth.ParseToken(token)

	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, err.Error())
	}

	ctx = context.WithValue(ctx, "user_id", int(id))

	return ctx, nil
}
