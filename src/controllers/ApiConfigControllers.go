package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type configDto struct {
	Running bool      `json:"running"`
	Name    string    `json:"apiName"`
	TimeNow time.Time `json:"timeNow"`
}

func ApiConfigControllers(ctx *fiber.Ctx) error {
	config := configDto{
		Name:    "To Do Api",
		Running: true,
		TimeNow: time.Now(),
	}
	return ctx.Status(http.StatusOK).JSON(config)
}
