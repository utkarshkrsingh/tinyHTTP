package http

type Response struct {
	Version    string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
}
