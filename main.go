package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)
import "github.com/aligoren/fiber_project/routes"

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
