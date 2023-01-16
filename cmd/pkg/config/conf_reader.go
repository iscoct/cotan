package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type confReader struct {
	IP   string
	Port int
}

var instance *confReader

func createConfReader() {
	ip := "127.0.0.1"
	port := 8080

	err := godotenv.Load()

	if err != nil {
		ip = os.Getenv("IP")
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}

	instance = &confReader{Port: port, IP: ip}
}

func GetConfReader() *confReader {
	if instance == nil {
		createConfReader()
	}

	return instance
}
