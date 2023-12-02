package logger

import (
	"log"
	"os"
)

type Log struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewLogger() *Log {
	file, err := os.OpenFile("../../log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	infoLog := log.New(file, "INFO: \t", log.Ldate|log.Ltime)
	errorLog := log.New(file, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile)
	
	infoLog.Println("Logger initialized successfully!")
	errorLog.Println("Logger initialized successfully!")

	return &Log{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
