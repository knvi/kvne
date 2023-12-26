package get

import (
	"errors"
	"io"
	"time"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) error {
	if len(args) != 1 {
		return errors.New("ERR wrong number of arguments for 'get' command")
	}

	k := args[0]
	obj := storage.Get(k)
	if obj == nil {
		c.Write([]byte("$-1" + coder.CRLF))
		return nil
	}

	if obj.Expire != -1 && obj.Expire <= time.Now().UnixMilli() {
		c.Write([]byte("$-1" + coder.CRLF))
		return nil
	}

	c.Write(coder.Encode(obj.Value, false))

	return nil
}