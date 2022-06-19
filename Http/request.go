package http

import "net/http"

type Request struct {
	*http.Request
}
