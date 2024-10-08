package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/Caps1d/task-manager-cloud-app/task/config"
	"github.com/Caps1d/task-manager-cloud-app/task/internals/models"
	"github.com/Caps1d/task-manager-cloud-app/task/pb"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

var lis *net.Listener

type server struct {
	tasks    models.ITaskModel
	projects models.IProjectModel
	infoLog  *log.Logger
	errorLog *log.Logger
	pb.UnimplementedTaskServiceServer
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(cfg.DBUrl)
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Print("DB connection established...")
	defer db.Close()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		infoLog.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &server{tasks: &models.TaskModel{DB: db}, projects: &models.ProjectModel{DB: db}, infoLog: infoLog, errorLog: errorLog})
	infoLog.Printf("Starting Task server on: %v", cfg.Port)
	if err := s.Serve(lis); err != nil {
		infoLog.Fatalf("Failed to serve: %v", err)
	}
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Printf("Ping sent to %v", dsn)

	return conn, nil
}
