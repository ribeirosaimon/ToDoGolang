package main

import (
	"ToDoGolang/src/routers"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.ConfigureRouters(app)
	err := app.Listen(":4000")
	if err != nil {
		fmt.Println(fmt.Sprint(err))
	}
}
