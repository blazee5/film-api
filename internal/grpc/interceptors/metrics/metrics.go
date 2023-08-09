package metricsgrpc

import (
	"context"
	"github.com/blazee5/film-api/internal/metrics"
	"google.golang.org/grpc"
)

func MetricsInterceptor(ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	metrics.IncRequests()

	h, err := handler(ctx, req)

	return h, err
}
