package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Файл не найден"})
	}

	src, err := file.Open()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка открытия файла"})
	}
	defer src.Close()

	data, err := c.imageService.CompressImage(ctx.Context(), src)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка сжатия изображения"})
	}

	err = c.s3Service.UploadFile(ctx.Context(), file.Filename, data, "image/jpeg")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка загрузки файла в S3"})
	}

	return ctx.JSON(fiber.Map{"message": "Файл загружен"})
}
