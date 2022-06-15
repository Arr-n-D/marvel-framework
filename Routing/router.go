package routing

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
	switch handler.(type) {
	case string:
		r.addRoute(&Route{
			uri:    path,
			method: "POST",
			action: handler.(string),
			Router: r,
		})
	// case func
	case func():
		r.addRoute(&Route{
			uri:    path,
			method: "POST",
			action: handler.(func()),
			Router: r,
		})
	}
}

// func get
func (r *Router) Get(path string, handler interface{}) {
	// check type of handler
	switch handler.(type) {
	case string:
		r.addRoute(&Route{
			uri:    path,
			method: "GET",
			action: handler.(string),
			Router: r,
		})
	// case func
	case func():
		r.addRoute(&Route{
			uri:    path,
			method: "GET",
			action: handler.(func()),
			Router: r,
		})
	}
}
