package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/jun06t/grpc-sample/unary/proto"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &helloHandler{})
	health.RegisterHealthServer(s, &healthHandler{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
