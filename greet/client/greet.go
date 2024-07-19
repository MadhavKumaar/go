package main

import (
	"context"
	"fmt"

	pb "github.com/MadhavKumaar/go/greet/proto"
)

func goGreet(c pb.GreetServiceClient) {
	fmt.Println("doGreet not involed")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Madhav",
	})
	if err != nil {
		fmt.Printf("Could not greet: %v\n", err)
	}

	fmt.Printf("Greetings: %s\n", res.Result)
}
