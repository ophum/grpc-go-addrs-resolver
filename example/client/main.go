package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "google.golang.org/grpc/examples/features/proto/echo"

	_ "github.com/ophum/grpc-go-addrs-resolver"
)

func callUnaryEcho(c pb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		log.Printf("could not greet: %v", err)
		return
	}
	fmt.Println(r.Message)
}

func main() {
	conn, err := grpc.NewClient("addrs:///localhost:50051,localhost:50052",
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ec := pb.NewEchoClient(conn)
	for i := 0; i < 10; i++ {
		callUnaryEcho(ec, "hello")
	}
}
