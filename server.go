package main

import (
	"log"
	"net"

	"github.com/utkarshkrsingh/tinyHTTP/httpcore"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	req, err := httpcore.ReadRequest(conn)

	res := httpcore.GenerateResponse(req, err)

	_, err = conn.Write([]byte(res))
	if err != nil {
		log.Printf("unable to write response: %v\n", err)
	}
}
