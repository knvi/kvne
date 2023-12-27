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
		return config.RESP_TTL_KEY_NOT_EXIST
	}

	exp, isSet := storage.GetExpiration(obj)
	if !isSet {
		return config.RESP_TTL_NO_EXPIRE
	}

	remain := exp - time.Now().UnixMilli()

	if remain < 0 {
		return config.RESP_TTL_NO_EXPIRE
	}

	c.Write(coder.Encode(int64(remain/1000), false))
	return nil
}