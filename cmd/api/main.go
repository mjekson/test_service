package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"test_service.mjekson.ru/internal/data"
)

const version = "1.0.0"

type application struct {
	config Config
	logger *log.Logger
	models data.Models
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "", "path to config file") //../../configs/api.toml
}

func main() {
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	cfg := NewConfig()

	_, err := toml.DecodeFile(configPath, &cfg)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.Fatal(err)
	}

}

func openDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.Db.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Db.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.Db.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
