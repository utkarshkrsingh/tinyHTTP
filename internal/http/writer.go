package http

import (
	"fmt"
	"net"
	"strings"
)

func WriteResponse(conn net.Conn, res *Response) error {
	var resMsg strings.Builder
	fmt.Fprintf(&resMsg, `%s %d %s\r\n`,
		res.Version,
		res.StatusCode,
		res.StatusText,
		)

	for key, value := range res.Headers {
		fmt.Fprintf(&resMsg, "%s:%s\r\n", key, value)
	}

	fmt.Fprintf(&resMsg, "\r\n")

	if len(res.Body) > 0 {
		resMsg.Write(res.Body)
	}
	_, err := conn.Write([]byte(resMsg.String()))
	return err
}
