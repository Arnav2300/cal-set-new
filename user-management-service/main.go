package main

import (
	"user-management-service/config"
	"user-management-service/server"
)

func main() {
	config.PrintLogo()
	server.StartServer(config.Port)
}
