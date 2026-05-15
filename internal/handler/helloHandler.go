package handler

import (
	"strconv"

	"github.com/utkarshkrsingh/tinyhttp/internal/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(req *http.Request) *http.Response {
	body := []byte("Hello!")

	return &http.Response{
		Version:    "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Length": strconv.Itoa(len(body)),
			"Connection":     "close",
		},
		Body: body,
	}
}
