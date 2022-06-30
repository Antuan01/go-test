package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Antuan01/go-test/controllers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Post("/report/post", controllers.CreatePostReport)

	api.Post("/report/comment", controllers.CreateCommentReport)

}