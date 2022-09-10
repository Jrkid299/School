// cmd/api/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// the application version number
const version = "1.0.0"

// the cofiguration settings
type config struct {
	port int
	env  string // development, staging, production. etc
}

// dependency Injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	// read in the flags that are needed to populate the our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (Development | staging| production)")
	flag.Parse()

	// Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	//Create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create our new servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	// Create our http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	//start our server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
