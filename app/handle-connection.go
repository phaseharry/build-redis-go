package main

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/redis-starter-go/internal/command"
	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

// when a client has connected, hold onto that connection and repeated
// accept requests from them until connection has been terminated
func handleConnection(conn net.Conn) error {
	defer conn.Close()
	// let request message be a max of 1024 bytes for now
	reqBuffer := make([]byte, 1024)
	for {
		bytesRead, err := conn.Read(reqBuffer)
		if err != nil {
			fmt.Println("Error reading request bytes: ", err.Error())
			break
		}
		fmt.Printf("read in request %d bytes\n", bytesRead)
		req, err := resp.NewRequest(reqBuffer[:bytesRead])
		if err != nil {
			fmt.Println("Error creating request: ", err.Error())
			break
		}

		res, err := command.RunCommand(req)
		if err != nil {
			errMessage := fmt.Sprintf("Error running command: %s, %s", req.Command, err.Error())
			fmt.Println(errMessage)
			fmt.Fprintf(conn, "+ERROR: %s\r\n", errMessage)
			break
		}

		conn.Write(res.Content)
	}

	return nil
}
