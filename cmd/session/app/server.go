package app

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"mock-grpc/pkg/interceptor"
	"mock-grpc/pkg/proto"
	"net"
	"os"
)

func NewSessionServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			server := NewServer()
			return Run(server)
		},
	}
	return cmd
}

func NewServer() *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.NewUnaryLoggingInterceptor(),
			interceptor.NewUnaryHeaderChainInterceptor()),
	)
	secretKey := os.Getenv("SECRET_KEY")
	if len(secretKey) == 0 {
		secretKey = "None"
	}
	proto.RegisterSessionServer(server, &SessionService{SecretKey: secretKey})
	return server
}

func Run(server *grpc.Server) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	return server.Serve(listener)
}
