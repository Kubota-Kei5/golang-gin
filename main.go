package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/logger"
)

func main() {
	router := gin.Default()

	router.Use(ginzap.Ginzap(logger.ZapLogger, time.RFC3339, true))

	router.Use(ginzap.RecoveryWithZap(logger.ZapLogger, true))

	router.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/panic", func(c *gin.Context){
		panic("An unexpected error occurred!")
	})

	srv := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.ZapLogger.Fatal("Failed to start server", zap.Error(err))
	}

	go func() {
		logger.Info("start server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()
	
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	defer logger.Sync()
	logger.Info("stop")
}