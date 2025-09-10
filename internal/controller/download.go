package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Download(ctx *fiber.Ctx) error {
	fileName := ctx.Params("name")
	if fileName == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Имя файла не указано"})
	}

	data, err := c.s3Service.DownloadFile(ctx.Context(), fileName)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Файл не найден"})
	}

	ctx.Set("Content-Type", "image/jpeg")
	ctx.Set("Content-Disposition", "attachment; filename="+fileName)

	return ctx.Send(data)
}
