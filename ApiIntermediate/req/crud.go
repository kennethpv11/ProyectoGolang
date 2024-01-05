package req

import (
	"apiintermediate/models"
	"apiintermediate/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func CreateElectricityRpc(message models.ElectronicDevice) (*string, error) {
	conn, err := grpc.Dial("dns:localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	request := proto.ElectricityRequest{
		ValueMin: int32(message.ValueMin),
		ValueMax: int32(message.ValueMax),
		Measure:  message.Measure,
		Id:       message.Id.String(),
	}
	stub := proto.NewElectricityServiceClient(conn)
	response, err := stub.CreateElecticity(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return &response.Message, nil
}
