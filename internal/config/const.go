package config

import (
	"fmt"
)

var RESP_NIL = []byte("$-1\r\n")
var RESP_OK = []byte("+OK\r\n")
var RESP_ZERO = []byte(":0\r\n")
var RESP_ONE = []byte(":1\r\n")
var RESP_TTL_KEY_NOT_EXIST = []byte(":-2\r\n")
var RESP_TTL_NO_EXPIRE = []byte(":-1\r\n")

var AOF_FILE = "./kvne.aof"

func ErrWrongNumberOfArguments(cmd string) error {
    return fmt.Errorf("ERR wrong number of arguments for '%s' command", cmd)
}