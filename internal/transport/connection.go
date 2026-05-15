package transport

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

type Connection struct {
	conn   net.Conn
	reader *bufio.Reader
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		conn: conn,
		reader: bufio.NewReader(conn),
	}
}

func (c *Connection) ReadRequest() ([]byte, error) {
	var data []byte

	for {
		line, err := c.reader.ReadBytes('\n')
		if err != nil {
			return nil, err
		}

		data =  append(data, line...)

		if bytes.Contains(data, []byte("\r\n\r\n")) {
			break
		}
	}

	fmt.Println(string(data))

	return data, nil
}
