package image

import (
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"
)

func (s *Service) CompressBatch(ctx context.Context, files map[string][]byte) (map[string][]byte, error) {
	l := logger.FromContext(ctx)
	compressedFiles := make(map[string][]byte)

	for fileName, data := range files {
		compressedData, err := s.CompressImage(ctx, bytes.NewReader(data))
		if err != nil {
			l.Errorf("service.CompressBatch: failed to compress file %s: %v", fileName, err)
			return nil, fmt.Errorf("service.CompressBatch: failed to compress file %s: %w", fileName, err)
		}
		compressedFiles[fileName] = compressedData
	}

	return compressedFiles, nil
}
