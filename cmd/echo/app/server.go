package app

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"mock-grpc/pkg/interceptor"
	"mock-grpc/pkg/proto"
	"net"
	"os"
)

func NewEchoServerCommand() *cobra.Command {
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
	proto.RegisterEchoServer(server, &EchoService{
		Voice:      getVoice(),
		EchoClient: getClient(),
	})
	return server
}

func Run(server *grpc.Server) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	return server.Serve(listener)
}

func getVoice() string {
	voice := os.Getenv("ECHO_VOICE")
	if len(voice) == 0 {
		voice = "..."
	}
	return voice
}

func getClient() proto.EchoClient {
	address := os.Getenv("ECHO_ADDRESS")
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	return proto.NewEchoClient(conn)
}
