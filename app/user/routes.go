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
	id := r.GetId()

	// refactor into a helper function -> DRY
	exists, err := s.users.Exist(id)
	if !exists {
		s.errorLog.Println("User: sql error in GetUser handler, user doesn't exist")
		return nil, err
	}

	user, err := s.users.Get(id)
	if err != nil {
		s.errorLog.Printf("User: error in GetUser handler %v", err)
		return nil, err
	}

	data := &pb.GetUser{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	}

	return &pb.GetUserResponse{Data: data}, nil
}

func (s *server) UpdateUser(ctx context.Context, r *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id := r.GetUserID()
	email := r.GetEmail()

	// DRY
	exists, err := s.users.Exist(id)
	if !exists {
		s.errorLog.Println("User: error in UpdateUser handler, user doesn't exist")
		return nil, err
	}

	if len(email) > 0 {
		err := s.users.UpdateEmail(id, email)
		if err != nil {
			s.errorLog.Printf("User: error in UpdateUser handler while updating email %v", err)
			return &pb.UpdateUserResponse{
				Success: false,
			}, err
		}
	}

	return &pb.UpdateUserResponse{
		Success: true,
	}, nil
}

func (s *server) CreateTeam(ctx context.Context, r *pb.CreateTeamRequest) (*pb.CreateTeamResponse, error) {
	id := r.GetName()
	manager := r.GetManager()

	err := s.teams.Insert(id, manager)
	if err != nil {
		s.errorLog.Printf("User: sql error in CreateTeam handler %v", err)
		return nil, err
	}

	return &pb.CreateTeamResponse{Success: true}, nil
}

func (s *server) GetTeam(ctx context.Context, r *pb.GetTeamRequest) (*pb.GetTeamResponse, error) {
	id := r.GetId()

	team, err := s.teams.GetTeam(id)
	if err != nil {
		s.errorLog.Printf("User: sql error in GetTeam handler %v", err)
		return nil, err
	}

	pbTeam := &pb.Team{
		Id:      team.ID,
		Name:    team.Name,
		Manager: team.Manager,
	}
	for _, member := range team.Members {
		pbTeam.Members = append(pbTeam.Members, &pb.Member{Id: member.ID, Name: member.Name, Email: member.Email, Username: member.Username, Role: member.Role})
	}

	return &pb.GetTeamResponse{Team: pbTeam}, nil
}

func (s *server) UpdateTeam(ctx context.Context, r *pb.UpdateTeamRequest) (*pb.UpdateTeamResponse, error) {
	id := r.GetId()
	name := r.GetName()
	manager := r.GetManager()
	userID := r.GetUserId()
	role := r.GetRole()

	if len(name) > 0 {
		err := s.teams.UpdateName(id, name)
		if err != nil {
			s.errorLog.Printf("User: sql error in UpdateTeam handler, name case %v", err)
			return nil, err
		}
	}

	if len(manager) > 0 {
		err := s.teams.UpdateManager(id, manager)
		if err != nil {
			s.errorLog.Printf("User: sql error in UpdateTeam handler, manager case %v", err)
			return nil, err
		}
	}

	if len(role) > 0 {
		err := s.teams.UpdateMemberRole(id, userID, role)
		if err != nil {
			s.errorLog.Printf("User: sql error in UpdateTeam handler, role case %v", err)
			return nil, err
		}
	}

	return &pb.UpdateTeamResponse{Success: true}, nil
}

func (s *server) DeleteTeam(ctx context.Context, r *pb.DeleteTeamRequest) (*pb.DeleteTeamResponse, error) {
	id := r.GetId()

	err := s.teams.Delete(id)
	if err != nil {
		s.errorLog.Printf("User: sql error in DeleteTeam handler %v", err)
		return nil, err
	}

	return &pb.DeleteTeamResponse{Success: true}, nil
}
