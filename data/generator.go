package data

import (
	"golang_grpc_unary_stream/pb"
)

func NewPlant() *pb.Plant {
	return &pb.Plant{
		Id:   randomPlantId(),
		Name: randomPlantName(),
		Category: &pb.Category{
			Category: randomPlantCategory(),
		},
		Price:       randomFloat(50.0, 500.0),
		Avatar:      randomPlantImage(),
		Description: randomPlantDescription(),
	}
}
