package s3

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"
	"path/filepath"
)

func (s *Service) DownloadBatchCompress(ctx context.Context, fileNames []string) ([]byte, error) {
	l := logger.FromContext(ctx)
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	fileData := make(map[string][]byte)
	for _, fileName := range fileNames {
		if fileName == "" {
			continue
		}

		data, err := s.DownloadFile(ctx, fileName)
		if err != nil {
			l.Errorf("service.DownloadBatchCompress: failed to download file %s: %v", fileName, err)
			return nil, fmt.Errorf("service.DownloadBatchCompress: failed to download file %s: %w", fileName, err)
		}
		fileData[fileName] = data
	}

	compressedFiles, err := s.imageService.CompressBatch(ctx, fileData)
	if err != nil {
		l.Errorf("service.DownloadBatchCompress: failed to compress files: %v", err)
		return nil, fmt.Errorf("service.DownloadBatchCompress: failed to compress files: %w", err)
	}

	for fileName, data := range compressedFiles {
		f, err := zipWriter.Create(filepath.Base(fileName))
		if err != nil {
			l.Errorf("service.DownloadBatchCompress: failed to create zip entry for %s: %v", fileName, err)
			return nil, fmt.Errorf("service.DownloadBatchCompress: failed to create zip entry for %s: %w", fileName, err)
		}

		_, err = f.Write(data)
		if err != nil {
			l.Errorf("service.DownloadBatchCompress: failed to write zip entry for %s: %v", fileName, err)
			return nil, fmt.Errorf("service.DownloadBatchCompress: failed to write zip entry for %s: %w", fileName, err)
		}
	}

	err = zipWriter.Close()
	if err != nil {
		l.Errorf("service.DownloadBatchCompress: failed to close zip: %v", err)
		return nil, fmt.Errorf("service.DownloadBatchCompress: failed to close zip: %w", err)
	}

	return buf.Bytes(), nil
}
