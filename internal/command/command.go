package command

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

func RunCommand(req resp.Request) (*resp.Response, error) {
	command, arguments := req.Command, req.Arguments
	switch strings.ToLower(command) {
	case "echo":
		return echo(arguments)
	case "ping":
		return ping()
	default:
		return nil, fmt.Errorf("Unknown command: %s", command)
	}
}
