package app

import (
	"context"
	"mock-grpc/pkg/proto"
)

type EchoService struct {
	Voice string
	EchoClient proto.EchoClient
}

func (e *EchoService) Echo1(ctx context.Context, request *proto.EchoRequest) (*proto.EchoResponse, error) {
	return &proto.EchoResponse{Message: request.Message}, nil
}

func (e *EchoService) Echo2(ctx context.Context, request *proto.EchoRequest) (*proto.EchoResponse, error) {
	return &proto.EchoResponse{Message: e.Voice}, nil
}

func (e *EchoService) ChainedEcho(ctx context.Context, request *proto.ChainedEchoRequest) (*proto.EchoResponse, error) {
	switch request.EchoType {
	case proto.ChainedEchoRequest_Echo1:
		return e.EchoClient.Echo1(ctx, request.EchoRequest)
	case proto.ChainedEchoRequest_Echo2:
		return e.EchoClient.Echo2(ctx, request.EchoRequest)
	default:
		return e.EchoClient.Echo1(ctx, request.EchoRequest)
	}
}
