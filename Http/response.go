package http

import "net/http"

type Response struct {
	http.ResponseWriter
}
