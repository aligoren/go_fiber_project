package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {

	app.Post("/login", Login)

	app.Get("/users", GetUsers)

}
