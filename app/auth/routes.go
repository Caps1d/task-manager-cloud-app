package main

import (
	"context"
	"log"

	"github.com/Caps1d/task-manager-cloud-app/auth/internals/sessions"
	"github.com/Caps1d/task-manager-cloud-app/auth/pb"
)

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.infoLog.Print("New auth register request")

	email := r.GetEmail()
	username := r.GetUsername()
	password := r.GetPassword()

	// register
	err := s.users.Insert(username, email, password)
	if err != nil {
		s.errorLog.Print("Auth: registration failed")
		return nil, err
	}

	// call user service

	s.infoLog.Printf("Auth: user - %v registered", username)

	return &pb.RegisterResponse{
		Success: true,
	}, nil
}

func (s *server) IsAuthenticated(ctx context.Context, r *pb.IsAuthenticatedRequest) (*pb.IsAuthenticatedResponse, error) {
	s.infoLog.Printf("pb IsAuthenticatedRequest, received: %v", r.GetSessionID())

	// check if user is already logged in
	someID := r.GetSessionID()
	userID, err := sessions.Get(s.kv, someID)
	if err != nil {
		s.errorLog.Printf("AUTH: Error at IsAuthenticated, failed to get sessionID from kv")
		return &pb.IsAuthenticatedResponse{
			Success: false,
		}, err
	}

	return &pb.IsAuthenticatedResponse{
		UserID:  userID,
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

	log.Printf("userID: %v", userID)

	// create session field in Redis
	// store sessionID key with userID value
	sessionID, err := sessions.Create(s.kv, userID)
	if err != nil {
		log.Printf("Error at session create: %v", err)
	}
	log.Printf("sessionID: %v", sessionID)

	// return sessionID to client
	// generate cookie in api-gateway with sessionID
	return &pb.LoginResponse{
		Id:      sessionID,
		Success: true,
	}, nil
}

func (s *server) Logout(ctx context.Context, r *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	s.infoLog.Printf("New pb logout request")

	sessionID := r.GetSessionID()

	if _, err := sessions.Get(s.kv, sessionID); err != nil {
		s.errorLog.Printf("Coudn't find session with this id, %v", err)
		return &pb.LogoutResponse{
			Success: false,
		}, nil
	}

	if err := sessions.Destroy(s.kv, sessionID); err != nil {
		s.errorLog.Printf("Error while destroying session")
		return &pb.LogoutResponse{
			Success: false,
		}, err
	}

	if id, _ := sessions.Get(s.kv, sessionID); id == 0 {
		s.infoLog.Printf("pb Logout: Session destroyed successfully")
	}

	return &pb.LogoutResponse{
		Success: true,
	}, nil
}
