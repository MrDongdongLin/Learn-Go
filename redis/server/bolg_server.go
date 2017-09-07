package server

import (
	mux "github.com/julienschmidt/httprouter"

	"LearnGo/redis/controller"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  mux.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controller.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/:todoId",
		controller.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controller.TodoCreate,
	},
	Route{
		"TodoDownload",
		"GET",
		"/todos.json",
		controller.TodoDownload,
	},
}
