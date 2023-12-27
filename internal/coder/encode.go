package coder

import (
	"bytes"
	"fmt"
)

var CRLF = "\r\n"

func encodeString(value string) []byte {
	return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value))
}

func Encode(value interface{}, isBase bool) []byte {
	switch v := value.(type) {
	case string:
		if isBase {
			return []byte(fmt.Sprintf("+%s%s", v, CRLF))
		} 
		return []byte(fmt.Sprintf("$%d%s%s%s", len(v), CRLF, v, CRLF))
	case int64, int32, int16, int8, int: 
		return []byte(fmt.Sprintf(":%d%s", v, CRLF))
	case error:
		return []byte(fmt.Sprintf("-%s%s", v, CRLF))
	case []string:
		var res []byte
		buf := bytes.NewBuffer(res)

		for _, s := range value.([]string) {
			buf.Write(encodeString(s))
		}

		return []byte(fmt.Sprintf("*%d%s%s", len(v), CRLF, buf.String()))
	default:
		return []byte(fmt.Sprintf("$-1%s", CRLF))
	}
}