package main

import (
	"net/http"
	"os"

	"github.com/Caps1d/task-manager-cloud-app/api-gateway/internal/config"
	authpb "github.com/Caps1d/task-manager-cloud-app/auth/pb"
	userpb "github.com/Caps1d/task-manager-cloud-app/user/pb"
	"github.com/go-playground/form/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var e *echo.Echo
var authClient authpb.AuthServiceClient
var userClient userpb.UserServiceClient

const loggerHeaders = "${time_rfc3339} ${level} ${prefix} ${short_file} ${line}"

type Application struct {
	cfg         *config.Config
	infoLog     echo.Logger
	errorLog    echo.Logger
	formDecoder *form.Decoder
	//sessionManager? managing sessions with Auth svc
}

var app *Application

func main() {
	e = echo.New()

	infoLog := e.Logger
	infoLog.SetLevel(log.INFO)
	infoLog.SetHeader(loggerHeaders)

	errorLog := e.Logger
	errorLog.SetOutput(os.Stderr)
	errorLog.SetHeader(loggerHeaders)

	cfg, err := config.NewConfig()
	if err != nil {
		errorLog.Fatal(err)
	}

	app = &Application{
		cfg:      &cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	// Connect to Auth service
	infoLog.Printf("grpc dial Auth at: %v", cfg.AuthSvcUrl)
	// look into security settings for grpc
	authConn, err := grpc.NewClient(cfg.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorLog.Fatalf("Did not connect to auth server: %v", err)
	}
	defer authConn.Close()
	authClient = authpb.NewAuthServiceClient(authConn)

	// Connect to User service
	infoLog.Printf("grpc dial User at: %v", cfg.UserSvcUrl)
	userConn, err := grpc.NewClient(cfg.UserSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		errorLog.Fatalf("Did not connect to user server: %v", err)
	}
	defer userConn.Close()
	userClient = userpb.NewUserServiceClient(userConn)

	app.initRoutes()

	infoLog.Infof("Starting server on port%v", cfg.Port)
	if err := e.Start(cfg.Port); err != http.ErrServerClosed {
		errorLog.Fatal(err)
	}
}
