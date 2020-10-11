package main

import (
	"context"
	"flag"
	"golang_grpc_unary_stream/data"
	"golang_grpc_unary_stream/pb"

	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	// String defines a string flag with specified name, default value, and usage string
	serverAddress := flag.String("address", "", "Server address")

	// Parse parses the command-line flags from os.Args[1:]
	flag.Parse()
	log.Printf("Dial Server %s", *serverAddress)

	// Dial creates a client connection to the given target.
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Can't dial server : ", err)
	}

	plantClient := pb.NewPlantServiceClient(conn)

	plant := data.NewPlant()
	req := &pb.CreatePlantRequest{
		Plant: plant,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := plantClient.CreatePlant(ctx, req)
	if err != nil {
		// FromError returns a Status representing err
		// if it was produced from this package or has a method `GRPCStatus() *Status`
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Printf("Plant already Exists")
		} else {
			log.Fatal("Can't create plant: ", err)
		}
		return
	}

	log.Printf("Created plant with id: %s", res.Id)

}
