package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/utkarshkrsingh/tinyhttp/internal/handler"
	"github.com/utkarshkrsingh/tinyhttp/internal/server"
)

func main() {
	addr := flag.String("addr", "4221", "HTTP network address")
	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:"+*addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer listener.Close()

	router := handler.NewRouter()
	router.Register("GET", "/", &handler.HomeHandler{})
	router.Register("GET", "/hello", &handler.HelloHandler{})

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go server.HandleConnection(conn, router)
	}
}
