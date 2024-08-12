package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello!!")
}

func check(c *fiber.Ctx) error {
	return c.SendString("Runing!!")
}

func creatRoutes(app *fiber.App) {
	app.Get("/api", hello)
	app.Get("/api/check", check)
}

func main() {
	app := fiber.New()
	creatRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
