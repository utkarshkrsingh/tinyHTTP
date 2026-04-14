package httpcore

import (
	"fmt"

	"github.com/utkarshkrsingh/tinyHTTP/protocol"
)

// generateResponse builds and sends a simple HTTP response
// based on the incoming request
func GenerateResponse(request *Request, err error) string {

	// Default response values (assume success)
	statusCode := protocol.StatusOk // 200
	statusText := "OK"
	headers := ""
	body := "Hello"

	// If there is any error while parsing the request, return, 400
	if err != nil {
		statusCode = protocol.StatusBadRequest // 400
		statusText = "Bad Request"
		body = "Bad Request"

		// If the requested path is not "/", return 404
	} else if request.Path != "/" {
		statusCode = protocol.StatusNotFound //404
		statusText = "Not Found"
		body = "Not Found"

		// If the method is not GET, return 405
	} else if request.Method != protocol.MethodGet {
		statusCode = protocol.StatusMethodNotAllowed // 405
		statusText = "Method Not Allowed"
		headers += "Allow: GET\r\n"
		body = "Method Not Allowed"
	}

	// Build raw HTTP response string manually:
	// - Status line
	// - Headers
	// - Blank line
	// - Body
	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, statusText) +
		fmt.Sprintf("Content-Length: %d\r\n", len(body)) + // Required for HTTP/1.1
		"Content-Type: text/plain\r\n" + // Response content type
		headers + // Adds additional header
		"\r\n" + // End of headers
		body // Response body

	return response
}
