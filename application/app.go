package application

import "log"

// Application defines app
type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
