package main

import (
	"github.com/gofiber/fiber/v2"
	"learnGo/src/routers"
)

func main() {

	app := fiber.New()
	routers.ConfigureRouters(app)
	app.Listen(":3000")
}
