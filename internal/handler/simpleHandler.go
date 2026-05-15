package handler

import (
	"strconv"

	"github.com/utkarshkrsingh/tinyhttp/internal/http"
)

type SimpleHandler struct {}

func(h *SimpleHandler) ServeHTTP(req *http.Request) *http.Response {
	body := []byte("Hello, World!")
	if req.Path == "/" {
		return &http.Response{
			Version: "HTTP/1.1",
			StatusCode: 200,
			StatusText: "OK",
			Headers: map[string]string{
				"Content-Length": strconv.Itoa(len(body)),
			},
			Body: body,
		}
	}

	return &http.Response{
		Version: "HTTP/1.1",
		StatusCode: 404,
		StatusText: "Not Found",
	}
}
