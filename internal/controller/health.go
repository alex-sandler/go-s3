package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Health(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"service": "s3-service",
		"message": "Service is running",
	})
}
