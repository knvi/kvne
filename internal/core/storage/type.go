package storage

import (
	"strconv"

	"github.com/knvi/kvne/internal/config"
)

func DeduceType(v string) (uint8, uint8) {
	o_type := config.OBJ_TYPE_STRING
	if _, err := strconv.ParseInt(v, 10, 64); err == nil {
		o_type = config.OBJ_ENCODING_INT
	}

	return o_type, config.OBJ_ENCODING_RAW
}

func GetType(t uint8) uint8 {
	return t & 0b11110000
}

func GetEncoding(t uint8) uint8 {
	return t & 0b00001111
}

func AssertType(t1 uint8, t2 uint8) bool {
	return GetType(t1) == t2
}

func AssertEncoding(t1 uint8, t2 uint8) bool {
	return GetEncoding(t1) == t2
}