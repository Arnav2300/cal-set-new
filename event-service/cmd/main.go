package main

import (
	"event-service/internal/logger"
	"event-service/internal/server"
)

func init() {
	logger.InitLogger()
}

func main() {
	server.Serve()
}
