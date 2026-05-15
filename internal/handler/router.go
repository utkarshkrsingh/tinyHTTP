package handler

import "github.com/utkarshkrsingh/tinyhttp/internal/http"

type Router struct {
	routes map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]Handler),
	}
}

func (r *Router) Register(method, path string, h Handler) {
	key := routeKey(method, path)
	r.routes[key] = h
}

func (r *Router) ServeHTTP(req *http.Request) *http.Response {
	key := routeKey(req.Method, req.Path)

	h, ok := r.routes[key]
	if !ok {
		return &http.Response{
			Version: "HTTP/1.1",
			StatusCode: 404,
			StatusText: "Not Found",
			Headers: map[string]string{
				"Content-Length": "0",
				"Connection": "close",
			},
		}
	}

	return h.ServeHTTP(req)
}

func routeKey(method, path string) string {
	return method + " " + path
}
