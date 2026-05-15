package http

import (
	"errors"
	"strings"
)

func ParseRequest(data []byte) (*Request, error) {
	request := &Request{}

	lines := strings.Split(string(data), "\r\n")

	requestLine := strings.Split(lines[0], " ")
	if len(requestLine) != 3 {
		return nil, errors.New("more then three fields in request line")
	}

	request.Method = requestLine[0]
	request.Path = requestLine[1]
	request.Version = requestLine[2]

	headers := make(map[string]string)

	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("malformed header")
		}

		key := strings.ToLower(parts[0])
		value := parts[1]

		headers[key] = value
	}

	request.Headers = headers

	return request, nil
}
