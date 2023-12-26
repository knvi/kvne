package set

import (
	"errors"
	"io"
	"strconv"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) error {
	if len(args) < 2 || len(args) == 3 || len(args) > 4 {
		return errors.New("ERR wrong number of arguments for 'get' command");
	}

	var k, v string
	var expire int64 = -1

	k, v = args[0], args[1]
	if len(args) > 2 {
		ttl, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			return err
		}
		expire = ttl * 1000
	}

	storage.Put(k, storage.Add(v, expire))
	c.Write(coder.Encode("OK", true))
	return nil
}