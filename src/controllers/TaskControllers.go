package controllers

import (
	"ToDoGolang/src/payloads"
	"ToDoGolang/src/response"
	"ToDoGolang/src/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetTask(ctx *fiber.Ctx) error {
	id := fmt.Sprint(ctx.Params("id"))
	responseTask := services.GetTask(id)
	return ctx.Status(responseTask.Code).JSON(responseTask.Data)
}

func GetAllTasks(ctx *fiber.Ctx) error {
	responseTasks := services.GetAllTasks()
	return ctx.Status(responseTasks.Code).JSON(responseTasks.Data)
}

func PostTask(ctx *fiber.Ctx) error {
	taskDto := payloads.TaskDTO{}
	if err := ctx.BodyParser(&taskDto); err != nil {
		return ctx.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "error in body parse"})
	}
	responseTasks := services.SaveTask(taskDto)
	return ctx.Status(responseTasks.Code).JSON(responseTasks.Data)
}

func PutTask(ctx *fiber.Ctx) error {
	id := fmt.Sprint(ctx.Params("id"))
	if id == "" {
		return ctx.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "error in body parse"})
	}
	taskDto := payloads.TaskDTO{}
	if err := ctx.BodyParser(&taskDto); err != nil {
		return ctx.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "error in body parse"})
	}
	responseTasks := services.EditTask(id, taskDto)
	return ctx.Status(responseTasks.Code).JSON(responseTasks.Data)
}

func DeleteTask(ctx *fiber.Ctx) error {
	id := fmt.Sprint(ctx.Params("id"))
	if id == "" {
		return ctx.Status(http.StatusConflict).JSON(response.ErrorResponse{Message: "error in body parse"})
	}
	responseTasks := services.DeleteTask(id)
	return ctx.SendStatus(responseTasks.Code)
}
