package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Routes(app *fiber.App) {
	app.Get("/health", c.Health)
	app.Post("/upload", c.Upload)
	app.Get("/files/:name", c.Download)
	app.Get("/files", c.List)
	app.Delete("/files/:name", c.Delete)
}
