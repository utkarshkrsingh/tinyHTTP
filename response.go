package main

import (
	"fmt"
	"net"
	"net/http"
)

// generateResponse builds and sends a simple HTTP response
// based on the incoming request
func generateResponse(conn net.Conn, request *Request) {

	// Default response values (assume success)
	statusCode := http.StatusOK // 200
	statusText := "OK"
	body := "Hello"

	// If the requested path is not "/", return 404
	if request.Path != "/" {
		statusCode = http.StatusNotFound //404
		statusText = "Not Found"
		body = "Not Found"

		// If the method is not GET, return 405
	} else if request.Method != http.MethodGet {
		statusCode = http.StatusMethodNotAllowed // 405
		statusText = "Method Not Allowed"
		body = "Method Not Allowed"
	}

	// Build raw HTTP response string manually:
	// - Status line
	// - Headers
	// - Blank line
	// - Body
	reponse := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, statusText) +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) + // Required for HTTP/1.1
		"Content-Type: text/plain\r\n" + // Response content type
		"\r\n" + // End of headers
		body // Response body

	// Send the response back to the client
	conn.Write([]byte(reponse))
}
