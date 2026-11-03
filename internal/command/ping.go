package command

import "github.com/codecrafters-io/redis-starter-go/internal/resp"

func ping() (*resp.Response, error) {
	response, err := resp.NewResponse(
		resp.SimpleStrings,
		[]byte("PONG"),
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}
