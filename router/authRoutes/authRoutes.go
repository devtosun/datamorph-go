package authroutes

import (
    "github.com/gofiber/fiber/v2"
    loginHandler "github.com/devtosun/datamorph-go/handler/authHandlers"
)

func SetupNoteRoutes(router fiber.Router) {
    auth := router.Group("/")
    // Login
    auth.Post("/login", loginHandler.Login)
    // // Read all Notes
    // note.Get("/", noteHandler.GetNotes)
    // // // Read one Note
    // note.Get("/:noteId", noteHandler.GetNote)
    // // // Update one Note
    // note.Put("/:noteId", noteHandler.UpdateNote)
    // // // Delete one Note
    // note.Delete("/:noteId", noteHandler.DeleteNote)
}