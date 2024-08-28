package main

import (
	"context"

	"github.com/Caps1d/task-manager-cloud-app/user/pb"
)

func (s *server) Register(ctx context.Context, r *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.infoLog.Print("New auth register request")

	email := r.GetEmail()
	username := r.GetUsername()
	name := r.GetName()

	s.infoLog.Printf("User: received %v, %v, %v", email, username, name)

	// register
	err := s.users.Insert(name, email, username)
	if err != nil {
		s.errorLog.Printf("User: error at register handler %v", err)
		return nil, err
	}

	s.infoLog.Printf("User: %v succesfully registered", username)

	return &pb.RegisterResponse{
		Success: true,
	}, nil
}

func (s *server) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	uid := r.GetId()

	user, err := s.users.Get(uid)
	if err != nil {
		s.errorLog.Printf("User: error at get handler %v", err)
		return nil, err
	}

	data := &pb.GetUser{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		TeamID:   user.TeamID,
	}

	return &pb.GetUserResponse{Data: data}, nil
}
