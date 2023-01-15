package logger

import (
	"log"
	"os"
	"strconv"

	"cotan/cmd/pkg/config"
)

var Standard Logger

func init() {
	confReader := config.GetConfReader()

	machine_format := "port: " + strconv.Itoa(confReader.Port) + " - ip: " + confReader.IP

	info_logger := log.New(os.Stdout, "INFO - "+machine_format+": ", log.LstdFlags)
	warning_logger := log.New(os.Stdout, "WARNING - "+machine_format+": ", log.LstdFlags)
	error_logger := log.New(os.Stderr, "ERROR - "+machine_format+": ", log.LstdFlags)

	Standard = &Log{
		info: info_logger, warning: warning_logger, error_logger: error_logger,
	}
}
