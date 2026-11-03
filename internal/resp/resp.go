package resp

type DataType = string

const (
	SimpleStrings   DataType = "SimpleStrings"
	SimpleErrors    DataType = "SimpleErrors"
	Integer         DataType = "Integer"
	BulkStrings     DataType = "BulkStrings"
	Arrays          DataType = "Arrays"
	NullBulkStrings DataType = "NullBulkStrings"
)

var dataTypePrefixByteMap map[DataType]byte = map[DataType]byte{
	"SimpleStrings":   '+',
	"SimpleErrors":    '-',
	"Integer":         ':',
	"BulkStrings":     '$',
	"Arrays":          '*',
	"NullBulkStrings": '$',
}

var prefixByteDataTypeMap map[string]DataType = map[string]DataType{
	"+":   "SimpleStrings",
	"-":   "SimpleErrors",
	":":   "Integer",
	"$":   "BulkStrings",
	"*":   "Arrays",
	"$-1": "NullBulkStrings",
}

// CRLF (\r\n) is always used to separate each part of a request and response
// serves as a the RESP protocol's terminator and is always there between each component of a request and response
const CARRIAGE_RETURN byte = '\r'
const LINE_FEED byte = '\n'
