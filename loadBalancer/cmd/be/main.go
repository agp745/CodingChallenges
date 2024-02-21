package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var hostname, port string
	flag.StringVar(&hostname, "h", "127.0.0.1", "hostname")
	flag.StringVar(&port, "p", "70", "port")
	flag.Parse()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Printf("Recieved connection from %v\n"+
			"%v / %v\n"+
			"Host: %v\n"+
			"User-Agent: %v\n"+
			"Accept: %v\n\n"+
			"Replied with a hello message\n\n",
			c.Context().RemoteAddr(), c.Method(), c.Protocol(), c.Hostname(), string(c.Context().UserAgent()), c.Accepts())

		return c.SendString("Hello from server!\n")
	})

	app.Listen(fmt.Sprintf(":%s", port))
}
