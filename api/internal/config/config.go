package config

import "flag"

type Config struct {
	Addr  string
	DBUrl string
	// StaticDir string
}

func NewConfig() Config {
	var cfg Config

	fs := flag.NewFlagSet("config", flag.ExitOnError)

	fs.StringVar(&cfg.Addr, "addr", ":8080", "HTTP network address")
	fs.StringVar(&cfg.DBUrl, "db", "postgres://web:test@localhost:5432/task-manager", "Database URL")

	flag.Parse()

	return cfg
}
