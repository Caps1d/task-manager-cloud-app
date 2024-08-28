package main

import (
	"context"

	"github.com/Caps1d/task-manager-cloud-app/user/pb"
)

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.infoLog.Print("New auth register request")

	email := r.GetEmail()
	username := r.GetUsername()
	password := r.GetPassword()

	// register

	return &pb.RegisterResponse{
		Success: true,
	}, nil
}
