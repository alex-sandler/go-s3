package main

import (
	"context"
	"go-s3/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	if err := app.Run(ctx); err != nil {
		log.Fatalf("failed to start service: %v", err)
	}
}
