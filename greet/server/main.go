package main

import (
	"fmt"
	"net"

	pb "github.com/MadhavKumaar/go/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
	}

	fmt.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		fmt.Printf("Failed to server: %v\n", err)
	}

}
