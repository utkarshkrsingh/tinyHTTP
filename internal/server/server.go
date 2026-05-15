package server

import (
	"net"

	"github.com/utkarshkrsingh/tinyhttp/internal/handler"
	"github.com/utkarshkrsingh/tinyhttp/internal/http"
	"github.com/utkarshkrsingh/tinyhttp/internal/transport"
)

func HandleConnection(conn net.Conn, handler handler.Handler) {
	defer conn.Close()
	c := transport.NewConnection(conn)

	raw, err := c.ReadRequest()
	if err != nil {
		return
	}

	req, err := http.ParseRequest(raw)
	if err != nil {
		return
	}

	res := handler.ServeHTTP(req)

	err = http.WriteResponse(conn, res)
}
