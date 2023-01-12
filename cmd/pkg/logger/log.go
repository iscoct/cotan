package logger

import (
	"log"
	"os"
)

type Log struct {
	info         *log.Logger
	warning      *log.Logger
	error_logger *log.Logger
}

func (log *Log) Info(args ...interface{}) {
	log.info.Println(args...)
}

func (log *Log) Warning(args ...interface{}) {
	log.warning.Println(args...)
}

func (log *Log) Error(args ...interface{}) {
	log.error_logger.Println(args...)
}

func (log *Log) Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func NewLoggerWithDefaultOutputs() Logger {
	info_logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	warning_logger := log.New(os.Stdout, "WARNING: ", log.LstdFlags)
	error_logger := log.New(os.Stderr, "ERROR: ", log.LstdFlags)

	return &Log{
		info: info_logger, warning: warning_logger, error_logger: error_logger,
	}
}
