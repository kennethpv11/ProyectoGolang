package main

import (
	"CRUDSimulador/proto"
	"CRUDSimulador/service"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverGrpc := grpc.NewServer()
	proto.RegisterElectricityServiceServer(serverGrpc, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := serverGrpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
