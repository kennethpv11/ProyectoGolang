package service

import (
	"CRUDSimulador/models"
	"CRUDSimulador/proto"
	"context"
	"log"

	"github.com/gofrs/uuid"
	"github.com/uptrace/bun"
)

type Server struct {
	proto.ElectricityServiceServer
	db *bun.DB
}

func (s *Server) CreateElecticity(ctx context.Context, request *proto.ElectricityRequest) (*proto.ElectricityResponse, error) {
	id, err := uuid.FromString(request.Id)
	if err != nil {
		return nil, err
	}
	log.Printf("Received: %v", request.GetId())
	toInsert := models.ElectricityInfo{
		ValueMin: int(request.ValueMin),
		ValueMax: int(request.ValueMax),
		Id:       id,
		Measure:  request.Measure,
	}
	_, errDb := s.db.NewInsert().Model(&toInsert).Exec(ctx)
	if errDb != nil {
		log.Printf("error when insert value: %v", errDb.Error())
		return nil, errDb
	}
	return &proto.ElectricityResponse{Message: "success"}, nil
}

func NewServer(con *bun.DB) *Server {
	return &Server{
		db: con,
	}
}
