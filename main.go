package main

import (
	"log"
	"net"
)

// Define the port the server will listen on
var (
	port = "4220"
)

func main() {
	// Start a TCP server listening on localhost:4220
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		// If the server fails to start, log the error and exit
		log.Fatalf("error : %v\n", err)
	}
	// Ensure the listener is closed when the program exists
	defer listener.Close()

	// Infinite loop to keep accepting incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			// If there's an error accepting a connection, log it and continue
			log.Printf("error : %v\n", err)
			continue
		}

		// Handle each client connection in a separate goroutine (concurrently)
		go handleClient(conn)
	}
}
