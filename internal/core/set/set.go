package set

import (
	"errors"
	"io"
	"strconv"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) []byte {
	if len(args) < 2 || len(args) == 3 || len(args) > 4 {
		return coder.Encode(config.ErrWrongNumberOfArguments("set"), false)
	}

	var k, v string
	var expire int64 = -1

	k, v = args[0], args[1]
	o_type, o_enc := storage.DeduceType(v)
	if len(args) > 2 {
		ttl, err := strconv.ParseInt(args[3], 10, 64)
		if err != nil {
			return coder.Encode(errors.New("ERR value is not an integer or out of range"), false)
		}
		expire = ttl * 1000
	}

	storage.Put(k, storage.NewObject(v, int64(expire), o_type, o_enc))
	c.Write(coder.Encode("OK", true))
	return nil
}