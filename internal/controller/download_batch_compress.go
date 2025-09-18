package controller

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) DownloadBatchCompress(ctx *fiber.Ctx) error {
	fileNames := strings.Split(ctx.Query("names"), ",")
	if len(fileNames) == 0 {
		return ctx.Status(400).JSON(fiber.Map{"error": "Имена файлов не указаны"})
	}

	data, err := c.s3Service.DownloadBatchCompress(ctx.Context(), fileNames)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка скачивания файлов"})
	}

	ctx.Set("Content-Type", "application/zip")
	ctx.Set("Content-Disposition", "attachment; filename=images_compressed.zip")

	return ctx.Send(data)
}
