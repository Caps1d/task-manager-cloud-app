package main

import (
	"context"
	"net/http"
	"time"

	// "fmt"
	"log"
	"os"

	"github.com/Caps1d/task-manager-cloud-app/api/internal/config"
	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	cfg *config.Config
	// tasks         models.TaskModelInterface
	// users         models.UserModelInterface
	// notifications models.NotificationModelInterface
	infoLog     *log.Logger
	errorLog    *log.Logger
	formDecoder *form.Decoder
	//sessionManager?
}

func main() {
	cfg := config.NewConfig()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// db, err := openDB(cfg.Addr)
	// if err != nil {
	// 	errorLog.Fatal(err)
	// }
	// infoLog.Print("DB connection established...")
	// defer db.Close()

	formDecoder := form.NewDecoder()

	// app
	app := &Application{
		cfg:         &cfg,
		infoLog:     infoLog,
		errorLog:    errorLog,
		formDecoder: formDecoder,
	}

	srv := &http.Server{
		Addr:         cfg.Addr,
		ErrorLog:     app.errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("String starting on server %v", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conn, nil
}
