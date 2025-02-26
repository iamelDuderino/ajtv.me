package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		Info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		Error: log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Llongfile),
	}
}
