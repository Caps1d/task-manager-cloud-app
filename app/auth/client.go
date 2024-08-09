package main

import (
	"log"
	"net"

	"github.com/Caps1d/task-manager-cloud-app/auth/config"
	"github.com/Caps1d/task-manager-cloud-app/auth/pb"
	"google.golang.org/grpc"
)

var lis *net.Listener

type server struct {
	pb.UnimplementedAuthServiceServer
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Printf("Starting Auth server on: %v", cfg.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
