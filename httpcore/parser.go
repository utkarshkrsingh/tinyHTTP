package httpcore

import (
	"strings"

	"github.com/utkarshkrsingh/tinyHTTP/protocol"
)

// parseMetaData parses the HTTP request line and headers from the raw text
func parseMetaData(metaData string) *Request {
	request := &Request{}

	// Split header into individual lines
	lines := strings.Split(metaData, protocol.CRLF)
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

		headers[key] = append(headers[key], values...)
	}

	request.Headers = headers
	return request
}
