package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

type server struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *server) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println(s.addr, "UnaryEcho")
	return &pb.EchoResponse{
		Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr),
	}, nil
}

func startServer(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{addr: addr})
	log.Printf("serving on %s\n", addr)
	return s.Serve(l)
}
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go func() {
		if err := startServer("localhost:50051"); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := startServer("localhost:50052"); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
}
