package main

import (
	"github.com/anilbolat/go-gin-rest-poc/controller"
	"github.com/anilbolat/go-gin-rest-poc/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var videoController controller.VideoController = controller.NewVideoController(service.NewVideoService())

func main() {
	server := gin.Default()
	server.GET("/test", func(context *gin.Context) {

		context.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})

	})

	server.GET("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(context *gin.Context) {
		context.JSON(http.StatusOK, videoController.Save(context))
	})

	err := server.Run(":8080")
	if err != nil {
		return
	}
}
