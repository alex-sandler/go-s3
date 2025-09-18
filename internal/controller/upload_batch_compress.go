package controller

import (
	"io"

	"github.com/gofiber/fiber/v2"
)

func (c *Controller) UploadBatchCompressed(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Ошибка получения формы"})
	}

	files := form.File["files"]
	if len(files) == 0 {
		return ctx.Status(400).JSON(fiber.Map{"error": "Файлы не найдены"})
	}

	fileData := make(map[string][]byte)
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка открытия файла: " + file.Filename})
		}
		defer src.Close()

		data, err := io.ReadAll(src)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка чтения файла: " + file.Filename})
		}

		fileData[file.Filename] = data
	}

	err = c.s3Service.UploadBatchCompress(ctx.Context(), fileData, "image/jpeg")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка загрузки файлов в S3"})
	}

	return ctx.JSON(fiber.Map{"message": "Файлы загружены"})
}
