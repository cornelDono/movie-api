package main

import (
	"context"
	"fmt"
	"log"
	"movie_api.net/internal/data"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

const version = "1.0.0"

type postgresConfig struct {
	user     string
	password string
	host     string
	port     string
	dbName   string
}
type config struct {
	port int
	env  string
	db   postgresConfig
}

type application struct {
	config config
	logger *log.Logger
	models data.Models
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	cfg, err := loadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	dbpool, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer dbpool.Close()

	app := application{
		config: cfg,
		logger: logger,
		models: data.NewModels(dbpool),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s severs on %s", cfg.env, srv.Addr)
	listenErr := srv.ListenAndServe()
	logger.Fatal(listenErr)
}

func openDB(cfg config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.db.user, cfg.db.password, cfg.db.host, cfg.db.port, cfg.db.dbName)
	dbpool, err := pgxpool.New(context.Background(), dsn)
	dbpool.Config().MaxConnIdleTime, _ = time.ParseDuration("15m")
	dbpool.Config().MaxConns = 50

	if err != nil {
		return nil, err
	}

	return dbpool, nil
}

func loadConfig() (config, error) {
	cfg := config{
		port: 4000,
		env:  "development",
		db: postgresConfig{
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			dbName:   os.Getenv("DB_NAME"),
		},
	}

	return cfg, nil
}
