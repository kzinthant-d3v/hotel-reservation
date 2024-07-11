package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/kzinthant-d3v/hotel-reservation/api"
)

func main() {
	port := flag.String("port", ":5000", "The port of api");
	flag.Parse()

  app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleListUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*port)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "hello foo"})
}