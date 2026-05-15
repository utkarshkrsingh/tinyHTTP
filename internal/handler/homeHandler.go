package handler

import (
	"strconv"

	"github.com/utkarshkrsingh/tinyhttp/internal/http"
)

type HomeHandler struct {}

func (h *HomeHandler) ServeHTTP(req *http.Request) *http.Response {
	body := []byte("Welcome Home")

	return &http.Response{
		Version: "HTTP/1.1",
		StatusCode: 200,
		StatusText: "OK",
		Headers: map[string]string{
			"Content-Length": strconv.Itoa(len(body)),
			"Connection": "close",
		},
		Body: body,
	}
}
