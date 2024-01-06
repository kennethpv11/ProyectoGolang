package main

import (
	"CRUDSimulador/bd"
	"CRUDSimulador/proto"
	"CRUDSimulador/service"
	"context"
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
	con := bd.ConectToBd()
	bd.Migrations(context.Background(), *con)
	defer bd.CloseToBd(con)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverGrpc := grpc.NewServer()
	proto.RegisterElectricityServiceServer(serverGrpc, service.NewServer(con))
	log.Printf("server listening at %v", lis.Addr())
	if err := serverGrpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
