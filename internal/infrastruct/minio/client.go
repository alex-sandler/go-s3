package minio

import (
	"context"
	"fmt"
	"go-s3/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	mc *minio.Client
}

func NewMinioClient() *Client {
	return &Client{}
}

func (m *Client) GetClient() *minio.Client {
	return m.mc
}

func (m *Client) Init(ctx context.Context, cfg *config.Config) error {
	client, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.RootUser, cfg.Minio.RootPassword, ""),
		Secure: cfg.Minio.UseSSL,
	})

	if err != nil {
		return fmt.Errorf("minio.Init: failed to create minio client: %w", err)
	}

	m.mc = client

	exists, err := m.mc.BucketExists(ctx, cfg.Minio.BucketName)
	if err != nil {
		return fmt.Errorf("minio.Init: failed to check if bucket exists: %w", err)
	}
	if !exists {
		err = m.mc.MakeBucket(ctx, cfg.Minio.BucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("minio.Init: failed to create bucket: %w", err)
		}
	}

	return nil
}
