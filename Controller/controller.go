package controller

type Controller struct {
	request  *Request
	response *Response
	router   *Router
	app      *App
}
