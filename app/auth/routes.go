package main

import (
	"context"

	"github.com/Caps1d/task-manager-cloud-app/auth/pb"
)

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.infoLog.Print("New pb register request")

	email := r.GetEmail()
	username := r.GetUsername()
	password := r.GetPassword()

	// register
	err := s.users.Insert(username, email, password)
	if err != nil {
		s.errorLog.Print("Auth: registration failed")
		return nil, err
	}

	s.infoLog.Printf("Auth: registered with %v", email)

	// Needs to be reworked
	return &pb.RegisterResponse{
		Success: true,
	}, nil
}

func (s *server) Login(ctx context.Context, r *pb.LoginRequest) (*pb.LoginResponse, error) {
	s.infoLog.Printf("New pb login request, received: %v", r.GetEmail())

	email := r.GetEmail()
	password := r.GetPassword()

	// authenticate
	userID, err := s.users.Authenticate(email, password)
	if err != nil {
		return nil, err
	}

	// generate sessionID
	// create session field in Redis
	// store sessionID key with userID value
	// return sessionID to client
	// generate cookie in api-gateway with sessionID

	// Needs to support sessions
	// return cookie or sessionID
	return &pb.LoginResponse{
		Success: true,
		Id:      int32(userID),
	}, nil
}
