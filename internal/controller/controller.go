package controller

import "go-s3/internal/service"

type Controller struct {
	s3Service    service.S3Service
	imageService service.ImageService
}

func NewController(s3Service service.S3Service, imageService service.ImageService) *Controller {
	return &Controller{
		s3Service:    s3Service,
		imageService: imageService,
	}
}
