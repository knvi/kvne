package ttl

import (
	"io"
	"time"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) []byte {
	if len(args) != 1 {
		return coder.Encode(config.ErrWrongNumberOfArguments("ttl"), false)
	}

	k := args[0]
	obj := storage.Get(k)
	if obj == nil {
		c.Write([]byte(":-2\r\n"))
		return nil
	}

	if obj.Expire == -1 {
		c.Write([]byte(":-1\r\n"))
		return nil
	}

	remain := obj.Expire - time.Now().UnixMilli()

	if remain < 0 {
		c.Write([]byte(":-2\r\n"))
		return nil
	}

	c.Write(coder.Encode(int64(remain/1000), false))
	return nil
}