package controller

import "github.com/gofiber/fiber/v2"

func (c *Controller) Routes(app *fiber.App) {
	app.Get("/health", c.Health)
	app.Post("/upload", c.Upload)
	app.Post("/upload-batch", c.UploadBatch)
	app.Post("/upload-batch-compress", c.UploadBatchCompressed)
	app.Get("/files/:name", c.Download)
	app.Get("/files", c.List)
	app.Get("/download-batch", c.DownloadBatch)
	app.Get("/download-batch-compress", c.DownloadBatchCompress)
	app.Delete("/files/:name", c.Delete)
}
