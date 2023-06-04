package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserServer struct {
	// define pb file vals
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed")
	}

	s := grpc.NewServer()
	lis.Addr()
	fmt.Print(s)
}
