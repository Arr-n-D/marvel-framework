package routing

import (
	"net/http"
)

type Router struct {
	Routes []*Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) addRoute(route *Route) {
	r.Routes = append(r.Routes, route)
}

// our handler can be either a string or a function
func (r *Router) Post(path string, handler interface{}) {
	// check type of handler
	r.addRoute(&Route{
		uri:    path,
		method: "POST",
		action: handler,
		Router: r,
	})

}

// func get
func (r *Router) Get(path string, handler interface{}) {

	r.addRoute(&Route{
		uri:    path,
		method: "GET",
		action: handler,
		Router: r,
	})

}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// regexp.MustCompile(`\{(\d+)\}`)
	for _, route := range r.Routes {
		if route.GetMethod() == request.Method {
			r.DispatchRoute(route, writer, request)
		}
	}

	// http.NotFound(writer, request)
}

func (r *Router) DispatchRoute(route *Route, writer http.ResponseWriter, request *http.Request) {
	// check type of handler
	if handler, ok := route.GetAction().(func(http.ResponseWriter, *http.Request)); ok {
		handler(writer, request)
	}
}
