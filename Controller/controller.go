package controller

import (
	http "github.com/Arr-n-D/marvel/Http"
	routing "github.com/Arr-n-D/marvel/Routing"

	app "github.com/Arr-n-D/marvel/App"
)

type Controller struct {
	request  *http.Request
	response *http.Response
	router   *routing.Router
	app      *app.Application
}
