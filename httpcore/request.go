package httpcore

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/utkarshkrsingh/tinyHTTP/protocol"
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
)

// readRequest read data from a TCP connection and parse it into HTTP requesr.
// It incrementally reads from the connection because data may arrive in chunks.
func ReadRequest(conn net.Conn) (*Request, error) {
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
			return nil, fmt.Errorf("unable to read: %s", err)
		}

		// Append the newly read bytes to the accumulated message buffer
		msg.Write(buff[:n])
		data := msg.String()

		// Parse headers once we detect the header-body separated
		if !headerParsed {
			parts := strings.SplitN(data, protocol.DBL_CRLF, 2)
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
		parts := strings.SplitN(data, protocol.DBL_CRLF, 2)
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

	return req, nil
}
