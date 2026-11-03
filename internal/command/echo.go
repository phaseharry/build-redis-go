package command

import (
	"strings"

	"github.com/codecrafters-io/redis-starter-go/internal/resp"
)

func echo(arguments []string) (*resp.Response, error) {
	response, err := resp.NewResponse(
		resp.SimpleStrings,
		[]byte(strings.Join(arguments, " ")),
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
