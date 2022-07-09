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

// func (r *Router) Group(fn func(*Router)) {
// 	router := NewRouter()
// 	fn(router)
// 	r.addRoute(router.Routes[0])
// }

func (r *Router) updatePrefix(route *Route, prefix string) {
	route.uri = prefix + route.uri
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
		if route.GetMethod() == request.Method && route.GetUri() == request.URL.Path {
			r.DispatchRoute(route, writer, request)
		}
	}

	// http.NotFound(writer, request)
}

func (r *Router) DispatchRoute(route *Route, writer http.ResponseWriter, request *http.Request) {
	r.handleFunction(route, writer, request)
}

// func (r *Router) handleString(route *Route, writer http.ResponseWriter, request *http.Request) {
// 	parts := strings.Split(route.GetAction().(string), "\\")
// 	directory := strings.Join(parts[:len(parts)-1], "\\")
// 	controller := strings.Split(parts[len(parts)-1], "@")[0]
// 	action := strings.Split(parts[len(parts)-1], "@")[1]

// 	if _, err := os.Stat(directory); os.IsNotExist(err) {
// 		log.Println("Directory does not exist", directory)
// 		return
// 	}

// 	// make sure the controller exists
// 	if _, err := os.Stat(directory + "\\" + controller + ".go"); os.IsNotExist(err) {
// 		log.Println("Controller does not exist", directory+"\\"+controller+".go")
// 		return
// 	}

// 	// instantiate the controller
// 	controllerInstance := NewController(directory, controller)
// }

func (r *Router) handleFunction(route *Route, writer http.ResponseWriter, request *http.Request) {
	route.GetAction().(func(http.ResponseWriter, *http.Request))(writer, request)
}
