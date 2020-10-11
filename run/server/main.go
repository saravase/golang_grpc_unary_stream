package main

import (
	"flag"
	"fmt"
	"golang_grpc_unary_stream/pb"
	"golang_grpc_unary_stream/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	// Int defines an int flag with specified name, default value, and usage string
	port := flag.Int("port", 0, "Server port")
	// Parse parses the command-line flags from os.Args[1:]
	flag.Parse()
	log.Printf("Start the server on port: %d", *port)

	plantServer := service.NewPlantServer(service.NewInMemoryPlantStore())
	// NewServer creates a gRPC server which has no service registered and has not started to accept requests yet
	grpcServer := grpc.NewServer()
	pb.RegisterPlantServiceServer(grpcServer, plantServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	// Listen announces on the local network address
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Can't start the server:", err)
	}

	// Serve accepts incoming connections on the listener lis,
	// creating a new ServerTransport and service goroutine for each
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Can't start the server", err)
	}
}
