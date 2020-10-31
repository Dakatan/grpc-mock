package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"log"
)

func NewUnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Printf("req: %v", req)
		resp, err = handler(ctx, req)
		log.Printf("res: %v", resp)
		return
	}
}
