package main

import (
	"os"
)

const (
	defaultDataDir = "/opt/mock-otp-server/data/"
	defaultHTTPPort = "8085"
)

func GetConfig() (string, string) {
	dataDir := os.Getenv("DATA_DIR")
	if dataDir == "" {
		dataDir = defaultDataDir
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = defaultHTTPPort
	}

	return dataDir, httpPort
}
