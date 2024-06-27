package main

import (
	"net/http"
	"time"

	// "fmt"
	"log"
	"os"

	"github.com/Caps1d/task-manager-cloud-app/api-gateway/internal/config"
	"github.com/go-playground/form/v4"
)

type Application struct {
	cfg *config.Config
	// tasks         models.TaskModelInterface -> figure out how to interact with microservices
	// users         models.UserModelInterface
	// notifications models.NotificationModelInterface
	infoLog     *log.Logger
	errorLog    *log.Logger
	formDecoder *form.Decoder
	//sessionManager?
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfg, err := config.NewConfig()
	if err != nil {
		errorLog.Fatal(err)
	}
	formDecoder := form.NewDecoder()

	// app
	app := &Application{
		cfg:         &cfg,
		infoLog:     infoLog,
		errorLog:    errorLog,
		formDecoder: formDecoder,
	}

	srv := &http.Server{
		Addr:         cfg.Port,
		ErrorLog:     app.errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on port%v", cfg.Port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
