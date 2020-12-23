package main

import (
	"dam-webhook/application"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

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

	app := &application.Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  routes(app),
	}

	infoLog.Printf("Starting server on %s", addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
