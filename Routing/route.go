package routing

const VERBS = "GET|POST|PUT|DELETE|PATCH|HEAD|OPTIONS"

// declare a union type for action of type string or func
type Action interface {
	String() string
	Func() func()
}

type Route struct {
	uri    string
	method string
	action Action
	Router *Router
}

func NewRoute(uri string, method string, action Action, router *Router) *Route {
	return &Route{uri, method, action, router}
}

// get route uri
func (r *Route) getUri() string {
	return r.uri
}

// get route method
func (r *Route) getMethod() string {
	return r.method
}

// get route action
func (r *Route) getAction() Action {
	return r.action
}
