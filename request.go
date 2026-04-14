package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// Request represents a parsed HTTP request
type Request struct {
	Method  string
	Path    string
	Version string
	Headers map[string][]string
	Body    BodyData
}

// BodyData represents the raw HTTP request body.
// It is stored as bytes because the body may not be text (e.g., binary data).
type BodyData []byte

var (
	// bufferSize defines how many bytes are read per socket read operation
	bufferSize = 1024

	// CRLF is thr standard HTTP line separator ("\r\n")
	CRLF = "\r\n"

	// DBL_CRLF separates HTTP headers from the body ("\r\n\r\n")
	DBL_CRLF = "\r\n\r\n"
)

// handleClient reads data from a TCP connection and parses it into an HTTP request.
// It incrementally reads from the connection because data may arrive in chunks.
func handleClient(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, bufferSize)
	var msg strings.Builder

	var req *Request
	var contentLength int

	// Tracks whether header have fully parsed
	headerParsed := false

	for {
		// Read incoming data from the connection
		n, err := conn.Read(buff)
		if err != nil {
			log.Printf("error : %v\n", err)
			return
		}

		// Append the newly read bytes to the accumulated message buffer
		msg.Write(buff[:n])
		data := msg.String()

		// Parse headers once we detect the header-body separated
		if !headerParsed {
			parts := strings.SplitN(data, DBL_CRLF, 2)
			if len(parts) < 2 {
				// Header not fully received yet
				continue
			}

			metaData := parts[0]
			req = parseMetaData(metaData)

			// Extract Content-Length header if present to know body size
			for key, val := range req.Headers {
				if strings.EqualFold(key, "Content-Length") && len(val) > 0 {
					num, err := strconv.Atoi(val[0])
					if err == nil {
						contentLength = num
					}
				}
			}
			headerParsed = true
		}

		// Extract body if headers are already parsed
		parts := strings.SplitN(data, DBL_CRLF, 2)
		if len(parts) < 2 {
			continue
		}

		bodyData := parts[1]

		// Wait until full body is received based on Content-Length
		if len(bodyData) >= contentLength {
			req.Body = []byte(BodyData(bodyData[:contentLength]))
			log.Println(bodyData)
			break
		}
	}

	// generateResponse sends response back to the client
	generateResponse(conn, req)
}

// parseMetaData parses the HTTP request line and headers from the raw text
func parseMetaData(metaData string) *Request {
	request := &Request{}

	// Split header into individual lines
	lines := strings.Split(metaData, CRLF)
	if len(lines) == 0 {
		return request
	}

	// Parse the request line: METHOD PATH VERSION
	startLine := strings.Fields(lines[0])

	// Ensure the request line have at least 3 components
	if len(startLine) >= 3 {
		request.Method = startLine[0]
		request.Path = startLine[1]
		request.Version = startLine[2]
	}

	headers := make(map[string][]string)

	// Parse each header line
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)

		// Skip empty lines
		if line == "" {
			continue
		}

		// Split header into key and value (only at first colon)
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Some header can have multiple comma-seperated values
		values := strings.Split(value, ",")

		headers[key] = values
	}

	request.Headers = headers
	return request
}
