package expire

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) error {
	if len(args) < 2 {
		return errors.New("ERR wrong number of arguments for 'expire' command")
	}

	k := args[0]
	ttl, sec := strconv.ParseInt(args[1], 10, 64)
	if sec != nil {
		return errors.New("ERR value is not an integer or out of range")
	}

	obj := storage.Get(k)
	if obj == nil {
		c.Write([]byte(":0\r\n"))
		return nil
	}

	obj.Expire = time.Now().UnixMilli() + ttl*1000
	c.Write([]byte(":1\r\n"))
	return nil
}