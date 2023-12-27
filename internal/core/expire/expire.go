package expire

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) []byte {
	if len(args) < 2 {
		return coder.Encode(config.ErrWrongNumberOfArguments("expire"), false)
	}

	k := args[0]
	ttl, sec := strconv.ParseInt(args[1], 10, 64)
	if sec != nil {
		return coder.Encode(errors.New("ERR value is not an integer or out of range"), false)
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