package main

import (
	"fmt"
	"net"
)

// when a client has connected, hold onto that connection and repeated
// accept requests from them until connection has been terminated
func handleConnection(conn net.Conn) error {
	defer conn.Close()
	// let request message be a max of 1024 bytes for now
	req := make([]byte, 1024)
	for {
		bytesRead, err := conn.Read(req)
		if err != nil {
			fmt.Println("Error reading request: ", err.Error())
			break
		}
		fmt.Printf("read in request %d bytes\n", bytesRead)
		command := req[:bytesRead]
		fmt.Printf("command: %s", string(command))
		// doing exact match to redis RESP contract for now.
		if string(command) == "*1\r\n$4\r\nPING\r\n" {
			conn.Write([]byte("+PONG\r\n"))
			continue
		}
		conn.Write([]byte("+ERROR: unknown command\r\n"))
	}

	return nil
}
