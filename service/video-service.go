package service

import "github.com/anilbolat/go-gin-rest-poc/entity"

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (service *videoService) Save(newVideo entity.Video) entity.Video {
	service.videos = append(service.videos, newVideo)
	return newVideo
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
