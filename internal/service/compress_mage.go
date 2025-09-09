package service

import (
	"bytes"
	"context"
	"fmt"
	"go-s3/internal/infrastruct/logger"
	"image"
	"image/jpeg"
	"io"

	"golang.org/x/image/draw"
)

func (i *ImageService) CompressImage(ctx context.Context, src io.Reader) ([]byte, error) {
	l := logger.FromContext(ctx)

	img, _, err := image.Decode(src)
	if err != nil {
		l.Errorf("service.CompressImage: failed to decode image: %v", err)
		return nil, fmt.Errorf("service.CompressImage: failed to decode image: %w", err)
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if width <= i.maxWidth {
		newHeight := height * i.maxWidth / width
		newImg := image.NewRGBA(image.Rect(0, 0, i.maxWidth, newHeight))
		draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, bounds, draw.Over, nil)
		img = newImg
	}

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: i.quality})
	if err != nil {
		l.Errorf("service.CompressImage: failed to encode image: %v", err)
		return nil, fmt.Errorf("service.CompressImage: failed to encode image: %w", err)
	}

	return buf.Bytes(), nil

}
