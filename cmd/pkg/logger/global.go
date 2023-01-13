package logger

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var Standard Logger

func readConfig(config_path string) error {
	rootPath := os.Getenv("ROOT_PATH")

	viper.SetConfigName(config_path)
	viper.AddConfigPath(rootPath + "/cmd/pkg/config")
	err := viper.ReadInConfig()

	return err
}

func init() {
	var ip string
	var port int

	err := readConfig("basic_conf")

	if err != nil {
		ip = viper.GetString("server.ip")
		port = viper.GetInt("server.port")
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
