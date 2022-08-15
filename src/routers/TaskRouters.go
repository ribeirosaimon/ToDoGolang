package routers

import (
	"learnGo/src/controllers"
	"net/http"
)

var TaskRoutes = []Router{
	{
		URI:           "/tasks",
		Method:        http.MethodGet,
		Function:      controllers.GetAllTasks,
		Authenticated: false,
	},
	{
		URI:           "/task/:id",
		Method:        http.MethodGet,
		Function:      controllers.GetTask,
		Authenticated: false,
	},
	{
		URI:           "/task",
		Method:        http.MethodPost,
		Function:      controllers.PostTask,
		Authenticated: false,
	},
	{
		URI:           "/task/:id",
		Method:        http.MethodPut,
		Function:      controllers.PutTask,
		Authenticated: false,
	},
	{
		URI:           "/task/:id",
		Method:        http.MethodDelete,
		Function:      controllers.DeleteTask,
		Authenticated: false,
	},
}
