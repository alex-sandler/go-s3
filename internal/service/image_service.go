package service

type ImageService struct {
	quality  int
	maxWidth int
}

func NewImageService(quality int, maxWidth int) *ImageService {
	return &ImageService{
		quality:  quality,
		maxWidth: maxWidth,
	}
}
