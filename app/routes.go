package app

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"Upload", "POST", "/upload", Upload},
		Route{"Search", "GET", "/search", Search},
	}
	return routes
}
