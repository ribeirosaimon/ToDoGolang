package services

import (
	"learnGo/src/payloads"
	"learnGo/src/repository"
	"learnGo/src/response"
	"net/http"
)

const errorMessage = "Task not Found"

func GetTask(id string) response.HttpResponse {
	task, err := repository.FindTask(id)
	if err != nil {
		return response.HttpResponse{
			Code: http.StatusConflict,
			Data: response.ErrorResponse{Message: errorMessage},
		}
	}
	return response.HttpResponse{
		Code: http.StatusOK,
		Data: task,
	}

}

func GetAllTasks() response.HttpResponse {
	tasks, err := repository.FindAllTasks()
	if err != nil {
		return response.HttpResponse{
			Code: http.StatusConflict,
			Data: response.ErrorResponse{Message: errorMessage},
		}
	}
	return response.HttpResponse{
		Code: http.StatusOK,
		Data: tasks,
	}

}

func SaveTask(dto payloads.TaskDTO) response.HttpResponse {
	task, err := repository.SaveTask(dto)
	if err != nil {
		return response.HttpResponse{
			Code: http.StatusConflict,
			Data: response.ErrorResponse{Message: errorMessage},
		}
	}
	return response.HttpResponse{
		Code: http.StatusOK,
		Data: task,
	}
}

func EditTask(id string, dto payloads.TaskDTO) response.HttpResponse {
	task, err := repository.EditTask(id, dto)
	if err != nil {
		return response.HttpResponse{
			Code: http.StatusConflict,
			Data: response.ErrorResponse{Message: errorMessage},
		}
	}
	return response.HttpResponse{
		Code: http.StatusOK,
		Data: task,
	}
}

func DeleteTask(id string) response.HttpResponse {
	isDeleted, err := repository.DeleteTask(id)
	if err != nil {
		return response.HttpResponse{
			Code: http.StatusConflict,
			Data: response.ErrorResponse{Message: errorMessage},
		}
	}
	return response.HttpResponse{
		Code: http.StatusOK,
		Data: isDeleted,
	}
}
