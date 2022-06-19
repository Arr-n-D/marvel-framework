package routing

const VERBS = "GET|POST|PUT|DELETE|PATCH|HEAD|OPTIONS"

// declare a union type for action of type string or func
type Route struct {
	uri    string
	method string
	action interface{}
	Router *Router
}

func NewRoute(uri string, method string, action interface{}, router *Router) *Route {
	return &Route{uri, method, action, router}
}

// get route uri
func (r *Route) GetUri() string {
	return r.uri
}

// get route method
func (r *Route) GetMethod() string {
	return r.method
}

func (r *Route) GetAction() interface{} {
	return r.action
}

func (r *Route) setRouter(router *Router) {
	r.Router = router
}
