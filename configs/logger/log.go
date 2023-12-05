package logger

import (
	"log"
	"os"
)

type Log struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger

	File *os.File
}

func NewLogger(path string) *Log {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	infoLog := log.New(file, "INFO: \t", log.Ldate|log.Ltime)
	errorLog := log.New(file, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)

	return &Log{
		InfoLog:  infoLog,
		ErrorLog: errorLog,

		File: file,
	}
}
