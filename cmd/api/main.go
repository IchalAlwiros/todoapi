package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"goapi.railway.app/cmd/api/database"
	"goapi.railway.app/cmd/api/models"
)

const version = "0.0.1"

type config struct {
	port int
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {


	// CONNECT DB
	if err := database.ConnectionDB(); err != nil {
		log.Fatal("Cannot connect to DB", err.Error())
	}



	// MIGRATION MODEL
	if err := database.DB.AutoMigrate(models.Todo{}); err != nil {
		log.Fatal("Cannot migrate to DB", err)
	}



	var cfg config

	// Try to read environment variable for port (given by railway). Otherwise use default
	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil {
		intPort = 4000
	}

	// Set the port to run the API on
	cfg.port = intPort

	// create the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// create the application
	app := &application{
		config: cfg,
		logger: logger,
	}

	// create the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("server started", "addr", srv.Addr)

	// Start the server
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
