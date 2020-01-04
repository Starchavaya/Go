package main

import (
	"context"
	pb "github.com/Starchavaya/Go/task3/api/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewIMotherClient(conn)

	// Contact the server and print out its response.
	r, err := c.ListMother(context.Background(), &pb.ListMotherRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("11233333Mothers: %s", r.GetMothers())
}
