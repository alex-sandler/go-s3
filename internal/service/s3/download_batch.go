package s3

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"
	"path/filepath"
)

func (s *Service) DownloadBatch(ctx context.Context, fileNames []string) ([]byte, error) {
	l := logger.FromContext(ctx)
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	for _, fileName := range fileNames {
		if fileName == "" {
			continue
		}

		data, err := s.DownloadFile(ctx, fileName)
		if err != nil {
			l.Errorf("service.DownloadBatch: failed to download file %s: %v", fileName, err)
			return nil, fmt.Errorf("service.DownloadBatch: failed to download file %s: %w", fileName, err)
		}

		f, err := zipWriter.Create(filepath.Base(fileName))
		if err != nil {
			l.Errorf("service.DownloadBatch: failed to create zip entry for %s: %v", fileName, err)
			return nil, fmt.Errorf("service.DownloadBatch: failed to create zip entry for %s: %w", fileName, err)
		}

		_, err = f.Write(data)
		if err != nil {
			l.Errorf("service.DownloadBatch: failed to write zip entry for %s: %v", fileName, err)
			return nil, fmt.Errorf("service.DownloadBatch: failed to write zip entry for %s: %w", fileName, err)
		}
	}

	err := zipWriter.Close()
	if err != nil {
		l.Errorf("service.DownloadBatch: failed to close zip: %v", err)
		return nil, fmt.Errorf("service.DownloadBatch: failed to close zip: %w", err)
	}

	return buf.Bytes(), nil
}
