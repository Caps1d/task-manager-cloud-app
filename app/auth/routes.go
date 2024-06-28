package main

import (
	"context"
	"log"

	"github.com/Caps1d/task-manager-cloud-app/auth/pb"
)

func (s *server) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: %v", r.GetEmail())
	return &pb.LoginResponse{
		Status: 1,
		Token:  "Bingo!",
	}, nil
}
