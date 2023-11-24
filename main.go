package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/hello", helloHandler)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}

func helloHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
