package main

import (
	"github.com/azkaainurridho514/api_notes_app/database"
	"github.com/azkaainurridho514/api_notes_app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)
func main() {
	database.Connect()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	router.SetupRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
	return c.SendStatus(404) 
	})
 	// app.Listen(":8080")
 	app.Listen("192.168.126.114:8080")
}