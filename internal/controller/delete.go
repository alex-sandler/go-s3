package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	fileName := ctx.Params("name")
	if fileName == "" {
		return ctx.Status(400).JSON(fiber.Map{"error": "Имя файла не указано"})
	}

	if err := c.s3Service.DeleteFile(ctx.Context(), fileName); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": "Ошибка удаления файла"})
	}

	return ctx.JSON(fiber.Map{"message": "Файл удалён"})
}
