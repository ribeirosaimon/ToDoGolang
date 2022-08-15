package routers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	URI           string
	Method        string
	Function      func(ctx *fiber.Ctx) error
	Authenticated bool
}

func ConfigureRouters(app *fiber.App) *fiber.App {
	tasksRoutes := TaskRoutes
	configRouter := ConfigRouters

	for index, route := range tasksRoutes {
		fmt.Printf("Task Router number: %d successful add\n", index)
		app.Add(route.Method, route.URI, route.Function)
	}
	for index, route := range configRouter {
		fmt.Printf("Config Router number: %d successful add\n", index)
		app.Add(route.Method, route.URI, route.Function)
	}
	return app
}
