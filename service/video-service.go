package service

import (
	"gitlab.com/voideus/go-rest-api/entity"
	"gitlab.com/voideus/go-rest-api/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	Update(entity.Video)
	Delete(entity.Video)
}

type videoService struct {
	videoRepository repository.VideoRepository
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	// service.videos = append(service.videos, video)
	service.videoRepository.Save(video)
	return video

}

func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

func (service *videoService) Update(video entity.Video) {
	service.videoRepository.Update(video)
}

func (service *videoService) Delete(video entity.Video) {
	service.videoRepository.Delete(video)
}
