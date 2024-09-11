package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
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

	s.infoLog.Printf("Auth: user - %v registered", username)

	return &pb.RegisterResponse{
		Success: true,
	}, nil
}

func (s *server) IsAuthenticated(ctx context.Context, r *pb.IsAuthenticatedRequest) (*pb.IsAuthenticatedResponse, error) {
	s.infoLog.Printf("pb IsAuthenticatedRequest, received: %v", r.GetSessionID())

	// check if user is already logged in
	cryptoText := r.GetSessionID()
	// decrypt the sessionID
	encryptedID, _ := base64.StdEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher([]byte(s.cfg.EncKey))
	if err != nil {
		s.errorLog.Printf("AUTH: error at NewCipher in IsAuthenticated handler, %s", err)
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, []byte(s.cfg.InitVec))
	mode.CryptBlocks(encryptedID, encryptedID)
	// unpad the decrypted id
	sessionID := encryptedID[:(len(encryptedID) - int(encryptedID[len(encryptedID)-1]))]

	userID, err := sessions.Get(s.kv, string(sessionID))
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
		s.errorLog.Printf("Error at session create: %v", err)
	}
	s.infoLog.Printf("sessionID: %v", sessionID)

	// encrypt sessionID
	var plainTextBlock []byte
	length := len(sessionID)

	copy(plainTextBlock, sessionID)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	block, err := aes.NewCipher([]byte(s.cfg.EncKey))
	if err != nil {
		s.errorLog.Printf("AUTH: error at NewCipher in Login handler, %s", err)
		return nil, err
	}
	cryptoSessionID := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(s.cfg.InitVec))
	mode.CryptBlocks(cryptoSessionID, plainTextBlock)

	return &pb.LoginResponse{
		Id:      base64.StdEncoding.EncodeToString(cryptoSessionID),
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
