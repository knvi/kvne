package coder

import "fmt"

var CRLF = "\r\n"

func Encode(value interface{}, isBase bool) []byte {
	switch v := value.(type) {
	case string:
		if isBase {
			return []byte(fmt.Sprintf("+%s%s", v, CRLF))
		} 
		return []byte(fmt.Sprintf("$%d%s%s%s", len(v), CRLF, v, CRLF))
	case int: 
		return []byte(fmt.Sprintf(":%d%s", v, CRLF))
	default:
		return []byte(fmt.Sprintf("$-1%s", CRLF))
	}
}