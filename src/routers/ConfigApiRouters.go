package routers

import (
	"ToDoGolang/src/controllers"
	"net/http"
)

var ConfigRouters = []Router{
	{
		URI:           "/config",
		Method:        http.MethodGet,
		Function:      controllers.ApiConfigControllers,
		Authenticated: false,
	},
}
