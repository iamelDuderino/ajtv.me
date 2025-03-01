package logger

import (
	"log"
	"os"
)

type Logger struct {
	Info  *log.Logger
	Error *log.Logger
}

func NewLogger(info, err string) *Logger {
	return &Logger{
		Info:  log.New(os.Stdout, info+": ", log.LstdFlags),
		Error: log.New(os.Stderr, err+": ", log.LstdFlags|log.Llongfile),
	}
}
