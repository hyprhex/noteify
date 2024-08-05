package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Application version
// TODO: we will make it automatically at build time
const version = "1.0.0"

// Application Config
type config struct {
	port int
	env  string
}

// Application dependencies
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Config Instance
	var cfg config

	// Use command-line flags to change port and env
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Enviroment (development|staging|production)")
	flag.Parse()

	// New logger prefixed with date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Application dependencies instance
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Server settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
