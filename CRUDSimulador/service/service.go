package service

import (
	"CRUDSimulador/proto"
	"context"
	"log"
)

type Server struct {
	proto.ElectricityServiceServer
}

func (s *Server) CreateElectricity(ctx context.Context, request *proto.ElectricityRequest) (*proto.ElectricityResponse, error) {
	log.Printf("Received: %v", request.Id)
	return &proto.ElectricityResponse{Message: request.Id}, nil
}
