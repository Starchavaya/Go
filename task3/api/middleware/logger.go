package middleware

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func LoggerMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Info()
	return handler(ctx, req)
}
