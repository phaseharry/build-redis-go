package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		// let request message be a max of 32 bytes for now
		req := make([]byte, 32)
		bytesRead, err := conn.Read(req)
		if err != nil {
			fmt.Println("Error reading request: ", err.Error())
			continue
		}
		fmt.Printf("read in request %d bytes\n", bytesRead)
		command := req[:bytesRead]
		fmt.Printf("command: %s", string(command))
		if string(command) == "PING" {
			conn.Write([]byte("+PONG\r\n"))
			continue
		}
		conn.Write([]byte("+ERROR: unknown command\r\n"))
	}
}
