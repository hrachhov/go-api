package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hwllo!!")
}

func creatRoutes(app *fiber.App) {
	app.Get("/api", hello)
}

func main() {
	app := fiber.New()
	creatRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
