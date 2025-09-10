package service

import (
	"context"
	"io"
)

type S3Service interface {
	UploadFile(ctx context.Context, fileName string, data []byte, contentType string) error
	ListFiles(ctx context.Context) ([]string, error)
	DownloadFile(ctx context.Context, fileName string) ([]byte, error)
	DeleteFile(ctx context.Context, fileName string) error
}

type ImageService interface {
	CompressImage(ctx context.Context, src io.Reader) ([]byte, error)
}
