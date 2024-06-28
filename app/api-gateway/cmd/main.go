package main

import (
	"net/http"
	"os"

	"github.com/Caps1d/task-manager-cloud-app/api-gateway/internal/config"
	pb "github.com/Caps1d/task-manager-cloud-app/auth/pb"
	"github.com/go-playground/form/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

var e *echo.Echo
var authClient pb.AuthServiceClient

const loggerHeaders = "${time_rfc3339} ${level} ${prefix} ${short_file} ${line}"

type Application struct {
	cfg *config.Config
	// tasks         models.TaskModelInterface -> figure out how to interact with microservices
	// users         models.UserModelInterface
	// notifications models.NotificationModelInterface
	infoLog     echo.Logger
	errorLog    echo.Logger
	formDecoder *form.Decoder
	//sessionManager?
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

	infoLog.Printf("grpc dial Auth at: %v", cfg.AuthSvcUrl)
	authConn, err := grpc.Dial(cfg.AuthSvcUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		errorLog.Fatalf("Did not connect to auth server: %v", err)
	}
	defer authConn.Close()
	authClient = pb.NewAuthServiceClient(authConn)

	initRoutes()

	infoLog.Infof("Starting server on port%v", cfg.Port)
	if err := e.Start(cfg.Port); err != http.ErrServerClosed {
		errorLog.Fatal(err)
	}
}
