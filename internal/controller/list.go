package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) List(ctx *fiber.Ctx) error {
	files, err := c.s3Service.ListFiles(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка получения списка файлов"})
	}

	return ctx.JSON(fiber.Map{
		"files": files,
		"count": len(files),
	})
}
