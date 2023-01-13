package logger

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Standard Logger

func init() {
	var ip string
	var port int

	err := godotenv.Load()

	if err != nil {
		ip = os.Getenv("IP")
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	} else {
		ip = "127.0.0.1"
		port = 8080
	}

	machine_format := "port: " + strconv.Itoa(port) + " - ip: " + ip

	info_logger := log.New(os.Stdout, "INFO - "+machine_format+": ", log.LstdFlags)
	warning_logger := log.New(os.Stdout, "WARNING - "+machine_format+": ", log.LstdFlags)
	error_logger := log.New(os.Stderr, "ERROR - "+machine_format+": ", log.LstdFlags)

	Standard = &Log{
		info: info_logger, warning: warning_logger, error_logger: error_logger,
	}
}
