package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Printf("Recieved connection from %v\n"+
			"%v / %v\n"+
			"Host: %v\n"+
			"User-Agent: %v\n"+
			"Accept: %v\n\n"+
			"Replied with a hello message\n\n",
			c.Context().RemoteAddr(), c.Method(), c.Protocol(), c.Hostname(), string(c.Context().UserAgent()), c.Accepts())

		return c.SendString("Hello from server on port 71")
	})

	app.Listen(":71")
}
