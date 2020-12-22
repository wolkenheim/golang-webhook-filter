package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func initConfig() {
	fmt.Printf("%v", os.Getenv("APP_ENV"))
	if os.Getenv("APP_ENV") == "production" || os.Getenv("APP_ENV") == "staging" {
		readConfig(os.Getenv("APP_ENV"))
	} else {
		readConfig("local")
	}
}

func main() {
	initConfig()
	addr := viper.GetString("server.port")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
