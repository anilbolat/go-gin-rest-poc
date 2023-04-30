package controller

import (
	"github.com/anilbolat/go-gin-rest-poc/entity"
	"github.com/anilbolat/go-gin-rest-poc/service"
	"github.com/gin-gonic/gin"
	"log"
)

type VideoControllerInterface interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type videoController struct {
	videoService service.VideoServiceInterface
}

func NewVideoController(videoService service.VideoServiceInterface) VideoControllerInterface {
	return &videoController{
		videoService: videoService,
	}
}

func (videoController *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		log.Panic("cannot get video from the request. ", err)
	}

	videoController.videoService.Save(video)
	return video
}

func (videoController *videoController) FindAll() []entity.Video {
	return videoController.videoService.FindAll()
}
