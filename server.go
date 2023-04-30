package main

import (
	"github.com/anilbolat/go-gin-rest-poc/controller"
	"github.com/anilbolat/go-gin-rest-poc/middlewares"
	"github.com/anilbolat/go-gin-rest-poc/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"io"
	"net/http"
	"os"
)

var vc controller.VideoControllerInterface = controller.NewVideoController(service.NewVideoService())

func main() {
	// setupLogOutput()  // to write logs into a file.

	// server := gin.Default() // this has default middlewares.

	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/test", func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})

	})
	server.GET("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, vc.FindAll())
	})
	server.POST("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, vc.Save(context))
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func setupLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}
