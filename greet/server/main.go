package greet

import (
	"log"
	"net"

	pb "github.com/MadhavKumaar/go-learn/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to listen: %v\n", err)
	}
	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to server: %v\n", err)
	}

}
