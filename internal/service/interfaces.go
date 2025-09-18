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
	UploadBatch(ctx context.Context, files map[string][]byte, contentType string) error
	UploadBatchCompress(ctx context.Context, files map[string][]byte, contentType string) error
	DownloadBatch(ctx context.Context, fileNames []string) ([]byte, error)
	DownloadBatchCompress(ctx context.Context, fileNames []string) ([]byte, error)
}

type ImageService interface {
	CompressImage(ctx context.Context, src io.Reader) ([]byte, error)
	CompressBatch(ctx context.Context, files map[string][]byte) (map[string][]byte, error)
}
