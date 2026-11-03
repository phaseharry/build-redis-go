package resp

import (
	"fmt"
)

type Response struct {
	Content []byte
}

func NewResponse(dataType string, message []byte) (*Response, error) {
	res := &Response{}
	content := []byte{}

	dataTypePrefix, ok := dataTypePrefixByteMap[dataType]
	if !ok {
		return res, fmt.Errorf("Invalid response dataType %s", dataType)
	}

	content = append(content, dataTypePrefix)
	content = append(content, message...)
	content = append(content, []byte{CARRIAGE_RETURN, LINE_FEED}...)
	res.Content = content

	return res, nil
}
