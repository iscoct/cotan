package logger

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var Standard Logger

func readConfig(config_path string) {
	rootPath := os.Getenv("ROOT_PATH")

	viper.SetConfigName(config_path)
	viper.AddConfigPath(rootPath + "/cmd/pkg/config")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func init() {
	readConfig("basic_conf")

	ip := viper.GetString("server.ip")
	port := viper.GetInt("server.port")
	machine_format := "port: " + strconv.Itoa(port) + " - ip: " + ip

	info_logger := log.New(os.Stdout, "INFO - "+machine_format+": ", log.LstdFlags)
	warning_logger := log.New(os.Stdout, "WARNING - "+machine_format+": ", log.LstdFlags)
	error_logger := log.New(os.Stderr, "ERROR - "+machine_format+": ", log.LstdFlags)

	Standard = &Log{
		info: info_logger, warning: warning_logger, error_logger: error_logger,
	}
}
