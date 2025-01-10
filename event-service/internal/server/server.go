package server

import (
	"event-service/internal/config"
	"event-service/internal/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Serve() {
	config.PrintLogo()
	if logger.Log == nil {
		panic("Logger is not initialized!")
	}
	router := gin.Default()
	logger.Log.Infof("Server started on port %s", config.Port)
	if err := router.Run(config.Port); err != nil {
		// logger.Log.Fatalf("Failed to start server: %v", err)
		fmt.Print(err)
	}
}
