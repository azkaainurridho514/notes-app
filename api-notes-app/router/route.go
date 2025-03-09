package router

import (
	"github.com/azkaainurridho514/api_notes_app/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// auth =======================================================
	auth := api.Group("/auth")
	auth.Get("/login", handler.Login)
	auth.Post("/register", handler.Register)
	// note ====================================================
	note := api.Group("/notes")
	note.Get("/all", handler.GetAllNotes)
	note.Post("/create", handler.CreateNote)
	note.Put("/update", handler.UpdateNote)
	note.Delete("/delete", handler.DeleteNote)
}