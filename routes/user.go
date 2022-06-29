package routes

import (
	"tayson/blockchain/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUser(app *fiber.App) {
	app.Get("/api/users", controllers.GetUsers)
}
