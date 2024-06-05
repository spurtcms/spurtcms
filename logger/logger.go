package logger

import (
	"fmt"
	"log"
	"os"
)

var LogFile = "logs/error.log"

const (
	InfoLevel = iota
	WarningLevel
	ErrorLevel
)

type Logger struct {
	Level       int
	InfoLogger  *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
}

var logger *Logger

func init() {

	makedir := os.Mkdir("logs", os.ModePerm)

	if makedir != nil {

		fmt.Println(makedir)

	}

	_, err := os.Stat(LogFile)

	var file *os.File

	if os.IsNotExist(err) {
		// Create the file since it does not exist
		_, err = os.Create(LogFile)

		if err != nil {
			fmt.Printf("Failed to create log file: %v", err)
		}

		fmt.Println("Created new log file.")

	}

	file, err = os.OpenFile(LogFile, os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {

		fmt.Printf("Failed to open log file: %v", err)
	}

	log.SetOutput(file)

	logger = &Logger{
		Level:       InfoLevel,
		InfoLogger:  log.New(file, "INFO ", log.LstdFlags|log.Lshortfile),
		WarnLogger:  log.New(file, "WARN ", log.LstdFlags|log.Lshortfile),
		ErrorLogger: log.New(file, "ERROR ", log.LstdFlags|log.Lshortfile),
	}
}

func ErrorLOG() *log.Logger {
	return logger.ErrorLogger
}

func WarnLOG() *log.Logger {
	return logger.WarnLogger
}
