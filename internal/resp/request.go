package resp

import (
	"fmt"
	"strconv"
)

type Request struct {
	DataType  string  // first byte tells us the request type
	Command   string  // first arguement (and maybe second argument) once we split the rawBytes by CRLF is the redis command called
	Arguments []Token // after first (and maybe second), the rest are arguments that that passed to the command to run
}

type Token struct {
	DataType string
	Length   int
	Value    string
}

func NewRequest(rawBytes []byte) (Request, error) {
	req := Request{}
	lines, err := parseLines(rawBytes)
	if err != nil {
		return req, err
	}

	dataType := ""
	respHeader := lines[0]
	if respHeader == "$-1" {
		dataType, _ = prefixByteDataTypeMap[respHeader]
	} else {
		firstByte := string(respHeader[0])
		_, ok := prefixByteDataTypeMap[firstByte]
		if !ok {
			return req, fmt.Errorf("RESP header's first byte value: %s does not match any supported data type.", firstByte)
		}
		dataType = prefixByteDataTypeMap[firstByte]
	}

	req.DataType = dataType
	req.Command = lines[1]
	// assign the rest of the lines after the first 2 lines to Arguments

	return req, nil
}

// https://redis.io/docs/latest/develop/reference/protocol-spec/#resp-protocol-description
// *2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n
//
// * - first byte = request type
// 2 - length / number of arguments  (2 arguments)
func parseLines(rawBytes []byte) ([]Token, error) {
	// convert incoming *2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n -> [*2, $4, ECHO, $s, hey]
	// remaining := bytes.Split(rawBytes, []byte{CARRIAGE_RETURN, LINE_FEED})
	remaining := rawBytes

	// assuming there should always be at least 2 lines.
	// the resp header and a command line
	if len(remaining) < 2 {
		return nil, fmt.Errorf("invalid request: not enough lines")
	}

	tokens := []Token{}
	// *1\r\n$4\r\nPING\r\n

	for len(remaining) > 0 {
		token := Token{}
		offset := 0
		dataTypeByte := remaining[offset]

		// handling BulkStrings & NullBulkStrings which has a size value after the '$' byte
		if dataTypeByte == '$' {
			// if dataTypeByte is '$', check the next 2 bytes to see if it's "-1" because "$-1" is the NullBulkStrings
			if len(remaining) >= 3 && string(remaining[offset:3]) == "$-1" {
				offset += 3
				dataType, _ := prefixByteDataTypeMap[string(remaining[offset:3])]
				token.DataType = dataType
			} else {
				if !isDigit(remaining[offset]) {
					return nil, fmt.Errorf("Expected digit byte following $ to indicate the length of the value")
				}
				startingDigitOffset := offset
				for offset+1 < len(remaining) && isDigit(remaining[offset+1]) {
					offset += 1
				}
				length, _ := strconv.Atoi(string(remaining[startingDigitOffset:offset]))
				fmt.Printf("%d - %d, length: %d", startingDigitOffset, offset, length)
				token.Length = length
			}
		} else {
			offset += 1
			dataType, ok := prefixByteDataTypeMap[string(dataTypeByte)]
			if !ok {
				return nil, fmt.Errorf("Invalid data type byte: %s", dataType)
			}
			token.DataType = dataType
		}

		carriageReturn := remaining[offset]
		if carriageReturn != CARRIAGE_RETURN {
			return nil, fmt.Errorf("Invalid byte: %s, expected a carriage return", carriageReturn)
		}
		offset += 1

		lineFeed := remaining[offset]
		if lineFeed != LINE_FEED {
			return nil, fmt.Errorf("Invalid byte: %s, expected a line feed", lineFeed)
		}
		offset += 1

		// for array type, we consider that there are no more values
		if dataTypeByte == '*' {

		}

		remaining = remaining[offset:]

		value := []byte{}

		tokens = append(tokens, token)
	}
	return tokens, nil
}

func isDigit(char byte) bool {
	return char > '0' && char < '9'
}
