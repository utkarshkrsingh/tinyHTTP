package handler

import "github.com/utkarshkrsingh/tinyhttp/internal/http"

type Handler interface {
	ServeHTTP(req *http.Request) *http.Response
}
